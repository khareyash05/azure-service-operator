// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

// The storage account.
type StorageAccount_STATUS_ARM struct {
	// ExtendedLocation: The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: The identity of the resource.
	Identity *Identity_STATUS_ARM `json:"identity,omitempty"`

	// Kind: Gets the Kind.
	Kind *StorageAccount_Kind_STATUS `json:"kind,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the storage account.
	Properties *StorageAccountProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: Gets the SKU.
	Sku *Sku_STATUS_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The complex type of the extended location.
type ExtendedLocation_STATUS_ARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType_STATUS `json:"type,omitempty"`
}

// Identity for the resource.
type Identity_STATUS_ARM struct {
	// PrincipalId: The principal ID of resource identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of resource.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type.
	Type *Identity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: Gets or sets a list of key value pairs that describe the set of User Assigned identities that
	// will be used with this storage account. The key is the ARM resource identifier of the identity. Only 1 User Assigned
	// identity is permitted here.
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

// The SKU of the storage account.
type Sku_STATUS_ARM struct {
	// Name: The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called
	//  accountType.
	Name *SkuName_STATUS `json:"name,omitempty"`

	// Tier: The SKU tier. This is based on the SKU name.
	Tier *Tier_STATUS `json:"tier,omitempty"`
}

type StorageAccount_Kind_STATUS string

const (
	StorageAccount_Kind_STATUS_BlobStorage      = StorageAccount_Kind_STATUS("BlobStorage")
	StorageAccount_Kind_STATUS_BlockBlobStorage = StorageAccount_Kind_STATUS("BlockBlobStorage")
	StorageAccount_Kind_STATUS_FileStorage      = StorageAccount_Kind_STATUS("FileStorage")
	StorageAccount_Kind_STATUS_Storage          = StorageAccount_Kind_STATUS("Storage")
	StorageAccount_Kind_STATUS_StorageV2        = StorageAccount_Kind_STATUS("StorageV2")
)

// Properties of the storage account.
type StorageAccountProperties_STATUS_ARM struct {
	// AccessTier: Required for storage accounts where kind = BlobStorage. The access tier is used for billing. The 'Premium'
	// access tier is the default value for premium block blobs storage account type and it cannot be changed for the premium
	// block blobs storage account type.
	AccessTier *StorageAccountProperties_AccessTier_STATUS `json:"accessTier,omitempty"`

	// AllowBlobPublicAccess: Allow or disallow public access to all blobs or containers in the storage account. The default
	// interpretation is true for this property.
	AllowBlobPublicAccess *bool `json:"allowBlobPublicAccess,omitempty"`

	// AllowCrossTenantReplication: Allow or disallow cross AAD tenant object replication. The default interpretation is true
	// for this property.
	AllowCrossTenantReplication *bool `json:"allowCrossTenantReplication,omitempty"`

	// AllowSharedKeyAccess: Indicates whether the storage account permits requests to be authorized with the account access
	// key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure
	// Active Directory (Azure AD). The default value is null, which is equivalent to true.
	AllowSharedKeyAccess *bool `json:"allowSharedKeyAccess,omitempty"`

	// AllowedCopyScope: Restrict copy to and from Storage Accounts within an AAD tenant or with Private Links to the same VNet.
	AllowedCopyScope *StorageAccountProperties_AllowedCopyScope_STATUS `json:"allowedCopyScope,omitempty"`

	// AzureFilesIdentityBasedAuthentication: Provides the identity based authentication settings for Azure Files.
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication_STATUS_ARM `json:"azureFilesIdentityBasedAuthentication,omitempty"`

	// BlobRestoreStatus: Blob restore status
	BlobRestoreStatus *BlobRestoreStatus_STATUS_ARM `json:"blobRestoreStatus,omitempty"`

	// CreationTime: Gets the creation date and time of the storage account in UTC.
	CreationTime *string `json:"creationTime,omitempty"`

	// CustomDomain: Gets the custom domain the user assigned to this storage account.
	CustomDomain *CustomDomain_STATUS_ARM `json:"customDomain,omitempty"`

	// DefaultToOAuthAuthentication: A boolean flag which indicates whether the default authentication is OAuth or not. The
	// default interpretation is false for this property.
	DefaultToOAuthAuthentication *bool `json:"defaultToOAuthAuthentication,omitempty"`

	// DnsEndpointType: Allows you to specify the type of endpoint. Set this to AzureDNSZone to create a large number of
	// accounts in a single subscription, which creates accounts in an Azure DNS Zone and the endpoint URL will have an
	// alphanumeric DNS Zone identifier.
	DnsEndpointType *StorageAccountProperties_DnsEndpointType_STATUS `json:"dnsEndpointType,omitempty"`

	// Encryption: Encryption settings to be used for server-side encryption for the storage account.
	Encryption *Encryption_STATUS_ARM `json:"encryption,omitempty"`

	// FailoverInProgress: If the failover is in progress, the value will be true, otherwise, it will be null.
	FailoverInProgress *bool `json:"failoverInProgress,omitempty"`

	// GeoReplicationStats: Geo Replication Stats
	GeoReplicationStats *GeoReplicationStats_STATUS_ARM `json:"geoReplicationStats,omitempty"`

	// ImmutableStorageWithVersioning: The property is immutable and can only be set to true at the account creation time. When
	// set to true, it enables object level immutability for all the containers in the account by default.
	ImmutableStorageWithVersioning *ImmutableStorageAccount_STATUS_ARM `json:"immutableStorageWithVersioning,omitempty"`

	// IsHnsEnabled: Account HierarchicalNamespace enabled if sets to true.
	IsHnsEnabled *bool `json:"isHnsEnabled,omitempty"`

	// IsLocalUserEnabled: Enables local users feature, if set to true
	IsLocalUserEnabled *bool `json:"isLocalUserEnabled,omitempty"`

	// IsNfsV3Enabled: NFS 3.0 protocol support enabled if set to true.
	IsNfsV3Enabled *bool `json:"isNfsV3Enabled,omitempty"`

	// IsSftpEnabled: Enables Secure File Transfer Protocol, if set to true
	IsSftpEnabled *bool `json:"isSftpEnabled,omitempty"`

	// KeyCreationTime: Storage account keys creation time.
	KeyCreationTime *KeyCreationTime_STATUS_ARM `json:"keyCreationTime,omitempty"`

	// KeyPolicy: KeyPolicy assigned to the storage account.
	KeyPolicy *KeyPolicy_STATUS_ARM `json:"keyPolicy,omitempty"`

	// LargeFileSharesState: Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
	LargeFileSharesState *StorageAccountProperties_LargeFileSharesState_STATUS `json:"largeFileSharesState,omitempty"`

	// LastGeoFailoverTime: Gets the timestamp of the most recent instance of a failover to the secondary location. Only the
	// most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only
	// available if the accountType is Standard_GRS or Standard_RAGRS.
	LastGeoFailoverTime *string `json:"lastGeoFailoverTime,omitempty"`

	// MinimumTlsVersion: Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS
	// 1.0 for this property.
	MinimumTlsVersion *StorageAccountProperties_MinimumTlsVersion_STATUS `json:"minimumTlsVersion,omitempty"`

	// NetworkAcls: Network rule set
	NetworkAcls *NetworkRuleSet_STATUS_ARM `json:"networkAcls,omitempty"`

	// PrimaryEndpoints: Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that
	// Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
	PrimaryEndpoints *Endpoints_STATUS_ARM `json:"primaryEndpoints,omitempty"`

	// PrimaryLocation: Gets the location of the primary data center for the storage account.
	PrimaryLocation *string `json:"primaryLocation,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connection associated with the specified storage account
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Gets the status of the storage account at the time the operation was called.
	ProvisioningState *StorageAccountProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Allow or disallow public network access to Storage Account. Value is optional but if passed in,
	// must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *StorageAccountProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// RoutingPreference: Maintains information about the network routing choice opted by the user for data transfer
	RoutingPreference *RoutingPreference_STATUS_ARM `json:"routingPreference,omitempty"`

	// SasPolicy: SasPolicy assigned to the storage account.
	SasPolicy *SasPolicy_STATUS_ARM `json:"sasPolicy,omitempty"`

	// SecondaryEndpoints: Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the
	// secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
	SecondaryEndpoints *Endpoints_STATUS_ARM `json:"secondaryEndpoints,omitempty"`

	// SecondaryLocation: Gets the location of the geo-replicated secondary for the storage account. Only available if the
	// accountType is Standard_GRS or Standard_RAGRS.
	SecondaryLocation *string `json:"secondaryLocation,omitempty"`

	// StatusOfPrimary: Gets the status indicating whether the primary location of the storage account is available or
	// unavailable.
	StatusOfPrimary *StorageAccountProperties_StatusOfPrimary_STATUS `json:"statusOfPrimary,omitempty"`

	// StatusOfSecondary: Gets the status indicating whether the secondary location of the storage account is available or
	// unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
	StatusOfSecondary *StorageAccountProperties_StatusOfSecondary_STATUS `json:"statusOfSecondary,omitempty"`

	// StorageAccountSkuConversionStatus: This property is readOnly and is set by server during asynchronous storage account
	// sku conversion operations.
	StorageAccountSkuConversionStatus *StorageAccountSkuConversionStatus_STATUS_ARM `json:"storageAccountSkuConversionStatus,omitempty"`

	// SupportsHttpsTrafficOnly: Allows https traffic only to storage service if sets to true.
	SupportsHttpsTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`
}

// Settings for Azure Files identity based authentication.
type AzureFilesIdentityBasedAuthentication_STATUS_ARM struct {
	// ActiveDirectoryProperties: Required if directoryServiceOptions are AD, optional if they are AADKERB.
	ActiveDirectoryProperties *ActiveDirectoryProperties_STATUS_ARM `json:"activeDirectoryProperties,omitempty"`

	// DefaultSharePermission: Default share permission for users using Kerberos authentication if RBAC role is not assigned.
	DefaultSharePermission *AzureFilesIdentityBasedAuthentication_DefaultSharePermission_STATUS `json:"defaultSharePermission,omitempty"`

	// DirectoryServiceOptions: Indicates the directory service used. Note that this enum may be extended in the future.
	DirectoryServiceOptions *AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_STATUS `json:"directoryServiceOptions,omitempty"`
}

// Blob restore status.
type BlobRestoreStatus_STATUS_ARM struct {
	// FailureReason: Failure reason when blob restore is failed.
	FailureReason *string `json:"failureReason,omitempty"`

	// Parameters: Blob restore request parameters.
	Parameters *BlobRestoreParameters_STATUS_ARM `json:"parameters,omitempty"`

	// RestoreId: Id for tracking blob restore request.
	RestoreId *string `json:"restoreId,omitempty"`

	// Status: The status of blob restore progress. Possible values are: - InProgress: Indicates that blob restore is ongoing.
	// - Complete: Indicates that blob restore has been completed successfully. - Failed: Indicates that blob restore is failed.
	Status *BlobRestoreStatus_Status_STATUS `json:"status,omitempty"`
}

// The custom domain assigned to this storage account. This can be set via Update.
type CustomDomain_STATUS_ARM struct {
	// Name: Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
	Name *string `json:"name,omitempty"`

	// UseSubDomainName: Indicates whether indirect CName validation is enabled. Default value is false. This should only be
	// set on updates.
	UseSubDomainName *bool `json:"useSubDomainName,omitempty"`
}

// The encryption settings on the storage account.
type Encryption_STATUS_ARM struct {
	// Identity: The identity to be used with service-side encryption at rest.
	Identity *EncryptionIdentity_STATUS_ARM `json:"identity,omitempty"`

	// KeySource: The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage,
	// Microsoft.Keyvault
	KeySource *Encryption_KeySource_STATUS `json:"keySource,omitempty"`

	// Keyvaultproperties: Properties provided by key vault.
	Keyvaultproperties *KeyVaultProperties_STATUS_ARM `json:"keyvaultproperties,omitempty"`

	// RequireInfrastructureEncryption: A boolean indicating whether or not the service applies a secondary layer of encryption
	// with platform managed keys for data at rest.
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`

	// Services: List of services which support encryption.
	Services *EncryptionServices_STATUS_ARM `json:"services,omitempty"`
}

