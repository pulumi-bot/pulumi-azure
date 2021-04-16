// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managementgroups

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Management Group.
//
// !> **Note:** Configuring `subscriptionIds` is not supported when using the `management.GroupSubscriptionAssociation` resource, results will be unpredictable.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.LookupSubscription(ctx, nil, nil)
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
// ## Import
//
// Management Groups can be imported using the `management group resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:managementgroups/managementGroup:ManagementGroup example /providers/Microsoft.Management/managementGroups/group1
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

func (*ManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroup)(nil))
}

func (i *ManagementGroup) ToManagementGroupOutput() ManagementGroupOutput {
	return i.ToManagementGroupOutputWithContext(context.Background())
}

func (i *ManagementGroup) ToManagementGroupOutputWithContext(ctx context.Context) ManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupOutput)
}

func (i *ManagementGroup) ToManagementGroupPtrOutput() ManagementGroupPtrOutput {
	return i.ToManagementGroupPtrOutputWithContext(context.Background())
}

func (i *ManagementGroup) ToManagementGroupPtrOutputWithContext(ctx context.Context) ManagementGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupPtrOutput)
}

type ManagementGroupPtrInput interface {
	pulumi.Input

	ToManagementGroupPtrOutput() ManagementGroupPtrOutput
	ToManagementGroupPtrOutputWithContext(ctx context.Context) ManagementGroupPtrOutput
}

type managementGroupPtrType ManagementGroupArgs

func (*managementGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroup)(nil))
}

func (i *managementGroupPtrType) ToManagementGroupPtrOutput() ManagementGroupPtrOutput {
	return i.ToManagementGroupPtrOutputWithContext(context.Background())
}

func (i *managementGroupPtrType) ToManagementGroupPtrOutputWithContext(ctx context.Context) ManagementGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupPtrOutput)
}

// ManagementGroupArrayInput is an input type that accepts ManagementGroupArray and ManagementGroupArrayOutput values.
// You can construct a concrete instance of `ManagementGroupArrayInput` via:
//
//          ManagementGroupArray{ ManagementGroupArgs{...} }
type ManagementGroupArrayInput interface {
	pulumi.Input

	ToManagementGroupArrayOutput() ManagementGroupArrayOutput
	ToManagementGroupArrayOutputWithContext(context.Context) ManagementGroupArrayOutput
}

type ManagementGroupArray []ManagementGroupInput

func (ManagementGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ManagementGroup)(nil))
}

func (i ManagementGroupArray) ToManagementGroupArrayOutput() ManagementGroupArrayOutput {
	return i.ToManagementGroupArrayOutputWithContext(context.Background())
}

func (i ManagementGroupArray) ToManagementGroupArrayOutputWithContext(ctx context.Context) ManagementGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupArrayOutput)
}

// ManagementGroupMapInput is an input type that accepts ManagementGroupMap and ManagementGroupMapOutput values.
// You can construct a concrete instance of `ManagementGroupMapInput` via:
//
//          ManagementGroupMap{ "key": ManagementGroupArgs{...} }
type ManagementGroupMapInput interface {
	pulumi.Input

	ToManagementGroupMapOutput() ManagementGroupMapOutput
	ToManagementGroupMapOutputWithContext(context.Context) ManagementGroupMapOutput
}

type ManagementGroupMap map[string]ManagementGroupInput

func (ManagementGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ManagementGroup)(nil))
}

func (i ManagementGroupMap) ToManagementGroupMapOutput() ManagementGroupMapOutput {
	return i.ToManagementGroupMapOutputWithContext(context.Background())
}

func (i ManagementGroupMap) ToManagementGroupMapOutputWithContext(ctx context.Context) ManagementGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupMapOutput)
}

type ManagementGroupOutput struct {
	*pulumi.OutputState
}

func (ManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroup)(nil))
}

func (o ManagementGroupOutput) ToManagementGroupOutput() ManagementGroupOutput {
	return o
}

func (o ManagementGroupOutput) ToManagementGroupOutputWithContext(ctx context.Context) ManagementGroupOutput {
	return o
}

func (o ManagementGroupOutput) ToManagementGroupPtrOutput() ManagementGroupPtrOutput {
	return o.ToManagementGroupPtrOutputWithContext(context.Background())
}

func (o ManagementGroupOutput) ToManagementGroupPtrOutputWithContext(ctx context.Context) ManagementGroupPtrOutput {
	return o.ApplyT(func(v ManagementGroup) *ManagementGroup {
		return &v
	}).(ManagementGroupPtrOutput)
}

type ManagementGroupPtrOutput struct {
	*pulumi.OutputState
}

func (ManagementGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroup)(nil))
}

func (o ManagementGroupPtrOutput) ToManagementGroupPtrOutput() ManagementGroupPtrOutput {
	return o
}

func (o ManagementGroupPtrOutput) ToManagementGroupPtrOutputWithContext(ctx context.Context) ManagementGroupPtrOutput {
	return o
}

type ManagementGroupArrayOutput struct{ *pulumi.OutputState }

func (ManagementGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroup)(nil))
}

func (o ManagementGroupArrayOutput) ToManagementGroupArrayOutput() ManagementGroupArrayOutput {
	return o
}

func (o ManagementGroupArrayOutput) ToManagementGroupArrayOutputWithContext(ctx context.Context) ManagementGroupArrayOutput {
	return o
}

func (o ManagementGroupArrayOutput) Index(i pulumi.IntInput) ManagementGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementGroup {
		return vs[0].([]ManagementGroup)[vs[1].(int)]
	}).(ManagementGroupOutput)
}

type ManagementGroupMapOutput struct{ *pulumi.OutputState }

func (ManagementGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagementGroup)(nil))
}

func (o ManagementGroupMapOutput) ToManagementGroupMapOutput() ManagementGroupMapOutput {
	return o
}

func (o ManagementGroupMapOutput) ToManagementGroupMapOutputWithContext(ctx context.Context) ManagementGroupMapOutput {
	return o
}

func (o ManagementGroupMapOutput) MapIndex(k pulumi.StringInput) ManagementGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagementGroup {
		return vs[0].(map[string]ManagementGroup)[vs[1].(string)]
	}).(ManagementGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupOutput{})
	pulumi.RegisterOutputType(ManagementGroupPtrOutput{})
	pulumi.RegisterOutputType(ManagementGroupArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupMapOutput{})
}
