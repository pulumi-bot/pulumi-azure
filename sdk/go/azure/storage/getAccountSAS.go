// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account.
//
// Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account.
//
// Note that this is an [Account SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-an-account-sas)
// and *not* a [Service SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-a-service-sas).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               pulumi.String("westus"),
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("GRS"),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("sasUrlQueryString", exampleAccountSAS.ApplyT(func(exampleAccountSAS storage.GetAccountSASResult) (string, error) {
// 			return exampleAccountSAS.Sas, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
func GetAccountSAS(ctx *pulumi.Context, args *GetAccountSASArgs, opts ...pulumi.InvokeOption) (*GetAccountSASResult, error) {
	var rv GetAccountSASResult
	err := ctx.Invoke("azure:storage/getAccountSAS:getAccountSAS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountSAS.
type GetAccountSASArgs struct {
	// The connection string for the storage account to which this SAS applies. Typically directly from the `primaryConnectionString` attribute of a `storage.Account` resource.
	ConnectionString string `pulumi:"connectionString"`
	// The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
	Expiry string `pulumi:"expiry"`
	// Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// A `permissions` block as defined below.
	Permissions GetAccountSASPermissions `pulumi:"permissions"`
	// A `resourceTypes` block as defined below.
	ResourceTypes GetAccountSASResourceTypes `pulumi:"resourceTypes"`
	// A `services` block as defined below.
	Services GetAccountSASServices `pulumi:"services"`
	// Specifies the signed storage service version to use to authorize requests made with this account SAS. Defaults to `2017-07-29`.
	SignedVersion *string `pulumi:"signedVersion"`
	// The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
	Start string `pulumi:"start"`
}

// A collection of values returned by getAccountSAS.
type GetAccountSASResult struct {
	ConnectionString string `pulumi:"connectionString"`
	Expiry           string `pulumi:"expiry"`
	HttpsOnly        *bool  `pulumi:"httpsOnly"`
	// The provider-assigned unique ID for this managed resource.
	Id            string                     `pulumi:"id"`
	Permissions   GetAccountSASPermissions   `pulumi:"permissions"`
	ResourceTypes GetAccountSASResourceTypes `pulumi:"resourceTypes"`
	// The computed Account Shared Access Signature (SAS).
	Sas           string                `pulumi:"sas"`
	Services      GetAccountSASServices `pulumi:"services"`
	SignedVersion *string               `pulumi:"signedVersion"`
	Start         string                `pulumi:"start"`
}
