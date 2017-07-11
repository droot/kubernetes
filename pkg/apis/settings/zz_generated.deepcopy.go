// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package settings

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_settings_ContainerSelector, InType: reflect.TypeOf(&ContainerSelector{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_settings_PodPreset, InType: reflect.TypeOf(&PodPreset{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_settings_PodPresetList, InType: reflect.TypeOf(&PodPresetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_settings_PodPresetSpec, InType: reflect.TypeOf(&PodPresetSpec{})},
	)
}

// DeepCopy_settings_ContainerSelector is an autogenerated deepcopy function.
func DeepCopy_settings_ContainerSelector(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ContainerSelector)
		out := out.(*ContainerSelector)
		*out = *in
		if in.MatchNames != nil {
			in, out := &in.MatchNames, &out.MatchNames
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_settings_PodPreset is an autogenerated deepcopy function.
func DeepCopy_settings_PodPreset(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodPreset)
		out := out.(*PodPreset)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_settings_PodPresetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_settings_PodPresetList is an autogenerated deepcopy function.
func DeepCopy_settings_PodPresetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodPresetList)
		out := out.(*PodPresetList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PodPreset, len(*in))
			for i := range *in {
				if err := DeepCopy_settings_PodPreset(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_settings_PodPresetSpec is an autogenerated deepcopy function.
func DeepCopy_settings_PodPresetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodPresetSpec)
		out := out.(*PodPresetSpec)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Selector); err != nil {
			return err
		} else {
			out.Selector = *newVal.(*v1.LabelSelector)
		}
		if in.AppContainerSelector != nil {
			in, out := &in.AppContainerSelector, &out.AppContainerSelector
			*out = new(ContainerSelector)
			if err := DeepCopy_settings_ContainerSelector(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api.EnvVar, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.EnvFrom != nil {
			in, out := &in.EnvFrom, &out.EnvFrom
			*out = make([]api.EnvFromSource, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_EnvFromSource(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Volumes != nil {
			in, out := &in.Volumes, &out.Volumes
			*out = make([]api.Volume, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_Volume(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.VolumeMounts != nil {
			in, out := &in.VolumeMounts, &out.VolumeMounts
			*out = make([]api.VolumeMount, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}
