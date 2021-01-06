// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sets a MySQL Configuration value on a MySQL Server.
//
// ## Disclaimers
//
// > **Note:** Since this resource is provisioned by default, the Azure Provider will not check for the presence of an existing resource prior to attempting to create it.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mysql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		exampleServer, err := mysql.NewServer(ctx, "exampleServer", &mysql.ServerArgs{
// 			Location:                        exampleResourceGroup.Location,
// 			ResourceGroupName:               exampleResourceGroup.Name,
// 			AdministratorLogin:              pulumi.String("mysqladminun"),
// 			AdministratorLoginPassword:      pulumi.String("H@Sh1CoR3!"),
// 			SkuName:                         pulumi.String("B_Gen5_2"),
// 			StorageMb:                       pulumi.Int(5120),
// 			Version:                         pulumi.String("5.7"),
// 			AutoGrowEnabled:                 pulumi.Bool(true),
// 			BackupRetentionDays:             pulumi.Int(7),
// 			GeoRedundantBackupEnabled:       pulumi.Bool(true),
// 			InfrastructureEncryptionEnabled: pulumi.Bool(true),
// 			PublicNetworkAccessEnabled:      pulumi.Bool(false),
// 			SslEnforcementEnabled:           pulumi.Bool(true),
// 			SslMinimalTlsVersionEnforced:    pulumi.String("TLS1_2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mysql.NewConfiguration(ctx, "exampleConfiguration", &mysql.ConfigurationArgs{
// 			Name:              pulumi.String("interactive_timeout"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        exampleServer.Name,
// 			Value:             pulumi.String("600"),
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
// MySQL Configurations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mysql/configuration:Configuration interactive_timeout /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/servers/server1/configurations/interactive_timeout
// ```
type Configuration struct {
	pulumi.CustomResourceState

	// Specifies the name of the MySQL Configuration, which needs [to be a valid MySQL configuration name](https://dev.mysql.com/doc/refman/5.7/en/server-configuration.html). Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Specifies the value of the MySQL Configuration. See the MySQL documentation for valid values.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource Configuration
	err := ctx.RegisterResource("azure:mysql/configuration:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("azure:mysql/configuration:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
	// Specifies the name of the MySQL Configuration, which needs [to be a valid MySQL configuration name](https://dev.mysql.com/doc/refman/5.7/en/server-configuration.html). Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// Specifies the value of the MySQL Configuration. See the MySQL documentation for valid values.
	Value *string `pulumi:"value"`
}

type ConfigurationState struct {
	// Specifies the name of the MySQL Configuration, which needs [to be a valid MySQL configuration name](https://dev.mysql.com/doc/refman/5.7/en/server-configuration.html). Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// Specifies the value of the MySQL Configuration. See the MySQL documentation for valid values.
	Value pulumi.StringPtrInput
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	// Specifies the name of the MySQL Configuration, which needs [to be a valid MySQL configuration name](https://dev.mysql.com/doc/refman/5.7/en/server-configuration.html). Changing this forces a new resource to be created.
	Name string `pulumi:"name"`
	// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// Specifies the value of the MySQL Configuration. See the MySQL documentation for valid values.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// Specifies the name of the MySQL Configuration, which needs [to be a valid MySQL configuration name](https://dev.mysql.com/doc/refman/5.7/en/server-configuration.html). Changing this forces a new resource to be created.
	Name pulumi.StringInput
	// The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// Specifies the value of the MySQL Configuration. See the MySQL documentation for valid values.
	Value pulumi.StringInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}

type ConfigurationInput interface {
	pulumi.Input

	ToConfigurationOutput() ConfigurationOutput
	ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput
}

func (*Configuration) ElementType() reflect.Type {
	return reflect.TypeOf((*Configuration)(nil))
}

func (i *Configuration) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i *Configuration) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

func (i *Configuration) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return i.ToConfigurationPtrOutputWithContext(context.Background())
}

