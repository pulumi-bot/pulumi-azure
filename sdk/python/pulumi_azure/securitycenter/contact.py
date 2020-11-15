# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Contact']


class Contact(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_notifications: Optional[pulumi.Input[bool]] = None,
                 alerts_to_admins: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages the subscription's Security Center Contact.

        > **NOTE:** Owner access permission is required.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.securitycenter.Contact("example",
            alert_notifications=True,
            alerts_to_admins=True,
            email="contact@example.com",
            phone="+1-555-555-5555")
        ```

        ## Import

        The contact can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:securitycenter/contact:Contact example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/securityContacts/default1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] alert_notifications: Whether to send security alerts notifications to the security contact.
        :param pulumi.Input[bool] alerts_to_admins: Whether to send security alerts notifications to subscription admins.
        :param pulumi.Input[str] email: The email of the Security Center Contact.
        :param pulumi.Input[str] phone: The phone number of the Security Center Contact.
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

            if alert_notifications is None:
                raise TypeError("Missing required property 'alert_notifications'")
            __props__['alert_notifications'] = alert_notifications
            if alerts_to_admins is None:
                raise TypeError("Missing required property 'alerts_to_admins'")
            __props__['alerts_to_admins'] = alerts_to_admins
            if email is None:
                raise TypeError("Missing required property 'email'")
            __props__['email'] = email
            __props__['phone'] = phone
        super(Contact, __self__).__init__(
            'azure:securitycenter/contact:Contact',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_notifications: Optional[pulumi.Input[bool]] = None,
            alerts_to_admins: Optional[pulumi.Input[bool]] = None,
            email: Optional[pulumi.Input[str]] = None,
            phone: Optional[pulumi.Input[str]] = None) -> 'Contact':
        """
        Get an existing Contact resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] alert_notifications: Whether to send security alerts notifications to the security contact.
        :param pulumi.Input[bool] alerts_to_admins: Whether to send security alerts notifications to subscription admins.
        :param pulumi.Input[str] email: The email of the Security Center Contact.
        :param pulumi.Input[str] phone: The phone number of the Security Center Contact.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alert_notifications"] = alert_notifications
        __props__["alerts_to_admins"] = alerts_to_admins
        __props__["email"] = email
        __props__["phone"] = phone
        return Contact(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertNotifications")
    def alert_notifications(self) -> pulumi.Output[bool]:
        """
        Whether to send security alerts notifications to the security contact.
        """
        return pulumi.get(self, "alert_notifications")

    @property
    @pulumi.getter(name="alertsToAdmins")
    def alerts_to_admins(self) -> pulumi.Output[bool]:
        """
        Whether to send security alerts notifications to subscription admins.
        """
        return pulumi.get(self, "alerts_to_admins")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        The email of the Security Center Contact.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def phone(self) -> pulumi.Output[Optional[str]]:
        """
        The phone number of the Security Center Contact.
        """
        return pulumi.get(self, "phone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

