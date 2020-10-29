// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Managed Kubernetes Cluster (also known as AKS / Azure Kubernetes Service)
type KubernetesCluster struct {
	pulumi.CustomResourceState

	// A `addonProfile` block as defined below.
	AddonProfile KubernetesClusterAddonProfileOutput `pulumi:"addonProfile"`
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges pulumi.StringArrayOutput `pulumi:"apiServerAuthorizedIpRanges"`
	// A `autoScalerProfile` block as defined below.
	AutoScalerProfile KubernetesClusterAutoScalerProfileOutput `pulumi:"autoScalerProfile"`
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePoolOutput `pulumi:"defaultNodePool"`
	// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
	DiskEncryptionSetId pulumi.StringPtrOutput `pulumi:"diskEncryptionSetId"`
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix pulumi.StringOutput `pulumi:"dnsPrefix"`
	// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
	EnablePodSecurityPolicy pulumi.BoolPtrOutput `pulumi:"enablePodSecurityPolicy"`
	// The FQDN of the Azure Kubernetes Managed Cluster.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// A `identity` block as defined below. Changing this forces a new resource to be created.
	Identity KubernetesClusterIdentityPtrOutput `pulumi:"identity"`
	// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigRaw pulumi.StringOutput `pulumi:"kubeAdminConfigRaw"`
	// A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigs KubernetesClusterKubeAdminConfigArrayOutput `pulumi:"kubeAdminConfigs"`
	// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
	KubeConfigRaw pulumi.StringOutput `pulumi:"kubeConfigRaw"`
	// A `kubeConfig` block as defined below.
	KubeConfigs KubernetesClusterKubeConfigArrayOutput `pulumi:"kubeConfigs"`
	// A `kubeletIdentity` block as defined below.
	KubeletIdentities KubernetesClusterKubeletIdentityArrayOutput `pulumi:"kubeletIdentities"`
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// A `linuxProfile` block as defined below.
	LinuxProfile KubernetesClusterLinuxProfilePtrOutput `pulumi:"linuxProfile"`
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile KubernetesClusterNetworkProfileOutput `pulumi:"networkProfile"`
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup pulumi.StringOutput `pulumi:"nodeResourceGroup"`
	// Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
	PrivateClusterEnabled pulumi.BoolOutput `pulumi:"privateClusterEnabled"`
	// The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
	PrivateFqdn pulumi.StringOutput `pulumi:"privateFqdn"`
	// Deprecated: Deprecated in favor of `private_cluster_enabled`
	PrivateLinkEnabled pulumi.BoolOutput `pulumi:"privateLinkEnabled"`
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
	RoleBasedAccessControl KubernetesClusterRoleBasedAccessControlOutput `pulumi:"roleBasedAccessControl"`
	// A `servicePrincipal` block as documented below.
	ServicePrincipal KubernetesClusterServicePrincipalPtrOutput `pulumi:"servicePrincipal"`
	// The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
	SkuTier pulumi.StringPtrOutput `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `windowsProfile` block as defined below.
	WindowsProfile KubernetesClusterWindowsProfileOutput `pulumi:"windowsProfile"`
}

// NewKubernetesCluster registers a new resource with the given unique name, arguments, and options.
func NewKubernetesCluster(ctx *pulumi.Context,
	name string, args *KubernetesClusterArgs, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	if args == nil || args.DefaultNodePool == nil {
		return nil, errors.New("missing required argument 'DefaultNodePool'")
	}
	if args == nil || args.DnsPrefix == nil {
		return nil, errors.New("missing required argument 'DnsPrefix'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &KubernetesClusterArgs{}
	}
	var resource KubernetesCluster
	err := ctx.RegisterResource("azure:containerservice/kubernetesCluster:KubernetesCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesCluster gets an existing KubernetesCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesClusterState, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	var resource KubernetesCluster
	err := ctx.ReadResource("azure:containerservice/kubernetesCluster:KubernetesCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesCluster resources.
type kubernetesClusterState struct {
	// A `addonProfile` block as defined below.
	AddonProfile *KubernetesClusterAddonProfile `pulumi:"addonProfile"`
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges []string `pulumi:"apiServerAuthorizedIpRanges"`
	// A `autoScalerProfile` block as defined below.
	AutoScalerProfile *KubernetesClusterAutoScalerProfile `pulumi:"autoScalerProfile"`
	// A `defaultNodePool` block as defined below.
	DefaultNodePool *KubernetesClusterDefaultNodePool `pulumi:"defaultNodePool"`
	// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
	EnablePodSecurityPolicy *bool `pulumi:"enablePodSecurityPolicy"`
	// The FQDN of the Azure Kubernetes Managed Cluster.
	Fqdn *string `pulumi:"fqdn"`
	// A `identity` block as defined below. Changing this forces a new resource to be created.
	Identity *KubernetesClusterIdentity `pulumi:"identity"`
	// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigRaw *string `pulumi:"kubeAdminConfigRaw"`
	// A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigs []KubernetesClusterKubeAdminConfig `pulumi:"kubeAdminConfigs"`
	// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
	KubeConfigRaw *string `pulumi:"kubeConfigRaw"`
	// A `kubeConfig` block as defined below.
	KubeConfigs []KubernetesClusterKubeConfig `pulumi:"kubeConfigs"`
	// A `kubeletIdentity` block as defined below.
	KubeletIdentities []KubernetesClusterKubeletIdentity `pulumi:"kubeletIdentities"`
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// A `linuxProfile` block as defined below.
	LinuxProfile *KubernetesClusterLinuxProfile `pulumi:"linuxProfile"`
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile *KubernetesClusterNetworkProfile `pulumi:"networkProfile"`
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
	PrivateClusterEnabled *bool `pulumi:"privateClusterEnabled"`
	// The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
	PrivateFqdn *string `pulumi:"privateFqdn"`
	// Deprecated: Deprecated in favor of `private_cluster_enabled`
	PrivateLinkEnabled *bool `pulumi:"privateLinkEnabled"`
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
	RoleBasedAccessControl *KubernetesClusterRoleBasedAccessControl `pulumi:"roleBasedAccessControl"`
	// A `servicePrincipal` block as documented below.
	ServicePrincipal *KubernetesClusterServicePrincipal `pulumi:"servicePrincipal"`
	// The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
	SkuTier *string `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `windowsProfile` block as defined below.
	WindowsProfile *KubernetesClusterWindowsProfile `pulumi:"windowsProfile"`
}

