// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables you to manage DNS AAAA Records within Azure Private DNS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/privatedns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testResourceGroup, err := core.NewResourceGroup(ctx, "testResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testZone, err := privatedns.NewZone(ctx, "testZone", &privatedns.ZoneArgs{
// 			ResourceGroupName: testResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatedns.NewAAAARecord(ctx, "testAAAARecord", &privatedns.AAAARecordArgs{
// 			ZoneName:          testZone.Name,
// 			ResourceGroupName: testResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: pulumi.StringArray{
// 				pulumi.String("fd5d:70bc:930e:d008:0000:0000:0000:7334"),
// 				pulumi.String("fd5d:70bc:930e:d008::7335"),
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
// Private DNS AAAA Records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:privatedns/aAAARecord:AAAARecord test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1/AAAA/myrecord1
// ```
type AAAARecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS AAAA Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS A Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of IPv6 Addresses.
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewAAAARecord registers a new resource with the given unique name, arguments, and options.
func NewAAAARecord(ctx *pulumi.Context,
	name string, args *AAAARecordArgs, opts ...pulumi.ResourceOption) (*AAAARecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Records == nil {
		return nil, errors.New("invalid value for required argument 'Records'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Ttl == nil {
		return nil, errors.New("invalid value for required argument 'Ttl'")
	}
	if args.ZoneName == nil {
		return nil, errors.New("invalid value for required argument 'ZoneName'")
	}
	var resource AAAARecord
	err := ctx.RegisterResource("azure:privatedns/aAAARecord:AAAARecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAAAARecord gets an existing AAAARecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAAAARecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AAAARecordState, opts ...pulumi.ResourceOption) (*AAAARecord, error) {
	var resource AAAARecord
	err := ctx.ReadResource("azure:privatedns/aAAARecord:AAAARecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AAAARecord resources.
type aaaarecordState struct {
	// The FQDN of the DNS AAAA Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS A Record.
	Name *string `pulumi:"name"`
	// A list of IPv6 Addresses.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type AAAARecordState struct {
	// The FQDN of the DNS AAAA Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS A Record.
	Name pulumi.StringPtrInput
	// A list of IPv6 Addresses.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (AAAARecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*aaaarecordState)(nil)).Elem()
}

type aaaarecordArgs struct {
	// The name of the DNS A Record.
	Name *string `pulumi:"name"`
	// A list of IPv6 Addresses.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a AAAARecord resource.
type AAAARecordArgs struct {
	// The name of the DNS A Record.
	Name pulumi.StringPtrInput
	// A list of IPv6 Addresses.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (AAAARecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aaaarecordArgs)(nil)).Elem()
}

type AAAARecordInput interface {
	pulumi.Input

	ToAAAARecordOutput() AAAARecordOutput
	ToAAAARecordOutputWithContext(ctx context.Context) AAAARecordOutput
}

func (*AAAARecord) ElementType() reflect.Type {
	return reflect.TypeOf((*AAAARecord)(nil))
}

func (i *AAAARecord) ToAAAARecordOutput() AAAARecordOutput {
	return i.ToAAAARecordOutputWithContext(context.Background())
}

func (i *AAAARecord) ToAAAARecordOutputWithContext(ctx context.Context) AAAARecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AAAARecordOutput)
}

func (i *AAAARecord) ToAAAARecordPtrOutput() AAAARecordPtrOutput {
	return i.ToAAAARecordPtrOutputWithContext(context.Background())
}

func (i *AAAARecord) ToAAAARecordPtrOutputWithContext(ctx context.Context) AAAARecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AAAARecordPtrOutput)
}

type AAAARecordPtrInput interface {
	pulumi.Input

	ToAAAARecordPtrOutput() AAAARecordPtrOutput
	ToAAAARecordPtrOutputWithContext(ctx context.Context) AAAARecordPtrOutput
}

type aaaarecordPtrType AAAARecordArgs

func (*aaaarecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AAAARecord)(nil))
}

func (i *aaaarecordPtrType) ToAAAARecordPtrOutput() AAAARecordPtrOutput {
	return i.ToAAAARecordPtrOutputWithContext(context.Background())
}

func (i *aaaarecordPtrType) ToAAAARecordPtrOutputWithContext(ctx context.Context) AAAARecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AAAARecordPtrOutput)
}

