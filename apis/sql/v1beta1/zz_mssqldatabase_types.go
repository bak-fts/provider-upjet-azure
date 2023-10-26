// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ImportInitParameters struct {

	// Specifies the name of the SQL administrator.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// Specifies the type of authentication used to access the server. Valid values are SQL or ADPassword.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The resource id for the storage account used to store BACPAC file. If set, private endpoint connection will be created for the storage account. Must match storage account used for storage_uri parameter.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Specifies the type of access key for the storage account. Valid values are StorageAccessKey or SharedAccessKey.
	StorageKeyType *string `json:"storageKeyType,omitempty" tf:"storage_key_type,omitempty"`

	// Specifies the blob URI of the .bacpac file.
	StorageURI *string `json:"storageUri,omitempty" tf:"storage_uri,omitempty"`
}

type ImportObservation struct {

	// Specifies the name of the SQL administrator.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// Specifies the type of authentication used to access the server. Valid values are SQL or ADPassword.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The resource id for the storage account used to store BACPAC file. If set, private endpoint connection will be created for the storage account. Must match storage account used for storage_uri parameter.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Specifies the type of access key for the storage account. Valid values are StorageAccessKey or SharedAccessKey.
	StorageKeyType *string `json:"storageKeyType,omitempty" tf:"storage_key_type,omitempty"`

	// Specifies the blob URI of the .bacpac file.
	StorageURI *string `json:"storageUri,omitempty" tf:"storage_uri,omitempty"`
}

type ImportParameters struct {

	// Specifies the name of the SQL administrator.
	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin" tf:"administrator_login,omitempty"`

	// Specifies the password of the SQL administrator.
	// +kubebuilder:validation:Required
	AdministratorLoginPasswordSecretRef v1.SecretKeySelector `json:"administratorLoginPasswordSecretRef" tf:"-"`

	// Specifies the type of authentication used to access the server. Valid values are SQL or ADPassword.
	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType" tf:"authentication_type,omitempty"`

	// The resource id for the storage account used to store BACPAC file. If set, private endpoint connection will be created for the storage account. Must match storage account used for storage_uri parameter.
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Specifies the access key for the storage account.
	// +kubebuilder:validation:Required
	StorageKeySecretRef v1.SecretKeySelector `json:"storageKeySecretRef" tf:"-"`

	// Specifies the type of access key for the storage account. Valid values are StorageAccessKey or SharedAccessKey.
	// +kubebuilder:validation:Optional
	StorageKeyType *string `json:"storageKeyType" tf:"storage_key_type,omitempty"`

	// Specifies the blob URI of the .bacpac file.
	// +kubebuilder:validation:Optional
	StorageURI *string `json:"storageUri" tf:"storage_uri,omitempty"`
}

type LongTermRetentionPolicyInitParameters struct {

	// The monthly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 120 months. e.g. P1Y, P1M, P4W or P30D.
	MonthlyRetention *string `json:"monthlyRetention,omitempty" tf:"monthly_retention,omitempty"`

	// The week of year to take the yearly backup. Value has to be between 1 and 52.
	WeekOfYear *float64 `json:"weekOfYear,omitempty" tf:"week_of_year,omitempty"`

	// The weekly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 520 weeks. e.g. P1Y, P1M, P1W or P7D.
	WeeklyRetention *string `json:"weeklyRetention,omitempty" tf:"weekly_retention,omitempty"`

	// The yearly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 10 years. e.g. P1Y, P12M, P52W or P365D.
	YearlyRetention *string `json:"yearlyRetention,omitempty" tf:"yearly_retention,omitempty"`
}

