// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Windows Virtual Machine within a Dev Test Lab.
//
// ## Import
//
// DevTest Windows Virtual Machines can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:devtest/windowsVirtualMachine:WindowsVirtualMachine machine1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualmachines/machine1
// ```
type WindowsVirtualMachine struct {
	pulumi.CustomResourceState

	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim pulumi.BoolPtrOutput `pulumi:"allowClaim"`
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress pulumi.BoolPtrOutput `pulumi:"disallowPublicIpAddress"`
	// The FQDN of the Virtual Machine.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// A `galleryImageReference` block as defined below.
	GalleryImageReference WindowsVirtualMachineGalleryImageReferenceOutput `pulumi:"galleryImageReference"`
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules WindowsVirtualMachineInboundNatRuleArrayOutput `pulumi:"inboundNatRules"`
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringOutput `pulumi:"labName"`
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName pulumi.StringOutput `pulumi:"labSubnetName"`
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId pulumi.StringOutput `pulumi:"labVirtualNetworkId"`
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Any notes about the Virtual Machine.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password pulumi.StringOutput `pulumi:"password"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size pulumi.StringOutput `pulumi:"size"`
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewWindowsVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewWindowsVirtualMachine(ctx *pulumi.Context,
	name string, args *WindowsVirtualMachineArgs, opts ...pulumi.ResourceOption) (*WindowsVirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GalleryImageReference == nil {
		return nil, errors.New("invalid value for required argument 'GalleryImageReference'")
	}
	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.LabSubnetName == nil {
		return nil, errors.New("invalid value for required argument 'LabSubnetName'")
	}
	if args.LabVirtualNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'LabVirtualNetworkId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.StorageType == nil {
		return nil, errors.New("invalid value for required argument 'StorageType'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource WindowsVirtualMachine
	err := ctx.RegisterResource("azure:devtest/windowsVirtualMachine:WindowsVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWindowsVirtualMachine gets an existing WindowsVirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWindowsVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WindowsVirtualMachineState, opts ...pulumi.ResourceOption) (*WindowsVirtualMachine, error) {
	var resource WindowsVirtualMachine
	err := ctx.ReadResource("azure:devtest/windowsVirtualMachine:WindowsVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WindowsVirtualMachine resources.
type windowsVirtualMachineState struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim *bool `pulumi:"allowClaim"`
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress *bool `pulumi:"disallowPublicIpAddress"`
	// The FQDN of the Virtual Machine.
	Fqdn *string `pulumi:"fqdn"`
	// A `galleryImageReference` block as defined below.
	GalleryImageReference *WindowsVirtualMachineGalleryImageReference `pulumi:"galleryImageReference"`
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules []WindowsVirtualMachineInboundNatRule `pulumi:"inboundNatRules"`
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName *string `pulumi:"labName"`
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName *string `pulumi:"labSubnetName"`
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId *string `pulumi:"labVirtualNetworkId"`
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Any notes about the Virtual Machine.
	Notes *string `pulumi:"notes"`
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password *string `pulumi:"password"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size *string `pulumi:"size"`
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType *string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username *string `pulumi:"username"`
}

type WindowsVirtualMachineState struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim pulumi.BoolPtrInput
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress pulumi.BoolPtrInput
	// The FQDN of the Virtual Machine.
	Fqdn pulumi.StringPtrInput
	// A `galleryImageReference` block as defined below.
	GalleryImageReference WindowsVirtualMachineGalleryImageReferencePtrInput
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules WindowsVirtualMachineInboundNatRuleArrayInput
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringPtrInput
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName pulumi.StringPtrInput
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId pulumi.StringPtrInput
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Any notes about the Virtual Machine.
	Notes pulumi.StringPtrInput
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password pulumi.StringPtrInput
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size pulumi.StringPtrInput
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier pulumi.StringPtrInput
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username pulumi.StringPtrInput
}

func (WindowsVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*windowsVirtualMachineState)(nil)).Elem()
}

