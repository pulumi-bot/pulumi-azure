# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProductPolicyArgs', 'ProductPolicy']

@pulumi.input_type
class ProductPolicyArgs:
    def __init__(__self__, *,
                 api_management_name: pulumi.Input[str],
                 product_id: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProductPolicy resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] product_id: The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        pulumi.set(__self__, "api_management_name", api_management_name)
        pulumi.set(__self__, "product_id", product_id)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if xml_content is not None:
            pulumi.set(__self__, "xml_content", xml_content)
        if xml_link is not None:
            pulumi.set(__self__, "xml_link", xml_link)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Input[str]:
        """
        The name of the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @api_management_name.setter
    def api_management_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_management_name", value)

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> pulumi.Input[str]:
        """
        The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "product_id")

    @product_id.setter
    def product_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "product_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="xmlContent")
    def xml_content(self) -> Optional[pulumi.Input[str]]:
        """
        The XML Content for this Policy.
        """
        return pulumi.get(self, "xml_content")

    @xml_content.setter
    def xml_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "xml_content", value)

    @property
    @pulumi.getter(name="xmlLink")
    def xml_link(self) -> Optional[pulumi.Input[str]]:
        """
        A link to a Policy XML Document, which must be publicly available.
        """
        return pulumi.get(self, "xml_link")

    @xml_link.setter
    def xml_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "xml_link", value)


@pulumi.input_type
class _ProductPolicyState:
    def __init__(__self__, *,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProductPolicy resources.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] product_id: The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        if api_management_name is not None:
            pulumi.set(__self__, "api_management_name", api_management_name)
        if product_id is not None:
            pulumi.set(__self__, "product_id", product_id)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if xml_content is not None:
            pulumi.set(__self__, "xml_content", xml_content)
        if xml_link is not None:
            pulumi.set(__self__, "xml_link", xml_link)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @api_management_name.setter
    def api_management_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_management_name", value)

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "product_id")

    @product_id.setter
    def product_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_id", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="xmlContent")
    def xml_content(self) -> Optional[pulumi.Input[str]]:
        """
        The XML Content for this Policy.
        """
        return pulumi.get(self, "xml_content")

    @xml_content.setter
    def xml_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "xml_content", value)

    @property
    @pulumi.getter(name="xmlLink")
    def xml_link(self) -> Optional[pulumi.Input[str]]:
        """
        A link to a Policy XML Document, which must be publicly available.
        """
        return pulumi.get(self, "xml_link")

    @xml_link.setter
    def xml_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "xml_link", value)


class ProductPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an API Management Product Policy

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_product = azure.apimanagement.get_product(product_id="my-product",
            api_management_name="example-apim",
            resource_group_name="search-service")
        example_product_policy = azure.apimanagement.ProductPolicy("exampleProductPolicy",
            product_id=example_product.product_id,
            api_management_name=example_product.api_management_name,
            resource_group_name=example_product.resource_group_name,
            xml_content=\"\"\"<policies>
          <inbound>
            <find-and-replace from="xyz" to="abc" />
          </inbound>
        </policies>
        \"\"\")
        ```

        ## Import

        API Management Product Policy can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/productPolicy:ProductPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/products/exampleId/policies/policy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] product_id: The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProductPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an API Management Product Policy

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_product = azure.apimanagement.get_product(product_id="my-product",
            api_management_name="example-apim",
            resource_group_name="search-service")
        example_product_policy = azure.apimanagement.ProductPolicy("exampleProductPolicy",
            product_id=example_product.product_id,
            api_management_name=example_product.api_management_name,
            resource_group_name=example_product.resource_group_name,
            xml_content=\"\"\"<policies>
          <inbound>
            <find-and-replace from="xyz" to="abc" />
          </inbound>
        </policies>
        \"\"\")
        ```

        ## Import

        API Management Product Policy can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/productPolicy:ProductPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/products/exampleId/policies/policy
        ```

        :param str resource_name: The name of the resource.
        :param ProductPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProductPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProductPolicyArgs.__new__(ProductPolicyArgs)

            if api_management_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_management_name'")
            __props__.__dict__["api_management_name"] = api_management_name
            if product_id is None and not opts.urn:
                raise TypeError("Missing required property 'product_id'")
            __props__.__dict__["product_id"] = product_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["xml_content"] = xml_content
            __props__.__dict__["xml_link"] = xml_link
        super(ProductPolicy, __self__).__init__(
            'azure:apimanagement/productPolicy:ProductPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            product_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            xml_content: Optional[pulumi.Input[str]] = None,
            xml_link: Optional[pulumi.Input[str]] = None) -> 'ProductPolicy':
        """
        Get an existing ProductPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] product_id: The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProductPolicyState.__new__(_ProductPolicyState)

        __props__.__dict__["api_management_name"] = api_management_name
        __props__.__dict__["product_id"] = product_id
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["xml_content"] = xml_content
        __props__.__dict__["xml_link"] = xml_link
        return ProductPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Output[str]:
        """
        The name of the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> pulumi.Output[str]:
        """
        The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "product_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="xmlContent")
    def xml_content(self) -> pulumi.Output[str]:
        """
        The XML Content for this Policy.
        """
        return pulumi.get(self, "xml_content")

    @property
    @pulumi.getter(name="xmlLink")
    def xml_link(self) -> pulumi.Output[Optional[str]]:
        """
        A link to a Policy XML Document, which must be publicly available.
        """
        return pulumi.get(self, "xml_link")

