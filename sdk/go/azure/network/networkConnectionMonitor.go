// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Network Connection Monitor.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleNetworkWatcher, err := network.NewNetworkWatcher(ctx, "exampleNetworkWatcher", &network.NetworkWatcherArgs{
// 			Location:          pulumi.String(exampleResourceGroup.Location),
// 			ResourceGroupName: pulumi.String(exampleResourceGroup.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		srcVirtualMachine, err := compute.LookupVirtualMachine(ctx, &compute.LookupVirtualMachineArgs{
// 			Name:              "example-vm",
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		srcExtension, err := compute.NewExtension(ctx, "srcExtension", &compute.ExtensionArgs{
// 			VirtualMachineId:        pulumi.String(srcVirtualMachine.Id),
// 			Publisher:               pulumi.String("Microsoft.Azure.NetworkWatcher"),
// 			Type:                    pulumi.String("NetworkWatcherAgentLinux"),
// 			TypeHandlerVersion:      pulumi.String("1.4"),
// 			AutoUpgradeMinorVersion: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkConnectionMonitor(ctx, "exampleNetworkConnectionMonitor", &network.NetworkConnectionMonitorArgs{
// 			NetworkWatcherName: exampleNetworkWatcher.Name,
// 			ResourceGroupName:  pulumi.String(exampleResourceGroup.Name),
// 			Location:           exampleNetworkWatcher.Location,
// 			AutoStart:          pulumi.Bool(false),
// 			IntervalInSeconds:  pulumi.Int(30),
// 			Source: &network.NetworkConnectionMonitorSourceArgs{
// 				VirtualMachineId: pulumi.String(srcVirtualMachine.Id),
// 				Port:             pulumi.Int(20020),
// 			},
// 			Destination: &network.NetworkConnectionMonitorDestinationArgs{
// 				Address: pulumi.String("mycompany.io"),
// 				Port:    pulumi.Int(443),
// 			},
// 			Tags: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			srcExtension,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NetworkConnectionMonitor struct {
	pulumi.CustomResourceState

	// Will the connection monitor start automatically once created? Changing this forces a new Network Connection Monitor to be created.
	AutoStart pulumi.BoolPtrOutput `pulumi:"autoStart"`
	// A `destination` block as defined below.
	Destination NetworkConnectionMonitorDestinationOutput `pulumi:"destination"`
	// Monitoring interval in seconds.
	IntervalInSeconds pulumi.IntPtrOutput `pulumi:"intervalInSeconds"`
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Network Connection Monitor. Changing this forces a new Network Connection Monitor to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new Network Connection Monitor to be created.
	NetworkWatcherName pulumi.StringOutput `pulumi:"networkWatcherName"`
	// The name of the Resource Group where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `source` block as defined below.
	Source NetworkConnectionMonitorSourceOutput `pulumi:"source"`
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewNetworkConnectionMonitor registers a new resource with the given unique name, arguments, and options.
func NewNetworkConnectionMonitor(ctx *pulumi.Context,
	name string, args *NetworkConnectionMonitorArgs, opts ...pulumi.ResourceOption) (*NetworkConnectionMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.NetworkWatcherName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkWatcherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource NetworkConnectionMonitor
	err := ctx.RegisterResource("azure:network/networkConnectionMonitor:NetworkConnectionMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkConnectionMonitor gets an existing NetworkConnectionMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkConnectionMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkConnectionMonitorState, opts ...pulumi.ResourceOption) (*NetworkConnectionMonitor, error) {
	var resource NetworkConnectionMonitor
	err := ctx.ReadResource("azure:network/networkConnectionMonitor:NetworkConnectionMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkConnectionMonitor resources.
type networkConnectionMonitorState struct {
	// Will the connection monitor start automatically once created? Changing this forces a new Network Connection Monitor to be created.
	AutoStart *bool `pulumi:"autoStart"`
	// A `destination` block as defined below.
	Destination *NetworkConnectionMonitorDestination `pulumi:"destination"`
	// Monitoring interval in seconds.
	IntervalInSeconds *int `pulumi:"intervalInSeconds"`
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Network Connection Monitor. Changing this forces a new Network Connection Monitor to be created.
	Name *string `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new Network Connection Monitor to be created.
	NetworkWatcherName *string `pulumi:"networkWatcherName"`
	// The name of the Resource Group where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `source` block as defined below.
	Source *NetworkConnectionMonitorSource `pulumi:"source"`
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags map[string]string `pulumi:"tags"`
}

type NetworkConnectionMonitorState struct {
	// Will the connection monitor start automatically once created? Changing this forces a new Network Connection Monitor to be created.
	AutoStart pulumi.BoolPtrInput
	// A `destination` block as defined below.
	Destination NetworkConnectionMonitorDestinationPtrInput
	// Monitoring interval in seconds.
	IntervalInSeconds pulumi.IntPtrInput
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Network Connection Monitor. Changing this forces a new Network Connection Monitor to be created.
	Name pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new Network Connection Monitor to be created.
	NetworkWatcherName pulumi.StringPtrInput
	// The name of the Resource Group where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `source` block as defined below.
	Source NetworkConnectionMonitorSourcePtrInput
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags pulumi.StringMapInput
}

func (NetworkConnectionMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectionMonitorState)(nil)).Elem()
}

type networkConnectionMonitorArgs struct {
	// Will the connection monitor start automatically once created? Changing this forces a new Network Connection Monitor to be created.
	AutoStart *bool `pulumi:"autoStart"`
	// A `destination` block as defined below.
	Destination NetworkConnectionMonitorDestination `pulumi:"destination"`
	// Monitoring interval in seconds.
	IntervalInSeconds *int `pulumi:"intervalInSeconds"`
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Network Connection Monitor. Changing this forces a new Network Connection Monitor to be created.
	Name *string `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new Network Connection Monitor to be created.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the Resource Group where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `source` block as defined below.
	Source NetworkConnectionMonitorSource `pulumi:"source"`
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkConnectionMonitor resource.
type NetworkConnectionMonitorArgs struct {
	// Will the connection monitor start automatically once created? Changing this forces a new Network Connection Monitor to be created.
	AutoStart pulumi.BoolPtrInput
	// A `destination` block as defined below.
	Destination NetworkConnectionMonitorDestinationInput
	// Monitoring interval in seconds.
	IntervalInSeconds pulumi.IntPtrInput
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Network Connection Monitor. Changing this forces a new Network Connection Monitor to be created.
	Name pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new Network Connection Monitor to be created.
	NetworkWatcherName pulumi.StringInput
	// The name of the Resource Group where the Network Connection Monitor should exist. Changing this forces a new Network Connection Monitor to be created.
	ResourceGroupName pulumi.StringInput
	// A `source` block as defined below.
	Source NetworkConnectionMonitorSourceInput
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags pulumi.StringMapInput
}

func (NetworkConnectionMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectionMonitorArgs)(nil)).Elem()
}
