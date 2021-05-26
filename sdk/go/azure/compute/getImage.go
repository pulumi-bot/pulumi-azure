// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Image.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "search-api"
// 		search, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
// 			Name:              &opt0,
// 			ResourceGroupName: "packerimages",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("imageId", search.Id)
// 		return nil
// 	})
// }
// ```
func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	var rv LookupImageResult
	err := ctx.Invoke("azure:compute/getImage:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImage.
type LookupImageArgs struct {
	// The name of the Image.
	Name *string `pulumi:"name"`
	// Regex pattern of the image to match.
	NameRegex *string `pulumi:"nameRegex"`
	// The Name of the Resource Group where this Image exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// By default when matching by regex, images are sorted by name in ascending order and the first match is chosen, to sort descending, set this flag.
	SortDescending *bool `pulumi:"sortDescending"`
}

// A collection of values returned by getImage.
type LookupImageResult struct {
	// a collection of `dataDisk` blocks as defined below.
	DataDisks []GetImageDataDisk `pulumi:"dataDisks"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// the Azure Location where this Image exists.
	Location string `pulumi:"location"`
	// the name of the Image.
	Name      *string `pulumi:"name"`
	NameRegex *string `pulumi:"nameRegex"`
	// a `osDisk` block as defined below.
	OsDisks           []GetImageOsDisk `pulumi:"osDisks"`
	ResourceGroupName string           `pulumi:"resourceGroupName"`
	SortDescending    *bool            `pulumi:"sortDescending"`
	// a mapping of tags to assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// is zone resiliency enabled?
	ZoneResilient bool `pulumi:"zoneResilient"`
}

func LookupImageApply(ctx *pulumi.Context, args LookupImageApplyInput, opts ...pulumi.InvokeOption) LookupImageResultOutput {
	return args.ToLookupImageApplyOutput().ApplyT(func(v LookupImageArgs) (LookupImageResult, error) {
		r, err := LookupImage(ctx, &v, opts...)
		return *r, err

	}).(LookupImageResultOutput)
}

// LookupImageApplyInput is an input type that accepts LookupImageApplyArgs and LookupImageApplyOutput values.
// You can construct a concrete instance of `LookupImageApplyInput` via:
//
//          LookupImageApplyArgs{...}
type LookupImageApplyInput interface {
	pulumi.Input

	ToLookupImageApplyOutput() LookupImageApplyOutput
	ToLookupImageApplyOutputWithContext(context.Context) LookupImageApplyOutput
}

// A collection of arguments for invoking getImage.
type LookupImageApplyArgs struct {
	// The name of the Image.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Regex pattern of the image to match.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The Name of the Resource Group where this Image exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// By default when matching by regex, images are sorted by name in ascending order and the first match is chosen, to sort descending, set this flag.
	SortDescending pulumi.BoolPtrInput `pulumi:"sortDescending"`
}

func (LookupImageApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageArgs)(nil)).Elem()
}

func (i LookupImageApplyArgs) ToLookupImageApplyOutput() LookupImageApplyOutput {
	return i.ToLookupImageApplyOutputWithContext(context.Background())
}

func (i LookupImageApplyArgs) ToLookupImageApplyOutputWithContext(ctx context.Context) LookupImageApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupImageApplyOutput)
}

// A collection of arguments for invoking getImage.
type LookupImageApplyOutput struct{ *pulumi.OutputState }

func (LookupImageApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageArgs)(nil)).Elem()
}

func (o LookupImageApplyOutput) ToLookupImageApplyOutput() LookupImageApplyOutput {
	return o
}

func (o LookupImageApplyOutput) ToLookupImageApplyOutputWithContext(ctx context.Context) LookupImageApplyOutput {
	return o
}

// The name of the Image.
func (o LookupImageApplyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageArgs) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Regex pattern of the image to match.
func (o LookupImageApplyOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageArgs) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// The Name of the Resource Group where this Image exists.
func (o LookupImageApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// By default when matching by regex, images are sorted by name in ascending order and the first match is chosen, to sort descending, set this flag.
func (o LookupImageApplyOutput) SortDescending() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupImageArgs) *bool { return v.SortDescending }).(pulumi.BoolPtrOutput)
}

// A collection of values returned by getImage.
type LookupImageResultOutput struct{ *pulumi.OutputState }

func (LookupImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageResult)(nil)).Elem()
}

func (o LookupImageResultOutput) ToLookupImageResultOutput() LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToLookupImageResultOutputWithContext(ctx context.Context) LookupImageResultOutput {
	return o
}

// a collection of `dataDisk` blocks as defined below.
func (o LookupImageResultOutput) DataDisks() GetImageDataDiskArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []GetImageDataDisk { return v.DataDisks }).(GetImageDataDiskArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// the Azure Location where this Image exists.
func (o LookupImageResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Location }).(pulumi.StringOutput)
}

// the name of the Image.
func (o LookupImageResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupImageResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// a `osDisk` block as defined below.
func (o LookupImageResultOutput) OsDisks() GetImageOsDiskArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []GetImageOsDisk { return v.OsDisks }).(GetImageOsDiskArrayOutput)
}

func (o LookupImageResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) SortDescending() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *bool { return v.SortDescending }).(pulumi.BoolPtrOutput)
}

// a mapping of tags to assigned to the resource.
func (o LookupImageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// is zone resiliency enabled?
func (o LookupImageResultOutput) ZoneResilient() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.ZoneResilient }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageApplyOutput{})
	pulumi.RegisterOutputType(LookupImageResultOutput{})
}
