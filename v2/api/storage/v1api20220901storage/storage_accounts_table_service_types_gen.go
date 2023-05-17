// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccountstableservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccountstableservices/status,storageaccountstableservices/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220901.StorageAccountsTableService
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default
type StorageAccountsTableService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccounts_TableService_Spec   `json:"spec,omitempty"`
	Status            StorageAccounts_TableService_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsTableService{}

// GetConditions returns the conditions of the resource
func (service *StorageAccountsTableService) GetConditions() conditions.Conditions {
	return service.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (service *StorageAccountsTableService) SetConditions(conditions conditions.Conditions) {
	service.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &StorageAccountsTableService{}

// AzureName returns the Azure name of the resource (always "default")
func (service *StorageAccountsTableService) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
func (service StorageAccountsTableService) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (service *StorageAccountsTableService) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (service *StorageAccountsTableService) GetSpec() genruntime.ConvertibleSpec {
	return &service.Spec
}

// GetStatus returns the status of this resource
func (service *StorageAccountsTableService) GetStatus() genruntime.ConvertibleStatus {
	return &service.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/tableServices"
func (service *StorageAccountsTableService) GetType() string {
	return "Microsoft.Storage/storageAccounts/tableServices"
}

// NewEmptyStatus returns a new empty (blank) status
func (service *StorageAccountsTableService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccounts_TableService_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (service *StorageAccountsTableService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  service.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (service *StorageAccountsTableService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccounts_TableService_STATUS); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccounts_TableService_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// Hub marks that this StorageAccountsTableService is the hub type for conversion
func (service *StorageAccountsTableService) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsTableService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion,
		Kind:    "StorageAccountsTableService",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220901.StorageAccountsTableService
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default
type StorageAccountsTableServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsTableService `json:"items"`
}

// Storage version of v1api20220901.StorageAccounts_TableService_Spec
type StorageAccounts_TableService_Spec struct {
	Cors            *CorsRules `json:"cors,omitempty"`
	OriginalVersion string     `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccount resource
	Owner       *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccount"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccounts_TableService_Spec{}

// ConvertSpecFrom populates our StorageAccounts_TableService_Spec from the provided source
func (service *StorageAccounts_TableService_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(service)
}

// ConvertSpecTo populates the provided destination from our StorageAccounts_TableService_Spec
func (service *StorageAccounts_TableService_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(service)
}

// Storage version of v1api20220901.StorageAccounts_TableService_STATUS
type StorageAccounts_TableService_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Cors        *CorsRules_STATUS      `json:"cors,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccounts_TableService_STATUS{}

// ConvertStatusFrom populates our StorageAccounts_TableService_STATUS from the provided source
func (service *StorageAccounts_TableService_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(service)
}

// ConvertStatusTo populates the provided destination from our StorageAccounts_TableService_STATUS
func (service *StorageAccounts_TableService_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(service)
}

func init() {
	SchemeBuilder.Register(&StorageAccountsTableService{}, &StorageAccountsTableServiceList{})
}
