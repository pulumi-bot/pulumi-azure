// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loganalytics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Log Analytics (formally Operational Insights) Windows Performance Counter DataSource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/loganalytics"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/operationalinsights"
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
// 		exampleAnalyticsWorkspace, err := operationalinsights.NewAnalyticsWorkspace(ctx, "exampleAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("PerGB2018"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = loganalytics.NewDataSourceWindowsPerformanceCounter(ctx, "exampleDataSourceWindowsPerformanceCounter", &loganalytics.DataSourceWindowsPerformanceCounterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			WorkspaceName:     exampleAnalyticsWorkspace.Name,
// 			ObjectName:        pulumi.String("CPU"),
// 			InstanceName:      pulumi.String("*"),
// 			CounterName:       pulumi.String("CPU"),
// 			IntervalSeconds:   pulumi.Int(10),
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
// Log Analytics Windows Performance Counter DataSources can be imported using the `resource id`, e.g.
type DataSourceWindowsPerformanceCounter struct {
	pulumi.CustomResourceState

	// The friendly name of the performance counter.
	CounterName pulumi.StringOutput `pulumi:"counterName"`
	// The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The time of sample interval in seconds. Supports values between 10 and 2147483647.
	IntervalSeconds pulumi.IntOutput `pulumi:"intervalSeconds"`
	// The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The object name of the Log Analytics Windows Performance Counter DataSource.
	ObjectName pulumi.StringOutput `pulumi:"objectName"`
	// The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	WorkspaceName pulumi.StringOutput `pulumi:"workspaceName"`
}

// NewDataSourceWindowsPerformanceCounter registers a new resource with the given unique name, arguments, and options.
func NewDataSourceWindowsPerformanceCounter(ctx *pulumi.Context,
	name string, args *DataSourceWindowsPerformanceCounterArgs, opts ...pulumi.ResourceOption) (*DataSourceWindowsPerformanceCounter, error) {
	if args == nil || args.CounterName == nil {
		return nil, errors.New("missing required argument 'CounterName'")
	}
	if args == nil || args.InstanceName == nil {
		return nil, errors.New("missing required argument 'InstanceName'")
	}
	if args == nil || args.IntervalSeconds == nil {
		return nil, errors.New("missing required argument 'IntervalSeconds'")
	}
	if args == nil || args.ObjectName == nil {
		return nil, errors.New("missing required argument 'ObjectName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &DataSourceWindowsPerformanceCounterArgs{}
	}
	var resource DataSourceWindowsPerformanceCounter
	err := ctx.RegisterResource("azure:loganalytics/dataSourceWindowsPerformanceCounter:DataSourceWindowsPerformanceCounter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSourceWindowsPerformanceCounter gets an existing DataSourceWindowsPerformanceCounter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSourceWindowsPerformanceCounter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceWindowsPerformanceCounterState, opts ...pulumi.ResourceOption) (*DataSourceWindowsPerformanceCounter, error) {
	var resource DataSourceWindowsPerformanceCounter
	err := ctx.ReadResource("azure:loganalytics/dataSourceWindowsPerformanceCounter:DataSourceWindowsPerformanceCounter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSourceWindowsPerformanceCounter resources.
type dataSourceWindowsPerformanceCounterState struct {
	// The friendly name of the performance counter.
	CounterName *string `pulumi:"counterName"`
	// The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
	InstanceName *string `pulumi:"instanceName"`
	// The time of sample interval in seconds. Supports values between 10 and 2147483647.
	IntervalSeconds *int `pulumi:"intervalSeconds"`
	// The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	Name *string `pulumi:"name"`
	// The object name of the Log Analytics Windows Performance Counter DataSource.
	ObjectName *string `pulumi:"objectName"`
	// The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	WorkspaceName *string `pulumi:"workspaceName"`
}

type DataSourceWindowsPerformanceCounterState struct {
	// The friendly name of the performance counter.
	CounterName pulumi.StringPtrInput
	// The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
	InstanceName pulumi.StringPtrInput
	// The time of sample interval in seconds. Supports values between 10 and 2147483647.
	IntervalSeconds pulumi.IntPtrInput
	// The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	Name pulumi.StringPtrInput
	// The object name of the Log Analytics Windows Performance Counter DataSource.
	ObjectName pulumi.StringPtrInput
	// The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	WorkspaceName pulumi.StringPtrInput
}

func (DataSourceWindowsPerformanceCounterState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceWindowsPerformanceCounterState)(nil)).Elem()
}

type dataSourceWindowsPerformanceCounterArgs struct {
	// The friendly name of the performance counter.
	CounterName string `pulumi:"counterName"`
	// The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
	InstanceName string `pulumi:"instanceName"`
	// The time of sample interval in seconds. Supports values between 10 and 2147483647.
	IntervalSeconds int `pulumi:"intervalSeconds"`
	// The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	Name *string `pulumi:"name"`
	// The object name of the Log Analytics Windows Performance Counter DataSource.
	ObjectName string `pulumi:"objectName"`
	// The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataSourceWindowsPerformanceCounter resource.
type DataSourceWindowsPerformanceCounterArgs struct {
	// The friendly name of the performance counter.
	CounterName pulumi.StringInput
	// The name of the virtual machine instance to which the Windows Performance Counter DataSource be applied. Specify a `*` will apply to all instances.
	InstanceName pulumi.StringInput
	// The time of sample interval in seconds. Supports values between 10 and 2147483647.
	IntervalSeconds pulumi.IntInput
	// The Name which should be used for this Log Analytics Windows Performance Counter DataSource. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	Name pulumi.StringPtrInput
	// The object name of the Log Analytics Windows Performance Counter DataSource.
	ObjectName pulumi.StringInput
	// The name of the Resource Group where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the Log Analytics Workspace where the Log Analytics Windows Performance Counter DataSource should exist. Changing this forces a new Log Analytics Windows Performance Counter DataSource to be created.
	WorkspaceName pulumi.StringInput
}

func (DataSourceWindowsPerformanceCounterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceWindowsPerformanceCounterArgs)(nil)).Elem()
}
