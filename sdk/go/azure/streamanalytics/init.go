// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

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
	case "azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF":
		r, err = NewFunctionJavaScriptUDF(ctx, name, nil, pulumi.URN_(urn))
	case "azure:streamanalytics/job:Job":
		r, err = NewJob(ctx, name, nil, pulumi.URN_(urn))
	case "azure:streamanalytics/outputBlob:OutputBlob":
		r, err = NewOutputBlob(ctx, name, nil, pulumi.URN_(urn))
	case "azure:streamanalytics/outputEventHub:OutputEventHub":
		r, err = NewOutputEventHub(ctx, name, nil, pulumi.URN_(urn))
	case "azure:streamanalytics/outputMssql:OutputMssql":
		r, err = NewOutputMssql(ctx, name, nil, pulumi.URN_(urn))
	case "azure:streamanalytics/outputServiceBusQueue:OutputServiceBusQueue":
		r, err = NewOutputServiceBusQueue(ctx, name, nil, pulumi.URN_(urn))
	case "azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic":
		r, err = NewOutputServicebusTopic(ctx, name, nil, pulumi.URN_(urn))
	case "azure:streamanalytics/referenceInputBlob:ReferenceInputBlob":
		r, err = NewReferenceInputBlob(ctx, name, nil, pulumi.URN_(urn))
	case "azure:streamanalytics/streamInputBlob:StreamInputBlob":
		r, err = NewStreamInputBlob(ctx, name, nil, pulumi.URN_(urn))
	case "azure:streamanalytics/streamInputEventHub:StreamInputEventHub":
		r, err = NewStreamInputEventHub(ctx, name, nil, pulumi.URN_(urn))
	case "azure:streamanalytics/streamInputIotHub:StreamInputIotHub":
		r, err = NewStreamInputIotHub(ctx, name, nil, pulumi.URN_(urn))
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
		"streamanalytics/functionJavaScriptUDF",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/job",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/outputBlob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/outputEventHub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/outputMssql",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/outputServiceBusQueue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/outputServicebusTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/referenceInputBlob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/streamInputBlob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/streamInputEventHub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/streamInputIotHub",
		&module{version},
	)
}