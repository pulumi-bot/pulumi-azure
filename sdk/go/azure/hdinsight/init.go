// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

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
	case "azure:hdinsight/hBaseCluster:HBaseCluster":
		r = &HBaseCluster{}
	case "azure:hdinsight/hadoopCluster:HadoopCluster":
		r = &HadoopCluster{}
	case "azure:hdinsight/interactiveQueryCluster:InteractiveQueryCluster":
		r = &InteractiveQueryCluster{}
	case "azure:hdinsight/kafkaCluster:KafkaCluster":
		r = &KafkaCluster{}
	case "azure:hdinsight/mLServicesCluster:MLServicesCluster":
		r = &MLServicesCluster{}
	case "azure:hdinsight/rServerCluster:RServerCluster":
		r = &RServerCluster{}
	case "azure:hdinsight/sparkCluster:SparkCluster":
		r = &SparkCluster{}
	case "azure:hdinsight/stormCluster:StormCluster":
		r = &StormCluster{}
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
		"hdinsight/hBaseCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"hdinsight/hadoopCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"hdinsight/interactiveQueryCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"hdinsight/kafkaCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"hdinsight/mLServicesCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"hdinsight/rServerCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"hdinsight/sparkCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"hdinsight/stormCluster",
		&module{version},
	)
}
