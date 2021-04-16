// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Container Registry Webhook.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/containerservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		acr, err := containerservice.NewRegistry(ctx, "acr", &containerservice.RegistryArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 			Sku:               pulumi.String("Standard"),
// 			AdminEnabled:      pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = containerservice.NewRegistryWebhook(ctx, "webhook", &containerservice.RegistryWebhookArgs{
// 			ResourceGroupName: rg.Name,
// 			RegistryName:      acr.Name,
// 			Location:          rg.Location,
// 			ServiceUri:        pulumi.String("https://mywebhookreceiver.example/mytag"),
// 			Status:            pulumi.String("enabled"),
// 			Scope:             pulumi.String("mytag:*"),
// 			Actions: pulumi.StringArray{
// 				pulumi.String("push"),
// 			},
// 			CustomHeaders: pulumi.StringMap{
// 				"Content-Type": pulumi.String("application/json"),
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
// Container Registry Webhooks can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:containerservice/registryWebook:RegistryWebook example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/mygroup1/providers/Microsoft.ContainerRegistry/registries/myregistry1/webhooks/mywebhook1
// ```
//
// Deprecated: azure.containerservice.RegistryWebook has been deprecated in favor of azure.containerservice.RegistryWebhook
type RegistryWebook struct {
	pulumi.CustomResourceState

	// A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
	Actions pulumi.StringArrayOutput `pulumi:"actions"`
	// Custom headers that will be added to the webhook notifications request.
	CustomHeaders pulumi.StringMapOutput `pulumi:"customHeaders"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
	RegistryName pulumi.StringOutput `pulumi:"registryName"`
	// The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Specifies the service URI for the Webhook to post notifications.
	ServiceUri pulumi.StringOutput `pulumi:"serviceUri"`
	// Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	Tags   pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRegistryWebook registers a new resource with the given unique name, arguments, and options.
func NewRegistryWebook(ctx *pulumi.Context,
	name string, args *RegistryWebookArgs, opts ...pulumi.ResourceOption) (*RegistryWebook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceUri == nil {
		return nil, errors.New("invalid value for required argument 'ServiceUri'")
	}
	var resource RegistryWebook
	err := ctx.RegisterResource("azure:containerservice/registryWebook:RegistryWebook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryWebook gets an existing RegistryWebook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryWebook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryWebookState, opts ...pulumi.ResourceOption) (*RegistryWebook, error) {
	var resource RegistryWebook
	err := ctx.ReadResource("azure:containerservice/registryWebook:RegistryWebook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryWebook resources.
type registryWebookState struct {
	// A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
	Actions []string `pulumi:"actions"`
	// Custom headers that will be added to the webhook notifications request.
	CustomHeaders map[string]string `pulumi:"customHeaders"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
	RegistryName *string `pulumi:"registryName"`
	// The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
	Scope *string `pulumi:"scope"`
	// Specifies the service URI for the Webhook to post notifications.
	ServiceUri *string `pulumi:"serviceUri"`
	// Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
	Status *string           `pulumi:"status"`
	Tags   map[string]string `pulumi:"tags"`
}

type RegistryWebookState struct {
	// A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
	Actions pulumi.StringArrayInput
	// Custom headers that will be added to the webhook notifications request.
	CustomHeaders pulumi.StringMapInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
	RegistryName pulumi.StringPtrInput
	// The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
	Scope pulumi.StringPtrInput
	// Specifies the service URI for the Webhook to post notifications.
	ServiceUri pulumi.StringPtrInput
	// Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
	Status pulumi.StringPtrInput
	Tags   pulumi.StringMapInput
}

func (RegistryWebookState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryWebookState)(nil)).Elem()
}

