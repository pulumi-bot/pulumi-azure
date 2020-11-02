// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blueprint

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Blueprint Assignment resource
//
// > **NOTE:** Azure Blueprints are in Preview and potentially subject to breaking change without notice.
//
// > **NOTE:** Azure Blueprint Assignments can only be applied to Subscriptions.  Assignments to Management Groups is not currently supported by the service or by this provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/blueprint"
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
// 		exampleSubscription, err := core.GetSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleDefinition, err := blueprint.GetDefinition(ctx, &blueprint.GetDefinitionArgs{
// 			Name:    "exampleBlueprint",
// 			ScopeId: exampleSubscription.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		examplePublishedVersion, err := blueprint.GetPublishedVersion(ctx, &blueprint.GetPublishedVersionArgs{
// 			ScopeId:       exampleDefinition.ScopeId,
// 			BlueprintName: exampleDefinition.Name,
// 			Version:       "v1.0.0",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("westeurope"),
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleUserAssignedIdentity, err := authorization.NewUserAssignedIdentity(ctx, "exampleUserAssignedIdentity", &authorization.UserAssignedIdentityArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		operator, err := authorization.NewAssignment(ctx, "operator", &authorization.AssignmentArgs{
// 			Scope:              pulumi.String(exampleSubscription.Id),
// 			RoleDefinitionName: pulumi.String("Blueprint Operator"),
// 			PrincipalId:        exampleUserAssignedIdentity.PrincipalId,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		owner, err := authorization.NewAssignment(ctx, "owner", &authorization.AssignmentArgs{
// 			Scope:              pulumi.String(exampleSubscription.Id),
// 			RoleDefinitionName: pulumi.String("Owner"),
// 			PrincipalId:        exampleUserAssignedIdentity.PrincipalId,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = blueprint.NewAssignment(ctx, "exampleAssignment", &blueprint.AssignmentArgs{
// 			TargetSubscriptionId: pulumi.String(exampleSubscription.Id),
// 			VersionId:            pulumi.String(examplePublishedVersion.Id),
// 			Location:             exampleResourceGroup.Location,
// 			LockMode:             pulumi.String("AllResourcesDoNotDelete"),
// 			LockExcludePrincipals: pulumi.StringArray{
// 				pulumi.String(current.ObjectId),
// 			},
// 			Identity: &blueprint.AssignmentIdentityArgs{
// 				Type: pulumi.String("UserAssigned"),
// 				IdentityIds: pulumi.StringArray{
// 					exampleUserAssignedIdentity.ID(),
// 				},
// 			},
// 			ResourceGroups:  pulumi.String(fmt.Sprintf("%v%v%v%v%v", "    {\n", "      \"ResourceGroup\": {\n", "        \"name\": \"exampleRG-bp\"\n", "      }\n", "    }\n")),
// 			ParameterValues: pulumi.String(fmt.Sprintf("%v%v%v%v%v", "    {\n", "      \"allowedlocationsforresourcegroups_listOfAllowedLocations\": {\n", "        \"value\": [\"westus\", \"westus2\", \"eastus\", \"centralus\", \"centraluseuap\", \"southcentralus\", \"northcentralus\", \"westcentralus\", \"eastus2\", \"eastus2euap\", \"brazilsouth\", \"brazilus\", \"northeurope\", \"westeurope\", \"eastasia\", \"southeastasia\", \"japanwest\", \"japaneast\", \"koreacentral\", \"koreasouth\", \"indiasouth\", \"indiawest\", \"indiacentral\", \"australiaeast\", \"australiasoutheast\", \"canadacentral\", \"canadaeast\", \"uknorth\", \"uksouth2\", \"uksouth\", \"ukwest\", \"francecentral\", \"francesouth\", \"australiacentral\", \"australiacentral2\", \"uaecentral\", \"uaenorth\", \"southafricanorth\", \"southafricawest\", \"switzerlandnorth\", \"switzerlandwest\", \"germanynorth\", \"germanywestcentral\", \"norwayeast\", \"norwaywest\"]\n", "      }\n", "    }\n")),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			operator,
// 			owner,
// 		}))
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
// Azure Blueprint Assignments can be imported using the `resource id`, e.g.
type Assignment struct {
	pulumi.CustomResourceState

	// The name of the blueprint assigned
	BlueprintName pulumi.StringOutput `pulumi:"blueprintName"`
	// The Description on the Blueprint
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name of the blueprint
	DisplayName pulumi.StringOutput         `pulumi:"displayName"`
	Identity    AssignmentIdentityPtrOutput `pulumi:"identity"`
	// The Azure location of the Assignment.
	Location pulumi.StringOutput `pulumi:"location"`
	// a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
	LockExcludePrincipals pulumi.StringArrayOutput `pulumi:"lockExcludePrincipals"`
	// The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
	LockMode pulumi.StringPtrOutput `pulumi:"lockMode"`
	// The name of the Blueprint Assignment
	Name pulumi.StringOutput `pulumi:"name"`
	// a JSON string to supply Blueprint Assignment parameter values.
	ParameterValues pulumi.StringPtrOutput `pulumi:"parameterValues"`
	// a JSON string to supply the Blueprint Resource Group information.
	ResourceGroups pulumi.StringPtrOutput `pulumi:"resourceGroups"`
	// The Subscription ID the Blueprint Published Version is to be applied to.
	TargetSubscriptionId pulumi.StringOutput `pulumi:"targetSubscriptionId"`
	// The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
	Type pulumi.StringOutput `pulumi:"type"`
	// The ID of the Published Version of the blueprint to be assigned.
	VersionId pulumi.StringOutput `pulumi:"versionId"`
}

// NewAssignment registers a new resource with the given unique name, arguments, and options.
func NewAssignment(ctx *pulumi.Context,
	name string, args *AssignmentArgs, opts ...pulumi.ResourceOption) (*Assignment, error) {
	if args == nil || args.TargetSubscriptionId == nil {
		return nil, errors.New("missing required argument 'TargetSubscriptionId'")
	}
	if args == nil || args.VersionId == nil {
		return nil, errors.New("missing required argument 'VersionId'")
	}
	if args == nil {
		args = &AssignmentArgs{}
	}
	var resource Assignment
	err := ctx.RegisterResource("azure:blueprint/assignment:Assignment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:blueprint/assignment:Assignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assignment resources.
type assignmentState struct {
	// The name of the blueprint assigned
	BlueprintName *string `pulumi:"blueprintName"`
	// The Description on the Blueprint
	Description *string `pulumi:"description"`
	// The display name of the blueprint
	DisplayName *string             `pulumi:"displayName"`
	Identity    *AssignmentIdentity `pulumi:"identity"`
	// The Azure location of the Assignment.
	Location *string `pulumi:"location"`
	// a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
	LockExcludePrincipals []string `pulumi:"lockExcludePrincipals"`
	// The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
	LockMode *string `pulumi:"lockMode"`
	// The name of the Blueprint Assignment
	Name *string `pulumi:"name"`
	// a JSON string to supply Blueprint Assignment parameter values.
	ParameterValues *string `pulumi:"parameterValues"`
	// a JSON string to supply the Blueprint Resource Group information.
	ResourceGroups *string `pulumi:"resourceGroups"`
	// The Subscription ID the Blueprint Published Version is to be applied to.
	TargetSubscriptionId *string `pulumi:"targetSubscriptionId"`
	// The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
	Type *string `pulumi:"type"`
	// The ID of the Published Version of the blueprint to be assigned.
	VersionId *string `pulumi:"versionId"`
}

type AssignmentState struct {
	// The name of the blueprint assigned
	BlueprintName pulumi.StringPtrInput
	// The Description on the Blueprint
	Description pulumi.StringPtrInput
	// The display name of the blueprint
	DisplayName pulumi.StringPtrInput
	Identity    AssignmentIdentityPtrInput
	// The Azure location of the Assignment.
	Location pulumi.StringPtrInput
	// a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
	LockExcludePrincipals pulumi.StringArrayInput
	// The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
	LockMode pulumi.StringPtrInput
	// The name of the Blueprint Assignment
	Name pulumi.StringPtrInput
	// a JSON string to supply Blueprint Assignment parameter values.
	ParameterValues pulumi.StringPtrInput
	// a JSON string to supply the Blueprint Resource Group information.
	ResourceGroups pulumi.StringPtrInput
	// The Subscription ID the Blueprint Published Version is to be applied to.
	TargetSubscriptionId pulumi.StringPtrInput
	// The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
	Type pulumi.StringPtrInput
	// The ID of the Published Version of the blueprint to be assigned.
	VersionId pulumi.StringPtrInput
}

func (AssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentState)(nil)).Elem()
}

type assignmentArgs struct {
	Identity *AssignmentIdentity `pulumi:"identity"`
	// The Azure location of the Assignment.
	Location *string `pulumi:"location"`
	// a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
	LockExcludePrincipals []string `pulumi:"lockExcludePrincipals"`
	// The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
	LockMode *string `pulumi:"lockMode"`
	// The name of the Blueprint Assignment
	Name *string `pulumi:"name"`
	// a JSON string to supply Blueprint Assignment parameter values.
	ParameterValues *string `pulumi:"parameterValues"`
	// a JSON string to supply the Blueprint Resource Group information.
	ResourceGroups *string `pulumi:"resourceGroups"`
	// The Subscription ID the Blueprint Published Version is to be applied to.
	TargetSubscriptionId string `pulumi:"targetSubscriptionId"`
	// The ID of the Published Version of the blueprint to be assigned.
	VersionId string `pulumi:"versionId"`
}

// The set of arguments for constructing a Assignment resource.
type AssignmentArgs struct {
	Identity AssignmentIdentityPtrInput
	// The Azure location of the Assignment.
	Location pulumi.StringPtrInput
	// a list of up to 5 Principal IDs that are permitted to bypass the locks applied by the Blueprint.
	LockExcludePrincipals pulumi.StringArrayInput
	// The locking mode of the Blueprint Assignment.  One of `None` (Default), `AllResourcesReadOnly`, or `AlResourcesDoNotDelete`.
	LockMode pulumi.StringPtrInput
	// The name of the Blueprint Assignment
	Name pulumi.StringPtrInput
	// a JSON string to supply Blueprint Assignment parameter values.
	ParameterValues pulumi.StringPtrInput
	// a JSON string to supply the Blueprint Resource Group information.
	ResourceGroups pulumi.StringPtrInput
	// The Subscription ID the Blueprint Published Version is to be applied to.
	TargetSubscriptionId pulumi.StringInput
	// The ID of the Published Version of the blueprint to be assigned.
	VersionId pulumi.StringInput
}

func (AssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentArgs)(nil)).Elem()
}
