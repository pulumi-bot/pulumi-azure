// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages automated shutdown schedules for Azure VMs that are not within an Azure DevTest Lab. While this is part of the DevTest Labs service in Azure,
// this resource applies only to standard VMs, not DevTest Lab VMs. To manage automated shutdown schedules for DevTest Lab VMs, reference the
// `devtest.Schedule` resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/devtest"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("eastus"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNetworkInterface, err := network.NewNetworkInterface(ctx, "exampleNetworkInterface", &network.NetworkInterfaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpConfigurations: network.NetworkInterfaceIpConfigurationArray{
// 				&network.NetworkInterfaceIpConfigurationArgs{
// 					Name:                       pulumi.String("testconfiguration1"),
// 					SubnetId:                   exampleSubnet.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewLinuxVirtualMachine(ctx, "exampleLinuxVirtualMachine", &compute.LinuxVirtualMachineArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			NetworkInterfaceIds: pulumi.StringArray{
// 				exampleNetworkInterface.ID(),
// 			},
// 			Size: pulumi.String("Standard_B2s"),
// 			SourceImageReference: &compute.LinuxVirtualMachineSourceImageReferenceArgs{
// 				Publisher: pulumi.String("Canonical"),
// 				Offer:     pulumi.String("UbuntuServer"),
// 				Sku:       pulumi.String("16.04-LTS"),
// 				Version:   pulumi.String("latest"),
// 			},
// 			OsDisk: &compute.LinuxVirtualMachineOsDiskArgs{
// 				Name:            pulumi.String(fmt.Sprintf("%v%v%v", "myosdisk-", "%", "d")),
// 				Caching:         pulumi.String("ReadWrite"),
// 				ManagedDiskType: pulumi.String("Standard_LRS"),
// 			},
// 			AdminUsername:                 pulumi.String("testadmin"),
// 			AdminPassword:                 pulumi.String("Password1234!"),
// 			DisablePasswordAuthentication: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = devtest.NewGlobalVMShutdownSchedule(ctx, "exampleGlobalVMShutdownSchedule", &devtest.GlobalVMShutdownScheduleArgs{
// 			VirtualMachineId:    pulumi.Any(azurerm_virtual_machine.Example.Id),
// 			Location:            exampleResourceGroup.Location,
// 			Enabled:             pulumi.Bool(true),
// 			DailyRecurrenceTime: pulumi.String("1100"),
// 			Timezone:            pulumi.String("Pacific Standard Time"),
// 			NotificationSettings: &devtest.GlobalVMShutdownScheduleNotificationSettingsArgs{
// 				Enabled:       pulumi.Bool(true),
// 				TimeInMinutes: pulumi.Int(60),
// 				WebhookUrl:    pulumi.String("https://sample-webhook-url.example.com"),
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
// An existing Dev Test Global Shutdown Schedule can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:devtest/globalVMShutdownSchedule:GlobalVMShutdownSchedule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-rg/providers/Microsoft.DevTestLab/schedules/shutdown-computevm-SampleVM
// ```
//
//  The name of the resource within the `resource id` will always follow the format `shutdown-computevm-<VM Name>` where `<VM Name>` is replaced by the name of the target Virtual Machine
type GlobalVMShutdownSchedule struct {
	pulumi.CustomResourceState

	// The time each day when the schedule takes effect. Must match the format HHmm where HH is 00-23 and mm is 00-59 (e.g. 0930, 2300, etc.)
	DailyRecurrenceTime pulumi.StringOutput `pulumi:"dailyRecurrenceTime"`
	// Whether to enable the schedule. Possible values are `true` and `false`. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The location where the schedule is created. Changing this forces a new resource to be created.
	Location             pulumi.StringOutput                                `pulumi:"location"`
	NotificationSettings GlobalVMShutdownScheduleNotificationSettingsOutput `pulumi:"notificationSettings"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The time zone ID (e.g. Pacific Standard time). Refer to this guide for a [full list of accepted time zone names](https://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/).
	Timezone pulumi.StringOutput `pulumi:"timezone"`
	// The resource ID of the target ARM-based Virtual Machine. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
}

// NewGlobalVMShutdownSchedule registers a new resource with the given unique name, arguments, and options.
func NewGlobalVMShutdownSchedule(ctx *pulumi.Context,
	name string, args *GlobalVMShutdownScheduleArgs, opts ...pulumi.ResourceOption) (*GlobalVMShutdownSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DailyRecurrenceTime == nil {
		return nil, errors.New("invalid value for required argument 'DailyRecurrenceTime'")
	}
	if args.NotificationSettings == nil {
		return nil, errors.New("invalid value for required argument 'NotificationSettings'")
	}
	if args.Timezone == nil {
		return nil, errors.New("invalid value for required argument 'Timezone'")
	}
	if args.VirtualMachineId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineId'")
	}
	var resource GlobalVMShutdownSchedule
	err := ctx.RegisterResource("azure:devtest/globalVMShutdownSchedule:GlobalVMShutdownSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalVMShutdownSchedule gets an existing GlobalVMShutdownSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalVMShutdownSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalVMShutdownScheduleState, opts ...pulumi.ResourceOption) (*GlobalVMShutdownSchedule, error) {
	var resource GlobalVMShutdownSchedule
	err := ctx.ReadResource("azure:devtest/globalVMShutdownSchedule:GlobalVMShutdownSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalVMShutdownSchedule resources.
type globalVMShutdownScheduleState struct {
	// The time each day when the schedule takes effect. Must match the format HHmm where HH is 00-23 and mm is 00-59 (e.g. 0930, 2300, etc.)
	DailyRecurrenceTime *string `pulumi:"dailyRecurrenceTime"`
	// Whether to enable the schedule. Possible values are `true` and `false`. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The location where the schedule is created. Changing this forces a new resource to be created.
	Location             *string                                       `pulumi:"location"`
	NotificationSettings *GlobalVMShutdownScheduleNotificationSettings `pulumi:"notificationSettings"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The time zone ID (e.g. Pacific Standard time). Refer to this guide for a [full list of accepted time zone names](https://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/).
	Timezone *string `pulumi:"timezone"`
	// The resource ID of the target ARM-based Virtual Machine. Changing this forces a new resource to be created.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type GlobalVMShutdownScheduleState struct {
	// The time each day when the schedule takes effect. Must match the format HHmm where HH is 00-23 and mm is 00-59 (e.g. 0930, 2300, etc.)
	DailyRecurrenceTime pulumi.StringPtrInput
	// Whether to enable the schedule. Possible values are `true` and `false`. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The location where the schedule is created. Changing this forces a new resource to be created.
	Location             pulumi.StringPtrInput
	NotificationSettings GlobalVMShutdownScheduleNotificationSettingsPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The time zone ID (e.g. Pacific Standard time). Refer to this guide for a [full list of accepted time zone names](https://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/).
	Timezone pulumi.StringPtrInput
	// The resource ID of the target ARM-based Virtual Machine. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringPtrInput
}

func (GlobalVMShutdownScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalVMShutdownScheduleState)(nil)).Elem()
}

