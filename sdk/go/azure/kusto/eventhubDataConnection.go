// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) EventHub Data Connection
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventhub"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/kusto"
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
// 		cluster, err := kusto.NewCluster(ctx, "cluster", &kusto.ClusterArgs{
// 			Location:          rg.Location,
// 			ResourceGroupName: rg.Name,
// 			Sku: &kusto.ClusterSkuArgs{
// 				Name:     pulumi.String("Standard_D13_v2"),
// 				Capacity: pulumi.Int(2),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		database, err := kusto.NewDatabase(ctx, "database", &kusto.DatabaseArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 			ClusterName:       cluster.Name,
// 			HotCachePeriod:    pulumi.String("P7D"),
// 			SoftDeletePeriod:  pulumi.String("P31D"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		eventhubNs, err := eventhub.NewEventHubNamespace(ctx, "eventhubNs", &eventhub.EventHubNamespaceArgs{
// 			Location:          rg.Location,
// 			ResourceGroupName: rg.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		eventhub, err := eventhub.NewEventHub(ctx, "eventhub", &eventhub.EventHubArgs{
// 			NamespaceName:     eventhubNs.Name,
// 			ResourceGroupName: rg.Name,
// 			PartitionCount:    pulumi.Int(1),
// 			MessageRetention:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		consumerGroup, err := eventhub.NewConsumerGroup(ctx, "consumerGroup", &eventhub.ConsumerGroupArgs{
// 			NamespaceName:     eventhubNs.Name,
// 			EventhubName:      eventhub.Name,
// 			ResourceGroupName: rg.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kusto.NewEventhubDataConnection(ctx, "eventhubConnection", &kusto.EventhubDataConnectionArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 			ClusterName:       cluster.Name,
// 			DatabaseName:      database.Name,
// 			EventhubId:        pulumi.Any(azurerm_eventhub.Evenhub.Id),
// 			ConsumerGroup:     consumerGroup.Name,
// 			TableName:         pulumi.String("my-table"),
// 			MappingRuleName:   pulumi.String("my-table-mapping"),
// 			DataFormat:        pulumi.String("JSON"),
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
// Kusto EventHub Data Connections can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:kusto/eventhubDataConnection:EventhubDataConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/eventHubConnection1
// ```
type EventhubDataConnection struct {
	pulumi.CustomResourceState

	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
	Compression pulumi.StringPtrOutput `pulumi:"compression"`
	// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup pulumi.StringOutput `pulumi:"consumerGroup"`
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat pulumi.StringPtrOutput `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Specifies a list of system properties for the Event Hub.
	EventSystemProperties pulumi.StringArrayOutput `pulumi:"eventSystemProperties"`
	// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventhubId pulumi.StringOutput `pulumi:"eventhubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
}

// NewEventhubDataConnection registers a new resource with the given unique name, arguments, and options.
func NewEventhubDataConnection(ctx *pulumi.Context,
	name string, args *EventhubDataConnectionArgs, opts ...pulumi.ResourceOption) (*EventhubDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ConsumerGroup == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroup'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.EventhubId == nil {
		return nil, errors.New("invalid value for required argument 'EventhubId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource EventhubDataConnection
	err := ctx.RegisterResource("azure:kusto/eventhubDataConnection:EventhubDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventhubDataConnection gets an existing EventhubDataConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventhubDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventhubDataConnectionState, opts ...pulumi.ResourceOption) (*EventhubDataConnection, error) {
	var resource EventhubDataConnection
	err := ctx.ReadResource("azure:kusto/eventhubDataConnection:EventhubDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventhubDataConnection resources.
type eventhubDataConnectionState struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName *string `pulumi:"clusterName"`
	// Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
	Compression *string `pulumi:"compression"`
	// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup *string `pulumi:"consumerGroup"`
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat *string `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// Specifies a list of system properties for the Event Hub.
	EventSystemProperties []string `pulumi:"eventSystemProperties"`
	// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventhubId *string `pulumi:"eventhubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `pulumi:"tableName"`
}

type EventhubDataConnectionState struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringPtrInput
	// Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
	Compression pulumi.StringPtrInput
	// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup pulumi.StringPtrInput
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat pulumi.StringPtrInput
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// Specifies a list of system properties for the Event Hub.
	EventSystemProperties pulumi.StringArrayInput
	// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventhubId pulumi.StringPtrInput
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrInput
	// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrInput
}

func (EventhubDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventhubDataConnectionState)(nil)).Elem()
}

