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

package v1beta1

import (
	unsafe "unsafe"

	v1 "k8s.io/api/authorization/v1"
	v1beta1 "k8s.io/api/authorization/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	authorization "k8s.io/kubernetes/pkg/apis/authorization"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1beta1.LocalSubjectAccessReview)(nil), (*authorization.LocalSubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(a.(*v1beta1.LocalSubjectAccessReview), b.(*authorization.LocalSubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.LocalSubjectAccessReview)(nil), (*v1beta1.LocalSubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_LocalSubjectAccessReview_To_v1beta1_LocalSubjectAccessReview(a.(*authorization.LocalSubjectAccessReview), b.(*v1beta1.LocalSubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.NonResourceAttributes)(nil), (*authorization.NonResourceAttributes)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NonResourceAttributes_To_authorization_NonResourceAttributes(a.(*v1beta1.NonResourceAttributes), b.(*authorization.NonResourceAttributes), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.NonResourceAttributes)(nil), (*v1beta1.NonResourceAttributes)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_NonResourceAttributes_To_v1beta1_NonResourceAttributes(a.(*authorization.NonResourceAttributes), b.(*v1beta1.NonResourceAttributes), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.NonResourceRule)(nil), (*authorization.NonResourceRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NonResourceRule_To_authorization_NonResourceRule(a.(*v1beta1.NonResourceRule), b.(*authorization.NonResourceRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.NonResourceRule)(nil), (*v1beta1.NonResourceRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_NonResourceRule_To_v1beta1_NonResourceRule(a.(*authorization.NonResourceRule), b.(*v1beta1.NonResourceRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ResourceAttributes)(nil), (*authorization.ResourceAttributes)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ResourceAttributes_To_authorization_ResourceAttributes(a.(*v1beta1.ResourceAttributes), b.(*authorization.ResourceAttributes), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.ResourceAttributes)(nil), (*v1beta1.ResourceAttributes)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ResourceAttributes_To_v1beta1_ResourceAttributes(a.(*authorization.ResourceAttributes), b.(*v1beta1.ResourceAttributes), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ResourceRule)(nil), (*authorization.ResourceRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ResourceRule_To_authorization_ResourceRule(a.(*v1beta1.ResourceRule), b.(*authorization.ResourceRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.ResourceRule)(nil), (*v1beta1.ResourceRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_ResourceRule_To_v1beta1_ResourceRule(a.(*authorization.ResourceRule), b.(*v1beta1.ResourceRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SelfSubjectAccessReview)(nil), (*authorization.SelfSubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(a.(*v1beta1.SelfSubjectAccessReview), b.(*authorization.SelfSubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SelfSubjectAccessReview)(nil), (*v1beta1.SelfSubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SelfSubjectAccessReview_To_v1beta1_SelfSubjectAccessReview(a.(*authorization.SelfSubjectAccessReview), b.(*v1beta1.SelfSubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SelfSubjectAccessReviewSpec)(nil), (*authorization.SelfSubjectAccessReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(a.(*v1beta1.SelfSubjectAccessReviewSpec), b.(*authorization.SelfSubjectAccessReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SelfSubjectAccessReviewSpec)(nil), (*v1beta1.SelfSubjectAccessReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec(a.(*authorization.SelfSubjectAccessReviewSpec), b.(*v1beta1.SelfSubjectAccessReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SelfSubjectRulesReview)(nil), (*authorization.SelfSubjectRulesReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(a.(*v1beta1.SelfSubjectRulesReview), b.(*authorization.SelfSubjectRulesReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SelfSubjectRulesReview)(nil), (*v1beta1.SelfSubjectRulesReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SelfSubjectRulesReview_To_v1beta1_SelfSubjectRulesReview(a.(*authorization.SelfSubjectRulesReview), b.(*v1beta1.SelfSubjectRulesReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SelfSubjectRulesReviewSpec)(nil), (*authorization.SelfSubjectRulesReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(a.(*v1beta1.SelfSubjectRulesReviewSpec), b.(*authorization.SelfSubjectRulesReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SelfSubjectRulesReviewSpec)(nil), (*v1beta1.SelfSubjectRulesReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SelfSubjectRulesReviewSpec_To_v1beta1_SelfSubjectRulesReviewSpec(a.(*authorization.SelfSubjectRulesReviewSpec), b.(*v1beta1.SelfSubjectRulesReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SubjectAccessReview)(nil), (*authorization.SubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SubjectAccessReview_To_authorization_SubjectAccessReview(a.(*v1beta1.SubjectAccessReview), b.(*authorization.SubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SubjectAccessReview)(nil), (*v1beta1.SubjectAccessReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SubjectAccessReview_To_v1beta1_SubjectAccessReview(a.(*authorization.SubjectAccessReview), b.(*v1beta1.SubjectAccessReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SubjectAccessReviewSpec)(nil), (*authorization.SubjectAccessReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(a.(*v1beta1.SubjectAccessReviewSpec), b.(*authorization.SubjectAccessReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SubjectAccessReviewSpec)(nil), (*v1beta1.SubjectAccessReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(a.(*authorization.SubjectAccessReviewSpec), b.(*v1beta1.SubjectAccessReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SubjectAccessReviewStatus)(nil), (*authorization.SubjectAccessReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(a.(*v1beta1.SubjectAccessReviewStatus), b.(*authorization.SubjectAccessReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SubjectAccessReviewStatus)(nil), (*v1beta1.SubjectAccessReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(a.(*authorization.SubjectAccessReviewStatus), b.(*v1beta1.SubjectAccessReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SubjectRulesReviewStatus)(nil), (*authorization.SubjectRulesReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(a.(*v1beta1.SubjectRulesReviewStatus), b.(*authorization.SubjectRulesReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authorization.SubjectRulesReviewStatus)(nil), (*v1beta1.SubjectRulesReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authorization_SubjectRulesReviewStatus_To_v1beta1_SubjectRulesReviewStatus(a.(*authorization.SubjectRulesReviewStatus), b.(*v1beta1.SubjectRulesReviewStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in *v1beta1.LocalSubjectAccessReview, out *authorization.LocalSubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview is an autogenerated conversion function.
func Convert_v1beta1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in *v1beta1.LocalSubjectAccessReview, out *authorization.LocalSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1beta1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in, out, s)
}

