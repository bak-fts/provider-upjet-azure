/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this RedisCache
func (mg *RedisCache) GetTerraformResourceType() string {
	return "azurerm_redis_cache"
}

// GetConnectionDetailsMapping for this RedisCache
func (tr *RedisCache) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"primary_access_key": "status.atProvider.primaryAccessKey", "primary_connection_string": "status.atProvider.primaryConnectionString", "redis_configuration[*].aof_storage_connection_string_0": "spec.forProvider.redisConfiguration[*].aofStorageConnectionString0SecretRef", "redis_configuration[*].aof_storage_connection_string_1": "spec.forProvider.redisConfiguration[*].aofStorageConnectionString1SecretRef", "redis_configuration[*].rdb_storage_connection_string": "spec.forProvider.redisConfiguration[*].rdbStorageConnectionStringSecretRef", "secondary_access_key": "status.atProvider.secondaryAccessKey", "secondary_connection_string": "status.atProvider.secondaryConnectionString"}
}

// GetObservation of this RedisCache
func (tr *RedisCache) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RedisCache
func (tr *RedisCache) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RedisCache
func (tr *RedisCache) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RedisCache
func (tr *RedisCache) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RedisCache
func (tr *RedisCache) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RedisCache using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RedisCache) LateInitialize(attrs []byte) (bool, error) {
	params := &RedisCacheParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RedisCache) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this RedisEnterpriseCluster
func (mg *RedisEnterpriseCluster) GetTerraformResourceType() string {
	return "azurerm_redis_enterprise_cluster"
}

// GetConnectionDetailsMapping for this RedisEnterpriseCluster
func (tr *RedisEnterpriseCluster) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this RedisEnterpriseCluster
func (tr *RedisEnterpriseCluster) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RedisEnterpriseCluster
func (tr *RedisEnterpriseCluster) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RedisEnterpriseCluster
func (tr *RedisEnterpriseCluster) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RedisEnterpriseCluster
func (tr *RedisEnterpriseCluster) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RedisEnterpriseCluster
func (tr *RedisEnterpriseCluster) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RedisEnterpriseCluster using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RedisEnterpriseCluster) LateInitialize(attrs []byte) (bool, error) {
	params := &RedisEnterpriseClusterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RedisEnterpriseCluster) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this RedisEnterpriseDatabase
func (mg *RedisEnterpriseDatabase) GetTerraformResourceType() string {
	return "azurerm_redis_enterprise_database"
}

// GetConnectionDetailsMapping for this RedisEnterpriseDatabase
func (tr *RedisEnterpriseDatabase) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"primary_access_key": "status.atProvider.primaryAccessKey", "secondary_access_key": "status.atProvider.secondaryAccessKey"}
}

// GetObservation of this RedisEnterpriseDatabase
func (tr *RedisEnterpriseDatabase) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RedisEnterpriseDatabase
func (tr *RedisEnterpriseDatabase) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RedisEnterpriseDatabase
func (tr *RedisEnterpriseDatabase) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RedisEnterpriseDatabase
func (tr *RedisEnterpriseDatabase) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RedisEnterpriseDatabase
func (tr *RedisEnterpriseDatabase) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RedisEnterpriseDatabase using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RedisEnterpriseDatabase) LateInitialize(attrs []byte) (bool, error) {
	params := &RedisEnterpriseDatabaseParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RedisEnterpriseDatabase) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this RedisFirewallRule
func (mg *RedisFirewallRule) GetTerraformResourceType() string {
	return "azurerm_redis_firewall_rule"
}

// GetConnectionDetailsMapping for this RedisFirewallRule
func (tr *RedisFirewallRule) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this RedisFirewallRule
func (tr *RedisFirewallRule) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RedisFirewallRule
func (tr *RedisFirewallRule) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RedisFirewallRule
func (tr *RedisFirewallRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RedisFirewallRule
func (tr *RedisFirewallRule) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RedisFirewallRule
func (tr *RedisFirewallRule) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RedisFirewallRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RedisFirewallRule) LateInitialize(attrs []byte) (bool, error) {
	params := &RedisFirewallRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RedisFirewallRule) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this RedisLinkedServer
func (mg *RedisLinkedServer) GetTerraformResourceType() string {
	return "azurerm_redis_linked_server"
}

// GetConnectionDetailsMapping for this RedisLinkedServer
func (tr *RedisLinkedServer) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this RedisLinkedServer
func (tr *RedisLinkedServer) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RedisLinkedServer
func (tr *RedisLinkedServer) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RedisLinkedServer
func (tr *RedisLinkedServer) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RedisLinkedServer
func (tr *RedisLinkedServer) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RedisLinkedServer
func (tr *RedisLinkedServer) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RedisLinkedServer using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RedisLinkedServer) LateInitialize(attrs []byte) (bool, error) {
	params := &RedisLinkedServerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RedisLinkedServer) GetTerraformSchemaVersion() int {
	return 1
}