type eventhubDataConnectionArgs struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName string `pulumi:"clusterName"`
	// Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
	Compression *string `pulumi:"compression"`
	// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup string `pulumi:"consumerGroup"`
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat *string `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// Specifies a list of system properties for the Event Hub.
	EventSystemProperties []string `pulumi:"eventSystemProperties"`
	// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventhubId string `pulumi:"eventhubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a EventhubDataConnection resource.
type EventhubDataConnectionArgs struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput
	// Specifies compression type for the connection. Allowed values: `GZip` and `None`. Defaults to `None`. Changing this forces a new resource to be created.
	Compression pulumi.StringPtrInput
	// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup pulumi.StringInput
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat pulumi.StringPtrInput
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// Specifies a list of system properties for the Event Hub.
	EventSystemProperties pulumi.StringArrayInput
	// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventhubId pulumi.StringInput
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrInput
	// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrInput
}

func (EventhubDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventhubDataConnectionArgs)(nil)).Elem()
}

type EventhubDataConnectionInput interface {
	pulumi.Input

	ToEventhubDataConnectionOutput() EventhubDataConnectionOutput
	ToEventhubDataConnectionOutputWithContext(ctx context.Context) EventhubDataConnectionOutput
}

func (*EventhubDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*EventhubDataConnection)(nil))
}

func (i *EventhubDataConnection) ToEventhubDataConnectionOutput() EventhubDataConnectionOutput {
	return i.ToEventhubDataConnectionOutputWithContext(context.Background())
}

func (i *EventhubDataConnection) ToEventhubDataConnectionOutputWithContext(ctx context.Context) EventhubDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventhubDataConnectionOutput)
}

func (i *EventhubDataConnection) ToEventhubDataConnectionPtrOutput() EventhubDataConnectionPtrOutput {
	return i.ToEventhubDataConnectionPtrOutputWithContext(context.Background())
}

func (i *EventhubDataConnection) ToEventhubDataConnectionPtrOutputWithContext(ctx context.Context) EventhubDataConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventhubDataConnectionPtrOutput)
}

type EventhubDataConnectionPtrInput interface {
	pulumi.Input

	ToEventhubDataConnectionPtrOutput() EventhubDataConnectionPtrOutput
	ToEventhubDataConnectionPtrOutputWithContext(ctx context.Context) EventhubDataConnectionPtrOutput
}

type eventhubDataConnectionPtrType EventhubDataConnectionArgs

func (*eventhubDataConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventhubDataConnection)(nil))
}

func (i *eventhubDataConnectionPtrType) ToEventhubDataConnectionPtrOutput() EventhubDataConnectionPtrOutput {
	return i.ToEventhubDataConnectionPtrOutputWithContext(context.Background())
}

func (i *eventhubDataConnectionPtrType) ToEventhubDataConnectionPtrOutputWithContext(ctx context.Context) EventhubDataConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventhubDataConnectionPtrOutput)
}

// EventhubDataConnectionArrayInput is an input type that accepts EventhubDataConnectionArray and EventhubDataConnectionArrayOutput values.
// You can construct a concrete instance of `EventhubDataConnectionArrayInput` via:
//
//          EventhubDataConnectionArray{ EventhubDataConnectionArgs{...} }
type EventhubDataConnectionArrayInput interface {
	pulumi.Input

	ToEventhubDataConnectionArrayOutput() EventhubDataConnectionArrayOutput
	ToEventhubDataConnectionArrayOutputWithContext(context.Context) EventhubDataConnectionArrayOutput
}

type EventhubDataConnectionArray []EventhubDataConnectionInput

func (EventhubDataConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EventhubDataConnection)(nil))
}

func (i EventhubDataConnectionArray) ToEventhubDataConnectionArrayOutput() EventhubDataConnectionArrayOutput {
	return i.ToEventhubDataConnectionArrayOutputWithContext(context.Background())
}

func (i EventhubDataConnectionArray) ToEventhubDataConnectionArrayOutputWithContext(ctx context.Context) EventhubDataConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventhubDataConnectionArrayOutput)
}

