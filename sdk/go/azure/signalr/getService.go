// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalr

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Azure SignalR service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/signalr"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := signalr.LookupService(ctx, &signalr.LookupServiceArgs{
// 			Name:              "test-signalr",
// 			ResourceGroupName: "signalr-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure:signalr/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type LookupServiceArgs struct {
	// Specifies the name of the SignalR service.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the SignalR service is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getService.
type LookupServiceResult struct {
	// The FQDN of the SignalR service.
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The publicly accessible IP of the SignalR service.
	IpAddress string `pulumi:"ipAddress"`
	// Specifies the supported Azure location where the SignalR service exists.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// The primary access key of the SignalR service.
	PrimaryAccessKey string `pulumi:"primaryAccessKey"`
	// The primary connection string of the SignalR service.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// The publicly accessible port of the SignalR service which is designed for browser/client use.
	PublicPort        int    `pulumi:"publicPort"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The secondary access key of the SignalR service.
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
	// The secondary connection string of the SignalR service.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// The publicly accessible port of the SignalR service which is designed for customer server side use.
	ServerPort int               `pulumi:"serverPort"`
	Tags       map[string]string `pulumi:"tags"`
}