type KubernetesClusterState struct {
	// A `addonProfile` block as defined below.
	AddonProfile KubernetesClusterAddonProfilePtrInput
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges pulumi.StringArrayInput
	// A `autoScalerProfile` block as defined below.
	AutoScalerProfile KubernetesClusterAutoScalerProfilePtrInput
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePoolPtrInput
	// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
	DiskEncryptionSetId pulumi.StringPtrInput
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix pulumi.StringPtrInput
	// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// The FQDN of the Azure Kubernetes Managed Cluster.
	Fqdn pulumi.StringPtrInput
	// A `identity` block as defined below. Changing this forces a new resource to be created.
	Identity KubernetesClusterIdentityPtrInput
	// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigRaw pulumi.StringPtrInput
	// A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigs KubernetesClusterKubeAdminConfigArrayInput
	// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
	KubeConfigRaw pulumi.StringPtrInput
	// A `kubeConfig` block as defined below.
	KubeConfigs KubernetesClusterKubeConfigArrayInput
	// A `kubeletIdentity` block as defined below.
	KubeletIdentities KubernetesClusterKubeletIdentityArrayInput
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion pulumi.StringPtrInput
	// A `linuxProfile` block as defined below.
	LinuxProfile KubernetesClusterLinuxProfilePtrInput
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkProfile` block as defined below.
	NetworkProfile KubernetesClusterNetworkProfilePtrInput
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup pulumi.StringPtrInput
	// Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
	PrivateClusterEnabled pulumi.BoolPtrInput
	// The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
	PrivateFqdn pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of `private_cluster_enabled`
	PrivateLinkEnabled pulumi.BoolPtrInput
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
	RoleBasedAccessControl KubernetesClusterRoleBasedAccessControlPtrInput
	// A `servicePrincipal` block as documented below.
	ServicePrincipal KubernetesClusterServicePrincipalPtrInput
	// The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
	SkuTier pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `windowsProfile` block as defined below.
	WindowsProfile KubernetesClusterWindowsProfilePtrInput
}

func (KubernetesClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterState)(nil)).Elem()
}

type kubernetesClusterArgs struct {
	// A `addonProfile` block as defined below.
	AddonProfile *KubernetesClusterAddonProfile `pulumi:"addonProfile"`
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges []string `pulumi:"apiServerAuthorizedIpRanges"`
	// A `autoScalerProfile` block as defined below.
	AutoScalerProfile *KubernetesClusterAutoScalerProfile `pulumi:"autoScalerProfile"`
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePool `pulumi:"defaultNodePool"`
	// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix string `pulumi:"dnsPrefix"`
	// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
	EnablePodSecurityPolicy *bool `pulumi:"enablePodSecurityPolicy"`
	// A `identity` block as defined below. Changing this forces a new resource to be created.
	Identity *KubernetesClusterIdentity `pulumi:"identity"`
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// A `linuxProfile` block as defined below.
	LinuxProfile *KubernetesClusterLinuxProfile `pulumi:"linuxProfile"`
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile *KubernetesClusterNetworkProfile `pulumi:"networkProfile"`
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
	PrivateClusterEnabled *bool `pulumi:"privateClusterEnabled"`
	// Deprecated: Deprecated in favor of `private_cluster_enabled`
	PrivateLinkEnabled *bool `pulumi:"privateLinkEnabled"`
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
	RoleBasedAccessControl *KubernetesClusterRoleBasedAccessControl `pulumi:"roleBasedAccessControl"`
	// A `servicePrincipal` block as documented below.
	ServicePrincipal *KubernetesClusterServicePrincipal `pulumi:"servicePrincipal"`
	// The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
	SkuTier *string `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `windowsProfile` block as defined below.
	WindowsProfile *KubernetesClusterWindowsProfile `pulumi:"windowsProfile"`
}

