// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a HPC Cache Access Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/hpc"
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
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
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
// 		_, err = hpc.NewCacheAccessPolicy(ctx, "exampleCacheAccessPolicy", &hpc.CacheAccessPolicyArgs{
// 			HpcCacheId: exampleCache.ID(),
// 			AccessRules: hpc.CacheAccessPolicyAccessRuleArray{
// 				&hpc.CacheAccessPolicyAccessRuleArgs{
// 					Scope:  pulumi.String("default"),
// 					Access: pulumi.String("rw"),
// 				},
// 			},
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
// HPC Cache Access Policys can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:hpc/cacheAccessPolicy:CacheAccessPolicy example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.StorageCache/caches/cache1/cacheAccessPolicies/policy1
// ```
type CacheAccessPolicy struct {
	pulumi.CustomResourceState

	// Up to three `accessRule` blocks as defined below.
	AccessRules CacheAccessPolicyAccessRuleArrayOutput `pulumi:"accessRules"`
	// The ID of the HPC Cache that this HPC Cache Access Policy resides in. Changing this forces a new HPC Cache Access Policy to be created.
	HpcCacheId pulumi.StringOutput `pulumi:"hpcCacheId"`
	// The name which should be used for this HPC Cache Access Policy. Changing this forces a new HPC Cache Access Policy to be created.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewCacheAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewCacheAccessPolicy(ctx *pulumi.Context,
	name string, args *CacheAccessPolicyArgs, opts ...pulumi.ResourceOption) (*CacheAccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessRules == nil {
		return nil, errors.New("invalid value for required argument 'AccessRules'")
	}
	if args.HpcCacheId == nil {
		return nil, errors.New("invalid value for required argument 'HpcCacheId'")
	}
	var resource CacheAccessPolicy
	err := ctx.RegisterResource("azure:hpc/cacheAccessPolicy:CacheAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCacheAccessPolicy gets an existing CacheAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCacheAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CacheAccessPolicyState, opts ...pulumi.ResourceOption) (*CacheAccessPolicy, error) {
	var resource CacheAccessPolicy
	err := ctx.ReadResource("azure:hpc/cacheAccessPolicy:CacheAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CacheAccessPolicy resources.
type cacheAccessPolicyState struct {
	// Up to three `accessRule` blocks as defined below.
	AccessRules []CacheAccessPolicyAccessRule `pulumi:"accessRules"`
	// The ID of the HPC Cache that this HPC Cache Access Policy resides in. Changing this forces a new HPC Cache Access Policy to be created.
	HpcCacheId *string `pulumi:"hpcCacheId"`
	// The name which should be used for this HPC Cache Access Policy. Changing this forces a new HPC Cache Access Policy to be created.
	Name *string `pulumi:"name"`
}

type CacheAccessPolicyState struct {
	// Up to three `accessRule` blocks as defined below.
	AccessRules CacheAccessPolicyAccessRuleArrayInput
	// The ID of the HPC Cache that this HPC Cache Access Policy resides in. Changing this forces a new HPC Cache Access Policy to be created.
	HpcCacheId pulumi.StringPtrInput
	// The name which should be used for this HPC Cache Access Policy. Changing this forces a new HPC Cache Access Policy to be created.
	Name pulumi.StringPtrInput
}

func (CacheAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheAccessPolicyState)(nil)).Elem()
}

type cacheAccessPolicyArgs struct {
	// Up to three `accessRule` blocks as defined below.
	AccessRules []CacheAccessPolicyAccessRule `pulumi:"accessRules"`
	// The ID of the HPC Cache that this HPC Cache Access Policy resides in. Changing this forces a new HPC Cache Access Policy to be created.
	HpcCacheId string `pulumi:"hpcCacheId"`
	// The name which should be used for this HPC Cache Access Policy. Changing this forces a new HPC Cache Access Policy to be created.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a CacheAccessPolicy resource.
type CacheAccessPolicyArgs struct {
	// Up to three `accessRule` blocks as defined below.
	AccessRules CacheAccessPolicyAccessRuleArrayInput
	// The ID of the HPC Cache that this HPC Cache Access Policy resides in. Changing this forces a new HPC Cache Access Policy to be created.
	HpcCacheId pulumi.StringInput
	// The name which should be used for this HPC Cache Access Policy. Changing this forces a new HPC Cache Access Policy to be created.
	Name pulumi.StringPtrInput
}

func (CacheAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheAccessPolicyArgs)(nil)).Elem()
}

