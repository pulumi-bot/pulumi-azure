// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Virtual Machine.
func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure:compute/getVirtualMachine:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualMachine.
type LookupVirtualMachineArgs struct {
	// Specifies the name of the Virtual Machine.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Virtual Machine is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getVirtualMachine.
type LookupVirtualMachineResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A `identity` block as defined below.
	Identities        []GetVirtualMachineIdentity `pulumi:"identities"`
	Location          string                      `pulumi:"location"`
	Name              string                      `pulumi:"name"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
}
