// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Security Alert Policy for a MSSQL Server.
//
// > **NOTE** Security Alert Policy is currently only available for MS SQL databases.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mssql"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/sql"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		exampleSqlServer, err := sql.NewSqlServer(ctx, "exampleSqlServer", &sql.SqlServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   exampleResourceGroup.Location,
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("4dm1n157r470r"),
// 			AdministratorLoginPassword: pulumi.String("4-v3ry-53cr37-p455w0rd"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("GRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mssql.NewServerSecurityAlertPolicy(ctx, "exampleServerSecurityAlertPolicy", &mssql.ServerSecurityAlertPolicyArgs{
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			ServerName:              exampleSqlServer.Name,
// 			State:                   pulumi.String("Enabled"),
// 			StorageEndpoint:         exampleAccount.PrimaryBlobEndpoint,
// 			StorageAccountAccessKey: exampleAccount.PrimaryAccessKey,
// 			DisabledAlerts: pulumi.StringArray{
// 				pulumi.String("Sql_Injection"),
// 				pulumi.String("Data_Exfiltration"),
// 			},
// 			RetentionDays: pulumi.Int(20),
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
// MS SQL Server Security Alert Policy can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mssql/serverSecurityAlertPolicy:ServerSecurityAlertPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/securityAlertPolicies/Default
// ```
type ServerSecurityAlertPolicy struct {
	pulumi.CustomResourceState

	// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
	DisabledAlerts pulumi.StringArrayOutput `pulumi:"disabledAlerts"`
	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
	EmailAccountAdmins pulumi.BoolPtrOutput `pulumi:"emailAccountAdmins"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses pulumi.StringArrayOutput `pulumi:"emailAddresses"`
	// The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
	// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`.
	State pulumi.StringOutput `pulumi:"state"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey pulumi.StringPtrOutput `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
}

