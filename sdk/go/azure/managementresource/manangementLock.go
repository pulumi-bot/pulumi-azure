// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managementresource

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Management Lock which is scoped to a Subscription, Resource Group or Resource.
//
// ## Example Usage
// ### Subscription Level Lock)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.LookupSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = management.NewLock(ctx, "subscription_level", &management.LockArgs{
// 			Scope:     pulumi.String(current.Id),
// 			LockLevel: pulumi.String("CanNotDelete"),
// 			Notes:     pulumi.String("Items can't be deleted in this subscription!"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Example Usage (Resource Group Level Lock)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = management.NewLock(ctx, "resource_group_level", &management.LockArgs{
// 			Scope:     example.ID(),
// 			LockLevel: pulumi.String("ReadOnly"),
// 			Notes:     pulumi.String("This Resource Group is Read-Only"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Resource Level Lock)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/management"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:             exampleResourceGroup.Location,
// 			ResourceGroupName:    exampleResourceGroup.Name,
// 			AllocationMethod:     pulumi.String("Static"),
// 			IdleTimeoutInMinutes: pulumi.Int(30),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = management.NewLock(ctx, "public_ip", &management.LockArgs{
// 			Scope:     examplePublicIp.ID(),
// 			LockLevel: pulumi.String("CanNotDelete"),
// 			Notes:     pulumi.String("Locked because it's needed by a third-party"),
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
// Management Locks can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:managementresource/manangementLock:ManangementLock lock1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Authorization/locks/lock1
// ```
//
// Deprecated: azure.managementresource.ManangementLock has been deprecated in favor of azure.management.Lock
type ManangementLock struct {
	pulumi.CustomResourceState

	// Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel pulumi.StringOutput `pulumi:"lockLevel"`
	// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
	Scope pulumi.StringOutput `pulumi:"scope"`
}

// NewManangementLock registers a new resource with the given unique name, arguments, and options.
func NewManangementLock(ctx *pulumi.Context,
	name string, args *ManangementLockArgs, opts ...pulumi.ResourceOption) (*ManangementLock, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LockLevel == nil {
		return nil, errors.New("invalid value for required argument 'LockLevel'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource ManangementLock
	err := ctx.RegisterResource("azure:managementresource/manangementLock:ManangementLock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManangementLock gets an existing ManangementLock resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManangementLock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManangementLockState, opts ...pulumi.ResourceOption) (*ManangementLock, error) {
	var resource ManangementLock
	err := ctx.ReadResource("azure:managementresource/manangementLock:ManangementLock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManangementLock resources.
type manangementLockState struct {
	// Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel *string `pulumi:"lockLevel"`
	// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
	Notes *string `pulumi:"notes"`
	// Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
	Scope *string `pulumi:"scope"`
}

type ManangementLockState struct {
	// Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel pulumi.StringPtrInput
	// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
	Notes pulumi.StringPtrInput
	// Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
	Scope pulumi.StringPtrInput
}

func (ManangementLockState) ElementType() reflect.Type {
	return reflect.TypeOf((*manangementLockState)(nil)).Elem()
}

type manangementLockArgs struct {
	// Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel string `pulumi:"lockLevel"`
	// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
	Notes *string `pulumi:"notes"`
	// Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a ManangementLock resource.
type ManangementLockArgs struct {
	// Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel pulumi.StringInput
	// Specifies the name of the Management Lock. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
	Notes pulumi.StringPtrInput
	// Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
	Scope pulumi.StringInput
}

func (ManangementLockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*manangementLockArgs)(nil)).Elem()
}

type ManangementLockInput interface {
	pulumi.Input

	ToManangementLockOutput() ManangementLockOutput
	ToManangementLockOutputWithContext(ctx context.Context) ManangementLockOutput
}

func (*ManangementLock) ElementType() reflect.Type {
	return reflect.TypeOf((*ManangementLock)(nil))
}

func (i *ManangementLock) ToManangementLockOutput() ManangementLockOutput {
	return i.ToManangementLockOutputWithContext(context.Background())
}

func (i *ManangementLock) ToManangementLockOutputWithContext(ctx context.Context) ManangementLockOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManangementLockOutput)
}

func (i *ManangementLock) ToManangementLockPtrOutput() ManangementLockPtrOutput {
	return i.ToManangementLockPtrOutputWithContext(context.Background())
}

func (i *ManangementLock) ToManangementLockPtrOutputWithContext(ctx context.Context) ManangementLockPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManangementLockPtrOutput)
}

type ManangementLockPtrInput interface {
	pulumi.Input

	ToManangementLockPtrOutput() ManangementLockPtrOutput
	ToManangementLockPtrOutputWithContext(ctx context.Context) ManangementLockPtrOutput
}

type manangementLockPtrType ManangementLockArgs

func (*manangementLockPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManangementLock)(nil))
}

