// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure IoT Time Series Insights Access Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/iot"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
//
// ## Import
//
// Azure IoT Time Series Insights Access Policy can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/timeSeriesInsightsAccessPolicy:TimeSeriesInsightsAccessPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/environment1/accessPolicies/example
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

type TimeSeriesInsightsAccessPolicyInput interface {
	pulumi.Input

	ToTimeSeriesInsightsAccessPolicyOutput() TimeSeriesInsightsAccessPolicyOutput
	ToTimeSeriesInsightsAccessPolicyOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyOutput
}

func (*TimeSeriesInsightsAccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsAccessPolicy)(nil))
}

func (i *TimeSeriesInsightsAccessPolicy) ToTimeSeriesInsightsAccessPolicyOutput() TimeSeriesInsightsAccessPolicyOutput {
	return i.ToTimeSeriesInsightsAccessPolicyOutputWithContext(context.Background())
}

func (i *TimeSeriesInsightsAccessPolicy) ToTimeSeriesInsightsAccessPolicyOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsAccessPolicyOutput)
}

func (i *TimeSeriesInsightsAccessPolicy) ToTimeSeriesInsightsAccessPolicyPtrOutput() TimeSeriesInsightsAccessPolicyPtrOutput {
	return i.ToTimeSeriesInsightsAccessPolicyPtrOutputWithContext(context.Background())
}

func (i *TimeSeriesInsightsAccessPolicy) ToTimeSeriesInsightsAccessPolicyPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsAccessPolicyPtrOutput)
}

type TimeSeriesInsightsAccessPolicyPtrInput interface {
	pulumi.Input

	ToTimeSeriesInsightsAccessPolicyPtrOutput() TimeSeriesInsightsAccessPolicyPtrOutput
	ToTimeSeriesInsightsAccessPolicyPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyPtrOutput
}

type timeSeriesInsightsAccessPolicyPtrType TimeSeriesInsightsAccessPolicyArgs

func (*timeSeriesInsightsAccessPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesInsightsAccessPolicy)(nil))
}

func (i *timeSeriesInsightsAccessPolicyPtrType) ToTimeSeriesInsightsAccessPolicyPtrOutput() TimeSeriesInsightsAccessPolicyPtrOutput {
	return i.ToTimeSeriesInsightsAccessPolicyPtrOutputWithContext(context.Background())
}

func (i *timeSeriesInsightsAccessPolicyPtrType) ToTimeSeriesInsightsAccessPolicyPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsAccessPolicyPtrOutput)
}

// TimeSeriesInsightsAccessPolicyArrayInput is an input type that accepts TimeSeriesInsightsAccessPolicyArray and TimeSeriesInsightsAccessPolicyArrayOutput values.
// You can construct a concrete instance of `TimeSeriesInsightsAccessPolicyArrayInput` via:
//
//          TimeSeriesInsightsAccessPolicyArray{ TimeSeriesInsightsAccessPolicyArgs{...} }
type TimeSeriesInsightsAccessPolicyArrayInput interface {
	pulumi.Input

	ToTimeSeriesInsightsAccessPolicyArrayOutput() TimeSeriesInsightsAccessPolicyArrayOutput
	ToTimeSeriesInsightsAccessPolicyArrayOutputWithContext(context.Context) TimeSeriesInsightsAccessPolicyArrayOutput
}

type TimeSeriesInsightsAccessPolicyArray []TimeSeriesInsightsAccessPolicyInput

func (TimeSeriesInsightsAccessPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TimeSeriesInsightsAccessPolicy)(nil))
}

func (i TimeSeriesInsightsAccessPolicyArray) ToTimeSeriesInsightsAccessPolicyArrayOutput() TimeSeriesInsightsAccessPolicyArrayOutput {
	return i.ToTimeSeriesInsightsAccessPolicyArrayOutputWithContext(context.Background())
}

func (i TimeSeriesInsightsAccessPolicyArray) ToTimeSeriesInsightsAccessPolicyArrayOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsAccessPolicyArrayOutput)
}

// TimeSeriesInsightsAccessPolicyMapInput is an input type that accepts TimeSeriesInsightsAccessPolicyMap and TimeSeriesInsightsAccessPolicyMapOutput values.
// You can construct a concrete instance of `TimeSeriesInsightsAccessPolicyMapInput` via:
//
//          TimeSeriesInsightsAccessPolicyMap{ "key": TimeSeriesInsightsAccessPolicyArgs{...} }
type TimeSeriesInsightsAccessPolicyMapInput interface {
	pulumi.Input

	ToTimeSeriesInsightsAccessPolicyMapOutput() TimeSeriesInsightsAccessPolicyMapOutput
	ToTimeSeriesInsightsAccessPolicyMapOutputWithContext(context.Context) TimeSeriesInsightsAccessPolicyMapOutput
}

