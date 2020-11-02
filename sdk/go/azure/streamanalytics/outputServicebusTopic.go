// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Stream Analytics Output to a ServiceBus Topic.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/servicebus"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/streamanalytics"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleJob, err := streamanalytics.LookupJob(ctx, &streamanalytics.LookupJobArgs{
// 			Name:              "example-job",
// 			ResourceGroupName: azurerm_resource_group.Example.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleNamespace, err := servicebus.NewNamespace(ctx, "exampleNamespace", &servicebus.NamespaceArgs{
// 			Location:          pulumi.String(exampleResourceGroup.Location),
// 			ResourceGroupName: pulumi.String(exampleResourceGroup.Name),
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTopic, err := servicebus.NewTopic(ctx, "exampleTopic", &servicebus.TopicArgs{
// 			ResourceGroupName:  pulumi.String(exampleResourceGroup.Name),
// 			NamespaceName:      exampleNamespace.Name,
// 			EnablePartitioning: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = streamanalytics.NewOutputServicebusTopic(ctx, "exampleOutputServicebusTopic", &streamanalytics.OutputServicebusTopicArgs{
// 			StreamAnalyticsJobName: pulumi.String(exampleJob.Name),
// 			ResourceGroupName:      pulumi.String(exampleJob.ResourceGroupName),
// 			TopicName:              exampleTopic.Name,
// 			ServicebusNamespace:    exampleNamespace.Name,
// 			SharedAccessPolicyKey:  exampleNamespace.DefaultPrimaryKey,
// 			SharedAccessPolicyName: pulumi.String("RootManageSharedAccessKey"),
// 			Serialization: &streamanalytics.OutputServicebusTopicSerializationArgs{
// 				Format: pulumi.String("Avro"),
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
// Stream Analytics Output ServiceBus Topic's can be imported using the `resource id`, e.g.
type OutputServicebusTopic struct {
	pulumi.CustomResourceState

	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerializationOutput `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringOutput `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringOutput `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringOutput `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
	// The name of the Service Bus Topic.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewOutputServicebusTopic registers a new resource with the given unique name, arguments, and options.
func NewOutputServicebusTopic(ctx *pulumi.Context,
	name string, args *OutputServicebusTopicArgs, opts ...pulumi.ResourceOption) (*OutputServicebusTopic, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Serialization == nil {
		return nil, errors.New("missing required argument 'Serialization'")
	}
	if args == nil || args.ServicebusNamespace == nil {
		return nil, errors.New("missing required argument 'ServicebusNamespace'")
	}
	if args == nil || args.SharedAccessPolicyKey == nil {
		return nil, errors.New("missing required argument 'SharedAccessPolicyKey'")
	}
	if args == nil || args.SharedAccessPolicyName == nil {
		return nil, errors.New("missing required argument 'SharedAccessPolicyName'")
	}
	if args == nil || args.StreamAnalyticsJobName == nil {
		return nil, errors.New("missing required argument 'StreamAnalyticsJobName'")
	}
	if args == nil || args.TopicName == nil {
		return nil, errors.New("missing required argument 'TopicName'")
	}
	if args == nil {
		args = &OutputServicebusTopicArgs{}
	}
	var resource OutputServicebusTopic
	err := ctx.RegisterResource("azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputServicebusTopic gets an existing OutputServicebusTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputServicebusTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputServicebusTopicState, opts ...pulumi.ResourceOption) (*OutputServicebusTopic, error) {
	var resource OutputServicebusTopic
	err := ctx.ReadResource("azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputServicebusTopic resources.
type outputServicebusTopicState struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization *OutputServicebusTopicSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace *string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey *string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
	// The name of the Service Bus Topic.
	TopicName *string `pulumi:"topicName"`
}

type OutputServicebusTopicState struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerializationPtrInput
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringPtrInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringPtrInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
	// The name of the Service Bus Topic.
	TopicName pulumi.StringPtrInput
}

func (OutputServicebusTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputServicebusTopicState)(nil)).Elem()
}

type outputServicebusTopicArgs struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
	// The name of the Service Bus Topic.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a OutputServicebusTopic resource.
type OutputServicebusTopicArgs struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `serialization` block as defined below.
	Serialization OutputServicebusTopicSerializationInput
	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
	// The name of the Service Bus Topic.
	TopicName pulumi.StringInput
}

func (OutputServicebusTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputServicebusTopicArgs)(nil)).Elem()
}