func autoConvert_authorization_LocalSubjectAccessReview_To_v1beta1_LocalSubjectAccessReview(in *authorization.LocalSubjectAccessReview, out *v1beta1.LocalSubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authorization_LocalSubjectAccessReview_To_v1beta1_LocalSubjectAccessReview is an autogenerated conversion function.
func Convert_authorization_LocalSubjectAccessReview_To_v1beta1_LocalSubjectAccessReview(in *authorization.LocalSubjectAccessReview, out *v1beta1.LocalSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_LocalSubjectAccessReview_To_v1beta1_LocalSubjectAccessReview(in, out, s)
}

func autoConvert_v1beta1_NonResourceAttributes_To_authorization_NonResourceAttributes(in *v1beta1.NonResourceAttributes, out *authorization.NonResourceAttributes, s conversion.Scope) error {
	out.Path = in.Path
	out.Verb = in.Verb
	return nil
}

// Convert_v1beta1_NonResourceAttributes_To_authorization_NonResourceAttributes is an autogenerated conversion function.
func Convert_v1beta1_NonResourceAttributes_To_authorization_NonResourceAttributes(in *v1beta1.NonResourceAttributes, out *authorization.NonResourceAttributes, s conversion.Scope) error {
	return autoConvert_v1beta1_NonResourceAttributes_To_authorization_NonResourceAttributes(in, out, s)
}

func autoConvert_authorization_NonResourceAttributes_To_v1beta1_NonResourceAttributes(in *authorization.NonResourceAttributes, out *v1beta1.NonResourceAttributes, s conversion.Scope) error {
	out.Path = in.Path
	out.Verb = in.Verb
	return nil
}

// Convert_authorization_NonResourceAttributes_To_v1beta1_NonResourceAttributes is an autogenerated conversion function.
func Convert_authorization_NonResourceAttributes_To_v1beta1_NonResourceAttributes(in *authorization.NonResourceAttributes, out *v1beta1.NonResourceAttributes, s conversion.Scope) error {
	return autoConvert_authorization_NonResourceAttributes_To_v1beta1_NonResourceAttributes(in, out, s)
}

func autoConvert_v1beta1_NonResourceRule_To_authorization_NonResourceRule(in *v1beta1.NonResourceRule, out *authorization.NonResourceRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_v1beta1_NonResourceRule_To_authorization_NonResourceRule is an autogenerated conversion function.
func Convert_v1beta1_NonResourceRule_To_authorization_NonResourceRule(in *v1beta1.NonResourceRule, out *authorization.NonResourceRule, s conversion.Scope) error {
	return autoConvert_v1beta1_NonResourceRule_To_authorization_NonResourceRule(in, out, s)
}

func autoConvert_authorization_NonResourceRule_To_v1beta1_NonResourceRule(in *authorization.NonResourceRule, out *v1beta1.NonResourceRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_authorization_NonResourceRule_To_v1beta1_NonResourceRule is an autogenerated conversion function.
func Convert_authorization_NonResourceRule_To_v1beta1_NonResourceRule(in *authorization.NonResourceRule, out *v1beta1.NonResourceRule, s conversion.Scope) error {
	return autoConvert_authorization_NonResourceRule_To_v1beta1_NonResourceRule(in, out, s)
}

func autoConvert_v1beta1_ResourceAttributes_To_authorization_ResourceAttributes(in *v1beta1.ResourceAttributes, out *authorization.ResourceAttributes, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Verb = in.Verb
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	out.Subresource = in.Subresource
	out.Name = in.Name
	out.FieldSelector = (*authorization.FieldSelectorAttributes)(unsafe.Pointer(in.FieldSelector))
	out.LabelSelector = (*authorization.LabelSelectorAttributes)(unsafe.Pointer(in.LabelSelector))
	return nil
}

// Convert_v1beta1_ResourceAttributes_To_authorization_ResourceAttributes is an autogenerated conversion function.
func Convert_v1beta1_ResourceAttributes_To_authorization_ResourceAttributes(in *v1beta1.ResourceAttributes, out *authorization.ResourceAttributes, s conversion.Scope) error {
	return autoConvert_v1beta1_ResourceAttributes_To_authorization_ResourceAttributes(in, out, s)
}

func autoConvert_authorization_ResourceAttributes_To_v1beta1_ResourceAttributes(in *authorization.ResourceAttributes, out *v1beta1.ResourceAttributes, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Verb = in.Verb
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	out.Subresource = in.Subresource
	out.Name = in.Name
	out.FieldSelector = (*v1.FieldSelectorAttributes)(unsafe.Pointer(in.FieldSelector))
	out.LabelSelector = (*v1.LabelSelectorAttributes)(unsafe.Pointer(in.LabelSelector))
	return nil
}

// Convert_authorization_ResourceAttributes_To_v1beta1_ResourceAttributes is an autogenerated conversion function.
func Convert_authorization_ResourceAttributes_To_v1beta1_ResourceAttributes(in *authorization.ResourceAttributes, out *v1beta1.ResourceAttributes, s conversion.Scope) error {
	return autoConvert_authorization_ResourceAttributes_To_v1beta1_ResourceAttributes(in, out, s)
}

func autoConvert_v1beta1_ResourceRule_To_authorization_ResourceRule(in *v1beta1.ResourceRule, out *authorization.ResourceRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	return nil
}

// Convert_v1beta1_ResourceRule_To_authorization_ResourceRule is an autogenerated conversion function.
func Convert_v1beta1_ResourceRule_To_authorization_ResourceRule(in *v1beta1.ResourceRule, out *authorization.ResourceRule, s conversion.Scope) error {
	return autoConvert_v1beta1_ResourceRule_To_authorization_ResourceRule(in, out, s)
}

func autoConvert_authorization_ResourceRule_To_v1beta1_ResourceRule(in *authorization.ResourceRule, out *v1beta1.ResourceRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	return nil
}

// Convert_authorization_ResourceRule_To_v1beta1_ResourceRule is an autogenerated conversion function.
func Convert_authorization_ResourceRule_To_v1beta1_ResourceRule(in *authorization.ResourceRule, out *v1beta1.ResourceRule, s conversion.Scope) error {
	return autoConvert_authorization_ResourceRule_To_v1beta1_ResourceRule(in, out, s)
}

func autoConvert_v1beta1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in *v1beta1.SelfSubjectAccessReview, out *authorization.SelfSubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview is an autogenerated conversion function.
func Convert_v1beta1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in *v1beta1.SelfSubjectAccessReview, out *authorization.SelfSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1beta1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in, out, s)
}

func autoConvert_authorization_SelfSubjectAccessReview_To_v1beta1_SelfSubjectAccessReview(in *authorization.SelfSubjectAccessReview, out *v1beta1.SelfSubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authorization_SelfSubjectAccessReview_To_v1beta1_SelfSubjectAccessReview is an autogenerated conversion function.
func Convert_authorization_SelfSubjectAccessReview_To_v1beta1_SelfSubjectAccessReview(in *authorization.SelfSubjectAccessReview, out *v1beta1.SelfSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectAccessReview_To_v1beta1_SelfSubjectAccessReview(in, out, s)
}

func autoConvert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in *v1beta1.SelfSubjectAccessReviewSpec, out *authorization.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*authorization.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*authorization.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	return nil
}

// Convert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in *v1beta1.SelfSubjectAccessReviewSpec, out *authorization.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in, out, s)
}

