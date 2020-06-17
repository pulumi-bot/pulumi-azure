// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managedapplication

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Managed Application Definition.
//
// ## Example Usage
type Definition struct {
	pulumi.CustomResourceState

	// One or more `authorization` block defined below.
	Authorizations DefinitionAuthorizationArrayOutput `pulumi:"authorizations"`
	// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
	CreateUiDefinition pulumi.StringPtrOutput `pulumi:"createUiDefinition"`
	// Specifies the managed application definition description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the managed application definition display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel pulumi.StringOutput `pulumi:"lockLevel"`
	// Specifies the inline main template json which has resources to be provisioned.
	MainTemplate pulumi.StringPtrOutput `pulumi:"mainTemplate"`
	// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Is the package enabled? Defaults to `true`.
	PackageEnabled pulumi.BoolPtrOutput `pulumi:"packageEnabled"`
	// Specifies the managed application definition package file Uri.
	PackageFileUri pulumi.StringPtrOutput `pulumi:"packageFileUri"`
	// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDefinition registers a new resource with the given unique name, arguments, and options.
func NewDefinition(ctx *pulumi.Context,
	name string, args *DefinitionArgs, opts ...pulumi.ResourceOption) (*Definition, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.LockLevel == nil {
		return nil, errors.New("missing required argument 'LockLevel'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DefinitionArgs{}
	}
	var resource Definition
	err := ctx.RegisterResource("azure:managedapplication/definition:Definition", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:managedapplication/definition:Definition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Definition resources.
type definitionState struct {
	// One or more `authorization` block defined below.
	Authorizations []DefinitionAuthorization `pulumi:"authorizations"`
	// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
	CreateUiDefinition *string `pulumi:"createUiDefinition"`
	// Specifies the managed application definition description.
	Description *string `pulumi:"description"`
	// Specifies the managed application definition display name.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel *string `pulumi:"lockLevel"`
	// Specifies the inline main template json which has resources to be provisioned.
	MainTemplate *string `pulumi:"mainTemplate"`
	// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Is the package enabled? Defaults to `true`.
	PackageEnabled *bool `pulumi:"packageEnabled"`
	// Specifies the managed application definition package file Uri.
	PackageFileUri *string `pulumi:"packageFileUri"`
	// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type DefinitionState struct {
	// One or more `authorization` block defined below.
	Authorizations DefinitionAuthorizationArrayInput
	// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
	CreateUiDefinition pulumi.StringPtrInput
	// Specifies the managed application definition description.
	Description pulumi.StringPtrInput
	// Specifies the managed application definition display name.
	DisplayName pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel pulumi.StringPtrInput
	// Specifies the inline main template json which has resources to be provisioned.
	MainTemplate pulumi.StringPtrInput
	// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Is the package enabled? Defaults to `true`.
	PackageEnabled pulumi.BoolPtrInput
	// Specifies the managed application definition package file Uri.
	PackageFileUri pulumi.StringPtrInput
	// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionState)(nil)).Elem()
}

type definitionArgs struct {
	// One or more `authorization` block defined below.
	Authorizations []DefinitionAuthorization `pulumi:"authorizations"`
	// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
	CreateUiDefinition *string `pulumi:"createUiDefinition"`
	// Specifies the managed application definition description.
	Description *string `pulumi:"description"`
	// Specifies the managed application definition display name.
	DisplayName string `pulumi:"displayName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel string `pulumi:"lockLevel"`
	// Specifies the inline main template json which has resources to be provisioned.
	MainTemplate *string `pulumi:"mainTemplate"`
	// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Is the package enabled? Defaults to `true`.
	PackageEnabled *bool `pulumi:"packageEnabled"`
	// Specifies the managed application definition package file Uri.
	PackageFileUri *string `pulumi:"packageFileUri"`
	// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Definition resource.
type DefinitionArgs struct {
	// One or more `authorization` block defined below.
	Authorizations DefinitionAuthorizationArrayInput
	// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
	CreateUiDefinition pulumi.StringPtrInput
	// Specifies the managed application definition description.
	Description pulumi.StringPtrInput
	// Specifies the managed application definition display name.
	DisplayName pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel pulumi.StringInput
	// Specifies the inline main template json which has resources to be provisioned.
	MainTemplate pulumi.StringPtrInput
	// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Is the package enabled? Defaults to `true`.
	PackageEnabled pulumi.BoolPtrInput
	// Specifies the managed application definition package file Uri.
	PackageFileUri pulumi.StringPtrInput
	// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionArgs)(nil)).Elem()
}
