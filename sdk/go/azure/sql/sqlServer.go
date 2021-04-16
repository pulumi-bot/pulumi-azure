// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Microsoft SQL Azure Database Server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/sql"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewSqlServer(ctx, "exampleSqlServer", &sql.SqlServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   exampleResourceGroup.Location,
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("mradministrator"),
// 			AdministratorLoginPassword: pulumi.String("thisIsDog11"),
// 			ExtendedAuditingPolicy: &sql.SqlServerExtendedAuditingPolicyArgs{
// 				StorageEndpoint:                    exampleAccount.PrimaryBlobEndpoint,
// 				StorageAccountAccessKey:            exampleAccount.PrimaryAccessKey,
// 				StorageAccountAccessKeyIsSecondary: pulumi.Bool(true),
// 				RetentionInDays:                    pulumi.Int(6),
// 			},
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("production"),
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
// SQL Servers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sql/sqlServer:SqlServer example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver
// ```
type SqlServer struct {
	pulumi.CustomResourceState

	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringOutput `pulumi:"administratorLogin"`
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword pulumi.StringOutput `pulumi:"administratorLoginPassword"`
	// The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
	ConnectionPolicy pulumi.StringPtrOutput `pulumi:"connectionPolicy"`
	// A `extendedAuditingPolicy` block as defined below.
	//
	// Deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
	ExtendedAuditingPolicy SqlServerExtendedAuditingPolicyOutput `pulumi:"extendedAuditingPolicy"`
	// The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
	FullyQualifiedDomainName pulumi.StringOutput `pulumi:"fullyQualifiedDomainName"`
	// An `identity` block as defined below.
	Identity SqlServerIdentityPtrOutput `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Microsoft SQL Server.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewSqlServer registers a new resource with the given unique name, arguments, and options.
func NewSqlServer(ctx *pulumi.Context,
	name string, args *SqlServerArgs, opts ...pulumi.ResourceOption) (*SqlServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdministratorLogin == nil {
		return nil, errors.New("invalid value for required argument 'AdministratorLogin'")
	}
	if args.AdministratorLoginPassword == nil {
		return nil, errors.New("invalid value for required argument 'AdministratorLoginPassword'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource SqlServer
	err := ctx.RegisterResource("azure:sql/sqlServer:SqlServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlServer gets an existing SqlServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlServerState, opts ...pulumi.ResourceOption) (*SqlServer, error) {
	var resource SqlServer
	err := ctx.ReadResource("azure:sql/sqlServer:SqlServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlServer resources.
type sqlServerState struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
	ConnectionPolicy *string `pulumi:"connectionPolicy"`
	// A `extendedAuditingPolicy` block as defined below.
	//
	// Deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
	ExtendedAuditingPolicy *SqlServerExtendedAuditingPolicy `pulumi:"extendedAuditingPolicy"`
	// The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
	FullyQualifiedDomainName *string `pulumi:"fullyQualifiedDomainName"`
	// An `identity` block as defined below.
	Identity *SqlServerIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Microsoft SQL Server.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
	Version *string `pulumi:"version"`
}

type SqlServerState struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringPtrInput
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword pulumi.StringPtrInput
	// The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
	ConnectionPolicy pulumi.StringPtrInput
	// A `extendedAuditingPolicy` block as defined below.
	//
	// Deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
	ExtendedAuditingPolicy SqlServerExtendedAuditingPolicyPtrInput
	// The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
	FullyQualifiedDomainName pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity SqlServerIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Microsoft SQL Server.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
	Version pulumi.StringPtrInput
}

func (SqlServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerState)(nil)).Elem()
}

type sqlServerArgs struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin string `pulumi:"administratorLogin"`
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword string `pulumi:"administratorLoginPassword"`
	// The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
	ConnectionPolicy *string `pulumi:"connectionPolicy"`
	// A `extendedAuditingPolicy` block as defined below.
	//
	// Deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
	ExtendedAuditingPolicy *SqlServerExtendedAuditingPolicy `pulumi:"extendedAuditingPolicy"`
	// An `identity` block as defined below.
	Identity *SqlServerIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Microsoft SQL Server.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a SqlServer resource.
type SqlServerArgs struct {
	// The administrator login name for the new server. Changing this forces a new resource to be created.
	AdministratorLogin pulumi.StringInput
	// The password associated with the `administratorLogin` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
	AdministratorLoginPassword pulumi.StringInput
	// The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
	ConnectionPolicy pulumi.StringPtrInput
	// A `extendedAuditingPolicy` block as defined below.
	//
	// Deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
	ExtendedAuditingPolicy SqlServerExtendedAuditingPolicyPtrInput
	// An `identity` block as defined below.
	Identity SqlServerIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Microsoft SQL Server.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
	Version pulumi.StringInput
}

func (SqlServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerArgs)(nil)).Elem()
}

type SqlServerInput interface {
	pulumi.Input

	ToSqlServerOutput() SqlServerOutput
	ToSqlServerOutputWithContext(ctx context.Context) SqlServerOutput
}

func (*SqlServer) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServer)(nil))
}

func (i *SqlServer) ToSqlServerOutput() SqlServerOutput {
	return i.ToSqlServerOutputWithContext(context.Background())
}

func (i *SqlServer) ToSqlServerOutputWithContext(ctx context.Context) SqlServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerOutput)
}

func (i *SqlServer) ToSqlServerPtrOutput() SqlServerPtrOutput {
	return i.ToSqlServerPtrOutputWithContext(context.Background())
}

func (i *SqlServer) ToSqlServerPtrOutputWithContext(ctx context.Context) SqlServerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerPtrOutput)
}

type SqlServerPtrInput interface {
	pulumi.Input

	ToSqlServerPtrOutput() SqlServerPtrOutput
	ToSqlServerPtrOutputWithContext(ctx context.Context) SqlServerPtrOutput
}

type sqlServerPtrType SqlServerArgs

func (*sqlServerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServer)(nil))
}

func (i *sqlServerPtrType) ToSqlServerPtrOutput() SqlServerPtrOutput {
	return i.ToSqlServerPtrOutputWithContext(context.Background())
}

func (i *sqlServerPtrType) ToSqlServerPtrOutputWithContext(ctx context.Context) SqlServerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerPtrOutput)
}

// SqlServerArrayInput is an input type that accepts SqlServerArray and SqlServerArrayOutput values.
// You can construct a concrete instance of `SqlServerArrayInput` via:
//
//          SqlServerArray{ SqlServerArgs{...} }
type SqlServerArrayInput interface {
	pulumi.Input

	ToSqlServerArrayOutput() SqlServerArrayOutput
	ToSqlServerArrayOutputWithContext(context.Context) SqlServerArrayOutput
}

type SqlServerArray []SqlServerInput

func (SqlServerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SqlServer)(nil))
}

func (i SqlServerArray) ToSqlServerArrayOutput() SqlServerArrayOutput {
	return i.ToSqlServerArrayOutputWithContext(context.Background())
}

func (i SqlServerArray) ToSqlServerArrayOutputWithContext(ctx context.Context) SqlServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerArrayOutput)
}

// SqlServerMapInput is an input type that accepts SqlServerMap and SqlServerMapOutput values.
// You can construct a concrete instance of `SqlServerMapInput` via:
//
//          SqlServerMap{ "key": SqlServerArgs{...} }
type SqlServerMapInput interface {
	pulumi.Input

	ToSqlServerMapOutput() SqlServerMapOutput
	ToSqlServerMapOutputWithContext(context.Context) SqlServerMapOutput
}

type SqlServerMap map[string]SqlServerInput

func (SqlServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SqlServer)(nil))
}

func (i SqlServerMap) ToSqlServerMapOutput() SqlServerMapOutput {
	return i.ToSqlServerMapOutputWithContext(context.Background())
}

func (i SqlServerMap) ToSqlServerMapOutputWithContext(ctx context.Context) SqlServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerMapOutput)
}

type SqlServerOutput struct {
	*pulumi.OutputState
}

func (SqlServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServer)(nil))
}

func (o SqlServerOutput) ToSqlServerOutput() SqlServerOutput {
	return o
}

func (o SqlServerOutput) ToSqlServerOutputWithContext(ctx context.Context) SqlServerOutput {
	return o
}

func (o SqlServerOutput) ToSqlServerPtrOutput() SqlServerPtrOutput {
	return o.ToSqlServerPtrOutputWithContext(context.Background())
}

func (o SqlServerOutput) ToSqlServerPtrOutputWithContext(ctx context.Context) SqlServerPtrOutput {
	return o.ApplyT(func(v SqlServer) *SqlServer {
		return &v
	}).(SqlServerPtrOutput)
}

type SqlServerPtrOutput struct {
	*pulumi.OutputState
}

func (SqlServerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServer)(nil))
}

func (o SqlServerPtrOutput) ToSqlServerPtrOutput() SqlServerPtrOutput {
	return o
}

func (o SqlServerPtrOutput) ToSqlServerPtrOutputWithContext(ctx context.Context) SqlServerPtrOutput {
	return o
}

type SqlServerArrayOutput struct{ *pulumi.OutputState }

func (SqlServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SqlServer)(nil))
}

func (o SqlServerArrayOutput) ToSqlServerArrayOutput() SqlServerArrayOutput {
	return o
}

func (o SqlServerArrayOutput) ToSqlServerArrayOutputWithContext(ctx context.Context) SqlServerArrayOutput {
	return o
}

func (o SqlServerArrayOutput) Index(i pulumi.IntInput) SqlServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SqlServer {
		return vs[0].([]SqlServer)[vs[1].(int)]
	}).(SqlServerOutput)
}

type SqlServerMapOutput struct{ *pulumi.OutputState }

func (SqlServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SqlServer)(nil))
}

func (o SqlServerMapOutput) ToSqlServerMapOutput() SqlServerMapOutput {
	return o
}

func (o SqlServerMapOutput) ToSqlServerMapOutputWithContext(ctx context.Context) SqlServerMapOutput {
	return o
}

func (o SqlServerMapOutput) MapIndex(k pulumi.StringInput) SqlServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SqlServer {
		return vs[0].(map[string]SqlServer)[vs[1].(string)]
	}).(SqlServerOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlServerOutput{})
	pulumi.RegisterOutputType(SqlServerPtrOutput{})
	pulumi.RegisterOutputType(SqlServerArrayOutput{})
	pulumi.RegisterOutputType(SqlServerMapOutput{})
}
