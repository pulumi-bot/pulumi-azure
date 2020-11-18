// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package management

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Management Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleParent, err := management.NewGroup(ctx, "exampleParent", &management.GroupArgs{
// 			DisplayName: pulumi.String("ParentGroup"),
// 			SubscriptionIds: pulumi.StringArray{
// 				pulumi.String(current.SubscriptionId),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = management.NewGroup(ctx, "exampleChild", &management.GroupArgs{
// 			DisplayName:             pulumi.String("ChildGroup"),
// 			ParentManagementGroupId: exampleParent.ID(),
// 			SubscriptionIds: pulumi.StringArray{
// 				pulumi.String(current.SubscriptionId),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Group struct {
	pulumi.CustomResourceState

	// A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `name`
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Parent Management Group. Changing this forces a new resource to be created.
	ParentManagementGroupId pulumi.StringOutput `pulumi:"parentManagementGroupId"`
	// A list of Subscription GUIDs which should be assigned to the Management Group.
	SubscriptionIds pulumi.StringArrayOutput `pulumi:"subscriptionIds"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:managementgroups/managementGroup:ManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource Group
	err := ctx.RegisterResource("azure:management/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azure:management/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
	DisplayName *string `pulumi:"displayName"`
	// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `name`
	GroupId *string `pulumi:"groupId"`
	// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Parent Management Group. Changing this forces a new resource to be created.
	ParentManagementGroupId *string `pulumi:"parentManagementGroupId"`
	// A list of Subscription GUIDs which should be assigned to the Management Group.
	SubscriptionIds []string `pulumi:"subscriptionIds"`
}

type GroupState struct {
	// A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
	DisplayName pulumi.StringPtrInput
	// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `name`
	GroupId pulumi.StringPtrInput
	// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Parent Management Group. Changing this forces a new resource to be created.
	ParentManagementGroupId pulumi.StringPtrInput
	// A list of Subscription GUIDs which should be assigned to the Management Group.
	SubscriptionIds pulumi.StringArrayInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
	DisplayName *string `pulumi:"displayName"`
	// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `name`
	GroupId *string `pulumi:"groupId"`
	// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Parent Management Group. Changing this forces a new resource to be created.
	ParentManagementGroupId *string `pulumi:"parentManagementGroupId"`
	// A list of Subscription GUIDs which should be assigned to the Management Group.
	SubscriptionIds []string `pulumi:"subscriptionIds"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// A friendly name for this Management Group. If not specified, this'll be the same as the `name`.
	DisplayName pulumi.StringPtrInput
	// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `name`
	GroupId pulumi.StringPtrInput
	// The name or UUID for this Management Group, which needs to be unique across your tenant. A new UUID will be generated if not provided. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Parent Management Group. Changing this forces a new resource to be created.
	ParentManagementGroupId pulumi.StringPtrInput
	// A list of Subscription GUIDs which should be assigned to the Management Group.
	SubscriptionIds pulumi.StringArrayInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil)).Elem()
}

func (i Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct {
	*pulumi.OutputState
}

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupOutput)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}
