// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Assigns a given Principal (User or Group) to a given Role.
//
// ## Example Usage
// ### Using A Built-In Role)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := core.GetSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleClientConfig, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Scope:              pulumi.String(primary.Id),
// 			RoleDefinitionName: pulumi.String("Reader"),
// 			PrincipalId:        pulumi.String(exampleClientConfig.ObjectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Custom Role & Service Principal)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := core.GetSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleClientConfig, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleRoleDefinition, err := authorization.NewRoleDefinition(ctx, "exampleRoleDefinition", &authorization.RoleDefinitionArgs{
// 			RoleDefinitionId: pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			Scope:            pulumi.String(primary.Id),
// 			Permissions: authorization.RoleDefinitionPermissionArray{
// 				&authorization.RoleDefinitionPermissionArgs{
// 					Actions: pulumi.StringArray{
// 						pulumi.String("Microsoft.Resources/subscriptions/resourceGroups/read"),
// 					},
// 					NotActions: []interface{}{},
// 				},
// 			},
// 			AssignableScopes: pulumi.StringArray{
// 				pulumi.String(primary.Id),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Name:             pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			Scope:            pulumi.String(primary.Id),
// 			RoleDefinitionId: exampleRoleDefinition.RoleDefinitionResourceId,
// 			PrincipalId:      pulumi.String(exampleClientConfig.ObjectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Custom Role & User)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := core.GetSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleClientConfig, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleRoleDefinition, err := authorization.NewRoleDefinition(ctx, "exampleRoleDefinition", &authorization.RoleDefinitionArgs{
// 			RoleDefinitionId: pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			Scope:            pulumi.String(primary.Id),
// 			Permissions: authorization.RoleDefinitionPermissionArray{
// 				&authorization.RoleDefinitionPermissionArgs{
// 					Actions: pulumi.StringArray{
// 						pulumi.String("Microsoft.Resources/subscriptions/resourceGroups/read"),
// 					},
// 					NotActions: []interface{}{},
// 				},
// 			},
// 			AssignableScopes: pulumi.StringArray{
// 				pulumi.String(primary.Id),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Name:             pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			Scope:            pulumi.String(primary.Id),
// 			RoleDefinitionId: exampleRoleDefinition.RoleDefinitionResourceId,
// 			PrincipalId:      pulumi.String(exampleClientConfig.ObjectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Custom Role & Management Group)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := core.GetSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleClientConfig, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = management.LookupGroup(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleRoleDefinition, err := authorization.NewRoleDefinition(ctx, "exampleRoleDefinition", &authorization.RoleDefinitionArgs{
// 			RoleDefinitionId: pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			Scope:            pulumi.String(primary.Id),
// 			Permissions: authorization.RoleDefinitionPermissionArray{
// 				&authorization.RoleDefinitionPermissionArgs{
// 					Actions: pulumi.StringArray{
// 						pulumi.String("Microsoft.Resources/subscriptions/resourceGroups/read"),
// 					},
// 					NotActions: []interface{}{},
// 				},
// 			},
// 			AssignableScopes: pulumi.StringArray{
// 				pulumi.String(primary.Id),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Name:             pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			Scope:            pulumi.Any(data.Azurerm_management_group.Primary.Id),
// 			RoleDefinitionId: exampleRoleDefinition.RoleDefinitionResourceId,
// 			PrincipalId:      pulumi.String(exampleClientConfig.ObjectId),
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
// Role Assignments can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:authorization/assignment:Assignment example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleAssignments/00000000-0000-0000-0000-000000000000
// ```
//
//  - for scope `Subscription`, the id format is `/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleAssignments/00000000-0000-0000-0000-000000000000` - for scope `Resource Group`, the id format is `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Authorization/roleAssignments/00000000-0000-0000-0000-000000000000`
type Assignment struct {
	pulumi.CustomResourceState

	// A unique UUID/GUID for this Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Principal (User, Group or Service Principal) to assign the Role Definition to. Changing this forces a new resource to be created.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The type of the `principalId`, e.g. User, Group, Service Principal, Application, etc.
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// The Scoped-ID of the Role Definition. Changing this forces a new resource to be created. Conflicts with `roleDefinitionName`.
	RoleDefinitionId pulumi.StringOutput `pulumi:"roleDefinitionId"`
	// The name of a built-in Role. Changing this forces a new resource to be created. Conflicts with `roleDefinitionId`.
	RoleDefinitionName pulumi.StringOutput `pulumi:"roleDefinitionName"`
	// The scope at which the Role Assignment applies to, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`, or `/providers/Microsoft.Management/managementGroups/myMG`. Changing this forces a new resource to be created.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// If the `principalId` is a newly provisioned `Service Principal` set this value to `true` to skip the `Azure Active Directory` check which may fail due to replication lag. This argument is only valid if the `principalId` is a `Service Principal` identity. If it is not a `Service Principal` identity it will cause the role assignment to fail. Defaults to `false`.
	SkipServicePrincipalAadCheck pulumi.BoolOutput `pulumi:"skipServicePrincipalAadCheck"`
}

