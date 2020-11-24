// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Create a failover group of databases on a collection of Azure SQL servers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/sql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("uksouth"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primary, err := sql.NewSqlServer(ctx, "primary", &sql.SqlServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   exampleResourceGroup.Location,
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("sqladmin"),
// 			AdministratorLoginPassword: pulumi.String(fmt.Sprintf("%v%v%v%v", "pa", "$", "$", "w0rd")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondary, err := sql.NewSqlServer(ctx, "secondary", &sql.SqlServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   pulumi.String("northeurope"),
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("sqladmin"),
// 			AdministratorLoginPassword: pulumi.String(fmt.Sprintf("%v%v%v%v", "pa", "$", "$", "w0rd")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		db1, err := sql.NewDatabase(ctx, "db1", &sql.DatabaseArgs{
// 			ResourceGroupName: primary.ResourceGroupName,
// 			Location:          primary.Location,
// 			ServerName:        primary.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewFailoverGroup(ctx, "exampleFailoverGroup", &sql.FailoverGroupArgs{
// 			ResourceGroupName: primary.ResourceGroupName,
// 			ServerName:        primary.Name,
// 			Databases: pulumi.StringArray{
// 				db1.ID(),
// 			},
// 			PartnerServers: sql.FailoverGroupPartnerServerArray{
// 				&sql.FailoverGroupPartnerServerArgs{
// 					Id: secondary.ID(),
// 				},
// 			},
// 			ReadWriteEndpointFailoverPolicy: &sql.FailoverGroupReadWriteEndpointFailoverPolicyArgs{
// 				Mode:         pulumi.String("Automatic"),
// 				GraceMinutes: pulumi.Int(60),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// SQL Failover Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sql/failoverGroup:FailoverGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/failovergroups/group1
// ```
type FailoverGroup struct {
	pulumi.CustomResourceState

	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayOutput `pulumi:"databases"`
	// the location of the failover group.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServerArrayOutput `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyOutput `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyOutput `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// local replication role of the failover group instance.
	Role pulumi.StringOutput `pulumi:"role"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewFailoverGroup registers a new resource with the given unique name, arguments, and options.
func NewFailoverGroup(ctx *pulumi.Context,
	name string, args *FailoverGroupArgs, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PartnerServers == nil {
		return nil, errors.New("invalid value for required argument 'PartnerServers'")
	}
	if args.ReadWriteEndpointFailoverPolicy == nil {
		return nil, errors.New("invalid value for required argument 'ReadWriteEndpointFailoverPolicy'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	var resource FailoverGroup
	err := ctx.RegisterResource("azure:sql/failoverGroup:FailoverGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFailoverGroup gets an existing FailoverGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFailoverGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FailoverGroupState, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	var resource FailoverGroup
	err := ctx.ReadResource("azure:sql/failoverGroup:FailoverGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FailoverGroup resources.
type failoverGroupState struct {
	// A list of database ids to add to the failover group
	Databases []string `pulumi:"databases"`
	// the location of the failover group.
	Location *string `pulumi:"location"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers []FailoverGroupPartnerServer `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy *FailoverGroupReadWriteEndpointFailoverPolicy `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy *FailoverGroupReadonlyEndpointFailoverPolicy `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// local replication role of the failover group instance.
	Role *string `pulumi:"role"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type FailoverGroupState struct {
	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayInput
	// the location of the failover group.
	Location pulumi.StringPtrInput
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServerArrayInput
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyPtrInput
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyPtrInput
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringPtrInput
	// local replication role of the failover group instance.
	Role pulumi.StringPtrInput
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (FailoverGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupState)(nil)).Elem()
}

type failoverGroupArgs struct {
	// A list of database ids to add to the failover group
	Databases []string `pulumi:"databases"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers []FailoverGroupPartnerServer `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicy `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy *FailoverGroupReadonlyEndpointFailoverPolicy `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FailoverGroup resource.
type FailoverGroupArgs struct {
	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayInput
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServerArrayInput
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyInput
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyPtrInput
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringInput
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (FailoverGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupArgs)(nil)).Elem()
}

type FailoverGroupInput interface {
	pulumi.Input

	ToFailoverGroupOutput() FailoverGroupOutput
	ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput
}

func (FailoverGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroup)(nil)).Elem()
}

func (i FailoverGroup) ToFailoverGroupOutput() FailoverGroupOutput {
	return i.ToFailoverGroupOutputWithContext(context.Background())
}

func (i FailoverGroup) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupOutput)
}

type FailoverGroupOutput struct {
	*pulumi.OutputState
}

func (FailoverGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupOutput)(nil)).Elem()
}

func (o FailoverGroupOutput) ToFailoverGroupOutput() FailoverGroupOutput {
	return o
}

func (o FailoverGroupOutput) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FailoverGroupOutput{})
}
