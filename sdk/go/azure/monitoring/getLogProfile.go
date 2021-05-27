// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access the properties of a Log Profile.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/monitoring"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := monitoring.LookupLogProfile(ctx, &monitoring.LookupLogProfileArgs{
// 			Name: "test-logprofile",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("logProfileStorageAccountId", example.StorageAccountId)
// 		return nil
// 	})
// }
// ```
func LookupLogProfile(ctx *pulumi.Context, args *LookupLogProfileArgs, opts ...pulumi.InvokeOption) (*LookupLogProfileResult, error) {
	var rv LookupLogProfileResult
	err := ctx.Invoke("azure:monitoring/getLogProfile:getLogProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogProfile.
type LookupLogProfileArgs struct {
	// Specifies the Name of the Log Profile.
	Name string `pulumi:"name"`
}

// A collection of values returned by getLogProfile.
type LookupLogProfileResult struct {
	// List of categories of the logs.
	Categories []string `pulumi:"categories"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of regions for which Activity Log events are stored or streamed.
	Locations         []string                       `pulumi:"locations"`
	Name              string                         `pulumi:"name"`
	RetentionPolicies []GetLogProfileRetentionPolicy `pulumi:"retentionPolicies"`
	// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to.
	ServicebusRuleId string `pulumi:"servicebusRuleId"`
	// The resource id of the storage account in which the Activity Log is stored.
	StorageAccountId string `pulumi:"storageAccountId"`
}

func LookupLogProfileOutput(ctx *pulumi.Context, args LookupLogProfileOutputArgs, opts ...pulumi.InvokeOption) LookupLogProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLogProfileResult, error) {
			args := v.(LookupLogProfileArgs)
			r, err := LookupLogProfile(ctx, &args, opts...)
			return *r, err
		}).(LookupLogProfileResultOutput)
}

// A collection of arguments for invoking getLogProfile.
type LookupLogProfileOutputArgs struct {
	// Specifies the Name of the Log Profile.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupLogProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogProfileArgs)(nil)).Elem()
}

// A collection of values returned by getLogProfile.
type LookupLogProfileResultOutput struct{ *pulumi.OutputState }

func (LookupLogProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogProfileResult)(nil)).Elem()
}

func (o LookupLogProfileResultOutput) ToLookupLogProfileResultOutput() LookupLogProfileResultOutput {
	return o
}

func (o LookupLogProfileResultOutput) ToLookupLogProfileResultOutputWithContext(ctx context.Context) LookupLogProfileResultOutput {
	return o
}

// List of categories of the logs.
func (o LookupLogProfileResultOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLogProfileResult) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLogProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of regions for which Activity Log events are stored or streamed.
func (o LookupLogProfileResultOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLogProfileResult) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o LookupLogProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLogProfileResultOutput) RetentionPolicies() GetLogProfileRetentionPolicyArrayOutput {
	return o.ApplyT(func(v LookupLogProfileResult) []GetLogProfileRetentionPolicy { return v.RetentionPolicies }).(GetLogProfileRetentionPolicyArrayOutput)
}

// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to.
func (o LookupLogProfileResultOutput) ServicebusRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogProfileResult) string { return v.ServicebusRuleId }).(pulumi.StringOutput)
}

// The resource id of the storage account in which the Activity Log is stored.
func (o LookupLogProfileResultOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogProfileResult) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLogProfileResultOutput{})
}
