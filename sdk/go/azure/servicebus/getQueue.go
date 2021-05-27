// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Service Bus Queue.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/servicebus"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := servicebus.LookupQueue(ctx, &servicebus.LookupQueueArgs{
// 			Name:              "existing",
// 			ResourceGroupName: "existing",
// 			NamespaceName:     "existing",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupQueue(ctx *pulumi.Context, args *LookupQueueArgs, opts ...pulumi.InvokeOption) (*LookupQueueResult, error) {
	var rv LookupQueueResult
	err := ctx.Invoke("azure:servicebus/getQueue:getQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueue.
type LookupQueueArgs struct {
	// The name of this Service Bus Queue.
	Name string `pulumi:"name"`
	// The name of the ServiceBus Namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the Resource Group where the Service Bus Queue exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getQueue.
type LookupQueueResult struct {
	// The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle string `pulumi:"autoDeleteOnIdle"`
	// Boolean flag which controls whether the Queue has dead letter support when a message expires.
	DeadLetteringOnMessageExpiration bool `pulumi:"deadLetteringOnMessageExpiration"`
	// The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on a message itself.
	DefaultMessageTtl string `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which duplicates can be detected.
	DuplicateDetectionHistoryTimeWindow string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls whether server-side batched operations are enabled.
	EnableBatchedOperations bool `pulumi:"enableBatchedOperations"`
	// Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage.
	EnableExpress bool `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers.
	EnablePartitioning bool `pulumi:"enablePartitioning"`
	// The name of a Queue or Topic to automatically forward dead lettered messages to.
	ForwardDeadLetteredMessagesTo string `pulumi:"forwardDeadLetteredMessagesTo"`
	// The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
	ForwardTo string `pulumi:"forwardTo"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers.
	LockDuration string `pulumi:"lockDuration"`
	// Integer value which controls when a message is automatically dead lettered.
	MaxDeliveryCount int `pulumi:"maxDeliveryCount"`
	// Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes int    `pulumi:"maxSizeInMegabytes"`
	Name               string `pulumi:"name"`
	NamespaceName      string `pulumi:"namespaceName"`
	// Boolean flag which controls whether the Queue requires duplicate detection.
	RequiresDuplicateDetection bool `pulumi:"requiresDuplicateDetection"`
	// Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages.
	RequiresSession   bool   `pulumi:"requiresSession"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`.
	Status string `pulumi:"status"`
}

func LookupQueueOutput(ctx *pulumi.Context, args LookupQueueOutputArgs, opts ...pulumi.InvokeOption) LookupQueueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueueResult, error) {
			args := v.(LookupQueueArgs)
			r, err := LookupQueue(ctx, &args, opts...)
			return *r, err
		}).(LookupQueueResultOutput)
}

// A collection of arguments for invoking getQueue.
type LookupQueueOutputArgs struct {
	// The name of this Service Bus Queue.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the ServiceBus Namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the Resource Group where the Service Bus Queue exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueArgs)(nil)).Elem()
}

// A collection of values returned by getQueue.
type LookupQueueResultOutput struct{ *pulumi.OutputState }

func (LookupQueueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueResult)(nil)).Elem()
}

func (o LookupQueueResultOutput) ToLookupQueueResultOutput() LookupQueueResultOutput {
	return o
}

func (o LookupQueueResultOutput) ToLookupQueueResultOutputWithContext(ctx context.Context) LookupQueueResultOutput {
	return o
}

// The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
func (o LookupQueueResultOutput) AutoDeleteOnIdle() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.AutoDeleteOnIdle }).(pulumi.StringOutput)
}

// Boolean flag which controls whether the Queue has dead letter support when a message expires.
func (o LookupQueueResultOutput) DeadLetteringOnMessageExpiration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQueueResult) bool { return v.DeadLetteringOnMessageExpiration }).(pulumi.BoolOutput)
}

// The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on a message itself.
func (o LookupQueueResultOutput) DefaultMessageTtl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.DefaultMessageTtl }).(pulumi.StringOutput)
}

// The ISO 8601 timespan duration during which duplicates can be detected.
func (o LookupQueueResultOutput) DuplicateDetectionHistoryTimeWindow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.DuplicateDetectionHistoryTimeWindow }).(pulumi.StringOutput)
}

// Boolean flag which controls whether server-side batched operations are enabled.
func (o LookupQueueResultOutput) EnableBatchedOperations() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQueueResult) bool { return v.EnableBatchedOperations }).(pulumi.BoolOutput)
}

// Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage.
func (o LookupQueueResultOutput) EnableExpress() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQueueResult) bool { return v.EnableExpress }).(pulumi.BoolOutput)
}

// Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers.
func (o LookupQueueResultOutput) EnablePartitioning() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQueueResult) bool { return v.EnablePartitioning }).(pulumi.BoolOutput)
}

// The name of a Queue or Topic to automatically forward dead lettered messages to.
func (o LookupQueueResultOutput) ForwardDeadLetteredMessagesTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.ForwardDeadLetteredMessagesTo }).(pulumi.StringOutput)
}

// The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
func (o LookupQueueResultOutput) ForwardTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.ForwardTo }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupQueueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers.
func (o LookupQueueResultOutput) LockDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.LockDuration }).(pulumi.StringOutput)
}

// Integer value which controls when a message is automatically dead lettered.
func (o LookupQueueResultOutput) MaxDeliveryCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQueueResult) int { return v.MaxDeliveryCount }).(pulumi.IntOutput)
}

// Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
func (o LookupQueueResultOutput) MaxSizeInMegabytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQueueResult) int { return v.MaxSizeInMegabytes }).(pulumi.IntOutput)
}

func (o LookupQueueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQueueResultOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.NamespaceName }).(pulumi.StringOutput)
}

// Boolean flag which controls whether the Queue requires duplicate detection.
func (o LookupQueueResultOutput) RequiresDuplicateDetection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQueueResult) bool { return v.RequiresDuplicateDetection }).(pulumi.BoolOutput)
}

// Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages.
func (o LookupQueueResultOutput) RequiresSession() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQueueResult) bool { return v.RequiresSession }).(pulumi.BoolOutput)
}

func (o LookupQueueResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`.
func (o LookupQueueResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueResultOutput{})
}
