//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1alpha1 "k8s.io/kube-controller-manager/config/v1alpha1"
	config "k8s.io/kubernetes/pkg/controller/volume/persistentvolume/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*configv1alpha1.GroupResource)(nil), (*v1.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_GroupResource_To_v1_GroupResource(a.(*configv1alpha1.GroupResource), b.(*v1.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.GroupResource)(nil), (*configv1alpha1.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GroupResource_To_v1alpha1_GroupResource(a.(*v1.GroupResource), b.(*configv1alpha1.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*configv1alpha1.PersistentVolumeRecyclerConfiguration)(nil), (*config.PersistentVolumeRecyclerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration(a.(*configv1alpha1.PersistentVolumeRecyclerConfiguration), b.(*config.PersistentVolumeRecyclerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PersistentVolumeRecyclerConfiguration)(nil), (*configv1alpha1.PersistentVolumeRecyclerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(a.(*config.PersistentVolumeRecyclerConfiguration), b.(*configv1alpha1.PersistentVolumeRecyclerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*configv1alpha1.VolumeConfiguration)(nil), (*config.VolumeConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration(a.(*configv1alpha1.VolumeConfiguration), b.(*config.VolumeConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.VolumeConfiguration)(nil), (*configv1alpha1.VolumeConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(a.(*config.VolumeConfiguration), b.(*configv1alpha1.VolumeConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.PersistentVolumeBinderControllerConfiguration)(nil), (*configv1alpha1.PersistentVolumeBinderControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PersistentVolumeBinderControllerConfiguration_To_v1alpha1_PersistentVolumeBinderControllerConfiguration(a.(*config.PersistentVolumeBinderControllerConfiguration), b.(*configv1alpha1.PersistentVolumeBinderControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*configv1alpha1.PersistentVolumeBinderControllerConfiguration)(nil), (*config.PersistentVolumeBinderControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PersistentVolumeBinderControllerConfiguration_To_config_PersistentVolumeBinderControllerConfiguration(a.(*configv1alpha1.PersistentVolumeBinderControllerConfiguration), b.(*config.PersistentVolumeBinderControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_GroupResource_To_v1_GroupResource(in *configv1alpha1.GroupResource, out *v1.GroupResource, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_v1alpha1_GroupResource_To_v1_GroupResource is an autogenerated conversion function.
func Convert_v1alpha1_GroupResource_To_v1_GroupResource(in *configv1alpha1.GroupResource, out *v1.GroupResource, s conversion.Scope) error {
	return autoConvert_v1alpha1_GroupResource_To_v1_GroupResource(in, out, s)
}

func autoConvert_v1_GroupResource_To_v1alpha1_GroupResource(in *v1.GroupResource, out *configv1alpha1.GroupResource, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_v1_GroupResource_To_v1alpha1_GroupResource is an autogenerated conversion function.
func Convert_v1_GroupResource_To_v1alpha1_GroupResource(in *v1.GroupResource, out *configv1alpha1.GroupResource, s conversion.Scope) error {
	return autoConvert_v1_GroupResource_To_v1alpha1_GroupResource(in, out, s)
}

func autoConvert_v1alpha1_PersistentVolumeBinderControllerConfiguration_To_config_PersistentVolumeBinderControllerConfiguration(in *configv1alpha1.PersistentVolumeBinderControllerConfiguration, out *config.PersistentVolumeBinderControllerConfiguration, s conversion.Scope) error {
	out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
	if err := Convert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration(&in.VolumeConfiguration, &out.VolumeConfiguration, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_config_PersistentVolumeBinderControllerConfiguration_To_v1alpha1_PersistentVolumeBinderControllerConfiguration(in *config.PersistentVolumeBinderControllerConfiguration, out *configv1alpha1.PersistentVolumeBinderControllerConfiguration, s conversion.Scope) error {
	out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
	if err := Convert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(&in.VolumeConfiguration, &out.VolumeConfiguration, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration(in *configv1alpha1.PersistentVolumeRecyclerConfiguration, out *config.PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
	out.MaximumRetry = in.MaximumRetry
	out.MinimumTimeoutNFS = in.MinimumTimeoutNFS
	out.PodTemplateFilePathNFS = in.PodTemplateFilePathNFS
	out.IncrementTimeoutNFS = in.IncrementTimeoutNFS
	out.PodTemplateFilePathHostPath = in.PodTemplateFilePathHostPath
	out.MinimumTimeoutHostPath = in.MinimumTimeoutHostPath
	out.IncrementTimeoutHostPath = in.IncrementTimeoutHostPath
	return nil
}

// Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration(in *configv1alpha1.PersistentVolumeRecyclerConfiguration, out *config.PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration(in, out, s)
}

func autoConvert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(in *config.PersistentVolumeRecyclerConfiguration, out *configv1alpha1.PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
	out.MaximumRetry = in.MaximumRetry
	out.MinimumTimeoutNFS = in.MinimumTimeoutNFS
	out.PodTemplateFilePathNFS = in.PodTemplateFilePathNFS
	out.IncrementTimeoutNFS = in.IncrementTimeoutNFS
	out.PodTemplateFilePathHostPath = in.PodTemplateFilePathHostPath
	out.MinimumTimeoutHostPath = in.MinimumTimeoutHostPath
	out.IncrementTimeoutHostPath = in.IncrementTimeoutHostPath
	return nil
}

// Convert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration is an autogenerated conversion function.
func Convert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(in *config.PersistentVolumeRecyclerConfiguration, out *configv1alpha1.PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
	return autoConvert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration(in *configv1alpha1.VolumeConfiguration, out *config.VolumeConfiguration, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableHostPathProvisioning, &out.EnableHostPathProvisioning, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableDynamicProvisioning, &out.EnableDynamicProvisioning, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_config_PersistentVolumeRecyclerConfiguration(&in.PersistentVolumeRecyclerConfiguration, &out.PersistentVolumeRecyclerConfiguration, s); err != nil {
		return err
	}
	out.FlexVolumePluginDir = in.FlexVolumePluginDir
	return nil
}

// Convert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration(in *configv1alpha1.VolumeConfiguration, out *config.VolumeConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeConfiguration_To_config_VolumeConfiguration(in, out, s)
}

func autoConvert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(in *config.VolumeConfiguration, out *configv1alpha1.VolumeConfiguration, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableHostPathProvisioning, &out.EnableHostPathProvisioning, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableDynamicProvisioning, &out.EnableDynamicProvisioning, s); err != nil {
		return err
	}
	if err := Convert_config_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(&in.PersistentVolumeRecyclerConfiguration, &out.PersistentVolumeRecyclerConfiguration, s); err != nil {
		return err
	}
	out.FlexVolumePluginDir = in.FlexVolumePluginDir
	return nil
}

// Convert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration is an autogenerated conversion function.
func Convert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(in *config.VolumeConfiguration, out *configv1alpha1.VolumeConfiguration, s conversion.Scope) error {
	return autoConvert_config_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(in, out, s)
}
