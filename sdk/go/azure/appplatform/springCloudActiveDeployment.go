// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Active Azure Spring Cloud Deployment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appplatform"
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
// 		exampleSpringCloudService, err := appplatform.NewSpringCloudService(ctx, "exampleSpringCloudService", &appplatform.SpringCloudServiceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSpringCloudApp, err := appplatform.NewSpringCloudApp(ctx, "exampleSpringCloudApp", &appplatform.SpringCloudAppArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServiceName:       exampleSpringCloudService.Name,
// 			Identity: &appplatform.SpringCloudAppIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSpringCloudJavaDeployment, err := appplatform.NewSpringCloudJavaDeployment(ctx, "exampleSpringCloudJavaDeployment", &appplatform.SpringCloudJavaDeploymentArgs{
// 			SpringCloudAppId: exampleSpringCloudApp.ID(),
// 			Cpu:              pulumi.Int(2),
// 			MemoryInGb:       pulumi.Int(4),
// 			InstanceCount:    pulumi.Int(2),
// 			JvmOptions:       pulumi.String("-XX:+PrintGC"),
// 			RuntimeVersion:   pulumi.String("Java_11"),
// 			EnvironmentVariables: pulumi.StringMap{
// 				"Env": pulumi.String("Staging"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appplatform.NewSpringCloudActiveDeployment(ctx, "exampleSpringCloudActiveDeployment", &appplatform.SpringCloudActiveDeploymentArgs{
// 			SpringCloudAppId: exampleSpringCloudApp.ID(),
// 			DeploymentName:   exampleSpringCloudJavaDeployment.Name,
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
// Spring Cloud Active Deployment can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appplatform/springCloudActiveDeployment:SpringCloudActiveDeployment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/Spring/service1/apps/app1
// ```
type SpringCloudActiveDeployment struct {
	pulumi.CustomResourceState

	// Specifies the name of Spring Cloud Deployment which is going to be active.
	DeploymentName pulumi.StringOutput `pulumi:"deploymentName"`
	// Specifies the id of the Spring Cloud Application. Changing this forces a new resource to be created.
	SpringCloudAppId pulumi.StringOutput `pulumi:"springCloudAppId"`
}

// NewSpringCloudActiveDeployment registers a new resource with the given unique name, arguments, and options.
func NewSpringCloudActiveDeployment(ctx *pulumi.Context,
	name string, args *SpringCloudActiveDeploymentArgs, opts ...pulumi.ResourceOption) (*SpringCloudActiveDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeploymentName == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentName'")
	}
	if args.SpringCloudAppId == nil {
		return nil, errors.New("invalid value for required argument 'SpringCloudAppId'")
	}
	var resource SpringCloudActiveDeployment
	err := ctx.RegisterResource("azure:appplatform/springCloudActiveDeployment:SpringCloudActiveDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpringCloudActiveDeployment gets an existing SpringCloudActiveDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpringCloudActiveDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpringCloudActiveDeploymentState, opts ...pulumi.ResourceOption) (*SpringCloudActiveDeployment, error) {
	var resource SpringCloudActiveDeployment
	err := ctx.ReadResource("azure:appplatform/springCloudActiveDeployment:SpringCloudActiveDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpringCloudActiveDeployment resources.
type springCloudActiveDeploymentState struct {
	// Specifies the name of Spring Cloud Deployment which is going to be active.
	DeploymentName *string `pulumi:"deploymentName"`
	// Specifies the id of the Spring Cloud Application. Changing this forces a new resource to be created.
	SpringCloudAppId *string `pulumi:"springCloudAppId"`
}

type SpringCloudActiveDeploymentState struct {
	// Specifies the name of Spring Cloud Deployment which is going to be active.
	DeploymentName pulumi.StringPtrInput
	// Specifies the id of the Spring Cloud Application. Changing this forces a new resource to be created.
	SpringCloudAppId pulumi.StringPtrInput
}

func (SpringCloudActiveDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudActiveDeploymentState)(nil)).Elem()
}

type springCloudActiveDeploymentArgs struct {
	// Specifies the name of Spring Cloud Deployment which is going to be active.
	DeploymentName string `pulumi:"deploymentName"`
	// Specifies the id of the Spring Cloud Application. Changing this forces a new resource to be created.
	SpringCloudAppId string `pulumi:"springCloudAppId"`
}

