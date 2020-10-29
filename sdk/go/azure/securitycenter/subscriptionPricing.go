// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securitycenter

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the Pricing Tier for Azure Security Center in the current subscription.
//
// > **NOTE:** This resource requires the `Owner` permission on the Subscription.
//
// > **NOTE:** Deletion of this resource does not change or reset the pricing tier to `Free`
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
	if args == nil || args.Tier == nil {
		return nil, errors.New("missing required argument 'Tier'")
	}
	if args == nil {
		args = &SubscriptionPricingArgs{}
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
