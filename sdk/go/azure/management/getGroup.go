// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package management

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Management Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "00000000-0000-0000-0000-000000000000"
// 		example, err := management.LookupGroup(ctx, &management.LookupGroupArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("displayName", example.DisplayName)
// 		return nil
// 	})
// }
// ```
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("azure:management/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// Specifies the display name of this Management Group.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the name or UUID of this Management Group.
	//
	// Deprecated: Deprecated in favour of `name`
	GroupId *string `pulumi:"groupId"`
	// Specifies the name or UUID of this Management Group.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	DisplayName string `pulumi:"displayName"`
	// Deprecated: Deprecated in favour of `name`
	GroupId string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The ID of any Parent Management Group.
	ParentManagementGroupId string `pulumi:"parentManagementGroupId"`
	// A list of Subscription IDs which are assigned to the Management Group.
	SubscriptionIds []string `pulumi:"subscriptionIds"`
}

func LookupGroupApply(ctx *pulumi.Context, args LookupGroupApplyInput, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return args.ToLookupGroupApplyOutput().ApplyT(func(v LookupGroupArgs) (LookupGroupResult, error) {
		r, err := LookupGroup(ctx, &v, opts...)
		return *r, err

	}).(LookupGroupResultOutput)
}

// LookupGroupApplyInput is an input type that accepts LookupGroupApplyArgs and LookupGroupApplyOutput values.
// You can construct a concrete instance of `LookupGroupApplyInput` via:
//
//          LookupGroupApplyArgs{...}
type LookupGroupApplyInput interface {
	pulumi.Input

	ToLookupGroupApplyOutput() LookupGroupApplyOutput
	ToLookupGroupApplyOutputWithContext(context.Context) LookupGroupApplyOutput
}

// A collection of arguments for invoking getGroup.
type LookupGroupApplyArgs struct {
	// Specifies the display name of this Management Group.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// Specifies the name or UUID of this Management Group.
	//
	// Deprecated: Deprecated in favour of `name`
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	// Specifies the name or UUID of this Management Group.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupGroupApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

func (i LookupGroupApplyArgs) ToLookupGroupApplyOutput() LookupGroupApplyOutput {
	return i.ToLookupGroupApplyOutputWithContext(context.Background())
}

func (i LookupGroupApplyArgs) ToLookupGroupApplyOutputWithContext(ctx context.Context) LookupGroupApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupGroupApplyOutput)
}

// A collection of arguments for invoking getGroup.
type LookupGroupApplyOutput struct{ *pulumi.OutputState }

func (LookupGroupApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

func (o LookupGroupApplyOutput) ToLookupGroupApplyOutput() LookupGroupApplyOutput {
	return o
}

func (o LookupGroupApplyOutput) ToLookupGroupApplyOutputWithContext(ctx context.Context) LookupGroupApplyOutput {
	return o
}

// Specifies the display name of this Management Group.
func (o LookupGroupApplyOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupArgs) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Specifies the name or UUID of this Management Group.
//
// Deprecated: Deprecated in favour of `name`
func (o LookupGroupApplyOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupArgs) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// Specifies the name or UUID of this Management Group.
func (o LookupGroupApplyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupArgs) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A collection of values returned by getGroup.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Deprecated: Deprecated in favour of `name`
func (o LookupGroupResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of any Parent Management Group.
func (o LookupGroupResultOutput) ParentManagementGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.ParentManagementGroupId }).(pulumi.StringOutput)
}

// A list of Subscription IDs which are assigned to the Management Group.
func (o LookupGroupResultOutput) SubscriptionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.SubscriptionIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupApplyOutput{})
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
