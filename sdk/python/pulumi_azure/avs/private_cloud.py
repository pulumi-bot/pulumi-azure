# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['PrivateCloudArgs', 'PrivateCloud']

@pulumi.input_type
class PrivateCloudArgs:
    def __init__(__self__, *,
                 management_cluster: pulumi.Input['PrivateCloudManagementClusterArgs'],
                 network_subnet_cidr: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 sku_name: pulumi.Input[str],
                 internet_connection_enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nsxt_password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vcenter_password: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrivateCloud resource.
        :param pulumi.Input['PrivateCloudManagementClusterArgs'] management_cluster: A `management_cluster` block as defined below.
               > **NOTE :** `internet_connection_enabled` and `management_cluster.0.size` cannot be updated at the same time.
        :param pulumi.Input[str] network_subnet_cidr: The subnet which should be unique across virtual network in your subscription as well as on-premise. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Vmware Private Cloud should exist. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] sku_name: The Name of the SKU used for this Private Cloud. Possible values are `av20`, `av36` and `av36t`. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[bool] internet_connection_enabled: Is the Private Cluster connected to the internet? This field can not updated with `management_cluster.0.size` together.
               > **NOTE :** `internet_connection_enabled` and `management_cluster.0.size` cannot be updated at the same time.
        :param pulumi.Input[str] location: The Azure Region where the Vmware Private Cloud should exist. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] name: The name which should be used for this Vmware Private Cloud. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] nsxt_password: The password of the NSX-T Manager. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Vmware Private Cloud.
        :param pulumi.Input[str] vcenter_password: The password of the vCenter admin. Changing this forces a new Vmware Private Cloud to be created.
        """
        pulumi.set(__self__, "management_cluster", management_cluster)
        pulumi.set(__self__, "network_subnet_cidr", network_subnet_cidr)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku_name", sku_name)
        if internet_connection_enabled is not None:
            pulumi.set(__self__, "internet_connection_enabled", internet_connection_enabled)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nsxt_password is not None:
            pulumi.set(__self__, "nsxt_password", nsxt_password)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vcenter_password is not None:
            pulumi.set(__self__, "vcenter_password", vcenter_password)

    @property
    @pulumi.getter(name="managementCluster")
    def management_cluster(self) -> pulumi.Input['PrivateCloudManagementClusterArgs']:
        """
        A `management_cluster` block as defined below.
        > **NOTE :** `internet_connection_enabled` and `management_cluster.0.size` cannot be updated at the same time.
        """
        return pulumi.get(self, "management_cluster")

    @management_cluster.setter
    def management_cluster(self, value: pulumi.Input['PrivateCloudManagementClusterArgs']):
        pulumi.set(self, "management_cluster", value)

    @property
    @pulumi.getter(name="networkSubnetCidr")
    def network_subnet_cidr(self) -> pulumi.Input[str]:
        """
        The subnet which should be unique across virtual network in your subscription as well as on-premise. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "network_subnet_cidr")

    @network_subnet_cidr.setter
    def network_subnet_cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_subnet_cidr", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the Vmware Private Cloud should exist. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Input[str]:
        """
        The Name of the SKU used for this Private Cloud. Possible values are `av20`, `av36` and `av36t`. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "sku_name")

    @sku_name.setter
    def sku_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "sku_name", value)

    @property
    @pulumi.getter(name="internetConnectionEnabled")
    def internet_connection_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the Private Cluster connected to the internet? This field can not updated with `management_cluster.0.size` together.
        > **NOTE :** `internet_connection_enabled` and `management_cluster.0.size` cannot be updated at the same time.
        """
        return pulumi.get(self, "internet_connection_enabled")

    @internet_connection_enabled.setter
    def internet_connection_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internet_connection_enabled", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure Region where the Vmware Private Cloud should exist. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this Vmware Private Cloud. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nsxtPassword")
    def nsxt_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the NSX-T Manager. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "nsxt_password")

    @nsxt_password.setter
    def nsxt_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nsxt_password", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags which should be assigned to the Vmware Private Cloud.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vcenterPassword")
    def vcenter_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the vCenter admin. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "vcenter_password")

    @vcenter_password.setter
    def vcenter_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vcenter_password", value)


