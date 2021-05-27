// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing SQL database.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mssql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := mssql.LookupDatabase(ctx, &mssql.LookupDatabaseArgs{
// 			Name:     "example-mssql-db",
// 			ServerId: "example-mssql-server-id",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("databaseId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure:mssql/getDatabase:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabase.
type LookupDatabaseArgs struct {
	// The name of the Ms SQL Database.
	Name string `pulumi:"name"`
	// The id of the Ms SQL Server on which to create the database.
	ServerId string `pulumi:"serverId"`
}

// A collection of values returned by getDatabase.
type LookupDatabaseResult struct {
	// The collation of the database.
	Collation string `pulumi:"collation"`
	// The id of the elastic pool containing this database.
	ElasticPoolId string `pulumi:"elasticPoolId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The license type to apply for this database.
	LicenseType string `pulumi:"licenseType"`
	// The max size of the database in gigabytes.
	MaxSizeGb int    `pulumi:"maxSizeGb"`
	Name      string `pulumi:"name"`
	// The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed.
	ReadReplicaCount int `pulumi:"readReplicaCount"`
	// If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica.
	ReadScale bool   `pulumi:"readScale"`
	ServerId  string `pulumi:"serverId"`
	// The name of the sku of the database.
	SkuName string `pulumi:"skuName"`
	// The storage account type used to store backups for this database.
	StorageAccountType string `pulumi:"storageAccountType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant bool `pulumi:"zoneRedundant"`
}

func LookupDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseResult, error) {
			args := v.(LookupDatabaseArgs)
			r, err := LookupDatabase(ctx, &args, opts...)
			return *r, err
		}).(LookupDatabaseResultOutput)
}

// A collection of arguments for invoking getDatabase.
type LookupDatabaseOutputArgs struct {
	// The name of the Ms SQL Database.
	Name pulumi.StringInput `pulumi:"name"`
	// The id of the Ms SQL Server on which to create the database.
	ServerId pulumi.StringInput `pulumi:"serverId"`
}

func (LookupDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getDatabase.
type LookupDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutput() LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseResultOutput {
	return o
}

// The collation of the database.
func (o LookupDatabaseResultOutput) Collation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Collation }).(pulumi.StringOutput)
}

// The id of the elastic pool containing this database.
func (o LookupDatabaseResultOutput) ElasticPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ElasticPoolId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// The license type to apply for this database.
func (o LookupDatabaseResultOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.LicenseType }).(pulumi.StringOutput)
}

// The max size of the database in gigabytes.
func (o LookupDatabaseResultOutput) MaxSizeGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseResult) int { return v.MaxSizeGb }).(pulumi.IntOutput)
}

func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed.
func (o LookupDatabaseResultOutput) ReadReplicaCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseResult) int { return v.ReadReplicaCount }).(pulumi.IntOutput)
}

// If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica.
func (o LookupDatabaseResultOutput) ReadScale() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseResult) bool { return v.ReadScale }).(pulumi.BoolOutput)
}

func (o LookupDatabaseResultOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ServerId }).(pulumi.StringOutput)
}

// The name of the sku of the database.
func (o LookupDatabaseResultOutput) SkuName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.SkuName }).(pulumi.StringOutput)
}

// The storage account type used to store backups for this database.
func (o LookupDatabaseResultOutput) StorageAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.StorageAccountType }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o LookupDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
func (o LookupDatabaseResultOutput) ZoneRedundant() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseResult) bool { return v.ZoneRedundant }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}
