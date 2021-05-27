// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing ServiceBus Queue Authorisation Rule within a ServiceBus Queue.
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
// 		example, err := servicebus.LookupQueueAuthorizationRule(ctx, &servicebus.LookupQueueAuthorizationRuleArgs{
// 			Name:              "example-tfex_name",
// 			ResourceGroupName: "example-resources",
// 			QueueName:         "example-servicebus_queue",
// 			NamespaceName:     "example-namespace",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupQueueAuthorizationRule(ctx *pulumi.Context, args *LookupQueueAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupQueueAuthorizationRuleResult, error) {
	var rv LookupQueueAuthorizationRuleResult
	err := ctx.Invoke("azure:servicebus/getQueueAuthorizationRule:getQueueAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueueAuthorizationRule.
type LookupQueueAuthorizationRuleArgs struct {
	// The name of this ServiceBus Queue Authorisation Rule.
	Name string `pulumi:"name"`
	// The name of the ServiceBus Namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the ServiceBus Queue.
	QueueName string `pulumi:"queueName"`
	// The name of the Resource Group where the ServiceBus Queue Authorisation Rule exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getQueueAuthorizationRule.
type LookupQueueAuthorizationRuleResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	Listen        bool   `pulumi:"listen"`
	Manage        bool   `pulumi:"manage"`
	Name          string `pulumi:"name"`
	NamespaceName string `pulumi:"namespaceName"`
	// The Primary Connection String for the ServiceBus Queue authorization Rule.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// The Primary Key for the ServiceBus Queue authorization Rule.
	PrimaryKey        string `pulumi:"primaryKey"`
	QueueName         string `pulumi:"queueName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the ServiceBus Queue authorization Rule.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the ServiceBus Queue authorization Rule.
	SecondaryKey string `pulumi:"secondaryKey"`
	Send         bool   `pulumi:"send"`
}

func LookupQueueAuthorizationRuleOutput(ctx *pulumi.Context, args LookupQueueAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupQueueAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueueAuthorizationRuleResult, error) {
			args := v.(LookupQueueAuthorizationRuleArgs)
			r, err := LookupQueueAuthorizationRule(ctx, &args, opts...)
			return *r, err
		}).(LookupQueueAuthorizationRuleResultOutput)
}

// A collection of arguments for invoking getQueueAuthorizationRule.
type LookupQueueAuthorizationRuleOutputArgs struct {
	// The name of this ServiceBus Queue Authorisation Rule.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the ServiceBus Namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the ServiceBus Queue.
	QueueName pulumi.StringInput `pulumi:"queueName"`
	// The name of the Resource Group where the ServiceBus Queue Authorisation Rule exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueueAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueAuthorizationRuleArgs)(nil)).Elem()
}

// A collection of values returned by getQueueAuthorizationRule.
type LookupQueueAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupQueueAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupQueueAuthorizationRuleResultOutput) ToLookupQueueAuthorizationRuleResultOutput() LookupQueueAuthorizationRuleResultOutput {
	return o
}

func (o LookupQueueAuthorizationRuleResultOutput) ToLookupQueueAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupQueueAuthorizationRuleResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupQueueAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQueueAuthorizationRuleResultOutput) Listen() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) bool { return v.Listen }).(pulumi.BoolOutput)
}

func (o LookupQueueAuthorizationRuleResultOutput) Manage() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) bool { return v.Manage }).(pulumi.BoolOutput)
}

func (o LookupQueueAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQueueAuthorizationRuleResultOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.NamespaceName }).(pulumi.StringOutput)
}

// The Primary Connection String for the ServiceBus Queue authorization Rule.
func (o LookupQueueAuthorizationRuleResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

// The Primary Key for the ServiceBus Queue authorization Rule.
func (o LookupQueueAuthorizationRuleResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o LookupQueueAuthorizationRuleResultOutput) QueueName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.QueueName }).(pulumi.StringOutput)
}

func (o LookupQueueAuthorizationRuleResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The Secondary Connection String for the ServiceBus Queue authorization Rule.
func (o LookupQueueAuthorizationRuleResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

// The Secondary Key for the ServiceBus Queue authorization Rule.
func (o LookupQueueAuthorizationRuleResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func (o LookupQueueAuthorizationRuleResultOutput) Send() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) bool { return v.Send }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueAuthorizationRuleResultOutput{})
}
