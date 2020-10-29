// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) Cluster
type Cluster struct {
	pulumi.CustomResourceState

	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionUri pulumi.StringOutput `pulumi:"dataIngestionUri"`
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrOutput `pulumi:"enableDiskEncryption"`
	// Specifies if the purge operations are enabled.
	EnablePurge pulumi.BoolPtrOutput `pulumi:"enablePurge"`
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrOutput `pulumi:"enableStreamingIngest"`
	// A identity block.
	Identity ClusterIdentityOutput `pulumi:"identity"`
	// An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
	LanguageExtensions pulumi.StringArrayOutput `pulumi:"languageExtensions"`
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// An `optimizedAutoScale` block as defined below.
	OptimizedAutoScale ClusterOptimizedAutoScalePtrOutput `pulumi:"optimizedAutoScale"`
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku ClusterSkuOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies a list of tenant IDs that are trusted by the cluster.
	TrustedExternalTenants pulumi.StringArrayOutput `pulumi:"trustedExternalTenants"`
	// The FQDN of the Azure Kusto Cluster.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// A `virtualNetworkConfiguration` block as defined below.
	VirtualNetworkConfiguration ClusterVirtualNetworkConfigurationPtrOutput `pulumi:"virtualNetworkConfiguration"`
	// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("azure:kusto/cluster:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:kusto/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionUri *string `pulumi:"dataIngestionUri"`
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption *bool `pulumi:"enableDiskEncryption"`
	// Specifies if the purge operations are enabled.
	EnablePurge *bool `pulumi:"enablePurge"`
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest *bool `pulumi:"enableStreamingIngest"`
	// A identity block.
	Identity *ClusterIdentity `pulumi:"identity"`
	// An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
	LanguageExtensions []string `pulumi:"languageExtensions"`
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// An `optimizedAutoScale` block as defined below.
	OptimizedAutoScale *ClusterOptimizedAutoScale `pulumi:"optimizedAutoScale"`
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku *ClusterSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies a list of tenant IDs that are trusted by the cluster.
	TrustedExternalTenants []string `pulumi:"trustedExternalTenants"`
	// The FQDN of the Azure Kusto Cluster.
	Uri *string `pulumi:"uri"`
	// A `virtualNetworkConfiguration` block as defined below.
	VirtualNetworkConfiguration *ClusterVirtualNetworkConfiguration `pulumi:"virtualNetworkConfiguration"`
	// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
	Zones []string `pulumi:"zones"`
}

type ClusterState struct {
	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionUri pulumi.StringPtrInput
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrInput
	// Specifies if the purge operations are enabled.
	EnablePurge pulumi.BoolPtrInput
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrInput
	// A identity block.
	Identity ClusterIdentityPtrInput
	// An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
	LanguageExtensions pulumi.StringArrayInput
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// An `optimizedAutoScale` block as defined below.
	OptimizedAutoScale ClusterOptimizedAutoScalePtrInput
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `sku` block as defined below.
	Sku ClusterSkuPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies a list of tenant IDs that are trusted by the cluster.
	TrustedExternalTenants pulumi.StringArrayInput
	// The FQDN of the Azure Kusto Cluster.
	Uri pulumi.StringPtrInput
	// A `virtualNetworkConfiguration` block as defined below.
	VirtualNetworkConfiguration ClusterVirtualNetworkConfigurationPtrInput
	// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringArrayInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption *bool `pulumi:"enableDiskEncryption"`
	// Specifies if the purge operations are enabled.
	EnablePurge *bool `pulumi:"enablePurge"`
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest *bool `pulumi:"enableStreamingIngest"`
	// A identity block.
	Identity *ClusterIdentity `pulumi:"identity"`
	// An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
	LanguageExtensions []string `pulumi:"languageExtensions"`
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// An `optimizedAutoScale` block as defined below.
	OptimizedAutoScale *ClusterOptimizedAutoScale `pulumi:"optimizedAutoScale"`
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku ClusterSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies a list of tenant IDs that are trusted by the cluster.
	TrustedExternalTenants []string `pulumi:"trustedExternalTenants"`
	// A `virtualNetworkConfiguration` block as defined below.
	VirtualNetworkConfiguration *ClusterVirtualNetworkConfiguration `pulumi:"virtualNetworkConfiguration"`
	// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrInput
	// Specifies if the purge operations are enabled.
	EnablePurge pulumi.BoolPtrInput
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrInput
	// A identity block.
	Identity ClusterIdentityPtrInput
	// An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
	LanguageExtensions pulumi.StringArrayInput
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// An `optimizedAutoScale` block as defined below.
	OptimizedAutoScale ClusterOptimizedAutoScalePtrInput
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `sku` block as defined below.
	Sku ClusterSkuInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies a list of tenant IDs that are trusted by the cluster.
	TrustedExternalTenants pulumi.StringArrayInput
	// A `virtualNetworkConfiguration` block as defined below.
	VirtualNetworkConfiguration ClusterVirtualNetworkConfigurationPtrInput
	// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringArrayInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
