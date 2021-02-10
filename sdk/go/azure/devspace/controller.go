// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devspace

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// DevSpace Controller's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:devspace/controller:Controller controller1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevSpaces/controllers/controller1Name
// ```
type Controller struct {
	pulumi.CustomResourceState

	// DNS name for accessing DataPlane services.
	DataPlaneFqdn pulumi.StringOutput `pulumi:"dataPlaneFqdn"`
	// The host suffix for the DevSpace Controller.
	HostSuffix pulumi.StringOutput `pulumi:"hostSuffix"`
	// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Base64 encoding of `kubeConfigRaw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostCredentialsBase64 pulumi.StringOutput `pulumi:"targetContainerHostCredentialsBase64"`
	// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostResourceId pulumi.StringOutput `pulumi:"targetContainerHostResourceId"`
}

// NewController registers a new resource with the given unique name, arguments, and options.
func NewController(ctx *pulumi.Context,
	name string, args *ControllerArgs, opts ...pulumi.ResourceOption) (*Controller, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	if args.TargetContainerHostCredentialsBase64 == nil {
		return nil, errors.New("invalid value for required argument 'TargetContainerHostCredentialsBase64'")
	}
	if args.TargetContainerHostResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetContainerHostResourceId'")
	}
	var resource Controller
	err := ctx.RegisterResource("azure:devspace/controller:Controller", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetController gets an existing Controller resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControllerState, opts ...pulumi.ResourceOption) (*Controller, error) {
	var resource Controller
	err := ctx.ReadResource("azure:devspace/controller:Controller", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Controller resources.
type controllerState struct {
	// DNS name for accessing DataPlane services.
	DataPlaneFqdn *string `pulumi:"dataPlaneFqdn"`
	// The host suffix for the DevSpace Controller.
	HostSuffix *string `pulumi:"hostSuffix"`
	// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Base64 encoding of `kubeConfigRaw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostCredentialsBase64 *string `pulumi:"targetContainerHostCredentialsBase64"`
	// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostResourceId *string `pulumi:"targetContainerHostResourceId"`
}

type ControllerState struct {
	// DNS name for accessing DataPlane services.
	DataPlaneFqdn pulumi.StringPtrInput
	// The host suffix for the DevSpace Controller.
	HostSuffix pulumi.StringPtrInput
	// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
	SkuName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Base64 encoding of `kubeConfigRaw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostCredentialsBase64 pulumi.StringPtrInput
	// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostResourceId pulumi.StringPtrInput
}

func (ControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerState)(nil)).Elem()
}

type controllerArgs struct {
	// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Base64 encoding of `kubeConfigRaw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostCredentialsBase64 string `pulumi:"targetContainerHostCredentialsBase64"`
	// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostResourceId string `pulumi:"targetContainerHostResourceId"`
}

// The set of arguments for constructing a Controller resource.
type ControllerArgs struct {
	// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the SKU Name for this DevSpace Controller. Possible values are `S1`.
	SkuName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Base64 encoding of `kubeConfigRaw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostCredentialsBase64 pulumi.StringInput
	// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostResourceId pulumi.StringInput
}

func (ControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controllerArgs)(nil)).Elem()
}

type ControllerInput interface {
	pulumi.Input

	ToControllerOutput() ControllerOutput
	ToControllerOutputWithContext(ctx context.Context) ControllerOutput
}

func (*Controller) ElementType() reflect.Type {
	return reflect.TypeOf((*Controller)(nil))
}

func (i *Controller) ToControllerOutput() ControllerOutput {
	return i.ToControllerOutputWithContext(context.Background())
}

func (i *Controller) ToControllerOutputWithContext(ctx context.Context) ControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControllerOutput)
}

type ControllerOutput struct {
	*pulumi.OutputState
}

func (ControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Controller)(nil))
}

func (o ControllerOutput) ToControllerOutput() ControllerOutput {
	return o
}

func (o ControllerOutput) ToControllerOutputWithContext(ctx context.Context) ControllerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ControllerOutput{})
}
