// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing SQL Azure Database Server.
//
// ## Example Usage
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
// 		example, err := sql.LookupServer(ctx, &sql.LookupServerArgs{
// 			Name:              "examplesqlservername",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("sqlServerId", example.Id)
// 		return nil
// 	})
// }
// ```
func GetServer(ctx *pulumi.Context, args *GetServerArgs, opts ...pulumi.InvokeOption) (*GetServerResult, error) {
	var rv GetServerResult
	err := ctx.Invoke("azure:sql/getServer:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServer.
type GetServerArgs struct {
	// The name of the SQL Server.
	Name string `pulumi:"name"`
	// Specifies the name of the Resource Group where the SQL Server exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getServer.
type GetServerResult struct {
	// The administrator username of the SQL Server.
	AdministratorLogin string `pulumi:"administratorLogin"`
	// The fully qualified domain name of the SQL Server.
	Fqdn string `pulumi:"fqdn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An `identity` block as defined below.
	Identities []GetServerIdentity `pulumi:"identities"`
	// The location of the Resource Group in which the SQL Server exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The version of the SQL Server.
	Version string `pulumi:"version"`
}