type windowsVirtualMachineArgs struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim *bool `pulumi:"allowClaim"`
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress *bool `pulumi:"disallowPublicIpAddress"`
	// A `galleryImageReference` block as defined below.
	GalleryImageReference WindowsVirtualMachineGalleryImageReference `pulumi:"galleryImageReference"`
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules []WindowsVirtualMachineInboundNatRule `pulumi:"inboundNatRules"`
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName string `pulumi:"labName"`
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName string `pulumi:"labSubnetName"`
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId string `pulumi:"labVirtualNetworkId"`
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Any notes about the Virtual Machine.
	Notes *string `pulumi:"notes"`
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password string `pulumi:"password"`
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size string `pulumi:"size"`
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType string `pulumi:"storageType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a WindowsVirtualMachine resource.
type WindowsVirtualMachineArgs struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim pulumi.BoolPtrInput
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress pulumi.BoolPtrInput
	// A `galleryImageReference` block as defined below.
	GalleryImageReference WindowsVirtualMachineGalleryImageReferenceInput
	// One or more `inboundNatRule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules WindowsVirtualMachineInboundNatRuleArrayInput
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName pulumi.StringInput
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName pulumi.StringInput
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId pulumi.StringInput
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Any notes about the Virtual Machine.
	Notes pulumi.StringPtrInput
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password pulumi.StringInput
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size pulumi.StringInput
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username pulumi.StringInput
}

func (WindowsVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*windowsVirtualMachineArgs)(nil)).Elem()
}

type WindowsVirtualMachineInput interface {
	pulumi.Input

	ToWindowsVirtualMachineOutput() WindowsVirtualMachineOutput
	ToWindowsVirtualMachineOutputWithContext(ctx context.Context) WindowsVirtualMachineOutput
}

func (*WindowsVirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsVirtualMachine)(nil))
}

func (i *WindowsVirtualMachine) ToWindowsVirtualMachineOutput() WindowsVirtualMachineOutput {
	return i.ToWindowsVirtualMachineOutputWithContext(context.Background())
}

func (i *WindowsVirtualMachine) ToWindowsVirtualMachineOutputWithContext(ctx context.Context) WindowsVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsVirtualMachineOutput)
}

func (i *WindowsVirtualMachine) ToWindowsVirtualMachinePtrOutput() WindowsVirtualMachinePtrOutput {
	return i.ToWindowsVirtualMachinePtrOutputWithContext(context.Background())
}

func (i *WindowsVirtualMachine) ToWindowsVirtualMachinePtrOutputWithContext(ctx context.Context) WindowsVirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsVirtualMachinePtrOutput)
}

type WindowsVirtualMachinePtrInput interface {
	pulumi.Input

	ToWindowsVirtualMachinePtrOutput() WindowsVirtualMachinePtrOutput
	ToWindowsVirtualMachinePtrOutputWithContext(ctx context.Context) WindowsVirtualMachinePtrOutput
}

type windowsVirtualMachinePtrType WindowsVirtualMachineArgs

func (*windowsVirtualMachinePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsVirtualMachine)(nil))
}

func (i *windowsVirtualMachinePtrType) ToWindowsVirtualMachinePtrOutput() WindowsVirtualMachinePtrOutput {
	return i.ToWindowsVirtualMachinePtrOutputWithContext(context.Background())
}

func (i *windowsVirtualMachinePtrType) ToWindowsVirtualMachinePtrOutputWithContext(ctx context.Context) WindowsVirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsVirtualMachineOutput).ToWindowsVirtualMachinePtrOutput()
}

// WindowsVirtualMachineArrayInput is an input type that accepts WindowsVirtualMachineArray and WindowsVirtualMachineArrayOutput values.
// You can construct a concrete instance of `WindowsVirtualMachineArrayInput` via:
//
//          WindowsVirtualMachineArray{ WindowsVirtualMachineArgs{...} }
type WindowsVirtualMachineArrayInput interface {
	pulumi.Input

	ToWindowsVirtualMachineArrayOutput() WindowsVirtualMachineArrayOutput
	ToWindowsVirtualMachineArrayOutputWithContext(context.Context) WindowsVirtualMachineArrayOutput
}

type WindowsVirtualMachineArray []WindowsVirtualMachineInput

func (WindowsVirtualMachineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsVirtualMachine)(nil))
}

func (i WindowsVirtualMachineArray) ToWindowsVirtualMachineArrayOutput() WindowsVirtualMachineArrayOutput {
	return i.ToWindowsVirtualMachineArrayOutputWithContext(context.Background())
}

func (i WindowsVirtualMachineArray) ToWindowsVirtualMachineArrayOutputWithContext(ctx context.Context) WindowsVirtualMachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsVirtualMachineArrayOutput)
}

