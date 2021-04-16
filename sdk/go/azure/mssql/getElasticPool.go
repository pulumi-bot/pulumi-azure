// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing SQL elastic pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mssql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := mssql.LookupElasticPool(ctx, &mssql.LookupElasticPoolArgs{
// 			Name:              "mssqlelasticpoolname",
// 			ResourceGroupName: "example-resources",
// 			ServerName:        "example-sql-server",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("elasticpoolId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupElasticPool(ctx *pulumi.Context, args *LookupElasticPoolArgs, opts ...pulumi.InvokeOption) (*LookupElasticPoolResult, error) {
	var rv LookupElasticPoolResult
	err := ctx.Invoke("azure:mssql/getElasticPool:getElasticPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getElasticPool.
type LookupElasticPoolArgs struct {
	// The name of the elastic pool.
	Name string `pulumi:"name"`
	// The name of the resource group which contains the elastic pool.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SQL Server which contains the elastic pool.
	ServerName string `pulumi:"serverName"`
}

// A collection of values returned by getElasticPool.
type LookupElasticPoolResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The license type to apply for this database.
	LicenseType string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists.
	Location string `pulumi:"location"`
	// The max data size of the elastic pool in bytes.
	MaxSizeBytes int `pulumi:"maxSizeBytes"`
	// The max data size of the elastic pool in gigabytes.
	MaxSizeGb float64 `pulumi:"maxSizeGb"`
	Name      string  `pulumi:"name"`
	// The maximum capacity any one database can consume.
	PerDbMaxCapacity int `pulumi:"perDbMaxCapacity"`
	// The minimum capacity all databases are guaranteed.
	PerDbMinCapacity  int    `pulumi:"perDbMinCapacity"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this elastic pool is zone redundant.
	ZoneRedundant bool `pulumi:"zoneRedundant"`
}
