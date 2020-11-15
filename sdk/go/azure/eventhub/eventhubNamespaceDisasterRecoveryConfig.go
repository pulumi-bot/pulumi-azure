// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Disaster Recovery Config for an Event Hub Namespace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventhub"
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
// 		primary, err := eventhub.NewEventHubNamespace(ctx, "primary", &eventhub.EventHubNamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondary, err := eventhub.NewEventHubNamespace(ctx, "secondary", &eventhub.EventHubNamespaceArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventhub.NewEventhubNamespaceDisasterRecoveryConfig(ctx, "exampleEventhubNamespaceDisasterRecoveryConfig", &eventhub.EventhubNamespaceDisasterRecoveryConfigArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			NamespaceName:      primary.Name,
// 			PartnerNamespaceId: secondary.ID(),
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
// EventHubs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventhub/eventhubNamespaceDisasterRecoveryConfig:EventhubNamespaceDisasterRecoveryConfig config1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/disasterRecoveryConfigs/config1
// ```
type EventhubNamespaceDisasterRecoveryConfig struct {
	pulumi.CustomResourceState

	// An alternate name to use when the Disaster Recovery Config's name is the same as the replicated namespace's name.
	AlternateName pulumi.StringPtrOutput `pulumi:"alternateName"`
	// Specifies the name of the Disaster Recovery Config. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the primary EventHub Namespace to replicate. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The ID of the EventHub Namespace to replicate to.
	PartnerNamespaceId pulumi.StringOutput `pulumi:"partnerNamespaceId"`
	// The name of the resource group in which the Disaster Recovery Config exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewEventhubNamespaceDisasterRecoveryConfig registers a new resource with the given unique name, arguments, and options.
func NewEventhubNamespaceDisasterRecoveryConfig(ctx *pulumi.Context,
	name string, args *EventhubNamespaceDisasterRecoveryConfigArgs, opts ...pulumi.ResourceOption) (*EventhubNamespaceDisasterRecoveryConfig, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.PartnerNamespaceId == nil {
		return nil, errors.New("missing required argument 'PartnerNamespaceId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EventhubNamespaceDisasterRecoveryConfigArgs{}
	}
	var resource EventhubNamespaceDisasterRecoveryConfig
	err := ctx.RegisterResource("azure:eventhub/eventhubNamespaceDisasterRecoveryConfig:EventhubNamespaceDisasterRecoveryConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventhubNamespaceDisasterRecoveryConfig gets an existing EventhubNamespaceDisasterRecoveryConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventhubNamespaceDisasterRecoveryConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventhubNamespaceDisasterRecoveryConfigState, opts ...pulumi.ResourceOption) (*EventhubNamespaceDisasterRecoveryConfig, error) {
	var resource EventhubNamespaceDisasterRecoveryConfig
	err := ctx.ReadResource("azure:eventhub/eventhubNamespaceDisasterRecoveryConfig:EventhubNamespaceDisasterRecoveryConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventhubNamespaceDisasterRecoveryConfig resources.
type eventhubNamespaceDisasterRecoveryConfigState struct {
	// An alternate name to use when the Disaster Recovery Config's name is the same as the replicated namespace's name.
	AlternateName *string `pulumi:"alternateName"`
	// Specifies the name of the Disaster Recovery Config. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the primary EventHub Namespace to replicate. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// The ID of the EventHub Namespace to replicate to.
	PartnerNamespaceId *string `pulumi:"partnerNamespaceId"`
	// The name of the resource group in which the Disaster Recovery Config exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type EventhubNamespaceDisasterRecoveryConfigState struct {
	// An alternate name to use when the Disaster Recovery Config's name is the same as the replicated namespace's name.
	AlternateName pulumi.StringPtrInput
	// Specifies the name of the Disaster Recovery Config. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the primary EventHub Namespace to replicate. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// The ID of the EventHub Namespace to replicate to.
	PartnerNamespaceId pulumi.StringPtrInput
	// The name of the resource group in which the Disaster Recovery Config exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (EventhubNamespaceDisasterRecoveryConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventhubNamespaceDisasterRecoveryConfigState)(nil)).Elem()
}

type eventhubNamespaceDisasterRecoveryConfigArgs struct {
	// An alternate name to use when the Disaster Recovery Config's name is the same as the replicated namespace's name.
	AlternateName *string `pulumi:"alternateName"`
	// Specifies the name of the Disaster Recovery Config. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the primary EventHub Namespace to replicate. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// The ID of the EventHub Namespace to replicate to.
	PartnerNamespaceId string `pulumi:"partnerNamespaceId"`
	// The name of the resource group in which the Disaster Recovery Config exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EventhubNamespaceDisasterRecoveryConfig resource.
type EventhubNamespaceDisasterRecoveryConfigArgs struct {
	// An alternate name to use when the Disaster Recovery Config's name is the same as the replicated namespace's name.
	AlternateName pulumi.StringPtrInput
	// Specifies the name of the Disaster Recovery Config. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the primary EventHub Namespace to replicate. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// The ID of the EventHub Namespace to replicate to.
	PartnerNamespaceId pulumi.StringInput
	// The name of the resource group in which the Disaster Recovery Config exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (EventhubNamespaceDisasterRecoveryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventhubNamespaceDisasterRecoveryConfigArgs)(nil)).Elem()
}

type EventhubNamespaceDisasterRecoveryConfigInput interface {
	pulumi.Input

	ToEventhubNamespaceDisasterRecoveryConfigOutput() EventhubNamespaceDisasterRecoveryConfigOutput
	ToEventhubNamespaceDisasterRecoveryConfigOutputWithContext(ctx context.Context) EventhubNamespaceDisasterRecoveryConfigOutput
}

func (EventhubNamespaceDisasterRecoveryConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*EventhubNamespaceDisasterRecoveryConfig)(nil)).Elem()
}

func (i EventhubNamespaceDisasterRecoveryConfig) ToEventhubNamespaceDisasterRecoveryConfigOutput() EventhubNamespaceDisasterRecoveryConfigOutput {
	return i.ToEventhubNamespaceDisasterRecoveryConfigOutputWithContext(context.Background())
}

func (i EventhubNamespaceDisasterRecoveryConfig) ToEventhubNamespaceDisasterRecoveryConfigOutputWithContext(ctx context.Context) EventhubNamespaceDisasterRecoveryConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventhubNamespaceDisasterRecoveryConfigOutput)
}

type EventhubNamespaceDisasterRecoveryConfigOutput struct {
	*pulumi.OutputState
}

func (EventhubNamespaceDisasterRecoveryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventhubNamespaceDisasterRecoveryConfigOutput)(nil)).Elem()
}

func (o EventhubNamespaceDisasterRecoveryConfigOutput) ToEventhubNamespaceDisasterRecoveryConfigOutput() EventhubNamespaceDisasterRecoveryConfigOutput {
	return o
}

func (o EventhubNamespaceDisasterRecoveryConfigOutput) ToEventhubNamespaceDisasterRecoveryConfigOutputWithContext(ctx context.Context) EventhubNamespaceDisasterRecoveryConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventhubNamespaceDisasterRecoveryConfigOutput{})
}
