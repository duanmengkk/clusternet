/*
Copyright The Clusternet Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ReplicaStatusApplyConfiguration represents a declarative configuration of the ReplicaStatus type for use
// with apply.
type ReplicaStatusApplyConfiguration struct {
	ObservedGeneration  *int64 `json:"observedGeneration,omitempty"`
	Replicas            *int32 `json:"replicas,omitempty"`
	UpdatedReplicas     *int32 `json:"updatedReplicas,omitempty"`
	CurrentReplicas     *int32 `json:"currentReplicas,omitempty"`
	ReadyReplicas       *int32 `json:"readyReplicas,omitempty"`
	AvailableReplicas   *int32 `json:"availableReplicas,omitempty"`
	UnavailableReplicas *int32 `json:"unavailableReplicas,omitempty"`
	Active              *int32 `json:"active,omitempty"`
	Succeeded           *int32 `json:"succeeded,omitempty"`
	Failed              *int32 `json:"failed,omitempty"`
}

// ReplicaStatusApplyConfiguration constructs a declarative configuration of the ReplicaStatus type for use with
// apply.
func ReplicaStatus() *ReplicaStatusApplyConfiguration {
	return &ReplicaStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *ReplicaStatusApplyConfiguration) WithObservedGeneration(value int64) *ReplicaStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *ReplicaStatusApplyConfiguration) WithReplicas(value int32) *ReplicaStatusApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithUpdatedReplicas sets the UpdatedReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdatedReplicas field is set to the value of the last call.
func (b *ReplicaStatusApplyConfiguration) WithUpdatedReplicas(value int32) *ReplicaStatusApplyConfiguration {
	b.UpdatedReplicas = &value
	return b
}

// WithCurrentReplicas sets the CurrentReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentReplicas field is set to the value of the last call.
func (b *ReplicaStatusApplyConfiguration) WithCurrentReplicas(value int32) *ReplicaStatusApplyConfiguration {
	b.CurrentReplicas = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *ReplicaStatusApplyConfiguration) WithReadyReplicas(value int32) *ReplicaStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// WithAvailableReplicas sets the AvailableReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailableReplicas field is set to the value of the last call.
func (b *ReplicaStatusApplyConfiguration) WithAvailableReplicas(value int32) *ReplicaStatusApplyConfiguration {
	b.AvailableReplicas = &value
	return b
}

// WithUnavailableReplicas sets the UnavailableReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnavailableReplicas field is set to the value of the last call.
func (b *ReplicaStatusApplyConfiguration) WithUnavailableReplicas(value int32) *ReplicaStatusApplyConfiguration {
	b.UnavailableReplicas = &value
	return b
}

// WithActive sets the Active field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Active field is set to the value of the last call.
func (b *ReplicaStatusApplyConfiguration) WithActive(value int32) *ReplicaStatusApplyConfiguration {
	b.Active = &value
	return b
}

// WithSucceeded sets the Succeeded field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Succeeded field is set to the value of the last call.
func (b *ReplicaStatusApplyConfiguration) WithSucceeded(value int32) *ReplicaStatusApplyConfiguration {
	b.Succeeded = &value
	return b
}

// WithFailed sets the Failed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Failed field is set to the value of the last call.
func (b *ReplicaStatusApplyConfiguration) WithFailed(value int32) *ReplicaStatusApplyConfiguration {
	b.Failed = &value
	return b
}
