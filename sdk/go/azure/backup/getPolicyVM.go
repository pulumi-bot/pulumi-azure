// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing VM Backup Policy.
//
// ## Example Usage
//
//
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
// 		_, err := backup.LookupPolicyVM(ctx, &backup.LookupPolicyVMArgs{
// 			Name:              "policy",
// 			RecoveryVaultName: "recovery_vault",
// 			ResourceGroupName: "resource_group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupPolicyVM(ctx *pulumi.Context, args *LookupPolicyVMArgs, opts ...pulumi.InvokeOption) (*LookupPolicyVMResult, error) {
	var rv LookupPolicyVMResult
	err := ctx.Invoke("azure:backup/getPolicyVM:getPolicyVM", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyVM.
type LookupPolicyVMArgs struct {
	// Specifies the name of the VM Backup Policy.
	Name string `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which the VM Backup Policy resides.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getPolicyVM.
type LookupPolicyVMResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}