// The set of arguments for constructing a KubernetesCluster resource.
type KubernetesClusterArgs struct {
	// A `addonProfile` block as defined below.
	AddonProfile KubernetesClusterAddonProfilePtrInput
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges pulumi.StringArrayInput
	// A `autoScalerProfile` block as defined below.
	AutoScalerProfile KubernetesClusterAutoScalerProfilePtrInput
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePoolInput
	// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
	DiskEncryptionSetId pulumi.StringPtrInput
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix pulumi.StringInput
	// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// A `identity` block as defined below. Changing this forces a new resource to be created.
	Identity KubernetesClusterIdentityPtrInput
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion pulumi.StringPtrInput
	// A `linuxProfile` block as defined below.
	LinuxProfile KubernetesClusterLinuxProfilePtrInput
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkProfile` block as defined below.
	NetworkProfile KubernetesClusterNetworkProfilePtrInput
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup pulumi.StringPtrInput
	// Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
	PrivateClusterEnabled pulumi.BoolPtrInput
	// Deprecated: Deprecated in favor of `private_cluster_enabled`
	PrivateLinkEnabled pulumi.BoolPtrInput
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
	RoleBasedAccessControl KubernetesClusterRoleBasedAccessControlPtrInput
	// A `servicePrincipal` block as documented below.
	ServicePrincipal KubernetesClusterServicePrincipalPtrInput
	// The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
	SkuTier pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `windowsProfile` block as defined below.
	WindowsProfile KubernetesClusterWindowsProfilePtrInput
}

func (KubernetesClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterArgs)(nil)).Elem()
}
