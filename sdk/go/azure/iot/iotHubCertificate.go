// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an IotHub Device Provisioning Service Certificate.
//
// ## Import
//
// IoTHub Device Provisioning Service Certificates can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/iotHubCertificate:IotHubCertificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/example/certificates/example
// ```
type IotHubCertificate struct {
	pulumi.CustomResourceState

	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent pulumi.StringOutput `pulumi:"certificateContent"`
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName pulumi.StringOutput `pulumi:"iotDpsName"`
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIotHubCertificate registers a new resource with the given unique name, arguments, and options.
func NewIotHubCertificate(ctx *pulumi.Context,
	name string, args *IotHubCertificateArgs, opts ...pulumi.ResourceOption) (*IotHubCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateContent == nil {
		return nil, errors.New("invalid value for required argument 'CertificateContent'")
	}
	if args.IotDpsName == nil {
		return nil, errors.New("invalid value for required argument 'IotDpsName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IotHubCertificate
	err := ctx.RegisterResource("azure:iot/iotHubCertificate:IotHubCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotHubCertificate gets an existing IotHubCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotHubCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubCertificateState, opts ...pulumi.ResourceOption) (*IotHubCertificate, error) {
	var resource IotHubCertificate
	err := ctx.ReadResource("azure:iot/iotHubCertificate:IotHubCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotHubCertificate resources.
type iotHubCertificateState struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent *string `pulumi:"certificateContent"`
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName *string `pulumi:"iotDpsName"`
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IotHubCertificateState struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent pulumi.StringPtrInput
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName pulumi.StringPtrInput
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IotHubCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubCertificateState)(nil)).Elem()
}

type iotHubCertificateArgs struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent string `pulumi:"certificateContent"`
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName string `pulumi:"iotDpsName"`
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IotHubCertificate resource.
type IotHubCertificateArgs struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent pulumi.StringInput
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName pulumi.StringInput
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (IotHubCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubCertificateArgs)(nil)).Elem()
}

type IotHubCertificateInput interface {
	pulumi.Input

	ToIotHubCertificateOutput() IotHubCertificateOutput
	ToIotHubCertificateOutputWithContext(ctx context.Context) IotHubCertificateOutput
}

func (*IotHubCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubCertificate)(nil))
}

func (i *IotHubCertificate) ToIotHubCertificateOutput() IotHubCertificateOutput {
	return i.ToIotHubCertificateOutputWithContext(context.Background())
}

func (i *IotHubCertificate) ToIotHubCertificateOutputWithContext(ctx context.Context) IotHubCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubCertificateOutput)
}

type IotHubCertificateOutput struct {
	*pulumi.OutputState
}

func (IotHubCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubCertificate)(nil))
}

func (o IotHubCertificateOutput) ToIotHubCertificateOutput() IotHubCertificateOutput {
	return o
}

func (o IotHubCertificateOutput) ToIotHubCertificateOutputWithContext(ctx context.Context) IotHubCertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IotHubCertificateOutput{})
}