// AAAARecordArrayInput is an input type that accepts AAAARecordArray and AAAARecordArrayOutput values.
// You can construct a concrete instance of `AAAARecordArrayInput` via:
//
//          AAAARecordArray{ AAAARecordArgs{...} }
type AAAARecordArrayInput interface {
	pulumi.Input

	ToAAAARecordArrayOutput() AAAARecordArrayOutput
	ToAAAARecordArrayOutputWithContext(context.Context) AAAARecordArrayOutput
}

type AAAARecordArray []AAAARecordInput

func (AAAARecordArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AAAARecord)(nil))
}

func (i AAAARecordArray) ToAAAARecordArrayOutput() AAAARecordArrayOutput {
	return i.ToAAAARecordArrayOutputWithContext(context.Background())
}

func (i AAAARecordArray) ToAAAARecordArrayOutputWithContext(ctx context.Context) AAAARecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AAAARecordArrayOutput)
}

// AAAARecordMapInput is an input type that accepts AAAARecordMap and AAAARecordMapOutput values.
// You can construct a concrete instance of `AAAARecordMapInput` via:
//
//          AAAARecordMap{ "key": AAAARecordArgs{...} }
type AAAARecordMapInput interface {
	pulumi.Input

	ToAAAARecordMapOutput() AAAARecordMapOutput
	ToAAAARecordMapOutputWithContext(context.Context) AAAARecordMapOutput
}

type AAAARecordMap map[string]AAAARecordInput

func (AAAARecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AAAARecord)(nil))
}

func (i AAAARecordMap) ToAAAARecordMapOutput() AAAARecordMapOutput {
	return i.ToAAAARecordMapOutputWithContext(context.Background())
}

func (i AAAARecordMap) ToAAAARecordMapOutputWithContext(ctx context.Context) AAAARecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AAAARecordMapOutput)
}

type AAAARecordOutput struct {
	*pulumi.OutputState
}

func (AAAARecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AAAARecord)(nil))
}

func (o AAAARecordOutput) ToAAAARecordOutput() AAAARecordOutput {
	return o
}

func (o AAAARecordOutput) ToAAAARecordOutputWithContext(ctx context.Context) AAAARecordOutput {
	return o
}

func (o AAAARecordOutput) ToAAAARecordPtrOutput() AAAARecordPtrOutput {
	return o.ToAAAARecordPtrOutputWithContext(context.Background())
}

func (o AAAARecordOutput) ToAAAARecordPtrOutputWithContext(ctx context.Context) AAAARecordPtrOutput {
	return o.ApplyT(func(v AAAARecord) *AAAARecord {
		return &v
	}).(AAAARecordPtrOutput)
}

type AAAARecordPtrOutput struct {
	*pulumi.OutputState
}

func (AAAARecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AAAARecord)(nil))
}

func (o AAAARecordPtrOutput) ToAAAARecordPtrOutput() AAAARecordPtrOutput {
	return o
}

func (o AAAARecordPtrOutput) ToAAAARecordPtrOutputWithContext(ctx context.Context) AAAARecordPtrOutput {
	return o
}

type AAAARecordArrayOutput struct{ *pulumi.OutputState }

func (AAAARecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AAAARecord)(nil))
}

func (o AAAARecordArrayOutput) ToAAAARecordArrayOutput() AAAARecordArrayOutput {
	return o
}

func (o AAAARecordArrayOutput) ToAAAARecordArrayOutputWithContext(ctx context.Context) AAAARecordArrayOutput {
	return o
}

func (o AAAARecordArrayOutput) Index(i pulumi.IntInput) AAAARecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AAAARecord {
		return vs[0].([]AAAARecord)[vs[1].(int)]
	}).(AAAARecordOutput)
}

type AAAARecordMapOutput struct{ *pulumi.OutputState }

func (AAAARecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AAAARecord)(nil))
}

func (o AAAARecordMapOutput) ToAAAARecordMapOutput() AAAARecordMapOutput {
	return o
}

func (o AAAARecordMapOutput) ToAAAARecordMapOutputWithContext(ctx context.Context) AAAARecordMapOutput {
	return o
}

func (o AAAARecordMapOutput) MapIndex(k pulumi.StringInput) AAAARecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AAAARecord {
		return vs[0].(map[string]AAAARecord)[vs[1].(string)]
	}).(AAAARecordOutput)
}

func init() {
	pulumi.RegisterOutputType(AAAARecordOutput{})
	pulumi.RegisterOutputType(AAAARecordPtrOutput{})
	pulumi.RegisterOutputType(AAAARecordArrayOutput{})
	pulumi.RegisterOutputType(AAAARecordMapOutput{})
}
