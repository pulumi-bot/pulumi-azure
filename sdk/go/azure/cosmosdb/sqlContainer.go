// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a SQL Container within a Cosmos DB Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/cosmosdb_sql_container.html.markdown.
type SqlContainer struct {
	s *pulumi.ResourceState
}

// NewSqlContainer registers a new resource with the given unique name, arguments, and options.
func NewSqlContainer(ctx *pulumi.Context,
	name string, args *SqlContainerArgs, opts ...pulumi.ResourceOpt) (*SqlContainer, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accountName"] = nil
		inputs["databaseName"] = nil
		inputs["name"] = nil
		inputs["partitionKeyPath"] = nil
		inputs["resourceGroupName"] = nil
		inputs["uniqueKeys"] = nil
	} else {
		inputs["accountName"] = args.AccountName
		inputs["databaseName"] = args.DatabaseName
		inputs["name"] = args.Name
		inputs["partitionKeyPath"] = args.PartitionKeyPath
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["uniqueKeys"] = args.UniqueKeys
	}
	s, err := ctx.RegisterResource("azure:cosmosdb/sqlContainer:SqlContainer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SqlContainer{s: s}, nil
}

// GetSqlContainer gets an existing SqlContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlContainer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SqlContainerState, opts ...pulumi.ResourceOpt) (*SqlContainer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accountName"] = state.AccountName
		inputs["databaseName"] = state.DatabaseName
		inputs["name"] = state.Name
		inputs["partitionKeyPath"] = state.PartitionKeyPath
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["uniqueKeys"] = state.UniqueKeys
	}
	s, err := ctx.ReadResource("azure:cosmosdb/sqlContainer:SqlContainer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SqlContainer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SqlContainer) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SqlContainer) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
func (r *SqlContainer) AccountName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accountName"])
}

// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
func (r *SqlContainer) DatabaseName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["databaseName"])
}

// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
func (r *SqlContainer) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Define a partition key. Changing this forces a new resource to be created.
func (r *SqlContainer) PartitionKeyPath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["partitionKeyPath"])
}

// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
func (r *SqlContainer) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
func (r *SqlContainer) UniqueKeys() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["uniqueKeys"])
}

// Input properties used for looking up and filtering SqlContainer resources.
type SqlContainerState struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName interface{}
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName interface{}
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name interface{}
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath interface{}
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys interface{}
}

// The set of arguments for constructing a SqlContainer resource.
type SqlContainerArgs struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName interface{}
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName interface{}
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name interface{}
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath interface{}
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys interface{}
}