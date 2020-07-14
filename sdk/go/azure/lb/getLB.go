// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Load Balancer
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := lb.GetLB(ctx, "azure:lb:getLB", &lb.GetLBArgs{
// 			Name:              "example-lb",
// 			ResourceGroupName: "example-resources",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("loadbalancerId", example.Id)
// 		return nil
// 	})
// }
// ```
func GetLB(ctx *pulumi.Context, args *GetLBArgs, opts ...pulumi.InvokeOption) (*GetLBResult, error) {
	var rv GetLBResult
	err := ctx.Invoke("azure:lb/getLB:getLB", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLB.
type GetLBArgs struct {
	// Specifies the name of the Load Balancer.
	Name string `pulumi:"name"`
	// The name of the Resource Group in which the Load Balancer exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getLB.
type GetLBResult struct {
	// (Optional) A `frontendIpConfiguration` block as documented below.
	FrontendIpConfigurations []GetLBFrontendIpConfiguration `pulumi:"frontendIpConfigurations"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure location where the Load Balancer exists.
	Location string `pulumi:"location"`
	// The name of the Frontend IP Configuration.
	Name string `pulumi:"name"`
	// Private IP Address to assign to the Load Balancer.
	PrivateIpAddress string `pulumi:"privateIpAddress"`
	// The list of private IP address assigned to the load balancer in `frontendIpConfiguration` blocks, if any.
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	// The SKU of the Load Balancer.
	Sku string `pulumi:"sku"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}
