// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bot

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Slack integration for a Bot Channel
//
// > **Note** A bot can only have a single Slack Channel associated with it.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/bot"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleChannelsRegistration, err := bot.NewChannelsRegistration(ctx, "exampleChannelsRegistration", &bot.ChannelsRegistrationArgs{
// 			Location:          pulumi.String("global"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("F0"),
// 			MicrosoftAppId:    pulumi.String(current.ClientId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bot.NewChannelSlack(ctx, "exampleChannelSlack", &bot.ChannelSlackArgs{
// 			BotName:           exampleChannelsRegistration.Name,
// 			Location:          exampleChannelsRegistration.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ClientId:          pulumi.String("exampleId"),
// 			ClientSecret:      pulumi.String("exampleSecret"),
// 			VerificationToken: pulumi.String("exampleVerificationToken"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ChannelSlack struct {
	pulumi.CustomResourceState

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName pulumi.StringOutput `pulumi:"botName"`
	// The Client ID that will be used to authenticate with Slack.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The Client Secret that will be used to authenticate with Slack.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// The Slack Landing Page URL.
	LandingPageUrl pulumi.StringPtrOutput `pulumi:"landingPageUrl"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Verification Token that will be used to authenticate with Slack.
	VerificationToken pulumi.StringOutput `pulumi:"verificationToken"`
}

// NewChannelSlack registers a new resource with the given unique name, arguments, and options.
func NewChannelSlack(ctx *pulumi.Context,
	name string, args *ChannelSlackArgs, opts ...pulumi.ResourceOption) (*ChannelSlack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.BotName == nil {
		return nil, errors.New("invalid value for required argument 'BotName'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VerificationToken == nil {
		return nil, errors.New("invalid value for required argument 'VerificationToken'")
	}
	var resource ChannelSlack
	err := ctx.RegisterResource("azure:bot/channelSlack:ChannelSlack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannelSlack gets an existing ChannelSlack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannelSlack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelSlackState, opts ...pulumi.ResourceOption) (*ChannelSlack, error) {
	var resource ChannelSlack
	err := ctx.ReadResource("azure:bot/channelSlack:ChannelSlack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChannelSlack resources.
type channelSlackState struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName *string `pulumi:"botName"`
	// The Client ID that will be used to authenticate with Slack.
	ClientId *string `pulumi:"clientId"`
	// The Client Secret that will be used to authenticate with Slack.
	ClientSecret *string `pulumi:"clientSecret"`
	// The Slack Landing Page URL.
	LandingPageUrl *string `pulumi:"landingPageUrl"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Verification Token that will be used to authenticate with Slack.
	VerificationToken *string `pulumi:"verificationToken"`
}

type ChannelSlackState struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName pulumi.StringPtrInput
	// The Client ID that will be used to authenticate with Slack.
	ClientId pulumi.StringPtrInput
	// The Client Secret that will be used to authenticate with Slack.
	ClientSecret pulumi.StringPtrInput
	// The Slack Landing Page URL.
	LandingPageUrl pulumi.StringPtrInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Verification Token that will be used to authenticate with Slack.
	VerificationToken pulumi.StringPtrInput
}

func (ChannelSlackState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelSlackState)(nil)).Elem()
}

type channelSlackArgs struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName string `pulumi:"botName"`
	// The Client ID that will be used to authenticate with Slack.
	ClientId string `pulumi:"clientId"`
	// The Client Secret that will be used to authenticate with Slack.
	ClientSecret string `pulumi:"clientSecret"`
	// The Slack Landing Page URL.
	LandingPageUrl *string `pulumi:"landingPageUrl"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Verification Token that will be used to authenticate with Slack.
	VerificationToken string `pulumi:"verificationToken"`
}

// The set of arguments for constructing a ChannelSlack resource.
type ChannelSlackArgs struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName pulumi.StringInput
	// The Client ID that will be used to authenticate with Slack.
	ClientId pulumi.StringInput
	// The Client Secret that will be used to authenticate with Slack.
	ClientSecret pulumi.StringInput
	// The Slack Landing Page URL.
	LandingPageUrl pulumi.StringPtrInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Verification Token that will be used to authenticate with Slack.
	VerificationToken pulumi.StringInput
}

func (ChannelSlackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelSlackArgs)(nil)).Elem()
}
