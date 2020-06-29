// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an EventGrid Topic
//
// > **Note:** at this time EventGrid Topic's are only available in a limited number of regions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventgrid"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US 2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventgrid.NewTopic(ctx, "exampleTopic", &eventgrid.TopicArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
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
// Deprecated: azure.eventhub.EventGridTopic has been deprecated in favor of azure.eventgrid.Topic
type EventGridTopic struct {
	pulumi.CustomResourceState

	// The Endpoint associated with the EventGrid Topic.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues EventGridTopicInputMappingDefaultValuesPtrOutput `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields EventGridTopicInputMappingFieldsPtrOutput `pulumi:"inputMappingFields"`
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema pulumi.StringPtrOutput `pulumi:"inputSchema"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Primary Shared Access Key associated with the EventGrid Topic.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Shared Access Key associated with the EventGrid Topic.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// A mapping of tags to assign to the resource.
	// ---
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewEventGridTopic registers a new resource with the given unique name, arguments, and options.
func NewEventGridTopic(ctx *pulumi.Context,
	name string, args *EventGridTopicArgs, opts ...pulumi.ResourceOption) (*EventGridTopic, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EventGridTopicArgs{}
	}
	var resource EventGridTopic
	err := ctx.RegisterResource("azure:eventhub/eventGridTopic:EventGridTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventGridTopic gets an existing EventGridTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventGridTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventGridTopicState, opts ...pulumi.ResourceOption) (*EventGridTopic, error) {
	var resource EventGridTopic
	err := ctx.ReadResource("azure:eventhub/eventGridTopic:EventGridTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventGridTopic resources.
type eventGridTopicState struct {
	// The Endpoint associated with the EventGrid Topic.
	Endpoint *string `pulumi:"endpoint"`
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues *EventGridTopicInputMappingDefaultValues `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields *EventGridTopicInputMappingFields `pulumi:"inputMappingFields"`
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema *string `pulumi:"inputSchema"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Primary Shared Access Key associated with the EventGrid Topic.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Shared Access Key associated with the EventGrid Topic.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// A mapping of tags to assign to the resource.
	// ---
	Tags map[string]string `pulumi:"tags"`
}

type EventGridTopicState struct {
	// The Endpoint associated with the EventGrid Topic.
	Endpoint pulumi.StringPtrInput
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues EventGridTopicInputMappingDefaultValuesPtrInput
	// A `inputMappingFields` block as defined below.
	InputMappingFields EventGridTopicInputMappingFieldsPtrInput
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Primary Shared Access Key associated with the EventGrid Topic.
	PrimaryAccessKey pulumi.StringPtrInput
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Shared Access Key associated with the EventGrid Topic.
	SecondaryAccessKey pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// ---
	Tags pulumi.StringMapInput
}

func (EventGridTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridTopicState)(nil)).Elem()
}

type eventGridTopicArgs struct {
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues *EventGridTopicInputMappingDefaultValues `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields *EventGridTopicInputMappingFields `pulumi:"inputMappingFields"`
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema *string `pulumi:"inputSchema"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	// ---
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EventGridTopic resource.
type EventGridTopicArgs struct {
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues EventGridTopicInputMappingDefaultValuesPtrInput
	// A `inputMappingFields` block as defined below.
	InputMappingFields EventGridTopicInputMappingFieldsPtrInput
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	// ---
	Tags pulumi.StringMapInput
}

func (EventGridTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridTopicArgs)(nil)).Elem()
}
