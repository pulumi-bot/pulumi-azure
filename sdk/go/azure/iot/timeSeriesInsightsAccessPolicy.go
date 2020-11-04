// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure IoT Time Series Insights Access Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/iot"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTimeSeriesInsightsStandardEnvironment, err := iot.NewTimeSeriesInsightsStandardEnvironment(ctx, "exampleTimeSeriesInsightsStandardEnvironment", &iot.TimeSeriesInsightsStandardEnvironmentArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("S1_1"),
// 			DataRetentionTime: pulumi.String("P30D"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewTimeSeriesInsightsAccessPolicy(ctx, "exampleTimeSeriesInsightsAccessPolicy", &iot.TimeSeriesInsightsAccessPolicyArgs{
// 			TimeSeriesInsightsEnvironmentId: exampleTimeSeriesInsightsStandardEnvironment.Name,
// 			PrincipalObjectId:               pulumi.String("aGUID"),
// 			Roles: pulumi.StringArray{
// 				pulumi.String("Reader"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type TimeSeriesInsightsAccessPolicy struct {
	pulumi.CustomResourceState

	// The description of the Azure IoT Time Series Insights Access Policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the name of the Azure IoT Time Series Insights Access Policy. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the principal in Azure Active Directory.
	PrincipalObjectId pulumi.StringOutput `pulumi:"principalObjectId"`
	// A list of roles to apply to the Access Policy. Valid values include `Contributor` and `Reader`.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentId pulumi.StringOutput `pulumi:"timeSeriesInsightsEnvironmentId"`
}

// NewTimeSeriesInsightsAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewTimeSeriesInsightsAccessPolicy(ctx *pulumi.Context,
	name string, args *TimeSeriesInsightsAccessPolicyArgs, opts ...pulumi.ResourceOption) (*TimeSeriesInsightsAccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.PrincipalObjectId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalObjectId'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.TimeSeriesInsightsEnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'TimeSeriesInsightsEnvironmentId'")
	}
	var resource TimeSeriesInsightsAccessPolicy
	err := ctx.RegisterResource("azure:iot/timeSeriesInsightsAccessPolicy:TimeSeriesInsightsAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTimeSeriesInsightsAccessPolicy gets an existing TimeSeriesInsightsAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTimeSeriesInsightsAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TimeSeriesInsightsAccessPolicyState, opts ...pulumi.ResourceOption) (*TimeSeriesInsightsAccessPolicy, error) {
	var resource TimeSeriesInsightsAccessPolicy
	err := ctx.ReadResource("azure:iot/timeSeriesInsightsAccessPolicy:TimeSeriesInsightsAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TimeSeriesInsightsAccessPolicy resources.
type timeSeriesInsightsAccessPolicyState struct {
	// The description of the Azure IoT Time Series Insights Access Policy.
	Description *string `pulumi:"description"`
	// Specifies the name of the Azure IoT Time Series Insights Access Policy. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The id of the principal in Azure Active Directory.
	PrincipalObjectId *string `pulumi:"principalObjectId"`
	// A list of roles to apply to the Access Policy. Valid values include `Contributor` and `Reader`.
	Roles []string `pulumi:"roles"`
	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentId *string `pulumi:"timeSeriesInsightsEnvironmentId"`
}

type TimeSeriesInsightsAccessPolicyState struct {
	// The description of the Azure IoT Time Series Insights Access Policy.
	Description pulumi.StringPtrInput
	// Specifies the name of the Azure IoT Time Series Insights Access Policy. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The id of the principal in Azure Active Directory.
	PrincipalObjectId pulumi.StringPtrInput
	// A list of roles to apply to the Access Policy. Valid values include `Contributor` and `Reader`.
	Roles pulumi.StringArrayInput
	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentId pulumi.StringPtrInput
}

func (TimeSeriesInsightsAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesInsightsAccessPolicyState)(nil)).Elem()
}

type timeSeriesInsightsAccessPolicyArgs struct {
	// The description of the Azure IoT Time Series Insights Access Policy.
	Description *string `pulumi:"description"`
	// Specifies the name of the Azure IoT Time Series Insights Access Policy. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The id of the principal in Azure Active Directory.
	PrincipalObjectId string `pulumi:"principalObjectId"`
	// A list of roles to apply to the Access Policy. Valid values include `Contributor` and `Reader`.
	Roles []string `pulumi:"roles"`
	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentId string `pulumi:"timeSeriesInsightsEnvironmentId"`
}

// The set of arguments for constructing a TimeSeriesInsightsAccessPolicy resource.
type TimeSeriesInsightsAccessPolicyArgs struct {
	// The description of the Azure IoT Time Series Insights Access Policy.
	Description pulumi.StringPtrInput
	// Specifies the name of the Azure IoT Time Series Insights Access Policy. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The id of the principal in Azure Active Directory.
	PrincipalObjectId pulumi.StringInput
	// A list of roles to apply to the Access Policy. Valid values include `Contributor` and `Reader`.
	Roles pulumi.StringArrayInput
	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentId pulumi.StringInput
}

func (TimeSeriesInsightsAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesInsightsAccessPolicyArgs)(nil)).Elem()
}
