/*
Copyright 2016 The Kubernetes Authors.

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

package settings

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/api"
)

// +genclient=true

// PodPreset is a policy resource that defines additional runtime
// requirements for a Pod.
type PodPreset struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// +optional
	Spec PodPresetSpec
}

// PodPresetSpec is a description of a pod preset.
type PodPresetSpec struct {
	// Selector is a label query over a set of resources, in this case pods.
	// Required.
	Selector metav1.LabelSelector

	// AppContainerSelector is a ContainerSelector to select app containers of a
	// PodSpec for injecting information into.
	// If not specified, it will default to EmptyContainerSelector which matches
	// any app-container in the PodSpec.
	// +optional
	AppContainerSelector *ContainerSelector `json:"omitempty"`

	// Env defines the collection of EnvVar to inject into containers.
	// +optional
	Env []api.EnvVar
	// EnvFrom defines the collection of EnvFromSource to inject into containers.
	// +optional
	EnvFrom []api.EnvFromSource
	// Volumes defines the collection of Volume to inject into the pod.
	// +optional
	Volumes []api.Volume
	// VolumeMounts defines the collection of VolumeMount to inject into containers.
	// +optional
	VolumeMounts []api.VolumeMount
}

// PodPresetList is a list of PodPreset objects.
type PodPresetList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta

	Items []PodPreset
}

// ContainerSelector represents container selector criterion that can be applied on
// list of v1.Container to select containers.
type ContainerSelector struct {
	// MatchNames is a list of container names to be matched to. This selector
	// returns a match if container name matches any of the names in this list.
	// Special string “*” is treated as match ANY.
	// Empty or nil list is treated as match NOTHING.
	MatchNames []string

	// in future matchExpressions support can also be added to enable set expresssions.
}
