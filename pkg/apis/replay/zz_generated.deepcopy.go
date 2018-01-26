// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

package replay

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Harvester) DeepCopyInto(out *Harvester) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Harvester.
func (in *Harvester) DeepCopy() *Harvester {
	if in == nil {
		return nil
	}
	out := new(Harvester)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Harvester) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarvesterList) DeepCopyInto(out *HarvesterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Harvester, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarvesterList.
func (in *HarvesterList) DeepCopy() *HarvesterList {
	if in == nil {
		return nil
	}
	out := new(HarvesterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HarvesterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarvesterSpec) DeepCopyInto(out *HarvesterSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarvesterSpec.
func (in *HarvesterSpec) DeepCopy() *HarvesterSpec {
	if in == nil {
		return nil
	}
	out := new(HarvesterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Silo) DeepCopyInto(out *Silo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Silo.
func (in *Silo) DeepCopy() *Silo {
	if in == nil {
		return nil
	}
	out := new(Silo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Silo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiloElasticsearchOutput) DeepCopyInto(out *SiloElasticsearchOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiloElasticsearchOutput.
func (in *SiloElasticsearchOutput) DeepCopy() *SiloElasticsearchOutput {
	if in == nil {
		return nil
	}
	out := new(SiloElasticsearchOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiloFileOutput) DeepCopyInto(out *SiloFileOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiloFileOutput.
func (in *SiloFileOutput) DeepCopy() *SiloFileOutput {
	if in == nil {
		return nil
	}
	out := new(SiloFileOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiloHttpOutput) DeepCopyInto(out *SiloHttpOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiloHttpOutput.
func (in *SiloHttpOutput) DeepCopy() *SiloHttpOutput {
	if in == nil {
		return nil
	}
	out := new(SiloHttpOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiloKafkaOutput) DeepCopyInto(out *SiloKafkaOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiloKafkaOutput.
func (in *SiloKafkaOutput) DeepCopy() *SiloKafkaOutput {
	if in == nil {
		return nil
	}
	out := new(SiloKafkaOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiloList) DeepCopyInto(out *SiloList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Silo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiloList.
func (in *SiloList) DeepCopy() *SiloList {
	if in == nil {
		return nil
	}
	out := new(SiloList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiloList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiloSpec) DeepCopyInto(out *SiloSpec) {
	*out = *in
	out.Output = in.Output
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiloSpec.
func (in *SiloSpec) DeepCopy() *SiloSpec {
	if in == nil {
		return nil
	}
	out := new(SiloSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiloStatus) DeepCopyInto(out *SiloStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiloStatus.
func (in *SiloStatus) DeepCopy() *SiloStatus {
	if in == nil {
		return nil
	}
	out := new(SiloStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiloStdoutOutput) DeepCopyInto(out *SiloStdoutOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiloStdoutOutput.
func (in *SiloStdoutOutput) DeepCopy() *SiloStdoutOutput {
	if in == nil {
		return nil
	}
	out := new(SiloStdoutOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiloTcpOutput) DeepCopyInto(out *SiloTcpOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiloTcpOutput.
func (in *SiloTcpOutput) DeepCopy() *SiloTcpOutput {
	if in == nil {
		return nil
	}
	out := new(SiloTcpOutput)
	in.DeepCopyInto(out)
	return out
}
