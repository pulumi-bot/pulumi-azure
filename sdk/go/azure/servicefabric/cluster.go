// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Service Fabric Cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/servicefabric"
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
// 		_, err = servicefabric.NewCluster(ctx, "exampleCluster", &servicefabric.ClusterArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			Location:           exampleResourceGroup.Location,
// 			ReliabilityLevel:   pulumi.String("Bronze"),
// 			UpgradeMode:        pulumi.String("Manual"),
// 			ClusterCodeVersion: pulumi.String("7.1.456.959"),
// 			VmImage:            pulumi.String("Windows"),
// 			ManagementEndpoint: pulumi.String("https://example:80"),
// 			NodeTypes: servicefabric.ClusterNodeTypeArray{
// 				&servicefabric.ClusterNodeTypeArgs{
// 					Name:               pulumi.String("first"),
// 					InstanceCount:      pulumi.Int(3),
// 					IsPrimary:          pulumi.Bool(true),
// 					ClientEndpointPort: pulumi.Int(2020),
// 					HttpEndpointPort:   pulumi.Int(80),
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
//
// ## Import
//
// Service Fabric Clusters can be imported using the `resource id`, e.g.
type Cluster struct {
	pulumi.CustomResourceState

	// A List of one or more features which should be enabled, such as `DnsService`.
	AddOnFeatures pulumi.StringArrayOutput `pulumi:"addOnFeatures"`
	// An `azureActiveDirectory` block as defined below.
	AzureActiveDirectory ClusterAzureActiveDirectoryPtrOutput `pulumi:"azureActiveDirectory"`
	// A `certificate` block as defined below. Conflicts with `certificateCommonNames`.
	Certificate ClusterCertificatePtrOutput `pulumi:"certificate"`
	// A `certificateCommonNames` block as defined below. Conflicts with `certificate`.
	CertificateCommonNames ClusterCertificateCommonNamesPtrOutput `pulumi:"certificateCommonNames"`
	// A `clientCertificateCommonName` block as defined below.
	ClientCertificateCommonNames ClusterClientCertificateCommonNameArrayOutput `pulumi:"clientCertificateCommonNames"`
	// One or more `clientCertificateThumbprint` blocks as defined below.
	ClientCertificateThumbprints ClusterClientCertificateThumbprintArrayOutput `pulumi:"clientCertificateThumbprints"`
	// Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
	ClusterCodeVersion pulumi.StringOutput `pulumi:"clusterCodeVersion"`
	// The Cluster Endpoint for this Service Fabric Cluster.
	ClusterEndpoint pulumi.StringOutput `pulumi:"clusterEndpoint"`
	// A `diagnosticsConfig` block as defined below. Changing this forces a new resource to be created.
	DiagnosticsConfig ClusterDiagnosticsConfigPtrOutput `pulumi:"diagnosticsConfig"`
	// One or more `fabricSettings` blocks as defined below.
	FabricSettings ClusterFabricSettingArrayOutput `pulumi:"fabricSettings"`
	// Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
	ManagementEndpoint pulumi.StringOutput `pulumi:"managementEndpoint"`
	// The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `nodeType` blocks as defined below.
	NodeTypes ClusterNodeTypeArrayOutput `pulumi:"nodeTypes"`
	// Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
	ReliabilityLevel pulumi.StringOutput `pulumi:"reliabilityLevel"`
	// The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `reverseProxyCertificate` block as defined below.
	ReverseProxyCertificate ClusterReverseProxyCertificatePtrOutput `pulumi:"reverseProxyCertificate"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
	UpgradeMode pulumi.StringOutput `pulumi:"upgradeMode"`
	// Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
	VmImage pulumi.StringOutput `pulumi:"vmImage"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.ManagementEndpoint == nil {
		return nil, errors.New("missing required argument 'ManagementEndpoint'")
	}
	if args == nil || args.NodeTypes == nil {
		return nil, errors.New("missing required argument 'NodeTypes'")
	}
	if args == nil || args.ReliabilityLevel == nil {
		return nil, errors.New("missing required argument 'ReliabilityLevel'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.UpgradeMode == nil {
		return nil, errors.New("missing required argument 'UpgradeMode'")
	}
	if args == nil || args.VmImage == nil {
		return nil, errors.New("missing required argument 'VmImage'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("azure:servicefabric/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure:servicefabric/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// A List of one or more features which should be enabled, such as `DnsService`.
	AddOnFeatures []string `pulumi:"addOnFeatures"`
	// An `azureActiveDirectory` block as defined below.
	AzureActiveDirectory *ClusterAzureActiveDirectory `pulumi:"azureActiveDirectory"`
	// A `certificate` block as defined below. Conflicts with `certificateCommonNames`.
	Certificate *ClusterCertificate `pulumi:"certificate"`
	// A `certificateCommonNames` block as defined below. Conflicts with `certificate`.
	CertificateCommonNames *ClusterCertificateCommonNames `pulumi:"certificateCommonNames"`
	// A `clientCertificateCommonName` block as defined below.
	ClientCertificateCommonNames []ClusterClientCertificateCommonName `pulumi:"clientCertificateCommonNames"`
	// One or more `clientCertificateThumbprint` blocks as defined below.
	ClientCertificateThumbprints []ClusterClientCertificateThumbprint `pulumi:"clientCertificateThumbprints"`
	// Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
	ClusterCodeVersion *string `pulumi:"clusterCodeVersion"`
	// The Cluster Endpoint for this Service Fabric Cluster.
	ClusterEndpoint *string `pulumi:"clusterEndpoint"`
	// A `diagnosticsConfig` block as defined below. Changing this forces a new resource to be created.
	DiagnosticsConfig *ClusterDiagnosticsConfig `pulumi:"diagnosticsConfig"`
	// One or more `fabricSettings` blocks as defined below.
	FabricSettings []ClusterFabricSetting `pulumi:"fabricSettings"`
	// Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
	ManagementEndpoint *string `pulumi:"managementEndpoint"`
	// The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `nodeType` blocks as defined below.
	NodeTypes []ClusterNodeType `pulumi:"nodeTypes"`
	// Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
	ReliabilityLevel *string `pulumi:"reliabilityLevel"`
	// The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `reverseProxyCertificate` block as defined below.
	ReverseProxyCertificate *ClusterReverseProxyCertificate `pulumi:"reverseProxyCertificate"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
	UpgradeMode *string `pulumi:"upgradeMode"`
	// Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
	VmImage *string `pulumi:"vmImage"`
}

type ClusterState struct {
	// A List of one or more features which should be enabled, such as `DnsService`.
	AddOnFeatures pulumi.StringArrayInput
	// An `azureActiveDirectory` block as defined below.
	AzureActiveDirectory ClusterAzureActiveDirectoryPtrInput
	// A `certificate` block as defined below. Conflicts with `certificateCommonNames`.
	Certificate ClusterCertificatePtrInput
	// A `certificateCommonNames` block as defined below. Conflicts with `certificate`.
	CertificateCommonNames ClusterCertificateCommonNamesPtrInput
	// A `clientCertificateCommonName` block as defined below.
	ClientCertificateCommonNames ClusterClientCertificateCommonNameArrayInput
	// One or more `clientCertificateThumbprint` blocks as defined below.
	ClientCertificateThumbprints ClusterClientCertificateThumbprintArrayInput
	// Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
	ClusterCodeVersion pulumi.StringPtrInput
	// The Cluster Endpoint for this Service Fabric Cluster.
	ClusterEndpoint pulumi.StringPtrInput
	// A `diagnosticsConfig` block as defined below. Changing this forces a new resource to be created.
	DiagnosticsConfig ClusterDiagnosticsConfigPtrInput
	// One or more `fabricSettings` blocks as defined below.
	FabricSettings ClusterFabricSettingArrayInput
	// Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
	ManagementEndpoint pulumi.StringPtrInput
	// The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `nodeType` blocks as defined below.
	NodeTypes ClusterNodeTypeArrayInput
	// Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
	ReliabilityLevel pulumi.StringPtrInput
	// The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `reverseProxyCertificate` block as defined below.
	ReverseProxyCertificate ClusterReverseProxyCertificatePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
	UpgradeMode pulumi.StringPtrInput
	// Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
	VmImage pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// A List of one or more features which should be enabled, such as `DnsService`.
	AddOnFeatures []string `pulumi:"addOnFeatures"`
	// An `azureActiveDirectory` block as defined below.
	AzureActiveDirectory *ClusterAzureActiveDirectory `pulumi:"azureActiveDirectory"`
	// A `certificate` block as defined below. Conflicts with `certificateCommonNames`.
	Certificate *ClusterCertificate `pulumi:"certificate"`
	// A `certificateCommonNames` block as defined below. Conflicts with `certificate`.
	CertificateCommonNames *ClusterCertificateCommonNames `pulumi:"certificateCommonNames"`
	// A `clientCertificateCommonName` block as defined below.
	ClientCertificateCommonNames []ClusterClientCertificateCommonName `pulumi:"clientCertificateCommonNames"`
	// One or more `clientCertificateThumbprint` blocks as defined below.
	ClientCertificateThumbprints []ClusterClientCertificateThumbprint `pulumi:"clientCertificateThumbprints"`
	// Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
	ClusterCodeVersion *string `pulumi:"clusterCodeVersion"`
	// A `diagnosticsConfig` block as defined below. Changing this forces a new resource to be created.
	DiagnosticsConfig *ClusterDiagnosticsConfig `pulumi:"diagnosticsConfig"`
	// One or more `fabricSettings` blocks as defined below.
	FabricSettings []ClusterFabricSetting `pulumi:"fabricSettings"`
	// Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
	ManagementEndpoint string `pulumi:"managementEndpoint"`
	// The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `nodeType` blocks as defined below.
	NodeTypes []ClusterNodeType `pulumi:"nodeTypes"`
	// Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
	ReliabilityLevel string `pulumi:"reliabilityLevel"`
	// The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `reverseProxyCertificate` block as defined below.
	ReverseProxyCertificate *ClusterReverseProxyCertificate `pulumi:"reverseProxyCertificate"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
	UpgradeMode string `pulumi:"upgradeMode"`
	// Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
	VmImage string `pulumi:"vmImage"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// A List of one or more features which should be enabled, such as `DnsService`.
	AddOnFeatures pulumi.StringArrayInput
	// An `azureActiveDirectory` block as defined below.
	AzureActiveDirectory ClusterAzureActiveDirectoryPtrInput
	// A `certificate` block as defined below. Conflicts with `certificateCommonNames`.
	Certificate ClusterCertificatePtrInput
	// A `certificateCommonNames` block as defined below. Conflicts with `certificate`.
	CertificateCommonNames ClusterCertificateCommonNamesPtrInput
	// A `clientCertificateCommonName` block as defined below.
	ClientCertificateCommonNames ClusterClientCertificateCommonNameArrayInput
	// One or more `clientCertificateThumbprint` blocks as defined below.
	ClientCertificateThumbprints ClusterClientCertificateThumbprintArrayInput
	// Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
	ClusterCodeVersion pulumi.StringPtrInput
	// A `diagnosticsConfig` block as defined below. Changing this forces a new resource to be created.
	DiagnosticsConfig ClusterDiagnosticsConfigPtrInput
	// One or more `fabricSettings` blocks as defined below.
	FabricSettings ClusterFabricSettingArrayInput
	// Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
	ManagementEndpoint pulumi.StringInput
	// The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `nodeType` blocks as defined below.
	NodeTypes ClusterNodeTypeArrayInput
	// Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
	ReliabilityLevel pulumi.StringInput
	// The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `reverseProxyCertificate` block as defined below.
	ReverseProxyCertificate ClusterReverseProxyCertificatePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
	UpgradeMode pulumi.StringInput
	// Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
	VmImage pulumi.StringInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
