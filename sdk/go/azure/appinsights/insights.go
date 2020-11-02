// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Application Insights component.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
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
// 		exampleInsights, err := appinsights.NewInsights(ctx, "exampleInsights", &appinsights.InsightsArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationType:   pulumi.String("web"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("instrumentationKey", exampleInsights.InstrumentationKey)
// 		ctx.Export("appId", exampleInsights.AppId)
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Application Insights instances can be imported using the `resource id`, e.g.
type Insights struct {
	pulumi.CustomResourceState

	// The App ID associated with this Application Insights component.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
	ApplicationType pulumi.StringOutput `pulumi:"applicationType"`
	// The Connection String for this Application Insights component. (Sensitive)
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// Specifies the Application Insights component daily data volume cap in GB.
	DailyDataCapInGb pulumi.Float64Output `pulumi:"dailyDataCapInGb"`
	// Specifies if a notification email will be send when the daily data volume cap is met.
	DailyDataCapNotificationsDisabled pulumi.BoolOutput `pulumi:"dailyDataCapNotificationsDisabled"`
	// By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
	DisableIpMasking pulumi.BoolPtrOutput `pulumi:"disableIpMasking"`
	// The Instrumentation Key for this Application Insights component.
	InstrumentationKey pulumi.StringOutput `pulumi:"instrumentationKey"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Application Insights component. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to
	// create the Application Insights component.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
	RetentionInDays pulumi.IntPtrOutput `pulumi:"retentionInDays"`
	// Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
	SamplingPercentage pulumi.Float64PtrOutput `pulumi:"samplingPercentage"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewInsights registers a new resource with the given unique name, arguments, and options.
func NewInsights(ctx *pulumi.Context,
	name string, args *InsightsArgs, opts ...pulumi.ResourceOption) (*Insights, error) {
	if args == nil || args.ApplicationType == nil {
		return nil, errors.New("missing required argument 'ApplicationType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &InsightsArgs{}
	}
	var resource Insights
	err := ctx.RegisterResource("azure:appinsights/insights:Insights", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInsights gets an existing Insights resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInsights(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InsightsState, opts ...pulumi.ResourceOption) (*Insights, error) {
	var resource Insights
	err := ctx.ReadResource("azure:appinsights/insights:Insights", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Insights resources.
type insightsState struct {
	// The App ID associated with this Application Insights component.
	AppId *string `pulumi:"appId"`
	// Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
	ApplicationType *string `pulumi:"applicationType"`
	// The Connection String for this Application Insights component. (Sensitive)
	ConnectionString *string `pulumi:"connectionString"`
	// Specifies the Application Insights component daily data volume cap in GB.
	DailyDataCapInGb *float64 `pulumi:"dailyDataCapInGb"`
	// Specifies if a notification email will be send when the daily data volume cap is met.
	DailyDataCapNotificationsDisabled *bool `pulumi:"dailyDataCapNotificationsDisabled"`
	// By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
	DisableIpMasking *bool `pulumi:"disableIpMasking"`
	// The Instrumentation Key for this Application Insights component.
	InstrumentationKey *string `pulumi:"instrumentationKey"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Application Insights component. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the Application Insights component.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
	SamplingPercentage *float64 `pulumi:"samplingPercentage"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type InsightsState struct {
	// The App ID associated with this Application Insights component.
	AppId pulumi.StringPtrInput
	// Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
	ApplicationType pulumi.StringPtrInput
	// The Connection String for this Application Insights component. (Sensitive)
	ConnectionString pulumi.StringPtrInput
	// Specifies the Application Insights component daily data volume cap in GB.
	DailyDataCapInGb pulumi.Float64PtrInput
	// Specifies if a notification email will be send when the daily data volume cap is met.
	DailyDataCapNotificationsDisabled pulumi.BoolPtrInput
	// By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
	DisableIpMasking pulumi.BoolPtrInput
	// The Instrumentation Key for this Application Insights component.
	InstrumentationKey pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Application Insights component. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the Application Insights component.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
	RetentionInDays pulumi.IntPtrInput
	// Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
	SamplingPercentage pulumi.Float64PtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (InsightsState) ElementType() reflect.Type {
	return reflect.TypeOf((*insightsState)(nil)).Elem()
}

type insightsArgs struct {
	// Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
	ApplicationType string `pulumi:"applicationType"`
	// Specifies the Application Insights component daily data volume cap in GB.
	DailyDataCapInGb *float64 `pulumi:"dailyDataCapInGb"`
	// Specifies if a notification email will be send when the daily data volume cap is met.
	DailyDataCapNotificationsDisabled *bool `pulumi:"dailyDataCapNotificationsDisabled"`
	// By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
	DisableIpMasking *bool `pulumi:"disableIpMasking"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Application Insights component. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the Application Insights component.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
	SamplingPercentage *float64 `pulumi:"samplingPercentage"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Insights resource.
type InsightsArgs struct {
	// Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
	ApplicationType pulumi.StringInput
	// Specifies the Application Insights component daily data volume cap in GB.
	DailyDataCapInGb pulumi.Float64PtrInput
	// Specifies if a notification email will be send when the daily data volume cap is met.
	DailyDataCapNotificationsDisabled pulumi.BoolPtrInput
	// By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
	DisableIpMasking pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Application Insights component. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the Application Insights component.
	ResourceGroupName pulumi.StringInput
	// Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
	RetentionInDays pulumi.IntPtrInput
	// Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
	SamplingPercentage pulumi.Float64PtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (InsightsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*insightsArgs)(nil)).Elem()
}
