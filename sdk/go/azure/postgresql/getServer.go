// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing PostgreSQL Azure Database Server.
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
// 		example, err := postgresql.LookupServer(ctx, &postgresql.LookupServerArgs{
// 			Name:              "postgresql-server-1",
// 			ResourceGroupName: "api-rg-pro",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("postgresqlServerId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure:postgresql/getServer:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServer.
type LookupServerArgs struct {
	// The name of the PostgreSQL Server.
	Name string `pulumi:"name"`
	// Specifies the name of the Resource Group where the PostgreSQL Server exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getServer.
type LookupServerResult struct {
	// The administrator username of the PostgreSQL Server.
	AdministratorLogin string `pulumi:"administratorLogin"`
	// The fully qualified domain name of the PostgreSQL Server.
	Fqdn string `pulumi:"fqdn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The location of the Resource Group in which the PostgreSQL Server exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The version of the PostgreSQL Server.
	Version string `pulumi:"version"`
}
