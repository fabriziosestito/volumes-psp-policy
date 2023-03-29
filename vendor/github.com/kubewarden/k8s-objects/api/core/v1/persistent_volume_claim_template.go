// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// PersistentVolumeClaimTemplate PersistentVolumeClaimTemplate is used to produce PersistentVolumeClaim objects as part of an EphemeralVolumeSource.
//
// swagger:model PersistentVolumeClaimTemplate
type PersistentVolumeClaimTemplate struct {

	// May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.
	Metadata *apimachinery_pkg_apis_meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.
	// Required: true
	Spec *PersistentVolumeClaimSpec `json:"spec"`
}
