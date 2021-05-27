// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Storage Table Entity.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storage.LookupTableEntity(ctx, &storage.LookupTableEntityArgs{
// 			PartitionKey:       "example-parition-key",
// 			RowKey:             "example-row-key",
// 			StorageAccountName: "example-storage-account-name",
// 			TableName:          "example-table-name",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupTableEntity(ctx *pulumi.Context, args *LookupTableEntityArgs, opts ...pulumi.InvokeOption) (*LookupTableEntityResult, error) {
	var rv LookupTableEntityResult
	err := ctx.Invoke("azure:storage/getTableEntity:getTableEntity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTableEntity.
type LookupTableEntityArgs struct {
	// The key for the partition where the entity will be retrieved.
	PartitionKey string `pulumi:"partitionKey"`
	// The key for the row where the entity will be retrieved.
	RowKey string `pulumi:"rowKey"`
	// The name of the Storage Account where the Table exists.
	StorageAccountName string `pulumi:"storageAccountName"`
	// The name of the Table.
	TableName string `pulumi:"tableName"`
}

// A collection of values returned by getTableEntity.
type LookupTableEntityResult struct {
	// A map of key/value pairs that describe the entity to be stored in the storage table.
	Entity map[string]string `pulumi:"entity"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string `pulumi:"id"`
	PartitionKey       string `pulumi:"partitionKey"`
	RowKey             string `pulumi:"rowKey"`
	StorageAccountName string `pulumi:"storageAccountName"`
	TableName          string `pulumi:"tableName"`
}

func LookupTableEntityOutput(ctx *pulumi.Context, args LookupTableEntityOutputArgs, opts ...pulumi.InvokeOption) LookupTableEntityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTableEntityResult, error) {
			args := v.(LookupTableEntityArgs)
			r, err := LookupTableEntity(ctx, &args, opts...)
			return *r, err
		}).(LookupTableEntityResultOutput)
}

// A collection of arguments for invoking getTableEntity.
type LookupTableEntityOutputArgs struct {
	// The key for the partition where the entity will be retrieved.
	PartitionKey pulumi.StringInput `pulumi:"partitionKey"`
	// The key for the row where the entity will be retrieved.
	RowKey pulumi.StringInput `pulumi:"rowKey"`
	// The name of the Storage Account where the Table exists.
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
	// The name of the Table.
	TableName pulumi.StringInput `pulumi:"tableName"`
}

func (LookupTableEntityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableEntityArgs)(nil)).Elem()
}

// A collection of values returned by getTableEntity.
type LookupTableEntityResultOutput struct{ *pulumi.OutputState }

func (LookupTableEntityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableEntityResult)(nil)).Elem()
}

func (o LookupTableEntityResultOutput) ToLookupTableEntityResultOutput() LookupTableEntityResultOutput {
	return o
}

func (o LookupTableEntityResultOutput) ToLookupTableEntityResultOutputWithContext(ctx context.Context) LookupTableEntityResultOutput {
	return o
}

// A map of key/value pairs that describe the entity to be stored in the storage table.
func (o LookupTableEntityResultOutput) Entity() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTableEntityResult) map[string]string { return v.Entity }).(pulumi.StringMapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTableEntityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableEntityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTableEntityResultOutput) PartitionKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableEntityResult) string { return v.PartitionKey }).(pulumi.StringOutput)
}

func (o LookupTableEntityResultOutput) RowKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableEntityResult) string { return v.RowKey }).(pulumi.StringOutput)
}

func (o LookupTableEntityResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableEntityResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupTableEntityResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableEntityResult) string { return v.TableName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableEntityResultOutput{})
}