type LongTermRetentionPolicyObservation struct {

	// The monthly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 120 months. e.g. P1Y, P1M, P4W or P30D.
	MonthlyRetention *string `json:"monthlyRetention,omitempty" tf:"monthly_retention,omitempty"`

	// The week of year to take the yearly backup. Value has to be between 1 and 52.
	WeekOfYear *float64 `json:"weekOfYear,omitempty" tf:"week_of_year,omitempty"`

	// The weekly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 520 weeks. e.g. P1Y, P1M, P1W or P7D.
	WeeklyRetention *string `json:"weeklyRetention,omitempty" tf:"weekly_retention,omitempty"`

	// The yearly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 10 years. e.g. P1Y, P12M, P52W or P365D.
	YearlyRetention *string `json:"yearlyRetention,omitempty" tf:"yearly_retention,omitempty"`
}

type LongTermRetentionPolicyParameters struct {

	// The monthly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 120 months. e.g. P1Y, P1M, P4W or P30D.
	// +kubebuilder:validation:Optional
	MonthlyRetention *string `json:"monthlyRetention,omitempty" tf:"monthly_retention,omitempty"`

	// The week of year to take the yearly backup. Value has to be between 1 and 52.
	// +kubebuilder:validation:Optional
	WeekOfYear *float64 `json:"weekOfYear,omitempty" tf:"week_of_year,omitempty"`

	// The weekly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 520 weeks. e.g. P1Y, P1M, P1W or P7D.
	// +kubebuilder:validation:Optional
	WeeklyRetention *string `json:"weeklyRetention,omitempty" tf:"weekly_retention,omitempty"`

	// The yearly retention policy for an LTR backup in an ISO 8601 format. Valid value is between 1 to 10 years. e.g. P1Y, P12M, P52W or P365D.
	// +kubebuilder:validation:Optional
	YearlyRetention *string `json:"yearlyRetention,omitempty" tf:"yearly_retention,omitempty"`
}

