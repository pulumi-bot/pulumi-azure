// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bot

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Email integration for a Bot Channel
// 
// > **Note** A bot can only have a single Email Channel associated with it.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/bot_channel_email.html.markdown.
type ChannelEmail struct {
	s *pulumi.ResourceState
}

// NewChannelEmail registers a new resource with the given unique name, arguments, and options.
func NewChannelEmail(ctx *pulumi.Context,
	name string, args *ChannelEmailArgs, opts ...pulumi.ResourceOpt) (*ChannelEmail, error) {
	if args == nil || args.BotName == nil {
		return nil, errors.New("missing required argument 'BotName'")
	}
	if args == nil || args.EmailAddress == nil {
		return nil, errors.New("missing required argument 'EmailAddress'")
	}
	if args == nil || args.EmailPassword == nil {
		return nil, errors.New("missing required argument 'EmailPassword'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["botName"] = nil
		inputs["emailAddress"] = nil
		inputs["emailPassword"] = nil
		inputs["location"] = nil
		inputs["resourceGroupName"] = nil
	} else {
		inputs["botName"] = args.BotName
		inputs["emailAddress"] = args.EmailAddress
		inputs["emailPassword"] = args.EmailPassword
		inputs["location"] = args.Location
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	s, err := ctx.RegisterResource("azure:bot/channelEmail:ChannelEmail", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ChannelEmail{s: s}, nil
}

// GetChannelEmail gets an existing ChannelEmail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannelEmail(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ChannelEmailState, opts ...pulumi.ResourceOpt) (*ChannelEmail, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["botName"] = state.BotName
		inputs["emailAddress"] = state.EmailAddress
		inputs["emailPassword"] = state.EmailPassword
		inputs["location"] = state.Location
		inputs["resourceGroupName"] = state.ResourceGroupName
	}
	s, err := ctx.ReadResource("azure:bot/channelEmail:ChannelEmail", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ChannelEmail{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ChannelEmail) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ChannelEmail) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
func (r *ChannelEmail) BotName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["botName"])
}

// The email address that the Bot will authenticate with.
func (r *ChannelEmail) EmailAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["emailAddress"])
}

// The email password that the the Bot will authenticate with.
func (r *ChannelEmail) EmailPassword() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["emailPassword"])
}

// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *ChannelEmail) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
func (r *ChannelEmail) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Input properties used for looking up and filtering ChannelEmail resources.
type ChannelEmailState struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName interface{}
	// The email address that the Bot will authenticate with.
	EmailAddress interface{}
	// The email password that the the Bot will authenticate with.
	EmailPassword interface{}
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
}

// The set of arguments for constructing a ChannelEmail resource.
type ChannelEmailArgs struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName interface{}
	// The email address that the Bot will authenticate with.
	EmailAddress interface{}
	// The email password that the the Bot will authenticate with.
	EmailPassword interface{}
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
}