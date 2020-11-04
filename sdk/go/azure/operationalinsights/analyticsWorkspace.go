// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Log Analytics (formally Operational Insights) Workspace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/operationalinsights"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("East US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = operationalinsights.NewAnalyticsWorkspace(ctx, "exampleAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("PerGB2018"),
// 			RetentionInDays:   pulumi.Int(30),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AnalyticsWorkspace struct {
	pulumi.CustomResourceState

	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited).
	DailyQuotaGb pulumi.Float64PtrOutput `pulumi:"dailyQuotaGb"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	PortalUrl pulumi.StringOutput `pulumi:"portalUrl"`
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey pulumi.StringOutput `pulumi:"primarySharedKey"`
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	RetentionInDays pulumi.IntOutput `pulumi:"retentionInDays"`
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey pulumi.StringOutput `pulumi:"secondarySharedKey"`
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewAnalyticsWorkspace registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsWorkspace(ctx *pulumi.Context,
	name string, args *AnalyticsWorkspaceArgs, opts ...pulumi.ResourceOption) (*AnalyticsWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource AnalyticsWorkspace
	err := ctx.RegisterResource("azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsWorkspace gets an existing AnalyticsWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsWorkspaceState, opts ...pulumi.ResourceOption) (*AnalyticsWorkspace, error) {
	var resource AnalyticsWorkspace
	err := ctx.ReadResource("azure:operationalinsights/analyticsWorkspace:AnalyticsWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalyticsWorkspace resources.
type analyticsWorkspaceState struct {
	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited).
	DailyQuotaGb *float64 `pulumi:"dailyQuotaGb"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	PortalUrl *string `pulumi:"portalUrl"`
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey *string `pulumi:"primarySharedKey"`
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey *string `pulumi:"secondarySharedKey"`
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceId *string `pulumi:"workspaceId"`
}

type AnalyticsWorkspaceState struct {
	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited).
	DailyQuotaGb pulumi.Float64PtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	PortalUrl pulumi.StringPtrInput
	// The Primary shared key for the Log Analytics Workspace.
	PrimarySharedKey pulumi.StringPtrInput
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	RetentionInDays pulumi.IntPtrInput
	// The Secondary shared key for the Log Analytics Workspace.
	SecondarySharedKey pulumi.StringPtrInput
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceId pulumi.StringPtrInput
}

func (AnalyticsWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsWorkspaceState)(nil)).Elem()
}

type analyticsWorkspaceArgs struct {
	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited).
	DailyQuotaGb *float64 `pulumi:"dailyQuotaGb"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AnalyticsWorkspace resource.
type AnalyticsWorkspaceArgs struct {
	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited).
	DailyQuotaGb pulumi.Float64PtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Log Analytics Workspace. Workspace name should include 4-63 letters, digits or '-'. The '-' shouldn't be the first or the last symbol. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	RetentionInDays pulumi.IntPtrInput
	// Specifies the Sku of the Log Analytics Workspace. Possible values are `Free`, `PerNode`, `Premium`, `Standard`, `Standalone`, `Unlimited`, and `PerGB2018` (new Sku as of `2018-04-03`). Defaults to `PerGB2018`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AnalyticsWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsWorkspaceArgs)(nil)).Elem()
}
