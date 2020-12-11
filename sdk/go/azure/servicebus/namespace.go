// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a ServiceBus Namespace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/servicebus"
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
// 		_, err = servicebus.NewNamespace(ctx, "exampleNamespace", &servicebus.NamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 			Tags: pulumi.StringMap{
// 				"source": pulumi.String("example"),
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
// Service Bus Namespace can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:servicebus/namespace:Namespace example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1
// ```
type Namespace struct {
	pulumi.CustomResourceState

	// Specifies the capacity. When `sku` is `Premium`, capacity can be `1`, `2`, `4` or `8`. When `sku` is `Basic` or `Standard`, capacity can be `0` only.
	Capacity pulumi.IntPtrOutput `pulumi:"capacity"`
	// The primary connection string for the authorization
	// rule `RootManageSharedAccessKey`.
	DefaultPrimaryConnectionString pulumi.StringOutput `pulumi:"defaultPrimaryConnectionString"`
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultPrimaryKey pulumi.StringOutput `pulumi:"defaultPrimaryKey"`
	// The secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryConnectionString pulumi.StringOutput `pulumi:"defaultSecondaryConnectionString"`
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryKey pulumi.StringOutput `pulumi:"defaultSecondaryKey"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the ServiceBus Namespace resource . Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to
	// create the namespace.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Defines which tier to use. Options are basic, standard or premium. Changing this forces a new resource to be created.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Whether or not this resource is zone redundant. `sku` needs to be `Premium`. Defaults to `false`.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:eventhub/namespace:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure:servicebus/namespace:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure:servicebus/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// Specifies the capacity. When `sku` is `Premium`, capacity can be `1`, `2`, `4` or `8`. When `sku` is `Basic` or `Standard`, capacity can be `0` only.
	Capacity *int `pulumi:"capacity"`
	// The primary connection string for the authorization
	// rule `RootManageSharedAccessKey`.
	DefaultPrimaryConnectionString *string `pulumi:"defaultPrimaryConnectionString"`
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultPrimaryKey *string `pulumi:"defaultPrimaryKey"`
	// The secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryConnectionString *string `pulumi:"defaultSecondaryConnectionString"`
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryKey *string `pulumi:"defaultSecondaryKey"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the ServiceBus Namespace resource . Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the namespace.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Defines which tier to use. Options are basic, standard or premium. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this resource is zone redundant. `sku` needs to be `Premium`. Defaults to `false`.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type NamespaceState struct {
	// Specifies the capacity. When `sku` is `Premium`, capacity can be `1`, `2`, `4` or `8`. When `sku` is `Basic` or `Standard`, capacity can be `0` only.
	Capacity pulumi.IntPtrInput
	// The primary connection string for the authorization
	// rule `RootManageSharedAccessKey`.
	DefaultPrimaryConnectionString pulumi.StringPtrInput
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultPrimaryKey pulumi.StringPtrInput
	// The secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryConnectionString pulumi.StringPtrInput
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryKey pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the ServiceBus Namespace resource . Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the namespace.
	ResourceGroupName pulumi.StringPtrInput
	// Defines which tier to use. Options are basic, standard or premium. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Whether or not this resource is zone redundant. `sku` needs to be `Premium`. Defaults to `false`.
	ZoneRedundant pulumi.BoolPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Specifies the capacity. When `sku` is `Premium`, capacity can be `1`, `2`, `4` or `8`. When `sku` is `Basic` or `Standard`, capacity can be `0` only.
	Capacity *int `pulumi:"capacity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the ServiceBus Namespace resource . Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the namespace.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Defines which tier to use. Options are basic, standard or premium. Changing this forces a new resource to be created.
	Sku string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this resource is zone redundant. `sku` needs to be `Premium`. Defaults to `false`.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Specifies the capacity. When `sku` is `Premium`, capacity can be `1`, `2`, `4` or `8`. When `sku` is `Basic` or `Standard`, capacity can be `0` only.
	Capacity pulumi.IntPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the ServiceBus Namespace resource . Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the namespace.
	ResourceGroupName pulumi.StringInput
	// Defines which tier to use. Options are basic, standard or premium. Changing this forces a new resource to be created.
	Sku pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Whether or not this resource is zone redundant. `sku` needs to be `Premium`. Defaults to `false`.
	ZoneRedundant pulumi.BoolPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

type NamespacePtrInput interface {
	pulumi.Input

	ToNamespacePtrOutput() NamespacePtrOutput
	ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput
}

func (Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil)).Elem()
}

func (i Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

func (i Namespace) ToNamespacePtrOutput() NamespacePtrOutput {
	return i.ToNamespacePtrOutputWithContext(context.Background())
}

func (i Namespace) ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePtrOutput)
}

type NamespaceOutput struct {
	*pulumi.OutputState
}

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceOutput)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

type NamespacePtrOutput struct {
	*pulumi.OutputState
}

func (NamespacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespacePtrOutput) ToNamespacePtrOutput() NamespacePtrOutput {
	return o
}

func (o NamespacePtrOutput) ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
	pulumi.RegisterOutputType(NamespacePtrOutput{})
}