// The URIs that are used to perform a retrieval of a public blob, queue, table, web or dfs object.
type Endpoints_STATUS_ARM struct {
	// Blob: Gets the blob endpoint.
	Blob *string `json:"blob,omitempty"`

	// Dfs: Gets the dfs endpoint.
	Dfs *string `json:"dfs,omitempty"`

	// File: Gets the file endpoint.
	File *string `json:"file,omitempty"`

	// InternetEndpoints: Gets the internet routing storage endpoints
	InternetEndpoints *StorageAccountInternetEndpoints_STATUS_ARM `json:"internetEndpoints,omitempty"`

	// MicrosoftEndpoints: Gets the microsoft routing storage endpoints.
	MicrosoftEndpoints *StorageAccountMicrosoftEndpoints_STATUS_ARM `json:"microsoftEndpoints,omitempty"`

	// Queue: Gets the queue endpoint.
	Queue *string `json:"queue,omitempty"`

	// Table: Gets the table endpoint.
	Table *string `json:"table,omitempty"`

	// Web: Gets the web endpoint.
	Web *string `json:"web,omitempty"`
}

// The type of extendedLocation.
type ExtendedLocationType_STATUS string

const ExtendedLocationType_STATUS_EdgeZone = ExtendedLocationType_STATUS("EdgeZone")

