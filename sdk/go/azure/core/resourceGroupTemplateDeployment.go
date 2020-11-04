// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Resource Group Template Deployment.
//
// > **Note:** This resource will automatically attempt to delete resources deployed by the ARM Template when it is deleted. You can opt-out of this by setting the `deleteNestedItemsDuringDeletion` field within the `templateDeployment` block of the `features` block to `false`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.NewResourceGroupTemplateDeployment(ctx, "example", &core.ResourceGroupTemplateDeploymentArgs{
// 			DeploymentMode:    pulumi.String("Complete"),
// 			ResourceGroupName: pulumi.String("example-group"),
// 			TemplateContent:   pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"", "$", "schema\": \"https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#\",\n", "  \"contentVersion\": \"1.0.0.0\",\n", "  \"parameters\": {},\n", "  \"variables\": {},\n", "  \"resources\": [\n", "    {\n", "      \"type\": \"Microsoft.Network/virtualNetworks\",\n", "      \"apiVersion\": \"2020-05-01\",\n", "      \"name\": \"acctest-network\",\n", "      \"location\": \"[resourceGroup().location]\",\n", "      \"properties\": {\n", "        \"addressSpace\": {\n", "          \"addressPrefixes\": [\n", "            \"10.0.0.0/16\"\n", "          ]\n", "        }\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ResourceGroupTemplateDeployment struct {
	pulumi.CustomResourceState

	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
	DebugLevel pulumi.StringPtrOutput `pulumi:"debugLevel"`
	// The Deployment Mode for this Resource Group Template Deployment. Possible values are `Complete` (where resources in the Resource Group not specified in the ARM Template will be destroyed) and `Incremental` (where resources are additive only).
	DeploymentMode pulumi.StringOutput `pulumi:"deploymentMode"`
	// The name which should be used for this Resource Group Template Deployment. Changing this forces a new Resource Group Template Deployment to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The JSON Content of the Outputs of the ARM Template Deployment.
	OutputContent pulumi.StringOutput `pulumi:"outputContent"`
	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent pulumi.StringOutput `pulumi:"parametersContent"`
	// The name of the Resource Group where the Resource Group Template Deployment should exist. Changing this forces a new Resource Group Template Deployment to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Resource Group Template Deployment.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The contents of the ARM Template which should be deployed into this Resource Group.
	TemplateContent pulumi.StringOutput `pulumi:"templateContent"`
}

// NewResourceGroupTemplateDeployment registers a new resource with the given unique name, arguments, and options.
func NewResourceGroupTemplateDeployment(ctx *pulumi.Context,
	name string, args *ResourceGroupTemplateDeploymentArgs, opts ...pulumi.ResourceOption) (*ResourceGroupTemplateDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DeploymentMode == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentMode'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TemplateContent == nil {
		return nil, errors.New("invalid value for required argument 'TemplateContent'")
	}
	var resource ResourceGroupTemplateDeployment
	err := ctx.RegisterResource("azure:core/resourceGroupTemplateDeployment:ResourceGroupTemplateDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGroupTemplateDeployment gets an existing ResourceGroupTemplateDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroupTemplateDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupTemplateDeploymentState, opts ...pulumi.ResourceOption) (*ResourceGroupTemplateDeployment, error) {
	var resource ResourceGroupTemplateDeployment
	err := ctx.ReadResource("azure:core/resourceGroupTemplateDeployment:ResourceGroupTemplateDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGroupTemplateDeployment resources.
type resourceGroupTemplateDeploymentState struct {
	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
	DebugLevel *string `pulumi:"debugLevel"`
	// The Deployment Mode for this Resource Group Template Deployment. Possible values are `Complete` (where resources in the Resource Group not specified in the ARM Template will be destroyed) and `Incremental` (where resources are additive only).
	DeploymentMode *string `pulumi:"deploymentMode"`
	// The name which should be used for this Resource Group Template Deployment. Changing this forces a new Resource Group Template Deployment to be created.
	Name *string `pulumi:"name"`
	// The JSON Content of the Outputs of the ARM Template Deployment.
	OutputContent *string `pulumi:"outputContent"`
	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent *string `pulumi:"parametersContent"`
	// The name of the Resource Group where the Resource Group Template Deployment should exist. Changing this forces a new Resource Group Template Deployment to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Resource Group Template Deployment.
	Tags map[string]string `pulumi:"tags"`
	// The contents of the ARM Template which should be deployed into this Resource Group.
	TemplateContent *string `pulumi:"templateContent"`
}

type ResourceGroupTemplateDeploymentState struct {
	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
	DebugLevel pulumi.StringPtrInput
	// The Deployment Mode for this Resource Group Template Deployment. Possible values are `Complete` (where resources in the Resource Group not specified in the ARM Template will be destroyed) and `Incremental` (where resources are additive only).
	DeploymentMode pulumi.StringPtrInput
	// The name which should be used for this Resource Group Template Deployment. Changing this forces a new Resource Group Template Deployment to be created.
	Name pulumi.StringPtrInput
	// The JSON Content of the Outputs of the ARM Template Deployment.
	OutputContent pulumi.StringPtrInput
	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent pulumi.StringPtrInput
	// The name of the Resource Group where the Resource Group Template Deployment should exist. Changing this forces a new Resource Group Template Deployment to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Resource Group Template Deployment.
	Tags pulumi.StringMapInput
	// The contents of the ARM Template which should be deployed into this Resource Group.
	TemplateContent pulumi.StringPtrInput
}

func (ResourceGroupTemplateDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupTemplateDeploymentState)(nil)).Elem()
}

type resourceGroupTemplateDeploymentArgs struct {
	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
	DebugLevel *string `pulumi:"debugLevel"`
	// The Deployment Mode for this Resource Group Template Deployment. Possible values are `Complete` (where resources in the Resource Group not specified in the ARM Template will be destroyed) and `Incremental` (where resources are additive only).
	DeploymentMode string `pulumi:"deploymentMode"`
	// The name which should be used for this Resource Group Template Deployment. Changing this forces a new Resource Group Template Deployment to be created.
	Name *string `pulumi:"name"`
	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent *string `pulumi:"parametersContent"`
	// The name of the Resource Group where the Resource Group Template Deployment should exist. Changing this forces a new Resource Group Template Deployment to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Resource Group Template Deployment.
	Tags map[string]string `pulumi:"tags"`
	// The contents of the ARM Template which should be deployed into this Resource Group.
	TemplateContent string `pulumi:"templateContent"`
}

// The set of arguments for constructing a ResourceGroupTemplateDeployment resource.
type ResourceGroupTemplateDeploymentArgs struct {
	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
	DebugLevel pulumi.StringPtrInput
	// The Deployment Mode for this Resource Group Template Deployment. Possible values are `Complete` (where resources in the Resource Group not specified in the ARM Template will be destroyed) and `Incremental` (where resources are additive only).
	DeploymentMode pulumi.StringInput
	// The name which should be used for this Resource Group Template Deployment. Changing this forces a new Resource Group Template Deployment to be created.
	Name pulumi.StringPtrInput
	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent pulumi.StringPtrInput
	// The name of the Resource Group where the Resource Group Template Deployment should exist. Changing this forces a new Resource Group Template Deployment to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the Resource Group Template Deployment.
	Tags pulumi.StringMapInput
	// The contents of the ARM Template which should be deployed into this Resource Group.
	TemplateContent pulumi.StringInput
}

func (ResourceGroupTemplateDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupTemplateDeploymentArgs)(nil)).Elem()
}
