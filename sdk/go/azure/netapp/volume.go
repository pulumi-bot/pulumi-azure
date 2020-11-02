// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a NetApp Volume.
//
// ## NetApp Volume Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/netapp"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 			Delegations: network.SubnetDelegationArray{
// 				&network.SubnetDelegationArgs{
// 					Name: pulumi.String("netapp"),
// 					ServiceDelegation: &network.SubnetDelegationServiceDelegationArgs{
// 						Name: pulumi.String("Microsoft.Netapp/volumes"),
// 						Actions: pulumi.StringArray{
// 							pulumi.String("Microsoft.Network/networkinterfaces/*"),
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/join/action"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := netapp.NewAccount(ctx, "exampleAccount", &netapp.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePool, err := netapp.NewPool(ctx, "examplePool", &netapp.PoolArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AccountName:       exampleAccount.Name,
// 			ServiceLevel:      pulumi.String("Premium"),
// 			SizeInTb:          pulumi.Int(4),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = netapp.NewVolume(ctx, "exampleVolume", &netapp.VolumeArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AccountName:       exampleAccount.Name,
// 			PoolName:          examplePool.Name,
// 			VolumePath:        pulumi.String("my-unique-file-path"),
// 			ServiceLevel:      pulumi.String("Premium"),
// 			SubnetId:          exampleSubnet.ID(),
// 			Protocols: pulumi.StringArray{
// 				pulumi.String("NFSv4.1"),
// 			},
// 			StorageQuotaInGb: pulumi.Int(100),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// NetApp Volumes can be imported using the `resource id`, e.g.
type Volume struct {
	pulumi.CustomResourceState

	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// One or more `exportPolicyRule` block defined below.
	ExportPolicyRules VolumeExportPolicyRuleArrayOutput `pulumi:"exportPolicyRules"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A list of IPv4 Addresses which should be used to mount the volume.
	MountIpAddresses pulumi.StringArrayOutput `pulumi:"mountIpAddresses"`
	// The name of the NetApp Volume. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	PoolName pulumi.StringOutput `pulumi:"poolName"`
	// The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
	Protocols pulumi.StringArrayOutput `pulumi:"protocols"`
	// The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
	ServiceLevel pulumi.StringOutput `pulumi:"serviceLevel"`
	// The maximum Storage Quota allowed for a file system in Gigabytes.
	StorageQuotaInGb pulumi.IntOutput `pulumi:"storageQuotaInGb"`
	// The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
	VolumePath pulumi.StringOutput `pulumi:"volumePath"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.PoolName == nil {
		return nil, errors.New("missing required argument 'PoolName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceLevel == nil {
		return nil, errors.New("missing required argument 'ServiceLevel'")
	}
	if args == nil || args.StorageQuotaInGb == nil {
		return nil, errors.New("missing required argument 'StorageQuotaInGb'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	if args == nil || args.VolumePath == nil {
		return nil, errors.New("missing required argument 'VolumePath'")
	}
	if args == nil {
		args = &VolumeArgs{}
	}
	var resource Volume
	err := ctx.RegisterResource("azure:netapp/volume:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure:netapp/volume:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// One or more `exportPolicyRule` block defined below.
	ExportPolicyRules []VolumeExportPolicyRule `pulumi:"exportPolicyRules"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A list of IPv4 Addresses which should be used to mount the volume.
	MountIpAddresses []string `pulumi:"mountIpAddresses"`
	// The name of the NetApp Volume. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	PoolName *string `pulumi:"poolName"`
	// The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
	Protocols []string `pulumi:"protocols"`
	// The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
	ServiceLevel *string `pulumi:"serviceLevel"`
	// The maximum Storage Quota allowed for a file system in Gigabytes.
	StorageQuotaInGb *int `pulumi:"storageQuotaInGb"`
	// The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
	SubnetId *string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
	VolumePath *string `pulumi:"volumePath"`
}

type VolumeState struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// One or more `exportPolicyRule` block defined below.
	ExportPolicyRules VolumeExportPolicyRuleArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A list of IPv4 Addresses which should be used to mount the volume.
	MountIpAddresses pulumi.StringArrayInput
	// The name of the NetApp Volume. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	PoolName pulumi.StringPtrInput
	// The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
	Protocols pulumi.StringArrayInput
	// The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
	ServiceLevel pulumi.StringPtrInput
	// The maximum Storage Quota allowed for a file system in Gigabytes.
	StorageQuotaInGb pulumi.IntPtrInput
	// The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
	VolumePath pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// One or more `exportPolicyRule` block defined below.
	ExportPolicyRules []VolumeExportPolicyRule `pulumi:"exportPolicyRules"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the NetApp Volume. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	PoolName string `pulumi:"poolName"`
	// The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
	Protocols []string `pulumi:"protocols"`
	// The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
	ServiceLevel string `pulumi:"serviceLevel"`
	// The maximum Storage Quota allowed for a file system in Gigabytes.
	StorageQuotaInGb int `pulumi:"storageQuotaInGb"`
	// The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
	SubnetId string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
	VolumePath string `pulumi:"volumePath"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// One or more `exportPolicyRule` block defined below.
	ExportPolicyRules VolumeExportPolicyRuleArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the NetApp Volume. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	PoolName pulumi.StringInput
	// The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
	Protocols pulumi.StringArrayInput
	// The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
	ServiceLevel pulumi.StringInput
	// The maximum Storage Quota allowed for a file system in Gigabytes.
	StorageQuotaInGb pulumi.IntInput
	// The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
	SubnetId pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
	VolumePath pulumi.StringInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}