// Statistics related to replication for storage account's Blob, Table, Queue and File services. It is only available when
// geo-redundant replication is enabled for the storage account.
type GeoReplicationStats_STATUS_ARM struct {
	// CanFailover: A boolean flag which indicates whether or not account failover is supported for the account.
	CanFailover *bool `json:"canFailover,omitempty"`

	// LastSyncTime: All primary writes preceding this UTC date/time value are guaranteed to be available for read operations.
	// Primary writes following this point in time may or may not be available for reads. Element may be default value if value
	// of LastSyncTime is not available, this can happen if secondary is offline or we are in bootstrap.
	LastSyncTime *string `json:"lastSyncTime,omitempty"`

	// Status: The status of the secondary location. Possible values are: - Live: Indicates that the secondary location is
	// active and operational. - Bootstrap: Indicates initial synchronization from the primary location to the secondary
	// location is in progress.This typically occurs when replication is first enabled. - Unavailable: Indicates that the
	// secondary location is temporarily unavailable.
	Status *GeoReplicationStats_Status_STATUS `json:"status,omitempty"`
}

type Identity_Type_STATUS string

const (
	Identity_Type_STATUS_None                       = Identity_Type_STATUS("None")
	Identity_Type_STATUS_SystemAssigned             = Identity_Type_STATUS("SystemAssigned")
	Identity_Type_STATUS_SystemAssignedUserAssigned = Identity_Type_STATUS("SystemAssigned,UserAssigned")
	Identity_Type_STATUS_UserAssigned               = Identity_Type_STATUS("UserAssigned")
)

