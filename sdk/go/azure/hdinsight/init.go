// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

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
	case "azure:hdinsight/hBaseCluster:HBaseCluster":
		r, err = NewHBaseCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure:hdinsight/hadoopCluster:HadoopCluster":
		r, err = NewHadoopCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure:hdinsight/interactiveQueryCluster:InteractiveQueryCluster":
		r, err = NewInteractiveQueryCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure:hdinsight/kafkaCluster:KafkaCluster":
		r, err = NewKafkaCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure:hdinsight/mLServicesCluster:MLServicesCluster":
		r, err = NewMLServicesCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure:hdinsight/rServerCluster:RServerCluster":
		r, err = NewRServerCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure:hdinsight/sparkCluster:SparkCluster":
		r, err = NewSparkCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure:hdinsight/stormCluster:StormCluster":
		r, err = NewStormCluster(ctx, name, nil, pulumi.URN_(urn))
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
