// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manage a Dedicated Host within a Dedicated Host Group.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		exampleDedicatedHostGroup, err := compute.NewDedicatedHostGroup(ctx, "exampleDedicatedHostGroup", &compute.DedicatedHostGroupArgs{
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			Location:                 exampleResourceGroup.Location,
// 			PlatformFaultDomainCount: pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewDedicatedHost(ctx, "exampleDedicatedHost", &compute.DedicatedHostArgs{
// 			Location:             exampleResourceGroup.Location,
// 			DedicatedHostGroupId: exampleDedicatedHostGroup.ID(),
// 			SkuName:              pulumi.String("DSv3-Type1"),
// 			PlatformFaultDomain:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DedicatedHost struct {
	pulumi.CustomResourceState

	// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
	AutoReplaceOnFailure pulumi.BoolPtrOutput `pulumi:"autoReplaceOnFailure"`
	// Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
	DedicatedHostGroupId pulumi.StringOutput `pulumi:"dedicatedHostGroupId"`
	// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
	PlatformFaultDomain pulumi.IntOutput `pulumi:"platformFaultDomain"`
	// Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDedicatedHost registers a new resource with the given unique name, arguments, and options.
func NewDedicatedHost(ctx *pulumi.Context,
	name string, args *DedicatedHostArgs, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	if args == nil || args.DedicatedHostGroupId == nil {
		return nil, errors.New("missing required argument 'DedicatedHostGroupId'")
	}
	if args == nil || args.PlatformFaultDomain == nil {
		return nil, errors.New("missing required argument 'PlatformFaultDomain'")
	}
	if args == nil || args.SkuName == nil {
		return nil, errors.New("missing required argument 'SkuName'")
	}
	if args == nil {
		args = &DedicatedHostArgs{}
	}
	var resource DedicatedHost
	err := ctx.RegisterResource("azure:compute/dedicatedHost:DedicatedHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedHost gets an existing DedicatedHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostState, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	var resource DedicatedHost
	err := ctx.ReadResource("azure:compute/dedicatedHost:DedicatedHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedHost resources.
type dedicatedHostState struct {
	// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
	AutoReplaceOnFailure *bool `pulumi:"autoReplaceOnFailure"`
	// Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
	DedicatedHostGroupId *string `pulumi:"dedicatedHostGroupId"`
	// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
	LicenseType *string `pulumi:"licenseType"`
	// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
	PlatformFaultDomain *int `pulumi:"platformFaultDomain"`
	// Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type DedicatedHostState struct {
	// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
	AutoReplaceOnFailure pulumi.BoolPtrInput
	// Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
	DedicatedHostGroupId pulumi.StringPtrInput
	// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
	LicenseType pulumi.StringPtrInput
	// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
	PlatformFaultDomain pulumi.IntPtrInput
	// Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
	SkuName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DedicatedHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostState)(nil)).Elem()
}

type dedicatedHostArgs struct {
	// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
	AutoReplaceOnFailure *bool `pulumi:"autoReplaceOnFailure"`
	// Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
	DedicatedHostGroupId string `pulumi:"dedicatedHostGroupId"`
	// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
	LicenseType *string `pulumi:"licenseType"`
	// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
	PlatformFaultDomain int `pulumi:"platformFaultDomain"`
	// Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DedicatedHost resource.
type DedicatedHostArgs struct {
	// Should the Dedicated Host automatically be replaced in case of a Hardware Failure? Defaults to `true`.
	AutoReplaceOnFailure pulumi.BoolPtrInput
	// Specifies the ID of the Dedicated Host Group where the Dedicated Host should exist. Changing this forces a new resource to be created.
	DedicatedHostGroupId pulumi.StringInput
	// Specifies the software license type that will be applied to the VMs deployed on the Dedicated Host. Possible values are `None`, `Windows_Server_Hybrid` and `Windows_Server_Perpetual`. Defaults to `None`.
	LicenseType pulumi.StringPtrInput
	// Specify the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of this Dedicated Host. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specify the fault domain of the Dedicated Host Group in which to create the Dedicated Host. Changing this forces a new resource to be created.
	PlatformFaultDomain pulumi.IntInput
	// Specify the sku name of the Dedicated Host. Possible values are `DSv3-Type1`, `DSv3-Type2`, `ESv3-Type1`, `ESv3-Type2`,`FSv2-Type2`. Changing this forces a new resource to be created.
	SkuName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DedicatedHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostArgs)(nil)).Elem()
}
