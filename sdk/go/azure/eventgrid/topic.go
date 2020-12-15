// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
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
// ## Import
//
// EventGrid Topic's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventgrid/topic:Topic topic1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1
// ```
type Topic struct {
	pulumi.CustomResourceState

	// The Endpoint associated with the EventGrid Topic.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues TopicInputMappingDefaultValuesPtrOutput `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields TopicInputMappingFieldsPtrOutput `pulumi:"inputMappingFields"`
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
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:eventhub/eventGridTopic:EventGridTopic"),
		},
	})
	opts = append(opts, aliases)
	var resource Topic
	err := ctx.RegisterResource("azure:eventgrid/topic:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azure:eventgrid/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// The Endpoint associated with the EventGrid Topic.
	Endpoint *string `pulumi:"endpoint"`
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues *TopicInputMappingDefaultValues `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields *TopicInputMappingFields `pulumi:"inputMappingFields"`
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
	Tags map[string]string `pulumi:"tags"`
}

type TopicState struct {
	// The Endpoint associated with the EventGrid Topic.
	Endpoint pulumi.StringPtrInput
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues TopicInputMappingDefaultValuesPtrInput
	// A `inputMappingFields` block as defined below.
	InputMappingFields TopicInputMappingFieldsPtrInput
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
	Tags pulumi.StringMapInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues *TopicInputMappingDefaultValues `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields *TopicInputMappingFields `pulumi:"inputMappingFields"`
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema *string `pulumi:"inputSchema"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues TopicInputMappingDefaultValuesPtrInput
	// A `inputMappingFields` block as defined below.
	InputMappingFields TopicInputMappingFieldsPtrInput
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

func (i *Topic) ToTopicPtrOutput() TopicPtrOutput {
	return i.ToTopicPtrOutputWithContext(context.Background())
}

func (i *Topic) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPtrOutput)
}

type TopicPtrInput interface {
	pulumi.Input

	ToTopicPtrOutput() TopicPtrOutput
	ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput
}

type TopicOutput struct {
	*pulumi.OutputState
}

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

type TopicPtrOutput struct {
	*pulumi.OutputState
}

func (TopicPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil))
}

func (o TopicPtrOutput) ToTopicPtrOutput() TopicPtrOutput {
	return o
}

func (o TopicPtrOutput) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TopicOutput{})
	pulumi.RegisterOutputType(TopicPtrOutput{})
}
