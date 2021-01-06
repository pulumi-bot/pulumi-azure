// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a policy rule definition on a management group or your provider subscription.
//
// Policy definitions do not take effect until they are assigned to a scope using a Policy Assignment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/policy"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := policy.NewDefinition(ctx, "policy", &policy.DefinitionArgs{
// 			DisplayName: pulumi.String("acceptance test policy definition"),
// 			Metadata:    pulumi.String(fmt.Sprintf("%v%v%v%v%v", "    {\n", "    \"category\": \"General\"\n", "    }\n", "\n", "\n")),
// 			Mode:        pulumi.String("Indexed"),
// 			Parameters: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "	{\n", "    \"allowedLocations\": {\n", "      \"type\": \"Array\",\n", "      \"metadata\": {\n", "        \"description\": \"The list of allowed locations for resources.\",\n", "        \"displayName\": \"Allowed locations\",\n", "        \"strongType\": \"location\"\n", "      }\n", "    }\n", "  }\n", "\n")),
// 			PolicyRule: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "	{\n", "    \"if\": {\n", "      \"not\": {\n", "        \"field\": \"location\",\n", "        \"in\": \"[parameters('allowedLocations')]\"\n", "      }\n", "    },\n", "    \"then\": {\n", "      \"effect\": \"audit\"\n", "    }\n", "  }\n", "\n")),
// 			PolicyType: pulumi.String("Custom"),
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
// Policy Definitions can be imported using the `policy name`, e.g.
//
// ```sh
//  $ pulumi import azure:policy/definition:Definition examplePolicy /subscriptions/<SUBSCRIPTION_ID>/providers/Microsoft.Authorization/policyDefinitions/<POLICY_NAME>
// ```
//
//  or
//
// ```sh
//  $ pulumi import azure:policy/definition:Definition examplePolicy /providers/Microsoft.Management/managementgroups/<MANGAGEMENT_GROUP_ID>/providers/Microsoft.Authorization/policyDefinitions/<POLICY_NAME>
// ```
type Definition struct {
	pulumi.CustomResourceState

	// The description of the policy definition.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `management_group_name`
	ManagementGroupId pulumi.StringOutput `pulumi:"managementGroupId"`
	// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupName pulumi.StringOutput `pulumi:"managementGroupName"`
	// The metadata for the policy definition. This
	// is a JSON string representing additional metadata that should be stored
	// with the policy definition.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The policy mode that allows you to specify which resource
	// types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters for the policy definition. This field
	// is a JSON string that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrOutput `pulumi:"parameters"`
	// The policy rule for the policy definition. This
	// is a JSON string representing the rule that contains an if and
	// a then block.
	PolicyRule pulumi.StringPtrOutput `pulumi:"policyRule"`
	// The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
}

