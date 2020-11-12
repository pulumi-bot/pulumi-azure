# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ManagementPolicy']


class ManagementPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ManagementPolicyRuleArgs']]]]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Azure Storage Account Management Policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="westus")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS",
            account_kind="BlobStorage")
        example_management_policy = azure.storage.ManagementPolicy("exampleManagementPolicy",
            storage_account_id=example_account.id,
            rules=[
                azure.storage.ManagementPolicyRuleArgs(
                    name="rule1",
                    enabled=True,
                    filters=azure.storage.ManagementPolicyRuleFiltersArgs(
                        prefix_matches=["container1/prefix1"],
                        blob_types=["blockBlob"],
                    ),
                    actions=azure.storage.ManagementPolicyRuleActionsArgs(
                        base_blob=azure.storage.ManagementPolicyRuleActionsBaseBlobArgs(
                            tier_to_cool_after_days_since_modification_greater_than=10,
                            tier_to_archive_after_days_since_modification_greater_than=50,
                            delete_after_days_since_modification_greater_than=100,
                        ),
                        snapshot=azure.storage.ManagementPolicyRuleActionsSnapshotArgs(
                            delete_after_days_since_creation_greater_than=30,
                        ),
                    ),
                ),
                azure.storage.ManagementPolicyRuleArgs(
                    name="rule2",
                    enabled=False,
                    filters=azure.storage.ManagementPolicyRuleFiltersArgs(
                        prefix_matches=[
                            "container2/prefix1",
                            "container2/prefix2",
                        ],
                        blob_types=["blockBlob"],
                    ),
                    actions=azure.storage.ManagementPolicyRuleActionsArgs(
                        base_blob=azure.storage.ManagementPolicyRuleActionsBaseBlobArgs(
                            tier_to_cool_after_days_since_modification_greater_than=11,
                            tier_to_archive_after_days_since_modification_greater_than=51,
                            delete_after_days_since_modification_greater_than=101,
                        ),
                        snapshot=azure.storage.ManagementPolicyRuleActionsSnapshotArgs(
                            delete_after_days_since_creation_greater_than=31,
                        ),
                    ),
                ),
            ])
        ```

        ## Import

        Storage Account Management Policies can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:storage/managementPolicy:ManagementPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Storage/storageAccounts/myaccountname/managementPolicies/default
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ManagementPolicyRuleArgs']]]] rules: A `rule` block as documented below.
        :param pulumi.Input[str] storage_account_id: Specifies the id of the storage account to apply the management policy to.
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

            __props__['rules'] = rules
            if storage_account_id is None:
                raise TypeError("Missing required property 'storage_account_id'")
            __props__['storage_account_id'] = storage_account_id
        super(ManagementPolicy, __self__).__init__(
            'azure:storage/managementPolicy:ManagementPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ManagementPolicyRuleArgs']]]]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None) -> 'ManagementPolicy':
        """
        Get an existing ManagementPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ManagementPolicyRuleArgs']]]] rules: A `rule` block as documented below.
        :param pulumi.Input[str] storage_account_id: Specifies the id of the storage account to apply the management policy to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["rules"] = rules
        __props__["storage_account_id"] = storage_account_id
        return ManagementPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.ManagementPolicyRule']]]:
        """
        A `rule` block as documented below.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[str]:
        """
        Specifies the id of the storage account to apply the management policy to.
        """
        return pulumi.get(self, "storage_account_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

