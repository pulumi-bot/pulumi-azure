# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'ChannelDirectLineSiteArgs',
]

@pulumi.input_type
class ChannelDirectLineSiteArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 enhanced_authentication_enabled: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 key2: Optional[pulumi.Input[str]] = None,
                 trusted_origins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 v1_allowed: Optional[pulumi.Input[bool]] = None,
                 v3_allowed: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] name: The name of the site
        :param pulumi.Input[bool] enabled: Enables/Disables this site. Enabled by default
        :param pulumi.Input[bool] enhanced_authentication_enabled: Enables additional security measures for this site, see [Enhanced Directline Authentication Features](https://blog.botframework.com/2018/09/25/enhanced-direct-line-authentication-features). Disabled by default.
        :param pulumi.Input[str] id: Id for the site
        :param pulumi.Input[str] key: Primary key for accessing this site
        :param pulumi.Input[str] key2: Secondary key for accessing this site
        :param pulumi.Input[Sequence[pulumi.Input[str]]] trusted_origins: This field is required when `is_secure_site_enabled` is enabled. Determines which origins can establish a Directline conversation for this site.
        :param pulumi.Input[bool] v1_allowed: Enables v1 of the Directline protocol for this site. Enabled by default
        :param pulumi.Input[bool] v3_allowed: Enables v3 of the Directline protocol for this site. Enabled by default
        """
        pulumi.set(__self__, "name", name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if enhanced_authentication_enabled is not None:
            pulumi.set(__self__, "enhanced_authentication_enabled", enhanced_authentication_enabled)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if key2 is not None:
            pulumi.set(__self__, "key2", key2)
        if trusted_origins is not None:
            pulumi.set(__self__, "trusted_origins", trusted_origins)
        if v1_allowed is not None:
            pulumi.set(__self__, "v1_allowed", v1_allowed)
        if v3_allowed is not None:
            pulumi.set(__self__, "v3_allowed", v3_allowed)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the site
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables/Disables this site. Enabled by default
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="enhancedAuthenticationEnabled")
    def enhanced_authentication_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables additional security measures for this site, see [Enhanced Directline Authentication Features](https://blog.botframework.com/2018/09/25/enhanced-direct-line-authentication-features). Disabled by default.
        """
        return pulumi.get(self, "enhanced_authentication_enabled")

    @enhanced_authentication_enabled.setter
    def enhanced_authentication_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enhanced_authentication_enabled", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Id for the site
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Primary key for accessing this site
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def key2(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary key for accessing this site
        """
        return pulumi.get(self, "key2")

    @key2.setter
    def key2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key2", value)

    @property
    @pulumi.getter(name="trustedOrigins")
    def trusted_origins(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        This field is required when `is_secure_site_enabled` is enabled. Determines which origins can establish a Directline conversation for this site.
        """
        return pulumi.get(self, "trusted_origins")

    @trusted_origins.setter
    def trusted_origins(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "trusted_origins", value)

    @property
    @pulumi.getter(name="v1Allowed")
    def v1_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables v1 of the Directline protocol for this site. Enabled by default
        """
        return pulumi.get(self, "v1_allowed")

    @v1_allowed.setter
    def v1_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "v1_allowed", value)

    @property
    @pulumi.getter(name="v3Allowed")
    def v3_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables v3 of the Directline protocol for this site. Enabled by default
        """
        return pulumi.get(self, "v3_allowed")

    @v3_allowed.setter
    def v3_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "v3_allowed", value)


