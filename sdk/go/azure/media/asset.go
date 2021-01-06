// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Media Asset.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/media"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("GRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleServiceAccount, err := media.NewServiceAccount(ctx, "exampleServiceAccount", &media.ServiceAccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			StorageAccounts: media.ServiceAccountStorageAccountArray{
// 				&media.ServiceAccountStorageAccountArgs{
// 					Id:        exampleAccount.ID(),
// 					IsPrimary: pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = media.NewAsset(ctx, "exampleAsset", &media.AssetArgs{
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			MediaServicesAccountName: exampleServiceAccount.Name,
// 			Description:              pulumi.String("Asset description"),
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
// Media Assets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:media/asset:Asset example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/assets/asset1
// ```
type Asset struct {
	pulumi.CustomResourceState

	// The alternate ID of the Asset.
	AlternateId pulumi.StringPtrOutput `pulumi:"alternateId"`
	// The name of the asset blob container. Changing this forces a new Media Asset to be created.
	Container pulumi.StringOutput `pulumi:"container"`
	// The Asset description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
	MediaServicesAccountName pulumi.StringOutput `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Media Asset. Changing this forces a new Media Asset to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
}

// NewAsset registers a new resource with the given unique name, arguments, and options.
func NewAsset(ctx *pulumi.Context,
	name string, args *AssetArgs, opts ...pulumi.ResourceOption) (*Asset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MediaServicesAccountName == nil {
		return nil, errors.New("invalid value for required argument 'MediaServicesAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Asset
	err := ctx.RegisterResource("azure:media/asset:Asset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAsset gets an existing Asset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAsset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetState, opts ...pulumi.ResourceOption) (*Asset, error) {
	var resource Asset
	err := ctx.ReadResource("azure:media/asset:Asset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Asset resources.
type assetState struct {
	// The alternate ID of the Asset.
	AlternateId *string `pulumi:"alternateId"`
	// The name of the asset blob container. Changing this forces a new Media Asset to be created.
	Container *string `pulumi:"container"`
	// The Asset description.
	Description *string `pulumi:"description"`
	// Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
	MediaServicesAccountName *string `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Media Asset. Changing this forces a new Media Asset to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
	StorageAccountName *string `pulumi:"storageAccountName"`
}

type AssetState struct {
	// The alternate ID of the Asset.
	AlternateId pulumi.StringPtrInput
	// The name of the asset blob container. Changing this forces a new Media Asset to be created.
	Container pulumi.StringPtrInput
	// The Asset description.
	Description pulumi.StringPtrInput
	// Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
	MediaServicesAccountName pulumi.StringPtrInput
	// The name which should be used for this Media Asset. Changing this forces a new Media Asset to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
	StorageAccountName pulumi.StringPtrInput
}

func (AssetState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetState)(nil)).Elem()
}

type assetArgs struct {
	// The alternate ID of the Asset.
	AlternateId *string `pulumi:"alternateId"`
	// The name of the asset blob container. Changing this forces a new Media Asset to be created.
	Container *string `pulumi:"container"`
	// The Asset description.
	Description *string `pulumi:"description"`
	// Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
	MediaServicesAccountName string `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Media Asset. Changing this forces a new Media Asset to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
	StorageAccountName *string `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Asset resource.
type AssetArgs struct {
	// The alternate ID of the Asset.
	AlternateId pulumi.StringPtrInput
	// The name of the asset blob container. Changing this forces a new Media Asset to be created.
	Container pulumi.StringPtrInput
	// The Asset description.
	Description pulumi.StringPtrInput
	// Specifies the name of the Media Services Account. Changing this forces a new Media Asset to be created.
	MediaServicesAccountName pulumi.StringInput
	// The name which should be used for this Media Asset. Changing this forces a new Media Asset to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Media Asset should exist. Changing this forces a new Media Asset to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the storage account where to store the media asset. Changing this forces a new Media Asset to be created.
	StorageAccountName pulumi.StringPtrInput
}

func (AssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetArgs)(nil)).Elem()
}

type AssetInput interface {
	pulumi.Input

	ToAssetOutput() AssetOutput
	ToAssetOutputWithContext(ctx context.Context) AssetOutput
}

func (*Asset) ElementType() reflect.Type {
	return reflect.TypeOf((*Asset)(nil))
}

func (i *Asset) ToAssetOutput() AssetOutput {
	return i.ToAssetOutputWithContext(context.Background())
}

func (i *Asset) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetOutput)
}

func (i *Asset) ToAssetPtrOutput() AssetPtrOutput {
	return i.ToAssetPtrOutputWithContext(context.Background())
}

