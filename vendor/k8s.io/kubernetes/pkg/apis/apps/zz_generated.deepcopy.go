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

package apps

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_PetSet, InType: reflect.TypeOf(&PetSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_PetSetList, InType: reflect.TypeOf(&PetSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_PetSetSpec, InType: reflect.TypeOf(&PetSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_PetSetStatus, InType: reflect.TypeOf(&PetSetStatus{})},
	)
}

func DeepCopy_apps_PetSet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PetSet)
		out := out.(*PetSet)
		out.TypeMeta = in.TypeMeta
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_apps_PetSetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_apps_PetSetStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_apps_PetSetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PetSetList)
		out := out.(*PetSetList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PetSet, len(*in))
			for i := range *in {
				if err := DeepCopy_apps_PetSet(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_apps_PetSetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PetSetSpec)
		out := out.(*PetSetSpec)
		out.Replicas = in.Replicas
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = new(unversioned.LabelSelector)
			if err := unversioned.DeepCopy_unversioned_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Selector = nil
		}
		if err := api.DeepCopy_api_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if in.VolumeClaimTemplates != nil {
			in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
			*out = make([]api.PersistentVolumeClaim, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_PersistentVolumeClaim(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.VolumeClaimTemplates = nil
		}
		out.ServiceName = in.ServiceName
		return nil
	}
}

func DeepCopy_apps_PetSetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PetSetStatus)
		out := out.(*PetSetStatus)
		if in.ObservedGeneration != nil {
			in, out := &in.ObservedGeneration, &out.ObservedGeneration
			*out = new(int64)
			**out = **in
		} else {
			out.ObservedGeneration = nil
		}
		out.Replicas = in.Replicas
		return nil
	}
}