type MSSQLDatabaseInitParameters struct {

	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled. This property is only settable for General Purpose Serverless databases.
	AutoPauseDelayInMinutes *float64 `json:"autoPauseDelayInMinutes,omitempty" tf:"auto_pause_delay_in_minutes,omitempty"`

	// Specifies the collation of the database. Changing this forces a new resource to be created.
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// The create mode of the database. Possible values are Copy, Default, OnlineSecondary, PointInTimeRestore, Recovery, Restore, RestoreExternalBackup, RestoreExternalBackupSecondary, RestoreLongTermRetentionBackup and Secondary. Mutually exclusive with import. Changing this forces a new resource to be created.
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// The ID of the source database from which to create the new database. This should only be used for databases with create_mode values that use another database as reference. Changing this forces a new resource to be created.
	CreationSourceDatabaseID *string `json:"creationSourceDatabaseId,omitempty" tf:"creation_source_database_id,omitempty"`

	// Specifies the ID of the elastic pool containing this database.
	ElasticPoolID *string `json:"elasticPoolId,omitempty" tf:"elastic_pool_id,omitempty"`

	// A boolean that specifies if the Geo Backup Policy is enabled. Defaults to true.
	GeoBackupEnabled *bool `json:"geoBackupEnabled,omitempty" tf:"geo_backup_enabled,omitempty"`

	// A Database Import block as documented below. Mutually exclusive with create_mode.
	Import []ImportInitParameters `json:"import,omitempty" tf:"import,omitempty"`

	// A boolean that specifies if this is a ledger database. Defaults to false. Changing this forces a new resource to be created.
	LedgerEnabled *bool `json:"ledgerEnabled,omitempty" tf:"ledger_enabled,omitempty"`

	// Specifies the license type applied to this database. Possible values are LicenseIncluded and BasePrice.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// A long_term_retention_policy block as defined below.
	LongTermRetentionPolicy []LongTermRetentionPolicyInitParameters `json:"longTermRetentionPolicy,omitempty" tf:"long_term_retention_policy,omitempty"`

	// The name of the Public Maintenance Configuration window to apply to the database. Valid values include SQL_Default, SQL_EastUS_DB_1, SQL_EastUS2_DB_1, SQL_SoutheastAsia_DB_1, SQL_AustraliaEast_DB_1, SQL_NorthEurope_DB_1, SQL_SouthCentralUS_DB_1, SQL_WestUS2_DB_1, SQL_UKSouth_DB_1, SQL_WestEurope_DB_1, SQL_EastUS_DB_2, SQL_EastUS2_DB_2, SQL_WestUS2_DB_2, SQL_SoutheastAsia_DB_2, SQL_AustraliaEast_DB_2, SQL_NorthEurope_DB_2, SQL_SouthCentralUS_DB_2, SQL_UKSouth_DB_2, SQL_WestEurope_DB_2, SQL_AustraliaSoutheast_DB_1, SQL_BrazilSouth_DB_1, SQL_CanadaCentral_DB_1, SQL_CanadaEast_DB_1, SQL_CentralUS_DB_1, SQL_EastAsia_DB_1, SQL_FranceCentral_DB_1, SQL_GermanyWestCentral_DB_1, SQL_CentralIndia_DB_1, SQL_SouthIndia_DB_1, SQL_JapanEast_DB_1, SQL_JapanWest_DB_1, SQL_NorthCentralUS_DB_1, SQL_UKWest_DB_1, SQL_WestUS_DB_1, SQL_AustraliaSoutheast_DB_2, SQL_BrazilSouth_DB_2, SQL_CanadaCentral_DB_2, SQL_CanadaEast_DB_2, SQL_CentralUS_DB_2, SQL_EastAsia_DB_2, SQL_FranceCentral_DB_2, SQL_GermanyWestCentral_DB_2, SQL_CentralIndia_DB_2, SQL_SouthIndia_DB_2, SQL_JapanEast_DB_2, SQL_JapanWest_DB_2, SQL_NorthCentralUS_DB_2, SQL_UKWest_DB_2, SQL_WestUS_DB_2, SQL_WestCentralUS_DB_1, SQL_FranceSouth_DB_1, SQL_WestCentralUS_DB_2, SQL_FranceSouth_DB_2, SQL_SwitzerlandNorth_DB_1, SQL_SwitzerlandNorth_DB_2, SQL_BrazilSoutheast_DB_1, SQL_UAENorth_DB_1, SQL_BrazilSoutheast_DB_2, SQL_UAENorth_DB_2. Defaults to SQL_Default.
	MaintenanceConfigurationName *string `json:"maintenanceConfigurationName,omitempty" tf:"maintenance_configuration_name,omitempty"`

	// The max size of the database in gigabytes.
	MaxSizeGb *float64 `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`

	// Minimal capacity that database will always have allocated, if not paused. This property is only settable for General Purpose Serverless databases.
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
	ReadReplicaCount *float64 `json:"readReplicaCount,omitempty" tf:"read_replica_count,omitempty"`

	// If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
	ReadScale *bool `json:"readScale,omitempty" tf:"read_scale,omitempty"`

	// The ID of the database to be recovered. This property is only applicable when the create_mode is Recovery.
	RecoverDatabaseID *string `json:"recoverDatabaseId,omitempty" tf:"recover_database_id,omitempty"`

	// The ID of the database to be restored. This property is only applicable when the create_mode is Restore.
	RestoreDroppedDatabaseID *string `json:"restoreDroppedDatabaseId,omitempty" tf:"restore_dropped_database_id,omitempty"`

	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. This property is only settable for create_mode= PointInTimeRestore databases.
	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time,omitempty"`

	// Specifies the name of the sample schema to apply when creating this database. Possible value is AdventureWorksLT.
	SampleName *string `json:"sampleName,omitempty" tf:"sample_name,omitempty"`

	// A short_term_retention_policy block as defined below.
	ShortTermRetentionPolicy []ShortTermRetentionPolicyInitParameters `json:"shortTermRetentionPolicy,omitempty" tf:"short_term_retention_policy,omitempty"`

	// Specifies the name of the SKU used by the database. For example, GP_S_Gen5_2,HS_Gen4_1,BC_Gen5_2, ElasticPool, Basic,S0, P2 ,DW100c, DS100. Changing this from the HyperScale service tier to another service tier will create a new resource.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies the storage account type used to store backups for this database. Possible values are Geo, Local and Zone. The default value is Geo.
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Threat detection policy configuration. The threat_detection_policy block supports fields documented below.
	ThreatDetectionPolicy []ThreatDetectionPolicyInitParameters `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy,omitempty"`

	// If set to true, Transparent Data Encryption will be enabled on the database. Defaults to true.
	TransparentDataEncryptionEnabled *bool `json:"transparentDataEncryptionEnabled,omitempty" tf:"transparent_data_encryption_enabled,omitempty"`

	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones. This property is only settable for Premium and Business Critical databases.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type MSSQLDatabaseObservation struct {

	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled. This property is only settable for General Purpose Serverless databases.
	AutoPauseDelayInMinutes *float64 `json:"autoPauseDelayInMinutes,omitempty" tf:"auto_pause_delay_in_minutes,omitempty"`

	// Specifies the collation of the database. Changing this forces a new resource to be created.
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// The create mode of the database. Possible values are Copy, Default, OnlineSecondary, PointInTimeRestore, Recovery, Restore, RestoreExternalBackup, RestoreExternalBackupSecondary, RestoreLongTermRetentionBackup and Secondary. Mutually exclusive with import. Changing this forces a new resource to be created.
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// The ID of the source database from which to create the new database. This should only be used for databases with create_mode values that use another database as reference. Changing this forces a new resource to be created.
	CreationSourceDatabaseID *string `json:"creationSourceDatabaseId,omitempty" tf:"creation_source_database_id,omitempty"`

	// Specifies the ID of the elastic pool containing this database.
	ElasticPoolID *string `json:"elasticPoolId,omitempty" tf:"elastic_pool_id,omitempty"`

	// A boolean that specifies if the Geo Backup Policy is enabled. Defaults to true.
	GeoBackupEnabled *bool `json:"geoBackupEnabled,omitempty" tf:"geo_backup_enabled,omitempty"`

	// The ID of the MS SQL Database.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A Database Import block as documented below. Mutually exclusive with create_mode.
	Import []ImportObservation `json:"import,omitempty" tf:"import,omitempty"`

	// A boolean that specifies if this is a ledger database. Defaults to false. Changing this forces a new resource to be created.
	LedgerEnabled *bool `json:"ledgerEnabled,omitempty" tf:"ledger_enabled,omitempty"`

	// Specifies the license type applied to this database. Possible values are LicenseIncluded and BasePrice.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// A long_term_retention_policy block as defined below.
	LongTermRetentionPolicy []LongTermRetentionPolicyObservation `json:"longTermRetentionPolicy,omitempty" tf:"long_term_retention_policy,omitempty"`

	// The name of the Public Maintenance Configuration window to apply to the database. Valid values include SQL_Default, SQL_EastUS_DB_1, SQL_EastUS2_DB_1, SQL_SoutheastAsia_DB_1, SQL_AustraliaEast_DB_1, SQL_NorthEurope_DB_1, SQL_SouthCentralUS_DB_1, SQL_WestUS2_DB_1, SQL_UKSouth_DB_1, SQL_WestEurope_DB_1, SQL_EastUS_DB_2, SQL_EastUS2_DB_2, SQL_WestUS2_DB_2, SQL_SoutheastAsia_DB_2, SQL_AustraliaEast_DB_2, SQL_NorthEurope_DB_2, SQL_SouthCentralUS_DB_2, SQL_UKSouth_DB_2, SQL_WestEurope_DB_2, SQL_AustraliaSoutheast_DB_1, SQL_BrazilSouth_DB_1, SQL_CanadaCentral_DB_1, SQL_CanadaEast_DB_1, SQL_CentralUS_DB_1, SQL_EastAsia_DB_1, SQL_FranceCentral_DB_1, SQL_GermanyWestCentral_DB_1, SQL_CentralIndia_DB_1, SQL_SouthIndia_DB_1, SQL_JapanEast_DB_1, SQL_JapanWest_DB_1, SQL_NorthCentralUS_DB_1, SQL_UKWest_DB_1, SQL_WestUS_DB_1, SQL_AustraliaSoutheast_DB_2, SQL_BrazilSouth_DB_2, SQL_CanadaCentral_DB_2, SQL_CanadaEast_DB_2, SQL_CentralUS_DB_2, SQL_EastAsia_DB_2, SQL_FranceCentral_DB_2, SQL_GermanyWestCentral_DB_2, SQL_CentralIndia_DB_2, SQL_SouthIndia_DB_2, SQL_JapanEast_DB_2, SQL_JapanWest_DB_2, SQL_NorthCentralUS_DB_2, SQL_UKWest_DB_2, SQL_WestUS_DB_2, SQL_WestCentralUS_DB_1, SQL_FranceSouth_DB_1, SQL_WestCentralUS_DB_2, SQL_FranceSouth_DB_2, SQL_SwitzerlandNorth_DB_1, SQL_SwitzerlandNorth_DB_2, SQL_BrazilSoutheast_DB_1, SQL_UAENorth_DB_1, SQL_BrazilSoutheast_DB_2, SQL_UAENorth_DB_2. Defaults to SQL_Default.
	MaintenanceConfigurationName *string `json:"maintenanceConfigurationName,omitempty" tf:"maintenance_configuration_name,omitempty"`

	// The max size of the database in gigabytes.
	MaxSizeGb *float64 `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`

	// Minimal capacity that database will always have allocated, if not paused. This property is only settable for General Purpose Serverless databases.
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
	ReadReplicaCount *float64 `json:"readReplicaCount,omitempty" tf:"read_replica_count,omitempty"`

	// If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
	ReadScale *bool `json:"readScale,omitempty" tf:"read_scale,omitempty"`

	// The ID of the database to be recovered. This property is only applicable when the create_mode is Recovery.
	RecoverDatabaseID *string `json:"recoverDatabaseId,omitempty" tf:"recover_database_id,omitempty"`

	// The ID of the database to be restored. This property is only applicable when the create_mode is Restore.
	RestoreDroppedDatabaseID *string `json:"restoreDroppedDatabaseId,omitempty" tf:"restore_dropped_database_id,omitempty"`

	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. This property is only settable for create_mode= PointInTimeRestore databases.
	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time,omitempty"`

	// Specifies the name of the sample schema to apply when creating this database. Possible value is AdventureWorksLT.
	SampleName *string `json:"sampleName,omitempty" tf:"sample_name,omitempty"`

	// The id of the MS SQL Server on which to create the database. Changing this forces a new resource to be created.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// A short_term_retention_policy block as defined below.
	ShortTermRetentionPolicy []ShortTermRetentionPolicyObservation `json:"shortTermRetentionPolicy,omitempty" tf:"short_term_retention_policy,omitempty"`

	// Specifies the name of the SKU used by the database. For example, GP_S_Gen5_2,HS_Gen4_1,BC_Gen5_2, ElasticPool, Basic,S0, P2 ,DW100c, DS100. Changing this from the HyperScale service tier to another service tier will create a new resource.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies the storage account type used to store backups for this database. Possible values are Geo, Local and Zone. The default value is Geo.
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Threat detection policy configuration. The threat_detection_policy block supports fields documented below.
	ThreatDetectionPolicy []ThreatDetectionPolicyObservation `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy,omitempty"`

	// If set to true, Transparent Data Encryption will be enabled on the database. Defaults to true.
	TransparentDataEncryptionEnabled *bool `json:"transparentDataEncryptionEnabled,omitempty" tf:"transparent_data_encryption_enabled,omitempty"`

	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones. This property is only settable for Premium and Business Critical databases.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type MSSQLDatabaseParameters struct {

	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled. This property is only settable for General Purpose Serverless databases.
	// +kubebuilder:validation:Optional
	AutoPauseDelayInMinutes *float64 `json:"autoPauseDelayInMinutes,omitempty" tf:"auto_pause_delay_in_minutes,omitempty"`

	// Specifies the collation of the database. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// The create mode of the database. Possible values are Copy, Default, OnlineSecondary, PointInTimeRestore, Recovery, Restore, RestoreExternalBackup, RestoreExternalBackupSecondary, RestoreLongTermRetentionBackup and Secondary. Mutually exclusive with import. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// The ID of the source database from which to create the new database. This should only be used for databases with create_mode values that use another database as reference. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CreationSourceDatabaseID *string `json:"creationSourceDatabaseId,omitempty" tf:"creation_source_database_id,omitempty"`

	// Specifies the ID of the elastic pool containing this database.
	// +kubebuilder:validation:Optional
	ElasticPoolID *string `json:"elasticPoolId,omitempty" tf:"elastic_pool_id,omitempty"`

	// A boolean that specifies if the Geo Backup Policy is enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	GeoBackupEnabled *bool `json:"geoBackupEnabled,omitempty" tf:"geo_backup_enabled,omitempty"`

	// A Database Import block as documented below. Mutually exclusive with create_mode.
	// +kubebuilder:validation:Optional
	Import []ImportParameters `json:"import,omitempty" tf:"import,omitempty"`

	// A boolean that specifies if this is a ledger database. Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	LedgerEnabled *bool `json:"ledgerEnabled,omitempty" tf:"ledger_enabled,omitempty"`

	// Specifies the license type applied to this database. Possible values are LicenseIncluded and BasePrice.
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// A long_term_retention_policy block as defined below.
	// +kubebuilder:validation:Optional
	LongTermRetentionPolicy []LongTermRetentionPolicyParameters `json:"longTermRetentionPolicy,omitempty" tf:"long_term_retention_policy,omitempty"`

	// The name of the Public Maintenance Configuration window to apply to the database. Valid values include SQL_Default, SQL_EastUS_DB_1, SQL_EastUS2_DB_1, SQL_SoutheastAsia_DB_1, SQL_AustraliaEast_DB_1, SQL_NorthEurope_DB_1, SQL_SouthCentralUS_DB_1, SQL_WestUS2_DB_1, SQL_UKSouth_DB_1, SQL_WestEurope_DB_1, SQL_EastUS_DB_2, SQL_EastUS2_DB_2, SQL_WestUS2_DB_2, SQL_SoutheastAsia_DB_2, SQL_AustraliaEast_DB_2, SQL_NorthEurope_DB_2, SQL_SouthCentralUS_DB_2, SQL_UKSouth_DB_2, SQL_WestEurope_DB_2, SQL_AustraliaSoutheast_DB_1, SQL_BrazilSouth_DB_1, SQL_CanadaCentral_DB_1, SQL_CanadaEast_DB_1, SQL_CentralUS_DB_1, SQL_EastAsia_DB_1, SQL_FranceCentral_DB_1, SQL_GermanyWestCentral_DB_1, SQL_CentralIndia_DB_1, SQL_SouthIndia_DB_1, SQL_JapanEast_DB_1, SQL_JapanWest_DB_1, SQL_NorthCentralUS_DB_1, SQL_UKWest_DB_1, SQL_WestUS_DB_1, SQL_AustraliaSoutheast_DB_2, SQL_BrazilSouth_DB_2, SQL_CanadaCentral_DB_2, SQL_CanadaEast_DB_2, SQL_CentralUS_DB_2, SQL_EastAsia_DB_2, SQL_FranceCentral_DB_2, SQL_GermanyWestCentral_DB_2, SQL_CentralIndia_DB_2, SQL_SouthIndia_DB_2, SQL_JapanEast_DB_2, SQL_JapanWest_DB_2, SQL_NorthCentralUS_DB_2, SQL_UKWest_DB_2, SQL_WestUS_DB_2, SQL_WestCentralUS_DB_1, SQL_FranceSouth_DB_1, SQL_WestCentralUS_DB_2, SQL_FranceSouth_DB_2, SQL_SwitzerlandNorth_DB_1, SQL_SwitzerlandNorth_DB_2, SQL_BrazilSoutheast_DB_1, SQL_UAENorth_DB_1, SQL_BrazilSoutheast_DB_2, SQL_UAENorth_DB_2. Defaults to SQL_Default.
	// +kubebuilder:validation:Optional
	MaintenanceConfigurationName *string `json:"maintenanceConfigurationName,omitempty" tf:"maintenance_configuration_name,omitempty"`

	// The max size of the database in gigabytes.
	// +kubebuilder:validation:Optional
	MaxSizeGb *float64 `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`

	// Minimal capacity that database will always have allocated, if not paused. This property is only settable for General Purpose Serverless databases.
	// +kubebuilder:validation:Optional
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
	// +kubebuilder:validation:Optional
	ReadReplicaCount *float64 `json:"readReplicaCount,omitempty" tf:"read_replica_count,omitempty"`

	// If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
	// +kubebuilder:validation:Optional
	ReadScale *bool `json:"readScale,omitempty" tf:"read_scale,omitempty"`

	// The ID of the database to be recovered. This property is only applicable when the create_mode is Recovery.
	// +kubebuilder:validation:Optional
	RecoverDatabaseID *string `json:"recoverDatabaseId,omitempty" tf:"recover_database_id,omitempty"`

	// The ID of the database to be restored. This property is only applicable when the create_mode is Restore.
	// +kubebuilder:validation:Optional
	RestoreDroppedDatabaseID *string `json:"restoreDroppedDatabaseId,omitempty" tf:"restore_dropped_database_id,omitempty"`

	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. This property is only settable for create_mode= PointInTimeRestore databases.
	// +kubebuilder:validation:Optional
	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time,omitempty"`

	// Specifies the name of the sample schema to apply when creating this database. Possible value is AdventureWorksLT.
	// +kubebuilder:validation:Optional
	SampleName *string `json:"sampleName,omitempty" tf:"sample_name,omitempty"`

	// The id of the MS SQL Server on which to create the database. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLServer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a MSSQLServer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// A short_term_retention_policy block as defined below.
	// +kubebuilder:validation:Optional
	ShortTermRetentionPolicy []ShortTermRetentionPolicyParameters `json:"shortTermRetentionPolicy,omitempty" tf:"short_term_retention_policy,omitempty"`

	// Specifies the name of the SKU used by the database. For example, GP_S_Gen5_2,HS_Gen4_1,BC_Gen5_2, ElasticPool, Basic,S0, P2 ,DW100c, DS100. Changing this from the HyperScale service tier to another service tier will create a new resource.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies the storage account type used to store backups for this database. Possible values are Geo, Local and Zone. The default value is Geo.
	// +kubebuilder:validation:Optional
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Threat detection policy configuration. The threat_detection_policy block supports fields documented below.
	// +kubebuilder:validation:Optional
	ThreatDetectionPolicy []ThreatDetectionPolicyParameters `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy,omitempty"`

	// If set to true, Transparent Data Encryption will be enabled on the database. Defaults to true.
	// +kubebuilder:validation:Optional
	TransparentDataEncryptionEnabled *bool `json:"transparentDataEncryptionEnabled,omitempty" tf:"transparent_data_encryption_enabled,omitempty"`

	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones. This property is only settable for Premium and Business Critical databases.
	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type ShortTermRetentionPolicyInitParameters struct {

	// The hours between each differential backup. This is only applicable to live databases but not dropped databases. Value has to be 12 or 24. Defaults to 12 hours.
	BackupIntervalInHours *float64 `json:"backupIntervalInHours,omitempty" tf:"backup_interval_in_hours,omitempty"`

	// Point In Time Restore configuration. Value has to be between 7 and 35.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
}

type ShortTermRetentionPolicyObservation struct {

	// The hours between each differential backup. This is only applicable to live databases but not dropped databases. Value has to be 12 or 24. Defaults to 12 hours.
	BackupIntervalInHours *float64 `json:"backupIntervalInHours,omitempty" tf:"backup_interval_in_hours,omitempty"`

	// Point In Time Restore configuration. Value has to be between 7 and 35.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
}

type ShortTermRetentionPolicyParameters struct {

	// The hours between each differential backup. This is only applicable to live databases but not dropped databases. Value has to be 12 or 24. Defaults to 12 hours.
	// +kubebuilder:validation:Optional
	BackupIntervalInHours *float64 `json:"backupIntervalInHours,omitempty" tf:"backup_interval_in_hours,omitempty"`

	// Point In Time Restore configuration. Value has to be between 7 and 35.
	// +kubebuilder:validation:Optional
	RetentionDays *float64 `json:"retentionDays" tf:"retention_days,omitempty"`
}

type ThreatDetectionPolicyInitParameters struct {

	// Specifies a list of alerts which should be disabled. Possible values include Access_Anomaly, Sql_Injection and Sql_Injection_Vulnerability.
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// Should the account administrators be emailed when this alert is triggered? Possible values are Disabled and Enabled.
	EmailAccountAdmins *string `json:"emailAccountAdmins,omitempty" tf:"email_account_admins,omitempty"`

	// A list of email addresses which alerts should be sent to.
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// The State of the Policy. Possible values are Enabled, Disabled or New.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Specifies the blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. Required if state is Enabled.
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

type ThreatDetectionPolicyObservation struct {

	// Specifies a list of alerts which should be disabled. Possible values include Access_Anomaly, Sql_Injection and Sql_Injection_Vulnerability.
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// Should the account administrators be emailed when this alert is triggered? Possible values are Disabled and Enabled.
	EmailAccountAdmins *string `json:"emailAccountAdmins,omitempty" tf:"email_account_admins,omitempty"`

	// A list of email addresses which alerts should be sent to.
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// The State of the Policy. Possible values are Enabled, Disabled or New.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Specifies the blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. Required if state is Enabled.
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

type ThreatDetectionPolicyParameters struct {

	// Specifies a list of alerts which should be disabled. Possible values include Access_Anomaly, Sql_Injection and Sql_Injection_Vulnerability.
	// +kubebuilder:validation:Optional
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// Should the account administrators be emailed when this alert is triggered? Possible values are Disabled and Enabled.
	// +kubebuilder:validation:Optional
	EmailAccountAdmins *string `json:"emailAccountAdmins,omitempty" tf:"email_account_admins,omitempty"`

	// A list of email addresses which alerts should be sent to.
	// +kubebuilder:validation:Optional
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// Specifies the number of days to keep in the Threat Detection audit logs.
	// +kubebuilder:validation:Optional
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// The State of the Policy. Possible values are Enabled, Disabled or New.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Specifies the identifier key of the Threat Detection audit storage account. Required if state is Enabled.
	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// Specifies the blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. Required if state is Enabled.
	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

// MSSQLDatabaseSpec defines the desired state of MSSQLDatabase
type MSSQLDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLDatabaseParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MSSQLDatabaseInitParameters `json:"initProvider,omitempty"`
}

// MSSQLDatabaseStatus defines the observed state of MSSQLDatabase.
type MSSQLDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLDatabase is the Schema for the MSSQLDatabases API. Manages a MS SQL Database.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLDatabaseSpec   `json:"spec"`
	Status            MSSQLDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLDatabaseList contains a list of MSSQLDatabases
type MSSQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLDatabase `json:"items"`
}

// Repository type metadata.
var (
	MSSQLDatabase_Kind             = "MSSQLDatabase"
	MSSQLDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLDatabase_Kind}.String()
	MSSQLDatabase_KindAPIVersion   = MSSQLDatabase_Kind + "." + CRDGroupVersion.String()
	MSSQLDatabase_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLDatabase{}, &MSSQLDatabaseList{})
}