func autoConvert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec(in *authorization.SelfSubjectAccessReviewSpec, out *v1beta1.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*v1beta1.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*v1beta1.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	return nil
}

// Convert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec(in *authorization.SelfSubjectAccessReviewSpec, out *v1beta1.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectAccessReviewSpec_To_v1beta1_SelfSubjectAccessReviewSpec(in, out, s)
}

func autoConvert_v1beta1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(in *v1beta1.SelfSubjectRulesReview, out *authorization.SelfSubjectRulesReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview is an autogenerated conversion function.
func Convert_v1beta1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(in *v1beta1.SelfSubjectRulesReview, out *authorization.SelfSubjectRulesReview, s conversion.Scope) error {
	return autoConvert_v1beta1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(in, out, s)
}

func autoConvert_authorization_SelfSubjectRulesReview_To_v1beta1_SelfSubjectRulesReview(in *authorization.SelfSubjectRulesReview, out *v1beta1.SelfSubjectRulesReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authorization_SelfSubjectRulesReviewSpec_To_v1beta1_SelfSubjectRulesReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectRulesReviewStatus_To_v1beta1_SubjectRulesReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authorization_SelfSubjectRulesReview_To_v1beta1_SelfSubjectRulesReview is an autogenerated conversion function.
func Convert_authorization_SelfSubjectRulesReview_To_v1beta1_SelfSubjectRulesReview(in *authorization.SelfSubjectRulesReview, out *v1beta1.SelfSubjectRulesReview, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectRulesReview_To_v1beta1_SelfSubjectRulesReview(in, out, s)
}

func autoConvert_v1beta1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(in *v1beta1.SelfSubjectRulesReviewSpec, out *authorization.SelfSubjectRulesReviewSpec, s conversion.Scope) error {
	out.Namespace = in.Namespace
	return nil
}

// Convert_v1beta1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec is an autogenerated conversion function.
func Convert_v1beta1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(in *v1beta1.SelfSubjectRulesReviewSpec, out *authorization.SelfSubjectRulesReviewSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(in, out, s)
}

func autoConvert_authorization_SelfSubjectRulesReviewSpec_To_v1beta1_SelfSubjectRulesReviewSpec(in *authorization.SelfSubjectRulesReviewSpec, out *v1beta1.SelfSubjectRulesReviewSpec, s conversion.Scope) error {
	out.Namespace = in.Namespace
	return nil
}

// Convert_authorization_SelfSubjectRulesReviewSpec_To_v1beta1_SelfSubjectRulesReviewSpec is an autogenerated conversion function.
func Convert_authorization_SelfSubjectRulesReviewSpec_To_v1beta1_SelfSubjectRulesReviewSpec(in *authorization.SelfSubjectRulesReviewSpec, out *v1beta1.SelfSubjectRulesReviewSpec, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectRulesReviewSpec_To_v1beta1_SelfSubjectRulesReviewSpec(in, out, s)
}

func autoConvert_v1beta1_SubjectAccessReview_To_authorization_SubjectAccessReview(in *v1beta1.SubjectAccessReview, out *authorization.SubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_SubjectAccessReview_To_authorization_SubjectAccessReview is an autogenerated conversion function.
func Convert_v1beta1_SubjectAccessReview_To_authorization_SubjectAccessReview(in *v1beta1.SubjectAccessReview, out *authorization.SubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1beta1_SubjectAccessReview_To_authorization_SubjectAccessReview(in, out, s)
}

func autoConvert_authorization_SubjectAccessReview_To_v1beta1_SubjectAccessReview(in *authorization.SubjectAccessReview, out *v1beta1.SubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authorization_SubjectAccessReview_To_v1beta1_SubjectAccessReview is an autogenerated conversion function.
func Convert_authorization_SubjectAccessReview_To_v1beta1_SubjectAccessReview(in *authorization.SubjectAccessReview, out *v1beta1.SubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReview_To_v1beta1_SubjectAccessReview(in, out, s)
}

func autoConvert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in *v1beta1.SubjectAccessReviewSpec, out *authorization.SubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*authorization.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*authorization.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]authorization.ExtraValue)(unsafe.Pointer(&in.Extra))
	out.UID = in.UID
	return nil
}

// Convert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in *v1beta1.SubjectAccessReviewSpec, out *authorization.SubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in, out, s)
}

