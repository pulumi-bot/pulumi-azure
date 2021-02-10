// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an iot security solution.
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
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleIoTHub, err := iot.NewIoTHub(ctx, "exampleIoTHub", &iot.IoTHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku: &iot.IoTHubSkuArgs{
// 				Name:     pulumi.String("S1"),
// 				Capacity: pulumi.Int(1),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewSecuritySolution(ctx, "exampleSecuritySolution", &iot.SecuritySolutionArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			DisplayName:       pulumi.String("Iot Security Solution"),
// 			IothubIds: pulumi.StringArray{
// 				exampleIoTHub.ID(),
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
// Iot Security Solution can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/securitySolution:SecuritySolution example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Security/IoTSecuritySolutions/solution1
// ```
type SecuritySolution struct {
	pulumi.CustomResourceState

	// Specifies the Display Name for this Iot Security Solution.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Is the Iot Security Solution enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
	EventsToExports pulumi.StringArrayOutput `pulumi:"eventsToExports"`
	// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
	IothubIds pulumi.StringArrayOutput `pulumi:"iothubIds"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the Log Analytics Workspace ID to which the security data will be sent.
	LogAnalyticsWorkspaceId pulumi.StringPtrOutput `pulumi:"logAnalyticsWorkspaceId"`
	// Should ip addressed be unmasked in the log? Defaults to `false`.
	LogUnmaskedIpsEnabled pulumi.BoolPtrOutput `pulumi:"logUnmaskedIpsEnabled"`
	// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// An Azure Resource Graph query used to set the resources monitored.
	QueryForResources pulumi.StringOutput `pulumi:"queryForResources"`
	// A list of subscription Ids on which the user defined resources query should be executed.
	QuerySubscriptionIds pulumi.StringArrayOutput `pulumi:"querySubscriptionIds"`
	// A `recommendationsEnabled` block of options to enable or disable as defined below.
	RecommendationsEnabled SecuritySolutionRecommendationsEnabledOutput `pulumi:"recommendationsEnabled"`
	// Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSecuritySolution registers a new resource with the given unique name, arguments, and options.
func NewSecuritySolution(ctx *pulumi.Context,
	name string, args *SecuritySolutionArgs, opts ...pulumi.ResourceOption) (*SecuritySolution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.IothubIds == nil {
		return nil, errors.New("invalid value for required argument 'IothubIds'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource SecuritySolution
	err := ctx.RegisterResource("azure:iot/securitySolution:SecuritySolution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecuritySolution gets an existing SecuritySolution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecuritySolution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecuritySolutionState, opts ...pulumi.ResourceOption) (*SecuritySolution, error) {
	var resource SecuritySolution
	err := ctx.ReadResource("azure:iot/securitySolution:SecuritySolution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecuritySolution resources.
type securitySolutionState struct {
	// Specifies the Display Name for this Iot Security Solution.
	DisplayName *string `pulumi:"displayName"`
	// Is the Iot Security Solution enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
	EventsToExports []string `pulumi:"eventsToExports"`
	// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
	IothubIds []string `pulumi:"iothubIds"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the Log Analytics Workspace ID to which the security data will be sent.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// Should ip addressed be unmasked in the log? Defaults to `false`.
	LogUnmaskedIpsEnabled *bool `pulumi:"logUnmaskedIpsEnabled"`
	// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// An Azure Resource Graph query used to set the resources monitored.
	QueryForResources *string `pulumi:"queryForResources"`
	// A list of subscription Ids on which the user defined resources query should be executed.
	QuerySubscriptionIds []string `pulumi:"querySubscriptionIds"`
	// A `recommendationsEnabled` block of options to enable or disable as defined below.
	RecommendationsEnabled *SecuritySolutionRecommendationsEnabled `pulumi:"recommendationsEnabled"`
	// Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SecuritySolutionState struct {
	// Specifies the Display Name for this Iot Security Solution.
	DisplayName pulumi.StringPtrInput
	// Is the Iot Security Solution enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
	EventsToExports pulumi.StringArrayInput
	// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
	IothubIds pulumi.StringArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the Log Analytics Workspace ID to which the security data will be sent.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// Should ip addressed be unmasked in the log? Defaults to `false`.
	LogUnmaskedIpsEnabled pulumi.BoolPtrInput
	// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// An Azure Resource Graph query used to set the resources monitored.
	QueryForResources pulumi.StringPtrInput
	// A list of subscription Ids on which the user defined resources query should be executed.
	QuerySubscriptionIds pulumi.StringArrayInput
	// A `recommendationsEnabled` block of options to enable or disable as defined below.
	RecommendationsEnabled SecuritySolutionRecommendationsEnabledPtrInput
	// Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SecuritySolutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*securitySolutionState)(nil)).Elem()
}

type securitySolutionArgs struct {
	// Specifies the Display Name for this Iot Security Solution.
	DisplayName string `pulumi:"displayName"`
	// Is the Iot Security Solution enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
	EventsToExports []string `pulumi:"eventsToExports"`
	// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
	IothubIds []string `pulumi:"iothubIds"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the Log Analytics Workspace ID to which the security data will be sent.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// Should ip addressed be unmasked in the log? Defaults to `false`.
	LogUnmaskedIpsEnabled *bool `pulumi:"logUnmaskedIpsEnabled"`
	// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// An Azure Resource Graph query used to set the resources monitored.
	QueryForResources *string `pulumi:"queryForResources"`
	// A list of subscription Ids on which the user defined resources query should be executed.
	QuerySubscriptionIds []string `pulumi:"querySubscriptionIds"`
	// A `recommendationsEnabled` block of options to enable or disable as defined below.
	RecommendationsEnabled *SecuritySolutionRecommendationsEnabled `pulumi:"recommendationsEnabled"`
	// Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SecuritySolution resource.
type SecuritySolutionArgs struct {
	// Specifies the Display Name for this Iot Security Solution.
	DisplayName pulumi.StringInput
	// Is the Iot Security Solution enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
	EventsToExports pulumi.StringArrayInput
	// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
	IothubIds pulumi.StringArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the Log Analytics Workspace ID to which the security data will be sent.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// Should ip addressed be unmasked in the log? Defaults to `false`.
	LogUnmaskedIpsEnabled pulumi.BoolPtrInput
	// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// An Azure Resource Graph query used to set the resources monitored.
	QueryForResources pulumi.StringPtrInput
	// A list of subscription Ids on which the user defined resources query should be executed.
	QuerySubscriptionIds pulumi.StringArrayInput
	// A `recommendationsEnabled` block of options to enable or disable as defined below.
	RecommendationsEnabled SecuritySolutionRecommendationsEnabledPtrInput
	// Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SecuritySolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securitySolutionArgs)(nil)).Elem()
}

type SecuritySolutionInput interface {
	pulumi.Input

	ToSecuritySolutionOutput() SecuritySolutionOutput
	ToSecuritySolutionOutputWithContext(ctx context.Context) SecuritySolutionOutput
}

func (*SecuritySolution) ElementType() reflect.Type {
	return reflect.TypeOf((*SecuritySolution)(nil))
}

func (i *SecuritySolution) ToSecuritySolutionOutput() SecuritySolutionOutput {
	return i.ToSecuritySolutionOutputWithContext(context.Background())
}

func (i *SecuritySolution) ToSecuritySolutionOutputWithContext(ctx context.Context) SecuritySolutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecuritySolutionOutput)
}

type SecuritySolutionOutput struct {
	*pulumi.OutputState
}

func (SecuritySolutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecuritySolution)(nil))
}

func (o SecuritySolutionOutput) ToSecuritySolutionOutput() SecuritySolutionOutput {
	return o
}

func (o SecuritySolutionOutput) ToSecuritySolutionOutputWithContext(ctx context.Context) SecuritySolutionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecuritySolutionOutput{})
}
