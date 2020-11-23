// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Virtual Desktop Workspace Application Group Association.
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
// 		pooledbreadthfirst, err := desktopvirtualization.NewHostPool(ctx, "pooledbreadthfirst", &desktopvirtualization.HostPoolArgs{
// 			Location:          example.Location,
// 			ResourceGroupName: example.Name,
// 			Type:              pulumi.String("Pooled"),
// 			LoadBalancerType:  pulumi.String("BreadthFirst"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		remoteapp, err := desktopvirtualization.NewApplicationGroup(ctx, "remoteapp", &desktopvirtualization.ApplicationGroupArgs{
// 			Location:          example.Location,
// 			ResourceGroupName: example.Name,
// 			Type:              pulumi.String("RemoteApp"),
// 			HostPoolId:        pooledbreadthfirst.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		workspace, err := desktopvirtualization.NewWorkspace(ctx, "workspace", &desktopvirtualization.WorkspaceArgs{
// 			Location:          example.Location,
// 			ResourceGroupName: example.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = desktopvirtualization.NewWorkspaceApplicationGroupAssociation(ctx, "workspaceremoteapp", &desktopvirtualization.WorkspaceApplicationGroupAssociationArgs{
// 			WorkspaceId:        workspace.ID(),
// 			ApplicationGroupId: remoteapp.ID(),
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
// Associations between Virtual Desktop Workspaces and Virtual Desktop Application Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:desktopvirtualization/workspaceApplicationGroupAssociation:WorkspaceApplicationGroupAssociation association1 "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myGroup1/providers/Microsoft.DesktopVirtualization/workspaces/myworkspace|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/myapplicationgroup"
// ```
type WorkspaceApplicationGroupAssociation struct {
	pulumi.CustomResourceState

	// The resource ID for the Virtual Desktop Application Group.
	ApplicationGroupId pulumi.StringOutput `pulumi:"applicationGroupId"`
	// The resource ID for the Virtual Desktop Workspace.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewWorkspaceApplicationGroupAssociation registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceApplicationGroupAssociation(ctx *pulumi.Context,
	name string, args *WorkspaceApplicationGroupAssociationArgs, opts ...pulumi.ResourceOption) (*WorkspaceApplicationGroupAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationGroupId'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	var resource WorkspaceApplicationGroupAssociation
	err := ctx.RegisterResource("azure:desktopvirtualization/workspaceApplicationGroupAssociation:WorkspaceApplicationGroupAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceApplicationGroupAssociation gets an existing WorkspaceApplicationGroupAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceApplicationGroupAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceApplicationGroupAssociationState, opts ...pulumi.ResourceOption) (*WorkspaceApplicationGroupAssociation, error) {
	var resource WorkspaceApplicationGroupAssociation
	err := ctx.ReadResource("azure:desktopvirtualization/workspaceApplicationGroupAssociation:WorkspaceApplicationGroupAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceApplicationGroupAssociation resources.
type workspaceApplicationGroupAssociationState struct {
	// The resource ID for the Virtual Desktop Application Group.
	ApplicationGroupId *string `pulumi:"applicationGroupId"`
	// The resource ID for the Virtual Desktop Workspace.
	WorkspaceId *string `pulumi:"workspaceId"`
}

type WorkspaceApplicationGroupAssociationState struct {
	// The resource ID for the Virtual Desktop Application Group.
	ApplicationGroupId pulumi.StringPtrInput
	// The resource ID for the Virtual Desktop Workspace.
	WorkspaceId pulumi.StringPtrInput
}

func (WorkspaceApplicationGroupAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceApplicationGroupAssociationState)(nil)).Elem()
}

type workspaceApplicationGroupAssociationArgs struct {
	// The resource ID for the Virtual Desktop Application Group.
	ApplicationGroupId string `pulumi:"applicationGroupId"`
	// The resource ID for the Virtual Desktop Workspace.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceApplicationGroupAssociation resource.
type WorkspaceApplicationGroupAssociationArgs struct {
	// The resource ID for the Virtual Desktop Application Group.
	ApplicationGroupId pulumi.StringInput
	// The resource ID for the Virtual Desktop Workspace.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceApplicationGroupAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceApplicationGroupAssociationArgs)(nil)).Elem()
}

type WorkspaceApplicationGroupAssociationInput interface {
	pulumi.Input

	ToWorkspaceApplicationGroupAssociationOutput() WorkspaceApplicationGroupAssociationOutput
	ToWorkspaceApplicationGroupAssociationOutputWithContext(ctx context.Context) WorkspaceApplicationGroupAssociationOutput
}

func (WorkspaceApplicationGroupAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceApplicationGroupAssociation)(nil)).Elem()
}

func (i WorkspaceApplicationGroupAssociation) ToWorkspaceApplicationGroupAssociationOutput() WorkspaceApplicationGroupAssociationOutput {
	return i.ToWorkspaceApplicationGroupAssociationOutputWithContext(context.Background())
}

func (i WorkspaceApplicationGroupAssociation) ToWorkspaceApplicationGroupAssociationOutputWithContext(ctx context.Context) WorkspaceApplicationGroupAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceApplicationGroupAssociationOutput)
}

type WorkspaceApplicationGroupAssociationOutput struct {
	*pulumi.OutputState
}

func (WorkspaceApplicationGroupAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceApplicationGroupAssociationOutput)(nil)).Elem()
}

func (o WorkspaceApplicationGroupAssociationOutput) ToWorkspaceApplicationGroupAssociationOutput() WorkspaceApplicationGroupAssociationOutput {
	return o
}

func (o WorkspaceApplicationGroupAssociationOutput) ToWorkspaceApplicationGroupAssociationOutputWithContext(ctx context.Context) WorkspaceApplicationGroupAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkspaceApplicationGroupAssociationOutput{})
}
