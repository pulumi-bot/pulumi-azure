// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managementresource

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetSubscription(ctx, nil, nil)
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
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
