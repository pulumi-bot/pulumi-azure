// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Network Connection Monitors can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkConnectionMonitor:NetworkConnectionMonitor example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/networkWatchers/watcher1/connectionMonitors/connectionMonitor1
// ```
type NetworkConnectionMonitor struct {
	pulumi.CustomResourceState

	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	AutoStart pulumi.BoolOutput `pulumi:"autoStart"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Destination NetworkConnectionMonitorDestinationOutput `pulumi:"destination"`
	// A `endpoint` block as defined below.
	Endpoints NetworkConnectionMonitorEndpointArrayOutput `pulumi:"endpoints"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	IntervalInSeconds pulumi.IntOutput `pulumi:"intervalInSeconds"`
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherId pulumi.StringOutput `pulumi:"networkWatcherId"`
	// The description of the Network Connection Monitor.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	OutputWorkspaceResourceIds pulumi.StringArrayOutput `pulumi:"outputWorkspaceResourceIds"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Source NetworkConnectionMonitorSourceOutput `pulumi:"source"`
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `testConfiguration` block as defined below.
	TestConfigurations NetworkConnectionMonitorTestConfigurationArrayOutput `pulumi:"testConfigurations"`
	// A `testGroup` block as defined below.
	TestGroups NetworkConnectionMonitorTestGroupArrayOutput `pulumi:"testGroups"`
}

// NewNetworkConnectionMonitor registers a new resource with the given unique name, arguments, and options.
func NewNetworkConnectionMonitor(ctx *pulumi.Context,
	name string, args *NetworkConnectionMonitorArgs, opts ...pulumi.ResourceOption) (*NetworkConnectionMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endpoints == nil {
		return nil, errors.New("invalid value for required argument 'Endpoints'")
	}
	if args.NetworkWatcherId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkWatcherId'")
	}
	if args.TestConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'TestConfigurations'")
	}
	if args.TestGroups == nil {
		return nil, errors.New("invalid value for required argument 'TestGroups'")
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
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	AutoStart *bool `pulumi:"autoStart"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Destination *NetworkConnectionMonitorDestination `pulumi:"destination"`
	// A `endpoint` block as defined below.
	Endpoints []NetworkConnectionMonitorEndpoint `pulumi:"endpoints"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	IntervalInSeconds *int `pulumi:"intervalInSeconds"`
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherId *string `pulumi:"networkWatcherId"`
	// The description of the Network Connection Monitor.
	Notes *string `pulumi:"notes"`
	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	OutputWorkspaceResourceIds []string `pulumi:"outputWorkspaceResourceIds"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Source *NetworkConnectionMonitorSource `pulumi:"source"`
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags map[string]string `pulumi:"tags"`
	// A `testConfiguration` block as defined below.
	TestConfigurations []NetworkConnectionMonitorTestConfiguration `pulumi:"testConfigurations"`
	// A `testGroup` block as defined below.
	TestGroups []NetworkConnectionMonitorTestGroup `pulumi:"testGroups"`
}

type NetworkConnectionMonitorState struct {
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	AutoStart pulumi.BoolPtrInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Destination NetworkConnectionMonitorDestinationPtrInput
	// A `endpoint` block as defined below.
	Endpoints NetworkConnectionMonitorEndpointArrayInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	IntervalInSeconds pulumi.IntPtrInput
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherId pulumi.StringPtrInput
	// The description of the Network Connection Monitor.
	Notes pulumi.StringPtrInput
	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	OutputWorkspaceResourceIds pulumi.StringArrayInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Source NetworkConnectionMonitorSourcePtrInput
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags pulumi.StringMapInput
	// A `testConfiguration` block as defined below.
	TestConfigurations NetworkConnectionMonitorTestConfigurationArrayInput
	// A `testGroup` block as defined below.
	TestGroups NetworkConnectionMonitorTestGroupArrayInput
}

func (NetworkConnectionMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectionMonitorState)(nil)).Elem()
}

type networkConnectionMonitorArgs struct {
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	AutoStart *bool `pulumi:"autoStart"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Destination *NetworkConnectionMonitorDestination `pulumi:"destination"`
	// A `endpoint` block as defined below.
	Endpoints []NetworkConnectionMonitorEndpoint `pulumi:"endpoints"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	IntervalInSeconds *int `pulumi:"intervalInSeconds"`
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherId string `pulumi:"networkWatcherId"`
	// The description of the Network Connection Monitor.
	Notes *string `pulumi:"notes"`
	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	OutputWorkspaceResourceIds []string `pulumi:"outputWorkspaceResourceIds"`
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Source *NetworkConnectionMonitorSource `pulumi:"source"`
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags map[string]string `pulumi:"tags"`
	// A `testConfiguration` block as defined below.
	TestConfigurations []NetworkConnectionMonitorTestConfiguration `pulumi:"testConfigurations"`
	// A `testGroup` block as defined below.
	TestGroups []NetworkConnectionMonitorTestGroup `pulumi:"testGroups"`
}

// The set of arguments for constructing a NetworkConnectionMonitor resource.
type NetworkConnectionMonitorArgs struct {
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	AutoStart pulumi.BoolPtrInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Destination NetworkConnectionMonitorDestinationPtrInput
	// A `endpoint` block as defined below.
	Endpoints NetworkConnectionMonitorEndpointArrayInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	IntervalInSeconds pulumi.IntPtrInput
	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Network Connection Monitor. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherId pulumi.StringInput
	// The description of the Network Connection Monitor.
	Notes pulumi.StringPtrInput
	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	OutputWorkspaceResourceIds pulumi.StringArrayInput
	// Deprecated: The field belongs to the v1 network connection monitor, which is now deprecated in favour of v2 by Azure. Please check the document (https://www.terraform.io/docs/providers/azurerm/r/network_connection_monitor.html) for the v2 properties.
	Source NetworkConnectionMonitorSourcePtrInput
	// A mapping of tags which should be assigned to the Network Connection Monitor.
	Tags pulumi.StringMapInput
	// A `testConfiguration` block as defined below.
	TestConfigurations NetworkConnectionMonitorTestConfigurationArrayInput
	// A `testGroup` block as defined below.
	TestGroups NetworkConnectionMonitorTestGroupArrayInput
}

func (NetworkConnectionMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectionMonitorArgs)(nil)).Elem()
}

type NetworkConnectionMonitorInput interface {
	pulumi.Input

	ToNetworkConnectionMonitorOutput() NetworkConnectionMonitorOutput
	ToNetworkConnectionMonitorOutputWithContext(ctx context.Context) NetworkConnectionMonitorOutput
}

func (NetworkConnectionMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConnectionMonitor)(nil)).Elem()
}

func (i NetworkConnectionMonitor) ToNetworkConnectionMonitorOutput() NetworkConnectionMonitorOutput {
	return i.ToNetworkConnectionMonitorOutputWithContext(context.Background())
}

func (i NetworkConnectionMonitor) ToNetworkConnectionMonitorOutputWithContext(ctx context.Context) NetworkConnectionMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConnectionMonitorOutput)
}

type NetworkConnectionMonitorOutput struct {
	*pulumi.OutputState
}

func (NetworkConnectionMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConnectionMonitorOutput)(nil)).Elem()
}

func (o NetworkConnectionMonitorOutput) ToNetworkConnectionMonitorOutput() NetworkConnectionMonitorOutput {
	return o
}

func (o NetworkConnectionMonitorOutput) ToNetworkConnectionMonitorOutputWithContext(ctx context.Context) NetworkConnectionMonitorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkConnectionMonitorOutput{})
}
