// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datalake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Data Lake Store.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datalake"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datalake.NewStore(ctx, "exampleStore", &datalake.StoreArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			EncryptionState:   pulumi.String("Enabled"),
// 			EncryptionType:    pulumi.String("ServiceManaged"),
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
// Data Lake Store's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datalake/store:Store example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DataLakeStore/accounts/mydatalakeaccount
// ```
type Store struct {
	pulumi.CustomResourceState

	// Is Encryption enabled on this Data Lake Store Account? Possible values are `Enabled` or `Disabled`. Defaults to `Enabled`.
	EncryptionState pulumi.StringPtrOutput `pulumi:"encryptionState"`
	// The Encryption Type used for this Data Lake Store Account. Currently can be set to `ServiceManaged` when `encryptionState` is `Enabled` - and must be a blank string when it's Disabled.
	EncryptionType pulumi.StringOutput `pulumi:"encryptionType"`
	// The Endpoint for the Data Lake Store.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// are Azure Service IP's allowed through the firewall? Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallAllowAzureIps pulumi.StringPtrOutput `pulumi:"firewallAllowAzureIps"`
	// the state of the Firewall. Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallState pulumi.StringPtrOutput `pulumi:"firewallState"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The monthly commitment tier for Data Lake Store. Accepted values are `Consumption`, `Commitment_1TB`, `Commitment_10TB`, `Commitment_100TB`, `Commitment_500TB`, `Commitment_1PB` or `Commitment_5PB`.
	Tier pulumi.StringPtrOutput `pulumi:"tier"`
}

// NewStore registers a new resource with the given unique name, arguments, and options.
func NewStore(ctx *pulumi.Context,
	name string, args *StoreArgs, opts ...pulumi.ResourceOption) (*Store, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Store
	err := ctx.RegisterResource("azure:datalake/store:Store", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStore gets an existing Store resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StoreState, opts ...pulumi.ResourceOption) (*Store, error) {
	var resource Store
	err := ctx.ReadResource("azure:datalake/store:Store", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Store resources.
type storeState struct {
	// Is Encryption enabled on this Data Lake Store Account? Possible values are `Enabled` or `Disabled`. Defaults to `Enabled`.
	EncryptionState *string `pulumi:"encryptionState"`
	// The Encryption Type used for this Data Lake Store Account. Currently can be set to `ServiceManaged` when `encryptionState` is `Enabled` - and must be a blank string when it's Disabled.
	EncryptionType *string `pulumi:"encryptionType"`
	// The Endpoint for the Data Lake Store.
	Endpoint *string `pulumi:"endpoint"`
	// are Azure Service IP's allowed through the firewall? Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallAllowAzureIps *string `pulumi:"firewallAllowAzureIps"`
	// the state of the Firewall. Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallState *string `pulumi:"firewallState"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The monthly commitment tier for Data Lake Store. Accepted values are `Consumption`, `Commitment_1TB`, `Commitment_10TB`, `Commitment_100TB`, `Commitment_500TB`, `Commitment_1PB` or `Commitment_5PB`.
	Tier *string `pulumi:"tier"`
}

type StoreState struct {
	// Is Encryption enabled on this Data Lake Store Account? Possible values are `Enabled` or `Disabled`. Defaults to `Enabled`.
	EncryptionState pulumi.StringPtrInput
	// The Encryption Type used for this Data Lake Store Account. Currently can be set to `ServiceManaged` when `encryptionState` is `Enabled` - and must be a blank string when it's Disabled.
	EncryptionType pulumi.StringPtrInput
	// The Endpoint for the Data Lake Store.
	Endpoint pulumi.StringPtrInput
	// are Azure Service IP's allowed through the firewall? Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallAllowAzureIps pulumi.StringPtrInput
	// the state of the Firewall. Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallState pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The monthly commitment tier for Data Lake Store. Accepted values are `Consumption`, `Commitment_1TB`, `Commitment_10TB`, `Commitment_100TB`, `Commitment_500TB`, `Commitment_1PB` or `Commitment_5PB`.
	Tier pulumi.StringPtrInput
}

func (StoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*storeState)(nil)).Elem()
}

type storeArgs struct {
	// Is Encryption enabled on this Data Lake Store Account? Possible values are `Enabled` or `Disabled`. Defaults to `Enabled`.
	EncryptionState *string `pulumi:"encryptionState"`
	// The Encryption Type used for this Data Lake Store Account. Currently can be set to `ServiceManaged` when `encryptionState` is `Enabled` - and must be a blank string when it's Disabled.
	EncryptionType *string `pulumi:"encryptionType"`
	// are Azure Service IP's allowed through the firewall? Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallAllowAzureIps *string `pulumi:"firewallAllowAzureIps"`
	// the state of the Firewall. Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallState *string `pulumi:"firewallState"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The monthly commitment tier for Data Lake Store. Accepted values are `Consumption`, `Commitment_1TB`, `Commitment_10TB`, `Commitment_100TB`, `Commitment_500TB`, `Commitment_1PB` or `Commitment_5PB`.
	Tier *string `pulumi:"tier"`
}

// The set of arguments for constructing a Store resource.
type StoreArgs struct {
	// Is Encryption enabled on this Data Lake Store Account? Possible values are `Enabled` or `Disabled`. Defaults to `Enabled`.
	EncryptionState pulumi.StringPtrInput
	// The Encryption Type used for this Data Lake Store Account. Currently can be set to `ServiceManaged` when `encryptionState` is `Enabled` - and must be a blank string when it's Disabled.
	EncryptionType pulumi.StringPtrInput
	// are Azure Service IP's allowed through the firewall? Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallAllowAzureIps pulumi.StringPtrInput
	// the state of the Firewall. Possible values are `Enabled` and `Disabled`. Defaults to `Enabled.`
	FirewallState pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Lake Store.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The monthly commitment tier for Data Lake Store. Accepted values are `Consumption`, `Commitment_1TB`, `Commitment_10TB`, `Commitment_100TB`, `Commitment_500TB`, `Commitment_1PB` or `Commitment_5PB`.
	Tier pulumi.StringPtrInput
}

func (StoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storeArgs)(nil)).Elem()
}

type StoreInput interface {
	pulumi.Input

	ToStoreOutput() StoreOutput
	ToStoreOutputWithContext(ctx context.Context) StoreOutput
}

func (Store) ElementType() reflect.Type {
	return reflect.TypeOf((*Store)(nil)).Elem()
}

func (i Store) ToStoreOutput() StoreOutput {
	return i.ToStoreOutputWithContext(context.Background())
}

func (i Store) ToStoreOutputWithContext(ctx context.Context) StoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreOutput)
}

type StoreOutput struct {
	*pulumi.OutputState
}

func (StoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreOutput)(nil)).Elem()
}

func (o StoreOutput) ToStoreOutput() StoreOutput {
	return o
}

func (o StoreOutput) ToStoreOutputWithContext(ctx context.Context) StoreOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StoreOutput{})
}
