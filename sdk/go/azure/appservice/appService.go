// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an App Service (within an App Service Plan).
//
// > **Note:** When using Slots - the `appSettings`, `connectionString` and `siteConfig` blocks on the `appservice.AppService` resource will be overwritten when promoting a Slot using the `appservice.ActiveSlot` resource.
//
// ## Example Usage
//
// This example provisions a Windows App Service.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
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
// 		examplePlan, err := appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Standard"),
// 				Size: pulumi.String("S1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewAppService(ctx, "exampleAppService", &appservice.AppServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AppServicePlanId:  examplePlan.ID(),
// 			SiteConfig: &appservice.AppServiceSiteConfigArgs{
// 				DotnetFrameworkVersion: pulumi.String("v4.0"),
// 				ScmType:                pulumi.String("LocalGit"),
// 			},
// 			AppSettings: pulumi.StringMap{
// 				"SOME_KEY": pulumi.String("some-value"),
// 			},
// 			ConnectionStrings: appservice.AppServiceConnectionStringArray{
// 				&appservice.AppServiceConnectionStringArgs{
// 					Name:  pulumi.String("Database"),
// 					Type:  pulumi.String("SQLServer"),
// 					Value: pulumi.String("Server=some-server.mydomain.com;Integrated Security=SSPI"),
// 				},
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
// App Services can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/appService:AppService instance1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1
// ```
type AppService struct {
	pulumi.CustomResourceState

	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId pulumi.StringOutput `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapOutput `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings AppServiceAuthSettingsOutput `pulumi:"authSettings"`
	// A `backup` block as defined below.
	Backup AppServiceBackupPtrOutput `pulumi:"backup"`
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrOutput `pulumi:"clientAffinityEnabled"`
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled pulumi.BoolPtrOutput `pulumi:"clientCertEnabled"`
	// One or more `connectionString` blocks as defined below.
	ConnectionStrings AppServiceConnectionStringArrayOutput `pulumi:"connectionStrings"`
	// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
	CustomDomainVerificationId pulumi.StringOutput `pulumi:"customDomainVerificationId"`
	// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
	DefaultSiteHostname pulumi.StringOutput `pulumi:"defaultSiteHostname"`
	// Is the App Service Enabled?
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrOutput `pulumi:"httpsOnly"`
	// A Managed Service Identity block as defined below.
	Identity AppServiceIdentityOutput `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `logs` block as defined below.
	Logs AppServiceLogsOutput `pulumi:"logs"`
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of outbound IP addresses - such as `["52.23.25.3", "52.143.43.12"]`
	OutboundIpAddressLists pulumi.StringArrayOutput `pulumi:"outboundIpAddressLists"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
	OutboundIpAddresses pulumi.StringOutput `pulumi:"outboundIpAddresses"`
	// A list of outbound IP addresses - such as `["52.23.25.3", "52.143.43.12", "52.143.43.17"]` - not all of which are necessarily in use. Superset of `outboundIpAddressList`.
	PossibleOutboundIpAddressLists pulumi.StringArrayOutput `pulumi:"possibleOutboundIpAddressLists"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
	PossibleOutboundIpAddresses pulumi.StringOutput `pulumi:"possibleOutboundIpAddresses"`
	// The name of the resource group in which to create the App Service.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `siteConfig` block as defined below.
	SiteConfig AppServiceSiteConfigOutput `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials AppServiceSiteCredentialArrayOutput `pulumi:"siteCredentials"`
	// A Source Control block as defined below
	SourceControl AppServiceSourceControlOutput `pulumi:"sourceControl"`
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts AppServiceStorageAccountArrayOutput `pulumi:"storageAccounts"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAppService registers a new resource with the given unique name, arguments, and options.
func NewAppService(ctx *pulumi.Context,
	name string, args *AppServiceArgs, opts ...pulumi.ResourceOption) (*AppService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppServicePlanId == nil {
		return nil, errors.New("invalid value for required argument 'AppServicePlanId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource AppService
	err := ctx.RegisterResource("azure:appservice/appService:AppService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppService gets an existing AppService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceState, opts ...pulumi.ResourceOption) (*AppService, error) {
	var resource AppService
	err := ctx.ReadResource("azure:appservice/appService:AppService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppService resources.
type appServiceState struct {
	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId *string `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings map[string]string `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings *AppServiceAuthSettings `pulumi:"authSettings"`
	// A `backup` block as defined below.
	Backup *AppServiceBackup `pulumi:"backup"`
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
	// One or more `connectionString` blocks as defined below.
	ConnectionStrings []AppServiceConnectionString `pulumi:"connectionStrings"`
	// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
	CustomDomainVerificationId *string `pulumi:"customDomainVerificationId"`
	// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
	DefaultSiteHostname *string `pulumi:"defaultSiteHostname"`
	// Is the App Service Enabled?
	Enabled *bool `pulumi:"enabled"`
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// A Managed Service Identity block as defined below.
	Identity *AppServiceIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `logs` block as defined below.
	Logs *AppServiceLogs `pulumi:"logs"`
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of outbound IP addresses - such as `["52.23.25.3", "52.143.43.12"]`
	OutboundIpAddressLists []string `pulumi:"outboundIpAddressLists"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
	OutboundIpAddresses *string `pulumi:"outboundIpAddresses"`
	// A list of outbound IP addresses - such as `["52.23.25.3", "52.143.43.12", "52.143.43.17"]` - not all of which are necessarily in use. Superset of `outboundIpAddressList`.
	PossibleOutboundIpAddressLists []string `pulumi:"possibleOutboundIpAddressLists"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
	PossibleOutboundIpAddresses *string `pulumi:"possibleOutboundIpAddresses"`
	// The name of the resource group in which to create the App Service.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `siteConfig` block as defined below.
	SiteConfig *AppServiceSiteConfig `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials []AppServiceSiteCredential `pulumi:"siteCredentials"`
	// A Source Control block as defined below
	SourceControl *AppServiceSourceControl `pulumi:"sourceControl"`
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts []AppServiceStorageAccount `pulumi:"storageAccounts"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type AppServiceState struct {
	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId pulumi.StringPtrInput
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapInput
	// A `authSettings` block as defined below.
	AuthSettings AppServiceAuthSettingsPtrInput
	// A `backup` block as defined below.
	Backup AppServiceBackupPtrInput
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled pulumi.BoolPtrInput
	// One or more `connectionString` blocks as defined below.
	ConnectionStrings AppServiceConnectionStringArrayInput
	// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
	CustomDomainVerificationId pulumi.StringPtrInput
	// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
	DefaultSiteHostname pulumi.StringPtrInput
	// Is the App Service Enabled?
	Enabled pulumi.BoolPtrInput
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// A Managed Service Identity block as defined below.
	Identity AppServiceIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `logs` block as defined below.
	Logs AppServiceLogsPtrInput
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of outbound IP addresses - such as `["52.23.25.3", "52.143.43.12"]`
	OutboundIpAddressLists pulumi.StringArrayInput
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
	OutboundIpAddresses pulumi.StringPtrInput
	// A list of outbound IP addresses - such as `["52.23.25.3", "52.143.43.12", "52.143.43.17"]` - not all of which are necessarily in use. Superset of `outboundIpAddressList`.
	PossibleOutboundIpAddressLists pulumi.StringArrayInput
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
	PossibleOutboundIpAddresses pulumi.StringPtrInput
	// The name of the resource group in which to create the App Service.
	ResourceGroupName pulumi.StringPtrInput
	// A `siteConfig` block as defined below.
	SiteConfig AppServiceSiteConfigPtrInput
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials AppServiceSiteCredentialArrayInput
	// A Source Control block as defined below
	SourceControl AppServiceSourceControlPtrInput
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts AppServiceStorageAccountArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AppServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceState)(nil)).Elem()
}

type appServiceArgs struct {
	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId string `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings map[string]string `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings *AppServiceAuthSettings `pulumi:"authSettings"`
	// A `backup` block as defined below.
	Backup *AppServiceBackup `pulumi:"backup"`
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
	// One or more `connectionString` blocks as defined below.
	ConnectionStrings []AppServiceConnectionString `pulumi:"connectionStrings"`
	// Is the App Service Enabled?
	Enabled *bool `pulumi:"enabled"`
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// A Managed Service Identity block as defined below.
	Identity *AppServiceIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `logs` block as defined below.
	Logs *AppServiceLogs `pulumi:"logs"`
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the App Service.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `siteConfig` block as defined below.
	SiteConfig *AppServiceSiteConfig `pulumi:"siteConfig"`
	// A Source Control block as defined below
	SourceControl *AppServiceSourceControl `pulumi:"sourceControl"`
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts []AppServiceStorageAccount `pulumi:"storageAccounts"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AppService resource.
type AppServiceArgs struct {
	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId pulumi.StringInput
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapInput
	// A `authSettings` block as defined below.
	AuthSettings AppServiceAuthSettingsPtrInput
	// A `backup` block as defined below.
	Backup AppServiceBackupPtrInput
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled pulumi.BoolPtrInput
	// One or more `connectionString` blocks as defined below.
	ConnectionStrings AppServiceConnectionStringArrayInput
	// Is the App Service Enabled?
	Enabled pulumi.BoolPtrInput
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// A Managed Service Identity block as defined below.
	Identity AppServiceIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `logs` block as defined below.
	Logs AppServiceLogsPtrInput
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the App Service.
	ResourceGroupName pulumi.StringInput
	// A `siteConfig` block as defined below.
	SiteConfig AppServiceSiteConfigPtrInput
	// A Source Control block as defined below
	SourceControl AppServiceSourceControlPtrInput
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts AppServiceStorageAccountArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AppServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceArgs)(nil)).Elem()
}

