// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
		r = &DatasetAzureBlob{}
	case "azure:datafactory/datasetCosmosDBApi:DatasetCosmosDBApi":
		r = &DatasetCosmosDBApi{}
	case "azure:datafactory/datasetDelimitedText:DatasetDelimitedText":
		r = &DatasetDelimitedText{}
	case "azure:datafactory/datasetHttp:DatasetHttp":
		r = &DatasetHttp{}
	case "azure:datafactory/datasetJson:DatasetJson":
		r = &DatasetJson{}
	case "azure:datafactory/datasetMysql:DatasetMysql":
		r = &DatasetMysql{}
	case "azure:datafactory/datasetParquet:DatasetParquet":
		r = &DatasetParquet{}
	case "azure:datafactory/datasetPostgresql:DatasetPostgresql":
		r = &DatasetPostgresql{}
	case "azure:datafactory/datasetSnowflake:DatasetSnowflake":
		r = &DatasetSnowflake{}
	case "azure:datafactory/datasetSqlServerTable:DatasetSqlServerTable":
		r = &DatasetSqlServerTable{}
	case "azure:datafactory/factory:Factory":
		r = &Factory{}
	case "azure:datafactory/integrationRuntimeManaged:IntegrationRuntimeManaged":
		r = &IntegrationRuntimeManaged{}
	case "azure:datafactory/integrationRuntimeRule:IntegrationRuntimeRule":
		r = &IntegrationRuntimeRule{}
	case "azure:datafactory/integrationRuntimeSelfHosted:IntegrationRuntimeSelfHosted":
		r = &IntegrationRuntimeSelfHosted{}
	case "azure:datafactory/integrationRuntimeSsis:IntegrationRuntimeSsis":
		r = &IntegrationRuntimeSsis{}
	case "azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage":
		r = &LinkedServiceAzureBlobStorage{}
	case "azure:datafactory/linkedServiceAzureDatabricks:LinkedServiceAzureDatabricks":
		r = &LinkedServiceAzureDatabricks{}
	case "azure:datafactory/linkedServiceAzureFileStorage:LinkedServiceAzureFileStorage":
		r = &LinkedServiceAzureFileStorage{}
	case "azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction":
		r = &LinkedServiceAzureFunction{}
	case "azure:datafactory/linkedServiceAzureSearch:LinkedServiceAzureSearch":
		r = &LinkedServiceAzureSearch{}
	case "azure:datafactory/linkedServiceAzureSqlDatabase:LinkedServiceAzureSqlDatabase":
		r = &LinkedServiceAzureSqlDatabase{}
	case "azure:datafactory/linkedServiceAzureTableStorage:LinkedServiceAzureTableStorage":
		r = &LinkedServiceAzureTableStorage{}
	case "azure:datafactory/linkedServiceCosmosDb:LinkedServiceCosmosDb":
		r = &LinkedServiceCosmosDb{}
	case "azure:datafactory/linkedServiceDataLakeStorageGen2:LinkedServiceDataLakeStorageGen2":
		r = &LinkedServiceDataLakeStorageGen2{}
	case "azure:datafactory/linkedServiceKeyVault:LinkedServiceKeyVault":
		r = &LinkedServiceKeyVault{}
	case "azure:datafactory/linkedServiceKusto:LinkedServiceKusto":
		r = &LinkedServiceKusto{}
	case "azure:datafactory/linkedServiceMysql:LinkedServiceMysql":
		r = &LinkedServiceMysql{}
	case "azure:datafactory/linkedServicePostgresql:LinkedServicePostgresql":
		r = &LinkedServicePostgresql{}
	case "azure:datafactory/linkedServiceSftp:LinkedServiceSftp":
		r = &LinkedServiceSftp{}
	case "azure:datafactory/linkedServiceSnowflake:LinkedServiceSnowflake":
		r = &LinkedServiceSnowflake{}
	case "azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer":
		r = &LinkedServiceSqlServer{}
	case "azure:datafactory/linkedServiceSynapse:LinkedServiceSynapse":
		r = &LinkedServiceSynapse{}
	case "azure:datafactory/linkedServiceWeb:LinkedServiceWeb":
		r = &LinkedServiceWeb{}
	case "azure:datafactory/pipeline:Pipeline":
		r = &Pipeline{}
	case "azure:datafactory/triggerSchedule:TriggerSchedule":
		r = &TriggerSchedule{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
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
		"datafactory/datasetParquet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/datasetPostgresql",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/datasetSnowflake",
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
		"datafactory/integrationRuntimeRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/integrationRuntimeSelfHosted",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/integrationRuntimeSsis",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceAzureBlobStorage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceAzureDatabricks",
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
		"datafactory/linkedServiceAzureSearch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceAzureSqlDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datafactory/linkedServiceAzureTableStorage",
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
		"datafactory/linkedServiceKusto",
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
		"datafactory/linkedServiceSnowflake",
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
