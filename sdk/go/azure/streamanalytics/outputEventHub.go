// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Stream Analytics Output to an EventHub.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventhub"
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
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          pulumi.String(exampleResourceGroup.Location),
// 			ResourceGroupName: pulumi.String(exampleResourceGroup.Name),
// 			Sku:               pulumi.String("Standard"),
// 			Capacity:          pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHub, err := eventhub.NewEventHub(ctx, "exampleEventHub", &eventhub.EventHubArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			ResourceGroupName: pulumi.String(exampleResourceGroup.Name),
// 			PartitionCount:    pulumi.Int(2),
// 			MessageRetention:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = streamanalytics.NewOutputEventHub(ctx, "exampleOutputEventHub", &streamanalytics.OutputEventHubArgs{
// 			StreamAnalyticsJobName: pulumi.String(exampleJob.Name),
// 			ResourceGroupName:      pulumi.String(exampleJob.ResourceGroupName),
// 			EventhubName:           exampleEventHub.Name,
// 			ServicebusNamespace:    exampleEventHubNamespace.Name,
// 			SharedAccessPolicyKey:  exampleEventHubNamespace.DefaultPrimaryKey,
// 			SharedAccessPolicyName: pulumi.String("RootManageSharedAccessKey"),
// 			Serialization: &streamanalytics.OutputEventHubSerializationArgs{
// 				Type: pulumi.String("Avro"),
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
// Stream Analytics Outputs to an EventHub can be imported using the `resource id`, e.g.
type OutputEventHub struct {
	pulumi.CustomResourceState

	// The name of the Event Hub.
	EventhubName pulumi.StringOutput `pulumi:"eventhubName"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputEventHubSerializationOutput `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringOutput `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringOutput `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringOutput `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
}

// NewOutputEventHub registers a new resource with the given unique name, arguments, and options.
func NewOutputEventHub(ctx *pulumi.Context,
	name string, args *OutputEventHubArgs, opts ...pulumi.ResourceOption) (*OutputEventHub, error) {
	if args == nil || args.EventhubName == nil {
		return nil, errors.New("missing required argument 'EventhubName'")
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
		args = &OutputEventHubArgs{}
	}
	var resource OutputEventHub
	err := ctx.RegisterResource("azure:streamanalytics/outputEventHub:OutputEventHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputEventHub gets an existing OutputEventHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputEventHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputEventHubState, opts ...pulumi.ResourceOption) (*OutputEventHub, error) {
	var resource OutputEventHub
	err := ctx.ReadResource("azure:streamanalytics/outputEventHub:OutputEventHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputEventHub resources.
type outputEventHubState struct {
	// The name of the Event Hub.
	EventhubName *string `pulumi:"eventhubName"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization *OutputEventHubSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace *string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey *string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
}

type OutputEventHubState struct {
	// The name of the Event Hub.
	EventhubName pulumi.StringPtrInput
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `serialization` block as defined below.
	Serialization OutputEventHubSerializationPtrInput
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringPtrInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringPtrInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
}

func (OutputEventHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputEventHubState)(nil)).Elem()
}

type outputEventHubArgs struct {
	// The name of the Event Hub.
	EventhubName string `pulumi:"eventhubName"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputEventHubSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
}

// The set of arguments for constructing a OutputEventHub resource.
type OutputEventHubArgs struct {
	// The name of the Event Hub.
	EventhubName pulumi.StringInput
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `serialization` block as defined below.
	Serialization OutputEventHubSerializationInput
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
}

func (OutputEventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputEventHubArgs)(nil)).Elem()
}
