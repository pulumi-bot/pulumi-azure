# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SourceCodeToken']


class SourceCodeToken(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 token_secret: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an App Service source control token.

        > **NOTE:** Source Control Tokens are configured at the subscription level, not on each App Service - as such this can only be configured Subscription-wide

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.appservice.SourceCodeToken("example",
            token="7e57735e77e577e57",
            type="GitHub")
        ```

        ## Import

        App Service Source Control Token's can be imported using the `type`, e.g.

        ```sh
         $ pulumi import azure:appservice/sourceCodeToken:SourceCodeToken example GitHub
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] token: The OAuth access token.
        :param pulumi.Input[str] token_secret: The OAuth access token secret.
        :param pulumi.Input[str] type: The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
        """
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

            if token is None and not opts.urn:
                raise TypeError("Missing required property 'token'")
            __props__['token'] = token
            __props__['token_secret'] = token_secret
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
        super(SourceCodeToken, __self__).__init__(
            'azure:appservice/sourceCodeToken:SourceCodeToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            token: Optional[pulumi.Input[str]] = None,
            token_secret: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'SourceCodeToken':
        """
        Get an existing SourceCodeToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] token: The OAuth access token.
        :param pulumi.Input[str] token_secret: The OAuth access token secret.
        :param pulumi.Input[str] type: The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["token"] = token
        __props__["token_secret"] = token_secret
        __props__["type"] = type
        return SourceCodeToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        The OAuth access token.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="tokenSecret")
    def token_secret(self) -> pulumi.Output[Optional[str]]:
        """
        The OAuth access token secret.
        """
        return pulumi.get(self, "token_secret")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

