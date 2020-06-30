// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Data Share Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datashare"
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
// 		_, err = datashare.NewAccount(ctx, "exampleAccount", &datashare.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Identity: &datashare.AccountIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Account struct {
	pulumi.CustomResourceState

	// An `identity` block as defined below.
	Identity AccountIdentityOutput `pulumi:"identity"`
	// The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Data Share Account.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil || args.Identity == nil {
		return nil, errors.New("missing required argument 'Identity'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AccountArgs{}
	}
	var resource Account
	err := ctx.RegisterResource("azure:datashare/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure:datashare/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// An `identity` block as defined below.
	Identity *AccountIdentity `pulumi:"identity"`
	// The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Data Share Account.
	Tags map[string]string `pulumi:"tags"`
}

type AccountState struct {
	// An `identity` block as defined below.
	Identity AccountIdentityPtrInput
	// The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Data Share Account.
	Tags pulumi.StringMapInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// An `identity` block as defined below.
	Identity AccountIdentity `pulumi:"identity"`
	// The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Data Share Account.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// An `identity` block as defined below.
	Identity AccountIdentityInput
	// The Azure Region where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Data Share Account. Changing this forces a new Data Share Account to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Data Share Account should exist. Changing this forces a new Data Share Account to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the Data Share Account.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}
