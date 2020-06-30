// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService
{
    /// <summary>
    /// Manages a Managed Kubernetes Cluster (also known as AKS / Azure Kubernetes Service)
    /// 
    /// ## Example Usage
    /// 
    /// This example provisions a basic Managed Kubernetes Cluster.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleKubernetesCluster = new Azure.ContainerService.KubernetesCluster("exampleKubernetesCluster", new Azure.ContainerService.KubernetesClusterArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             DnsPrefix = "exampleaks1",
    ///             DefaultNodePool = new Azure.ContainerService.Inputs.KubernetesClusterDefaultNodePoolArgs
    ///             {
    ///                 Name = "default",
    ///                 NodeCount = 1,
    ///                 VmSize = "Standard_D2_v2",
    ///             },
    ///             Identity = new Azure.ContainerService.Inputs.KubernetesClusterIdentityArgs
    ///             {
    ///                 Type = "SystemAssigned",
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "Environment", "Production" },
    ///             },
    ///         });
    ///         this.ClientCertificate = exampleKubernetesCluster.KubeConfigs.Apply(kubeConfigs =&gt; kubeConfigs[0].ClientCertificate);
    ///         this.KubeConfig = exampleKubernetesCluster.KubeConfigRaw;
    ///     }
    /// 
    ///     [Output("clientCertificate")]
    ///     public Output&lt;string&gt; ClientCertificate { get; set; }
    ///     [Output("kubeConfig")]
    ///     public Output&lt;string&gt; KubeConfig { get; set; }
    /// }
    /// ```
    /// </summary>
    public partial class KubernetesCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// A `addon_profile` block as defined below.
        /// </summary>
        [Output("addonProfile")]
        public Output<Outputs.KubernetesClusterAddonProfile> AddonProfile { get; private set; } = null!;

        /// <summary>
        /// The IP ranges to whitelist for incoming traffic to the masters.
        /// </summary>
        [Output("apiServerAuthorizedIpRanges")]
        public Output<ImmutableArray<string>> ApiServerAuthorizedIpRanges { get; private set; } = null!;

        /// <summary>
        /// A `auto_scaler_profile` block as defined below.
        /// </summary>
        [Output("autoScalerProfile")]
        public Output<Outputs.KubernetesClusterAutoScalerProfile> AutoScalerProfile { get; private set; } = null!;

        /// <summary>
        /// A `default_node_pool` block as defined below.
        /// </summary>
        [Output("defaultNodePool")]
        public Output<Outputs.KubernetesClusterDefaultNodePool> DefaultNodePool { get; private set; } = null!;

        /// <summary>
        /// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
        /// </summary>
        [Output("diskEncryptionSetId")]
        public Output<string?> DiskEncryptionSetId { get; private set; } = null!;

        /// <summary>
        /// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("dnsPrefix")]
        public Output<string> DnsPrefix { get; private set; } = null!;

        /// <summary>
        /// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
        /// </summary>
        [Output("enablePodSecurityPolicy")]
        public Output<bool?> EnablePodSecurityPolicy { get; private set; } = null!;

        /// <summary>
        /// The FQDN of the Azure Kubernetes Managed Cluster.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.KubernetesClusterIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
        /// </summary>
        [Output("kubeAdminConfigRaw")]
        public Output<string> KubeAdminConfigRaw { get; private set; } = null!;

        /// <summary>
        /// A `kube_admin_config` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
        /// </summary>
        [Output("kubeAdminConfigs")]
        public Output<ImmutableArray<Outputs.KubernetesClusterKubeAdminConfig>> KubeAdminConfigs { get; private set; } = null!;

        /// <summary>
        /// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
        /// </summary>
        [Output("kubeConfigRaw")]
        public Output<string> KubeConfigRaw { get; private set; } = null!;

        /// <summary>
        /// A `kube_config` block as defined below.
        /// </summary>
        [Output("kubeConfigs")]
        public Output<ImmutableArray<Outputs.KubernetesClusterKubeConfig>> KubeConfigs { get; private set; } = null!;

        /// <summary>
        /// A `kubelet_identity` block as defined below.
        /// </summary>
        [Output("kubeletIdentities")]
        public Output<ImmutableArray<Outputs.KubernetesClusterKubeletIdentity>> KubeletIdentities { get; private set; } = null!;

        /// <summary>
        /// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
        /// </summary>
        [Output("kubernetesVersion")]
        public Output<string> KubernetesVersion { get; private set; } = null!;

        /// <summary>
        /// A `linux_profile` block as defined below.
        /// </summary>
        [Output("linuxProfile")]
        public Output<Outputs.KubernetesClusterLinuxProfile?> LinuxProfile { get; private set; } = null!;

        /// <summary>
        /// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `network_profile` block as defined below.
        /// </summary>
        [Output("networkProfile")]
        public Output<Outputs.KubernetesClusterNetworkProfile> NetworkProfile { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("nodeResourceGroup")]
        public Output<string> NodeResourceGroup { get; private set; } = null!;

        /// <summary>
        /// Should this Kubernetes Cluster have it's API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("privateClusterEnabled")]
        public Output<bool> PrivateClusterEnabled { get; private set; } = null!;

        /// <summary>
        /// The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
        /// </summary>
        [Output("privateFqdn")]
        public Output<string> PrivateFqdn { get; private set; } = null!;

        [Output("privateLinkEnabled")]
        public Output<bool> PrivateLinkEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `role_based_access_control` block. Changing this forces a new resource to be created.
        /// </summary>
        [Output("roleBasedAccessControl")]
        public Output<Outputs.KubernetesClusterRoleBasedAccessControl> RoleBasedAccessControl { get; private set; } = null!;

        /// <summary>
        /// A `service_principal` block as documented below.
        /// </summary>
        [Output("servicePrincipal")]
        public Output<Outputs.KubernetesClusterServicePrincipal?> ServicePrincipal { get; private set; } = null!;

        /// <summary>
        /// The SKU Tier that should be used for this Kubernetes Cluster. Changing this forces a new resource to be created. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
        /// </summary>
        [Output("skuTier")]
        public Output<string?> SkuTier { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A `windows_profile` block as defined below.
        /// </summary>
        [Output("windowsProfile")]
        public Output<Outputs.KubernetesClusterWindowsProfile> WindowsProfile { get; private set; } = null!;


        /// <summary>
        /// Create a KubernetesCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubernetesCluster(string name, KubernetesClusterArgs args, CustomResourceOptions? options = null)
            : base("azure:containerservice/kubernetesCluster:KubernetesCluster", name, args ?? new KubernetesClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KubernetesCluster(string name, Input<string> id, KubernetesClusterState? state = null, CustomResourceOptions? options = null)
            : base("azure:containerservice/kubernetesCluster:KubernetesCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KubernetesCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubernetesCluster Get(string name, Input<string> id, KubernetesClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new KubernetesCluster(name, id, state, options);
        }
    }

    public sealed class KubernetesClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `addon_profile` block as defined below.
        /// </summary>
        [Input("addonProfile")]
        public Input<Inputs.KubernetesClusterAddonProfileArgs>? AddonProfile { get; set; }

        [Input("apiServerAuthorizedIpRanges")]
        private InputList<string>? _apiServerAuthorizedIpRanges;

        /// <summary>
        /// The IP ranges to whitelist for incoming traffic to the masters.
        /// </summary>
        public InputList<string> ApiServerAuthorizedIpRanges
        {
            get => _apiServerAuthorizedIpRanges ?? (_apiServerAuthorizedIpRanges = new InputList<string>());
            set => _apiServerAuthorizedIpRanges = value;
        }

        /// <summary>
        /// A `auto_scaler_profile` block as defined below.
        /// </summary>
        [Input("autoScalerProfile")]
        public Input<Inputs.KubernetesClusterAutoScalerProfileArgs>? AutoScalerProfile { get; set; }

        /// <summary>
        /// A `default_node_pool` block as defined below.
        /// </summary>
        [Input("defaultNodePool", required: true)]
        public Input<Inputs.KubernetesClusterDefaultNodePoolArgs> DefaultNodePool { get; set; } = null!;

        /// <summary>
        /// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
        /// </summary>
        [Input("diskEncryptionSetId")]
        public Input<string>? DiskEncryptionSetId { get; set; }

        /// <summary>
        /// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dnsPrefix", required: true)]
        public Input<string> DnsPrefix { get; set; } = null!;

        /// <summary>
        /// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
        /// </summary>
        [Input("enablePodSecurityPolicy")]
        public Input<bool>? EnablePodSecurityPolicy { get; set; }

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.KubernetesClusterIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
        /// </summary>
        [Input("kubernetesVersion")]
        public Input<string>? KubernetesVersion { get; set; }

        /// <summary>
        /// A `linux_profile` block as defined below.
        /// </summary>
        [Input("linuxProfile")]
        public Input<Inputs.KubernetesClusterLinuxProfileArgs>? LinuxProfile { get; set; }

        /// <summary>
        /// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `network_profile` block as defined below.
        /// </summary>
        [Input("networkProfile")]
        public Input<Inputs.KubernetesClusterNetworkProfileArgs>? NetworkProfile { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("nodeResourceGroup")]
        public Input<string>? NodeResourceGroup { get; set; }

        /// <summary>
        /// Should this Kubernetes Cluster have it's API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("privateClusterEnabled")]
        public Input<bool>? PrivateClusterEnabled { get; set; }

        [Input("privateLinkEnabled")]
        public Input<bool>? PrivateLinkEnabled { get; set; }

        /// <summary>
        /// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `role_based_access_control` block. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleBasedAccessControl")]
        public Input<Inputs.KubernetesClusterRoleBasedAccessControlArgs>? RoleBasedAccessControl { get; set; }

        /// <summary>
        /// A `service_principal` block as documented below.
        /// </summary>
        [Input("servicePrincipal")]
        public Input<Inputs.KubernetesClusterServicePrincipalArgs>? ServicePrincipal { get; set; }

        /// <summary>
        /// The SKU Tier that should be used for this Kubernetes Cluster. Changing this forces a new resource to be created. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
        /// </summary>
        [Input("skuTier")]
        public Input<string>? SkuTier { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A `windows_profile` block as defined below.
        /// </summary>
        [Input("windowsProfile")]
        public Input<Inputs.KubernetesClusterWindowsProfileArgs>? WindowsProfile { get; set; }

        public KubernetesClusterArgs()
        {
        }
    }

    public sealed class KubernetesClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `addon_profile` block as defined below.
        /// </summary>
        [Input("addonProfile")]
        public Input<Inputs.KubernetesClusterAddonProfileGetArgs>? AddonProfile { get; set; }

        [Input("apiServerAuthorizedIpRanges")]
        private InputList<string>? _apiServerAuthorizedIpRanges;

        /// <summary>
        /// The IP ranges to whitelist for incoming traffic to the masters.
        /// </summary>
        public InputList<string> ApiServerAuthorizedIpRanges
        {
            get => _apiServerAuthorizedIpRanges ?? (_apiServerAuthorizedIpRanges = new InputList<string>());
            set => _apiServerAuthorizedIpRanges = value;
        }

        /// <summary>
        /// A `auto_scaler_profile` block as defined below.
        /// </summary>
        [Input("autoScalerProfile")]
        public Input<Inputs.KubernetesClusterAutoScalerProfileGetArgs>? AutoScalerProfile { get; set; }

        /// <summary>
        /// A `default_node_pool` block as defined below.
        /// </summary>
        [Input("defaultNodePool")]
        public Input<Inputs.KubernetesClusterDefaultNodePoolGetArgs>? DefaultNodePool { get; set; }

        /// <summary>
        /// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
        /// </summary>
        [Input("diskEncryptionSetId")]
        public Input<string>? DiskEncryptionSetId { get; set; }

        /// <summary>
        /// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dnsPrefix")]
        public Input<string>? DnsPrefix { get; set; }

        /// <summary>
        /// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
        /// </summary>
        [Input("enablePodSecurityPolicy")]
        public Input<bool>? EnablePodSecurityPolicy { get; set; }

        /// <summary>
        /// The FQDN of the Azure Kubernetes Managed Cluster.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.KubernetesClusterIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
        /// </summary>
        [Input("kubeAdminConfigRaw")]
        public Input<string>? KubeAdminConfigRaw { get; set; }

        [Input("kubeAdminConfigs")]
        private InputList<Inputs.KubernetesClusterKubeAdminConfigGetArgs>? _kubeAdminConfigs;

        /// <summary>
        /// A `kube_admin_config` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
        /// </summary>
        public InputList<Inputs.KubernetesClusterKubeAdminConfigGetArgs> KubeAdminConfigs
        {
            get => _kubeAdminConfigs ?? (_kubeAdminConfigs = new InputList<Inputs.KubernetesClusterKubeAdminConfigGetArgs>());
            set => _kubeAdminConfigs = value;
        }

        /// <summary>
        /// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
        /// </summary>
        [Input("kubeConfigRaw")]
        public Input<string>? KubeConfigRaw { get; set; }

        [Input("kubeConfigs")]
        private InputList<Inputs.KubernetesClusterKubeConfigGetArgs>? _kubeConfigs;

        /// <summary>
        /// A `kube_config` block as defined below.
        /// </summary>
        public InputList<Inputs.KubernetesClusterKubeConfigGetArgs> KubeConfigs
        {
            get => _kubeConfigs ?? (_kubeConfigs = new InputList<Inputs.KubernetesClusterKubeConfigGetArgs>());
            set => _kubeConfigs = value;
        }

        [Input("kubeletIdentities")]
        private InputList<Inputs.KubernetesClusterKubeletIdentityGetArgs>? _kubeletIdentities;

        /// <summary>
        /// A `kubelet_identity` block as defined below.
        /// </summary>
        public InputList<Inputs.KubernetesClusterKubeletIdentityGetArgs> KubeletIdentities
        {
            get => _kubeletIdentities ?? (_kubeletIdentities = new InputList<Inputs.KubernetesClusterKubeletIdentityGetArgs>());
            set => _kubeletIdentities = value;
        }

        /// <summary>
        /// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
        /// </summary>
        [Input("kubernetesVersion")]
        public Input<string>? KubernetesVersion { get; set; }

        /// <summary>
        /// A `linux_profile` block as defined below.
        /// </summary>
        [Input("linuxProfile")]
        public Input<Inputs.KubernetesClusterLinuxProfileGetArgs>? LinuxProfile { get; set; }

        /// <summary>
        /// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `network_profile` block as defined below.
        /// </summary>
        [Input("networkProfile")]
        public Input<Inputs.KubernetesClusterNetworkProfileGetArgs>? NetworkProfile { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("nodeResourceGroup")]
        public Input<string>? NodeResourceGroup { get; set; }

        /// <summary>
        /// Should this Kubernetes Cluster have it's API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("privateClusterEnabled")]
        public Input<bool>? PrivateClusterEnabled { get; set; }

        /// <summary>
        /// The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
        /// </summary>
        [Input("privateFqdn")]
        public Input<string>? PrivateFqdn { get; set; }

        [Input("privateLinkEnabled")]
        public Input<bool>? PrivateLinkEnabled { get; set; }

        /// <summary>
        /// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `role_based_access_control` block. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleBasedAccessControl")]
        public Input<Inputs.KubernetesClusterRoleBasedAccessControlGetArgs>? RoleBasedAccessControl { get; set; }

        /// <summary>
        /// A `service_principal` block as documented below.
        /// </summary>
        [Input("servicePrincipal")]
        public Input<Inputs.KubernetesClusterServicePrincipalGetArgs>? ServicePrincipal { get; set; }

        /// <summary>
        /// The SKU Tier that should be used for this Kubernetes Cluster. Changing this forces a new resource to be created. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
        /// </summary>
        [Input("skuTier")]
        public Input<string>? SkuTier { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A `windows_profile` block as defined below.
        /// </summary>
        [Input("windowsProfile")]
        public Input<Inputs.KubernetesClusterWindowsProfileGetArgs>? WindowsProfile { get; set; }

        public KubernetesClusterState()
        {
        }
    }
}
