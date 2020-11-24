// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loganalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Log Analytics (formally Operational Insights) Saved Search.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/loganalytics"
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
// 		_, err = loganalytics.NewSavedSearch(ctx, "exampleSavedSearch", &loganalytics.SavedSearchArgs{
// 			LogAnalyticsWorkspaceId: pulumi.Any(azurerm_log_analytics_workspace.Test.Id),
// 			Category:                pulumi.String("exampleCategory"),
// 			DisplayName:             pulumi.String("exampleDisplayName"),
// 			Query:                   pulumi.String("exampleQuery"),
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
// Log Analytics Saved Searches can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:loganalytics/savedSearch:SavedSearch search1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/savedSearches/search1
// ```
type SavedSearch struct {
	pulumi.CustomResourceState

	// The category that the Saved Search will be listed under. Changing this forces a new resource to be created.
	Category pulumi.StringOutput `pulumi:"category"`
	// The name that Saved Search will be displayed as. Changing this forces a new resource to be created.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The function alias if the query serves as a function. Changing this forces a new resource to be created.
	FunctionAlias pulumi.StringPtrOutput `pulumi:"functionAlias"`
	// The function parameters if the query serves as a function. Changing this forces a new resource to be created.
	FunctionParameters pulumi.StringArrayOutput `pulumi:"functionParameters"`
	// Specifies the ID of the Log Analytics Workspace that the Saved Search will be associated with. Changing this forces a new resource to be created.
	LogAnalyticsWorkspaceId pulumi.StringOutput `pulumi:"logAnalyticsWorkspaceId"`
	// Specifies the name of the Log Analytics Saved Search. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The query expression for the saved search. Changing this forces a new resource to be created.
	Query pulumi.StringOutput `pulumi:"query"`
	// A mapping of tags which should be assigned to the Logs Analytics Saved Search.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSavedSearch registers a new resource with the given unique name, arguments, and options.
func NewSavedSearch(ctx *pulumi.Context,
	name string, args *SavedSearchArgs, opts ...pulumi.ResourceOption) (*SavedSearch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.LogAnalyticsWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'LogAnalyticsWorkspaceId'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	var resource SavedSearch
	err := ctx.RegisterResource("azure:loganalytics/savedSearch:SavedSearch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSavedSearch gets an existing SavedSearch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSavedSearch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SavedSearchState, opts ...pulumi.ResourceOption) (*SavedSearch, error) {
	var resource SavedSearch
	err := ctx.ReadResource("azure:loganalytics/savedSearch:SavedSearch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SavedSearch resources.
type savedSearchState struct {
	// The category that the Saved Search will be listed under. Changing this forces a new resource to be created.
	Category *string `pulumi:"category"`
	// The name that Saved Search will be displayed as. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// The function alias if the query serves as a function. Changing this forces a new resource to be created.
	FunctionAlias *string `pulumi:"functionAlias"`
	// The function parameters if the query serves as a function. Changing this forces a new resource to be created.
	FunctionParameters []string `pulumi:"functionParameters"`
	// Specifies the ID of the Log Analytics Workspace that the Saved Search will be associated with. Changing this forces a new resource to be created.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// Specifies the name of the Log Analytics Saved Search. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The query expression for the saved search. Changing this forces a new resource to be created.
	Query *string `pulumi:"query"`
	// A mapping of tags which should be assigned to the Logs Analytics Saved Search.
	Tags map[string]string `pulumi:"tags"`
}

type SavedSearchState struct {
	// The category that the Saved Search will be listed under. Changing this forces a new resource to be created.
	Category pulumi.StringPtrInput
	// The name that Saved Search will be displayed as. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// The function alias if the query serves as a function. Changing this forces a new resource to be created.
	FunctionAlias pulumi.StringPtrInput
	// The function parameters if the query serves as a function. Changing this forces a new resource to be created.
	FunctionParameters pulumi.StringArrayInput
	// Specifies the ID of the Log Analytics Workspace that the Saved Search will be associated with. Changing this forces a new resource to be created.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// Specifies the name of the Log Analytics Saved Search. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The query expression for the saved search. Changing this forces a new resource to be created.
	Query pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Logs Analytics Saved Search.
	Tags pulumi.StringMapInput
}

func (SavedSearchState) ElementType() reflect.Type {
	return reflect.TypeOf((*savedSearchState)(nil)).Elem()
}

type savedSearchArgs struct {
	// The category that the Saved Search will be listed under. Changing this forces a new resource to be created.
	Category string `pulumi:"category"`
	// The name that Saved Search will be displayed as. Changing this forces a new resource to be created.
	DisplayName string `pulumi:"displayName"`
	// The function alias if the query serves as a function. Changing this forces a new resource to be created.
	FunctionAlias *string `pulumi:"functionAlias"`
	// The function parameters if the query serves as a function. Changing this forces a new resource to be created.
	FunctionParameters []string `pulumi:"functionParameters"`
	// Specifies the ID of the Log Analytics Workspace that the Saved Search will be associated with. Changing this forces a new resource to be created.
	LogAnalyticsWorkspaceId string `pulumi:"logAnalyticsWorkspaceId"`
	// Specifies the name of the Log Analytics Saved Search. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The query expression for the saved search. Changing this forces a new resource to be created.
	Query string `pulumi:"query"`
	// A mapping of tags which should be assigned to the Logs Analytics Saved Search.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SavedSearch resource.
type SavedSearchArgs struct {
	// The category that the Saved Search will be listed under. Changing this forces a new resource to be created.
	Category pulumi.StringInput
	// The name that Saved Search will be displayed as. Changing this forces a new resource to be created.
	DisplayName pulumi.StringInput
	// The function alias if the query serves as a function. Changing this forces a new resource to be created.
	FunctionAlias pulumi.StringPtrInput
	// The function parameters if the query serves as a function. Changing this forces a new resource to be created.
	FunctionParameters pulumi.StringArrayInput
	// Specifies the ID of the Log Analytics Workspace that the Saved Search will be associated with. Changing this forces a new resource to be created.
	LogAnalyticsWorkspaceId pulumi.StringInput
	// Specifies the name of the Log Analytics Saved Search. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The query expression for the saved search. Changing this forces a new resource to be created.
	Query pulumi.StringInput
	// A mapping of tags which should be assigned to the Logs Analytics Saved Search.
	Tags pulumi.StringMapInput
}

func (SavedSearchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*savedSearchArgs)(nil)).Elem()
}

type SavedSearchInput interface {
	pulumi.Input

	ToSavedSearchOutput() SavedSearchOutput
	ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput
}

func (SavedSearch) ElementType() reflect.Type {
	return reflect.TypeOf((*SavedSearch)(nil)).Elem()
}

func (i SavedSearch) ToSavedSearchOutput() SavedSearchOutput {
	return i.ToSavedSearchOutputWithContext(context.Background())
}

func (i SavedSearch) ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavedSearchOutput)
}

type SavedSearchOutput struct {
	*pulumi.OutputState
}

func (SavedSearchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SavedSearchOutput)(nil)).Elem()
}

func (o SavedSearchOutput) ToSavedSearchOutput() SavedSearchOutput {
	return o
}

func (o SavedSearchOutput) ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SavedSearchOutput{})
}
