/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and clustersecret-operator contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/sap/clustersecret-operator/pkg/apis/core.cs.sap.com/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterSecretConditionApplyConfiguration represents an declarative configuration of the ClusterSecretCondition type for use
// with apply.
type ClusterSecretConditionApplyConfiguration struct {
	Type               *v1alpha1.ClusterSecretConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus                  `json:"status,omitempty"`
	LastUpdateTime     *metav1.Time                         `json:"lastUpdateTime,omitempty"`
	LastTransitionTime *metav1.Time                         `json:"lastTransitionTime,omitempty"`
	Reason             *string                              `json:"reason,omitempty"`
	Message            *string                              `json:"message,omitempty"`
}

// ClusterSecretConditionApplyConfiguration constructs an declarative configuration of the ClusterSecretCondition type for use with
// apply.
func ClusterSecretCondition() *ClusterSecretConditionApplyConfiguration {
	return &ClusterSecretConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ClusterSecretConditionApplyConfiguration) WithType(value v1alpha1.ClusterSecretConditionType) *ClusterSecretConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ClusterSecretConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *ClusterSecretConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastUpdateTime sets the LastUpdateTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastUpdateTime field is set to the value of the last call.
func (b *ClusterSecretConditionApplyConfiguration) WithLastUpdateTime(value metav1.Time) *ClusterSecretConditionApplyConfiguration {
	b.LastUpdateTime = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *ClusterSecretConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *ClusterSecretConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *ClusterSecretConditionApplyConfiguration) WithReason(value string) *ClusterSecretConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ClusterSecretConditionApplyConfiguration) WithMessage(value string) *ClusterSecretConditionApplyConfiguration {
	b.Message = &value
	return b
}
