// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management Service.
//
// ## Disclaimers
//
// > **Note:** It's possible to define Custom Domains both within the `apimanagement.Service` resource via the `hostnameConfigurations` block and by using the `apimanagement.CustomDomain` resource. However it's not possible to use both methods to manage Custom Domains within an API Management Service, since there'll be conflicts.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/apimanagement"
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
// 		_, err = apimanagement.NewService(ctx, "exampleService", &apimanagement.ServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PublisherName:     pulumi.String("My Company"),
// 			PublisherEmail:    pulumi.String("company@exmaple.com"),
// 			SkuName:           pulumi.String("Developer_1"),
// 			Policy: &apimanagement.ServicePolicyArgs{
// 				XmlContent: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v", "    <policies>\n", "      <inbound />\n", "      <backend />\n", "      <outbound />\n", "      <on-error />\n", "    </policies>\n")),
// 			},
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
// API Management Services can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1
// ```
type Service struct {
	pulumi.CustomResourceState

	// One or more `additionalLocation` blocks as defined below.
	AdditionalLocations ServiceAdditionalLocationArrayOutput `pulumi:"additionalLocations"`
	// One or more (up to 10) `certificate` blocks as defined below.
	Certificates ServiceCertificateArrayOutput `pulumi:"certificates"`
	// The URL for the Developer Portal associated with this API Management service.
	DeveloperPortalUrl pulumi.StringOutput `pulumi:"developerPortalUrl"`
	// The URL of the Regional Gateway for the API Management Service in the specified region.
	GatewayRegionalUrl pulumi.StringOutput `pulumi:"gatewayRegionalUrl"`
	// The URL of the Gateway for the API Management Service.
	GatewayUrl pulumi.StringOutput `pulumi:"gatewayUrl"`
	// A `hostnameConfiguration` block as defined below.
	HostnameConfiguration ServiceHostnameConfigurationPtrOutput `pulumi:"hostnameConfiguration"`
	// An `identity` block is documented below.
	Identity ServiceIdentityPtrOutput `pulumi:"identity"`
	// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The URL for the Management API associated with this API Management service.
	ManagementApiUrl pulumi.StringOutput `pulumi:"managementApiUrl"`
	// The name of the API Management Service. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Email address from which the notification will be sent.
	NotificationSenderEmail pulumi.StringOutput `pulumi:"notificationSenderEmail"`
	// A `policy` block as defined below.
	Policy ServicePolicyOutput `pulumi:"policy"`
	// The URL for the Publisher Portal associated with this API Management service.
	PortalUrl pulumi.StringOutput `pulumi:"portalUrl"`
	// The Private IP addresses of the API Management Service.  Available only when the API Manager instance is using Virtual Network mode.
	PrivateIpAddresses pulumi.StringArrayOutput `pulumi:"privateIpAddresses"`
	// A `protocols` block as defined below.
	Protocols ServiceProtocolsOutput `pulumi:"protocols"`
	// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
	PublicIpAddresses pulumi.StringArrayOutput `pulumi:"publicIpAddresses"`
	// The email of publisher/company.
	PublisherEmail pulumi.StringOutput `pulumi:"publisherEmail"`
	// The name of publisher/company.
	PublisherName pulumi.StringOutput `pulumi:"publisherName"`
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
	ScmUrl pulumi.StringOutput `pulumi:"scmUrl"`
	// A `security` block as defined below.
	Security ServiceSecurityOutput `pulumi:"security"`
	// A `signIn` block as defined below.
	SignIn ServiceSignInOutput `pulumi:"signIn"`
	// A `signUp` block as defined below.
	SignUp ServiceSignUpOutput `pulumi:"signUp"`
	// `skuName` is a string consisting of two parts separated by an underscore(\_). The first part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `virtualNetworkConfiguration` block as defined below. Required when `virtualNetworkType` is `External` or `Internal`.
	VirtualNetworkConfiguration ServiceVirtualNetworkConfigurationPtrOutput `pulumi:"virtualNetworkConfiguration"`
	// The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`.
	// > **NOTE:** Please ensure that in the subnet, inbound port 3443 is open when `virtualNetworkType` is `Internal` or `External`. And please ensure other necessary ports are open according to [api management network configuration](https://docs.microsoft.com/en-us/azure/api-management/api-management-using-with-vnet#-common-network-configuration-issues).
	VirtualNetworkType pulumi.StringPtrOutput `pulumi:"virtualNetworkType"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublisherEmail == nil {
		return nil, errors.New("invalid value for required argument 'PublisherEmail'")
	}
	if args.PublisherName == nil {
		return nil, errors.New("invalid value for required argument 'PublisherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	var resource Service
	err := ctx.RegisterResource("azure:apimanagement/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure:apimanagement/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// One or more `additionalLocation` blocks as defined below.
	AdditionalLocations []ServiceAdditionalLocation `pulumi:"additionalLocations"`
	// One or more (up to 10) `certificate` blocks as defined below.
	Certificates []ServiceCertificate `pulumi:"certificates"`
	// The URL for the Developer Portal associated with this API Management service.
	DeveloperPortalUrl *string `pulumi:"developerPortalUrl"`
	// The URL of the Regional Gateway for the API Management Service in the specified region.
	GatewayRegionalUrl *string `pulumi:"gatewayRegionalUrl"`
	// The URL of the Gateway for the API Management Service.
	GatewayUrl *string `pulumi:"gatewayUrl"`
	// A `hostnameConfiguration` block as defined below.
	HostnameConfiguration *ServiceHostnameConfiguration `pulumi:"hostnameConfiguration"`
	// An `identity` block is documented below.
	Identity *ServiceIdentity `pulumi:"identity"`
	// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The URL for the Management API associated with this API Management service.
	ManagementApiUrl *string `pulumi:"managementApiUrl"`
	// The name of the API Management Service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Email address from which the notification will be sent.
	NotificationSenderEmail *string `pulumi:"notificationSenderEmail"`
	// A `policy` block as defined below.
	Policy *ServicePolicy `pulumi:"policy"`
	// The URL for the Publisher Portal associated with this API Management service.
	PortalUrl *string `pulumi:"portalUrl"`
	// The Private IP addresses of the API Management Service.  Available only when the API Manager instance is using Virtual Network mode.
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	// A `protocols` block as defined below.
	Protocols *ServiceProtocols `pulumi:"protocols"`
	// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
	PublicIpAddresses []string `pulumi:"publicIpAddresses"`
	// The email of publisher/company.
	PublisherEmail *string `pulumi:"publisherEmail"`
	// The name of publisher/company.
	PublisherName *string `pulumi:"publisherName"`
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
	ScmUrl *string `pulumi:"scmUrl"`
	// A `security` block as defined below.
	Security *ServiceSecurity `pulumi:"security"`
	// A `signIn` block as defined below.
	SignIn *ServiceSignIn `pulumi:"signIn"`
	// A `signUp` block as defined below.
	SignUp *ServiceSignUp `pulumi:"signUp"`
	// `skuName` is a string consisting of two parts separated by an underscore(\_). The first part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `virtualNetworkConfiguration` block as defined below. Required when `virtualNetworkType` is `External` or `Internal`.
	VirtualNetworkConfiguration *ServiceVirtualNetworkConfiguration `pulumi:"virtualNetworkConfiguration"`
	// The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`.
	// > **NOTE:** Please ensure that in the subnet, inbound port 3443 is open when `virtualNetworkType` is `Internal` or `External`. And please ensure other necessary ports are open according to [api management network configuration](https://docs.microsoft.com/en-us/azure/api-management/api-management-using-with-vnet#-common-network-configuration-issues).
	VirtualNetworkType *string `pulumi:"virtualNetworkType"`
}

