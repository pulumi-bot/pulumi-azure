// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Node Pool within a Kubernetes Cluster
//
// > **NOTE:** Multiple Node Pools are only supported when the Kubernetes Cluster is using Virtual Machine Scale Sets.
//
// ## Example Usage
//
// This example provisions a basic Kubernetes Node Pool.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/containerservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
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
// 		exampleKubernetesCluster, err := containerservice.NewKubernetesCluster(ctx, "exampleKubernetesCluster", &containerservice.KubernetesClusterArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DnsPrefix:         pulumi.String("exampleaks1"),
// 			DefaultNodePool: &containerservice.KubernetesClusterDefaultNodePoolArgs{
// 				Name:      pulumi.String("default"),
// 				NodeCount: pulumi.Int(1),
// 				VmSize:    pulumi.String("Standard_D2_v2"),
// 			},
// 			ServicePrincipal: &containerservice.KubernetesClusterServicePrincipalArgs{
// 				ClientId:     pulumi.String("00000000-0000-0000-0000-000000000000"),
// 				ClientSecret: pulumi.String("00000000000000000000000000000000"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = containerservice.NewKubernetesClusterNodePool(ctx, "exampleKubernetesClusterNodePool", &containerservice.KubernetesClusterNodePoolArgs{
// 			KubernetesClusterId: exampleKubernetesCluster.ID(),
// 			VmSize:              pulumi.String("Standard_DS2_v2"),
// 			NodeCount:           pulumi.Int(1),
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type KubernetesClusterNodePool struct {
	pulumi.CustomResourceState

	// A list of Availability Zones where the Nodes in this Node Pool should be created in.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// Whether to enable [auto-scaler](https://docs.microsoft.com/en-us/azure/aks/cluster-autoscaler). Defaults to `false`.
	EnableAutoScaling pulumi.BoolPtrOutput `pulumi:"enableAutoScaling"`
	// Should each node have a Public IP Address? Defaults to `false`.
	EnableNodePublicIp pulumi.BoolPtrOutput `pulumi:"enableNodePublicIp"`
	// The Eviction Policy which should be used for Virtual Machines within the Virtual Machine Scale Set powering this Node Pool. Possible values are `Deallocate` and `Delete`. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrOutput `pulumi:"evictionPolicy"`
	// The ID of the Kubernetes Cluster where this Node Pool should exist. Changing this forces a new resource to be created.
	KubernetesClusterId pulumi.StringOutput `pulumi:"kubernetesClusterId"`
	// The maximum number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be greater than or equal to `minCount`.
	MaxCount pulumi.IntPtrOutput `pulumi:"maxCount"`
	// The maximum number of pods that can run on each agent. Changing this forces a new resource to be created.
	MaxPods pulumi.IntOutput `pulumi:"maxPods"`
	// The minimum number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be less than or equal to `maxCount`.
	MinCount pulumi.IntPtrOutput `pulumi:"minCount"`
	// Should this Node Pool be used for System or User resources? Possible values are `System` and `User`. Defaults to `User`.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the Node Pool which should be created within the Kubernetes Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The initial number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be a value in the range `minCount` - `maxCount`.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// A map of Kubernetes labels which should be applied to nodes in this Node Pool. Changing this forces a new resource to be created.
	NodeLabels pulumi.StringMapOutput `pulumi:"nodeLabels"`
	// A list of Kubernetes taints which should be applied to nodes in the agent pool (e.g `key=value:NoSchedule`). Changing this forces a new resource to be created.
	NodeTaints pulumi.StringArrayOutput `pulumi:"nodeTaints"`
	// Version of Kubernetes used for the Agents. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade)
	OrchestratorVersion pulumi.StringOutput `pulumi:"orchestratorVersion"`
	// The Agent Operating System disk size in GB. Changing this forces a new resource to be created.
	OsDiskSizeGb pulumi.IntOutput `pulumi:"osDiskSizeGb"`
	// The Operating System which should be used for this Node Pool. Changing this forces a new resource to be created. Possible values are `Linux` and `Windows`. Defaults to `Linux`.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The Priority for Virtual Machines within the Virtual Machine Scale Set that powers this Node Pool. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
	Priority pulumi.StringPtrOutput `pulumi:"priority"`
	// The maximum price you're willing to pay in USD per Virtual Machine. Valid values are `-1` (the current on-demand price for a Virtual Machine) or a positive value with up to five decimal places. Changing this forces a new resource to be created.
	SpotMaxPrice pulumi.Float64PtrOutput `pulumi:"spotMaxPrice"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The SKU which should be used for the Virtual Machines used in this Node Pool. Changing this forces a new resource to be created.
	VmSize pulumi.StringOutput `pulumi:"vmSize"`
	// The ID of the Subnet where this Node Pool should exist.
	VnetSubnetId pulumi.StringPtrOutput `pulumi:"vnetSubnetId"`
}

// NewKubernetesClusterNodePool registers a new resource with the given unique name, arguments, and options.
func NewKubernetesClusterNodePool(ctx *pulumi.Context,
	name string, args *KubernetesClusterNodePoolArgs, opts ...pulumi.ResourceOption) (*KubernetesClusterNodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KubernetesClusterId == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesClusterId'")
	}
	if args.VmSize == nil {
		return nil, errors.New("invalid value for required argument 'VmSize'")
	}
	var resource KubernetesClusterNodePool
	err := ctx.RegisterResource("azure:containerservice/kubernetesClusterNodePool:KubernetesClusterNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesClusterNodePool gets an existing KubernetesClusterNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesClusterNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesClusterNodePoolState, opts ...pulumi.ResourceOption) (*KubernetesClusterNodePool, error) {
	var resource KubernetesClusterNodePool
	err := ctx.ReadResource("azure:containerservice/kubernetesClusterNodePool:KubernetesClusterNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesClusterNodePool resources.
type kubernetesClusterNodePoolState struct {
	// A list of Availability Zones where the Nodes in this Node Pool should be created in.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Whether to enable [auto-scaler](https://docs.microsoft.com/en-us/azure/aks/cluster-autoscaler). Defaults to `false`.
	EnableAutoScaling *bool `pulumi:"enableAutoScaling"`
	// Should each node have a Public IP Address? Defaults to `false`.
	EnableNodePublicIp *bool `pulumi:"enableNodePublicIp"`
	// The Eviction Policy which should be used for Virtual Machines within the Virtual Machine Scale Set powering this Node Pool. Possible values are `Deallocate` and `Delete`. Changing this forces a new resource to be created.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// The ID of the Kubernetes Cluster where this Node Pool should exist. Changing this forces a new resource to be created.
	KubernetesClusterId *string `pulumi:"kubernetesClusterId"`
	// The maximum number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be greater than or equal to `minCount`.
	MaxCount *int `pulumi:"maxCount"`
	// The maximum number of pods that can run on each agent. Changing this forces a new resource to be created.
	MaxPods *int `pulumi:"maxPods"`
	// The minimum number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be less than or equal to `maxCount`.
	MinCount *int `pulumi:"minCount"`
	// Should this Node Pool be used for System or User resources? Possible values are `System` and `User`. Defaults to `User`.
	Mode *string `pulumi:"mode"`
	// The name of the Node Pool which should be created within the Kubernetes Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The initial number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be a value in the range `minCount` - `maxCount`.
	NodeCount *int `pulumi:"nodeCount"`
	// A map of Kubernetes labels which should be applied to nodes in this Node Pool. Changing this forces a new resource to be created.
	NodeLabels map[string]string `pulumi:"nodeLabels"`
	// A list of Kubernetes taints which should be applied to nodes in the agent pool (e.g `key=value:NoSchedule`). Changing this forces a new resource to be created.
	NodeTaints []string `pulumi:"nodeTaints"`
	// Version of Kubernetes used for the Agents. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade)
	OrchestratorVersion *string `pulumi:"orchestratorVersion"`
	// The Agent Operating System disk size in GB. Changing this forces a new resource to be created.
	OsDiskSizeGb *int `pulumi:"osDiskSizeGb"`
	// The Operating System which should be used for this Node Pool. Changing this forces a new resource to be created. Possible values are `Linux` and `Windows`. Defaults to `Linux`.
	OsType *string `pulumi:"osType"`
	// The Priority for Virtual Machines within the Virtual Machine Scale Set that powers this Node Pool. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
	Priority *string `pulumi:"priority"`
	// The maximum price you're willing to pay in USD per Virtual Machine. Valid values are `-1` (the current on-demand price for a Virtual Machine) or a positive value with up to five decimal places. Changing this forces a new resource to be created.
	SpotMaxPrice *float64 `pulumi:"spotMaxPrice"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The SKU which should be used for the Virtual Machines used in this Node Pool. Changing this forces a new resource to be created.
	VmSize *string `pulumi:"vmSize"`
	// The ID of the Subnet where this Node Pool should exist.
	VnetSubnetId *string `pulumi:"vnetSubnetId"`
}

