// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows you to set a user or group as the AD administrator for an MySQL server in Azure
type ActiveDirectoryAdministrator struct {
	pulumi.CustomResourceState

	// The login name of the principal to set as the server administrator
	Login pulumi.StringOutput `pulumi:"login"`
	// The ID of the principal to set as the server administrator
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// The Azure Tenant ID
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewActiveDirectoryAdministrator registers a new resource with the given unique name, arguments, and options.
func NewActiveDirectoryAdministrator(ctx *pulumi.Context,
	name string, args *ActiveDirectoryAdministratorArgs, opts ...pulumi.ResourceOption) (*ActiveDirectoryAdministrator, error) {
	if args == nil || args.Login == nil {
		return nil, errors.New("missing required argument 'Login'")
	}
	if args == nil || args.ObjectId == nil {
		return nil, errors.New("missing required argument 'ObjectId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.TenantId == nil {
		return nil, errors.New("missing required argument 'TenantId'")
	}
	if args == nil {
		args = &ActiveDirectoryAdministratorArgs{}
	}
	var resource ActiveDirectoryAdministrator
	err := ctx.RegisterResource("azure:mysql/activeDirectoryAdministrator:ActiveDirectoryAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActiveDirectoryAdministrator gets an existing ActiveDirectoryAdministrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActiveDirectoryAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActiveDirectoryAdministratorState, opts ...pulumi.ResourceOption) (*ActiveDirectoryAdministrator, error) {
	var resource ActiveDirectoryAdministrator
	err := ctx.ReadResource("azure:mysql/activeDirectoryAdministrator:ActiveDirectoryAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActiveDirectoryAdministrator resources.
type activeDirectoryAdministratorState struct {
	// The login name of the principal to set as the server administrator
	Login *string `pulumi:"login"`
	// The ID of the principal to set as the server administrator
	ObjectId *string `pulumi:"objectId"`
	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// The Azure Tenant ID
	TenantId *string `pulumi:"tenantId"`
}

type ActiveDirectoryAdministratorState struct {
	// The login name of the principal to set as the server administrator
	Login pulumi.StringPtrInput
	// The ID of the principal to set as the server administrator
	ObjectId pulumi.StringPtrInput
	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// The Azure Tenant ID
	TenantId pulumi.StringPtrInput
}

func (ActiveDirectoryAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*activeDirectoryAdministratorState)(nil)).Elem()
}

type activeDirectoryAdministratorArgs struct {
	// The login name of the principal to set as the server administrator
	Login string `pulumi:"login"`
	// The ID of the principal to set as the server administrator
	ObjectId string `pulumi:"objectId"`
	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// The Azure Tenant ID
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ActiveDirectoryAdministrator resource.
type ActiveDirectoryAdministratorArgs struct {
	// The login name of the principal to set as the server administrator
	Login pulumi.StringInput
	// The ID of the principal to set as the server administrator
	ObjectId pulumi.StringInput
	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// The Azure Tenant ID
	TenantId pulumi.StringInput
}

func (ActiveDirectoryAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activeDirectoryAdministratorArgs)(nil)).Elem()
}
