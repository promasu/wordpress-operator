// +build !ignore_autogenerated

/*
Copyright 2018 Pressinfra SRL.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wordpress) DeepCopyInto(out *Wordpress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wordpress.
func (in *Wordpress) DeepCopy() *Wordpress {
	if in == nil {
		return nil
	}
	out := new(Wordpress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Wordpress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WordpressList) DeepCopyInto(out *WordpressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Wordpress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WordpressList.
func (in *WordpressList) DeepCopy() *WordpressList {
	if in == nil {
		return nil
	}
	out := new(WordpressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WordpressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WordpressPolicy) DeepCopyInto(out *WordpressPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WordpressPolicy.
func (in *WordpressPolicy) DeepCopy() *WordpressPolicy {
	if in == nil {
		return nil
	}
	out := new(WordpressPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WordpressPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WordpressPolicyList) DeepCopyInto(out *WordpressPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WordpressPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WordpressPolicyList.
func (in *WordpressPolicyList) DeepCopy() *WordpressPolicyList {
	if in == nil {
		return nil
	}
	out := new(WordpressPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WordpressPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WordpressPolicySpec) DeepCopyInto(out *WordpressPolicySpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WordpressPolicySpec.
func (in *WordpressPolicySpec) DeepCopy() *WordpressPolicySpec {
	if in == nil {
		return nil
	}
	out := new(WordpressPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WordpressPolicyStatus) DeepCopyInto(out *WordpressPolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WordpressPolicyStatus.
func (in *WordpressPolicyStatus) DeepCopy() *WordpressPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(WordpressPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WordpressRuntimeSpec) DeepCopyInto(out *WordpressRuntimeSpec) {
	*out = *in
	in.WebrootVolumeSpec.DeepCopyInto(&out.WebrootVolumeSpec)
	if in.MediaVolumeSpec != nil {
		in, out := &in.MediaVolumeSpec, &out.MediaVolumeSpec
		*out = new(WordpressVolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeMountsSpec != nil {
		in, out := &in.VolumeMountsSpec, &out.VolumeMountsSpec
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebPodTemplate != nil {
		in, out := &in.WebPodTemplate, &out.WebPodTemplate
		*out = new(corev1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CLIPodTemplate != nil {
		in, out := &in.CLIPodTemplate, &out.CLIPodTemplate
		*out = new(corev1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.IngressAnnotations != nil {
		in, out := &in.IngressAnnotations, &out.IngressAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceSpec != nil {
		in, out := &in.ServiceSpec, &out.ServiceSpec
		*out = new(corev1.ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WordpressRuntimeSpec.
func (in *WordpressRuntimeSpec) DeepCopy() *WordpressRuntimeSpec {
	if in == nil {
		return nil
	}
	out := new(WordpressRuntimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WordpressSpec) DeepCopyInto(out *WordpressSpec) {
	*out = *in
	in.WordpressRuntimeSpec.DeepCopyInto(&out.WordpressRuntimeSpec)
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]Domain, len(*in))
		copy(*out, *in)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WordpressSpec.
func (in *WordpressSpec) DeepCopy() *WordpressSpec {
	if in == nil {
		return nil
	}
	out := new(WordpressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WordpressStatus) DeepCopyInto(out *WordpressStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WordpressStatus.
func (in *WordpressStatus) DeepCopy() *WordpressStatus {
	if in == nil {
		return nil
	}
	out := new(WordpressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WordpressTemplateSpec) DeepCopyInto(out *WordpressTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WordpressTemplateSpec.
func (in *WordpressTemplateSpec) DeepCopy() *WordpressTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(WordpressTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WordpressVolumeSpec) DeepCopyInto(out *WordpressVolumeSpec) {
	*out = *in
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(corev1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(corev1.HostPathVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(corev1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WordpressVolumeSpec.
func (in *WordpressVolumeSpec) DeepCopy() *WordpressVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(WordpressVolumeSpec)
	in.DeepCopyInto(out)
	return out
}