// The set of arguments for constructing a SpringCloudActiveDeployment resource.
type SpringCloudActiveDeploymentArgs struct {
	// Specifies the name of Spring Cloud Deployment which is going to be active.
	DeploymentName pulumi.StringInput
	// Specifies the id of the Spring Cloud Application. Changing this forces a new resource to be created.
	SpringCloudAppId pulumi.StringInput
}

func (SpringCloudActiveDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudActiveDeploymentArgs)(nil)).Elem()
}

type SpringCloudActiveDeploymentInput interface {
	pulumi.Input

	ToSpringCloudActiveDeploymentOutput() SpringCloudActiveDeploymentOutput
	ToSpringCloudActiveDeploymentOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentOutput
}

func (*SpringCloudActiveDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((*SpringCloudActiveDeployment)(nil))
}

func (i *SpringCloudActiveDeployment) ToSpringCloudActiveDeploymentOutput() SpringCloudActiveDeploymentOutput {
	return i.ToSpringCloudActiveDeploymentOutputWithContext(context.Background())
}

func (i *SpringCloudActiveDeployment) ToSpringCloudActiveDeploymentOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudActiveDeploymentOutput)
}

func (i *SpringCloudActiveDeployment) ToSpringCloudActiveDeploymentPtrOutput() SpringCloudActiveDeploymentPtrOutput {
	return i.ToSpringCloudActiveDeploymentPtrOutputWithContext(context.Background())
}

func (i *SpringCloudActiveDeployment) ToSpringCloudActiveDeploymentPtrOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudActiveDeploymentPtrOutput)
}

type SpringCloudActiveDeploymentPtrInput interface {
	pulumi.Input

	ToSpringCloudActiveDeploymentPtrOutput() SpringCloudActiveDeploymentPtrOutput
	ToSpringCloudActiveDeploymentPtrOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentPtrOutput
}

type springCloudActiveDeploymentPtrType SpringCloudActiveDeploymentArgs

func (*springCloudActiveDeploymentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SpringCloudActiveDeployment)(nil))
}

func (i *springCloudActiveDeploymentPtrType) ToSpringCloudActiveDeploymentPtrOutput() SpringCloudActiveDeploymentPtrOutput {
	return i.ToSpringCloudActiveDeploymentPtrOutputWithContext(context.Background())
}

func (i *springCloudActiveDeploymentPtrType) ToSpringCloudActiveDeploymentPtrOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudActiveDeploymentPtrOutput)
}

// SpringCloudActiveDeploymentArrayInput is an input type that accepts SpringCloudActiveDeploymentArray and SpringCloudActiveDeploymentArrayOutput values.
// You can construct a concrete instance of `SpringCloudActiveDeploymentArrayInput` via:
//
//          SpringCloudActiveDeploymentArray{ SpringCloudActiveDeploymentArgs{...} }
type SpringCloudActiveDeploymentArrayInput interface {
	pulumi.Input

	ToSpringCloudActiveDeploymentArrayOutput() SpringCloudActiveDeploymentArrayOutput
	ToSpringCloudActiveDeploymentArrayOutputWithContext(context.Context) SpringCloudActiveDeploymentArrayOutput
}

type SpringCloudActiveDeploymentArray []SpringCloudActiveDeploymentInput

func (SpringCloudActiveDeploymentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SpringCloudActiveDeployment)(nil))
}

func (i SpringCloudActiveDeploymentArray) ToSpringCloudActiveDeploymentArrayOutput() SpringCloudActiveDeploymentArrayOutput {
	return i.ToSpringCloudActiveDeploymentArrayOutputWithContext(context.Background())
}

func (i SpringCloudActiveDeploymentArray) ToSpringCloudActiveDeploymentArrayOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudActiveDeploymentArrayOutput)
}

// SpringCloudActiveDeploymentMapInput is an input type that accepts SpringCloudActiveDeploymentMap and SpringCloudActiveDeploymentMapOutput values.
// You can construct a concrete instance of `SpringCloudActiveDeploymentMapInput` via:
//
//          SpringCloudActiveDeploymentMap{ "key": SpringCloudActiveDeploymentArgs{...} }
type SpringCloudActiveDeploymentMapInput interface {
	pulumi.Input

	ToSpringCloudActiveDeploymentMapOutput() SpringCloudActiveDeploymentMapOutput
	ToSpringCloudActiveDeploymentMapOutputWithContext(context.Context) SpringCloudActiveDeploymentMapOutput
}