type KubernetesClusterNodePoolState struct {
	// A list of Availability Zones where the Nodes in this Node Pool should be created in.
	AvailabilityZones pulumi.StringArrayInput
	// Whether to enable [auto-scaler](https://docs.microsoft.com/en-us/azure/aks/cluster-autoscaler). Defaults to `false`.
	EnableAutoScaling pulumi.BoolPtrInput
	// Should each node have a Public IP Address? Defaults to `false`.
	EnableNodePublicIp pulumi.BoolPtrInput
	// The Eviction Policy which should be used for Virtual Machines within the Virtual Machine Scale Set powering this Node Pool. Possible values are `Deallocate` and `Delete`. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrInput
	// The ID of the Kubernetes Cluster where this Node Pool should exist. Changing this forces a new resource to be created.
	KubernetesClusterId pulumi.StringPtrInput
	// The maximum number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be greater than or equal to `minCount`.
	MaxCount pulumi.IntPtrInput
	// The maximum number of pods that can run on each agent. Changing this forces a new resource to be created.
	MaxPods pulumi.IntPtrInput
	// The minimum number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be less than or equal to `maxCount`.
	MinCount pulumi.IntPtrInput
	// Should this Node Pool be used for System or User resources? Possible values are `System` and `User`. Defaults to `User`.
	Mode pulumi.StringPtrInput
	// The name of the Node Pool which should be created within the Kubernetes Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The initial number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be a value in the range `minCount` - `maxCount`.
	NodeCount pulumi.IntPtrInput
	// A map of Kubernetes labels which should be applied to nodes in this Node Pool. Changing this forces a new resource to be created.
	NodeLabels pulumi.StringMapInput
	// A list of Kubernetes taints which should be applied to nodes in the agent pool (e.g `key=value:NoSchedule`). Changing this forces a new resource to be created.
	NodeTaints pulumi.StringArrayInput
	// Version of Kubernetes used for the Agents. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade)
	OrchestratorVersion pulumi.StringPtrInput
	// The Agent Operating System disk size in GB. Changing this forces a new resource to be created.
	OsDiskSizeGb pulumi.IntPtrInput
	// The Operating System which should be used for this Node Pool. Changing this forces a new resource to be created. Possible values are `Linux` and `Windows`. Defaults to `Linux`.
	OsType pulumi.StringPtrInput
	// The Priority for Virtual Machines within the Virtual Machine Scale Set that powers this Node Pool. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
	Priority pulumi.StringPtrInput
	// The maximum price you're willing to pay in USD per Virtual Machine. Valid values are `-1` (the current on-demand price for a Virtual Machine) or a positive value with up to five decimal places. Changing this forces a new resource to be created.
	SpotMaxPrice pulumi.Float64PtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The SKU which should be used for the Virtual Machines used in this Node Pool. Changing this forces a new resource to be created.
	VmSize pulumi.StringPtrInput
	// The ID of the Subnet where this Node Pool should exist.
	VnetSubnetId pulumi.StringPtrInput
}