type registryWebookArgs struct {
	// A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
	Actions []string `pulumi:"actions"`
	// Custom headers that will be added to the webhook notifications request.
	CustomHeaders map[string]string `pulumi:"customHeaders"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
	Scope *string `pulumi:"scope"`
	// Specifies the service URI for the Webhook to post notifications.
	ServiceUri string `pulumi:"serviceUri"`
	// Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
	Status *string           `pulumi:"status"`
	Tags   map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RegistryWebook resource.
type RegistryWebookArgs struct {
	// A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
	Actions pulumi.StringArrayInput
	// Custom headers that will be added to the webhook notifications request.
	CustomHeaders pulumi.StringMapInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
	RegistryName pulumi.StringInput
	// The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
	Scope pulumi.StringPtrInput
	// Specifies the service URI for the Webhook to post notifications.
	ServiceUri pulumi.StringInput
	// Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
	Status pulumi.StringPtrInput
	Tags   pulumi.StringMapInput
}

func (RegistryWebookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryWebookArgs)(nil)).Elem()
}

type RegistryWebookInput interface {
	pulumi.Input

	ToRegistryWebookOutput() RegistryWebookOutput
	ToRegistryWebookOutputWithContext(ctx context.Context) RegistryWebookOutput
}

func (*RegistryWebook) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryWebook)(nil))
}

func (i *RegistryWebook) ToRegistryWebookOutput() RegistryWebookOutput {
	return i.ToRegistryWebookOutputWithContext(context.Background())
}

func (i *RegistryWebook) ToRegistryWebookOutputWithContext(ctx context.Context) RegistryWebookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryWebookOutput)
}

func (i *RegistryWebook) ToRegistryWebookPtrOutput() RegistryWebookPtrOutput {
	return i.ToRegistryWebookPtrOutputWithContext(context.Background())
}

func (i *RegistryWebook) ToRegistryWebookPtrOutputWithContext(ctx context.Context) RegistryWebookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryWebookPtrOutput)
}

type RegistryWebookPtrInput interface {
	pulumi.Input

	ToRegistryWebookPtrOutput() RegistryWebookPtrOutput
	ToRegistryWebookPtrOutputWithContext(ctx context.Context) RegistryWebookPtrOutput
}

type registryWebookPtrType RegistryWebookArgs

func (*registryWebookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryWebook)(nil))
}

func (i *registryWebookPtrType) ToRegistryWebookPtrOutput() RegistryWebookPtrOutput {
	return i.ToRegistryWebookPtrOutputWithContext(context.Background())
}

func (i *registryWebookPtrType) ToRegistryWebookPtrOutputWithContext(ctx context.Context) RegistryWebookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryWebookPtrOutput)
}

// RegistryWebookArrayInput is an input type that accepts RegistryWebookArray and RegistryWebookArrayOutput values.
// You can construct a concrete instance of `RegistryWebookArrayInput` via:
//
//          RegistryWebookArray{ RegistryWebookArgs{...} }
type RegistryWebookArrayInput interface {
	pulumi.Input

	ToRegistryWebookArrayOutput() RegistryWebookArrayOutput
	ToRegistryWebookArrayOutputWithContext(context.Context) RegistryWebookArrayOutput
}

type RegistryWebookArray []RegistryWebookInput

func (RegistryWebookArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RegistryWebook)(nil))
}

func (i RegistryWebookArray) ToRegistryWebookArrayOutput() RegistryWebookArrayOutput {
	return i.ToRegistryWebookArrayOutputWithContext(context.Background())
}

func (i RegistryWebookArray) ToRegistryWebookArrayOutputWithContext(ctx context.Context) RegistryWebookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryWebookArrayOutput)
}

// RegistryWebookMapInput is an input type that accepts RegistryWebookMap and RegistryWebookMapOutput values.
// You can construct a concrete instance of `RegistryWebookMapInput` via:
//
//          RegistryWebookMap{ "key": RegistryWebookArgs{...} }
type RegistryWebookMapInput interface {
	pulumi.Input

	ToRegistryWebookMapOutput() RegistryWebookMapOutput
	ToRegistryWebookMapOutputWithContext(context.Context) RegistryWebookMapOutput
}

type RegistryWebookMap map[string]RegistryWebookInput

func (RegistryWebookMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RegistryWebook)(nil))
}

func (i RegistryWebookMap) ToRegistryWebookMapOutput() RegistryWebookMapOutput {
	return i.ToRegistryWebookMapOutputWithContext(context.Background())
}

func (i RegistryWebookMap) ToRegistryWebookMapOutputWithContext(ctx context.Context) RegistryWebookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryWebookMapOutput)
}

type RegistryWebookOutput struct {
	*pulumi.OutputState
}

func (RegistryWebookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryWebook)(nil))
}

func (o RegistryWebookOutput) ToRegistryWebookOutput() RegistryWebookOutput {
	return o
}

func (o RegistryWebookOutput) ToRegistryWebookOutputWithContext(ctx context.Context) RegistryWebookOutput {
	return o
}

func (o RegistryWebookOutput) ToRegistryWebookPtrOutput() RegistryWebookPtrOutput {
	return o.ToRegistryWebookPtrOutputWithContext(context.Background())
}

func (o RegistryWebookOutput) ToRegistryWebookPtrOutputWithContext(ctx context.Context) RegistryWebookPtrOutput {
	return o.ApplyT(func(v RegistryWebook) *RegistryWebook {
		return &v
	}).(RegistryWebookPtrOutput)
}

type RegistryWebookPtrOutput struct {
	*pulumi.OutputState
}

func (RegistryWebookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryWebook)(nil))
}

func (o RegistryWebookPtrOutput) ToRegistryWebookPtrOutput() RegistryWebookPtrOutput {
	return o
}

func (o RegistryWebookPtrOutput) ToRegistryWebookPtrOutputWithContext(ctx context.Context) RegistryWebookPtrOutput {
	return o
}

type RegistryWebookArrayOutput struct{ *pulumi.OutputState }

func (RegistryWebookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryWebook)(nil))
}

func (o RegistryWebookArrayOutput) ToRegistryWebookArrayOutput() RegistryWebookArrayOutput {
	return o
}

func (o RegistryWebookArrayOutput) ToRegistryWebookArrayOutputWithContext(ctx context.Context) RegistryWebookArrayOutput {
	return o
}

func (o RegistryWebookArrayOutput) Index(i pulumi.IntInput) RegistryWebookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryWebook {
		return vs[0].([]RegistryWebook)[vs[1].(int)]
	}).(RegistryWebookOutput)
}

type RegistryWebookMapOutput struct{ *pulumi.OutputState }

func (RegistryWebookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RegistryWebook)(nil))
}

func (o RegistryWebookMapOutput) ToRegistryWebookMapOutput() RegistryWebookMapOutput {
	return o
}

func (o RegistryWebookMapOutput) ToRegistryWebookMapOutputWithContext(ctx context.Context) RegistryWebookMapOutput {
	return o
}

func (o RegistryWebookMapOutput) MapIndex(k pulumi.StringInput) RegistryWebookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RegistryWebook {
		return vs[0].(map[string]RegistryWebook)[vs[1].(string)]
	}).(RegistryWebookOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryWebookOutput{})
	pulumi.RegisterOutputType(RegistryWebookPtrOutput{})
	pulumi.RegisterOutputType(RegistryWebookArrayOutput{})
	pulumi.RegisterOutputType(RegistryWebookMapOutput{})
}