// NewDefinition registers a new resource with the given unique name, arguments, and options.
func NewDefinition(ctx *pulumi.Context,
	name string, args *DefinitionArgs, opts ...pulumi.ResourceOption) (*Definition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.PolicyType == nil {
		return nil, errors.New("invalid value for required argument 'PolicyType'")
	}
	var resource Definition
	err := ctx.RegisterResource("azure:policy/definition:Definition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefinition gets an existing Definition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefinitionState, opts ...pulumi.ResourceOption) (*Definition, error) {
	var resource Definition
	err := ctx.ReadResource("azure:policy/definition:Definition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Definition resources.
type definitionState struct {
	// The description of the policy definition.
	Description *string `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName *string `pulumi:"displayName"`
	// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `management_group_name`
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupName *string `pulumi:"managementGroupName"`
	// The metadata for the policy definition. This
	// is a JSON string representing additional metadata that should be stored
	// with the policy definition.
	Metadata *string `pulumi:"metadata"`
	// The policy mode that allows you to specify which resource
	// types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
	Mode *string `pulumi:"mode"`
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Parameters for the policy definition. This field
	// is a JSON string that allows you to parameterize your policy definition.
	Parameters *string `pulumi:"parameters"`
	// The policy rule for the policy definition. This
	// is a JSON string representing the rule that contains an if and
	// a then block.
	PolicyRule *string `pulumi:"policyRule"`
	// The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
	PolicyType *string `pulumi:"policyType"`
}

type DefinitionState struct {
	// The description of the policy definition.
	Description pulumi.StringPtrInput
	// The display name of the policy definition.
	DisplayName pulumi.StringPtrInput
	// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `management_group_name`
	ManagementGroupId pulumi.StringPtrInput
	// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupName pulumi.StringPtrInput
	// The metadata for the policy definition. This
	// is a JSON string representing additional metadata that should be stored
	// with the policy definition.
	Metadata pulumi.StringPtrInput
	// The policy mode that allows you to specify which resource
	// types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
	Mode pulumi.StringPtrInput
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Parameters for the policy definition. This field
	// is a JSON string that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrInput
	// The policy rule for the policy definition. This
	// is a JSON string representing the rule that contains an if and
	// a then block.
	PolicyRule pulumi.StringPtrInput
	// The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
	PolicyType pulumi.StringPtrInput
}

func (DefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionState)(nil)).Elem()
}

type definitionArgs struct {
	// The description of the policy definition.
	Description *string `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName string `pulumi:"displayName"`
	// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `management_group_name`
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupName *string `pulumi:"managementGroupName"`
	// The metadata for the policy definition. This
	// is a JSON string representing additional metadata that should be stored
	// with the policy definition.
	Metadata *string `pulumi:"metadata"`
	// The policy mode that allows you to specify which resource
	// types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
	Mode string `pulumi:"mode"`
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Parameters for the policy definition. This field
	// is a JSON string that allows you to parameterize your policy definition.
	Parameters *string `pulumi:"parameters"`
	// The policy rule for the policy definition. This
	// is a JSON string representing the rule that contains an if and
	// a then block.
	PolicyRule *string `pulumi:"policyRule"`
	// The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
	PolicyType string `pulumi:"policyType"`
}

// The set of arguments for constructing a Definition resource.
type DefinitionArgs struct {
	// The description of the policy definition.
	Description pulumi.StringPtrInput
	// The display name of the policy definition.
	DisplayName pulumi.StringInput
	// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	//
	// Deprecated: Deprecated in favour of `management_group_name`
	ManagementGroupId pulumi.StringPtrInput
	// The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupName pulumi.StringPtrInput
	// The metadata for the policy definition. This
	// is a JSON string representing additional metadata that should be stored
	// with the policy definition.
	Metadata pulumi.StringPtrInput
	// The policy mode that allows you to specify which resource
	// types will be evaluated. Possible values are `All`, `Indexed`, `Microsoft.ContainerService.Data`, `Microsoft.CustomerLockbox.Data`, `Microsoft.DataCatalog.Data`, `Microsoft.KeyVault.Data`, `Microsoft.Kubernetes.Data`, `Microsoft.MachineLearningServices.Data`, `Microsoft.Network.Data` and `Microsoft.Synapse.Data`.
	Mode pulumi.StringInput
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Parameters for the policy definition. This field
	// is a JSON string that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrInput
	// The policy rule for the policy definition. This
	// is a JSON string representing the rule that contains an if and
	// a then block.
	PolicyRule pulumi.StringPtrInput
	// The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
	PolicyType pulumi.StringInput
}

func (DefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionArgs)(nil)).Elem()
}

type DefinitionInput interface {
	pulumi.Input

	ToDefinitionOutput() DefinitionOutput
	ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput
}

func (*Definition) ElementType() reflect.Type {
	return reflect.TypeOf((*Definition)(nil))
}

func (i *Definition) ToDefinitionOutput() DefinitionOutput {
	return i.ToDefinitionOutputWithContext(context.Background())
}

func (i *Definition) ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionOutput)
}

func (i *Definition) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return i.ToDefinitionPtrOutputWithContext(context.Background())
}

func (i *Definition) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionPtrOutput)
}

type DefinitionPtrInput interface {
	pulumi.Input

	ToDefinitionPtrOutput() DefinitionPtrOutput
	ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput
}

type definitionPtrType DefinitionArgs