// NewServerSecurityAlertPolicy registers a new resource with the given unique name, arguments, and options.
func NewServerSecurityAlertPolicy(ctx *pulumi.Context,
	name string, args *ServerSecurityAlertPolicyArgs, opts ...pulumi.ResourceOption) (*ServerSecurityAlertPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	var resource ServerSecurityAlertPolicy
	err := ctx.RegisterResource("azure:mssql/serverSecurityAlertPolicy:ServerSecurityAlertPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerSecurityAlertPolicy gets an existing ServerSecurityAlertPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerSecurityAlertPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerSecurityAlertPolicyState, opts ...pulumi.ResourceOption) (*ServerSecurityAlertPolicy, error) {
	var resource ServerSecurityAlertPolicy
	err := ctx.ReadResource("azure:mssql/serverSecurityAlertPolicy:ServerSecurityAlertPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerSecurityAlertPolicy resources.
type serverSecurityAlertPolicyState struct {
	// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
	DisabledAlerts []string `pulumi:"disabledAlerts"`
	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
	EmailAccountAdmins *bool `pulumi:"emailAccountAdmins"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `pulumi:"emailAddresses"`
	// The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
	RetentionDays *int `pulumi:"retentionDays"`
	// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`.
	State *string `pulumi:"state"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
}

type ServerSecurityAlertPolicyState struct {
	// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
	DisabledAlerts pulumi.StringArrayInput
	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
	EmailAccountAdmins pulumi.BoolPtrInput
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses pulumi.StringArrayInput
	// The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
	RetentionDays pulumi.IntPtrInput
	// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`.
	State pulumi.StringPtrInput
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint pulumi.StringPtrInput
}

func (ServerSecurityAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverSecurityAlertPolicyState)(nil)).Elem()
}

type serverSecurityAlertPolicyArgs struct {
	// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
	DisabledAlerts []string `pulumi:"disabledAlerts"`
	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
	EmailAccountAdmins *bool `pulumi:"emailAccountAdmins"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `pulumi:"emailAddresses"`
	// The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
	RetentionDays *int `pulumi:"retentionDays"`
	// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`.
	State string `pulumi:"state"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
}

// The set of arguments for constructing a ServerSecurityAlertPolicy resource.
type ServerSecurityAlertPolicyArgs struct {
	// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
	DisabledAlerts pulumi.StringArrayInput
	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
	EmailAccountAdmins pulumi.BoolPtrInput
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses pulumi.StringArrayInput
	// The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
	RetentionDays pulumi.IntPtrInput
	// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Allowed values are: `Disabled`, `Enabled`.
	State pulumi.StringInput
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint pulumi.StringPtrInput
}

func (ServerSecurityAlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverSecurityAlertPolicyArgs)(nil)).Elem()
}

type ServerSecurityAlertPolicyInput interface {
	pulumi.Input

	ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput
	ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput
}

func (*ServerSecurityAlertPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerSecurityAlertPolicy)(nil))
}

func (i *ServerSecurityAlertPolicy) ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput {
	return i.ToServerSecurityAlertPolicyOutputWithContext(context.Background())
}

func (i *ServerSecurityAlertPolicy) ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSecurityAlertPolicyOutput)
}

func (i *ServerSecurityAlertPolicy) ToServerSecurityAlertPolicyPtrOutput() ServerSecurityAlertPolicyPtrOutput {
	return i.ToServerSecurityAlertPolicyPtrOutputWithContext(context.Background())
}

func (i *ServerSecurityAlertPolicy) ToServerSecurityAlertPolicyPtrOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSecurityAlertPolicyPtrOutput)
}

type ServerSecurityAlertPolicyPtrInput interface {
	pulumi.Input

	ToServerSecurityAlertPolicyPtrOutput() ServerSecurityAlertPolicyPtrOutput
	ToServerSecurityAlertPolicyPtrOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyPtrOutput
}

type serverSecurityAlertPolicyPtrType ServerSecurityAlertPolicyArgs

func (*serverSecurityAlertPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSecurityAlertPolicy)(nil))
}

func (i *serverSecurityAlertPolicyPtrType) ToServerSecurityAlertPolicyPtrOutput() ServerSecurityAlertPolicyPtrOutput {
	return i.ToServerSecurityAlertPolicyPtrOutputWithContext(context.Background())
}

func (i *serverSecurityAlertPolicyPtrType) ToServerSecurityAlertPolicyPtrOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSecurityAlertPolicyPtrOutput)
}

// ServerSecurityAlertPolicyArrayInput is an input type that accepts ServerSecurityAlertPolicyArray and ServerSecurityAlertPolicyArrayOutput values.
// You can construct a concrete instance of `ServerSecurityAlertPolicyArrayInput` via:
//
//          ServerSecurityAlertPolicyArray{ ServerSecurityAlertPolicyArgs{...} }
type ServerSecurityAlertPolicyArrayInput interface {
	pulumi.Input

	ToServerSecurityAlertPolicyArrayOutput() ServerSecurityAlertPolicyArrayOutput
	ToServerSecurityAlertPolicyArrayOutputWithContext(context.Context) ServerSecurityAlertPolicyArrayOutput
}

type ServerSecurityAlertPolicyArray []ServerSecurityAlertPolicyInput

func (ServerSecurityAlertPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServerSecurityAlertPolicy)(nil))
}

func (i ServerSecurityAlertPolicyArray) ToServerSecurityAlertPolicyArrayOutput() ServerSecurityAlertPolicyArrayOutput {
	return i.ToServerSecurityAlertPolicyArrayOutputWithContext(context.Background())
}

func (i ServerSecurityAlertPolicyArray) ToServerSecurityAlertPolicyArrayOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSecurityAlertPolicyArrayOutput)
}

// ServerSecurityAlertPolicyMapInput is an input type that accepts ServerSecurityAlertPolicyMap and ServerSecurityAlertPolicyMapOutput values.
// You can construct a concrete instance of `ServerSecurityAlertPolicyMapInput` via:
//
//          ServerSecurityAlertPolicyMap{ "key": ServerSecurityAlertPolicyArgs{...} }
type ServerSecurityAlertPolicyMapInput interface {
	pulumi.Input

	ToServerSecurityAlertPolicyMapOutput() ServerSecurityAlertPolicyMapOutput
	ToServerSecurityAlertPolicyMapOutputWithContext(context.Context) ServerSecurityAlertPolicyMapOutput
}

type ServerSecurityAlertPolicyMap map[string]ServerSecurityAlertPolicyInput

func (ServerSecurityAlertPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServerSecurityAlertPolicy)(nil))
}

func (i ServerSecurityAlertPolicyMap) ToServerSecurityAlertPolicyMapOutput() ServerSecurityAlertPolicyMapOutput {
	return i.ToServerSecurityAlertPolicyMapOutputWithContext(context.Background())
}

func (i ServerSecurityAlertPolicyMap) ToServerSecurityAlertPolicyMapOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSecurityAlertPolicyMapOutput)
}

type ServerSecurityAlertPolicyOutput struct {
	*pulumi.OutputState
}

func (ServerSecurityAlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerSecurityAlertPolicy)(nil))
}

func (o ServerSecurityAlertPolicyOutput) ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput {
	return o
}

func (o ServerSecurityAlertPolicyOutput) ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput {
	return o
}

func (o ServerSecurityAlertPolicyOutput) ToServerSecurityAlertPolicyPtrOutput() ServerSecurityAlertPolicyPtrOutput {
	return o.ToServerSecurityAlertPolicyPtrOutputWithContext(context.Background())
}

func (o ServerSecurityAlertPolicyOutput) ToServerSecurityAlertPolicyPtrOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyPtrOutput {
	return o.ApplyT(func(v ServerSecurityAlertPolicy) *ServerSecurityAlertPolicy {
		return &v
	}).(ServerSecurityAlertPolicyPtrOutput)
}

type ServerSecurityAlertPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (ServerSecurityAlertPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSecurityAlertPolicy)(nil))
}

func (o ServerSecurityAlertPolicyPtrOutput) ToServerSecurityAlertPolicyPtrOutput() ServerSecurityAlertPolicyPtrOutput {
	return o
}

func (o ServerSecurityAlertPolicyPtrOutput) ToServerSecurityAlertPolicyPtrOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyPtrOutput {
	return o
}

type ServerSecurityAlertPolicyArrayOutput struct{ *pulumi.OutputState }

func (ServerSecurityAlertPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerSecurityAlertPolicy)(nil))
}

func (o ServerSecurityAlertPolicyArrayOutput) ToServerSecurityAlertPolicyArrayOutput() ServerSecurityAlertPolicyArrayOutput {
	return o
}

func (o ServerSecurityAlertPolicyArrayOutput) ToServerSecurityAlertPolicyArrayOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyArrayOutput {
	return o
}

func (o ServerSecurityAlertPolicyArrayOutput) Index(i pulumi.IntInput) ServerSecurityAlertPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerSecurityAlertPolicy {
		return vs[0].([]ServerSecurityAlertPolicy)[vs[1].(int)]
	}).(ServerSecurityAlertPolicyOutput)
}

type ServerSecurityAlertPolicyMapOutput struct{ *pulumi.OutputState }

func (ServerSecurityAlertPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServerSecurityAlertPolicy)(nil))
}

func (o ServerSecurityAlertPolicyMapOutput) ToServerSecurityAlertPolicyMapOutput() ServerSecurityAlertPolicyMapOutput {
	return o
}

func (o ServerSecurityAlertPolicyMapOutput) ToServerSecurityAlertPolicyMapOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyMapOutput {
	return o
}

func (o ServerSecurityAlertPolicyMapOutput) MapIndex(k pulumi.StringInput) ServerSecurityAlertPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServerSecurityAlertPolicy {
		return vs[0].(map[string]ServerSecurityAlertPolicy)[vs[1].(string)]
	}).(ServerSecurityAlertPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerSecurityAlertPolicyOutput{})
	pulumi.RegisterOutputType(ServerSecurityAlertPolicyPtrOutput{})
	pulumi.RegisterOutputType(ServerSecurityAlertPolicyArrayOutput{})
	pulumi.RegisterOutputType(ServerSecurityAlertPolicyMapOutput{})
}
