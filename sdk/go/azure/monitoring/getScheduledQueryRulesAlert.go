// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access the properties of an AlertingAction scheduled query rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/monitoring"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := monitoring.LookupScheduledQueryRulesAlert(ctx, &monitoring.LookupScheduledQueryRulesAlertArgs{
// 			Name:              "tfex-queryrule",
// 			ResourceGroupName: "example-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("queryRuleId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupScheduledQueryRulesAlert(ctx *pulumi.Context, args *LookupScheduledQueryRulesAlertArgs, opts ...pulumi.InvokeOption) (*LookupScheduledQueryRulesAlertResult, error) {
	var rv LookupScheduledQueryRulesAlertResult
	err := ctx.Invoke("azure:monitoring/getScheduledQueryRulesAlert:getScheduledQueryRulesAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScheduledQueryRulesAlert.
type LookupScheduledQueryRulesAlertArgs struct {
	// Specifies the name of the scheduled query rule.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group where the scheduled query rule is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getScheduledQueryRulesAlert.
type LookupScheduledQueryRulesAlertResult struct {
	// An `action` block as defined below.
	Actions []GetScheduledQueryRulesAlertAction `pulumi:"actions"`
	// List of Resource IDs referred into query.
	AuthorizedResourceIds []string `pulumi:"authorizedResourceIds"`
	// The resource URI over which log search query is to be run.
	DataSourceId string `pulumi:"dataSourceId"`
	// The description of the scheduled query rule.
	Description string `pulumi:"description"`
	// Whether this scheduled query rule is enabled.
	Enabled bool `pulumi:"enabled"`
	// Frequency at which rule condition should be evaluated.
	Frequency int `pulumi:"frequency"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// Log search query.
	Query             string `pulumi:"query"`
	QueryType         string `pulumi:"queryType"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Severity of the alert.
	Severity int               `pulumi:"severity"`
	Tags     map[string]string `pulumi:"tags"`
	// Time for which alerts should be throttled or suppressed.
	Throttling int `pulumi:"throttling"`
	// Time window for which data needs to be fetched for query.
	TimeWindow int `pulumi:"timeWindow"`
	// A `trigger` block as defined below.
	Triggers []GetScheduledQueryRulesAlertTrigger `pulumi:"triggers"`
}
