// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bot

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
	case "azure:bot/channelDirectLine:ChannelDirectLine":
		r, err = NewChannelDirectLine(ctx, name, nil, pulumi.URN_(urn))
	case "azure:bot/channelEmail:ChannelEmail":
		r, err = NewChannelEmail(ctx, name, nil, pulumi.URN_(urn))
	case "azure:bot/channelSlack:ChannelSlack":
		r, err = NewChannelSlack(ctx, name, nil, pulumi.URN_(urn))
	case "azure:bot/channelTeams:ChannelTeams":
		r, err = NewChannelTeams(ctx, name, nil, pulumi.URN_(urn))
	case "azure:bot/channelsRegistration:ChannelsRegistration":
		r, err = NewChannelsRegistration(ctx, name, nil, pulumi.URN_(urn))
	case "azure:bot/connection:Connection":
		r, err = NewConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure:bot/webApp:WebApp":
		r, err = NewWebApp(ctx, name, nil, pulumi.URN_(urn))
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
		"bot/channelDirectLine",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"bot/channelEmail",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"bot/channelSlack",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"bot/channelTeams",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"bot/channelsRegistration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"bot/connection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"bot/webApp",
		&module{version},
	)
}