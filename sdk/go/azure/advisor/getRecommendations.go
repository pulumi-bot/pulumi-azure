// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package advisor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Advisor Recommendations.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/advisor"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := advisor.GetRecommendations(ctx, &advisor.GetRecommendationsArgs{
// 			FilterByCategories: []string{
// 				"security",
// 				"cost",
// 			},
// 			FilterByResourceGroups: []string{
// 				"example-resgroups",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("recommendations", example.Recommendations)
// 		return nil
// 	})
// }
// ```
func GetRecommendations(ctx *pulumi.Context, args *GetRecommendationsArgs, opts ...pulumi.InvokeOption) (*GetRecommendationsResult, error) {
	var rv GetRecommendationsResult
	err := ctx.Invoke("azure:advisor/getRecommendations:getRecommendations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecommendations.
type GetRecommendationsArgs struct {
	// Specifies a list of categories in which the Advisor Recommendations will be listed. Possible values are `HighAvailability`, `Security`, `Performance`, `Cost` and `OperationalExcellence`.
	FilterByCategories []string `pulumi:"filterByCategories"`
	// Specifies a list of resource groups about which the Advisor Recommendations will be listed.
	FilterByResourceGroups []string `pulumi:"filterByResourceGroups"`
}

// A collection of values returned by getRecommendations.
type GetRecommendationsResult struct {
	FilterByCategories     []string `pulumi:"filterByCategories"`
	FilterByResourceGroups []string `pulumi:"filterByResourceGroups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// One or more `recommendations` blocks as defined below.
	Recommendations []GetRecommendationsRecommendation `pulumi:"recommendations"`
}

func GetRecommendationsApply(ctx *pulumi.Context, args GetRecommendationsApplyInput, opts ...pulumi.InvokeOption) GetRecommendationsResultOutput {
	return args.ToGetRecommendationsApplyOutput().ApplyT(func(v GetRecommendationsArgs) (GetRecommendationsResult, error) {
		r, err := GetRecommendations(ctx, &v, opts...)
		return *r, err

	}).(GetRecommendationsResultOutput)
}

// GetRecommendationsApplyInput is an input type that accepts GetRecommendationsApplyArgs and GetRecommendationsApplyOutput values.
// You can construct a concrete instance of `GetRecommendationsApplyInput` via:
//
//          GetRecommendationsApplyArgs{...}
type GetRecommendationsApplyInput interface {
	pulumi.Input

	ToGetRecommendationsApplyOutput() GetRecommendationsApplyOutput
	ToGetRecommendationsApplyOutputWithContext(context.Context) GetRecommendationsApplyOutput
}

// A collection of arguments for invoking getRecommendations.
type GetRecommendationsApplyArgs struct {
	// Specifies a list of categories in which the Advisor Recommendations will be listed. Possible values are `HighAvailability`, `Security`, `Performance`, `Cost` and `OperationalExcellence`.
	FilterByCategories pulumi.StringArrayInput `pulumi:"filterByCategories"`
	// Specifies a list of resource groups about which the Advisor Recommendations will be listed.
	FilterByResourceGroups pulumi.StringArrayInput `pulumi:"filterByResourceGroups"`
}

func (GetRecommendationsApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecommendationsArgs)(nil)).Elem()
}

func (i GetRecommendationsApplyArgs) ToGetRecommendationsApplyOutput() GetRecommendationsApplyOutput {
	return i.ToGetRecommendationsApplyOutputWithContext(context.Background())
}

func (i GetRecommendationsApplyArgs) ToGetRecommendationsApplyOutputWithContext(ctx context.Context) GetRecommendationsApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRecommendationsApplyOutput)
}

// A collection of arguments for invoking getRecommendations.
type GetRecommendationsApplyOutput struct{ *pulumi.OutputState }

func (GetRecommendationsApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecommendationsArgs)(nil)).Elem()
}

func (o GetRecommendationsApplyOutput) ToGetRecommendationsApplyOutput() GetRecommendationsApplyOutput {
	return o
}

func (o GetRecommendationsApplyOutput) ToGetRecommendationsApplyOutputWithContext(ctx context.Context) GetRecommendationsApplyOutput {
	return o
}

// Specifies a list of categories in which the Advisor Recommendations will be listed. Possible values are `HighAvailability`, `Security`, `Performance`, `Cost` and `OperationalExcellence`.
func (o GetRecommendationsApplyOutput) FilterByCategories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRecommendationsArgs) []string { return v.FilterByCategories }).(pulumi.StringArrayOutput)
}

// Specifies a list of resource groups about which the Advisor Recommendations will be listed.
func (o GetRecommendationsApplyOutput) FilterByResourceGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRecommendationsArgs) []string { return v.FilterByResourceGroups }).(pulumi.StringArrayOutput)
}

// A collection of values returned by getRecommendations.
type GetRecommendationsResultOutput struct{ *pulumi.OutputState }

func (GetRecommendationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecommendationsResult)(nil)).Elem()
}

func (o GetRecommendationsResultOutput) ToGetRecommendationsResultOutput() GetRecommendationsResultOutput {
	return o
}

func (o GetRecommendationsResultOutput) ToGetRecommendationsResultOutputWithContext(ctx context.Context) GetRecommendationsResultOutput {
	return o
}

func (o GetRecommendationsResultOutput) FilterByCategories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRecommendationsResult) []string { return v.FilterByCategories }).(pulumi.StringArrayOutput)
}

func (o GetRecommendationsResultOutput) FilterByResourceGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRecommendationsResult) []string { return v.FilterByResourceGroups }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRecommendationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecommendationsResult) string { return v.Id }).(pulumi.StringOutput)
}

// One or more `recommendations` blocks as defined below.
func (o GetRecommendationsResultOutput) Recommendations() GetRecommendationsRecommendationArrayOutput {
	return o.ApplyT(func(v GetRecommendationsResult) []GetRecommendationsRecommendation { return v.Recommendations }).(GetRecommendationsRecommendationArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRecommendationsApplyOutput{})
	pulumi.RegisterOutputType(GetRecommendationsResultOutput{})
}
