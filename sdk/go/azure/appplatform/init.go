// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

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
	case "azure:appplatform/springCloudActiveDeployment:SpringCloudActiveDeployment":
		r, err = NewSpringCloudActiveDeployment(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appplatform/springCloudApp:SpringCloudApp":
		r, err = NewSpringCloudApp(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appplatform/springCloudCertificate:SpringCloudCertificate":
		r, err = NewSpringCloudCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appplatform/springCloudJavaDeployment:SpringCloudJavaDeployment":
		r, err = NewSpringCloudJavaDeployment(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appplatform/springCloudService:SpringCloudService":
		r, err = NewSpringCloudService(ctx, name, nil, pulumi.URN_(urn))
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
		"appplatform/springCloudActiveDeployment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appplatform/springCloudApp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appplatform/springCloudCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appplatform/springCloudJavaDeployment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appplatform/springCloudService",
		&module{version},
	)
}