type CacheAccessPolicyInput interface {
	pulumi.Input

	ToCacheAccessPolicyOutput() CacheAccessPolicyOutput
	ToCacheAccessPolicyOutputWithContext(ctx context.Context) CacheAccessPolicyOutput
}

func (*CacheAccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheAccessPolicy)(nil))
}

func (i *CacheAccessPolicy) ToCacheAccessPolicyOutput() CacheAccessPolicyOutput {
	return i.ToCacheAccessPolicyOutputWithContext(context.Background())
}

func (i *CacheAccessPolicy) ToCacheAccessPolicyOutputWithContext(ctx context.Context) CacheAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheAccessPolicyOutput)
}

func (i *CacheAccessPolicy) ToCacheAccessPolicyPtrOutput() CacheAccessPolicyPtrOutput {
	return i.ToCacheAccessPolicyPtrOutputWithContext(context.Background())
}

func (i *CacheAccessPolicy) ToCacheAccessPolicyPtrOutputWithContext(ctx context.Context) CacheAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheAccessPolicyPtrOutput)
}

type CacheAccessPolicyPtrInput interface {
	pulumi.Input

	ToCacheAccessPolicyPtrOutput() CacheAccessPolicyPtrOutput
	ToCacheAccessPolicyPtrOutputWithContext(ctx context.Context) CacheAccessPolicyPtrOutput
}

type cacheAccessPolicyPtrType CacheAccessPolicyArgs

func (*cacheAccessPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheAccessPolicy)(nil))
}

func (i *cacheAccessPolicyPtrType) ToCacheAccessPolicyPtrOutput() CacheAccessPolicyPtrOutput {
	return i.ToCacheAccessPolicyPtrOutputWithContext(context.Background())
}

func (i *cacheAccessPolicyPtrType) ToCacheAccessPolicyPtrOutputWithContext(ctx context.Context) CacheAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheAccessPolicyPtrOutput)
}

// CacheAccessPolicyArrayInput is an input type that accepts CacheAccessPolicyArray and CacheAccessPolicyArrayOutput values.
// You can construct a concrete instance of `CacheAccessPolicyArrayInput` via:
//
//          CacheAccessPolicyArray{ CacheAccessPolicyArgs{...} }
type CacheAccessPolicyArrayInput interface {
	pulumi.Input

	ToCacheAccessPolicyArrayOutput() CacheAccessPolicyArrayOutput
	ToCacheAccessPolicyArrayOutputWithContext(context.Context) CacheAccessPolicyArrayOutput
}

type CacheAccessPolicyArray []CacheAccessPolicyInput

func (CacheAccessPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CacheAccessPolicy)(nil))
}

func (i CacheAccessPolicyArray) ToCacheAccessPolicyArrayOutput() CacheAccessPolicyArrayOutput {
	return i.ToCacheAccessPolicyArrayOutputWithContext(context.Background())
}

func (i CacheAccessPolicyArray) ToCacheAccessPolicyArrayOutputWithContext(ctx context.Context) CacheAccessPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheAccessPolicyArrayOutput)
}

// CacheAccessPolicyMapInput is an input type that accepts CacheAccessPolicyMap and CacheAccessPolicyMapOutput values.
// You can construct a concrete instance of `CacheAccessPolicyMapInput` via:
//
//          CacheAccessPolicyMap{ "key": CacheAccessPolicyArgs{...} }
type CacheAccessPolicyMapInput interface {
	pulumi.Input

	ToCacheAccessPolicyMapOutput() CacheAccessPolicyMapOutput
	ToCacheAccessPolicyMapOutputWithContext(context.Context) CacheAccessPolicyMapOutput
}