type AppServiceInput interface {
	pulumi.Input

	ToAppServiceOutput() AppServiceOutput
	ToAppServiceOutputWithContext(ctx context.Context) AppServiceOutput
}

func (*AppService) ElementType() reflect.Type {
	return reflect.TypeOf((*AppService)(nil))
}

func (i *AppService) ToAppServiceOutput() AppServiceOutput {
	return i.ToAppServiceOutputWithContext(context.Background())
}

func (i *AppService) ToAppServiceOutputWithContext(ctx context.Context) AppServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceOutput)
}

func (i *AppService) ToAppServicePtrOutput() AppServicePtrOutput {
	return i.ToAppServicePtrOutputWithContext(context.Background())
}

func (i *AppService) ToAppServicePtrOutputWithContext(ctx context.Context) AppServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServicePtrOutput)
}

type AppServicePtrInput interface {
	pulumi.Input

	ToAppServicePtrOutput() AppServicePtrOutput
	ToAppServicePtrOutputWithContext(ctx context.Context) AppServicePtrOutput
}

type appServicePtrType AppServiceArgs

func (*appServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppService)(nil))
}

func (i *appServicePtrType) ToAppServicePtrOutput() AppServicePtrOutput {
	return i.ToAppServicePtrOutputWithContext(context.Background())
}

