// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manage an Azure Spring Cloud Application.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appplatform"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("Southeast Asia"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSpringCloudService, err := appplatform.NewSpringCloudService(ctx, "exampleSpringCloudService", &appplatform.SpringCloudServiceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appplatform.NewSpringCloudApp(ctx, "exampleSpringCloudApp", &appplatform.SpringCloudAppArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServiceName:       exampleSpringCloudService.Name,
// 			Identity: &appplatform.SpringCloudAppIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
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
// Spring Cloud Application can be imported using the `resource id`, e.g.
type SpringCloudApp struct {
	pulumi.CustomResourceState

	// An `identity` block as defined below.
	Identity SpringCloudAppIdentityPtrOutput `pulumi:"identity"`
	// Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewSpringCloudApp registers a new resource with the given unique name, arguments, and options.
func NewSpringCloudApp(ctx *pulumi.Context,
	name string, args *SpringCloudAppArgs, opts ...pulumi.ResourceOption) (*SpringCloudApp, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &SpringCloudAppArgs{}
	}
	var resource SpringCloudApp
	err := ctx.RegisterResource("azure:appplatform/springCloudApp:SpringCloudApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpringCloudApp gets an existing SpringCloudApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpringCloudApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpringCloudAppState, opts ...pulumi.ResourceOption) (*SpringCloudApp, error) {
	var resource SpringCloudApp
	err := ctx.ReadResource("azure:appplatform/springCloudApp:SpringCloudApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpringCloudApp resources.
type springCloudAppState struct {
	// An `identity` block as defined below.
	Identity *SpringCloudAppIdentity `pulumi:"identity"`
	// Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName *string `pulumi:"serviceName"`
}

type SpringCloudAppState struct {
	// An `identity` block as defined below.
	Identity SpringCloudAppIdentityPtrInput
	// Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName pulumi.StringPtrInput
}

func (SpringCloudAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudAppState)(nil)).Elem()
}

type springCloudAppArgs struct {
	// An `identity` block as defined below.
	Identity *SpringCloudAppIdentity `pulumi:"identity"`
	// Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a SpringCloudApp resource.
type SpringCloudAppArgs struct {
	// An `identity` block as defined below.
	Identity SpringCloudAppIdentityPtrInput
	// Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName pulumi.StringInput
}

func (SpringCloudAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudAppArgs)(nil)).Elem()
}
