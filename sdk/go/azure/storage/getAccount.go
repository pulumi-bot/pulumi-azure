// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Storage Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "packer-storage"
// 		example, err := storage.LookupAccount(ctx, "azure:storage:getAccount", &storage.LookupAccountArgs{
// 			Name:              "packerimages",
// 			ResourceGroupName: &opt0,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("storageAccountTier", example.AccountTier)
// 		return nil
// 	})
// }
// ```
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure:storage/getAccount:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccount.
type LookupAccountArgs struct {
	// Specifies the name of the Storage Account
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Storage Account is located in.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getAccount.
type LookupAccountResult struct {
	// The access tier for `BlobStorage` accounts.
	AccessTier string `pulumi:"accessTier"`
	// The Kind of account.
	AccountKind string `pulumi:"accountKind"`
	// The type of replication used for this storage account.
	AccountReplicationType string `pulumi:"accountReplicationType"`
	// The Tier of this storage account.
	AccountTier string `pulumi:"accountTier"`
	// A `customDomain` block as documented below.
	CustomDomains []GetAccountCustomDomain `pulumi:"customDomains"`
	// Is traffic only allowed via HTTPS? See [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information.
	EnableHttpsTrafficOnly bool `pulumi:"enableHttpsTrafficOnly"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Is Hierarchical Namespace enabled?
	IsHnsEnabled bool `pulumi:"isHnsEnabled"`
	// The Azure location where the Storage Account exists
	Location string `pulumi:"location"`
	// The Custom Domain Name used for the Storage Account.
	Name string `pulumi:"name"`
	// The primary access key for the Storage Account.
	PrimaryAccessKey string `pulumi:"primaryAccessKey"`
	// The connection string associated with the primary blob location
	PrimaryBlobConnectionString string `pulumi:"primaryBlobConnectionString"`
	// The endpoint URL for blob storage in the primary location.
	PrimaryBlobEndpoint string `pulumi:"primaryBlobEndpoint"`
	// The hostname with port if applicable for blob storage in the primary location.
	PrimaryBlobHost string `pulumi:"primaryBlobHost"`
	// The connection string associated with the primary location
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// The endpoint URL for DFS storage in the primary location.
	PrimaryDfsEndpoint string `pulumi:"primaryDfsEndpoint"`
	// The hostname with port if applicable for DFS storage in the primary location.
	PrimaryDfsHost string `pulumi:"primaryDfsHost"`
	// The endpoint URL for file storage in the primary location.
	PrimaryFileEndpoint string `pulumi:"primaryFileEndpoint"`
	// The hostname with port if applicable for file storage in the primary location.
	PrimaryFileHost string `pulumi:"primaryFileHost"`
	// The primary location of the Storage Account.
	PrimaryLocation string `pulumi:"primaryLocation"`
	// The endpoint URL for queue storage in the primary location.
	PrimaryQueueEndpoint string `pulumi:"primaryQueueEndpoint"`
	// The hostname with port if applicable for queue storage in the primary location.
	PrimaryQueueHost string `pulumi:"primaryQueueHost"`
	// The endpoint URL for table storage in the primary location.
	PrimaryTableEndpoint string `pulumi:"primaryTableEndpoint"`
	// The hostname with port if applicable for table storage in the primary location.
	PrimaryTableHost string `pulumi:"primaryTableHost"`
	// The endpoint URL for web storage in the primary location.
	PrimaryWebEndpoint string `pulumi:"primaryWebEndpoint"`
	// The hostname with port if applicable for web storage in the primary location.
	PrimaryWebHost    string `pulumi:"primaryWebHost"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The secondary access key for the Storage Account.
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
	// The connection string associated with the secondary blob location
	SecondaryBlobConnectionString string `pulumi:"secondaryBlobConnectionString"`
	// The endpoint URL for blob storage in the secondary location.
	SecondaryBlobEndpoint string `pulumi:"secondaryBlobEndpoint"`
	// The hostname with port if applicable for blob storage in the secondary location.
	SecondaryBlobHost string `pulumi:"secondaryBlobHost"`
	// The connection string associated with the secondary location
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// The endpoint URL for DFS storage in the secondary location.
	SecondaryDfsEndpoint string `pulumi:"secondaryDfsEndpoint"`
	// The hostname with port if applicable for DFS storage in the secondary location.
	SecondaryDfsHost string `pulumi:"secondaryDfsHost"`
	// The endpoint URL for file storage in the secondary location.
	SecondaryFileEndpoint string `pulumi:"secondaryFileEndpoint"`
	// The hostname with port if applicable for file storage in the secondary location.
	SecondaryFileHost string `pulumi:"secondaryFileHost"`
	// The secondary location of the Storage Account.
	SecondaryLocation string `pulumi:"secondaryLocation"`
	// The endpoint URL for queue storage in the secondary location.
	SecondaryQueueEndpoint string `pulumi:"secondaryQueueEndpoint"`
	// The hostname with port if applicable for queue storage in the secondary location.
	SecondaryQueueHost string `pulumi:"secondaryQueueHost"`
	// The endpoint URL for table storage in the secondary location.
	SecondaryTableEndpoint string `pulumi:"secondaryTableEndpoint"`
	// The hostname with port if applicable for table storage in the secondary location.
	SecondaryTableHost string `pulumi:"secondaryTableHost"`
	// The endpoint URL for web storage in the secondary location.
	SecondaryWebEndpoint string `pulumi:"secondaryWebEndpoint"`
	// The hostname with port if applicable for web storage in the secondary location.
	SecondaryWebHost string `pulumi:"secondaryWebHost"`
	// A mapping of tags to assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}