// This property enables and defines account-level immutability. Enabling the feature auto-enables Blob Versioning.
type ImmutableStorageAccount_STATUS_ARM struct {
	// Enabled: A boolean flag which enables account-level immutability. All the containers under such an account have
	// object-level immutability enabled by default.
	Enabled *bool `json:"enabled,omitempty"`

	// ImmutabilityPolicy: Specifies the default account-level immutability policy which is inherited and applied to objects
	// that do not possess an explicit immutability policy at the object level. The object-level immutability policy has higher
	// precedence than the container-level immutability policy, which has a higher precedence than the account-level
	// immutability policy.
	ImmutabilityPolicy *AccountImmutabilityPolicyProperties_STATUS_ARM `json:"immutabilityPolicy,omitempty"`
}

// Storage account keys creation time.
type KeyCreationTime_STATUS_ARM struct {
	Key1 *string `json:"key1,omitempty"`
	Key2 *string `json:"key2,omitempty"`
}

// KeyPolicy assigned to the storage account.
type KeyPolicy_STATUS_ARM struct {
	// KeyExpirationPeriodInDays: The key expiration period in days.
	KeyExpirationPeriodInDays *int `json:"keyExpirationPeriodInDays,omitempty"`
}

// Network rule set
type NetworkRuleSet_STATUS_ARM struct {
	// Bypass: Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of
	// Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics.
	Bypass *NetworkRuleSet_Bypass_STATUS `json:"bypass,omitempty"`

	// DefaultAction: Specifies the default action of allow or deny when no other rules match.
	DefaultAction *NetworkRuleSet_DefaultAction_STATUS `json:"defaultAction,omitempty"`

	// IpRules: Sets the IP ACL rules
	IpRules []IPRule_STATUS_ARM `json:"ipRules,omitempty"`

	// ResourceAccessRules: Sets the resource access rules
	ResourceAccessRules []ResourceAccessRule_STATUS_ARM `json:"resourceAccessRules,omitempty"`

	// VirtualNetworkRules: Sets the virtual network rules
	VirtualNetworkRules []VirtualNetworkRule_STATUS_ARM `json:"virtualNetworkRules,omitempty"`
}

