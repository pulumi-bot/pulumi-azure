// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Firewall Policy.
type FirewallPolicy struct {
	pulumi.CustomResourceState

	// The ID of the base Firewall Policy.
	BasePolicyId pulumi.StringPtrOutput `pulumi:"basePolicyId"`
	// A list of reference to child Firewall Policies of this Firewall Policy.
	ChildPolicies pulumi.StringArrayOutput `pulumi:"childPolicies"`
	// A `dns` block as defined below.
	Dns FirewallPolicyDnsPtrOutput `pulumi:"dns"`
	// A list of references to Azure Firewalls that this Firewall Policy is associated with.
	Firewalls pulumi.StringArrayOutput `pulumi:"firewalls"`
	// The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A list of references to Firewall Policy Rule Collection Groups that belongs to this Firewall Policy.
	RuleCollectionGroups pulumi.StringArrayOutput `pulumi:"ruleCollectionGroups"`
	// A mapping of tags which should be assigned to the Firewall Policy.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `threatIntelligenceAllowlist` block as defined below.
	ThreatIntelligenceAllowlist FirewallPolicyThreatIntelligenceAllowlistPtrOutput `pulumi:"threatIntelligenceAllowlist"`
	// The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
	ThreatIntelligenceMode pulumi.StringPtrOutput `pulumi:"threatIntelligenceMode"`
}

// NewFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &FirewallPolicyArgs{}
	}
	var resource FirewallPolicy
	err := ctx.RegisterResource("azure:network/firewallPolicy:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy gets an existing FirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("azure:network/firewallPolicy:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy resources.
type firewallPolicyState struct {
	// The ID of the base Firewall Policy.
	BasePolicyId *string `pulumi:"basePolicyId"`
	// A list of reference to child Firewall Policies of this Firewall Policy.
	ChildPolicies []string `pulumi:"childPolicies"`
	// A `dns` block as defined below.
	Dns *FirewallPolicyDns `pulumi:"dns"`
	// A list of references to Azure Firewalls that this Firewall Policy is associated with.
	Firewalls []string `pulumi:"firewalls"`
	// The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A list of references to Firewall Policy Rule Collection Groups that belongs to this Firewall Policy.
	RuleCollectionGroups []string `pulumi:"ruleCollectionGroups"`
	// A mapping of tags which should be assigned to the Firewall Policy.
	Tags map[string]string `pulumi:"tags"`
	// A `threatIntelligenceAllowlist` block as defined below.
	ThreatIntelligenceAllowlist *FirewallPolicyThreatIntelligenceAllowlist `pulumi:"threatIntelligenceAllowlist"`
	// The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
	ThreatIntelligenceMode *string `pulumi:"threatIntelligenceMode"`
}

type FirewallPolicyState struct {
	// The ID of the base Firewall Policy.
	BasePolicyId pulumi.StringPtrInput
	// A list of reference to child Firewall Policies of this Firewall Policy.
	ChildPolicies pulumi.StringArrayInput
	// A `dns` block as defined below.
	Dns FirewallPolicyDnsPtrInput
	// A list of references to Azure Firewalls that this Firewall Policy is associated with.
	Firewalls pulumi.StringArrayInput
	// The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A list of references to Firewall Policy Rule Collection Groups that belongs to this Firewall Policy.
	RuleCollectionGroups pulumi.StringArrayInput
	// A mapping of tags which should be assigned to the Firewall Policy.
	Tags pulumi.StringMapInput
	// A `threatIntelligenceAllowlist` block as defined below.
	ThreatIntelligenceAllowlist FirewallPolicyThreatIntelligenceAllowlistPtrInput
	// The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
	ThreatIntelligenceMode pulumi.StringPtrInput
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
	// The ID of the base Firewall Policy.
	BasePolicyId *string `pulumi:"basePolicyId"`
	// A `dns` block as defined below.
	Dns *FirewallPolicyDns `pulumi:"dns"`
	// The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Firewall Policy.
	Tags map[string]string `pulumi:"tags"`
	// A `threatIntelligenceAllowlist` block as defined below.
	ThreatIntelligenceAllowlist *FirewallPolicyThreatIntelligenceAllowlist `pulumi:"threatIntelligenceAllowlist"`
	// The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
	ThreatIntelligenceMode *string `pulumi:"threatIntelligenceMode"`
}

// The set of arguments for constructing a FirewallPolicy resource.
type FirewallPolicyArgs struct {
	// The ID of the base Firewall Policy.
	BasePolicyId pulumi.StringPtrInput
	// A `dns` block as defined below.
	Dns FirewallPolicyDnsPtrInput
	// The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the Firewall Policy.
	Tags pulumi.StringMapInput
	// A `threatIntelligenceAllowlist` block as defined below.
	ThreatIntelligenceAllowlist FirewallPolicyThreatIntelligenceAllowlistPtrInput
	// The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
	ThreatIntelligenceMode pulumi.StringPtrInput
}

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}