func (KubernetesClusterNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterNodePoolState)(nil)).Elem()
}

type kubernetesClusterNodePoolArgs struct {
	// A list of Availability Zones where the Nodes in this Node Pool should be created in.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Whether to enable [auto-scaler](https://docs.microsoft.com/en-us/azure/aks/cluster-autoscaler). Defaults to `false`.
	EnableAutoScaling *bool `pulumi:"enableAutoScaling"`
	// Should each node have a Public IP Address? Defaults to `false`.
	EnableNodePublicIp *bool `pulumi:"enableNodePublicIp"`
	// The Eviction Policy which should be used for Virtual Machines within the Virtual Machine Scale Set powering this Node Pool. Possible values are `Deallocate` and `Delete`. Changing this forces a new resource to be created.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// The ID of the Kubernetes Cluster where this Node Pool should exist. Changing this forces a new resource to be created.
	KubernetesClusterId string `pulumi:"kubernetesClusterId"`
	// The maximum number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be greater than or equal to `minCount`.
	MaxCount *int `pulumi:"maxCount"`
	// The maximum number of pods that can run on each agent. Changing this forces a new resource to be created.
	MaxPods *int `pulumi:"maxPods"`
	// The minimum number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be less than or equal to `maxCount`.
	MinCount *int `pulumi:"minCount"`
	// Should this Node Pool be used for System or User resources? Possible values are `System` and `User`. Defaults to `User`.
	Mode *string `pulumi:"mode"`
	// The name of the Node Pool which should be created within the Kubernetes Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The initial number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be a value in the range `minCount` - `maxCount`.
	NodeCount *int `pulumi:"nodeCount"`
	// A map of Kubernetes labels which should be applied to nodes in this Node Pool. Changing this forces a new resource to be created.
	NodeLabels map[string]string `pulumi:"nodeLabels"`
	// A list of Kubernetes taints which should be applied to nodes in the agent pool (e.g `key=value:NoSchedule`). Changing this forces a new resource to be created.
	NodeTaints []string `pulumi:"nodeTaints"`
	// Version of Kubernetes used for the Agents. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade)
	OrchestratorVersion *string `pulumi:"orchestratorVersion"`
	// The Agent Operating System disk size in GB. Changing this forces a new resource to be created.
	OsDiskSizeGb *int `pulumi:"osDiskSizeGb"`
	// The Operating System which should be used for this Node Pool. Changing this forces a new resource to be created. Possible values are `Linux` and `Windows`. Defaults to `Linux`.
	OsType *string `pulumi:"osType"`
	// The Priority for Virtual Machines within the Virtual Machine Scale Set that powers this Node Pool. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
	Priority *string `pulumi:"priority"`
	// The maximum price you're willing to pay in USD per Virtual Machine. Valid values are `-1` (the current on-demand price for a Virtual Machine) or a positive value with up to five decimal places. Changing this forces a new resource to be created.
	SpotMaxPrice *float64 `pulumi:"spotMaxPrice"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The SKU which should be used for the Virtual Machines used in this Node Pool. Changing this forces a new resource to be created.
	VmSize string `pulumi:"vmSize"`
	// The ID of the Subnet where this Node Pool should exist.
	VnetSubnetId *string `pulumi:"vnetSubnetId"`
}

// The set of arguments for constructing a KubernetesClusterNodePool resource.
type KubernetesClusterNodePoolArgs struct {
	// A list of Availability Zones where the Nodes in this Node Pool should be created in.
	AvailabilityZones pulumi.StringArrayInput
	// Whether to enable [auto-scaler](https://docs.microsoft.com/en-us/azure/aks/cluster-autoscaler). Defaults to `false`.
	EnableAutoScaling pulumi.BoolPtrInput
	// Should each node have a Public IP Address? Defaults to `false`.
	EnableNodePublicIp pulumi.BoolPtrInput
	// The Eviction Policy which should be used for Virtual Machines within the Virtual Machine Scale Set powering this Node Pool. Possible values are `Deallocate` and `Delete`. Changing this forces a new resource to be created.
	EvictionPolicy pulumi.StringPtrInput
	// The ID of the Kubernetes Cluster where this Node Pool should exist. Changing this forces a new resource to be created.
	KubernetesClusterId pulumi.StringInput
	// The maximum number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be greater than or equal to `minCount`.
	MaxCount pulumi.IntPtrInput
	// The maximum number of pods that can run on each agent. Changing this forces a new resource to be created.
	MaxPods pulumi.IntPtrInput
	// The minimum number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be less than or equal to `maxCount`.
	MinCount pulumi.IntPtrInput
	// Should this Node Pool be used for System or User resources? Possible values are `System` and `User`. Defaults to `User`.
	Mode pulumi.StringPtrInput
	// The name of the Node Pool which should be created within the Kubernetes Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The initial number of nodes which should exist within this Node Pool. Valid values are between `0` and `100` and must be a value in the range `minCount` - `maxCount`.
	NodeCount pulumi.IntPtrInput
	// A map of Kubernetes labels which should be applied to nodes in this Node Pool. Changing this forces a new resource to be created.
	NodeLabels pulumi.StringMapInput
	// A list of Kubernetes taints which should be applied to nodes in the agent pool (e.g `key=value:NoSchedule`). Changing this forces a new resource to be created.
	NodeTaints pulumi.StringArrayInput
	// Version of Kubernetes used for the Agents. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade)
	OrchestratorVersion pulumi.StringPtrInput
	// The Agent Operating System disk size in GB. Changing this forces a new resource to be created.
	OsDiskSizeGb pulumi.IntPtrInput
	// The Operating System which should be used for this Node Pool. Changing this forces a new resource to be created. Possible values are `Linux` and `Windows`. Defaults to `Linux`.
	OsType pulumi.StringPtrInput
	// The Priority for Virtual Machines within the Virtual Machine Scale Set that powers this Node Pool. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
	Priority pulumi.StringPtrInput
	// The maximum price you're willing to pay in USD per Virtual Machine. Valid values are `-1` (the current on-demand price for a Virtual Machine) or a positive value with up to five decimal places. Changing this forces a new resource to be created.
	SpotMaxPrice pulumi.Float64PtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The SKU which should be used for the Virtual Machines used in this Node Pool. Changing this forces a new resource to be created.
	VmSize pulumi.StringInput
	// The ID of the Subnet where this Node Pool should exist.
	VnetSubnetId pulumi.StringPtrInput
}

func (KubernetesClusterNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterNodePoolArgs)(nil)).Elem()
}

type KubernetesClusterNodePoolInput interface {
	pulumi.Input

	ToKubernetesClusterNodePoolOutput() KubernetesClusterNodePoolOutput
	ToKubernetesClusterNodePoolOutputWithContext(ctx context.Context) KubernetesClusterNodePoolOutput
}

func (KubernetesClusterNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterNodePool)(nil)).Elem()
}

func (i KubernetesClusterNodePool) ToKubernetesClusterNodePoolOutput() KubernetesClusterNodePoolOutput {
	return i.ToKubernetesClusterNodePoolOutputWithContext(context.Background())
}

func (i KubernetesClusterNodePool) ToKubernetesClusterNodePoolOutputWithContext(ctx context.Context) KubernetesClusterNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterNodePoolOutput)
}

type KubernetesClusterNodePoolOutput struct {
	*pulumi.OutputState
}

func (KubernetesClusterNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterNodePoolOutput)(nil)).Elem()
}

func (o KubernetesClusterNodePoolOutput) ToKubernetesClusterNodePoolOutput() KubernetesClusterNodePoolOutput {
	return o
}

func (o KubernetesClusterNodePoolOutput) ToKubernetesClusterNodePoolOutputWithContext(ctx context.Context) KubernetesClusterNodePoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KubernetesClusterNodePoolOutput{})
}
