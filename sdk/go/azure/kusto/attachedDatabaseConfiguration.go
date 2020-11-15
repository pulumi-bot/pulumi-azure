// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) Attached Database Configuration
//
// ## Import
//
// Kusto Attached Database Configurations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/AttachedDatabaseConfigurations/configuration1
// ```
type AttachedDatabaseConfiguration struct {
	pulumi.CustomResourceState

	// The list of databases from the `clusterResourceId` which are currently attached to the cluster.
	AttachedDatabaseNames pulumi.StringArrayOutput `pulumi:"attachedDatabaseNames"`
	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId pulumi.StringOutput `pulumi:"clusterResourceId"`
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
	DefaultPrincipalModificationKind pulumi.StringPtrOutput `pulumi:"defaultPrincipalModificationKind"`
	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewAttachedDatabaseConfiguration registers a new resource with the given unique name, arguments, and options.
func NewAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, args *AttachedDatabaseConfigurationArgs, opts ...pulumi.ResourceOption) (*AttachedDatabaseConfiguration, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.ClusterResourceId == nil {
		return nil, errors.New("missing required argument 'ClusterResourceId'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AttachedDatabaseConfigurationArgs{}
	}
	var resource AttachedDatabaseConfiguration
	err := ctx.RegisterResource("azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachedDatabaseConfiguration gets an existing AttachedDatabaseConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedDatabaseConfigurationState, opts ...pulumi.ResourceOption) (*AttachedDatabaseConfiguration, error) {
	var resource AttachedDatabaseConfiguration
	err := ctx.ReadResource("azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttachedDatabaseConfiguration resources.
type attachedDatabaseConfigurationState struct {
	// The list of databases from the `clusterResourceId` which are currently attached to the cluster.
	AttachedDatabaseNames []string `pulumi:"attachedDatabaseNames"`
	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName *string `pulumi:"clusterName"`
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId *string `pulumi:"clusterResourceId"`
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName *string `pulumi:"databaseName"`
	// The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
	DefaultPrincipalModificationKind *string `pulumi:"defaultPrincipalModificationKind"`
	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type AttachedDatabaseConfigurationState struct {
	// The list of databases from the `clusterResourceId` which are currently attached to the cluster.
	AttachedDatabaseNames pulumi.StringArrayInput
	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName pulumi.StringPtrInput
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId pulumi.StringPtrInput
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName pulumi.StringPtrInput
	// The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
	DefaultPrincipalModificationKind pulumi.StringPtrInput
	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (AttachedDatabaseConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDatabaseConfigurationState)(nil)).Elem()
}

type attachedDatabaseConfigurationArgs struct {
	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName string `pulumi:"clusterName"`
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId string `pulumi:"clusterResourceId"`
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName string `pulumi:"databaseName"`
	// The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
	DefaultPrincipalModificationKind *string `pulumi:"defaultPrincipalModificationKind"`
	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AttachedDatabaseConfiguration resource.
type AttachedDatabaseConfigurationArgs struct {
	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId pulumi.StringInput
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName pulumi.StringInput
	// The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
	DefaultPrincipalModificationKind pulumi.StringPtrInput
	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (AttachedDatabaseConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDatabaseConfigurationArgs)(nil)).Elem()
}

type AttachedDatabaseConfigurationInput interface {
	pulumi.Input

	ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput
	ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput
}

func (AttachedDatabaseConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*AttachedDatabaseConfiguration)(nil)).Elem()
}

func (i AttachedDatabaseConfiguration) ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput {
	return i.ToAttachedDatabaseConfigurationOutputWithContext(context.Background())
}

func (i AttachedDatabaseConfiguration) ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDatabaseConfigurationOutput)
}

type AttachedDatabaseConfigurationOutput struct {
	*pulumi.OutputState
}

func (AttachedDatabaseConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttachedDatabaseConfigurationOutput)(nil)).Elem()
}

func (o AttachedDatabaseConfigurationOutput) ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput {
	return o
}

func (o AttachedDatabaseConfigurationOutput) ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AttachedDatabaseConfigurationOutput{})
}