type SpringCloudActiveDeploymentMap map[string]SpringCloudActiveDeploymentInput

func (SpringCloudActiveDeploymentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SpringCloudActiveDeployment)(nil))
}

func (i SpringCloudActiveDeploymentMap) ToSpringCloudActiveDeploymentMapOutput() SpringCloudActiveDeploymentMapOutput {
	return i.ToSpringCloudActiveDeploymentMapOutputWithContext(context.Background())
}

func (i SpringCloudActiveDeploymentMap) ToSpringCloudActiveDeploymentMapOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudActiveDeploymentMapOutput)
}

type SpringCloudActiveDeploymentOutput struct {
	*pulumi.OutputState
}

func (SpringCloudActiveDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpringCloudActiveDeployment)(nil))
}

func (o SpringCloudActiveDeploymentOutput) ToSpringCloudActiveDeploymentOutput() SpringCloudActiveDeploymentOutput {
	return o
}

func (o SpringCloudActiveDeploymentOutput) ToSpringCloudActiveDeploymentOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentOutput {
	return o
}

func (o SpringCloudActiveDeploymentOutput) ToSpringCloudActiveDeploymentPtrOutput() SpringCloudActiveDeploymentPtrOutput {
	return o.ToSpringCloudActiveDeploymentPtrOutputWithContext(context.Background())
}

func (o SpringCloudActiveDeploymentOutput) ToSpringCloudActiveDeploymentPtrOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentPtrOutput {
	return o.ApplyT(func(v SpringCloudActiveDeployment) *SpringCloudActiveDeployment {
		return &v
	}).(SpringCloudActiveDeploymentPtrOutput)
}

type SpringCloudActiveDeploymentPtrOutput struct {
	*pulumi.OutputState
}

func (SpringCloudActiveDeploymentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpringCloudActiveDeployment)(nil))
}

func (o SpringCloudActiveDeploymentPtrOutput) ToSpringCloudActiveDeploymentPtrOutput() SpringCloudActiveDeploymentPtrOutput {
	return o
}

func (o SpringCloudActiveDeploymentPtrOutput) ToSpringCloudActiveDeploymentPtrOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentPtrOutput {
	return o
}

type SpringCloudActiveDeploymentArrayOutput struct{ *pulumi.OutputState }

func (SpringCloudActiveDeploymentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpringCloudActiveDeployment)(nil))
}

func (o SpringCloudActiveDeploymentArrayOutput) ToSpringCloudActiveDeploymentArrayOutput() SpringCloudActiveDeploymentArrayOutput {
	return o
}

func (o SpringCloudActiveDeploymentArrayOutput) ToSpringCloudActiveDeploymentArrayOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentArrayOutput {
	return o
}

func (o SpringCloudActiveDeploymentArrayOutput) Index(i pulumi.IntInput) SpringCloudActiveDeploymentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpringCloudActiveDeployment {
		return vs[0].([]SpringCloudActiveDeployment)[vs[1].(int)]
	}).(SpringCloudActiveDeploymentOutput)
}

type SpringCloudActiveDeploymentMapOutput struct{ *pulumi.OutputState }

func (SpringCloudActiveDeploymentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SpringCloudActiveDeployment)(nil))
}

func (o SpringCloudActiveDeploymentMapOutput) ToSpringCloudActiveDeploymentMapOutput() SpringCloudActiveDeploymentMapOutput {
	return o
}

func (o SpringCloudActiveDeploymentMapOutput) ToSpringCloudActiveDeploymentMapOutputWithContext(ctx context.Context) SpringCloudActiveDeploymentMapOutput {
	return o
}

func (o SpringCloudActiveDeploymentMapOutput) MapIndex(k pulumi.StringInput) SpringCloudActiveDeploymentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SpringCloudActiveDeployment {
		return vs[0].(map[string]SpringCloudActiveDeployment)[vs[1].(string)]
	}).(SpringCloudActiveDeploymentOutput)
}

func init() {
	pulumi.RegisterOutputType(SpringCloudActiveDeploymentOutput{})
	pulumi.RegisterOutputType(SpringCloudActiveDeploymentPtrOutput{})
	pulumi.RegisterOutputType(SpringCloudActiveDeploymentArrayOutput{})
	pulumi.RegisterOutputType(SpringCloudActiveDeploymentMapOutput{})
}
