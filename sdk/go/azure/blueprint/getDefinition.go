// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blueprint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Azure Blueprint Definition
//
// > **NOTE:** Azure Blueprints are in Preview and potentially subject to breaking change without notice.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/blueprint"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := current.TenantId
// 		root, err := management.LookupGroup(ctx, &management.LookupGroupArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = blueprint.GetDefinition(ctx, &blueprint.GetDefinitionArgs{
// 			Name:    "exampleManagementGroupBP",
// 			ScopeId: root.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDefinition(ctx *pulumi.Context, args *GetDefinitionArgs, opts ...pulumi.InvokeOption) (*GetDefinitionResult, error) {
	var rv GetDefinitionResult
	err := ctx.Invoke("azure:blueprint/getDefinition:getDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefinition.
type GetDefinitionArgs struct {
	// The name of the Blueprint.
	Name string `pulumi:"name"`
	// The ID of the Subscription or Management Group, as the scope at which the blueprint definition is stored.
	ScopeId string `pulumi:"scopeId"`
}

// A collection of values returned by getDefinition.
type GetDefinitionResult struct {
	// The description of the Blueprint Definition.
	Description string `pulumi:"description"`
	// The display name of the Blueprint Definition.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The timestamp of when this last modification was saved to the Blueprint Definition.
	LastModified string `pulumi:"lastModified"`
	Name         string `pulumi:"name"`
	ScopeId      string `pulumi:"scopeId"`
	// The target scope.
	TargetScope string `pulumi:"targetScope"`
	// The timestamp of when this Blueprint Definition was created.
	TimeCreated string `pulumi:"timeCreated"`
	// A list of versions published for this Blueprint Definition.
	Versions []string `pulumi:"versions"`
}

func GetDefinitionApply(ctx *pulumi.Context, args GetDefinitionApplyInput, opts ...pulumi.InvokeOption) GetDefinitionResultOutput {
	return args.ToGetDefinitionApplyOutput().ApplyT(func(v GetDefinitionArgs) (GetDefinitionResult, error) {
		r, err := GetDefinition(ctx, &v, opts...)
		return *r, err

	}).(GetDefinitionResultOutput)
}

// GetDefinitionApplyInput is an input type that accepts GetDefinitionApplyArgs and GetDefinitionApplyOutput values.
// You can construct a concrete instance of `GetDefinitionApplyInput` via:
//
//          GetDefinitionApplyArgs{...}
type GetDefinitionApplyInput interface {
	pulumi.Input

	ToGetDefinitionApplyOutput() GetDefinitionApplyOutput
	ToGetDefinitionApplyOutputWithContext(context.Context) GetDefinitionApplyOutput
}

// A collection of arguments for invoking getDefinition.
type GetDefinitionApplyArgs struct {
	// The name of the Blueprint.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the Subscription or Management Group, as the scope at which the blueprint definition is stored.
	ScopeId pulumi.StringInput `pulumi:"scopeId"`
}

func (GetDefinitionApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDefinitionArgs)(nil)).Elem()
}

func (i GetDefinitionApplyArgs) ToGetDefinitionApplyOutput() GetDefinitionApplyOutput {
	return i.ToGetDefinitionApplyOutputWithContext(context.Background())
}

func (i GetDefinitionApplyArgs) ToGetDefinitionApplyOutputWithContext(ctx context.Context) GetDefinitionApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDefinitionApplyOutput)
}

// A collection of arguments for invoking getDefinition.
type GetDefinitionApplyOutput struct{ *pulumi.OutputState }

func (GetDefinitionApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDefinitionArgs)(nil)).Elem()
}

func (o GetDefinitionApplyOutput) ToGetDefinitionApplyOutput() GetDefinitionApplyOutput {
	return o
}

func (o GetDefinitionApplyOutput) ToGetDefinitionApplyOutputWithContext(ctx context.Context) GetDefinitionApplyOutput {
	return o
}

// The name of the Blueprint.
func (o GetDefinitionApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefinitionArgs) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the Subscription or Management Group, as the scope at which the blueprint definition is stored.
func (o GetDefinitionApplyOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefinitionArgs) string { return v.ScopeId }).(pulumi.StringOutput)
}

// A collection of values returned by getDefinition.
type GetDefinitionResultOutput struct{ *pulumi.OutputState }

func (GetDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDefinitionResult)(nil)).Elem()
}

func (o GetDefinitionResultOutput) ToGetDefinitionResultOutput() GetDefinitionResultOutput {
	return o
}

func (o GetDefinitionResultOutput) ToGetDefinitionResultOutputWithContext(ctx context.Context) GetDefinitionResultOutput {
	return o
}

// The description of the Blueprint Definition.
func (o GetDefinitionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefinitionResult) string { return v.Description }).(pulumi.StringOutput)
}

// The display name of the Blueprint Definition.
func (o GetDefinitionResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefinitionResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The timestamp of when this last modification was saved to the Blueprint Definition.
func (o GetDefinitionResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefinitionResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o GetDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDefinitionResultOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefinitionResult) string { return v.ScopeId }).(pulumi.StringOutput)
}

// The target scope.
func (o GetDefinitionResultOutput) TargetScope() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefinitionResult) string { return v.TargetScope }).(pulumi.StringOutput)
}

// The timestamp of when this Blueprint Definition was created.
func (o GetDefinitionResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefinitionResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

// A list of versions published for this Blueprint Definition.
func (o GetDefinitionResultOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDefinitionResult) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDefinitionApplyOutput{})
	pulumi.RegisterOutputType(GetDefinitionResultOutput{})
}