type CacheAccessPolicyMap map[string]CacheAccessPolicyInput

func (CacheAccessPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CacheAccessPolicy)(nil))
}

func (i CacheAccessPolicyMap) ToCacheAccessPolicyMapOutput() CacheAccessPolicyMapOutput {
	return i.ToCacheAccessPolicyMapOutputWithContext(context.Background())
}

func (i CacheAccessPolicyMap) ToCacheAccessPolicyMapOutputWithContext(ctx context.Context) CacheAccessPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheAccessPolicyMapOutput)
}

type CacheAccessPolicyOutput struct {
	*pulumi.OutputState
}

func (CacheAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheAccessPolicy)(nil))
}

func (o CacheAccessPolicyOutput) ToCacheAccessPolicyOutput() CacheAccessPolicyOutput {
	return o
}

func (o CacheAccessPolicyOutput) ToCacheAccessPolicyOutputWithContext(ctx context.Context) CacheAccessPolicyOutput {
	return o
}

func (o CacheAccessPolicyOutput) ToCacheAccessPolicyPtrOutput() CacheAccessPolicyPtrOutput {
	return o.ToCacheAccessPolicyPtrOutputWithContext(context.Background())
}

func (o CacheAccessPolicyOutput) ToCacheAccessPolicyPtrOutputWithContext(ctx context.Context) CacheAccessPolicyPtrOutput {
	return o.ApplyT(func(v CacheAccessPolicy) *CacheAccessPolicy {
		return &v
	}).(CacheAccessPolicyPtrOutput)
}

type CacheAccessPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (CacheAccessPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheAccessPolicy)(nil))
}

func (o CacheAccessPolicyPtrOutput) ToCacheAccessPolicyPtrOutput() CacheAccessPolicyPtrOutput {
	return o
}

func (o CacheAccessPolicyPtrOutput) ToCacheAccessPolicyPtrOutputWithContext(ctx context.Context) CacheAccessPolicyPtrOutput {
	return o
}

type CacheAccessPolicyArrayOutput struct{ *pulumi.OutputState }

func (CacheAccessPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CacheAccessPolicy)(nil))
}

func (o CacheAccessPolicyArrayOutput) ToCacheAccessPolicyArrayOutput() CacheAccessPolicyArrayOutput {
	return o
}

func (o CacheAccessPolicyArrayOutput) ToCacheAccessPolicyArrayOutputWithContext(ctx context.Context) CacheAccessPolicyArrayOutput {
	return o
}

func (o CacheAccessPolicyArrayOutput) Index(i pulumi.IntInput) CacheAccessPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CacheAccessPolicy {
		return vs[0].([]CacheAccessPolicy)[vs[1].(int)]
	}).(CacheAccessPolicyOutput)
}

type CacheAccessPolicyMapOutput struct{ *pulumi.OutputState }

func (CacheAccessPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CacheAccessPolicy)(nil))
}

func (o CacheAccessPolicyMapOutput) ToCacheAccessPolicyMapOutput() CacheAccessPolicyMapOutput {
	return o
}

func (o CacheAccessPolicyMapOutput) ToCacheAccessPolicyMapOutputWithContext(ctx context.Context) CacheAccessPolicyMapOutput {
	return o
}

func (o CacheAccessPolicyMapOutput) MapIndex(k pulumi.StringInput) CacheAccessPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CacheAccessPolicy {
		return vs[0].(map[string]CacheAccessPolicy)[vs[1].(string)]
	}).(CacheAccessPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(CacheAccessPolicyOutput{})
	pulumi.RegisterOutputType(CacheAccessPolicyPtrOutput{})
	pulumi.RegisterOutputType(CacheAccessPolicyArrayOutput{})
	pulumi.RegisterOutputType(CacheAccessPolicyMapOutput{})
}
