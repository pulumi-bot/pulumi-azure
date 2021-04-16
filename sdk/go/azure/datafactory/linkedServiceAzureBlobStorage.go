// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Linked Service (connection) between an Azure Blob Storage Account and Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datafactory"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedServiceAzureBlobStorage(ctx, "exampleLinkedServiceAzureBlobStorage", &datafactory.LinkedServiceAzureBlobStorageArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			ConnectionString: exampleAccount.ApplyT(func(exampleAccount storage.LookupAccountResult) (string, error) {
// 				return exampleAccount.PrimaryConnectionString, nil
// 			}).(pulumi.StringOutput),
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
//  $ pulumi import azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceAzureBlobStorage struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string. Conflicts with `sasUri` and `serviceEndpoint`.
	ConnectionString pulumi.StringPtrOutput `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SAS URI. Conflicts with `connectionString` and `serviceEndpoint`.
	SasUri pulumi.StringPtrOutput `pulumi:"sasUri"`
	// The Service Endpoint. Conflicts with `connectionString` and `sasUri`. Required with `useManagedIdentity`.
	ServiceEndpoint pulumi.StringPtrOutput `pulumi:"serviceEndpoint"`
	// The service principal id in which to authenticate against the Azure Blob Storage account. Required if `servicePrincipalKey` is set.
	ServicePrincipalId pulumi.StringPtrOutput `pulumi:"servicePrincipalId"`
	// The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `servicePrincipalId` is set.
	ServicePrincipalKey pulumi.StringPtrOutput `pulumi:"servicePrincipalKey"`
	// The tenant id or name in which to authenticate against the Azure Blob Storage account.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `servicePrincipalId` and `servicePrincipalKey`.
	UseManagedIdentity pulumi.BoolPtrOutput `pulumi:"useManagedIdentity"`
}

// NewLinkedServiceAzureBlobStorage registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceAzureBlobStorage(ctx *pulumi.Context,
	name string, args *LinkedServiceAzureBlobStorageArgs, opts ...pulumi.ResourceOption) (*LinkedServiceAzureBlobStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource LinkedServiceAzureBlobStorage
	err := ctx.RegisterResource("azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceAzureBlobStorage gets an existing LinkedServiceAzureBlobStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceAzureBlobStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceAzureBlobStorageState, opts ...pulumi.ResourceOption) (*LinkedServiceAzureBlobStorage, error) {
	var resource LinkedServiceAzureBlobStorage
	err := ctx.ReadResource("azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceAzureBlobStorage resources.
type linkedServiceAzureBlobStorageState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The connection string. Conflicts with `sasUri` and `serviceEndpoint`.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SAS URI. Conflicts with `connectionString` and `serviceEndpoint`.
	SasUri *string `pulumi:"sasUri"`
	// The Service Endpoint. Conflicts with `connectionString` and `sasUri`. Required with `useManagedIdentity`.
	ServiceEndpoint *string `pulumi:"serviceEndpoint"`
	// The service principal id in which to authenticate against the Azure Blob Storage account. Required if `servicePrincipalKey` is set.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `servicePrincipalId` is set.
	ServicePrincipalKey *string `pulumi:"servicePrincipalKey"`
	// The tenant id or name in which to authenticate against the Azure Blob Storage account.
	TenantId *string `pulumi:"tenantId"`
	// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `servicePrincipalId` and `servicePrincipalKey`.
	UseManagedIdentity *bool `pulumi:"useManagedIdentity"`
}

type LinkedServiceAzureBlobStorageState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The connection string. Conflicts with `sasUri` and `serviceEndpoint`.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// The SAS URI. Conflicts with `connectionString` and `serviceEndpoint`.
	SasUri pulumi.StringPtrInput
	// The Service Endpoint. Conflicts with `connectionString` and `sasUri`. Required with `useManagedIdentity`.
	ServiceEndpoint pulumi.StringPtrInput
	// The service principal id in which to authenticate against the Azure Blob Storage account. Required if `servicePrincipalKey` is set.
	ServicePrincipalId pulumi.StringPtrInput
	// The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `servicePrincipalId` is set.
	ServicePrincipalKey pulumi.StringPtrInput
	// The tenant id or name in which to authenticate against the Azure Blob Storage account.
	TenantId pulumi.StringPtrInput
	// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `servicePrincipalId` and `servicePrincipalKey`.
	UseManagedIdentity pulumi.BoolPtrInput
}

func (LinkedServiceAzureBlobStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceAzureBlobStorageState)(nil)).Elem()
}

type linkedServiceAzureBlobStorageArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The connection string. Conflicts with `sasUri` and `serviceEndpoint`.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SAS URI. Conflicts with `connectionString` and `serviceEndpoint`.
	SasUri *string `pulumi:"sasUri"`
	// The Service Endpoint. Conflicts with `connectionString` and `sasUri`. Required with `useManagedIdentity`.
	ServiceEndpoint *string `pulumi:"serviceEndpoint"`
	// The service principal id in which to authenticate against the Azure Blob Storage account. Required if `servicePrincipalKey` is set.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `servicePrincipalId` is set.
	ServicePrincipalKey *string `pulumi:"servicePrincipalKey"`
	// The tenant id or name in which to authenticate against the Azure Blob Storage account.
	TenantId *string `pulumi:"tenantId"`
	// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `servicePrincipalId` and `servicePrincipalKey`.
	UseManagedIdentity *bool `pulumi:"useManagedIdentity"`
}

// The set of arguments for constructing a LinkedServiceAzureBlobStorage resource.
type LinkedServiceAzureBlobStorageArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The connection string. Conflicts with `sasUri` and `serviceEndpoint`.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// The SAS URI. Conflicts with `connectionString` and `serviceEndpoint`.
	SasUri pulumi.StringPtrInput
	// The Service Endpoint. Conflicts with `connectionString` and `sasUri`. Required with `useManagedIdentity`.
	ServiceEndpoint pulumi.StringPtrInput
	// The service principal id in which to authenticate against the Azure Blob Storage account. Required if `servicePrincipalKey` is set.
	ServicePrincipalId pulumi.StringPtrInput
	// The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `servicePrincipalId` is set.
	ServicePrincipalKey pulumi.StringPtrInput
	// The tenant id or name in which to authenticate against the Azure Blob Storage account.
	TenantId pulumi.StringPtrInput
	// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `servicePrincipalId` and `servicePrincipalKey`.
	UseManagedIdentity pulumi.BoolPtrInput
}

func (LinkedServiceAzureBlobStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceAzureBlobStorageArgs)(nil)).Elem()
}

type LinkedServiceAzureBlobStorageInput interface {
	pulumi.Input

	ToLinkedServiceAzureBlobStorageOutput() LinkedServiceAzureBlobStorageOutput
	ToLinkedServiceAzureBlobStorageOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStorageOutput
}

func (*LinkedServiceAzureBlobStorage) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceAzureBlobStorage)(nil))
}

func (i *LinkedServiceAzureBlobStorage) ToLinkedServiceAzureBlobStorageOutput() LinkedServiceAzureBlobStorageOutput {
	return i.ToLinkedServiceAzureBlobStorageOutputWithContext(context.Background())
}

func (i *LinkedServiceAzureBlobStorage) ToLinkedServiceAzureBlobStorageOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureBlobStorageOutput)
}

func (i *LinkedServiceAzureBlobStorage) ToLinkedServiceAzureBlobStoragePtrOutput() LinkedServiceAzureBlobStoragePtrOutput {
	return i.ToLinkedServiceAzureBlobStoragePtrOutputWithContext(context.Background())
}

func (i *LinkedServiceAzureBlobStorage) ToLinkedServiceAzureBlobStoragePtrOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureBlobStoragePtrOutput)
}

type LinkedServiceAzureBlobStoragePtrInput interface {
	pulumi.Input

	ToLinkedServiceAzureBlobStoragePtrOutput() LinkedServiceAzureBlobStoragePtrOutput
	ToLinkedServiceAzureBlobStoragePtrOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStoragePtrOutput
}

type linkedServiceAzureBlobStoragePtrType LinkedServiceAzureBlobStorageArgs

func (*linkedServiceAzureBlobStoragePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceAzureBlobStorage)(nil))
}

func (i *linkedServiceAzureBlobStoragePtrType) ToLinkedServiceAzureBlobStoragePtrOutput() LinkedServiceAzureBlobStoragePtrOutput {
	return i.ToLinkedServiceAzureBlobStoragePtrOutputWithContext(context.Background())
}

func (i *linkedServiceAzureBlobStoragePtrType) ToLinkedServiceAzureBlobStoragePtrOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureBlobStoragePtrOutput)
}

// LinkedServiceAzureBlobStorageArrayInput is an input type that accepts LinkedServiceAzureBlobStorageArray and LinkedServiceAzureBlobStorageArrayOutput values.
// You can construct a concrete instance of `LinkedServiceAzureBlobStorageArrayInput` via:
//
//          LinkedServiceAzureBlobStorageArray{ LinkedServiceAzureBlobStorageArgs{...} }
type LinkedServiceAzureBlobStorageArrayInput interface {
	pulumi.Input

	ToLinkedServiceAzureBlobStorageArrayOutput() LinkedServiceAzureBlobStorageArrayOutput
	ToLinkedServiceAzureBlobStorageArrayOutputWithContext(context.Context) LinkedServiceAzureBlobStorageArrayOutput
}

type LinkedServiceAzureBlobStorageArray []LinkedServiceAzureBlobStorageInput

func (LinkedServiceAzureBlobStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LinkedServiceAzureBlobStorage)(nil))
}

func (i LinkedServiceAzureBlobStorageArray) ToLinkedServiceAzureBlobStorageArrayOutput() LinkedServiceAzureBlobStorageArrayOutput {
	return i.ToLinkedServiceAzureBlobStorageArrayOutputWithContext(context.Background())
}

func (i LinkedServiceAzureBlobStorageArray) ToLinkedServiceAzureBlobStorageArrayOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureBlobStorageArrayOutput)
}

// LinkedServiceAzureBlobStorageMapInput is an input type that accepts LinkedServiceAzureBlobStorageMap and LinkedServiceAzureBlobStorageMapOutput values.
// You can construct a concrete instance of `LinkedServiceAzureBlobStorageMapInput` via:
//
//          LinkedServiceAzureBlobStorageMap{ "key": LinkedServiceAzureBlobStorageArgs{...} }
type LinkedServiceAzureBlobStorageMapInput interface {
	pulumi.Input

	ToLinkedServiceAzureBlobStorageMapOutput() LinkedServiceAzureBlobStorageMapOutput
	ToLinkedServiceAzureBlobStorageMapOutputWithContext(context.Context) LinkedServiceAzureBlobStorageMapOutput
}

type LinkedServiceAzureBlobStorageMap map[string]LinkedServiceAzureBlobStorageInput

func (LinkedServiceAzureBlobStorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LinkedServiceAzureBlobStorage)(nil))
}

func (i LinkedServiceAzureBlobStorageMap) ToLinkedServiceAzureBlobStorageMapOutput() LinkedServiceAzureBlobStorageMapOutput {
	return i.ToLinkedServiceAzureBlobStorageMapOutputWithContext(context.Background())
}

func (i LinkedServiceAzureBlobStorageMap) ToLinkedServiceAzureBlobStorageMapOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureBlobStorageMapOutput)
}

type LinkedServiceAzureBlobStorageOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceAzureBlobStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceAzureBlobStorage)(nil))
}

func (o LinkedServiceAzureBlobStorageOutput) ToLinkedServiceAzureBlobStorageOutput() LinkedServiceAzureBlobStorageOutput {
	return o
}

func (o LinkedServiceAzureBlobStorageOutput) ToLinkedServiceAzureBlobStorageOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStorageOutput {
	return o
}

func (o LinkedServiceAzureBlobStorageOutput) ToLinkedServiceAzureBlobStoragePtrOutput() LinkedServiceAzureBlobStoragePtrOutput {
	return o.ToLinkedServiceAzureBlobStoragePtrOutputWithContext(context.Background())
}

func (o LinkedServiceAzureBlobStorageOutput) ToLinkedServiceAzureBlobStoragePtrOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStoragePtrOutput {
	return o.ApplyT(func(v LinkedServiceAzureBlobStorage) *LinkedServiceAzureBlobStorage {
		return &v
	}).(LinkedServiceAzureBlobStoragePtrOutput)
}

type LinkedServiceAzureBlobStoragePtrOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceAzureBlobStoragePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceAzureBlobStorage)(nil))
}

func (o LinkedServiceAzureBlobStoragePtrOutput) ToLinkedServiceAzureBlobStoragePtrOutput() LinkedServiceAzureBlobStoragePtrOutput {
	return o
}

func (o LinkedServiceAzureBlobStoragePtrOutput) ToLinkedServiceAzureBlobStoragePtrOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStoragePtrOutput {
	return o
}

type LinkedServiceAzureBlobStorageArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceAzureBlobStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedServiceAzureBlobStorage)(nil))
}

func (o LinkedServiceAzureBlobStorageArrayOutput) ToLinkedServiceAzureBlobStorageArrayOutput() LinkedServiceAzureBlobStorageArrayOutput {
	return o
}

func (o LinkedServiceAzureBlobStorageArrayOutput) ToLinkedServiceAzureBlobStorageArrayOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStorageArrayOutput {
	return o
}

func (o LinkedServiceAzureBlobStorageArrayOutput) Index(i pulumi.IntInput) LinkedServiceAzureBlobStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedServiceAzureBlobStorage {
		return vs[0].([]LinkedServiceAzureBlobStorage)[vs[1].(int)]
	}).(LinkedServiceAzureBlobStorageOutput)
}

type LinkedServiceAzureBlobStorageMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceAzureBlobStorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedServiceAzureBlobStorage)(nil))
}

func (o LinkedServiceAzureBlobStorageMapOutput) ToLinkedServiceAzureBlobStorageMapOutput() LinkedServiceAzureBlobStorageMapOutput {
	return o
}

func (o LinkedServiceAzureBlobStorageMapOutput) ToLinkedServiceAzureBlobStorageMapOutputWithContext(ctx context.Context) LinkedServiceAzureBlobStorageMapOutput {
	return o
}

func (o LinkedServiceAzureBlobStorageMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceAzureBlobStorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkedServiceAzureBlobStorage {
		return vs[0].(map[string]LinkedServiceAzureBlobStorage)[vs[1].(string)]
	}).(LinkedServiceAzureBlobStorageOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceAzureBlobStorageOutput{})
	pulumi.RegisterOutputType(LinkedServiceAzureBlobStoragePtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceAzureBlobStorageArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceAzureBlobStorageMapOutput{})
}
