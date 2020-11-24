# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Dashboard']


class Dashboard(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dashboard_properties: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a shared dashboard in the Azure Portal.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        config = pulumi.Config()
        md_content = config.get("mdContent")
        if md_content is None:
            md_content = "# Hello all :)"
        video_link = config.get("videoLink")
        if video_link is None:
            video_link = "https://www.youtube.com/watch?v=......"
        current = azure.core.get_subscription()
        my_group = azure.core.ResourceGroup("my-group", location="uksouth")
        my_board = azure.dashboard.Dashboard("my-board",
            resource_group_name=my_group.name,
            location=my_group.location,
            tags={
                "source": "managed",
            },
            dashboard_properties=f\"\"\"{{
           "lenses": {{
                "0": {{
                    "order": 0,
                    "parts": {{
                        "0": {{
                            "position": {{
                                "x": 0,
                                "y": 0,
                                "rowSpan": 2,
                                "colSpan": 3
                            }},
                            "metadata": {{
                                "inputs": [],
                                "type": "Extension/HubsExtension/PartType/MarkdownPart",
                                "settings": {{
                                    "content": {{
                                        "settings": {{
                                            "content": "{md_content}",
                                            "subtitle": "",
                                            "title": ""
                                        }}
                                    }}
                                }}
                            }}
                        }},               
                        "1": {{
                            "position": {{
                                "x": 5,
                                "y": 0,
                                "rowSpan": 4,
                                "colSpan": 6
                            }},
                            "metadata": {{
                                "inputs": [],
                                "type": "Extension/HubsExtension/PartType/VideoPart",
                                "settings": {{
                                    "content": {{
                                        "settings": {{
                                            "title": "Important Information",
                                            "subtitle": "",
                                            "src": "{video_link}",
                                            "autoplay": true
                                        }}
                                    }}
                                }}
                            }}
                        }},
                        "2": {{
                            "position": {{
                                "x": 0,
                                "y": 4,
                                "rowSpan": 4,
                                "colSpan": 6
                            }},
                            "metadata": {{
                                "inputs": [
                                    {{
                                        "name": "ComponentId",
                                        "value": "/subscriptions/{current.subscription_id}/resourceGroups/myRG/providers/microsoft.insights/components/myWebApp"
                                    }}
                                ],
                                "type": "Extension/AppInsightsExtension/PartType/AppMapGalPt",
                                "settings": {{}},
                                "asset": {{
                                    "idInputName": "ComponentId",
                                    "type": "ApplicationInsights"
                                }}
                            }}
                        }}              
                    }}
                }}
            }},
            "metadata": {{
                "model": {{
                    "timeRange": {{
                        "value": {{
                            "relative": {{
                                "duration": 24,
                                "timeUnit": 1
                            }}
                        }},
                        "type": "MsPortalFx.Composition.Configuration.ValueTypes.TimeRange"
                    }},
                    "filterLocale": {{
                        "value": "en-us"
                    }},
                    "filters": {{
                        "value": {{
                            "MsPortalFx_TimeRange": {{
                                "model": {{
                                    "format": "utc",
                                    "granularity": "auto",
                                    "relative": "24h"
                                }},
                                "displayCache": {{
                                    "name": "UTC Time",
                                    "value": "Past 24 hours"
                                }},
                                "filteredPartIds": [
                                    "StartboardPart-UnboundPart-ae44fef5-76b8-46b0-86f0-2b3f47bad1c7"
                                ]
                            }}
                        }}
                    }}
                }}
            }}
        }}
        \"\"\")
        ```

        It is recommended to follow the steps outlined
        [here](https://docs.microsoft.com/en-us/azure/azure-portal/azure-portal-dashboards-create-programmatically#fetch-the-json-representation-of-the-dashboard) to create a Dashboard in the Portal and extract the relevant JSON to use in this resource. From the extracted JSON, the contents of the `properties: {}` object can used. Variables can be injected as needed - see above example.

        ## Import

        Dashboards can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:dashboard/dashboard:Dashboard my-board /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Portal/dashboards/00000000-0000-0000-0000-000000000000
        ```

         Note the URI in the above sample can be found using the Resource Explorer tool in the Azure Portal.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dashboard_properties: JSON data representing dashboard body. See above for details on how to obtain this from the Portal.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Shared Dashboard. This should be be 64 chars max, only alphanumeric and hyphens (no spaces). For a more friendly display name, add the `hidden-title` tag.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the dashboard.
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

            __props__['dashboard_properties'] = dashboard_properties
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(Dashboard, __self__).__init__(
            'azure:dashboard/dashboard:Dashboard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dashboard_properties: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Dashboard':
        """
        Get an existing Dashboard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dashboard_properties: JSON data representing dashboard body. See above for details on how to obtain this from the Portal.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Shared Dashboard. This should be be 64 chars max, only alphanumeric and hyphens (no spaces). For a more friendly display name, add the `hidden-title` tag.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the dashboard.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dashboard_properties"] = dashboard_properties
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return Dashboard(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dashboardProperties")
    def dashboard_properties(self) -> pulumi.Output[str]:
        """
        JSON data representing dashboard body. See above for details on how to obtain this from the Portal.
        """
        return pulumi.get(self, "dashboard_properties")

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
        Specifies the name of the Shared Dashboard. This should be be 64 chars max, only alphanumeric and hyphens (no spaces). For a more friendly display name, add the `hidden-title` tag.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to
        create the dashboard.
        """
        return pulumi.get(self, "resource_group_name")

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