func (i *manangementLockPtrType) ToManangementLockPtrOutput() ManangementLockPtrOutput {
	return i.ToManangementLockPtrOutputWithContext(context.Background())
}

func (i *manangementLockPtrType) ToManangementLockPtrOutputWithContext(ctx context.Context) ManangementLockPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManangementLockPtrOutput)
}

// ManangementLockArrayInput is an input type that accepts ManangementLockArray and ManangementLockArrayOutput values.
// You can construct a concrete instance of `ManangementLockArrayInput` via:
//
//          ManangementLockArray{ ManangementLockArgs{...} }
type ManangementLockArrayInput interface {
	pulumi.Input

	ToManangementLockArrayOutput() ManangementLockArrayOutput
	ToManangementLockArrayOutputWithContext(context.Context) ManangementLockArrayOutput
}

type ManangementLockArray []ManangementLockInput

func (ManangementLockArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ManangementLock)(nil))
}

func (i ManangementLockArray) ToManangementLockArrayOutput() ManangementLockArrayOutput {
	return i.ToManangementLockArrayOutputWithContext(context.Background())
}

func (i ManangementLockArray) ToManangementLockArrayOutputWithContext(ctx context.Context) ManangementLockArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManangementLockArrayOutput)
}

// ManangementLockMapInput is an input type that accepts ManangementLockMap and ManangementLockMapOutput values.
// You can construct a concrete instance of `ManangementLockMapInput` via:
//
//          ManangementLockMap{ "key": ManangementLockArgs{...} }
type ManangementLockMapInput interface {
	pulumi.Input

	ToManangementLockMapOutput() ManangementLockMapOutput
	ToManangementLockMapOutputWithContext(context.Context) ManangementLockMapOutput
}

type ManangementLockMap map[string]ManangementLockInput

func (ManangementLockMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ManangementLock)(nil))
}

func (i ManangementLockMap) ToManangementLockMapOutput() ManangementLockMapOutput {
	return i.ToManangementLockMapOutputWithContext(context.Background())
}

func (i ManangementLockMap) ToManangementLockMapOutputWithContext(ctx context.Context) ManangementLockMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManangementLockMapOutput)
}

type ManangementLockOutput struct {
	*pulumi.OutputState
}

func (ManangementLockOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManangementLock)(nil))
}

func (o ManangementLockOutput) ToManangementLockOutput() ManangementLockOutput {
	return o
}

func (o ManangementLockOutput) ToManangementLockOutputWithContext(ctx context.Context) ManangementLockOutput {
	return o
}

func (o ManangementLockOutput) ToManangementLockPtrOutput() ManangementLockPtrOutput {
	return o.ToManangementLockPtrOutputWithContext(context.Background())
}

func (o ManangementLockOutput) ToManangementLockPtrOutputWithContext(ctx context.Context) ManangementLockPtrOutput {
	return o.ApplyT(func(v ManangementLock) *ManangementLock {
		return &v
	}).(ManangementLockPtrOutput)
}

type ManangementLockPtrOutput struct {
	*pulumi.OutputState
}

func (ManangementLockPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManangementLock)(nil))
}

func (o ManangementLockPtrOutput) ToManangementLockPtrOutput() ManangementLockPtrOutput {
	return o
}

func (o ManangementLockPtrOutput) ToManangementLockPtrOutputWithContext(ctx context.Context) ManangementLockPtrOutput {
	return o
}

type ManangementLockArrayOutput struct{ *pulumi.OutputState }

func (ManangementLockArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManangementLock)(nil))
}

func (o ManangementLockArrayOutput) ToManangementLockArrayOutput() ManangementLockArrayOutput {
	return o
}

func (o ManangementLockArrayOutput) ToManangementLockArrayOutputWithContext(ctx context.Context) ManangementLockArrayOutput {
	return o
}

func (o ManangementLockArrayOutput) Index(i pulumi.IntInput) ManangementLockOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManangementLock {
		return vs[0].([]ManangementLock)[vs[1].(int)]
	}).(ManangementLockOutput)
}

type ManangementLockMapOutput struct{ *pulumi.OutputState }

func (ManangementLockMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManangementLock)(nil))
}

func (o ManangementLockMapOutput) ToManangementLockMapOutput() ManangementLockMapOutput {
	return o
}

func (o ManangementLockMapOutput) ToManangementLockMapOutputWithContext(ctx context.Context) ManangementLockMapOutput {
	return o
}

func (o ManangementLockMapOutput) MapIndex(k pulumi.StringInput) ManangementLockOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManangementLock {
		return vs[0].(map[string]ManangementLock)[vs[1].(string)]
	}).(ManangementLockOutput)
}

func init() {
	pulumi.RegisterOutputType(ManangementLockOutput{})
	pulumi.RegisterOutputType(ManangementLockPtrOutput{})
	pulumi.RegisterOutputType(ManangementLockArrayOutput{})
	pulumi.RegisterOutputType(ManangementLockMapOutput{})
}
