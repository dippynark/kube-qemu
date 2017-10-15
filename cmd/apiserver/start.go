/*
Copyright 2017 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io"
	"net"

	"github.com/spf13/cobra"

	"github.com/dippynark/kube-qemu/pkg/admission/hypervisorinitializer"
	"github.com/dippynark/kube-qemu/pkg/admission/plugin/banvirtualmachine"
	"github.com/dippynark/kube-qemu/pkg/apiserver"
	"k8s.io/api/rbac/v1alpha1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/client-go/informers"
	"k8s.io/metrics/pkg/client/clientset_generated/clientset"
)

const defaultEtcdPathPrefix = "/registry/hypervisor.lukeaddison.co.uk"

type HypervisorServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions
	Admission          *genericoptions.AdmissionOptions

	StdOut io.Writer
	StdErr io.Writer
}

func NewHypervisorServerOptions(out, errOut io.Writer) *HypervisorServerOptions {
	o := &HypervisorServerOptions{
		RecommendedOptions: genericoptions.NewRecommendedOptions(defaultEtcdPathPrefix, apiserver.Scheme, apiserver.Codecs.LegacyCodec(v1alpha1.SchemeGroupVersion)),
		Admission:          genericoptions.NewAdmissionOptions(),

		StdOut: out,
		StdErr: errOut,
	}

	return o
}

// NewCommandStartMaster provides a CLI handler for 'start master' command
func NewCommandStartHypervisorServer(out, errOut io.Writer, stopCh <-chan struct{}) *cobra.Command {
	o := NewHypervisorServerOptions(out, errOut)

	cmd := &cobra.Command{
		Short: "Launch a hypervisor API server",
		Long:  "Launch a hypervisor API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}
			if err := o.RunHypervisorServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.RecommendedOptions.AddFlags(flags)
	o.Admission.AddFlags(flags)

	return cmd
}

func (o HypervisorServerOptions) Validate(args []string) error {
	errors := []error{}
	errors = append(errors, o.RecommendedOptions.Validate()...)
	errors = append(errors, o.Admission.Validate()...)
	return utilerrors.NewAggregate(errors)
}

func (o *HypervisorServerOptions) Complete() error {
	return nil
}

func (o HypervisorServerOptions) Config() (*apiserver.Config, error) {
	// register admission plugins
	banvirtualmachine.Register(o.Admission.Plugins)

	// TODO have a "real" external address
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewRecommendedConfig(apiserver.Codecs)
	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	client, err := clientset.NewForConfig(serverConfig.LoopbackClientConfig)
	if err != nil {
		return nil, err
	}
	informerFactory := informers.NewSharedInformerFactory(client, serverConfig.LoopbackClientConfig.Timeout)
	admissionInitializer, err := hypervisorinitializer.New(informerFactory)
	if err != nil {
		return nil, err
	}

	if err := o.Admission.ApplyTo(&serverConfig.Config, serverConfig.SharedInformerFactory, admissionInitializer); err != nil {
		return nil, err
	}

	config := &apiserver.Config{
		GenericConfig: serverConfig,
		ExtraConfig:   apiserver.ExtraConfig{},
	}
	return config, nil
}

func (o HypervisorServerOptions) RunHypervisorServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}

	server.GenericAPIServer.AddPostStartHook("start-sample-server-informers", func(context genericapiserver.PostStartHookContext) error {
		config.GenericConfig.SharedInformerFactory.Start(context.StopCh)
		return nil
	})

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}
