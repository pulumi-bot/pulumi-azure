// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Automation Connection with type `AzureServicePrincipal`.
//
// ## Import
//
// Automation Connection can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:automation/connectionServicePrincipal:ConnectionServicePrincipal conn1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
// ```
type ConnectionServicePrincipal struct {
	pulumi.CustomResourceState

	// The (Client) ID of the Service Principal.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// The thumbprint of the Service Principal Certificate.
	CertificateThumbprint pulumi.StringOutput `pulumi:"certificateThumbprint"`
	// A description for this Connection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the name of the Connection. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The subscription GUID.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewConnectionServicePrincipal registers a new resource with the given unique name, arguments, and options.
func NewConnectionServicePrincipal(ctx *pulumi.Context,
	name string, args *ConnectionServicePrincipalArgs, opts ...pulumi.ResourceOption) (*ConnectionServicePrincipal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.CertificateThumbprint == nil {
		return nil, errors.New("invalid value for required argument 'CertificateThumbprint'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource ConnectionServicePrincipal
	err := ctx.RegisterResource("azure:automation/connectionServicePrincipal:ConnectionServicePrincipal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionServicePrincipal gets an existing ConnectionServicePrincipal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionServicePrincipal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionServicePrincipalState, opts ...pulumi.ResourceOption) (*ConnectionServicePrincipal, error) {
	var resource ConnectionServicePrincipal
	err := ctx.ReadResource("azure:automation/connectionServicePrincipal:ConnectionServicePrincipal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionServicePrincipal resources.
type connectionServicePrincipalState struct {
	// The (Client) ID of the Service Principal.
	ApplicationId *string `pulumi:"applicationId"`
	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// The thumbprint of the Service Principal Certificate.
	CertificateThumbprint *string `pulumi:"certificateThumbprint"`
	// A description for this Connection.
	Description *string `pulumi:"description"`
	// Specifies the name of the Connection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The subscription GUID.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId *string `pulumi:"tenantId"`
}

type ConnectionServicePrincipalState struct {
	// The (Client) ID of the Service Principal.
	ApplicationId pulumi.StringPtrInput
	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// The thumbprint of the Service Principal Certificate.
	CertificateThumbprint pulumi.StringPtrInput
	// A description for this Connection.
	Description pulumi.StringPtrInput
	// Specifies the name of the Connection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The subscription GUID.
	SubscriptionId pulumi.StringPtrInput
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId pulumi.StringPtrInput
}

func (ConnectionServicePrincipalState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionServicePrincipalState)(nil)).Elem()
}

type connectionServicePrincipalArgs struct {
	// The (Client) ID of the Service Principal.
	ApplicationId string `pulumi:"applicationId"`
	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The thumbprint of the Service Principal Certificate.
	CertificateThumbprint string `pulumi:"certificateThumbprint"`
	// A description for this Connection.
	Description *string `pulumi:"description"`
	// Specifies the name of the Connection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The subscription GUID.
	SubscriptionId string `pulumi:"subscriptionId"`
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ConnectionServicePrincipal resource.
type ConnectionServicePrincipalArgs struct {
	// The (Client) ID of the Service Principal.
	ApplicationId pulumi.StringInput
	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// The thumbprint of the Service Principal Certificate.
	CertificateThumbprint pulumi.StringInput
	// A description for this Connection.
	Description pulumi.StringPtrInput
	// Specifies the name of the Connection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The subscription GUID.
	SubscriptionId pulumi.StringInput
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId pulumi.StringInput
}

func (ConnectionServicePrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionServicePrincipalArgs)(nil)).Elem()
}

type ConnectionServicePrincipalInput interface {
	pulumi.Input

	ToConnectionServicePrincipalOutput() ConnectionServicePrincipalOutput
	ToConnectionServicePrincipalOutputWithContext(ctx context.Context) ConnectionServicePrincipalOutput
}

func (ConnectionServicePrincipal) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionServicePrincipal)(nil)).Elem()
}

func (i ConnectionServicePrincipal) ToConnectionServicePrincipalOutput() ConnectionServicePrincipalOutput {
	return i.ToConnectionServicePrincipalOutputWithContext(context.Background())
}

func (i ConnectionServicePrincipal) ToConnectionServicePrincipalOutputWithContext(ctx context.Context) ConnectionServicePrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionServicePrincipalOutput)
}

type ConnectionServicePrincipalOutput struct {
	*pulumi.OutputState
}

func (ConnectionServicePrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionServicePrincipalOutput)(nil)).Elem()
}

func (o ConnectionServicePrincipalOutput) ToConnectionServicePrincipalOutput() ConnectionServicePrincipalOutput {
	return o
}

func (o ConnectionServicePrincipalOutput) ToConnectionServicePrincipalOutputWithContext(ctx context.Context) ConnectionServicePrincipalOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectionServicePrincipalOutput{})
}