type globalVMShutdownScheduleArgs struct {
	// The time each day when the schedule takes effect. Must match the format HHmm where HH is 00-23 and mm is 00-59 (e.g. 0930, 2300, etc.)
	DailyRecurrenceTime string `pulumi:"dailyRecurrenceTime"`
	// Whether to enable the schedule. Possible values are `true` and `false`. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The location where the schedule is created. Changing this forces a new resource to be created.
	Location             *string                                      `pulumi:"location"`
	NotificationSettings GlobalVMShutdownScheduleNotificationSettings `pulumi:"notificationSettings"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The time zone ID (e.g. Pacific Standard time). Refer to this guide for a [full list of accepted time zone names](https://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/).
	Timezone string `pulumi:"timezone"`
	// The resource ID of the target ARM-based Virtual Machine. Changing this forces a new resource to be created.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a GlobalVMShutdownSchedule resource.
type GlobalVMShutdownScheduleArgs struct {
	// The time each day when the schedule takes effect. Must match the format HHmm where HH is 00-23 and mm is 00-59 (e.g. 0930, 2300, etc.)
	DailyRecurrenceTime pulumi.StringInput
	// Whether to enable the schedule. Possible values are `true` and `false`. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The location where the schedule is created. Changing this forces a new resource to be created.
	Location             pulumi.StringPtrInput
	NotificationSettings GlobalVMShutdownScheduleNotificationSettingsInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The time zone ID (e.g. Pacific Standard time). Refer to this guide for a [full list of accepted time zone names](https://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/).
	Timezone pulumi.StringInput
	// The resource ID of the target ARM-based Virtual Machine. Changing this forces a new resource to be created.
	VirtualMachineId pulumi.StringInput
}

func (GlobalVMShutdownScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalVMShutdownScheduleArgs)(nil)).Elem()
}

type GlobalVMShutdownScheduleInput interface {
	pulumi.Input

	ToGlobalVMShutdownScheduleOutput() GlobalVMShutdownScheduleOutput
	ToGlobalVMShutdownScheduleOutputWithContext(ctx context.Context) GlobalVMShutdownScheduleOutput
}

func (*GlobalVMShutdownSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalVMShutdownSchedule)(nil))
}

func (i *GlobalVMShutdownSchedule) ToGlobalVMShutdownScheduleOutput() GlobalVMShutdownScheduleOutput {
	return i.ToGlobalVMShutdownScheduleOutputWithContext(context.Background())
}

func (i *GlobalVMShutdownSchedule) ToGlobalVMShutdownScheduleOutputWithContext(ctx context.Context) GlobalVMShutdownScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalVMShutdownScheduleOutput)
}

func (i *GlobalVMShutdownSchedule) ToGlobalVMShutdownSchedulePtrOutput() GlobalVMShutdownSchedulePtrOutput {
	return i.ToGlobalVMShutdownSchedulePtrOutputWithContext(context.Background())
}

func (i *GlobalVMShutdownSchedule) ToGlobalVMShutdownSchedulePtrOutputWithContext(ctx context.Context) GlobalVMShutdownSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalVMShutdownSchedulePtrOutput)
}

type GlobalVMShutdownSchedulePtrInput interface {
	pulumi.Input

	ToGlobalVMShutdownSchedulePtrOutput() GlobalVMShutdownSchedulePtrOutput
	ToGlobalVMShutdownSchedulePtrOutputWithContext(ctx context.Context) GlobalVMShutdownSchedulePtrOutput
}

type globalVMShutdownSchedulePtrType GlobalVMShutdownScheduleArgs

func (*globalVMShutdownSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalVMShutdownSchedule)(nil))
}

func (i *globalVMShutdownSchedulePtrType) ToGlobalVMShutdownSchedulePtrOutput() GlobalVMShutdownSchedulePtrOutput {
	return i.ToGlobalVMShutdownSchedulePtrOutputWithContext(context.Background())
}

func (i *globalVMShutdownSchedulePtrType) ToGlobalVMShutdownSchedulePtrOutputWithContext(ctx context.Context) GlobalVMShutdownSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalVMShutdownScheduleOutput).ToGlobalVMShutdownSchedulePtrOutput()
}

// GlobalVMShutdownScheduleArrayInput is an input type that accepts GlobalVMShutdownScheduleArray and GlobalVMShutdownScheduleArrayOutput values.
// You can construct a concrete instance of `GlobalVMShutdownScheduleArrayInput` via:
//
//          GlobalVMShutdownScheduleArray{ GlobalVMShutdownScheduleArgs{...} }
type GlobalVMShutdownScheduleArrayInput interface {
	pulumi.Input

	ToGlobalVMShutdownScheduleArrayOutput() GlobalVMShutdownScheduleArrayOutput
	ToGlobalVMShutdownScheduleArrayOutputWithContext(context.Context) GlobalVMShutdownScheduleArrayOutput
}

type GlobalVMShutdownScheduleArray []GlobalVMShutdownScheduleInput

func (GlobalVMShutdownScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalVMShutdownSchedule)(nil))
}

func (i GlobalVMShutdownScheduleArray) ToGlobalVMShutdownScheduleArrayOutput() GlobalVMShutdownScheduleArrayOutput {
	return i.ToGlobalVMShutdownScheduleArrayOutputWithContext(context.Background())
}

func (i GlobalVMShutdownScheduleArray) ToGlobalVMShutdownScheduleArrayOutputWithContext(ctx context.Context) GlobalVMShutdownScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalVMShutdownScheduleArrayOutput)
}

// GlobalVMShutdownScheduleMapInput is an input type that accepts GlobalVMShutdownScheduleMap and GlobalVMShutdownScheduleMapOutput values.
// You can construct a concrete instance of `GlobalVMShutdownScheduleMapInput` via:
//
//          GlobalVMShutdownScheduleMap{ "key": GlobalVMShutdownScheduleArgs{...} }
type GlobalVMShutdownScheduleMapInput interface {
	pulumi.Input

	ToGlobalVMShutdownScheduleMapOutput() GlobalVMShutdownScheduleMapOutput
	ToGlobalVMShutdownScheduleMapOutputWithContext(context.Context) GlobalVMShutdownScheduleMapOutput
}

type GlobalVMShutdownScheduleMap map[string]GlobalVMShutdownScheduleInput

func (GlobalVMShutdownScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GlobalVMShutdownSchedule)(nil))
}

func (i GlobalVMShutdownScheduleMap) ToGlobalVMShutdownScheduleMapOutput() GlobalVMShutdownScheduleMapOutput {
	return i.ToGlobalVMShutdownScheduleMapOutputWithContext(context.Background())
}

func (i GlobalVMShutdownScheduleMap) ToGlobalVMShutdownScheduleMapOutputWithContext(ctx context.Context) GlobalVMShutdownScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalVMShutdownScheduleMapOutput)
}

type GlobalVMShutdownScheduleOutput struct {
	*pulumi.OutputState
}

func (GlobalVMShutdownScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalVMShutdownSchedule)(nil))
}

func (o GlobalVMShutdownScheduleOutput) ToGlobalVMShutdownScheduleOutput() GlobalVMShutdownScheduleOutput {
	return o
}

func (o GlobalVMShutdownScheduleOutput) ToGlobalVMShutdownScheduleOutputWithContext(ctx context.Context) GlobalVMShutdownScheduleOutput {
	return o
}

func (o GlobalVMShutdownScheduleOutput) ToGlobalVMShutdownSchedulePtrOutput() GlobalVMShutdownSchedulePtrOutput {
	return o.ToGlobalVMShutdownSchedulePtrOutputWithContext(context.Background())
}

func (o GlobalVMShutdownScheduleOutput) ToGlobalVMShutdownSchedulePtrOutputWithContext(ctx context.Context) GlobalVMShutdownSchedulePtrOutput {
	return o.ApplyT(func(v GlobalVMShutdownSchedule) *GlobalVMShutdownSchedule {
		return &v
	}).(GlobalVMShutdownSchedulePtrOutput)
}

type GlobalVMShutdownSchedulePtrOutput struct {
	*pulumi.OutputState
}

func (GlobalVMShutdownSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalVMShutdownSchedule)(nil))
}

func (o GlobalVMShutdownSchedulePtrOutput) ToGlobalVMShutdownSchedulePtrOutput() GlobalVMShutdownSchedulePtrOutput {
	return o
}

func (o GlobalVMShutdownSchedulePtrOutput) ToGlobalVMShutdownSchedulePtrOutputWithContext(ctx context.Context) GlobalVMShutdownSchedulePtrOutput {
	return o
}

type GlobalVMShutdownScheduleArrayOutput struct{ *pulumi.OutputState }

func (GlobalVMShutdownScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalVMShutdownSchedule)(nil))
}

func (o GlobalVMShutdownScheduleArrayOutput) ToGlobalVMShutdownScheduleArrayOutput() GlobalVMShutdownScheduleArrayOutput {
	return o
}

func (o GlobalVMShutdownScheduleArrayOutput) ToGlobalVMShutdownScheduleArrayOutputWithContext(ctx context.Context) GlobalVMShutdownScheduleArrayOutput {
	return o
}

func (o GlobalVMShutdownScheduleArrayOutput) Index(i pulumi.IntInput) GlobalVMShutdownScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GlobalVMShutdownSchedule {
		return vs[0].([]GlobalVMShutdownSchedule)[vs[1].(int)]
	}).(GlobalVMShutdownScheduleOutput)
}

type GlobalVMShutdownScheduleMapOutput struct{ *pulumi.OutputState }

func (GlobalVMShutdownScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GlobalVMShutdownSchedule)(nil))
}

func (o GlobalVMShutdownScheduleMapOutput) ToGlobalVMShutdownScheduleMapOutput() GlobalVMShutdownScheduleMapOutput {
	return o
}

func (o GlobalVMShutdownScheduleMapOutput) ToGlobalVMShutdownScheduleMapOutputWithContext(ctx context.Context) GlobalVMShutdownScheduleMapOutput {
	return o
}

func (o GlobalVMShutdownScheduleMapOutput) MapIndex(k pulumi.StringInput) GlobalVMShutdownScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GlobalVMShutdownSchedule {
		return vs[0].(map[string]GlobalVMShutdownSchedule)[vs[1].(string)]
	}).(GlobalVMShutdownScheduleOutput)
}

func init() {
	pulumi.RegisterOutputType(GlobalVMShutdownScheduleOutput{})
	pulumi.RegisterOutputType(GlobalVMShutdownSchedulePtrOutput{})
	pulumi.RegisterOutputType(GlobalVMShutdownScheduleArrayOutput{})
	pulumi.RegisterOutputType(GlobalVMShutdownScheduleMapOutput{})
}