func (i *appServicePtrType) ToAppServicePtrOutputWithContext(ctx context.Context) AppServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServicePtrOutput)
}

// AppServiceArrayInput is an input type that accepts AppServiceArray and AppServiceArrayOutput values.
// You can construct a concrete instance of `AppServiceArrayInput` via:
//
//          AppServiceArray{ AppServiceArgs{...} }
type AppServiceArrayInput interface {
	pulumi.Input

	ToAppServiceArrayOutput() AppServiceArrayOutput
	ToAppServiceArrayOutputWithContext(context.Context) AppServiceArrayOutput
}

type AppServiceArray []AppServiceInput

func (AppServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppService)(nil))
}

func (i AppServiceArray) ToAppServiceArrayOutput() AppServiceArrayOutput {
	return i.ToAppServiceArrayOutputWithContext(context.Background())
}

func (i AppServiceArray) ToAppServiceArrayOutputWithContext(ctx context.Context) AppServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceArrayOutput)
}

// AppServiceMapInput is an input type that accepts AppServiceMap and AppServiceMapOutput values.
// You can construct a concrete instance of `AppServiceMapInput` via:
//
//          AppServiceMap{ "key": AppServiceArgs{...} }
type AppServiceMapInput interface {
	pulumi.Input

	ToAppServiceMapOutput() AppServiceMapOutput
	ToAppServiceMapOutputWithContext(context.Context) AppServiceMapOutput
}

type AppServiceMap map[string]AppServiceInput

func (AppServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppService)(nil))
}

func (i AppServiceMap) ToAppServiceMapOutput() AppServiceMapOutput {
	return i.ToAppServiceMapOutputWithContext(context.Background())
}

func (i AppServiceMap) ToAppServiceMapOutputWithContext(ctx context.Context) AppServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceMapOutput)
}

type AppServiceOutput struct {
	*pulumi.OutputState
}

func (AppServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppService)(nil))
}

func (o AppServiceOutput) ToAppServiceOutput() AppServiceOutput {
	return o
}

func (o AppServiceOutput) ToAppServiceOutputWithContext(ctx context.Context) AppServiceOutput {
	return o
}

func (o AppServiceOutput) ToAppServicePtrOutput() AppServicePtrOutput {
	return o.ToAppServicePtrOutputWithContext(context.Background())
}

func (o AppServiceOutput) ToAppServicePtrOutputWithContext(ctx context.Context) AppServicePtrOutput {
	return o.ApplyT(func(v AppService) *AppService {
		return &v
	}).(AppServicePtrOutput)
}

type AppServicePtrOutput struct {
	*pulumi.OutputState
}

func (AppServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppService)(nil))
}

func (o AppServicePtrOutput) ToAppServicePtrOutput() AppServicePtrOutput {
	return o
}

func (o AppServicePtrOutput) ToAppServicePtrOutputWithContext(ctx context.Context) AppServicePtrOutput {
	return o
}

type AppServiceArrayOutput struct{ *pulumi.OutputState }

func (AppServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppService)(nil))
}

func (o AppServiceArrayOutput) ToAppServiceArrayOutput() AppServiceArrayOutput {
	return o
}

func (o AppServiceArrayOutput) ToAppServiceArrayOutputWithContext(ctx context.Context) AppServiceArrayOutput {
	return o
}

func (o AppServiceArrayOutput) Index(i pulumi.IntInput) AppServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppService {
		return vs[0].([]AppService)[vs[1].(int)]
	}).(AppServiceOutput)
}

type AppServiceMapOutput struct{ *pulumi.OutputState }

func (AppServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppService)(nil))
}

func (o AppServiceMapOutput) ToAppServiceMapOutput() AppServiceMapOutput {
	return o
}

func (o AppServiceMapOutput) ToAppServiceMapOutputWithContext(ctx context.Context) AppServiceMapOutput {
	return o
}

func (o AppServiceMapOutput) MapIndex(k pulumi.StringInput) AppServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppService {
		return vs[0].(map[string]AppService)[vs[1].(string)]
	}).(AppServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(AppServiceOutput{})
	pulumi.RegisterOutputType(AppServicePtrOutput{})
	pulumi.RegisterOutputType(AppServiceArrayOutput{})
	pulumi.RegisterOutputType(AppServiceMapOutput{})
}
