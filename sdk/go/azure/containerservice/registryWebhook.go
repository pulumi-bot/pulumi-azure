// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
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
//  $ pulumi import azure:containerservice/registryWebhook:RegistryWebhook example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/mygroup1/providers/Microsoft.ContainerRegistry/registries/myregistry1/webhooks/mywebhook1
// ```
type RegistryWebhook struct {
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

// NewRegistryWebhook registers a new resource with the given unique name, arguments, and options.
func NewRegistryWebhook(ctx *pulumi.Context,
	name string, args *RegistryWebhookArgs, opts ...pulumi.ResourceOption) (*RegistryWebhook, error) {
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:containerservice/registryWebook:RegistryWebook"),
		},
	})
	opts = append(opts, aliases)
	var resource RegistryWebhook
	err := ctx.RegisterResource("azure:containerservice/registryWebhook:RegistryWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryWebhook gets an existing RegistryWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryWebhookState, opts ...pulumi.ResourceOption) (*RegistryWebhook, error) {
	var resource RegistryWebhook
	err := ctx.ReadResource("azure:containerservice/registryWebhook:RegistryWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryWebhook resources.
type registryWebhookState struct {
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

type RegistryWebhookState struct {
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

func (RegistryWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryWebhookState)(nil)).Elem()
}

type registryWebhookArgs struct {
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

// The set of arguments for constructing a RegistryWebhook resource.
type RegistryWebhookArgs struct {
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

func (RegistryWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryWebhookArgs)(nil)).Elem()
}

type RegistryWebhookInput interface {
	pulumi.Input

	ToRegistryWebhookOutput() RegistryWebhookOutput
	ToRegistryWebhookOutputWithContext(ctx context.Context) RegistryWebhookOutput
}

func (RegistryWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryWebhook)(nil)).Elem()
}

func (i RegistryWebhook) ToRegistryWebhookOutput() RegistryWebhookOutput {
	return i.ToRegistryWebhookOutputWithContext(context.Background())
}

func (i RegistryWebhook) ToRegistryWebhookOutputWithContext(ctx context.Context) RegistryWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryWebhookOutput)
}

type RegistryWebhookOutput struct {
	*pulumi.OutputState
}

func (RegistryWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryWebhookOutput)(nil)).Elem()
}

func (o RegistryWebhookOutput) ToRegistryWebhookOutput() RegistryWebhookOutput {
	return o
}

func (o RegistryWebhookOutput) ToRegistryWebhookOutputWithContext(ctx context.Context) RegistryWebhookOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegistryWebhookOutput{})
}