// WindowsVirtualMachineMapInput is an input type that accepts WindowsVirtualMachineMap and WindowsVirtualMachineMapOutput values.
// You can construct a concrete instance of `WindowsVirtualMachineMapInput` via:
//
//          WindowsVirtualMachineMap{ "key": WindowsVirtualMachineArgs{...} }
type WindowsVirtualMachineMapInput interface {
	pulumi.Input

	ToWindowsVirtualMachineMapOutput() WindowsVirtualMachineMapOutput
	ToWindowsVirtualMachineMapOutputWithContext(context.Context) WindowsVirtualMachineMapOutput
}

type WindowsVirtualMachineMap map[string]WindowsVirtualMachineInput

func (WindowsVirtualMachineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WindowsVirtualMachine)(nil))
}

func (i WindowsVirtualMachineMap) ToWindowsVirtualMachineMapOutput() WindowsVirtualMachineMapOutput {
	return i.ToWindowsVirtualMachineMapOutputWithContext(context.Background())
}

func (i WindowsVirtualMachineMap) ToWindowsVirtualMachineMapOutputWithContext(ctx context.Context) WindowsVirtualMachineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsVirtualMachineMapOutput)
}

type WindowsVirtualMachineOutput struct {
	*pulumi.OutputState
}

func (WindowsVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsVirtualMachine)(nil))
}

func (o WindowsVirtualMachineOutput) ToWindowsVirtualMachineOutput() WindowsVirtualMachineOutput {
	return o
}

func (o WindowsVirtualMachineOutput) ToWindowsVirtualMachineOutputWithContext(ctx context.Context) WindowsVirtualMachineOutput {
	return o
}

func (o WindowsVirtualMachineOutput) ToWindowsVirtualMachinePtrOutput() WindowsVirtualMachinePtrOutput {
	return o.ToWindowsVirtualMachinePtrOutputWithContext(context.Background())
}

func (o WindowsVirtualMachineOutput) ToWindowsVirtualMachinePtrOutputWithContext(ctx context.Context) WindowsVirtualMachinePtrOutput {
	return o.ApplyT(func(v WindowsVirtualMachine) *WindowsVirtualMachine {
		return &v
	}).(WindowsVirtualMachinePtrOutput)
}

type WindowsVirtualMachinePtrOutput struct {
	*pulumi.OutputState
}

func (WindowsVirtualMachinePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsVirtualMachine)(nil))
}

func (o WindowsVirtualMachinePtrOutput) ToWindowsVirtualMachinePtrOutput() WindowsVirtualMachinePtrOutput {
	return o
}

func (o WindowsVirtualMachinePtrOutput) ToWindowsVirtualMachinePtrOutputWithContext(ctx context.Context) WindowsVirtualMachinePtrOutput {
	return o
}

type WindowsVirtualMachineArrayOutput struct{ *pulumi.OutputState }

func (WindowsVirtualMachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsVirtualMachine)(nil))
}

func (o WindowsVirtualMachineArrayOutput) ToWindowsVirtualMachineArrayOutput() WindowsVirtualMachineArrayOutput {
	return o
}

func (o WindowsVirtualMachineArrayOutput) ToWindowsVirtualMachineArrayOutputWithContext(ctx context.Context) WindowsVirtualMachineArrayOutput {
	return o
}

func (o WindowsVirtualMachineArrayOutput) Index(i pulumi.IntInput) WindowsVirtualMachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WindowsVirtualMachine {
		return vs[0].([]WindowsVirtualMachine)[vs[1].(int)]
	}).(WindowsVirtualMachineOutput)
}

type WindowsVirtualMachineMapOutput struct{ *pulumi.OutputState }

func (WindowsVirtualMachineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WindowsVirtualMachine)(nil))
}

func (o WindowsVirtualMachineMapOutput) ToWindowsVirtualMachineMapOutput() WindowsVirtualMachineMapOutput {
	return o
}

func (o WindowsVirtualMachineMapOutput) ToWindowsVirtualMachineMapOutputWithContext(ctx context.Context) WindowsVirtualMachineMapOutput {
	return o
}

func (o WindowsVirtualMachineMapOutput) MapIndex(k pulumi.StringInput) WindowsVirtualMachineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WindowsVirtualMachine {
		return vs[0].(map[string]WindowsVirtualMachine)[vs[1].(string)]
	}).(WindowsVirtualMachineOutput)
}

func init() {
	pulumi.RegisterOutputType(WindowsVirtualMachineOutput{})
	pulumi.RegisterOutputType(WindowsVirtualMachinePtrOutput{})
	pulumi.RegisterOutputType(WindowsVirtualMachineArrayOutput{})
	pulumi.RegisterOutputType(WindowsVirtualMachineMapOutput{})
}
