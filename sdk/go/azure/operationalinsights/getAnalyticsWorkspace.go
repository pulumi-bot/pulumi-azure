// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Log Analytics (formally Operational Insights) Workspace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/operationalinsights"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := operationalinsights.LookupAnalyticsWorkspace(ctx, "azure:operationalinsights:getAnalyticsWorkspace", &operationalinsights.LookupAnalyticsWorkspaceArgs{
// 			Name:              "acctest-01",
// 			ResourceGroupName: "acctest",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("logAnalyticsWorkspaceId", example.WorkspaceId)
// 		return nil
// 	})
// }
// ```
func LookupAnalyticsWorkspace(ctx *pulumi.Context, args *LookupAnalyticsWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupAnalyticsWorkspaceResult, error) {
	var rv LookupAnalyticsWorkspaceResult
	err := ctx.Invoke("azure:operationalinsights/getAnalyticsWorkspace:getAnalyticsWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAnalyticsWorkspace.
type LookupAnalyticsWorkspaceArgs struct {
	// Specifies the name of the Log Analytics Workspace.
	Name string `pulumi:"name"`
	// The name of the resource group in which the Log Analytics workspace is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getAnalyticsWorkspace.
type LookupAnalyticsWorkspaceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// The Portal URL for the Log Analytics Workspace.
	PortalUrl string `pulumi:"portalUrl"`
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey  string `pulumi:"primarySharedKey"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workspace data retention in days.
	RetentionInDays int `pulumi:"retentionInDays"`
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey string `pulumi:"secondarySharedKey"`
	// The Sku of the Log Analytics Workspace.
	Sku string `pulumi:"sku"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceId string `pulumi:"workspaceId"`
}
