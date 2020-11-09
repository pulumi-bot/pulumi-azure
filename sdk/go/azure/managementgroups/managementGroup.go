// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managementgroups

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
//
// Deprecated: azure.managementgroups.ManagementGroup has been deprecated in favor of azure.management.Group
type ManagementGroup struct {
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

// NewManagementGroup registers a new resource with the given unique name, arguments, and options.
func NewManagementGroup(ctx *pulumi.Context,
	name string, args *ManagementGroupArgs, opts ...pulumi.ResourceOption) (*ManagementGroup, error) {
	if args == nil {
		args = &ManagementGroupArgs{}
	}
	var resource ManagementGroup
	err := ctx.RegisterResource("azure:managementgroups/managementGroup:ManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementGroup gets an existing ManagementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementGroupState, opts ...pulumi.ResourceOption) (*ManagementGroup, error) {
	var resource ManagementGroup
	err := ctx.ReadResource("azure:managementgroups/managementGroup:ManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementGroup resources.
type managementGroupState struct {
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

type ManagementGroupState struct {
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

func (ManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupState)(nil)).Elem()
}

type managementGroupArgs struct {
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

// The set of arguments for constructing a ManagementGroup resource.
type ManagementGroupArgs struct {
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

func (ManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupArgs)(nil)).Elem()
}

type ManagementGroupInput interface {
	pulumi.Input

	ToManagementGroupOutput() ManagementGroupOutput
	ToManagementGroupOutputWithContext(ctx context.Context) ManagementGroupOutput
}

func (ManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroup)(nil)).Elem()
}

func (i ManagementGroup) ToManagementGroupOutput() ManagementGroupOutput {
	return i.ToManagementGroupOutputWithContext(context.Background())
}

func (i ManagementGroup) ToManagementGroupOutputWithContext(ctx context.Context) ManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupOutput)
}

type ManagementGroupOutput struct {
	*pulumi.OutputState
}

func (ManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupOutput)(nil)).Elem()
}

func (o ManagementGroupOutput) ToManagementGroupOutput() ManagementGroupOutput {
	return o
}

func (o ManagementGroupOutput) ToManagementGroupOutputWithContext(ctx context.Context) ManagementGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupOutput{})
}