func (i *Asset) ToAssetPtrOutputWithContext(ctx context.Context) AssetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetPtrOutput)
}

type AssetPtrInput interface {
	pulumi.Input

	ToAssetPtrOutput() AssetPtrOutput
	ToAssetPtrOutputWithContext(ctx context.Context) AssetPtrOutput
}

type assetPtrType AssetArgs

func (*assetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil))
}

func (i *assetPtrType) ToAssetPtrOutput() AssetPtrOutput {
	return i.ToAssetPtrOutputWithContext(context.Background())
}

func (i *assetPtrType) ToAssetPtrOutputWithContext(ctx context.Context) AssetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetOutput).ToAssetPtrOutput()
}

// AssetArrayInput is an input type that accepts AssetArray and AssetArrayOutput values.
// You can construct a concrete instance of `AssetArrayInput` via:
//
//          AssetArray{ AssetArgs{...} }
type AssetArrayInput interface {
	pulumi.Input

	ToAssetArrayOutput() AssetArrayOutput
	ToAssetArrayOutputWithContext(context.Context) AssetArrayOutput
}

type AssetArray []AssetInput

func (AssetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Asset)(nil))
}

func (i AssetArray) ToAssetArrayOutput() AssetArrayOutput {
	return i.ToAssetArrayOutputWithContext(context.Background())
}

func (i AssetArray) ToAssetArrayOutputWithContext(ctx context.Context) AssetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetArrayOutput)
}

// AssetMapInput is an input type that accepts AssetMap and AssetMapOutput values.
// You can construct a concrete instance of `AssetMapInput` via:
//
//          AssetMap{ "key": AssetArgs{...} }
type AssetMapInput interface {
	pulumi.Input

	ToAssetMapOutput() AssetMapOutput
	ToAssetMapOutputWithContext(context.Context) AssetMapOutput
}

type AssetMap map[string]AssetInput

func (AssetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Asset)(nil))
}

func (i AssetMap) ToAssetMapOutput() AssetMapOutput {
	return i.ToAssetMapOutputWithContext(context.Background())
}

func (i AssetMap) ToAssetMapOutputWithContext(ctx context.Context) AssetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetMapOutput)
}

type AssetOutput struct {
	*pulumi.OutputState
}

func (AssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Asset)(nil))
}

func (o AssetOutput) ToAssetOutput() AssetOutput {
	return o
}

func (o AssetOutput) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return o
}

func (o AssetOutput) ToAssetPtrOutput() AssetPtrOutput {
	return o.ToAssetPtrOutputWithContext(context.Background())
}

func (o AssetOutput) ToAssetPtrOutputWithContext(ctx context.Context) AssetPtrOutput {
	return o.ApplyT(func(v Asset) *Asset {
		return &v
	}).(AssetPtrOutput)
}

type AssetPtrOutput struct {
	*pulumi.OutputState
}

func (AssetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil))
}

func (o AssetPtrOutput) ToAssetPtrOutput() AssetPtrOutput {
	return o
}

func (o AssetPtrOutput) ToAssetPtrOutputWithContext(ctx context.Context) AssetPtrOutput {
	return o
}

type AssetArrayOutput struct{ *pulumi.OutputState }

func (AssetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Asset)(nil))
}

func (o AssetArrayOutput) ToAssetArrayOutput() AssetArrayOutput {
	return o
}

func (o AssetArrayOutput) ToAssetArrayOutputWithContext(ctx context.Context) AssetArrayOutput {
	return o
}

func (o AssetArrayOutput) Index(i pulumi.IntInput) AssetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Asset {
		return vs[0].([]Asset)[vs[1].(int)]
	}).(AssetOutput)
}

type AssetMapOutput struct{ *pulumi.OutputState }

func (AssetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Asset)(nil))
}

func (o AssetMapOutput) ToAssetMapOutput() AssetMapOutput {
	return o
}

func (o AssetMapOutput) ToAssetMapOutputWithContext(ctx context.Context) AssetMapOutput {
	return o
}

func (o AssetMapOutput) MapIndex(k pulumi.StringInput) AssetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Asset {
		return vs[0].(map[string]Asset)[vs[1].(string)]
	}).(AssetOutput)
}

func init() {
	pulumi.RegisterOutputType(AssetOutput{})
	pulumi.RegisterOutputType(AssetPtrOutput{})
	pulumi.RegisterOutputType(AssetArrayOutput{})
	pulumi.RegisterOutputType(AssetMapOutput{})
}
