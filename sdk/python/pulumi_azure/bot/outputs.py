# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'ChannelDirectLineSite',
]

@pulumi.output_type
class ChannelDirectLineSite(dict):
    def __init__(__self__, *,
                 name: str,
                 enabled: Optional[bool] = None,
                 enhanced_authentication_enabled: Optional[bool] = None,
                 id: Optional[str] = None,
                 key: Optional[str] = None,
                 key2: Optional[str] = None,
                 trusted_origins: Optional[Sequence[str]] = None,
                 v1_allowed: Optional[bool] = None,
                 v3_allowed: Optional[bool] = None):
        """
        :param str name: The name of the site
        :param bool enabled: Enables/Disables this site. Enabled by default
        :param bool enhanced_authentication_enabled: Enables additional security measures for this site, see [Enhanced Directline Authentication Features](https://blog.botframework.com/2018/09/25/enhanced-direct-line-authentication-features). Disabled by default.
        :param str id: Id for the site
        :param str key: Primary key for accessing this site
        :param str key2: Secondary key for accessing this site
        :param Sequence[str] trusted_origins: This field is required when `is_secure_site_enabled` is enabled. Determines which origins can establish a Directline conversation for this site.
        :param bool v1_allowed: Enables v1 of the Directline protocol for this site. Enabled by default
        :param bool v3_allowed: Enables v3 of the Directline protocol for this site. Enabled by default
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
    def name(self) -> str:
        """
        The name of the site
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Enables/Disables this site. Enabled by default
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="enhancedAuthenticationEnabled")
    def enhanced_authentication_enabled(self) -> Optional[bool]:
        """
        Enables additional security measures for this site, see [Enhanced Directline Authentication Features](https://blog.botframework.com/2018/09/25/enhanced-direct-line-authentication-features). Disabled by default.
        """
        return pulumi.get(self, "enhanced_authentication_enabled")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Id for the site
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        """
        Primary key for accessing this site
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def key2(self) -> Optional[str]:
        """
        Secondary key for accessing this site
        """
        return pulumi.get(self, "key2")

    @property
    @pulumi.getter(name="trustedOrigins")
    def trusted_origins(self) -> Optional[Sequence[str]]:
        """
        This field is required when `is_secure_site_enabled` is enabled. Determines which origins can establish a Directline conversation for this site.
        """
        return pulumi.get(self, "trusted_origins")

    @property
    @pulumi.getter(name="v1Allowed")
    def v1_allowed(self) -> Optional[bool]:
        """
        Enables v1 of the Directline protocol for this site. Enabled by default
        """
        return pulumi.get(self, "v1_allowed")

    @property
    @pulumi.getter(name="v3Allowed")
    def v3_allowed(self) -> Optional[bool]:
        """
        Enables v3 of the Directline protocol for this site. Enabled by default
        """
        return pulumi.get(self, "v3_allowed")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


