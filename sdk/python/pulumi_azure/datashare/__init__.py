# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .account import *
from .dataset_blob_storage import *
from .dataset_data_lake_gen1 import *
from .dataset_data_lake_gen2 import *
from .dataset_kusto_cluster import *
from .dataset_kusto_database import *
from .get_account import *
from .get_dataset_blob_storage import *
from .get_dataset_data_lake_gen1 import *
from .get_dataset_data_lake_gen2 import *
from .get_dataset_kusto_cluster import *
from .get_dataset_kusto_database import *
from .get_share import *
from .share import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure:datashare/account:Account":
                return Account(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datashare/datasetBlobStorage:DatasetBlobStorage":
                return DatasetBlobStorage(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datashare/datasetDataLakeGen1:DatasetDataLakeGen1":
                return DatasetDataLakeGen1(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datashare/datasetDataLakeGen2:DatasetDataLakeGen2":
                return DatasetDataLakeGen2(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datashare/datasetKustoCluster:DatasetKustoCluster":
                return DatasetKustoCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datashare/datasetKustoDatabase:DatasetKustoDatabase":
                return DatasetKustoDatabase(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datashare/share:Share":
                return Share(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure", "datashare/account", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datashare/datasetBlobStorage", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datashare/datasetDataLakeGen1", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datashare/datasetDataLakeGen2", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datashare/datasetKustoCluster", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datashare/datasetKustoDatabase", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datashare/share", _module_instance)

_register_module()
