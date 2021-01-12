// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package maintenance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a maintenance configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/maintenance"
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
// 		_, err = maintenance.NewConfiguration(ctx, "exampleConfiguration", &maintenance.ConfigurationArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Scope:             pulumi.String("All"),
// 			Tags: pulumi.StringMap{
// 				"Env": pulumi.String("prod"),
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
// Maintenance Configuration can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:maintenance/configuration:Configuration example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.maintenance/maintenanceconfigurations/example-mc
// ```
type Configuration struct {
	pulumi.CustomResourceState

	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Maintenance Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Maintenance Configuration should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The scope of the Maintenance Configuration. Possible values are `All`, `Host`, `Resource` or `InResource`. Default to `All`.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// A mapping of tags to assign to the resource. The key could not contain upper case letter.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Configuration
	err := ctx.RegisterResource("azure:maintenance/configuration:Configuration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:maintenance/configuration:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Maintenance Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Maintenance Configuration should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The scope of the Maintenance Configuration. Possible values are `All`, `Host`, `Resource` or `InResource`. Default to `All`.
	Scope *string `pulumi:"scope"`
	// A mapping of tags to assign to the resource. The key could not contain upper case letter.
	Tags map[string]string `pulumi:"tags"`
}

type ConfigurationState struct {
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Maintenance Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Maintenance Configuration should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The scope of the Maintenance Configuration. Possible values are `All`, `Host`, `Resource` or `InResource`. Default to `All`.
	Scope pulumi.StringPtrInput
	// A mapping of tags to assign to the resource. The key could not contain upper case letter.
	Tags pulumi.StringMapInput
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Maintenance Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Maintenance Configuration should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The scope of the Maintenance Configuration. Possible values are `All`, `Host`, `Resource` or `InResource`. Default to `All`.
	Scope *string `pulumi:"scope"`
	// A mapping of tags to assign to the resource. The key could not contain upper case letter.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Maintenance Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Maintenance Configuration should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The scope of the Maintenance Configuration. Possible values are `All`, `Host`, `Resource` or `InResource`. Default to `All`.
	Scope pulumi.StringPtrInput
	// A mapping of tags to assign to the resource. The key could not contain upper case letter.
	Tags pulumi.StringMapInput
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