func autoConvert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(in *authorization.SubjectAccessReviewSpec, out *v1beta1.SubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*v1beta1.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*v1beta1.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]v1beta1.ExtraValue)(unsafe.Pointer(&in.Extra))
	out.UID = in.UID
	return nil
}

// Convert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(in *authorization.SubjectAccessReviewSpec, out *v1beta1.SubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReviewSpec_To_v1beta1_SubjectAccessReviewSpec(in, out, s)
}

func autoConvert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in *v1beta1.SubjectAccessReviewStatus, out *authorization.SubjectAccessReviewStatus, s conversion.Scope) error {
	out.Allowed = in.Allowed
	out.Denied = in.Denied
	out.Reason = in.Reason
	out.EvaluationError = in.EvaluationError
	return nil
}

// Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus is an autogenerated conversion function.
func Convert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in *v1beta1.SubjectAccessReviewStatus, out *authorization.SubjectAccessReviewStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in, out, s)
}

func autoConvert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(in *authorization.SubjectAccessReviewStatus, out *v1beta1.SubjectAccessReviewStatus, s conversion.Scope) error {
	out.Allowed = in.Allowed
	out.Denied = in.Denied
	out.Reason = in.Reason
	out.EvaluationError = in.EvaluationError
	return nil
}

// Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus is an autogenerated conversion function.
func Convert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(in *authorization.SubjectAccessReviewStatus, out *v1beta1.SubjectAccessReviewStatus, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReviewStatus_To_v1beta1_SubjectAccessReviewStatus(in, out, s)
}

func autoConvert_v1beta1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(in *v1beta1.SubjectRulesReviewStatus, out *authorization.SubjectRulesReviewStatus, s conversion.Scope) error {
	out.ResourceRules = *(*[]authorization.ResourceRule)(unsafe.Pointer(&in.ResourceRules))
	out.NonResourceRules = *(*[]authorization.NonResourceRule)(unsafe.Pointer(&in.NonResourceRules))
	out.Incomplete = in.Incomplete
	out.EvaluationError = in.EvaluationError
	return nil
}

// Convert_v1beta1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus is an autogenerated conversion function.
func Convert_v1beta1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(in *v1beta1.SubjectRulesReviewStatus, out *authorization.SubjectRulesReviewStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(in, out, s)
}

func autoConvert_authorization_SubjectRulesReviewStatus_To_v1beta1_SubjectRulesReviewStatus(in *authorization.SubjectRulesReviewStatus, out *v1beta1.SubjectRulesReviewStatus, s conversion.Scope) error {
	out.ResourceRules = *(*[]v1beta1.ResourceRule)(unsafe.Pointer(&in.ResourceRules))
	out.NonResourceRules = *(*[]v1beta1.NonResourceRule)(unsafe.Pointer(&in.NonResourceRules))
	out.Incomplete = in.Incomplete
	out.EvaluationError = in.EvaluationError
	return nil
}

// Convert_authorization_SubjectRulesReviewStatus_To_v1beta1_SubjectRulesReviewStatus is an autogenerated conversion function.
func Convert_authorization_SubjectRulesReviewStatus_To_v1beta1_SubjectRulesReviewStatus(in *authorization.SubjectRulesReviewStatus, out *v1beta1.SubjectRulesReviewStatus, s conversion.Scope) error {
	return autoConvert_authorization_SubjectRulesReviewStatus_To_v1beta1_SubjectRulesReviewStatus(in, out, s)
}