type ServiceState struct {
	// One or more `additionalLocation` blocks as defined below.
	AdditionalLocations ServiceAdditionalLocationArrayInput
	// One or more (up to 10) `certificate` blocks as defined below.
	Certificates ServiceCertificateArrayInput
	// The URL for the Developer Portal associated with this API Management service.
	DeveloperPortalUrl pulumi.StringPtrInput
	// The URL of the Regional Gateway for the API Management Service in the specified region.
	GatewayRegionalUrl pulumi.StringPtrInput
	// The URL of the Gateway for the API Management Service.
	GatewayUrl pulumi.StringPtrInput
	// A `hostnameConfiguration` block as defined below.
	HostnameConfiguration ServiceHostnameConfigurationPtrInput
	// An `identity` block is documented below.
	Identity ServiceIdentityPtrInput
	// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The URL for the Management API associated with this API Management service.
	ManagementApiUrl pulumi.StringPtrInput
	// The name of the API Management Service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Email address from which the notification will be sent.
	NotificationSenderEmail pulumi.StringPtrInput
	// A `policy` block as defined below.
	Policy ServicePolicyPtrInput
	// The URL for the Publisher Portal associated with this API Management service.
	PortalUrl pulumi.StringPtrInput
	// The Private IP addresses of the API Management Service.  Available only when the API Manager instance is using Virtual Network mode.
	PrivateIpAddresses pulumi.StringArrayInput
	// A `protocols` block as defined below.
	Protocols ServiceProtocolsPtrInput
	// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
	PublicIpAddresses pulumi.StringArrayInput
	// The email of publisher/company.
	PublisherEmail pulumi.StringPtrInput
	// The name of publisher/company.
	PublisherName pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
	ScmUrl pulumi.StringPtrInput
	// A `security` block as defined below.
	Security ServiceSecurityPtrInput
	// A `signIn` block as defined below.
	SignIn ServiceSignInPtrInput
	// A `signUp` block as defined below.
	SignUp ServiceSignUpPtrInput
	// `skuName` is a string consisting of two parts separated by an underscore(\_). The first part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
	SkuName pulumi.StringPtrInput
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapInput
	// A `virtualNetworkConfiguration` block as defined below. Required when `virtualNetworkType` is `External` or `Internal`.
	VirtualNetworkConfiguration ServiceVirtualNetworkConfigurationPtrInput
	// The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`.
	// > **NOTE:** Please ensure that in the subnet, inbound port 3443 is open when `virtualNetworkType` is `Internal` or `External`. And please ensure other necessary ports are open according to [api management network configuration](https://docs.microsoft.com/en-us/azure/api-management/api-management-using-with-vnet#-common-network-configuration-issues).
	VirtualNetworkType pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// One or more `additionalLocation` blocks as defined below.
	AdditionalLocations []ServiceAdditionalLocation `pulumi:"additionalLocations"`
	// One or more (up to 10) `certificate` blocks as defined below.
	Certificates []ServiceCertificate `pulumi:"certificates"`
	// A `hostnameConfiguration` block as defined below.
	HostnameConfiguration *ServiceHostnameConfiguration `pulumi:"hostnameConfiguration"`
	// An `identity` block is documented below.
	Identity *ServiceIdentity `pulumi:"identity"`
	// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the API Management Service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Email address from which the notification will be sent.
	NotificationSenderEmail *string `pulumi:"notificationSenderEmail"`
	// A `policy` block as defined below.
	Policy *ServicePolicy `pulumi:"policy"`
	// A `protocols` block as defined below.
	Protocols *ServiceProtocols `pulumi:"protocols"`
	// The email of publisher/company.
	PublisherEmail string `pulumi:"publisherEmail"`
	// The name of publisher/company.
	PublisherName string `pulumi:"publisherName"`
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `security` block as defined below.
	Security *ServiceSecurity `pulumi:"security"`
	// A `signIn` block as defined below.
	SignIn *ServiceSignIn `pulumi:"signIn"`
	// A `signUp` block as defined below.
	SignUp *ServiceSignUp `pulumi:"signUp"`
	// `skuName` is a string consisting of two parts separated by an underscore(\_). The first part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
	SkuName string `pulumi:"skuName"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `virtualNetworkConfiguration` block as defined below. Required when `virtualNetworkType` is `External` or `Internal`.
	VirtualNetworkConfiguration *ServiceVirtualNetworkConfiguration `pulumi:"virtualNetworkConfiguration"`
	// The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`.
	// > **NOTE:** Please ensure that in the subnet, inbound port 3443 is open when `virtualNetworkType` is `Internal` or `External`. And please ensure other necessary ports are open according to [api management network configuration](https://docs.microsoft.com/en-us/azure/api-management/api-management-using-with-vnet#-common-network-configuration-issues).
	VirtualNetworkType *string `pulumi:"virtualNetworkType"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// One or more `additionalLocation` blocks as defined below.
	AdditionalLocations ServiceAdditionalLocationArrayInput
	// One or more (up to 10) `certificate` blocks as defined below.
	Certificates ServiceCertificateArrayInput
	// A `hostnameConfiguration` block as defined below.
	HostnameConfiguration ServiceHostnameConfigurationPtrInput
	// An `identity` block is documented below.
	Identity ServiceIdentityPtrInput
	// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the API Management Service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Email address from which the notification will be sent.
	NotificationSenderEmail pulumi.StringPtrInput
	// A `policy` block as defined below.
	Policy ServicePolicyPtrInput
	// A `protocols` block as defined below.
	Protocols ServiceProtocolsPtrInput
	// The email of publisher/company.
	PublisherEmail pulumi.StringInput
	// The name of publisher/company.
	PublisherName pulumi.StringInput
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `security` block as defined below.
	Security ServiceSecurityPtrInput
	// A `signIn` block as defined below.
	SignIn ServiceSignInPtrInput
	// A `signUp` block as defined below.
	SignUp ServiceSignUpPtrInput
	// `skuName` is a string consisting of two parts separated by an underscore(\_). The first part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
	SkuName pulumi.StringInput
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapInput
	// A `virtualNetworkConfiguration` block as defined below. Required when `virtualNetworkType` is `External` or `Internal`.
	VirtualNetworkConfiguration ServiceVirtualNetworkConfigurationPtrInput
	// The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`.
	// > **NOTE:** Please ensure that in the subnet, inbound port 3443 is open when `virtualNetworkType` is `Internal` or `External`. And please ensure other necessary ports are open according to [api management network configuration](https://docs.microsoft.com/en-us/azure/api-management/api-management-using-with-vnet#-common-network-configuration-issues).
	VirtualNetworkType pulumi.StringPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

func (i *Service) ToServicePtrOutput() ServicePtrOutput {
	return i.ToServicePtrOutputWithContext(context.Background())
}

func (i *Service) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePtrOutput)
}

type ServicePtrInput interface {
	pulumi.Input

	ToServicePtrOutput() ServicePtrOutput
	ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput
}

type ServiceOutput struct {
	*pulumi.OutputState
}

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

type ServicePtrOutput struct {
	*pulumi.OutputState
}

func (ServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil))
}

func (o ServicePtrOutput) ToServicePtrOutput() ServicePtrOutput {
	return o
}

func (o ServicePtrOutput) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServicePtrOutput{})
}
