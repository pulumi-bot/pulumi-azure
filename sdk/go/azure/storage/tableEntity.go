// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Entity within a Table in an Azure Storage Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_table_entity.html.markdown.
type TableEntity struct {
	s *pulumi.ResourceState
}

// NewTableEntity registers a new resource with the given unique name, arguments, and options.
func NewTableEntity(ctx *pulumi.Context,
	name string, args *TableEntityArgs, opts ...pulumi.ResourceOpt) (*TableEntity, error) {
	if args == nil || args.Entity == nil {
		return nil, errors.New("missing required argument 'Entity'")
	}
	if args == nil || args.PartitionKey == nil {
		return nil, errors.New("missing required argument 'PartitionKey'")
	}
	if args == nil || args.RowKey == nil {
		return nil, errors.New("missing required argument 'RowKey'")
	}
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	if args == nil || args.TableName == nil {
		return nil, errors.New("missing required argument 'TableName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["entity"] = nil
		inputs["partitionKey"] = nil
		inputs["rowKey"] = nil
		inputs["storageAccountName"] = nil
		inputs["tableName"] = nil
	} else {
		inputs["entity"] = args.Entity
		inputs["partitionKey"] = args.PartitionKey
		inputs["rowKey"] = args.RowKey
		inputs["storageAccountName"] = args.StorageAccountName
		inputs["tableName"] = args.TableName
	}
	s, err := ctx.RegisterResource("azure:storage/tableEntity:TableEntity", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TableEntity{s: s}, nil
}

// GetTableEntity gets an existing TableEntity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableEntity(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TableEntityState, opts ...pulumi.ResourceOpt) (*TableEntity, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["entity"] = state.Entity
		inputs["partitionKey"] = state.PartitionKey
		inputs["rowKey"] = state.RowKey
		inputs["storageAccountName"] = state.StorageAccountName
		inputs["tableName"] = state.TableName
	}
	s, err := ctx.ReadResource("azure:storage/tableEntity:TableEntity", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TableEntity{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TableEntity) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TableEntity) ID() pulumi.IDOutput {
	return r.s.ID()
}

// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
func (r *TableEntity) Entity() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["entity"])
}

// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
func (r *TableEntity) PartitionKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["partitionKey"])
}

// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
func (r *TableEntity) RowKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["rowKey"])
}

// Specifies the storage account in which to create the storage table entity.
// Changing this forces a new resource to be created.
func (r *TableEntity) StorageAccountName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["storageAccountName"])
}

// The name of the storage table in which to create the storage table entity. 
// Changing this forces a new resource to be created.
func (r *TableEntity) TableName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["tableName"])
}

// Input properties used for looking up and filtering TableEntity resources.
type TableEntityState struct {
	// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
	Entity interface{}
	// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
	PartitionKey interface{}
	// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
	RowKey interface{}
	// Specifies the storage account in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	StorageAccountName interface{}
	// The name of the storage table in which to create the storage table entity. 
	// Changing this forces a new resource to be created.
	TableName interface{}
}

// The set of arguments for constructing a TableEntity resource.
type TableEntityArgs struct {
	// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
	Entity interface{}
	// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
	PartitionKey interface{}
	// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
	RowKey interface{}
	// Specifies the storage account in which to create the storage table entity.
	// Changing this forces a new resource to be created.
	StorageAccountName interface{}
	// The name of the storage table in which to create the storage table entity. 
	// Changing this forces a new resource to be created.
	TableName interface{}
}