func (i *Configuration) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationPtrOutput)
}

type ConfigurationPtrInput interface {
	pulumi.Input

	ToConfigurationPtrOutput() ConfigurationPtrOutput
	ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput
}

type configurationPtrType ConfigurationArgs

func (*configurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil))
}

func (i *configurationPtrType) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return i.ToConfigurationPtrOutputWithContext(context.Background())
}

func (i *configurationPtrType) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput).ToConfigurationPtrOutput()
}

// ConfigurationArrayInput is an input type that accepts ConfigurationArray and ConfigurationArrayOutput values.
// You can construct a concrete instance of `ConfigurationArrayInput` via:
//
//          ConfigurationArray{ ConfigurationArgs{...} }
type ConfigurationArrayInput interface {
	pulumi.Input

	ToConfigurationArrayOutput() ConfigurationArrayOutput
	ToConfigurationArrayOutputWithContext(context.Context) ConfigurationArrayOutput
}

type ConfigurationArray []ConfigurationInput

func (ConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Configuration)(nil))
}

func (i ConfigurationArray) ToConfigurationArrayOutput() ConfigurationArrayOutput {
	return i.ToConfigurationArrayOutputWithContext(context.Background())
}

func (i ConfigurationArray) ToConfigurationArrayOutputWithContext(ctx context.Context) ConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationArrayOutput)
}

// ConfigurationMapInput is an input type that accepts ConfigurationMap and ConfigurationMapOutput values.
// You can construct a concrete instance of `ConfigurationMapInput` via:
//
//          ConfigurationMap{ "key": ConfigurationArgs{...} }
type ConfigurationMapInput interface {
	pulumi.Input

	ToConfigurationMapOutput() ConfigurationMapOutput
	ToConfigurationMapOutputWithContext(context.Context) ConfigurationMapOutput
}

type ConfigurationMap map[string]ConfigurationInput

func (ConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Configuration)(nil))
}

func (i ConfigurationMap) ToConfigurationMapOutput() ConfigurationMapOutput {
	return i.ToConfigurationMapOutputWithContext(context.Background())
}

func (i ConfigurationMap) ToConfigurationMapOutputWithContext(ctx context.Context) ConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationMapOutput)
}

type ConfigurationOutput struct {
	*pulumi.OutputState
}

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Configuration)(nil))
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return o.ToConfigurationPtrOutputWithContext(context.Background())
}

func (o ConfigurationOutput) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return o.ApplyT(func(v Configuration) *Configuration {
		return &v
	}).(ConfigurationPtrOutput)
}

type ConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (ConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil))
}

func (o ConfigurationPtrOutput) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return o
}

func (o ConfigurationPtrOutput) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return o
}

type ConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Configuration)(nil))
}

func (o ConfigurationArrayOutput) ToConfigurationArrayOutput() ConfigurationArrayOutput {
	return o
}

func (o ConfigurationArrayOutput) ToConfigurationArrayOutputWithContext(ctx context.Context) ConfigurationArrayOutput {
	return o
}

func (o ConfigurationArrayOutput) Index(i pulumi.IntInput) ConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Configuration {
		return vs[0].([]Configuration)[vs[1].(int)]
	}).(ConfigurationOutput)
}

type ConfigurationMapOutput struct{ *pulumi.OutputState }

func (ConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Configuration)(nil))
}

func (o ConfigurationMapOutput) ToConfigurationMapOutput() ConfigurationMapOutput {
	return o
}

func (o ConfigurationMapOutput) ToConfigurationMapOutputWithContext(ctx context.Context) ConfigurationMapOutput {
	return o
}

func (o ConfigurationMapOutput) MapIndex(k pulumi.StringInput) ConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Configuration {
		return vs[0].(map[string]Configuration)[vs[1].(string)]
	}).(ConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationOutput{})
	pulumi.RegisterOutputType(ConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationMapOutput{})
}