// NewAssignment registers a new resource with the given unique name, arguments, and options.
func NewAssignment(ctx *pulumi.Context,
	name string, args *AssignmentArgs, opts ...pulumi.ResourceOption) (*Assignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:role/assignment:Assignment"),
		},
	})
	opts = append(opts, aliases)
	var resource Assignment
	err := ctx.RegisterResource("azure:authorization/assignment:Assignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssignment gets an existing Assignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentState, opts ...pulumi.ResourceOption) (*Assignment, error) {
	var resource Assignment
	err := ctx.ReadResource("azure:authorization/assignment:Assignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assignment resources.
type assignmentState struct {
	// A unique UUID/GUID for this Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Principal (User, Group or Service Principal) to assign the Role Definition to. Changing this forces a new resource to be created.
	PrincipalId *string `pulumi:"principalId"`
	// The type of the `principalId`, e.g. User, Group, Service Principal, Application, etc.
	PrincipalType *string `pulumi:"principalType"`
	// The Scoped-ID of the Role Definition. Changing this forces a new resource to be created. Conflicts with `roleDefinitionName`.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The name of a built-in Role. Changing this forces a new resource to be created. Conflicts with `roleDefinitionId`.
	RoleDefinitionName *string `pulumi:"roleDefinitionName"`
	// The scope at which the Role Assignment applies to, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`, or `/providers/Microsoft.Management/managementGroups/myMG`. Changing this forces a new resource to be created.
	Scope *string `pulumi:"scope"`
	// If the `principalId` is a newly provisioned `Service Principal` set this value to `true` to skip the `Azure Active Directory` check which may fail due to replication lag. This argument is only valid if the `principalId` is a `Service Principal` identity. If it is not a `Service Principal` identity it will cause the role assignment to fail. Defaults to `false`.
	SkipServicePrincipalAadCheck *bool `pulumi:"skipServicePrincipalAadCheck"`
}

type AssignmentState struct {
	// A unique UUID/GUID for this Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Principal (User, Group or Service Principal) to assign the Role Definition to. Changing this forces a new resource to be created.
	PrincipalId pulumi.StringPtrInput
	// The type of the `principalId`, e.g. User, Group, Service Principal, Application, etc.
	PrincipalType pulumi.StringPtrInput
	// The Scoped-ID of the Role Definition. Changing this forces a new resource to be created. Conflicts with `roleDefinitionName`.
	RoleDefinitionId pulumi.StringPtrInput
	// The name of a built-in Role. Changing this forces a new resource to be created. Conflicts with `roleDefinitionId`.
	RoleDefinitionName pulumi.StringPtrInput
	// The scope at which the Role Assignment applies to, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`, or `/providers/Microsoft.Management/managementGroups/myMG`. Changing this forces a new resource to be created.
	Scope pulumi.StringPtrInput
	// If the `principalId` is a newly provisioned `Service Principal` set this value to `true` to skip the `Azure Active Directory` check which may fail due to replication lag. This argument is only valid if the `principalId` is a `Service Principal` identity. If it is not a `Service Principal` identity it will cause the role assignment to fail. Defaults to `false`.
	SkipServicePrincipalAadCheck pulumi.BoolPtrInput
}

func (AssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentState)(nil)).Elem()
}

