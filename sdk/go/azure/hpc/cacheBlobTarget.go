// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hpc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Blob Target within a HPC Cache.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/hpc"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi-azuread/sdk/v2/go/azuread"
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
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefix:      pulumi.String("10.0.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleCache, err := hpc.NewCache(ctx, "exampleCache", &hpc.CacheArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			CacheSizeInGb:     pulumi.Int(3072),
// 			SubnetId:          exampleSubnet.ID(),
// 			SkuName:           pulumi.String("Standard_2G"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName: exampleAccount.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := "HPC Cache Resource Provider"
// 		exampleServicePrincipal, err := azuread.LookupServicePrincipal(ctx, &azuread.LookupServicePrincipalArgs{
// 			DisplayName: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authorization.NewAssignment(ctx, "exampleStorageAccountContrib", &authorization.AssignmentArgs{
// 			Scope:              exampleAccount.ID(),
// 			RoleDefinitionName: pulumi.String("Storage Account Contributor"),
// 			PrincipalId:        pulumi.String(exampleServicePrincipal.ObjectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authorization.NewAssignment(ctx, "exampleStorageBlobDataContrib", &authorization.AssignmentArgs{
// 			Scope:              exampleAccount.ID(),
// 			RoleDefinitionName: pulumi.String("Storage Blob Data Contributor"),
// 			PrincipalId:        pulumi.String(exampleServicePrincipal.ObjectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = hpc.NewCacheBlobTarget(ctx, "exampleCacheBlobTarget", &hpc.CacheBlobTargetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			CacheName:          exampleCache.Name,
// 			StorageContainerId: exampleContainer.ResourceManagerId,
// 			NamespacePath:      pulumi.String("/blob_storage"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CacheBlobTarget struct {
	pulumi.CustomResourceState

	// The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
	CacheName pulumi.StringOutput `pulumi:"cacheName"`
	// The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The client-facing file path of the HPC Cache Blob Target.
	NamespacePath pulumi.StringOutput `pulumi:"namespacePath"`
	// The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
	StorageContainerId pulumi.StringOutput `pulumi:"storageContainerId"`
}

// NewCacheBlobTarget registers a new resource with the given unique name, arguments, and options.
func NewCacheBlobTarget(ctx *pulumi.Context,
	name string, args *CacheBlobTargetArgs, opts ...pulumi.ResourceOption) (*CacheBlobTarget, error) {
	if args == nil || args.CacheName == nil {
		return nil, errors.New("missing required argument 'CacheName'")
	}
	if args == nil || args.NamespacePath == nil {
		return nil, errors.New("missing required argument 'NamespacePath'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageContainerId == nil {
		return nil, errors.New("missing required argument 'StorageContainerId'")
	}
	if args == nil {
		args = &CacheBlobTargetArgs{}
	}
	var resource CacheBlobTarget
	err := ctx.RegisterResource("azure:hpc/cacheBlobTarget:CacheBlobTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCacheBlobTarget gets an existing CacheBlobTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCacheBlobTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CacheBlobTargetState, opts ...pulumi.ResourceOption) (*CacheBlobTarget, error) {
	var resource CacheBlobTarget
	err := ctx.ReadResource("azure:hpc/cacheBlobTarget:CacheBlobTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CacheBlobTarget resources.
type cacheBlobTargetState struct {
	// The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
	CacheName *string `pulumi:"cacheName"`
	// The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The client-facing file path of the HPC Cache Blob Target.
	NamespacePath *string `pulumi:"namespacePath"`
	// The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
	StorageContainerId *string `pulumi:"storageContainerId"`
}

type CacheBlobTargetState struct {
	// The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
	CacheName pulumi.StringPtrInput
	// The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The client-facing file path of the HPC Cache Blob Target.
	NamespacePath pulumi.StringPtrInput
	// The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
	StorageContainerId pulumi.StringPtrInput
}

func (CacheBlobTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheBlobTargetState)(nil)).Elem()
}

type cacheBlobTargetArgs struct {
	// The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
	CacheName string `pulumi:"cacheName"`
	// The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The client-facing file path of the HPC Cache Blob Target.
	NamespacePath string `pulumi:"namespacePath"`
	// The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
	StorageContainerId string `pulumi:"storageContainerId"`
}

// The set of arguments for constructing a CacheBlobTarget resource.
type CacheBlobTargetArgs struct {
	// The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
	CacheName pulumi.StringInput
	// The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The client-facing file path of the HPC Cache Blob Target.
	NamespacePath pulumi.StringInput
	// The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
	StorageContainerId pulumi.StringInput
}

func (CacheBlobTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheBlobTargetArgs)(nil)).Elem()
}
