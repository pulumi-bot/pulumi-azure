// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Virtual Desktop Workspace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/desktopvirtualization"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("eastus"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = desktopvirtualization.NewWorkspace(ctx, "workspace", &desktopvirtualization.WorkspaceArgs{
// 			Location:          example.Location,
// 			ResourceGroupName: example.Name,
// 			FriendlyName:      pulumi.String("FriendlyName"),
// 			Description:       pulumi.String("A description of my workspace"),
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
// Virtual Desktop Workspaces can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:desktopvirtualization/workspace:Workspace example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myGroup1/providers/Microsoft.DesktopVirtualization/workspaces/myworkspace
// ```
type Workspace struct {
	pulumi.CustomResourceState

	// A description for the Virtual Desktop Workspace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A friendly name for the Virtual Desktop Workspace.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// The location/region where the Virtual Desktop Workspace is located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Virtual Desktop Workspace. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to
	// create the Virtual Desktop Workspace. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Workspace
	err := ctx.RegisterResource("azure:desktopvirtualization/workspace:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure:desktopvirtualization/workspace:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// A description for the Virtual Desktop Workspace.
	Description *string `pulumi:"description"`
	// A friendly name for the Virtual Desktop Workspace.
	FriendlyName *string `pulumi:"friendlyName"`
	// The location/region where the Virtual Desktop Workspace is located. Changing the location/region forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Virtual Desktop Workspace. Changing the name
	// forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the Virtual Desktop Workspace. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type WorkspaceState struct {
	// A description for the Virtual Desktop Workspace.
	Description pulumi.StringPtrInput
	// A friendly name for the Virtual Desktop Workspace.
	FriendlyName pulumi.StringPtrInput
	// The location/region where the Virtual Desktop Workspace is located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Virtual Desktop Workspace. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the Virtual Desktop Workspace. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// A description for the Virtual Desktop Workspace.
	Description *string `pulumi:"description"`
	// A friendly name for the Virtual Desktop Workspace.
	FriendlyName *string `pulumi:"friendlyName"`
	// The location/region where the Virtual Desktop Workspace is located. Changing the location/region forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Virtual Desktop Workspace. Changing the name
	// forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the Virtual Desktop Workspace. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// A description for the Virtual Desktop Workspace.
	Description pulumi.StringPtrInput
	// A friendly name for the Virtual Desktop Workspace.
	FriendlyName pulumi.StringPtrInput
	// The location/region where the Virtual Desktop Workspace is located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Virtual Desktop Workspace. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the Virtual Desktop Workspace. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

type WorkspacePtrInput interface {
	pulumi.Input

	ToWorkspacePtrOutput() WorkspacePtrOutput
	ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput
}

func (Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((*Workspace)(nil)).Elem()
}

func (i Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

func (i Workspace) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return i.ToWorkspacePtrOutputWithContext(context.Background())
}

func (i Workspace) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePtrOutput)
}

type WorkspaceOutput struct {
	*pulumi.OutputState
}

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceOutput)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

type WorkspacePtrOutput struct {
	*pulumi.OutputState
}

func (WorkspacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (o WorkspacePtrOutput) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return o
}

func (o WorkspacePtrOutput) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
	pulumi.RegisterOutputType(WorkspacePtrOutput{})
}
