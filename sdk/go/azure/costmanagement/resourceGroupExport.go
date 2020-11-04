// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package costmanagement

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Cost Management Export for a Resource Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/costmanagement"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = costmanagement.NewResourceGroupExport(ctx, "exampleResourceGroupExport", &costmanagement.ResourceGroupExportArgs{
// 			ResourceGroupId:       exampleResourceGroup.ID(),
// 			RecurrenceType:        pulumi.String("Monthly"),
// 			RecurrencePeriodStart: pulumi.String("2020-08-18T00:00:00Z"),
// 			RecurrencePeriodEnd:   pulumi.String("2020-09-18T00:00:00Z"),
// 			DeliveryInfo: &costmanagement.ResourceGroupExportDeliveryInfoArgs{
// 				StorageAccountId: exampleAccount.ID(),
// 				ContainerName:    pulumi.String("examplecontainer"),
// 				RootFolderPath:   pulumi.String("/root/updated"),
// 			},
// 			Query: &costmanagement.ResourceGroupExportQueryArgs{
// 				Type:      pulumi.String("Usage"),
// 				TimeFrame: pulumi.String("WeekToDate"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ResourceGroupExport struct {
	pulumi.CustomResourceState

	// Is the cost management export active? Default is `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// A `deliveryInfo` block as defined below.
	DeliveryInfo ResourceGroupExportDeliveryInfoOutput `pulumi:"deliveryInfo"`
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `query` block as defined below.
	Query ResourceGroupExportQueryOutput `pulumi:"query"`
	// The date the export will stop capturing information.
	RecurrencePeriodEnd pulumi.StringOutput `pulumi:"recurrencePeriodEnd"`
	// The date the export will start capturing information.
	RecurrencePeriodStart pulumi.StringOutput `pulumi:"recurrencePeriodStart"`
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType pulumi.StringOutput `pulumi:"recurrenceType"`
	// The id of the resource group in which to export information.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
}

// NewResourceGroupExport registers a new resource with the given unique name, arguments, and options.
func NewResourceGroupExport(ctx *pulumi.Context,
	name string, args *ResourceGroupExportArgs, opts ...pulumi.ResourceOption) (*ResourceGroupExport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DeliveryInfo == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryInfo'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.RecurrencePeriodEnd == nil {
		return nil, errors.New("invalid value for required argument 'RecurrencePeriodEnd'")
	}
	if args.RecurrencePeriodStart == nil {
		return nil, errors.New("invalid value for required argument 'RecurrencePeriodStart'")
	}
	if args.RecurrenceType == nil {
		return nil, errors.New("invalid value for required argument 'RecurrenceType'")
	}
	if args.ResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupId'")
	}
	var resource ResourceGroupExport
	err := ctx.RegisterResource("azure:costmanagement/resourceGroupExport:ResourceGroupExport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGroupExport gets an existing ResourceGroupExport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroupExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupExportState, opts ...pulumi.ResourceOption) (*ResourceGroupExport, error) {
	var resource ResourceGroupExport
	err := ctx.ReadResource("azure:costmanagement/resourceGroupExport:ResourceGroupExport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGroupExport resources.
type resourceGroupExportState struct {
	// Is the cost management export active? Default is `true`.
	Active *bool `pulumi:"active"`
	// A `deliveryInfo` block as defined below.
	DeliveryInfo *ResourceGroupExportDeliveryInfo `pulumi:"deliveryInfo"`
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `query` block as defined below.
	Query *ResourceGroupExportQuery `pulumi:"query"`
	// The date the export will stop capturing information.
	RecurrencePeriodEnd *string `pulumi:"recurrencePeriodEnd"`
	// The date the export will start capturing information.
	RecurrencePeriodStart *string `pulumi:"recurrencePeriodStart"`
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType *string `pulumi:"recurrenceType"`
	// The id of the resource group in which to export information.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}

type ResourceGroupExportState struct {
	// Is the cost management export active? Default is `true`.
	Active pulumi.BoolPtrInput
	// A `deliveryInfo` block as defined below.
	DeliveryInfo ResourceGroupExportDeliveryInfoPtrInput
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `query` block as defined below.
	Query ResourceGroupExportQueryPtrInput
	// The date the export will stop capturing information.
	RecurrencePeriodEnd pulumi.StringPtrInput
	// The date the export will start capturing information.
	RecurrencePeriodStart pulumi.StringPtrInput
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType pulumi.StringPtrInput
	// The id of the resource group in which to export information.
	ResourceGroupId pulumi.StringPtrInput
}

func (ResourceGroupExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupExportState)(nil)).Elem()
}

type resourceGroupExportArgs struct {
	// Is the cost management export active? Default is `true`.
	Active *bool `pulumi:"active"`
	// A `deliveryInfo` block as defined below.
	DeliveryInfo ResourceGroupExportDeliveryInfo `pulumi:"deliveryInfo"`
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `query` block as defined below.
	Query ResourceGroupExportQuery `pulumi:"query"`
	// The date the export will stop capturing information.
	RecurrencePeriodEnd string `pulumi:"recurrencePeriodEnd"`
	// The date the export will start capturing information.
	RecurrencePeriodStart string `pulumi:"recurrencePeriodStart"`
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType string `pulumi:"recurrenceType"`
	// The id of the resource group in which to export information.
	ResourceGroupId string `pulumi:"resourceGroupId"`
}

// The set of arguments for constructing a ResourceGroupExport resource.
type ResourceGroupExportArgs struct {
	// Is the cost management export active? Default is `true`.
	Active pulumi.BoolPtrInput
	// A `deliveryInfo` block as defined below.
	DeliveryInfo ResourceGroupExportDeliveryInfoInput
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `query` block as defined below.
	Query ResourceGroupExportQueryInput
	// The date the export will stop capturing information.
	RecurrencePeriodEnd pulumi.StringInput
	// The date the export will start capturing information.
	RecurrencePeriodStart pulumi.StringInput
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType pulumi.StringInput
	// The id of the resource group in which to export information.
	ResourceGroupId pulumi.StringInput
}

func (ResourceGroupExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupExportArgs)(nil)).Elem()
}
