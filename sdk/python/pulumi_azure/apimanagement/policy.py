# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 api_management_id: pulumi.Input[str],
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input[str] api_management_id: The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy as a string.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        pulumi.set(__self__, "api_management_id", api_management_id)
        if xml_content is not None:
            pulumi.set(__self__, "xml_content", xml_content)
        if xml_link is not None:
            pulumi.set(__self__, "xml_link", xml_link)

    @property
    @pulumi.getter(name="apiManagementId")
    def api_management_id(self) -> pulumi.Input[str]:
        """
        The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        """
        return pulumi.get(self, "api_management_id")

    @api_management_id.setter
    def api_management_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_management_id", value)

    @property
    @pulumi.getter(name="xmlContent")
    def xml_content(self) -> Optional[pulumi.Input[str]]:
        """
        The XML Content for this Policy as a string.
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


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_id: Optional[pulumi.Input[str]] = None,
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a API Management service Policy.

        > **NOTE:** This resource will, upon creation, **overwrite any existing policy in the API Management service**, as there is no feasible way to test whether the policy has been modified from the default. Similarly, when this resource is destroyed, the API Management service will revert to its default policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.apimanagement.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            publisher_name="pub1",
            publisher_email="pub1@email.com",
            sku_name="Developer_1")
        example_named_value = azure.apimanagement.NamedValue("exampleNamedValue",
            resource_group_name=example_resource_group.name,
            api_management_name=example_service.name,
            display_name="ExampleProperty",
            value="Example Value")
        example_policy = azure.apimanagement.Policy("examplePolicy",
            api_management_id=example_service.id,
            xml_content=(lambda path: open(path).read())("example.xml"))
        ```

        ## Import

        API Management service Policys can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/policy:Policy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/policies/policy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_id: The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy as a string.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a API Management service Policy.

        > **NOTE:** This resource will, upon creation, **overwrite any existing policy in the API Management service**, as there is no feasible way to test whether the policy has been modified from the default. Similarly, when this resource is destroyed, the API Management service will revert to its default policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_service = azure.apimanagement.Service("exampleService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            publisher_name="pub1",
            publisher_email="pub1@email.com",
            sku_name="Developer_1")
        example_named_value = azure.apimanagement.NamedValue("exampleNamedValue",
            resource_group_name=example_resource_group.name,
            api_management_name=example_service.name,
            display_name="ExampleProperty",
            value="Example Value")
        example_policy = azure.apimanagement.Policy("examplePolicy",
            api_management_id=example_service.id,
            xml_content=(lambda path: open(path).read())("example.xml"))
        ```

        ## Import

        API Management service Policys can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/policy:Policy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/policies/policy
        ```

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_id: Optional[pulumi.Input[str]] = None,
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None,
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

            if api_management_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_management_id'")
            __props__['api_management_id'] = api_management_id
            __props__['xml_content'] = xml_content
            __props__['xml_link'] = xml_link
        super(Policy, __self__).__init__(
            'azure:apimanagement/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_id: Optional[pulumi.Input[str]] = None,
            xml_content: Optional[pulumi.Input[str]] = None,
            xml_link: Optional[pulumi.Input[str]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_id: The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy as a string.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_id"] = api_management_id
        __props__["xml_content"] = xml_content
        __props__["xml_link"] = xml_link
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiManagementId")
    def api_management_id(self) -> pulumi.Output[str]:
        """
        The ID of the API Management service. Changing this forces a new API Management service Policy to be created.
        """
        return pulumi.get(self, "api_management_id")

    @property
    @pulumi.getter(name="xmlContent")
    def xml_content(self) -> pulumi.Output[str]:
        """
        The XML Content for this Policy as a string.
        """
        return pulumi.get(self, "xml_content")

    @property
    @pulumi.getter(name="xmlLink")
    def xml_link(self) -> pulumi.Output[Optional[str]]:
        """
        A link to a Policy XML Document, which must be publicly available.
        """
        return pulumi.get(self, "xml_link")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

