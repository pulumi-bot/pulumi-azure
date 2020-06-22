// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iotcentral

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an IoT Central Application
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/iotcentral"
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
// 		_, err = iotcentral.NewApplication(ctx, "exampleApplication", &iotcentral.ApplicationArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			SubDomain:         pulumi.String("example-iotcentral-app-subdomain"),
// 			DisplayName:       pulumi.String("example-iotcentral-app-display-name"),
// 			Sku:               pulumi.String("S1"),
// 			Template:          pulumi.String("iotc-default@1.0.0"),
// 			Tags: pulumi.Map{
// 				"Foo": pulumi.String("Bar"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Application struct {
	pulumi.CustomResourceState

	// A `displayName` name. Custom display name for the IoT Central application. Default is resource name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// A `subDomain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
	SubDomain pulumi.StringOutput `pulumi:"subDomain"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `template` name. IoT Central application template name. Default is a custom application.
	Template pulumi.StringOutput `pulumi:"template"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SubDomain == nil {
		return nil, errors.New("missing required argument 'SubDomain'")
	}
	if args == nil {
		args = &ApplicationArgs{}
	}
	var resource Application
	err := ctx.RegisterResource("azure:iotcentral/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure:iotcentral/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// A `displayName` name. Custom display name for the IoT Central application. Default is resource name.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
	Sku *string `pulumi:"sku"`
	// A `subDomain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
	SubDomain *string `pulumi:"subDomain"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `template` name. IoT Central application template name. Default is a custom application.
	Template *string `pulumi:"template"`
}

type ApplicationState struct {
	// A `displayName` name. Custom display name for the IoT Central application. Default is resource name.
	DisplayName pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
	Sku pulumi.StringPtrInput
	// A `subDomain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
	SubDomain pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `template` name. IoT Central application template name. Default is a custom application.
	Template pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// A `displayName` name. Custom display name for the IoT Central application. Default is resource name.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
	Sku *string `pulumi:"sku"`
	// A `subDomain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
	SubDomain string `pulumi:"subDomain"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `template` name. IoT Central application template name. Default is a custom application.
	Template *string `pulumi:"template"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// A `displayName` name. Custom display name for the IoT Central application. Default is resource name.
	DisplayName pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource has to be create. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `sku` name. Possible values is `ST1`, `ST2`, Default value is `ST1`
	Sku pulumi.StringPtrInput
	// A `subDomain` name. Subdomain for the IoT Central URL. Each application must have a unique subdomain.
	SubDomain pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `template` name. IoT Central application template name. Default is a custom application.
	Template pulumi.StringPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}
