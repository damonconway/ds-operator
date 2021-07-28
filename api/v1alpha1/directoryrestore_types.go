/*

Copyright 2021 ForgeRock AS.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DirectoryRestoreSpec defines the desired state of DirectoryRestore
type DirectoryRestoreSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	SourcePVCName string `json:"sourcePvcName,required"`
	// Docker Image for the directory server.
	Image string `json:"image"`

	// Keystore references
	Keystore DirectoryKeystores `json:"keystore,required"`

	RestorePVC RestorePVC `json:"restore,required"`
}

type RestorePVC struct {
	Size                    string `json:"size"`
	StorageClassName        string `json:"storageClassName"`
	VolumeSnapshotClassName string `json:"volumeSnapshotClassName,required"`
}

// DirectoryRestoreStatus defines the observed state of DirectoryRestore
type DirectoryRestoreStatus struct {

	// StartTimestamp records the time a restore job was started.
	// The server's time is used for StartTimestamps
	// +optional
	// +nullable
	StartTimestamp *metav1.Time `json:"startTimestamp,omitempty"`

	// CompletionTimestamp records the time a restore was completed.
	// Completion time is recorded even on failed backups.
	// Completion time is recorded before uploading the backup object.
	// The server's time is used for CompletionTimestamps
	// +optional
	// +nullable
	CompletionTimestamp *metav1.Time `json:"completionTimestamp,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DirectoryRestore is the Schema for the directoryrestores API
type DirectoryRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DirectoryRestoreSpec   `json:"spec,omitempty"`
	Status DirectoryRestoreStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DirectoryRestoreList contains a list of DirectoryRestore
type DirectoryRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectoryRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DirectoryRestore{}, &DirectoryRestoreList{})
}