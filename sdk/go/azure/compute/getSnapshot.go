// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Snapshot.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := compute.LookupSnapshot(ctx, &compute.LookupSnapshotArgs{
// 			Name:              "my-snapshot",
// 			ResourceGroupName: "my-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	var rv LookupSnapshotResult
	err := ctx.Invoke("azure:compute/getSnapshot:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshot.
type LookupSnapshotArgs struct {
	// Specifies the name of the Snapshot.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Snapshot is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getSnapshot.
type LookupSnapshotResult struct {
	CreationOption string `pulumi:"creationOption"`
	// The size of the Snapshotted Disk in GB.
	DiskSizeGb         int                            `pulumi:"diskSizeGb"`
	EncryptionSettings []GetSnapshotEncryptionSetting `pulumi:"encryptionSettings"`
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	OsType            string `pulumi:"osType"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference to an existing snapshot.
	SourceResourceId string `pulumi:"sourceResourceId"`
	// The URI to a Managed or Unmanaged Disk.
	SourceUri string `pulumi:"sourceUri"`
	// The ID of an storage account.
	StorageAccountId string `pulumi:"storageAccountId"`
	TimeCreated      string `pulumi:"timeCreated"`
}
