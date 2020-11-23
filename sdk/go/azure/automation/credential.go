// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Automation Credential.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/automation"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := automation.NewAccount(ctx, "exampleAccount", &automation.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("Basic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = automation.NewCredential(ctx, "exampleCredential", &automation.CredentialArgs{
// 			ResourceGroupName:     exampleResourceGroup.Name,
// 			AutomationAccountName: exampleAccount.Name,
// 			Username:              pulumi.String("example_user"),
// 			Password:              pulumi.String("example_pwd"),
// 			Description:           pulumi.String("This is an example credential"),
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
// Automation Credentials can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:automation/credential:Credential credential1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/credentials/credential1
// ```
type Credential struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the Credential is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// The description associated with this Automation Credential.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the name of the Credential. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password associated with this Automation Credential.
	Password pulumi.StringOutput `pulumi:"password"`
	// The name of the resource group in which the Credential is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The username associated with this Automation Credential.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewCredential registers a new resource with the given unique name, arguments, and options.
func NewCredential(ctx *pulumi.Context,
	name string, args *CredentialArgs, opts ...pulumi.ResourceOption) (*Credential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource Credential
	err := ctx.RegisterResource("azure:automation/credential:Credential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCredential gets an existing Credential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CredentialState, opts ...pulumi.ResourceOption) (*Credential, error) {
	var resource Credential
	err := ctx.ReadResource("azure:automation/credential:Credential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Credential resources.
type credentialState struct {
	// The name of the automation account in which the Credential is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// The description associated with this Automation Credential.
	Description *string `pulumi:"description"`
	// Specifies the name of the Credential. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The password associated with this Automation Credential.
	Password *string `pulumi:"password"`
	// The name of the resource group in which the Credential is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The username associated with this Automation Credential.
	Username *string `pulumi:"username"`
}

type CredentialState struct {
	// The name of the automation account in which the Credential is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// The description associated with this Automation Credential.
	Description pulumi.StringPtrInput
	// Specifies the name of the Credential. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The password associated with this Automation Credential.
	Password pulumi.StringPtrInput
	// The name of the resource group in which the Credential is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The username associated with this Automation Credential.
	Username pulumi.StringPtrInput
}

func (CredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialState)(nil)).Elem()
}

type credentialArgs struct {
	// The name of the automation account in which the Credential is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The description associated with this Automation Credential.
	Description *string `pulumi:"description"`
	// Specifies the name of the Credential. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The password associated with this Automation Credential.
	Password string `pulumi:"password"`
	// The name of the resource group in which the Credential is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The username associated with this Automation Credential.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Credential resource.
type CredentialArgs struct {
	// The name of the automation account in which the Credential is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// The description associated with this Automation Credential.
	Description pulumi.StringPtrInput
	// Specifies the name of the Credential. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The password associated with this Automation Credential.
	Password pulumi.StringInput
	// The name of the resource group in which the Credential is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The username associated with this Automation Credential.
	Username pulumi.StringInput
}

func (CredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialArgs)(nil)).Elem()
}

type CredentialInput interface {
	pulumi.Input

	ToCredentialOutput() CredentialOutput
	ToCredentialOutputWithContext(ctx context.Context) CredentialOutput
}

func (Credential) ElementType() reflect.Type {
	return reflect.TypeOf((*Credential)(nil)).Elem()
}

func (i Credential) ToCredentialOutput() CredentialOutput {
	return i.ToCredentialOutputWithContext(context.Background())
}

func (i Credential) ToCredentialOutputWithContext(ctx context.Context) CredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialOutput)
}

type CredentialOutput struct {
	*pulumi.OutputState
}

func (CredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialOutput)(nil)).Elem()
}

func (o CredentialOutput) ToCredentialOutput() CredentialOutput {
	return o
}

func (o CredentialOutput) ToCredentialOutputWithContext(ctx context.Context) CredentialOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CredentialOutput{})
}
