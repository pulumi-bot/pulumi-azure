// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Policy within a Dev Test Policy Set.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/devtest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLab, err := devtest.NewLab(ctx, "exampleLab", &devtest.LabArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Tags: map[string]interface{}{
// 				"Sydney": "Australia",
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = devtest.NewPolicy(ctx, "examplePolicy", &devtest.PolicyArgs{
// 			PolicySetName:     pulumi.String("default"),
// 			LabName:           exampleLab.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			FactData:          pulumi.String(""),
// 			Threshold:         pulumi.String("999"),
// 			EvaluatorType:     pulumi.String("MaxValuePolicy"),
// 			Tags: map[string]interface{}{
// 				"Acceptance": "Test",
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Policy struct {
	pulumi.CustomResourceState

	// A description for the Policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Evaluation Type used for this Policy. Possible values include: 'AllowedValuesPolicy', 'MaxValuePolicy'. Changing this forces a new resource to be created.
	EvaluatorType pulumi.StringOutput `pulumi:"evaluatorType"`
	// The Fact Data for this Policy.
	FactData pulumi.StringPtrOutput `pulumi:"factData"`
	// Specifies the name of the Dev Test Lab in which the Policy should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringOutput `pulumi:"labName"`
	// Specifies the name of the Dev Test Policy. Possible values are `GalleryImage`, `LabPremiumVmCount`, `LabTargetCost`, `LabVmCount`, `LabVmSize`, `UserOwnedLabPremiumVmCount`, `UserOwnedLabVmCount` and `UserOwnedLabVmCountInSubnet`. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Policy Set within the Dev Test Lab where this policy should be created. Changing this forces a new resource to be created.
	PolicySetName pulumi.StringOutput `pulumi:"policySetName"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Threshold for this Policy.
	Threshold pulumi.StringOutput `pulumi:"threshold"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil || args.EvaluatorType == nil {
		return nil, errors.New("missing required argument 'EvaluatorType'")
	}
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.PolicySetName == nil {
		return nil, errors.New("missing required argument 'PolicySetName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Threshold == nil {
		return nil, errors.New("missing required argument 'Threshold'")
	}
	if args == nil {
		args = &PolicyArgs{}
	}
	var resource Policy
	err := ctx.RegisterResource("azure:devtest/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azure:devtest/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// A description for the Policy.
	Description *string `pulumi:"description"`
	// The Evaluation Type used for this Policy. Possible values include: 'AllowedValuesPolicy', 'MaxValuePolicy'. Changing this forces a new resource to be created.
	EvaluatorType *string `pulumi:"evaluatorType"`
	// The Fact Data for this Policy.
	FactData *string `pulumi:"factData"`
	// Specifies the name of the Dev Test Lab in which the Policy should be created. Changing this forces a new resource to be created.
	LabName *string `pulumi:"labName"`
	// Specifies the name of the Dev Test Policy. Possible values are `GalleryImage`, `LabPremiumVmCount`, `LabTargetCost`, `LabVmCount`, `LabVmSize`, `UserOwnedLabPremiumVmCount`, `UserOwnedLabVmCount` and `UserOwnedLabVmCountInSubnet`. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Policy Set within the Dev Test Lab where this policy should be created. Changing this forces a new resource to be created.
	PolicySetName *string `pulumi:"policySetName"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Threshold for this Policy.
	Threshold *string `pulumi:"threshold"`
}

type PolicyState struct {
	// A description for the Policy.
	Description pulumi.StringPtrInput
	// The Evaluation Type used for this Policy. Possible values include: 'AllowedValuesPolicy', 'MaxValuePolicy'. Changing this forces a new resource to be created.
	EvaluatorType pulumi.StringPtrInput
	// The Fact Data for this Policy.
	FactData pulumi.StringPtrInput
	// Specifies the name of the Dev Test Lab in which the Policy should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringPtrInput
	// Specifies the name of the Dev Test Policy. Possible values are `GalleryImage`, `LabPremiumVmCount`, `LabTargetCost`, `LabVmCount`, `LabVmSize`, `UserOwnedLabPremiumVmCount`, `UserOwnedLabVmCount` and `UserOwnedLabVmCountInSubnet`. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Policy Set within the Dev Test Lab where this policy should be created. Changing this forces a new resource to be created.
	PolicySetName pulumi.StringPtrInput
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Threshold for this Policy.
	Threshold pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// A description for the Policy.
	Description *string `pulumi:"description"`
	// The Evaluation Type used for this Policy. Possible values include: 'AllowedValuesPolicy', 'MaxValuePolicy'. Changing this forces a new resource to be created.
	EvaluatorType string `pulumi:"evaluatorType"`
	// The Fact Data for this Policy.
	FactData *string `pulumi:"factData"`
	// Specifies the name of the Dev Test Lab in which the Policy should be created. Changing this forces a new resource to be created.
	LabName string `pulumi:"labName"`
	// Specifies the name of the Dev Test Policy. Possible values are `GalleryImage`, `LabPremiumVmCount`, `LabTargetCost`, `LabVmCount`, `LabVmSize`, `UserOwnedLabPremiumVmCount`, `UserOwnedLabVmCount` and `UserOwnedLabVmCountInSubnet`. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Policy Set within the Dev Test Lab where this policy should be created. Changing this forces a new resource to be created.
	PolicySetName string `pulumi:"policySetName"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Threshold for this Policy.
	Threshold string `pulumi:"threshold"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// A description for the Policy.
	Description pulumi.StringPtrInput
	// The Evaluation Type used for this Policy. Possible values include: 'AllowedValuesPolicy', 'MaxValuePolicy'. Changing this forces a new resource to be created.
	EvaluatorType pulumi.StringInput
	// The Fact Data for this Policy.
	FactData pulumi.StringPtrInput
	// Specifies the name of the Dev Test Lab in which the Policy should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringInput
	// Specifies the name of the Dev Test Policy. Possible values are `GalleryImage`, `LabPremiumVmCount`, `LabTargetCost`, `LabVmCount`, `LabVmSize`, `UserOwnedLabPremiumVmCount`, `UserOwnedLabVmCount` and `UserOwnedLabVmCountInSubnet`. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Policy Set within the Dev Test Lab where this policy should be created. Changing this forces a new resource to be created.
	PolicySetName pulumi.StringInput
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Threshold for this Policy.
	Threshold pulumi.StringInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}