type assignmentArgs struct {
	// A unique UUID/GUID for this Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Principal (User, Group or Service Principal) to assign the Role Definition to. Changing this forces a new resource to be created.
	PrincipalId string `pulumi:"principalId"`
	// The Scoped-ID of the Role Definition. Changing this forces a new resource to be created. Conflicts with `roleDefinitionName`.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The name of a built-in Role. Changing this forces a new resource to be created. Conflicts with `roleDefinitionId`.
	RoleDefinitionName *string `pulumi:"roleDefinitionName"`
	// The scope at which the Role Assignment applies to, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`, or `/providers/Microsoft.Management/managementGroups/myMG`. Changing this forces a new resource to be created.
	Scope string `pulumi:"scope"`
	// If the `principalId` is a newly provisioned `Service Principal` set this value to `true` to skip the `Azure Active Directory` check which may fail due to replication lag. This argument is only valid if the `principalId` is a `Service Principal` identity. If it is not a `Service Principal` identity it will cause the role assignment to fail. Defaults to `false`.
	SkipServicePrincipalAadCheck *bool `pulumi:"skipServicePrincipalAadCheck"`
}

// The set of arguments for constructing a Assignment resource.
type AssignmentArgs struct {
	// A unique UUID/GUID for this Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Principal (User, Group or Service Principal) to assign the Role Definition to. Changing this forces a new resource to be created.
	PrincipalId pulumi.StringInput
	// The Scoped-ID of the Role Definition. Changing this forces a new resource to be created. Conflicts with `roleDefinitionName`.
	RoleDefinitionId pulumi.StringPtrInput
	// The name of a built-in Role. Changing this forces a new resource to be created. Conflicts with `roleDefinitionId`.
	RoleDefinitionName pulumi.StringPtrInput
	// The scope at which the Role Assignment applies to, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`, or `/providers/Microsoft.Management/managementGroups/myMG`. Changing this forces a new resource to be created.
	Scope pulumi.StringInput
	// If the `principalId` is a newly provisioned `Service Principal` set this value to `true` to skip the `Azure Active Directory` check which may fail due to replication lag. This argument is only valid if the `principalId` is a `Service Principal` identity. If it is not a `Service Principal` identity it will cause the role assignment to fail. Defaults to `false`.
	SkipServicePrincipalAadCheck pulumi.BoolPtrInput
}

func (AssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentArgs)(nil)).Elem()
}

type AssignmentInput interface {
	pulumi.Input

	ToAssignmentOutput() AssignmentOutput
	ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput
}

func (*Assignment) ElementType() reflect.Type {
	return reflect.TypeOf((*Assignment)(nil))
}

func (i *Assignment) ToAssignmentOutput() AssignmentOutput {
	return i.ToAssignmentOutputWithContext(context.Background())
}

func (i *Assignment) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentOutput)
}

func (i *Assignment) ToAssignmentPtrOutput() AssignmentPtrOutput {
	return i.ToAssignmentPtrOutputWithContext(context.Background())
}

func (i *Assignment) ToAssignmentPtrOutputWithContext(ctx context.Context) AssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPtrOutput)
}

type AssignmentPtrInput interface {
	pulumi.Input

	ToAssignmentPtrOutput() AssignmentPtrOutput
	ToAssignmentPtrOutputWithContext(ctx context.Context) AssignmentPtrOutput
}

type assignmentPtrType AssignmentArgs

func (*assignmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil))
}

func (i *assignmentPtrType) ToAssignmentPtrOutput() AssignmentPtrOutput {
	return i.ToAssignmentPtrOutputWithContext(context.Background())
}

func (i *assignmentPtrType) ToAssignmentPtrOutputWithContext(ctx context.Context) AssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPtrOutput)
}

