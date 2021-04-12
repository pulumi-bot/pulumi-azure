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

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 sku_name: pulumi.Input[str],
                 custom_subdomain_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_acls: Optional[pulumi.Input['AccountNetworkAclsArgs']] = None,
                 qna_runtime_endpoint: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[str] kind: Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
        :param pulumi.Input[str] custom_subdomain_name: The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
        :param pulumi.Input['AccountNetworkAclsArgs'] network_acls: A `network_acls` block as defined below.
        :param pulumi.Input[str] qna_runtime_endpoint: A URL to link a QnAMaker cognitive account to a QnA runtime.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku_name", sku_name)
        if custom_subdomain_name is not None:
            pulumi.set(__self__, "custom_subdomain_name", custom_subdomain_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_acls is not None:
            pulumi.set(__self__, "network_acls", network_acls)
        if qna_runtime_endpoint is not None:
            pulumi.set(__self__, "qna_runtime_endpoint", qna_runtime_endpoint)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        """
        Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Input[str]:
        """
        Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
        """
        return pulumi.get(self, "sku_name")

    @sku_name.setter
    def sku_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "sku_name", value)

    @property
    @pulumi.getter(name="customSubdomainName")
    def custom_subdomain_name(self) -> Optional[pulumi.Input[str]]:
        """
        The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "custom_subdomain_name")

    @custom_subdomain_name.setter
    def custom_subdomain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_subdomain_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkAcls")
    def network_acls(self) -> Optional[pulumi.Input['AccountNetworkAclsArgs']]:
        """
        A `network_acls` block as defined below.
        """
        return pulumi.get(self, "network_acls")

    @network_acls.setter
    def network_acls(self, value: Optional[pulumi.Input['AccountNetworkAclsArgs']]):
        pulumi.set(self, "network_acls", value)

    @property
    @pulumi.getter(name="qnaRuntimeEndpoint")
    def qna_runtime_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        A URL to link a QnAMaker cognitive account to a QnA runtime.
        """
        return pulumi.get(self, "qna_runtime_endpoint")

    @qna_runtime_endpoint.setter
    def qna_runtime_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "qna_runtime_endpoint", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_subdomain_name: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_acls: Optional[pulumi.Input[pulumi.InputType['AccountNetworkAclsArgs']]] = None,
                 qna_runtime_endpoint: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Cognitive Services Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.cognitive.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            kind="Face",
            sku_name="S0",
            tags={
                "Acceptance": "Test",
            })
        ```

        ## Import

        Cognitive Service Accounts can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:cognitive/account:Account account1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] custom_subdomain_name: The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
        :param pulumi.Input[str] kind: Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['AccountNetworkAclsArgs']] network_acls: A `network_acls` block as defined below.
        :param pulumi.Input[str] qna_runtime_endpoint: A URL to link a QnAMaker cognitive account to a QnA runtime.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Cognitive Services Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.cognitive.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            kind="Face",
            sku_name="S0",
            tags={
                "Acceptance": "Test",
            })
        ```

        ## Import

        Cognitive Service Accounts can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:cognitive/account:Account account1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
        ```

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_subdomain_name: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_acls: Optional[pulumi.Input[pulumi.InputType['AccountNetworkAclsArgs']]] = None,
                 qna_runtime_endpoint: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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

            __props__['custom_subdomain_name'] = custom_subdomain_name
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__['kind'] = kind
            __props__['location'] = location
            __props__['name'] = name
            __props__['network_acls'] = network_acls
            __props__['qna_runtime_endpoint'] = qna_runtime_endpoint
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku_name is None and not opts.urn:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['endpoint'] = None
            __props__['primary_access_key'] = None
            __props__['secondary_access_key'] = None
        super(Account, __self__).__init__(
            'azure:cognitive/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_subdomain_name: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            kind: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_acls: Optional[pulumi.Input[pulumi.InputType['AccountNetworkAclsArgs']]] = None,
            primary_access_key: Optional[pulumi.Input[str]] = None,
            qna_runtime_endpoint: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_access_key: Optional[pulumi.Input[str]] = None,
            sku_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] custom_subdomain_name: The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
        :param pulumi.Input[str] endpoint: The endpoint used to connect to the Cognitive Service Account.
        :param pulumi.Input[str] kind: Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['AccountNetworkAclsArgs']] network_acls: A `network_acls` block as defined below.
        :param pulumi.Input[str] primary_access_key: A primary access key which can be used to connect to the Cognitive Service Account.
        :param pulumi.Input[str] qna_runtime_endpoint: A URL to link a QnAMaker cognitive account to a QnA runtime.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_access_key: The secondary access key which can be used to connect to the Cognitive Service Account.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["custom_subdomain_name"] = custom_subdomain_name
        __props__["endpoint"] = endpoint
        __props__["kind"] = kind
        __props__["location"] = location
        __props__["name"] = name
        __props__["network_acls"] = network_acls
        __props__["primary_access_key"] = primary_access_key
        __props__["qna_runtime_endpoint"] = qna_runtime_endpoint
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_access_key"] = secondary_access_key
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customSubdomainName")
    def custom_subdomain_name(self) -> pulumi.Output[Optional[str]]:
        """
        The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "custom_subdomain_name")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint used to connect to the Cognitive Service Account.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkAcls")
    def network_acls(self) -> pulumi.Output[Optional['outputs.AccountNetworkAcls']]:
        """
        A `network_acls` block as defined below.
        """
        return pulumi.get(self, "network_acls")

    @property
    @pulumi.getter(name="primaryAccessKey")
    def primary_access_key(self) -> pulumi.Output[str]:
        """
        A primary access key which can be used to connect to the Cognitive Service Account.
        """
        return pulumi.get(self, "primary_access_key")

    @property
    @pulumi.getter(name="qnaRuntimeEndpoint")
    def qna_runtime_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        A URL to link a QnAMaker cognitive account to a QnA runtime.
        """
        return pulumi.get(self, "qna_runtime_endpoint")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryAccessKey")
    def secondary_access_key(self) -> pulumi.Output[str]:
        """
        The secondary access key which can be used to connect to the Cognitive Service Account.
        """
        return pulumi.get(self, "secondary_access_key")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Output[str]:
        """
        Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

