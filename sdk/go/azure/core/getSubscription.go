// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Subscription.
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
// 		current, err := core.LookupSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("currentSubscriptionDisplayName", current.DisplayName)
// 		return nil
// 	})
// }
// ```
func GetSubscription(ctx *pulumi.Context, args *GetSubscriptionArgs, opts ...pulumi.InvokeOption) (*GetSubscriptionResult, error) {
	var rv GetSubscriptionResult
	err := ctx.Invoke("azure:core/getSubscription:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubscription.
type GetSubscriptionArgs struct {
	// Specifies the ID of the subscription. If this argument is omitted, the subscription ID of the current Azure Resource Manager provider is used.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// A collection of values returned by getSubscription.
type GetSubscriptionResult struct {
	// The subscription display name.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The subscription location placement ID.
	LocationPlacementId string `pulumi:"locationPlacementId"`
	// The subscription quota ID.
	QuotaId string `pulumi:"quotaId"`
	// The subscription spending limit.
	SpendingLimit string `pulumi:"spendingLimit"`
	// The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
	State string `pulumi:"state"`
	// The subscription GUID.
	SubscriptionId string `pulumi:"subscriptionId"`
	// The subscription tenant ID.
	TenantId string `pulumi:"tenantId"`
}
