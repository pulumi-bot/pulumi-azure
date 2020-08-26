# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Key']


class Key(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 curve: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 key_opts: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 key_size: Optional[pulumi.Input[float]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_before_date: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Key Vault Key.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_random as random

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        server = random.RandomId("server",
            keepers={
                "ami_id": 1,
            },
            byte_length=8)
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="premium",
            access_policies=[azure.keyvault.KeyVaultAccessPolicyArgs(
                tenant_id=current.tenant_id,
                object_id=current.object_id,
                key_permissions=[
                    "create",
                    "get",
                ],
                secret_permissions=["set"],
            )],
            tags={
                "environment": "Production",
            })
        generated = azure.keyvault.Key("generated",
            key_vault_id=example_key_vault.id,
            key_type="RSA",
            key_size=2048,
            key_opts=[
                "decrypt",
                "encrypt",
                "sign",
                "unwrapKey",
                "verify",
                "wrapKey",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] curve: Specifies the curve to use when creating an `EC` key. Possible values are `P-256`, `P-384`, `P-521`, and `SECP256K1`. This field will be required in a future release if `key_type` is `EC` or `EC-HSM`. The API will default to `P-256` if nothing is specified. Changing this forces a new resource to be created.
        :param pulumi.Input[str] expiration_date: Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[List[pulumi.Input[str]]] key_opts: A list of JSON web key operations. Possible values include: `decrypt`, `encrypt`, `sign`, `unwrapKey`, `verify` and `wrapKey`. Please note these values are case sensitive.
        :param pulumi.Input[float] key_size: Specifies the Size of the RSA key to create in bytes. For example, 1024 or 2048. *Note*: This field is required if `key_type` is `RSA` or `RSA-HSM`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] key_type: Specifies the Key Type to use for this Key Vault Key. Possible values are `EC` (Elliptic Curve), `EC-HSM`, `Oct` (Octet), `RSA` and `RSA-HSM`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Key should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Key. Changing this forces a new resource to be created.
        :param pulumi.Input[str] not_before_date: Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            __props__['curve'] = curve
            __props__['expiration_date'] = expiration_date
            if key_opts is None:
                raise TypeError("Missing required property 'key_opts'")
            __props__['key_opts'] = key_opts
            __props__['key_size'] = key_size
            if key_type is None:
                raise TypeError("Missing required property 'key_type'")
            __props__['key_type'] = key_type
            if key_vault_id is None:
                raise TypeError("Missing required property 'key_vault_id'")
            __props__['key_vault_id'] = key_vault_id
            __props__['name'] = name
            __props__['not_before_date'] = not_before_date
            __props__['tags'] = tags
            __props__['e'] = None
            __props__['n'] = None
            __props__['version'] = None
            __props__['x'] = None
            __props__['y'] = None
        super(Key, __self__).__init__(
            'azure:keyvault/key:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            curve: Optional[pulumi.Input[str]] = None,
            e: Optional[pulumi.Input[str]] = None,
            expiration_date: Optional[pulumi.Input[str]] = None,
            key_opts: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            key_size: Optional[pulumi.Input[float]] = None,
            key_type: Optional[pulumi.Input[str]] = None,
            key_vault_id: Optional[pulumi.Input[str]] = None,
            n: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            not_before_date: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            version: Optional[pulumi.Input[str]] = None,
            x: Optional[pulumi.Input[str]] = None,
            y: Optional[pulumi.Input[str]] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] curve: Specifies the curve to use when creating an `EC` key. Possible values are `P-256`, `P-384`, `P-521`, and `SECP256K1`. This field will be required in a future release if `key_type` is `EC` or `EC-HSM`. The API will default to `P-256` if nothing is specified. Changing this forces a new resource to be created.
        :param pulumi.Input[str] e: The RSA public exponent of this Key Vault Key.
        :param pulumi.Input[str] expiration_date: Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[List[pulumi.Input[str]]] key_opts: A list of JSON web key operations. Possible values include: `decrypt`, `encrypt`, `sign`, `unwrapKey`, `verify` and `wrapKey`. Please note these values are case sensitive.
        :param pulumi.Input[float] key_size: Specifies the Size of the RSA key to create in bytes. For example, 1024 or 2048. *Note*: This field is required if `key_type` is `RSA` or `RSA-HSM`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] key_type: Specifies the Key Type to use for this Key Vault Key. Possible values are `EC` (Elliptic Curve), `EC-HSM`, `Oct` (Octet), `RSA` and `RSA-HSM`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Key should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] n: The RSA modulus of this Key Vault Key.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Key. Changing this forces a new resource to be created.
        :param pulumi.Input[str] not_before_date: Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] version: The current version of the Key Vault Key.
        :param pulumi.Input[str] x: The EC X component of this Key Vault Key.
        :param pulumi.Input[str] y: The EC Y component of this Key Vault Key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["curve"] = curve
        __props__["e"] = e
        __props__["expiration_date"] = expiration_date
        __props__["key_opts"] = key_opts
        __props__["key_size"] = key_size
        __props__["key_type"] = key_type
        __props__["key_vault_id"] = key_vault_id
        __props__["n"] = n
        __props__["name"] = name
        __props__["not_before_date"] = not_before_date
        __props__["tags"] = tags
        __props__["version"] = version
        __props__["x"] = x
        __props__["y"] = y
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def curve(self) -> pulumi.Output[str]:
        """
        Specifies the curve to use when creating an `EC` key. Possible values are `P-256`, `P-384`, `P-521`, and `SECP256K1`. This field will be required in a future release if `key_type` is `EC` or `EC-HSM`. The API will default to `P-256` if nothing is specified. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "curve")

    @property
    @pulumi.getter
    def e(self) -> pulumi.Output[str]:
        """
        The RSA public exponent of this Key Vault Key.
        """
        return pulumi.get(self, "e")

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> pulumi.Output[Optional[str]]:
        """
        Expiration UTC datetime (Y-m-d'T'H:M:S'Z').
        """
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="keyOpts")
    def key_opts(self) -> pulumi.Output[List[str]]:
        """
        A list of JSON web key operations. Possible values include: `decrypt`, `encrypt`, `sign`, `unwrapKey`, `verify` and `wrapKey`. Please note these values are case sensitive.
        """
        return pulumi.get(self, "key_opts")

    @property
    @pulumi.getter(name="keySize")
    def key_size(self) -> pulumi.Output[Optional[float]]:
        """
        Specifies the Size of the RSA key to create in bytes. For example, 1024 or 2048. *Note*: This field is required if `key_type` is `RSA` or `RSA-HSM`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "key_size")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Output[str]:
        """
        Specifies the Key Type to use for this Key Vault Key. Possible values are `EC` (Elliptic Curve), `EC-HSM`, `Oct` (Octet), `RSA` and `RSA-HSM`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> pulumi.Output[str]:
        """
        The ID of the Key Vault where the Key should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter
    def n(self) -> pulumi.Output[str]:
        """
        The RSA modulus of this Key Vault Key.
        """
        return pulumi.get(self, "n")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Key Vault Key. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notBeforeDate")
    def not_before_date(self) -> pulumi.Output[Optional[str]]:
        """
        Key not usable before the provided UTC datetime (Y-m-d'T'H:M:S'Z').
        """
        return pulumi.get(self, "not_before_date")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The current version of the Key Vault Key.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def x(self) -> pulumi.Output[str]:
        """
        The EC X component of this Key Vault Key.
        """
        return pulumi.get(self, "x")

    @property
    @pulumi.getter
    def y(self) -> pulumi.Output[str]:
        """
        The EC Y component of this Key Vault Key.
        """
        return pulumi.get(self, "y")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

