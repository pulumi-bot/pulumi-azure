// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Load Balancer's Backend Address Pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLB, err := lb.GetLB(ctx, &lb.GetLBArgs{
// 			Name:              "example-lb",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleBackendAddressPool, err := lb.LookupBackendAddressPool(ctx, &lb.LookupBackendAddressPoolArgs{
// 			Name:           "first",
// 			LoadbalancerId: exampleLB.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("backendAddressPoolId", exampleBackendAddressPool.Id)
// 		var splat0 []interface{}
// 		for _, val0 := range data.Azurerm_lb_backend_address_pool.Beap.Backend_ip_configurations {
// 			splat0 = append(splat0, val0.Id)
// 		}
// 		ctx.Export("backendIpConfigurationIds", splat0)
// 		return nil
// 	})
// }
// ```
func LookupBackendAddressPool(ctx *pulumi.Context, args *LookupBackendAddressPoolArgs, opts ...pulumi.InvokeOption) (*LookupBackendAddressPoolResult, error) {
	var rv LookupBackendAddressPoolResult
	err := ctx.Invoke("azure:lb/getBackendAddressPool:getBackendAddressPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackendAddressPool.
type LookupBackendAddressPoolArgs struct {
	// The ID of the Load Balancer in which the Backend Address Pool exists.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Specifies the name of the Backend Address Pool.
	Name string `pulumi:"name"`
}

// A collection of values returned by getBackendAddressPool.
type LookupBackendAddressPoolResult struct {
	// A list of `backendAddress` block as defined below.
	BackendAddresses []GetBackendAddressPoolBackendAddress `pulumi:"backendAddresses"`
	// A list of references to IP addresses defined in network interfaces.
	BackendIpConfigurations []GetBackendAddressPoolBackendIpConfiguration `pulumi:"backendIpConfigurations"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the Load Balancing Rules associated with this Backend Address Pool.
	LoadBalancingRules []string `pulumi:"loadBalancingRules"`
	LoadbalancerId     string   `pulumi:"loadbalancerId"`
	// The name of the Backend Address.
	Name string `pulumi:"name"`
	// A list of the Load Balancing Outbound Rules associated with this Backend Address Pool.
	OutboundRules []string `pulumi:"outboundRules"`
}
