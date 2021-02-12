// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages as an Azure Container Group instance.
//
// ## Example Usage
//
// This example provisions a Basic Container.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/containerservice"
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
// 		_, err = containerservice.NewGroup(ctx, "exampleGroup", &containerservice.GroupArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpAddressType:     pulumi.String("public"),
// 			DnsNameLabel:      pulumi.String("aci-label"),
// 			OsType:            pulumi.String("Linux"),
// 			Containers: containerservice.GroupContainerArray{
// 				&containerservice.GroupContainerArgs{
// 					Name:   pulumi.String("hello-world"),
// 					Image:  pulumi.String("microsoft/aci-helloworld:latest"),
// 					Cpu:    pulumi.Float64(0.5),
// 					Memory: pulumi.Float64(1.5),
// 					Ports: containerservice.GroupContainerPortArray{
// 						&containerservice.GroupContainerPortArgs{
// 							Port:     pulumi.Int(443),
// 							Protocol: pulumi.String("TCP"),
// 						},
// 					},
// 				},
// 				&containerservice.GroupContainerArgs{
// 					Name:   pulumi.String("sidecar"),
// 					Image:  pulumi.String("microsoft/aci-tutorial-sidecar"),
// 					Cpu:    pulumi.Float64(0.5),
// 					Memory: pulumi.Float64(1.5),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("testing"),
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
// Container Group's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:containerservice/group:Group containerGroup1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ContainerInstance/containerGroups/myContainerGroup1
// ```
type Group struct {
	pulumi.CustomResourceState

	// The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
	Containers GroupContainerArrayOutput `pulumi:"containers"`
	// A `diagnostics` block as documented below.
	Diagnostics GroupDiagnosticsPtrOutput `pulumi:"diagnostics"`
	// A `dnsConfig` block as documented below.
	DnsConfig GroupDnsConfigPtrOutput `pulumi:"dnsConfig"`
	// The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
	DnsNameLabel pulumi.StringPtrOutput `pulumi:"dnsNameLabel"`
	// The FQDN of the container group derived from `dnsNameLabel`.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// An `identity` block as defined below.
	Identity GroupIdentityOutput `pulumi:"identity"`
	// A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
	ImageRegistryCredentials GroupImageRegistryCredentialArrayOutput `pulumi:"imageRegistryCredentials"`
	// The IP address allocated to the container group.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
	IpAddressType pulumi.StringPtrOutput `pulumi:"ipAddressType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Container Group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network profile ID for deploying to virtual network.
	NetworkProfileId pulumi.StringPtrOutput `pulumi:"networkProfileId"`
	// The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
	RestartPolicy pulumi.StringPtrOutput `pulumi:"restartPolicy"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Containers == nil {
		return nil, errors.New("invalid value for required argument 'Containers'")
	}
	if args.OsType == nil {
		return nil, errors.New("invalid value for required argument 'OsType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Group
	err := ctx.RegisterResource("azure:containerservice/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azure:containerservice/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
	Containers []GroupContainer `pulumi:"containers"`
	// A `diagnostics` block as documented below.
	Diagnostics *GroupDiagnostics `pulumi:"diagnostics"`
	// A `dnsConfig` block as documented below.
	DnsConfig *GroupDnsConfig `pulumi:"dnsConfig"`
	// The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
	DnsNameLabel *string `pulumi:"dnsNameLabel"`
	// The FQDN of the container group derived from `dnsNameLabel`.
	Fqdn *string `pulumi:"fqdn"`
	// An `identity` block as defined below.
	Identity *GroupIdentity `pulumi:"identity"`
	// A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
	ImageRegistryCredentials []GroupImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// The IP address allocated to the container group.
	IpAddress *string `pulumi:"ipAddress"`
	// Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
	IpAddressType *string `pulumi:"ipAddressType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Container Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Network profile ID for deploying to virtual network.
	NetworkProfileId *string `pulumi:"networkProfileId"`
	// The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType *string `pulumi:"osType"`
	// The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
	RestartPolicy *string `pulumi:"restartPolicy"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type GroupState struct {
	// The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
	Containers GroupContainerArrayInput
	// A `diagnostics` block as documented below.
	Diagnostics GroupDiagnosticsPtrInput
	// A `dnsConfig` block as documented below.
	DnsConfig GroupDnsConfigPtrInput
	// The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
	DnsNameLabel pulumi.StringPtrInput
	// The FQDN of the container group derived from `dnsNameLabel`.
	Fqdn pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity GroupIdentityPtrInput
	// A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
	ImageRegistryCredentials GroupImageRegistryCredentialArrayInput
	// The IP address allocated to the container group.
	IpAddress pulumi.StringPtrInput
	// Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
	IpAddressType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Container Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Network profile ID for deploying to virtual network.
	NetworkProfileId pulumi.StringPtrInput
	// The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType pulumi.StringPtrInput
	// The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
	RestartPolicy pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
	Containers []GroupContainer `pulumi:"containers"`
	// A `diagnostics` block as documented below.
	Diagnostics *GroupDiagnostics `pulumi:"diagnostics"`
	// A `dnsConfig` block as documented below.
	DnsConfig *GroupDnsConfig `pulumi:"dnsConfig"`
	// The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
	DnsNameLabel *string `pulumi:"dnsNameLabel"`
	// An `identity` block as defined below.
	Identity *GroupIdentity `pulumi:"identity"`
	// A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
	ImageRegistryCredentials []GroupImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
	IpAddressType *string `pulumi:"ipAddressType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Container Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Network profile ID for deploying to virtual network.
	NetworkProfileId *string `pulumi:"networkProfileId"`
	// The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType string `pulumi:"osType"`
	// The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
	RestartPolicy *string `pulumi:"restartPolicy"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
	Containers GroupContainerArrayInput
	// A `diagnostics` block as documented below.
	Diagnostics GroupDiagnosticsPtrInput
	// A `dnsConfig` block as documented below.
	DnsConfig GroupDnsConfigPtrInput
	// The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
	DnsNameLabel pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity GroupIdentityPtrInput
	// A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
	ImageRegistryCredentials GroupImageRegistryCredentialArrayInput
	// Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
	IpAddressType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Container Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Network profile ID for deploying to virtual network.
	NetworkProfileId pulumi.StringPtrInput
	// The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType pulumi.StringInput
	// The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
	RestartPolicy pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct {
	*pulumi.OutputState
}

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}