class PrivateCloud(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 internet_connection_enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_cluster: Optional[pulumi.Input[pulumi.InputType['PrivateCloudManagementClusterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_subnet_cidr: Optional[pulumi.Input[str]] = None,
                 nsxt_password: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vcenter_password: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Vmware Private Clouds can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:avs/privateCloud:PrivateCloud example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/PrivateClouds/privateCloud1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] internet_connection_enabled: Is the Private Cluster connected to the internet? This field can not updated with `management_cluster.0.size` together.
               > **NOTE :** `internet_connection_enabled` and `management_cluster.0.size` cannot be updated at the same time.
        :param pulumi.Input[str] location: The Azure Region where the Vmware Private Cloud should exist. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[pulumi.InputType['PrivateCloudManagementClusterArgs']] management_cluster: A `management_cluster` block as defined below.
               > **NOTE :** `internet_connection_enabled` and `management_cluster.0.size` cannot be updated at the same time.
        :param pulumi.Input[str] name: The name which should be used for this Vmware Private Cloud. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] network_subnet_cidr: The subnet which should be unique across virtual network in your subscription as well as on-premise. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] nsxt_password: The password of the NSX-T Manager. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Vmware Private Cloud should exist. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] sku_name: The Name of the SKU used for this Private Cloud. Possible values are `av20`, `av36` and `av36t`. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Vmware Private Cloud.
        :param pulumi.Input[str] vcenter_password: The password of the vCenter admin. Changing this forces a new Vmware Private Cloud to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateCloudArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Vmware Private Clouds can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:avs/privateCloud:PrivateCloud example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/PrivateClouds/privateCloud1
        ```

        :param str resource_name: The name of the resource.
        :param PrivateCloudArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateCloudArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 internet_connection_enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_cluster: Optional[pulumi.Input[pulumi.InputType['PrivateCloudManagementClusterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_subnet_cidr: Optional[pulumi.Input[str]] = None,
                 nsxt_password: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vcenter_password: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['internet_connection_enabled'] = internet_connection_enabled
            __props__['location'] = location
            if management_cluster is None and not opts.urn:
                raise TypeError("Missing required property 'management_cluster'")
            __props__['management_cluster'] = management_cluster
            __props__['name'] = name
            if network_subnet_cidr is None and not opts.urn:
                raise TypeError("Missing required property 'network_subnet_cidr'")
            __props__['network_subnet_cidr'] = network_subnet_cidr
            __props__['nsxt_password'] = nsxt_password
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku_name is None and not opts.urn:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['vcenter_password'] = vcenter_password
            __props__['circuits'] = None
            __props__['hcx_cloud_manager_endpoint'] = None
            __props__['management_subnet_cidr'] = None
            __props__['nsxt_certificate_thumbprint'] = None
            __props__['nsxt_manager_endpoint'] = None
            __props__['provisioning_subnet_cidr'] = None
            __props__['vcenter_certificate_thumbprint'] = None
            __props__['vcsa_endpoint'] = None
            __props__['vmotion_subnet_cidr'] = None
        super(PrivateCloud, __self__).__init__(
            'azure:avs/privateCloud:PrivateCloud',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            circuits: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PrivateCloudCircuitArgs']]]]] = None,
            hcx_cloud_manager_endpoint: Optional[pulumi.Input[str]] = None,
            internet_connection_enabled: Optional[pulumi.Input[bool]] = None,
            location: Optional[pulumi.Input[str]] = None,
            management_cluster: Optional[pulumi.Input[pulumi.InputType['PrivateCloudManagementClusterArgs']]] = None,
            management_subnet_cidr: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_subnet_cidr: Optional[pulumi.Input[str]] = None,
            nsxt_certificate_thumbprint: Optional[pulumi.Input[str]] = None,
            nsxt_manager_endpoint: Optional[pulumi.Input[str]] = None,
            nsxt_password: Optional[pulumi.Input[str]] = None,
            provisioning_subnet_cidr: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sku_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vcenter_certificate_thumbprint: Optional[pulumi.Input[str]] = None,
            vcenter_password: Optional[pulumi.Input[str]] = None,
            vcsa_endpoint: Optional[pulumi.Input[str]] = None,
            vmotion_subnet_cidr: Optional[pulumi.Input[str]] = None) -> 'PrivateCloud':
        """
        Get an existing PrivateCloud resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PrivateCloudCircuitArgs']]]] circuits: A `circuit` block as defined below.
        :param pulumi.Input[str] hcx_cloud_manager_endpoint: The endpoint for the HCX Cloud Manager.
        :param pulumi.Input[bool] internet_connection_enabled: Is the Private Cluster connected to the internet? This field can not updated with `management_cluster.0.size` together.
               > **NOTE :** `internet_connection_enabled` and `management_cluster.0.size` cannot be updated at the same time.
        :param pulumi.Input[str] location: The Azure Region where the Vmware Private Cloud should exist. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[pulumi.InputType['PrivateCloudManagementClusterArgs']] management_cluster: A `management_cluster` block as defined below.
               > **NOTE :** `internet_connection_enabled` and `management_cluster.0.size` cannot be updated at the same time.
        :param pulumi.Input[str] management_subnet_cidr: The network used to access vCenter Server and NSX-T Manager.
        :param pulumi.Input[str] name: The name which should be used for this Vmware Private Cloud. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] network_subnet_cidr: The subnet which should be unique across virtual network in your subscription as well as on-premise. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] nsxt_certificate_thumbprint: The thumbprint of the NSX-T Manager SSL certificate.
        :param pulumi.Input[str] nsxt_manager_endpoint: The endpoint for the NSX-T Data Center manager.
        :param pulumi.Input[str] nsxt_password: The password of the NSX-T Manager. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] provisioning_subnet_cidr: The network which is used for virtual machine cold migration, cloning, and snapshot migration.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Vmware Private Cloud should exist. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] sku_name: The Name of the SKU used for this Private Cloud. Possible values are `av20`, `av36` and `av36t`. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Vmware Private Cloud.
        :param pulumi.Input[str] vcenter_certificate_thumbprint: The thumbprint of the vCenter Server SSL certificate.
        :param pulumi.Input[str] vcenter_password: The password of the vCenter admin. Changing this forces a new Vmware Private Cloud to be created.
        :param pulumi.Input[str] vcsa_endpoint: The endpoint for Virtual Center Server Appliance.
        :param pulumi.Input[str] vmotion_subnet_cidr: The network which is used for live migration of virtual machines.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["circuits"] = circuits
        __props__["hcx_cloud_manager_endpoint"] = hcx_cloud_manager_endpoint
        __props__["internet_connection_enabled"] = internet_connection_enabled
        __props__["location"] = location
        __props__["management_cluster"] = management_cluster
        __props__["management_subnet_cidr"] = management_subnet_cidr
        __props__["name"] = name
        __props__["network_subnet_cidr"] = network_subnet_cidr
        __props__["nsxt_certificate_thumbprint"] = nsxt_certificate_thumbprint
        __props__["nsxt_manager_endpoint"] = nsxt_manager_endpoint
        __props__["nsxt_password"] = nsxt_password
        __props__["provisioning_subnet_cidr"] = provisioning_subnet_cidr
        __props__["resource_group_name"] = resource_group_name
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        __props__["vcenter_certificate_thumbprint"] = vcenter_certificate_thumbprint
        __props__["vcenter_password"] = vcenter_password
        __props__["vcsa_endpoint"] = vcsa_endpoint
        __props__["vmotion_subnet_cidr"] = vmotion_subnet_cidr
        return PrivateCloud(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def circuits(self) -> pulumi.Output[Sequence['outputs.PrivateCloudCircuit']]:
        """
        A `circuit` block as defined below.
        """
        return pulumi.get(self, "circuits")

    @property
    @pulumi.getter(name="hcxCloudManagerEndpoint")
    def hcx_cloud_manager_endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint for the HCX Cloud Manager.
        """
        return pulumi.get(self, "hcx_cloud_manager_endpoint")

    @property
    @pulumi.getter(name="internetConnectionEnabled")
    def internet_connection_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Is the Private Cluster connected to the internet? This field can not updated with `management_cluster.0.size` together.
        > **NOTE :** `internet_connection_enabled` and `management_cluster.0.size` cannot be updated at the same time.
        """
        return pulumi.get(self, "internet_connection_enabled")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure Region where the Vmware Private Cloud should exist. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementCluster")
    def management_cluster(self) -> pulumi.Output['outputs.PrivateCloudManagementCluster']:
        """
        A `management_cluster` block as defined below.
        > **NOTE :** `internet_connection_enabled` and `management_cluster.0.size` cannot be updated at the same time.
        """
        return pulumi.get(self, "management_cluster")

    @property
    @pulumi.getter(name="managementSubnetCidr")
    def management_subnet_cidr(self) -> pulumi.Output[str]:
        """
        The network used to access vCenter Server and NSX-T Manager.
        """
        return pulumi.get(self, "management_subnet_cidr")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Vmware Private Cloud. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkSubnetCidr")
    def network_subnet_cidr(self) -> pulumi.Output[str]:
        """
        The subnet which should be unique across virtual network in your subscription as well as on-premise. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "network_subnet_cidr")

    @property
    @pulumi.getter(name="nsxtCertificateThumbprint")
    def nsxt_certificate_thumbprint(self) -> pulumi.Output[str]:
        """
        The thumbprint of the NSX-T Manager SSL certificate.
        """
        return pulumi.get(self, "nsxt_certificate_thumbprint")

    @property
    @pulumi.getter(name="nsxtManagerEndpoint")
    def nsxt_manager_endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint for the NSX-T Data Center manager.
        """
        return pulumi.get(self, "nsxt_manager_endpoint")

    @property
    @pulumi.getter(name="nsxtPassword")
    def nsxt_password(self) -> pulumi.Output[Optional[str]]:
        """
        The password of the NSX-T Manager. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "nsxt_password")

    @property
    @pulumi.getter(name="provisioningSubnetCidr")
    def provisioning_subnet_cidr(self) -> pulumi.Output[str]:
        """
        The network which is used for virtual machine cold migration, cloning, and snapshot migration.
        """
        return pulumi.get(self, "provisioning_subnet_cidr")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Vmware Private Cloud should exist. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Output[str]:
        """
        The Name of the SKU used for this Private Cloud. Possible values are `av20`, `av36` and `av36t`. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Vmware Private Cloud.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vcenterCertificateThumbprint")
    def vcenter_certificate_thumbprint(self) -> pulumi.Output[str]:
        """
        The thumbprint of the vCenter Server SSL certificate.
        """
        return pulumi.get(self, "vcenter_certificate_thumbprint")

    @property
    @pulumi.getter(name="vcenterPassword")
    def vcenter_password(self) -> pulumi.Output[Optional[str]]:
        """
        The password of the vCenter admin. Changing this forces a new Vmware Private Cloud to be created.
        """
        return pulumi.get(self, "vcenter_password")

    @property
    @pulumi.getter(name="vcsaEndpoint")
    def vcsa_endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint for Virtual Center Server Appliance.
        """
        return pulumi.get(self, "vcsa_endpoint")

    @property
    @pulumi.getter(name="vmotionSubnetCidr")
    def vmotion_subnet_cidr(self) -> pulumi.Output[str]:
        """
        The network which is used for live migration of virtual machines.
        """
        return pulumi.get(self, "vmotion_subnet_cidr")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

