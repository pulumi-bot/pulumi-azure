// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appconfiguration

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure App Configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appconfiguration"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appconfiguration.NewConfigurationStore(ctx, "appconf", &appconfiguration.ConfigurationStoreArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
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
// App Configurations can be imported using the `resource id`, e.g.
type ConfigurationStore struct {
	pulumi.CustomResourceState

	// The URL of the App Configuration.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// An `identity` block as defined below.
	Identity ConfigurationStoreIdentityPtrOutput `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the App Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `primaryReadKey` block as defined below containing the primary read access key.
	PrimaryReadKeys ConfigurationStorePrimaryReadKeyArrayOutput `pulumi:"primaryReadKeys"`
	// A `primaryWriteKey` block as defined below containing the primary write access key.
	PrimaryWriteKeys ConfigurationStorePrimaryWriteKeyArrayOutput `pulumi:"primaryWriteKeys"`
	// The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `secondaryReadKey` block as defined below containing the secondary read access key.
	SecondaryReadKeys ConfigurationStoreSecondaryReadKeyArrayOutput `pulumi:"secondaryReadKeys"`
	// A `secondaryWriteKey` block as defined below containing the secondary write access key.
	SecondaryWriteKeys ConfigurationStoreSecondaryWriteKeyArrayOutput `pulumi:"secondaryWriteKeys"`
	// The SKU name of the the App Configuration. Possible values are `free` and `standard`.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewConfigurationStore registers a new resource with the given unique name, arguments, and options.
func NewConfigurationStore(ctx *pulumi.Context,
	name string, args *ConfigurationStoreArgs, opts ...pulumi.ResourceOption) (*ConfigurationStore, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ConfigurationStoreArgs{}
	}
	var resource ConfigurationStore
	err := ctx.RegisterResource("azure:appconfiguration/configurationStore:ConfigurationStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationStore gets an existing ConfigurationStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationStoreState, opts ...pulumi.ResourceOption) (*ConfigurationStore, error) {
	var resource ConfigurationStore
	err := ctx.ReadResource("azure:appconfiguration/configurationStore:ConfigurationStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationStore resources.
type configurationStoreState struct {
	// The URL of the App Configuration.
	Endpoint *string `pulumi:"endpoint"`
	// An `identity` block as defined below.
	Identity *ConfigurationStoreIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the App Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `primaryReadKey` block as defined below containing the primary read access key.
	PrimaryReadKeys []ConfigurationStorePrimaryReadKey `pulumi:"primaryReadKeys"`
	// A `primaryWriteKey` block as defined below containing the primary write access key.
	PrimaryWriteKeys []ConfigurationStorePrimaryWriteKey `pulumi:"primaryWriteKeys"`
	// The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `secondaryReadKey` block as defined below containing the secondary read access key.
	SecondaryReadKeys []ConfigurationStoreSecondaryReadKey `pulumi:"secondaryReadKeys"`
	// A `secondaryWriteKey` block as defined below containing the secondary write access key.
	SecondaryWriteKeys []ConfigurationStoreSecondaryWriteKey `pulumi:"secondaryWriteKeys"`
	// The SKU name of the the App Configuration. Possible values are `free` and `standard`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ConfigurationStoreState struct {
	// The URL of the App Configuration.
	Endpoint pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity ConfigurationStoreIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the App Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `primaryReadKey` block as defined below containing the primary read access key.
	PrimaryReadKeys ConfigurationStorePrimaryReadKeyArrayInput
	// A `primaryWriteKey` block as defined below containing the primary write access key.
	PrimaryWriteKeys ConfigurationStorePrimaryWriteKeyArrayInput
	// The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `secondaryReadKey` block as defined below containing the secondary read access key.
	SecondaryReadKeys ConfigurationStoreSecondaryReadKeyArrayInput
	// A `secondaryWriteKey` block as defined below containing the secondary write access key.
	SecondaryWriteKeys ConfigurationStoreSecondaryWriteKeyArrayInput
	// The SKU name of the the App Configuration. Possible values are `free` and `standard`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ConfigurationStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationStoreState)(nil)).Elem()
}

type configurationStoreArgs struct {
	// An `identity` block as defined below.
	Identity *ConfigurationStoreIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the App Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU name of the the App Configuration. Possible values are `free` and `standard`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationStore resource.
type ConfigurationStoreArgs struct {
	// An `identity` block as defined below.
	Identity ConfigurationStoreIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the App Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The SKU name of the the App Configuration. Possible values are `free` and `standard`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ConfigurationStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationStoreArgs)(nil)).Elem()
}