// The Private Endpoint Connection resource.
type PrivateEndpointConnection_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`
}

// Routing preference defines the type of network, either microsoft or internet routing to be used to deliver the user
// data, the default option is microsoft routing
type RoutingPreference_STATUS_ARM struct {
	// PublishInternetEndpoints: A boolean flag which indicates whether internet routing storage endpoints are to be published
	PublishInternetEndpoints *bool `json:"publishInternetEndpoints,omitempty"`

	// PublishMicrosoftEndpoints: A boolean flag which indicates whether microsoft routing storage endpoints are to be published
	PublishMicrosoftEndpoints *bool `json:"publishMicrosoftEndpoints,omitempty"`

	// RoutingChoice: Routing Choice defines the kind of network routing opted by the user.
	RoutingChoice *RoutingPreference_RoutingChoice_STATUS `json:"routingChoice,omitempty"`
}

// SasPolicy assigned to the storage account.
type SasPolicy_STATUS_ARM struct {
	// ExpirationAction: The SAS expiration action. Can only be Log.
	ExpirationAction *SasPolicy_ExpirationAction_STATUS `json:"expirationAction,omitempty"`

	// SasExpirationPeriod: The SAS expiration period, DD.HH:MM:SS.
	SasExpirationPeriod *string `json:"sasExpirationPeriod,omitempty"`
}

// The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called
// accountType.
type SkuName_STATUS string

const (
	SkuName_STATUS_Premium_LRS     = SkuName_STATUS("Premium_LRS")
	SkuName_STATUS_Premium_ZRS     = SkuName_STATUS("Premium_ZRS")
	SkuName_STATUS_Standard_GRS    = SkuName_STATUS("Standard_GRS")
	SkuName_STATUS_Standard_GZRS   = SkuName_STATUS("Standard_GZRS")
	SkuName_STATUS_Standard_LRS    = SkuName_STATUS("Standard_LRS")
	SkuName_STATUS_Standard_RAGRS  = SkuName_STATUS("Standard_RAGRS")
	SkuName_STATUS_Standard_RAGZRS = SkuName_STATUS("Standard_RAGZRS")
	SkuName_STATUS_Standard_ZRS    = SkuName_STATUS("Standard_ZRS")
)

// This defines the sku conversion status object for asynchronous sku conversions.
type StorageAccountSkuConversionStatus_STATUS_ARM struct {
	// EndTime: This property represents the sku conversion end time.
	EndTime *string `json:"endTime,omitempty"`

	// SkuConversionStatus: This property indicates the current sku conversion status.
	SkuConversionStatus *StorageAccountSkuConversionStatus_SkuConversionStatus_STATUS `json:"skuConversionStatus,omitempty"`

	// StartTime: This property represents the sku conversion start time.
	StartTime *string `json:"startTime,omitempty"`

	// TargetSkuName: This property represents the target sku name to which the account sku is being converted asynchronously.
	TargetSkuName *SkuName_STATUS `json:"targetSkuName,omitempty"`
}

// The SKU tier. This is based on the SKU name.
type Tier_STATUS string

const (
	Tier_STATUS_Premium  = Tier_STATUS("Premium")
	Tier_STATUS_Standard = Tier_STATUS("Standard")
)

// UserAssignedIdentity for the resource.
type UserAssignedIdentity_STATUS_ARM struct {
	// ClientId: The client ID of the identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal ID of the identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

// This defines account-level immutability policy properties.
type AccountImmutabilityPolicyProperties_STATUS_ARM struct {
	// AllowProtectedAppendWrites: This property can only be changed for disabled and unlocked time-based retention policies.
	// When enabled, new blocks can be written to an append blob while maintaining immutability protection and compliance. Only
	// new blocks can be added and any existing blocks cannot be modified or deleted.
	AllowProtectedAppendWrites *bool `json:"allowProtectedAppendWrites,omitempty"`

	// ImmutabilityPeriodSinceCreationInDays: The immutability period for the blobs in the container since the policy creation,
	// in days.
	ImmutabilityPeriodSinceCreationInDays *int `json:"immutabilityPeriodSinceCreationInDays,omitempty"`

	// State: The ImmutabilityPolicy state defines the mode of the policy. Disabled state disables the policy, Unlocked state
	// allows increase and decrease of immutability retention time and also allows toggling allowProtectedAppendWrites
	// property, Locked state only allows the increase of the immutability retention time. A policy can only be created in a
	// Disabled or Unlocked state and can be toggled between the two states. Only a policy in an Unlocked state can transition
	// to a Locked state which cannot be reverted.
	State *AccountImmutabilityPolicyProperties_State_STATUS `json:"state,omitempty"`
}

// Settings properties for Active Directory (AD).
type ActiveDirectoryProperties_STATUS_ARM struct {
	// AccountType: Specifies the Active Directory account type for Azure Storage.
	AccountType *ActiveDirectoryProperties_AccountType_STATUS `json:"accountType,omitempty"`

	// AzureStorageSid: Specifies the security identifier (SID) for Azure Storage.
	AzureStorageSid *string `json:"azureStorageSid,omitempty"`

	// DomainGuid: Specifies the domain GUID.
	DomainGuid *string `json:"domainGuid,omitempty"`

	// DomainName: Specifies the primary domain that the AD DNS server is authoritative for.
	DomainName *string `json:"domainName,omitempty"`

	// DomainSid: Specifies the security identifier (SID).
	DomainSid *string `json:"domainSid,omitempty"`

	// ForestName: Specifies the Active Directory forest to get.
	ForestName *string `json:"forestName,omitempty"`

	// NetBiosDomainName: Specifies the NetBIOS domain name.
	NetBiosDomainName *string `json:"netBiosDomainName,omitempty"`

	// SamAccountName: Specifies the Active Directory SAMAccountName for Azure Storage.
	SamAccountName *string `json:"samAccountName,omitempty"`
}

// Blob restore parameters
type BlobRestoreParameters_STATUS_ARM struct {
	// BlobRanges: Blob ranges to restore.
	BlobRanges []BlobRestoreRange_STATUS_ARM `json:"blobRanges,omitempty"`

	// TimeToRestore: Restore blob to the specified time.
	TimeToRestore *string `json:"timeToRestore,omitempty"`
}

// Encryption identity for the storage account.
type EncryptionIdentity_STATUS_ARM struct {
	// FederatedIdentityClientId: ClientId of the multi-tenant application to be used in conjunction with the user-assigned
	// identity for cross-tenant customer-managed-keys server-side encryption on the storage account.
	FederatedIdentityClientId *string `json:"federatedIdentityClientId,omitempty"`

	// UserAssignedIdentity: Resource identifier of the UserAssigned identity to be associated with server-side encryption on
	// the storage account.
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// A list of services that support encryption.
type EncryptionServices_STATUS_ARM struct {
	// Blob: The encryption function of the blob storage service.
	Blob *EncryptionService_STATUS_ARM `json:"blob,omitempty"`

	// File: The encryption function of the file storage service.
	File *EncryptionService_STATUS_ARM `json:"file,omitempty"`

	// Queue: The encryption function of the queue storage service.
	Queue *EncryptionService_STATUS_ARM `json:"queue,omitempty"`

	// Table: The encryption function of the table storage service.
	Table *EncryptionService_STATUS_ARM `json:"table,omitempty"`
}

// IP rule with specific IP or IP range in CIDR format.
type IPRule_STATUS_ARM struct {
	// Action: The action of IP ACL rule.
	Action *IPRule_Action_STATUS `json:"action,omitempty"`

	// Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	Value *string `json:"value,omitempty"`
}

// Properties of key vault.
type KeyVaultProperties_STATUS_ARM struct {
	// CurrentVersionedKeyExpirationTimestamp: This is a read only property that represents the expiration time of the current
	// version of the customer managed key used for encryption.
	CurrentVersionedKeyExpirationTimestamp *string `json:"currentVersionedKeyExpirationTimestamp,omitempty"`

	// CurrentVersionedKeyIdentifier: The object identifier of the current versioned Key Vault Key in use.
	CurrentVersionedKeyIdentifier *string `json:"currentVersionedKeyIdentifier,omitempty"`

	// Keyname: The name of KeyVault key.
	Keyname *string `json:"keyname,omitempty"`

	// Keyvaulturi: The Uri of KeyVault.
	Keyvaulturi *string `json:"keyvaulturi,omitempty"`

	// Keyversion: The version of KeyVault key.
	Keyversion *string `json:"keyversion,omitempty"`

	// LastKeyRotationTimestamp: Timestamp of last rotation of the Key Vault Key.
	LastKeyRotationTimestamp *string `json:"lastKeyRotationTimestamp,omitempty"`
}

// Resource Access Rule.
type ResourceAccessRule_STATUS_ARM struct {
	// ResourceId: Resource Id
	ResourceId *string `json:"resourceId,omitempty"`

	// TenantId: Tenant Id
	TenantId *string `json:"tenantId,omitempty"`
}

// The URIs that are used to perform a retrieval of a public blob, file, web or dfs object via a internet routing endpoint.
type StorageAccountInternetEndpoints_STATUS_ARM struct {
	// Blob: Gets the blob endpoint.
	Blob *string `json:"blob,omitempty"`

	// Dfs: Gets the dfs endpoint.
	Dfs *string `json:"dfs,omitempty"`

	// File: Gets the file endpoint.
	File *string `json:"file,omitempty"`

	// Web: Gets the web endpoint.
	Web *string `json:"web,omitempty"`
}

// The URIs that are used to perform a retrieval of a public blob, queue, table, web or dfs object via a microsoft routing
// endpoint.
type StorageAccountMicrosoftEndpoints_STATUS_ARM struct {
	// Blob: Gets the blob endpoint.
	Blob *string `json:"blob,omitempty"`

	// Dfs: Gets the dfs endpoint.
	Dfs *string `json:"dfs,omitempty"`

	// File: Gets the file endpoint.
	File *string `json:"file,omitempty"`

	// Queue: Gets the queue endpoint.
	Queue *string `json:"queue,omitempty"`

	// Table: Gets the table endpoint.
	Table *string `json:"table,omitempty"`

	// Web: Gets the web endpoint.
	Web *string `json:"web,omitempty"`
}

// Virtual Network rule.
type VirtualNetworkRule_STATUS_ARM struct {
	// Action: The action of virtual network rule.
	Action *VirtualNetworkRule_Action_STATUS `json:"action,omitempty"`

	// Id: Resource ID of a subnet, for example:
	// /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
	Id *string `json:"id,omitempty"`

	// State: Gets the state of virtual network rule.
	State *VirtualNetworkRule_State_STATUS `json:"state,omitempty"`
}

// Blob range
type BlobRestoreRange_STATUS_ARM struct {
	// EndRange: Blob end range. This is exclusive. Empty means account end.
	EndRange *string `json:"endRange,omitempty"`

	// StartRange: Blob start range. This is inclusive. Empty means account start.
	StartRange *string `json:"startRange,omitempty"`
}

// A service that allows server-side encryption to be used.
type EncryptionService_STATUS_ARM struct {
	// Enabled: A boolean indicating whether or not the service encrypts the data as it is stored. Encryption at rest is
	// enabled by default today and cannot be disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// KeyType: Encryption key type to be used for the encryption service. 'Account' key type implies that an account-scoped
	// encryption key will be used. 'Service' key type implies that a default service key is used.
	KeyType *EncryptionService_KeyType_STATUS `json:"keyType,omitempty"`

	// LastEnabledTime: Gets a rough estimate of the date/time when the encryption was last enabled by the user. Data is
	// encrypted at rest by default today and cannot be disabled.
	LastEnabledTime *string `json:"lastEnabledTime,omitempty"`
}