type TimeSeriesInsightsAccessPolicyMap map[string]TimeSeriesInsightsAccessPolicyInput

func (TimeSeriesInsightsAccessPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TimeSeriesInsightsAccessPolicy)(nil))
}

func (i TimeSeriesInsightsAccessPolicyMap) ToTimeSeriesInsightsAccessPolicyMapOutput() TimeSeriesInsightsAccessPolicyMapOutput {
	return i.ToTimeSeriesInsightsAccessPolicyMapOutputWithContext(context.Background())
}

func (i TimeSeriesInsightsAccessPolicyMap) ToTimeSeriesInsightsAccessPolicyMapOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsAccessPolicyMapOutput)
}

type TimeSeriesInsightsAccessPolicyOutput struct {
	*pulumi.OutputState
}

func (TimeSeriesInsightsAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsAccessPolicy)(nil))
}

func (o TimeSeriesInsightsAccessPolicyOutput) ToTimeSeriesInsightsAccessPolicyOutput() TimeSeriesInsightsAccessPolicyOutput {
	return o
}

func (o TimeSeriesInsightsAccessPolicyOutput) ToTimeSeriesInsightsAccessPolicyOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyOutput {
	return o
}

func (o TimeSeriesInsightsAccessPolicyOutput) ToTimeSeriesInsightsAccessPolicyPtrOutput() TimeSeriesInsightsAccessPolicyPtrOutput {
	return o.ToTimeSeriesInsightsAccessPolicyPtrOutputWithContext(context.Background())
}

func (o TimeSeriesInsightsAccessPolicyOutput) ToTimeSeriesInsightsAccessPolicyPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyPtrOutput {
	return o.ApplyT(func(v TimeSeriesInsightsAccessPolicy) *TimeSeriesInsightsAccessPolicy {
		return &v
	}).(TimeSeriesInsightsAccessPolicyPtrOutput)
}

type TimeSeriesInsightsAccessPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (TimeSeriesInsightsAccessPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesInsightsAccessPolicy)(nil))
}

func (o TimeSeriesInsightsAccessPolicyPtrOutput) ToTimeSeriesInsightsAccessPolicyPtrOutput() TimeSeriesInsightsAccessPolicyPtrOutput {
	return o
}

func (o TimeSeriesInsightsAccessPolicyPtrOutput) ToTimeSeriesInsightsAccessPolicyPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyPtrOutput {
	return o
}

type TimeSeriesInsightsAccessPolicyArrayOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsAccessPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSeriesInsightsAccessPolicy)(nil))
}

func (o TimeSeriesInsightsAccessPolicyArrayOutput) ToTimeSeriesInsightsAccessPolicyArrayOutput() TimeSeriesInsightsAccessPolicyArrayOutput {
	return o
}

func (o TimeSeriesInsightsAccessPolicyArrayOutput) ToTimeSeriesInsightsAccessPolicyArrayOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyArrayOutput {
	return o
}

func (o TimeSeriesInsightsAccessPolicyArrayOutput) Index(i pulumi.IntInput) TimeSeriesInsightsAccessPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeSeriesInsightsAccessPolicy {
		return vs[0].([]TimeSeriesInsightsAccessPolicy)[vs[1].(int)]
	}).(TimeSeriesInsightsAccessPolicyOutput)
}

type TimeSeriesInsightsAccessPolicyMapOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsAccessPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TimeSeriesInsightsAccessPolicy)(nil))
}

func (o TimeSeriesInsightsAccessPolicyMapOutput) ToTimeSeriesInsightsAccessPolicyMapOutput() TimeSeriesInsightsAccessPolicyMapOutput {
	return o
}

func (o TimeSeriesInsightsAccessPolicyMapOutput) ToTimeSeriesInsightsAccessPolicyMapOutputWithContext(ctx context.Context) TimeSeriesInsightsAccessPolicyMapOutput {
	return o
}

func (o TimeSeriesInsightsAccessPolicyMapOutput) MapIndex(k pulumi.StringInput) TimeSeriesInsightsAccessPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TimeSeriesInsightsAccessPolicy {
		return vs[0].(map[string]TimeSeriesInsightsAccessPolicy)[vs[1].(string)]
	}).(TimeSeriesInsightsAccessPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(TimeSeriesInsightsAccessPolicyOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsAccessPolicyPtrOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsAccessPolicyArrayOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsAccessPolicyMapOutput{})
}
