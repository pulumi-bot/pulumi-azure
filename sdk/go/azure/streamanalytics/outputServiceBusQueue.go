// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Stream Analytics Output to a ServiceBus Queue.
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
// 		exampleResourceGroup, err := core.LookupResourceGroup(ctx, "azure:core:getResourceGroup", &core.LookupResourceGroupArgs{
// 			Name: "example-resources",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleJob, err := streamanalytics.LookupJob(ctx, "azure:streamanalytics:getJob", &streamanalytics.LookupJobArgs{
// 			Name:              "example-job",
// 			ResourceGroupName: azurerm_resource_group.Example.Name,
// 		})
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
// 		exampleQueue, err := servicebus.NewQueue(ctx, "exampleQueue", &servicebus.QueueArgs{
// 			ResourceGroupName:  pulumi.String(exampleResourceGroup.Name),
// 			NamespaceName:      exampleNamespace.Name,
// 			EnablePartitioning: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = streamanalytics.NewOutputServiceBusQueue(ctx, "exampleOutputServiceBusQueue", &streamanalytics.OutputServiceBusQueueArgs{
// 			StreamAnalyticsJobName: pulumi.String(exampleJob.Name),
// 			ResourceGroupName:      pulumi.String(exampleJob.ResourceGroupName),
// 			QueueName:              exampleQueue.Name,
// 			ServicebusNamespace:    exampleNamespace.Name,
// 			SharedAccessPolicyKey:  exampleNamespace.DefaultPrimaryKey,
// 			SharedAccessPolicyName: pulumi.String("RootManageSharedAccessKey"),
// 			Serialization: &streamanalytics.OutputServiceBusQueueSerializationArgs{
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
type OutputServiceBusQueue struct {
	pulumi.CustomResourceState

	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Service Bus Queue.
	QueueName pulumi.StringOutput `pulumi:"queueName"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputServiceBusQueueSerializationOutput `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringOutput `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringOutput `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringOutput `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
}

// NewOutputServiceBusQueue registers a new resource with the given unique name, arguments, and options.
func NewOutputServiceBusQueue(ctx *pulumi.Context,
	name string, args *OutputServiceBusQueueArgs, opts ...pulumi.ResourceOption) (*OutputServiceBusQueue, error) {
	if args == nil || args.QueueName == nil {
		return nil, errors.New("missing required argument 'QueueName'")
	}
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
	if args == nil {
		args = &OutputServiceBusQueueArgs{}
	}
	var resource OutputServiceBusQueue
	err := ctx.RegisterResource("azure:streamanalytics/outputServiceBusQueue:OutputServiceBusQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputServiceBusQueue gets an existing OutputServiceBusQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputServiceBusQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputServiceBusQueueState, opts ...pulumi.ResourceOption) (*OutputServiceBusQueue, error) {
	var resource OutputServiceBusQueue
	err := ctx.ReadResource("azure:streamanalytics/outputServiceBusQueue:OutputServiceBusQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputServiceBusQueue resources.
type outputServiceBusQueueState struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Service Bus Queue.
	QueueName *string `pulumi:"queueName"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization *OutputServiceBusQueueSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace *string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey *string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
}

type OutputServiceBusQueueState struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Service Bus Queue.
	QueueName pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `serialization` block as defined below.
	Serialization OutputServiceBusQueueSerializationPtrInput
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringPtrInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringPtrInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
}

func (OutputServiceBusQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputServiceBusQueueState)(nil)).Elem()
}

type outputServiceBusQueueArgs struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Service Bus Queue.
	QueueName string `pulumi:"queueName"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputServiceBusQueueSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
}

// The set of arguments for constructing a OutputServiceBusQueue resource.
type OutputServiceBusQueueArgs struct {
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Service Bus Queue.
	QueueName pulumi.StringInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `serialization` block as defined below.
	Serialization OutputServiceBusQueueSerializationInput
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
}

func (OutputServiceBusQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputServiceBusQueueArgs)(nil)).Elem()
}
