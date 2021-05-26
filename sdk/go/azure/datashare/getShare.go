// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Data Share.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datashare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := datashare.LookupAccount(ctx, &datashare.LookupAccountArgs{
// 			Name:              "example-account",
// 			ResourceGroupName: "example-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleShare, err := datashare.LookupShare(ctx, &datashare.LookupShareArgs{
// 			Name:      "existing",
// 			AccountId: exampleAccount.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", exampleShare.Id)
// 		return nil
// 	})
// }
// ```
func LookupShare(ctx *pulumi.Context, args *LookupShareArgs, opts ...pulumi.InvokeOption) (*LookupShareResult, error) {
	var rv LookupShareResult
	err := ctx.Invoke("azure:datashare/getShare:getShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getShare.
type LookupShareArgs struct {
	// The ID of the Data Share account in which the Data Share is created.
	AccountId string `pulumi:"accountId"`
	// The name of this Data Share.
	Name string `pulumi:"name"`
}

// A collection of values returned by getShare.
type LookupShareResult struct {
	AccountId string `pulumi:"accountId"`
	// The description of the Data Share.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The kind of the Data Share.
	Kind string `pulumi:"kind"`
	// The name of the snapshot schedule.
	Name string `pulumi:"name"`
	// A `snapshotSchedule` block as defined below.
	SnapshotSchedules []GetShareSnapshotSchedule `pulumi:"snapshotSchedules"`
	// The terms of the Data Share.
	Terms string `pulumi:"terms"`
}

func LookupShareApply(ctx *pulumi.Context, args LookupShareApplyInput, opts ...pulumi.InvokeOption) LookupShareResultOutput {
	return args.ToLookupShareApplyOutput().ApplyT(func(v LookupShareArgs) (LookupShareResult, error) {
		r, err := LookupShare(ctx, &v, opts...)
		return *r, err

	}).(LookupShareResultOutput)
}

// LookupShareApplyInput is an input type that accepts LookupShareApplyArgs and LookupShareApplyOutput values.
// You can construct a concrete instance of `LookupShareApplyInput` via:
//
//          LookupShareApplyArgs{...}
type LookupShareApplyInput interface {
	pulumi.Input

	ToLookupShareApplyOutput() LookupShareApplyOutput
	ToLookupShareApplyOutputWithContext(context.Context) LookupShareApplyOutput
}

// A collection of arguments for invoking getShare.
type LookupShareApplyArgs struct {
	// The ID of the Data Share account in which the Data Share is created.
	AccountId pulumi.StringInput `pulumi:"accountId"`
	// The name of this Data Share.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupShareApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareArgs)(nil)).Elem()
}

func (i LookupShareApplyArgs) ToLookupShareApplyOutput() LookupShareApplyOutput {
	return i.ToLookupShareApplyOutputWithContext(context.Background())
}

func (i LookupShareApplyArgs) ToLookupShareApplyOutputWithContext(ctx context.Context) LookupShareApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupShareApplyOutput)
}

// A collection of arguments for invoking getShare.
type LookupShareApplyOutput struct{ *pulumi.OutputState }

func (LookupShareApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareArgs)(nil)).Elem()
}

func (o LookupShareApplyOutput) ToLookupShareApplyOutput() LookupShareApplyOutput {
	return o
}

func (o LookupShareApplyOutput) ToLookupShareApplyOutputWithContext(ctx context.Context) LookupShareApplyOutput {
	return o
}

// The ID of the Data Share account in which the Data Share is created.
func (o LookupShareApplyOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareArgs) string { return v.AccountId }).(pulumi.StringOutput)
}

// The name of this Data Share.
func (o LookupShareApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareArgs) string { return v.Name }).(pulumi.StringOutput)
}

// A collection of values returned by getShare.
type LookupShareResultOutput struct{ *pulumi.OutputState }

func (LookupShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareResult)(nil)).Elem()
}

func (o LookupShareResultOutput) ToLookupShareResultOutput() LookupShareResultOutput {
	return o
}

func (o LookupShareResultOutput) ToLookupShareResultOutputWithContext(ctx context.Context) LookupShareResultOutput {
	return o
}

func (o LookupShareResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// The description of the Data Share.
func (o LookupShareResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the Data Share.
func (o LookupShareResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the snapshot schedule.
func (o LookupShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Name }).(pulumi.StringOutput)
}

// A `snapshotSchedule` block as defined below.
func (o LookupShareResultOutput) SnapshotSchedules() GetShareSnapshotScheduleArrayOutput {
	return o.ApplyT(func(v LookupShareResult) []GetShareSnapshotSchedule { return v.SnapshotSchedules }).(GetShareSnapshotScheduleArrayOutput)
}

// The terms of the Data Share.
func (o LookupShareResultOutput) Terms() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Terms }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupShareApplyOutput{})
	pulumi.RegisterOutputType(LookupShareResultOutput{})
}
