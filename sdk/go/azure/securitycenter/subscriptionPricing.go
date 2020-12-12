// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securitycenter

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the Pricing Tier for Azure Security Center in the current subscription.
//
// > **NOTE:** This resource requires the `Owner` permission on the Subscription.
//
// > **NOTE:** Deletion of this resource does not change or reset the pricing tier to `Free`
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/securitycenter"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := securitycenter.NewSubscriptionPricing(ctx, "example", &securitycenter.SubscriptionPricingArgs{
// 			ResourceType: pulumi.String("VirtualMachines"),
// 			Tier:         pulumi.String("Standard"),
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
// The pricing tier can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:securitycenter/subscriptionPricing:SubscriptionPricing example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/pricings/<resource_type>
// ```
type SubscriptionPricing struct {
	pulumi.CustomResourceState

	// The resource type this setting affects. Possible values are `AppServices`, `ContainerRegistry`, `KeyVaults`, `KubernetesService`, `SqlServers`, `SqlServerVirtualMachines`, `StorageAccounts`, and `VirtualMachines`.
	ResourceType pulumi.StringPtrOutput `pulumi:"resourceType"`
	// The pricing tier to use. Possible values are `Free` and `Standard`.
	Tier pulumi.StringOutput `pulumi:"tier"`
}

// NewSubscriptionPricing registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionPricing(ctx *pulumi.Context,
	name string, args *SubscriptionPricingArgs, opts ...pulumi.ResourceOption) (*SubscriptionPricing, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	var resource SubscriptionPricing
	err := ctx.RegisterResource("azure:securitycenter/subscriptionPricing:SubscriptionPricing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionPricing gets an existing SubscriptionPricing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionPricing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionPricingState, opts ...pulumi.ResourceOption) (*SubscriptionPricing, error) {
	var resource SubscriptionPricing
	err := ctx.ReadResource("azure:securitycenter/subscriptionPricing:SubscriptionPricing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionPricing resources.
type subscriptionPricingState struct {
	// The resource type this setting affects. Possible values are `AppServices`, `ContainerRegistry`, `KeyVaults`, `KubernetesService`, `SqlServers`, `SqlServerVirtualMachines`, `StorageAccounts`, and `VirtualMachines`.
	ResourceType *string `pulumi:"resourceType"`
	// The pricing tier to use. Possible values are `Free` and `Standard`.
	Tier *string `pulumi:"tier"`
}

type SubscriptionPricingState struct {
	// The resource type this setting affects. Possible values are `AppServices`, `ContainerRegistry`, `KeyVaults`, `KubernetesService`, `SqlServers`, `SqlServerVirtualMachines`, `StorageAccounts`, and `VirtualMachines`.
	ResourceType pulumi.StringPtrInput
	// The pricing tier to use. Possible values are `Free` and `Standard`.
	Tier pulumi.StringPtrInput
}

func (SubscriptionPricingState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionPricingState)(nil)).Elem()
}

type subscriptionPricingArgs struct {
	// The resource type this setting affects. Possible values are `AppServices`, `ContainerRegistry`, `KeyVaults`, `KubernetesService`, `SqlServers`, `SqlServerVirtualMachines`, `StorageAccounts`, and `VirtualMachines`.
	ResourceType *string `pulumi:"resourceType"`
	// The pricing tier to use. Possible values are `Free` and `Standard`.
	Tier string `pulumi:"tier"`
}

// The set of arguments for constructing a SubscriptionPricing resource.
type SubscriptionPricingArgs struct {
	// The resource type this setting affects. Possible values are `AppServices`, `ContainerRegistry`, `KeyVaults`, `KubernetesService`, `SqlServers`, `SqlServerVirtualMachines`, `StorageAccounts`, and `VirtualMachines`.
	ResourceType pulumi.StringPtrInput
	// The pricing tier to use. Possible values are `Free` and `Standard`.
	Tier pulumi.StringInput
}

func (SubscriptionPricingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionPricingArgs)(nil)).Elem()
}

type SubscriptionPricingInput interface {
	pulumi.Input

	ToSubscriptionPricingOutput() SubscriptionPricingOutput
	ToSubscriptionPricingOutputWithContext(ctx context.Context) SubscriptionPricingOutput
}

func (*SubscriptionPricing) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionPricing)(nil))
}

func (i *SubscriptionPricing) ToSubscriptionPricingOutput() SubscriptionPricingOutput {
	return i.ToSubscriptionPricingOutputWithContext(context.Background())
}

func (i *SubscriptionPricing) ToSubscriptionPricingOutputWithContext(ctx context.Context) SubscriptionPricingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionPricingOutput)
}

func (i *SubscriptionPricing) ToSubscriptionPricingPtrOutput() SubscriptionPricingPtrOutput {
	return i.ToSubscriptionPricingPtrOutputWithContext(context.Background())
}

func (i *SubscriptionPricing) ToSubscriptionPricingPtrOutputWithContext(ctx context.Context) SubscriptionPricingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionPricingPtrOutput)
}

type SubscriptionPricingPtrInput interface {
	pulumi.Input

	ToSubscriptionPricingPtrOutput() SubscriptionPricingPtrOutput
	ToSubscriptionPricingPtrOutputWithContext(ctx context.Context) SubscriptionPricingPtrOutput
}

type SubscriptionPricingOutput struct {
	*pulumi.OutputState
}

func (SubscriptionPricingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionPricing)(nil))
}

func (o SubscriptionPricingOutput) ToSubscriptionPricingOutput() SubscriptionPricingOutput {
	return o
}

func (o SubscriptionPricingOutput) ToSubscriptionPricingOutputWithContext(ctx context.Context) SubscriptionPricingOutput {
	return o
}

type SubscriptionPricingPtrOutput struct {
	*pulumi.OutputState
}

func (SubscriptionPricingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionPricing)(nil))
}

func (o SubscriptionPricingPtrOutput) ToSubscriptionPricingPtrOutput() SubscriptionPricingPtrOutput {
	return o
}

func (o SubscriptionPricingPtrOutput) ToSubscriptionPricingPtrOutputWithContext(ctx context.Context) SubscriptionPricingPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubscriptionPricingOutput{})
	pulumi.RegisterOutputType(SubscriptionPricingPtrOutput{})
}
