// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sentinel

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Sentinel MS Security Incident Alert Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/operationalinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/sentinel"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAnalyticsWorkspace, err := operationalinsights.NewAnalyticsWorkspace(ctx, "exampleAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("pergb2018"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sentinel.NewAlertRuleMsSecurityIncident(ctx, "exampleAlertRuleMsSecurityIncident", &sentinel.AlertRuleMsSecurityIncidentArgs{
// 			LogAnalyticsWorkspaceId: exampleAnalyticsWorkspace.ID(),
// 			ProductFilter:           pulumi.String("Microsoft Cloud App Security"),
// 			DisplayName:             pulumi.String("example rule"),
// 			SeverityFilters: pulumi.StringArray{
// 				pulumi.String("High"),
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
// Sentinel MS Security Incident Alert Rules can be imported using the `resource id`, e.g.
type AlertRuleMsSecurityIncident struct {
	pulumi.CustomResourceState

	// The description of this Sentinel MS Security Incident Alert Rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilters pulumi.StringArrayOutput `pulumi:"displayNameFilters"`
	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceId pulumi.StringOutput `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT` and `Microsoft Cloud App Security`.
	ProductFilter pulumi.StringOutput `pulumi:"productFilter"`
	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
	SeverityFilters pulumi.StringArrayOutput `pulumi:"severityFilters"`
	// Deprecated: this property has been renamed to display_name_filter to better match the SDK & API
	TextWhitelists pulumi.StringArrayOutput `pulumi:"textWhitelists"`
}

// NewAlertRuleMsSecurityIncident registers a new resource with the given unique name, arguments, and options.
func NewAlertRuleMsSecurityIncident(ctx *pulumi.Context,
	name string, args *AlertRuleMsSecurityIncidentArgs, opts ...pulumi.ResourceOption) (*AlertRuleMsSecurityIncident, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.LogAnalyticsWorkspaceId == nil {
		return nil, errors.New("missing required argument 'LogAnalyticsWorkspaceId'")
	}
	if args == nil || args.ProductFilter == nil {
		return nil, errors.New("missing required argument 'ProductFilter'")
	}
	if args == nil || args.SeverityFilters == nil {
		return nil, errors.New("missing required argument 'SeverityFilters'")
	}
	if args == nil {
		args = &AlertRuleMsSecurityIncidentArgs{}
	}
	var resource AlertRuleMsSecurityIncident
	err := ctx.RegisterResource("azure:sentinel/alertRuleMsSecurityIncident:AlertRuleMsSecurityIncident", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertRuleMsSecurityIncident gets an existing AlertRuleMsSecurityIncident resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertRuleMsSecurityIncident(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertRuleMsSecurityIncidentState, opts ...pulumi.ResourceOption) (*AlertRuleMsSecurityIncident, error) {
	var resource AlertRuleMsSecurityIncident
	err := ctx.ReadResource("azure:sentinel/alertRuleMsSecurityIncident:AlertRuleMsSecurityIncident", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertRuleMsSecurityIncident resources.
type alertRuleMsSecurityIncidentState struct {
	// The description of this Sentinel MS Security Incident Alert Rule.
	Description *string `pulumi:"description"`
	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName *string `pulumi:"displayName"`
	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilters []string `pulumi:"displayNameFilters"`
	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	Name *string `pulumi:"name"`
	// The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT` and `Microsoft Cloud App Security`.
	ProductFilter *string `pulumi:"productFilter"`
	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
	SeverityFilters []string `pulumi:"severityFilters"`
	// Deprecated: this property has been renamed to display_name_filter to better match the SDK & API
	TextWhitelists []string `pulumi:"textWhitelists"`
}

type AlertRuleMsSecurityIncidentState struct {
	// The description of this Sentinel MS Security Incident Alert Rule.
	Description pulumi.StringPtrInput
	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName pulumi.StringPtrInput
	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilters pulumi.StringArrayInput
	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	Name pulumi.StringPtrInput
	// The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT` and `Microsoft Cloud App Security`.
	ProductFilter pulumi.StringPtrInput
	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
	SeverityFilters pulumi.StringArrayInput
	// Deprecated: this property has been renamed to display_name_filter to better match the SDK & API
	TextWhitelists pulumi.StringArrayInput
}

func (AlertRuleMsSecurityIncidentState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleMsSecurityIncidentState)(nil)).Elem()
}

type alertRuleMsSecurityIncidentArgs struct {
	// The description of this Sentinel MS Security Incident Alert Rule.
	Description *string `pulumi:"description"`
	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName string `pulumi:"displayName"`
	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilters []string `pulumi:"displayNameFilters"`
	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceId string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	Name *string `pulumi:"name"`
	// The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT` and `Microsoft Cloud App Security`.
	ProductFilter string `pulumi:"productFilter"`
	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
	SeverityFilters []string `pulumi:"severityFilters"`
	// Deprecated: this property has been renamed to display_name_filter to better match the SDK & API
	TextWhitelists []string `pulumi:"textWhitelists"`
}

// The set of arguments for constructing a AlertRuleMsSecurityIncident resource.
type AlertRuleMsSecurityIncidentArgs struct {
	// The description of this Sentinel MS Security Incident Alert Rule.
	Description pulumi.StringPtrInput
	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName pulumi.StringInput
	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilters pulumi.StringArrayInput
	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceId pulumi.StringInput
	// The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	Name pulumi.StringPtrInput
	// The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT` and `Microsoft Cloud App Security`.
	ProductFilter pulumi.StringInput
	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
	SeverityFilters pulumi.StringArrayInput
	// Deprecated: this property has been renamed to display_name_filter to better match the SDK & API
	TextWhitelists pulumi.StringArrayInput
}

func (AlertRuleMsSecurityIncidentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleMsSecurityIncidentArgs)(nil)).Elem()
}