// AssignmentArrayInput is an input type that accepts AssignmentArray and AssignmentArrayOutput values.
// You can construct a concrete instance of `AssignmentArrayInput` via:
//
//          AssignmentArray{ AssignmentArgs{...} }
type AssignmentArrayInput interface {
	pulumi.Input

	ToAssignmentArrayOutput() AssignmentArrayOutput
	ToAssignmentArrayOutputWithContext(context.Context) AssignmentArrayOutput
}

type AssignmentArray []AssignmentInput

func (AssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Assignment)(nil))
}

func (i AssignmentArray) ToAssignmentArrayOutput() AssignmentArrayOutput {
	return i.ToAssignmentArrayOutputWithContext(context.Background())
}

func (i AssignmentArray) ToAssignmentArrayOutputWithContext(ctx context.Context) AssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentArrayOutput)
}

// AssignmentMapInput is an input type that accepts AssignmentMap and AssignmentMapOutput values.
// You can construct a concrete instance of `AssignmentMapInput` via:
//
//          AssignmentMap{ "key": AssignmentArgs{...} }
type AssignmentMapInput interface {
	pulumi.Input

	ToAssignmentMapOutput() AssignmentMapOutput
	ToAssignmentMapOutputWithContext(context.Context) AssignmentMapOutput
}

type AssignmentMap map[string]AssignmentInput

func (AssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Assignment)(nil))
}

func (i AssignmentMap) ToAssignmentMapOutput() AssignmentMapOutput {
	return i.ToAssignmentMapOutputWithContext(context.Background())
}

func (i AssignmentMap) ToAssignmentMapOutputWithContext(ctx context.Context) AssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentMapOutput)
}

type AssignmentOutput struct {
	*pulumi.OutputState
}

func (AssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Assignment)(nil))
}

func (o AssignmentOutput) ToAssignmentOutput() AssignmentOutput {
	return o
}

func (o AssignmentOutput) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return o
}

func (o AssignmentOutput) ToAssignmentPtrOutput() AssignmentPtrOutput {
	return o.ToAssignmentPtrOutputWithContext(context.Background())
}

func (o AssignmentOutput) ToAssignmentPtrOutputWithContext(ctx context.Context) AssignmentPtrOutput {
	return o.ApplyT(func(v Assignment) *Assignment {
		return &v
	}).(AssignmentPtrOutput)
}

type AssignmentPtrOutput struct {
	*pulumi.OutputState
}

func (AssignmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil))
}

func (o AssignmentPtrOutput) ToAssignmentPtrOutput() AssignmentPtrOutput {
	return o
}

func (o AssignmentPtrOutput) ToAssignmentPtrOutputWithContext(ctx context.Context) AssignmentPtrOutput {
	return o
}

type AssignmentArrayOutput struct{ *pulumi.OutputState }

func (AssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Assignment)(nil))
}

func (o AssignmentArrayOutput) ToAssignmentArrayOutput() AssignmentArrayOutput {
	return o
}

func (o AssignmentArrayOutput) ToAssignmentArrayOutputWithContext(ctx context.Context) AssignmentArrayOutput {
	return o
}

func (o AssignmentArrayOutput) Index(i pulumi.IntInput) AssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Assignment {
		return vs[0].([]Assignment)[vs[1].(int)]
	}).(AssignmentOutput)
}

type AssignmentMapOutput struct{ *pulumi.OutputState }

func (AssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Assignment)(nil))
}

func (o AssignmentMapOutput) ToAssignmentMapOutput() AssignmentMapOutput {
	return o
}

func (o AssignmentMapOutput) ToAssignmentMapOutputWithContext(ctx context.Context) AssignmentMapOutput {
	return o
}

func (o AssignmentMapOutput) MapIndex(k pulumi.StringInput) AssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Assignment {
		return vs[0].(map[string]Assignment)[vs[1].(string)]
	}).(AssignmentOutput)
}

func init() {
	pulumi.RegisterOutputType(AssignmentOutput{})
	pulumi.RegisterOutputType(AssignmentPtrOutput{})
	pulumi.RegisterOutputType(AssignmentArrayOutput{})
	pulumi.RegisterOutputType(AssignmentMapOutput{})
}
