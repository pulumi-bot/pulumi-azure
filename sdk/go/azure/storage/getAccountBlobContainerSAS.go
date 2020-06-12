// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account Blob Container.
//
// Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account Blob Container.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("westus"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		storage, err := storage.NewAccount(ctx, "storage", &storage.AccountArgs{
// 			ResourceGroupName:      rg.Name,
// 			Location:               rg.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		container, err := storage.NewContainer(ctx, "container", &storage.ContainerArgs{
// 			StorageAccountName:  storage.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("sasUrlQueryString", example.ApplyT(func(example storage.LookupAccountBlobContainerSASResult) (string, error) {
// 			return example.Sas, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
func GetAccountBlobContainerSAS(ctx *pulumi.Context, args *GetAccountBlobContainerSASArgs, opts ...pulumi.InvokeOption) (*GetAccountBlobContainerSASResult, error) {
	var rv GetAccountBlobContainerSASResult
	err := ctx.Invoke("azure:storage/getAccountBlobContainerSAS:getAccountBlobContainerSAS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountBlobContainerSAS.
type GetAccountBlobContainerSASArgs struct {
	// The `Cache-Control` response header that is sent when this SAS token is used.
	CacheControl *string `pulumi:"cacheControl"`
	// The connection string for the storage account to which this SAS applies. Typically directly from the `primaryConnectionString` attribute of an `storage.Account` resource.
	ConnectionString string `pulumi:"connectionString"`
	// Name of the container.
	ContainerName string `pulumi:"containerName"`
	// The `Content-Disposition` response header that is sent when this SAS token is used.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// The `Content-Encoding` response header that is sent when this SAS token is used.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// The `Content-Language` response header that is sent when this SAS token is used.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// The `Content-Type` response header that is sent when this SAS token is used.
	ContentType *string `pulumi:"contentType"`
	// The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
	Expiry string `pulumi:"expiry"`
	// Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// Single ipv4 address or range (connected with a dash) of ipv4 addresses.
	IpAddress *string `pulumi:"ipAddress"`
	// A `permissions` block as defined below.
	Permissions GetAccountBlobContainerSASPermissions `pulumi:"permissions"`
	// The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
	Start string `pulumi:"start"`
}

// A collection of values returned by getAccountBlobContainerSAS.
type GetAccountBlobContainerSASResult struct {
	CacheControl       *string `pulumi:"cacheControl"`
	ConnectionString   string  `pulumi:"connectionString"`
	ContainerName      string  `pulumi:"containerName"`
	ContentDisposition *string `pulumi:"contentDisposition"`
	ContentEncoding    *string `pulumi:"contentEncoding"`
	ContentLanguage    *string `pulumi:"contentLanguage"`
	ContentType        *string `pulumi:"contentType"`
	Expiry             string  `pulumi:"expiry"`
	HttpsOnly          *bool   `pulumi:"httpsOnly"`
	// The provider-assigned unique ID for this managed resource.
	Id          string                                `pulumi:"id"`
	IpAddress   *string                               `pulumi:"ipAddress"`
	Permissions GetAccountBlobContainerSASPermissions `pulumi:"permissions"`
	// The computed Blob Container Shared Access Signature (SAS).
	Sas   string `pulumi:"sas"`
	Start string `pulumi:"start"`
}
