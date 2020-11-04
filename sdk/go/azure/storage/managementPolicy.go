// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Storage Account Management Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("westus"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			AccountKind:            pulumi.String("BlobStorage"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewManagementPolicy(ctx, "exampleManagementPolicy", &storage.ManagementPolicyArgs{
// 			StorageAccountId: exampleAccount.ID(),
// 			Rules: storage.ManagementPolicyRuleArray{
// 				&storage.ManagementPolicyRuleArgs{
// 					Name:    pulumi.String("rule1"),
// 					Enabled: pulumi.Bool(true),
// 					Filters: &storage.ManagementPolicyRuleFiltersArgs{
// 						PrefixMatches: pulumi.StringArray{
// 							pulumi.String("container1/prefix1"),
// 						},
// 						BlobTypes: pulumi.StringArray{
// 							pulumi.String("blockBlob"),
// 						},
// 					},
// 					Actions: &storage.ManagementPolicyRuleActionsArgs{
// 						BaseBlob: &storage.ManagementPolicyRuleActionsBaseBlobArgs{
// 							TierToCoolAfterDaysSinceModificationGreaterThan:    pulumi.Int(10),
// 							TierToArchiveAfterDaysSinceModificationGreaterThan: pulumi.Int(50),
// 							DeleteAfterDaysSinceModificationGreaterThan:        pulumi.Int(100),
// 						},
// 						Snapshot: &storage.ManagementPolicyRuleActionsSnapshotArgs{
// 							DeleteAfterDaysSinceCreationGreaterThan: pulumi.Int(30),
// 						},
// 					},
// 				},
// 				&storage.ManagementPolicyRuleArgs{
// 					Name:    pulumi.String("rule2"),
// 					Enabled: pulumi.Bool(false),
// 					Filters: &storage.ManagementPolicyRuleFiltersArgs{
// 						PrefixMatches: pulumi.StringArray{
// 							pulumi.String("container2/prefix1"),
// 							pulumi.String("container2/prefix2"),
// 						},
// 						BlobTypes: pulumi.StringArray{
// 							pulumi.String("blockBlob"),
// 						},
// 					},
// 					Actions: &storage.ManagementPolicyRuleActionsArgs{
// 						BaseBlob: &storage.ManagementPolicyRuleActionsBaseBlobArgs{
// 							TierToCoolAfterDaysSinceModificationGreaterThan:    pulumi.Int(11),
// 							TierToArchiveAfterDaysSinceModificationGreaterThan: pulumi.Int(51),
// 							DeleteAfterDaysSinceModificationGreaterThan:        pulumi.Int(101),
// 						},
// 						Snapshot: &storage.ManagementPolicyRuleActionsSnapshotArgs{
// 							DeleteAfterDaysSinceCreationGreaterThan: pulumi.Int(31),
// 						},
// 					},
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
type ManagementPolicy struct {
	pulumi.CustomResourceState

	// A `rule` block as documented below.
	Rules ManagementPolicyRuleArrayOutput `pulumi:"rules"`
	// Specifies the id of the storage account to apply the management policy to.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
}

// NewManagementPolicy registers a new resource with the given unique name, arguments, and options.
func NewManagementPolicy(ctx *pulumi.Context,
	name string, args *ManagementPolicyArgs, opts ...pulumi.ResourceOption) (*ManagementPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	var resource ManagementPolicy
	err := ctx.RegisterResource("azure:storage/managementPolicy:ManagementPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementPolicy gets an existing ManagementPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementPolicyState, opts ...pulumi.ResourceOption) (*ManagementPolicy, error) {
	var resource ManagementPolicy
	err := ctx.ReadResource("azure:storage/managementPolicy:ManagementPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementPolicy resources.
type managementPolicyState struct {
	// A `rule` block as documented below.
	Rules []ManagementPolicyRule `pulumi:"rules"`
	// Specifies the id of the storage account to apply the management policy to.
	StorageAccountId *string `pulumi:"storageAccountId"`
}

type ManagementPolicyState struct {
	// A `rule` block as documented below.
	Rules ManagementPolicyRuleArrayInput
	// Specifies the id of the storage account to apply the management policy to.
	StorageAccountId pulumi.StringPtrInput
}

func (ManagementPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementPolicyState)(nil)).Elem()
}

type managementPolicyArgs struct {
	// A `rule` block as documented below.
	Rules []ManagementPolicyRule `pulumi:"rules"`
	// Specifies the id of the storage account to apply the management policy to.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// The set of arguments for constructing a ManagementPolicy resource.
type ManagementPolicyArgs struct {
	// A `rule` block as documented below.
	Rules ManagementPolicyRuleArrayInput
	// Specifies the id of the storage account to apply the management policy to.
	StorageAccountId pulumi.StringInput
}

func (ManagementPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementPolicyArgs)(nil)).Elem()
}
