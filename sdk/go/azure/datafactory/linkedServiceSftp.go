// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linked Service (connection) between a SFTP Server and Azure Data Factory.
//
// > **Note:** All arguments including the client secret will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datafactory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedServiceSftp(ctx, "exampleLinkedServiceSftp", &datafactory.LinkedServiceSftpArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			DataFactoryName:    exampleFactory.Name,
// 			AuthenticationType: pulumi.String("Basic"),
// 			Host:               pulumi.String("http://www.bing.com"),
// 			Port:               pulumi.Int(22),
// 			Username:           pulumi.String("foo"),
// 			Password:           pulumi.String("bar"),
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
// Data Factory Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceSftp:LinkedServiceSftp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceSftp struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType pulumi.StringOutput `pulumi:"authenticationType"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The SFTP server hostname.
	Host pulumi.StringOutput `pulumi:"host"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Password to logon to the SFTP Server for Basic Authentication.
	Password pulumi.StringOutput `pulumi:"password"`
	// The TCP port number that the SFTP server uses to lsiten for client connection. Default value is 22.
	Port pulumi.IntOutput `pulumi:"port"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The username used to log on to the SFTP server.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewLinkedServiceSftp registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceSftp(ctx *pulumi.Context,
	name string, args *LinkedServiceSftpArgs, opts ...pulumi.ResourceOption) (*LinkedServiceSftp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthenticationType == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationType'")
	}
	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource LinkedServiceSftp
	err := ctx.RegisterResource("azure:datafactory/linkedServiceSftp:LinkedServiceSftp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceSftp gets an existing LinkedServiceSftp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceSftp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceSftpState, opts ...pulumi.ResourceOption) (*LinkedServiceSftp, error) {
	var resource LinkedServiceSftp
	err := ctx.ReadResource("azure:datafactory/linkedServiceSftp:LinkedServiceSftp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceSftp resources.
type linkedServiceSftpState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType *string `pulumi:"authenticationType"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The SFTP server hostname.
	Host *string `pulumi:"host"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// Password to logon to the SFTP Server for Basic Authentication.
	Password *string `pulumi:"password"`
	// The TCP port number that the SFTP server uses to lsiten for client connection. Default value is 22.
	Port *int `pulumi:"port"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The username used to log on to the SFTP server.
	Username *string `pulumi:"username"`
}

type LinkedServiceSftpState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The SFTP server hostname.
	Host pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// Password to logon to the SFTP Server for Basic Authentication.
	Password pulumi.StringPtrInput
	// The TCP port number that the SFTP server uses to lsiten for client connection. Default value is 22.
	Port pulumi.IntPtrInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// The username used to log on to the SFTP server.
	Username pulumi.StringPtrInput
}

func (LinkedServiceSftpState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSftpState)(nil)).Elem()
}

type linkedServiceSftpArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType string `pulumi:"authenticationType"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The SFTP server hostname.
	Host string `pulumi:"host"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// Password to logon to the SFTP Server for Basic Authentication.
	Password string `pulumi:"password"`
	// The TCP port number that the SFTP server uses to lsiten for client connection. Default value is 22.
	Port int `pulumi:"port"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The username used to log on to the SFTP server.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a LinkedServiceSftp resource.
type LinkedServiceSftpArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The SFTP server hostname.
	Host pulumi.StringInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// Password to logon to the SFTP Server for Basic Authentication.
	Password pulumi.StringInput
	// The TCP port number that the SFTP server uses to lsiten for client connection. Default value is 22.
	Port pulumi.IntInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// The username used to log on to the SFTP server.
	Username pulumi.StringInput
}

func (LinkedServiceSftpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSftpArgs)(nil)).Elem()
}

type LinkedServiceSftpInput interface {
	pulumi.Input

	ToLinkedServiceSftpOutput() LinkedServiceSftpOutput
	ToLinkedServiceSftpOutputWithContext(ctx context.Context) LinkedServiceSftpOutput
}

func (*LinkedServiceSftp) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceSftp)(nil))
}

func (i *LinkedServiceSftp) ToLinkedServiceSftpOutput() LinkedServiceSftpOutput {
	return i.ToLinkedServiceSftpOutputWithContext(context.Background())
}

func (i *LinkedServiceSftp) ToLinkedServiceSftpOutputWithContext(ctx context.Context) LinkedServiceSftpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSftpOutput)
}

type LinkedServiceSftpOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceSftpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceSftp)(nil))
}

func (o LinkedServiceSftpOutput) ToLinkedServiceSftpOutput() LinkedServiceSftpOutput {
	return o
}

func (o LinkedServiceSftpOutput) ToLinkedServiceSftpOutputWithContext(ctx context.Context) LinkedServiceSftpOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceSftpOutput{})
}
