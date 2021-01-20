// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v2/go/azure"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure:datafactory/datasetAzureBlob:DatasetAzureBlob":
		r, err = NewDatasetAzureBlob(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/datasetCosmosDBApi:DatasetCosmosDBApi":
		r, err = NewDatasetCosmosDBApi(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/datasetDelimitedText:DatasetDelimitedText":
		r, err = NewDatasetDelimitedText(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/datasetHttp:DatasetHttp":
		r, err = NewDatasetHttp(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/datasetJson:DatasetJson":
		r, err = NewDatasetJson(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/datasetMysql:DatasetMysql":
		r, err = NewDatasetMysql(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/datasetPostgresql:DatasetPostgresql":
		r, err = NewDatasetPostgresql(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/datasetSqlServerTable:DatasetSqlServerTable":
		r, err = NewDatasetSqlServerTable(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/factory:Factory":
		r, err = NewFactory(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/integrationRuntimeManaged:IntegrationRuntimeManaged":
		r, err = NewIntegrationRuntimeManaged(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/integrationRuntimeSelfHosted:IntegrationRuntimeSelfHosted":
		r, err = NewIntegrationRuntimeSelfHosted(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage":
		r, err = NewLinkedServiceAzureBlobStorage(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceAzureFileStorage:LinkedServiceAzureFileStorage":
		r, err = NewLinkedServiceAzureFileStorage(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction":
		r, err = NewLinkedServiceAzureFunction(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceAzureSqlDatabase:LinkedServiceAzureSqlDatabase":
		r, err = NewLinkedServiceAzureSqlDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceCosmosDb:LinkedServiceCosmosDb":
		r, err = NewLinkedServiceCosmosDb(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceDataLakeStorageGen2:LinkedServiceDataLakeStorageGen2":
		r, err = NewLinkedServiceDataLakeStorageGen2(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceKeyVault:LinkedServiceKeyVault":
		r, err = NewLinkedServiceKeyVault(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceMysql:LinkedServiceMysql":
		r, err = NewLinkedServiceMysql(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServicePostgresql:LinkedServicePostgresql":
		r, err = NewLinkedServicePostgresql(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceSftp:LinkedServiceSftp":
		r, err = NewLinkedServiceSftp(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer":
		r, err = NewLinkedServiceSqlServer(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceSynapse:LinkedServiceSynapse":
		r, err = NewLinkedServiceSynapse(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/linkedServiceWeb:LinkedServiceWeb":
		r, err = NewLinkedServiceWeb(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/pipeline:Pipeline":
		r, err = NewPipeline(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datafactory/triggerSchedule:TriggerSchedule":
		r, err = NewTriggerSchedule(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/datasetAzureBlob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/datasetCosmosDBApi",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/datasetDelimitedText",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/datasetHttp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/datasetJson",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/datasetMysql",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/datasetPostgresql",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/datasetSqlServerTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/factory",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/integrationRuntimeManaged",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/integrationRuntimeSelfHosted",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceAzureBlobStorage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceAzureFileStorage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceAzureFunction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceAzureSqlDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceCosmosDb",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceDataLakeStorageGen2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceKeyVault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceMysql",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServicePostgresql",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceSftp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceSqlServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceSynapse",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceWeb",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/pipeline",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/triggerSchedule",
		&module{version},
	)
}