func (*definitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Definition)(nil))
}

func (i *definitionPtrType) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return i.ToDefinitionPtrOutputWithContext(context.Background())
}

func (i *definitionPtrType) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionOutput).ToDefinitionPtrOutput()
}

// DefinitionArrayInput is an input type that accepts DefinitionArray and DefinitionArrayOutput values.
// You can construct a concrete instance of `DefinitionArrayInput` via:
//
//          DefinitionArray{ DefinitionArgs{...} }
type DefinitionArrayInput interface {
	pulumi.Input

	ToDefinitionArrayOutput() DefinitionArrayOutput
	ToDefinitionArrayOutputWithContext(context.Context) DefinitionArrayOutput
}

type DefinitionArray []DefinitionInput

func (DefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Definition)(nil))
}

func (i DefinitionArray) ToDefinitionArrayOutput() DefinitionArrayOutput {
	return i.ToDefinitionArrayOutputWithContext(context.Background())
}

func (i DefinitionArray) ToDefinitionArrayOutputWithContext(ctx context.Context) DefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionArrayOutput)
}

// DefinitionMapInput is an input type that accepts DefinitionMap and DefinitionMapOutput values.
// You can construct a concrete instance of `DefinitionMapInput` via:
//
//          DefinitionMap{ "key": DefinitionArgs{...} }
type DefinitionMapInput interface {
	pulumi.Input

	ToDefinitionMapOutput() DefinitionMapOutput
	ToDefinitionMapOutputWithContext(context.Context) DefinitionMapOutput
}

type DefinitionMap map[string]DefinitionInput

func (DefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Definition)(nil))
}

func (i DefinitionMap) ToDefinitionMapOutput() DefinitionMapOutput {
	return i.ToDefinitionMapOutputWithContext(context.Background())
}

func (i DefinitionMap) ToDefinitionMapOutputWithContext(ctx context.Context) DefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionMapOutput)
}

type DefinitionOutput struct {
	*pulumi.OutputState
}

func (DefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Definition)(nil))
}

func (o DefinitionOutput) ToDefinitionOutput() DefinitionOutput {
	return o
}

func (o DefinitionOutput) ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput {
	return o
}

func (o DefinitionOutput) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return o.ToDefinitionPtrOutputWithContext(context.Background())
}

func (o DefinitionOutput) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return o.ApplyT(func(v Definition) *Definition {
		return &v
	}).(DefinitionPtrOutput)
}

type DefinitionPtrOutput struct {
	*pulumi.OutputState
}

func (DefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Definition)(nil))
}

func (o DefinitionPtrOutput) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return o
}

func (o DefinitionPtrOutput) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return o
}

type DefinitionArrayOutput struct{ *pulumi.OutputState }

func (DefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Definition)(nil))
}

func (o DefinitionArrayOutput) ToDefinitionArrayOutput() DefinitionArrayOutput {
	return o
}

func (o DefinitionArrayOutput) ToDefinitionArrayOutputWithContext(ctx context.Context) DefinitionArrayOutput {
	return o
}

func (o DefinitionArrayOutput) Index(i pulumi.IntInput) DefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Definition {
		return vs[0].([]Definition)[vs[1].(int)]
	}).(DefinitionOutput)
}

type DefinitionMapOutput struct{ *pulumi.OutputState }

func (DefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Definition)(nil))
}

func (o DefinitionMapOutput) ToDefinitionMapOutput() DefinitionMapOutput {
	return o
}

func (o DefinitionMapOutput) ToDefinitionMapOutputWithContext(ctx context.Context) DefinitionMapOutput {
	return o
}

func (o DefinitionMapOutput) MapIndex(k pulumi.StringInput) DefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Definition {
		return vs[0].(map[string]Definition)[vs[1].(string)]
	}).(DefinitionOutput)
}

func init() {
	pulumi.RegisterOutputType(DefinitionOutput{})
	pulumi.RegisterOutputType(DefinitionPtrOutput{})
	pulumi.RegisterOutputType(DefinitionArrayOutput{})
	pulumi.RegisterOutputType(DefinitionMapOutput{})
}
