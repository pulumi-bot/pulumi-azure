// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a MS Teams integration for a Bot Channel
//
// > **Note** A bot can only have a single MS Teams Channel associated with it.
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
// 		_, err = bot.NewChannelTeams(ctx, "exampleChannelTeams", &bot.ChannelTeamsArgs{
// 			BotName:           exampleChannelsRegistration.Name,
// 			Location:          exampleChannelsRegistration.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// The Microsoft Teams Integration for a Bot Channel can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:bot/channelTeams:ChannelTeams example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/MsTeamsChannel
// ```
type ChannelTeams struct {
	pulumi.CustomResourceState

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName pulumi.StringOutput `pulumi:"botName"`
	// Specifies the webhook for Microsoft Teams channel calls.
	CallingWebHook pulumi.StringOutput `pulumi:"callingWebHook"`
	// Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
	EnableCalling pulumi.BoolPtrOutput `pulumi:"enableCalling"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewChannelTeams registers a new resource with the given unique name, arguments, and options.
func NewChannelTeams(ctx *pulumi.Context,
	name string, args *ChannelTeamsArgs, opts ...pulumi.ResourceOption) (*ChannelTeams, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BotName == nil {
		return nil, errors.New("invalid value for required argument 'BotName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ChannelTeams
	err := ctx.RegisterResource("azure:bot/channelTeams:ChannelTeams", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannelTeams gets an existing ChannelTeams resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannelTeams(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelTeamsState, opts ...pulumi.ResourceOption) (*ChannelTeams, error) {
	var resource ChannelTeams
	err := ctx.ReadResource("azure:bot/channelTeams:ChannelTeams", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChannelTeams resources.
type channelTeamsState struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName *string `pulumi:"botName"`
	// Specifies the webhook for Microsoft Teams channel calls.
	CallingWebHook *string `pulumi:"callingWebHook"`
	// Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
	EnableCalling *bool `pulumi:"enableCalling"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ChannelTeamsState struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName pulumi.StringPtrInput
	// Specifies the webhook for Microsoft Teams channel calls.
	CallingWebHook pulumi.StringPtrInput
	// Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
	EnableCalling pulumi.BoolPtrInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (ChannelTeamsState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelTeamsState)(nil)).Elem()
}

type channelTeamsArgs struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName string `pulumi:"botName"`
	// Specifies the webhook for Microsoft Teams channel calls.
	CallingWebHook *string `pulumi:"callingWebHook"`
	// Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
	EnableCalling *bool `pulumi:"enableCalling"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ChannelTeams resource.
type ChannelTeamsArgs struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName pulumi.StringInput
	// Specifies the webhook for Microsoft Teams channel calls.
	CallingWebHook pulumi.StringPtrInput
	// Specifies whether to enable Microsoft Teams channel calls. This defaults to `false`.
	EnableCalling pulumi.BoolPtrInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (ChannelTeamsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelTeamsArgs)(nil)).Elem()
}

type ChannelTeamsInput interface {
	pulumi.Input

	ToChannelTeamsOutput() ChannelTeamsOutput
	ToChannelTeamsOutputWithContext(ctx context.Context) ChannelTeamsOutput
}

func (*ChannelTeams) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelTeams)(nil))
}

func (i *ChannelTeams) ToChannelTeamsOutput() ChannelTeamsOutput {
	return i.ToChannelTeamsOutputWithContext(context.Background())
}

func (i *ChannelTeams) ToChannelTeamsOutputWithContext(ctx context.Context) ChannelTeamsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelTeamsOutput)
}

func (i *ChannelTeams) ToChannelTeamsPtrOutput() ChannelTeamsPtrOutput {
	return i.ToChannelTeamsPtrOutputWithContext(context.Background())
}

func (i *ChannelTeams) ToChannelTeamsPtrOutputWithContext(ctx context.Context) ChannelTeamsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelTeamsPtrOutput)
}

type ChannelTeamsPtrInput interface {
	pulumi.Input

	ToChannelTeamsPtrOutput() ChannelTeamsPtrOutput
	ToChannelTeamsPtrOutputWithContext(ctx context.Context) ChannelTeamsPtrOutput
}

type ChannelTeamsOutput struct {
	*pulumi.OutputState
}

func (ChannelTeamsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelTeams)(nil))
}

func (o ChannelTeamsOutput) ToChannelTeamsOutput() ChannelTeamsOutput {
	return o
}

func (o ChannelTeamsOutput) ToChannelTeamsOutputWithContext(ctx context.Context) ChannelTeamsOutput {
	return o
}

type ChannelTeamsPtrOutput struct {
	*pulumi.OutputState
}

func (ChannelTeamsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelTeams)(nil))
}

func (o ChannelTeamsPtrOutput) ToChannelTeamsPtrOutput() ChannelTeamsPtrOutput {
	return o
}

func (o ChannelTeamsPtrOutput) ToChannelTeamsPtrOutputWithContext(ctx context.Context) ChannelTeamsPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ChannelTeamsOutput{})
	pulumi.RegisterOutputType(ChannelTeamsPtrOutput{})
}
