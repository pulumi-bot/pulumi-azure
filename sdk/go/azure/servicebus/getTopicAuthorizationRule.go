// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about a ServiceBus Topic Authorization Rule within a ServiceBus Topic.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicebus.LookupTopicAuthorizationRule(ctx, &servicebus.LookupTopicAuthorizationRuleArgs{
// 			Name:              "example-tfex_name",
// 			NamespaceName:     "example-namespace",
// 			ResourceGroupName: "example-resources",
// 			TopicName:         "example-servicebus_topic",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("servicebusAuthorizationRuleId", data.Azurem_servicebus_topic_authorization_rule.Example.Id)
// 		return nil
// 	})
// }
// ```
func LookupTopicAuthorizationRule(ctx *pulumi.Context, args *LookupTopicAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupTopicAuthorizationRuleResult, error) {
	var rv LookupTopicAuthorizationRuleResult
	err := ctx.Invoke("azure:servicebus/getTopicAuthorizationRule:getTopicAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTopicAuthorizationRule.
type LookupTopicAuthorizationRuleArgs struct {
	// The name of the ServiceBus Topic Authorization Rule resource.
	Name string `pulumi:"name"`
	// The name of the ServiceBus Namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the ServiceBus Namespace exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the ServiceBus Topic.
	TopicName string `pulumi:"topicName"`
}

// A collection of values returned by getTopicAuthorizationRule.
type LookupTopicAuthorizationRuleResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	Listen        bool   `pulumi:"listen"`
	Manage        bool   `pulumi:"manage"`
	Name          string `pulumi:"name"`
	NamespaceName string `pulumi:"namespaceName"`
	// The Primary Connection String for the ServiceBus Topic authorization Rule.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// The Primary Key for the ServiceBus Topic authorization Rule.
	PrimaryKey        string `pulumi:"primaryKey"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the ServiceBus Topic authorization Rule.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the ServiceBus Topic authorization Rule.
	SecondaryKey string `pulumi:"secondaryKey"`
	Send         bool   `pulumi:"send"`
	TopicName    string `pulumi:"topicName"`
}
