// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Service Fabric Mesh Secret.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/servicefabric"
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
// 		_, err = servicefabric.NewMeshSecret(ctx, "exampleMeshSecret", &servicefabric.MeshSecretArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
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
// Service Fabric Mesh Secret can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:servicefabric/meshSecret:MeshSecret secret1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceFabricMesh/secrets/secret1
// ```
type MeshSecret struct {
	pulumi.CustomResourceState

	// The type of the content stored in the secret value. Changing this forces a new resource to be created.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// A description of this Service Fabric Mesh Secret.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the Azure Region where the Service Fabric Mesh Secret should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the Service Fabric Mesh Secret exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewMeshSecret registers a new resource with the given unique name, arguments, and options.
func NewMeshSecret(ctx *pulumi.Context,
	name string, args *MeshSecretArgs, opts ...pulumi.ResourceOption) (*MeshSecret, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &MeshSecretArgs{}
	}
	var resource MeshSecret
	err := ctx.RegisterResource("azure:servicefabric/meshSecret:MeshSecret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMeshSecret gets an existing MeshSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMeshSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MeshSecretState, opts ...pulumi.ResourceOption) (*MeshSecret, error) {
	var resource MeshSecret
	err := ctx.ReadResource("azure:servicefabric/meshSecret:MeshSecret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MeshSecret resources.
type meshSecretState struct {
	// The type of the content stored in the secret value. Changing this forces a new resource to be created.
	ContentType *string `pulumi:"contentType"`
	// A description of this Service Fabric Mesh Secret.
	Description *string `pulumi:"description"`
	// Specifies the Azure Region where the Service Fabric Mesh Secret should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the Service Fabric Mesh Secret exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type MeshSecretState struct {
	// The type of the content stored in the secret value. Changing this forces a new resource to be created.
	ContentType pulumi.StringPtrInput
	// A description of this Service Fabric Mesh Secret.
	Description pulumi.StringPtrInput
	// Specifies the Azure Region where the Service Fabric Mesh Secret should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the Service Fabric Mesh Secret exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MeshSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*meshSecretState)(nil)).Elem()
}

type meshSecretArgs struct {
	// The type of the content stored in the secret value. Changing this forces a new resource to be created.
	ContentType *string `pulumi:"contentType"`
	// A description of this Service Fabric Mesh Secret.
	Description *string `pulumi:"description"`
	// Specifies the Azure Region where the Service Fabric Mesh Secret should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the Service Fabric Mesh Secret exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MeshSecret resource.
type MeshSecretArgs struct {
	// The type of the content stored in the secret value. Changing this forces a new resource to be created.
	ContentType pulumi.StringPtrInput
	// A description of this Service Fabric Mesh Secret.
	Description pulumi.StringPtrInput
	// Specifies the Azure Region where the Service Fabric Mesh Secret should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the Service Fabric Mesh Secret exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MeshSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*meshSecretArgs)(nil)).Elem()
}

type MeshSecretInput interface {
	pulumi.Input

	ToMeshSecretOutput() MeshSecretOutput
	ToMeshSecretOutputWithContext(ctx context.Context) MeshSecretOutput
}

func (MeshSecret) ElementType() reflect.Type {
	return reflect.TypeOf((*MeshSecret)(nil)).Elem()
}

func (i MeshSecret) ToMeshSecretOutput() MeshSecretOutput {
	return i.ToMeshSecretOutputWithContext(context.Background())
}

func (i MeshSecret) ToMeshSecretOutputWithContext(ctx context.Context) MeshSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshSecretOutput)
}

type MeshSecretOutput struct {
	*pulumi.OutputState
}

func (MeshSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MeshSecretOutput)(nil)).Elem()
}

func (o MeshSecretOutput) ToMeshSecretOutput() MeshSecretOutput {
	return o
}

func (o MeshSecretOutput) ToMeshSecretOutputWithContext(ctx context.Context) MeshSecretOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MeshSecretOutput{})
}