// EventhubDataConnectionMapInput is an input type that accepts EventhubDataConnectionMap and EventhubDataConnectionMapOutput values.
// You can construct a concrete instance of `EventhubDataConnectionMapInput` via:
//
//          EventhubDataConnectionMap{ "key": EventhubDataConnectionArgs{...} }
type EventhubDataConnectionMapInput interface {
	pulumi.Input

	ToEventhubDataConnectionMapOutput() EventhubDataConnectionMapOutput
	ToEventhubDataConnectionMapOutputWithContext(context.Context) EventhubDataConnectionMapOutput
}

type EventhubDataConnectionMap map[string]EventhubDataConnectionInput

func (EventhubDataConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EventhubDataConnection)(nil))
}

func (i EventhubDataConnectionMap) ToEventhubDataConnectionMapOutput() EventhubDataConnectionMapOutput {
	return i.ToEventhubDataConnectionMapOutputWithContext(context.Background())
}

func (i EventhubDataConnectionMap) ToEventhubDataConnectionMapOutputWithContext(ctx context.Context) EventhubDataConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventhubDataConnectionMapOutput)
}

type EventhubDataConnectionOutput struct {
	*pulumi.OutputState
}

func (EventhubDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventhubDataConnection)(nil))
}

func (o EventhubDataConnectionOutput) ToEventhubDataConnectionOutput() EventhubDataConnectionOutput {
	return o
}

func (o EventhubDataConnectionOutput) ToEventhubDataConnectionOutputWithContext(ctx context.Context) EventhubDataConnectionOutput {
	return o
}

func (o EventhubDataConnectionOutput) ToEventhubDataConnectionPtrOutput() EventhubDataConnectionPtrOutput {
	return o.ToEventhubDataConnectionPtrOutputWithContext(context.Background())
}

func (o EventhubDataConnectionOutput) ToEventhubDataConnectionPtrOutputWithContext(ctx context.Context) EventhubDataConnectionPtrOutput {
	return o.ApplyT(func(v EventhubDataConnection) *EventhubDataConnection {
		return &v
	}).(EventhubDataConnectionPtrOutput)
}

type EventhubDataConnectionPtrOutput struct {
	*pulumi.OutputState
}

func (EventhubDataConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventhubDataConnection)(nil))
}

func (o EventhubDataConnectionPtrOutput) ToEventhubDataConnectionPtrOutput() EventhubDataConnectionPtrOutput {
	return o
}

func (o EventhubDataConnectionPtrOutput) ToEventhubDataConnectionPtrOutputWithContext(ctx context.Context) EventhubDataConnectionPtrOutput {
	return o
}

type EventhubDataConnectionArrayOutput struct{ *pulumi.OutputState }

func (EventhubDataConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventhubDataConnection)(nil))
}

func (o EventhubDataConnectionArrayOutput) ToEventhubDataConnectionArrayOutput() EventhubDataConnectionArrayOutput {
	return o
}

func (o EventhubDataConnectionArrayOutput) ToEventhubDataConnectionArrayOutputWithContext(ctx context.Context) EventhubDataConnectionArrayOutput {
	return o
}

func (o EventhubDataConnectionArrayOutput) Index(i pulumi.IntInput) EventhubDataConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventhubDataConnection {
		return vs[0].([]EventhubDataConnection)[vs[1].(int)]
	}).(EventhubDataConnectionOutput)
}

type EventhubDataConnectionMapOutput struct{ *pulumi.OutputState }

func (EventhubDataConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventhubDataConnection)(nil))
}

func (o EventhubDataConnectionMapOutput) ToEventhubDataConnectionMapOutput() EventhubDataConnectionMapOutput {
	return o
}

func (o EventhubDataConnectionMapOutput) ToEventhubDataConnectionMapOutputWithContext(ctx context.Context) EventhubDataConnectionMapOutput {
	return o
}

func (o EventhubDataConnectionMapOutput) MapIndex(k pulumi.StringInput) EventhubDataConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventhubDataConnection {
		return vs[0].(map[string]EventhubDataConnection)[vs[1].(string)]
	}).(EventhubDataConnectionOutput)
}

func init() {
	pulumi.RegisterOutputType(EventhubDataConnectionOutput{})
	pulumi.RegisterOutputType(EventhubDataConnectionPtrOutput{})
	pulumi.RegisterOutputType(EventhubDataConnectionArrayOutput{})
	pulumi.RegisterOutputType(EventhubDataConnectionMapOutput{})
}
