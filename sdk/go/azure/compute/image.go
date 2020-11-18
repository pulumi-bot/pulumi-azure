// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a custom virtual machine image that can be used to create virtual machines.
//
// ## Example Usage
// ### Creating From VHD
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewImage(ctx, "exampleImage", &compute.ImageArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			OsDisk: &compute.ImageOsDiskArgs{
// 				OsType:  pulumi.String("Linux"),
// 				OsState: pulumi.String("Generalized"),
// 				BlobUri: pulumi.String("{blob_uri}"),
// 				SizeGb:  pulumi.Int(30),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Creating From Virtual Machine (VM Must Be Generalized Beforehand)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewImage(ctx, "exampleImage", &compute.ImageArgs{
// 			Location:               pulumi.String("West US"),
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			SourceVirtualMachineId: pulumi.String("{vm_id}"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Image struct {
	pulumi.CustomResourceState

	// One or more `dataDisk` elements as defined below.
	DataDisks ImageDataDiskArrayOutput `pulumi:"dataDisks"`
	// The HyperVGenerationType of the VirtualMachine created from the image as `V1`, `V2`. The default is `V1`.
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `osDisk` elements as defined below.
	OsDisk ImageOsDiskPtrOutput `pulumi:"osDisk"`
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId pulumi.StringPtrOutput `pulumi:"sourceVirtualMachineId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient pulumi.BoolPtrOutput `pulumi:"zoneResilient"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Image
	err := ctx.RegisterResource("azure:compute/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("azure:compute/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// One or more `dataDisk` elements as defined below.
	DataDisks []ImageDataDisk `pulumi:"dataDisks"`
	// The HyperVGenerationType of the VirtualMachine created from the image as `V1`, `V2`. The default is `V1`.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `osDisk` elements as defined below.
	OsDisk *ImageOsDisk `pulumi:"osDisk"`
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId *string `pulumi:"sourceVirtualMachineId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient *bool `pulumi:"zoneResilient"`
}

type ImageState struct {
	// One or more `dataDisk` elements as defined below.
	DataDisks ImageDataDiskArrayInput
	// The HyperVGenerationType of the VirtualMachine created from the image as `V1`, `V2`. The default is `V1`.
	HyperVGeneration pulumi.StringPtrInput
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `osDisk` elements as defined below.
	OsDisk ImageOsDiskPtrInput
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient pulumi.BoolPtrInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// One or more `dataDisk` elements as defined below.
	DataDisks []ImageDataDisk `pulumi:"dataDisks"`
	// The HyperVGenerationType of the VirtualMachine created from the image as `V1`, `V2`. The default is `V1`.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `osDisk` elements as defined below.
	OsDisk *ImageOsDisk `pulumi:"osDisk"`
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId *string `pulumi:"sourceVirtualMachineId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient *bool `pulumi:"zoneResilient"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// One or more `dataDisk` elements as defined below.
	DataDisks ImageDataDiskArrayInput
	// The HyperVGenerationType of the VirtualMachine created from the image as `V1`, `V2`. The default is `V1`.
	HyperVGeneration pulumi.StringPtrInput
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `osDisk` elements as defined below.
	OsDisk ImageOsDiskPtrInput
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient pulumi.BoolPtrInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (Image) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil)).Elem()
}

func (i Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct {
	*pulumi.OutputState
}

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageOutput)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ImageOutput{})
}
