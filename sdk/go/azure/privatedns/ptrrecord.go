// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables you to manage DNS PTR Records within Azure Private DNS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/privatedns"
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
// 		exampleZone, err := privatedns.NewZone(ctx, "exampleZone", &privatedns.ZoneArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatedns.NewPTRRecord(ctx, "examplePTRRecord", &privatedns.PTRRecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: pulumi.StringArray{
// 				pulumi.String("test.example.com"),
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
// Private DNS PTR Records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:privatedns/pTRRecord:PTRRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/2.0.192.in-addr.arpa/PTR/15
// ```
type PTRRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS PTR Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS PTR Record. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of Fully Qualified Domain Names.
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewPTRRecord registers a new resource with the given unique name, arguments, and options.
func NewPTRRecord(ctx *pulumi.Context,
	name string, args *PTRRecordArgs, opts ...pulumi.ResourceOption) (*PTRRecord, error) {
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
	var resource PTRRecord
	err := ctx.RegisterResource("azure:privatedns/pTRRecord:PTRRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPTRRecord gets an existing PTRRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPTRRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PTRRecordState, opts ...pulumi.ResourceOption) (*PTRRecord, error) {
	var resource PTRRecord
	err := ctx.ReadResource("azure:privatedns/pTRRecord:PTRRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PTRRecord resources.
type ptrrecordState struct {
	// The FQDN of the DNS PTR Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS PTR Record. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// List of Fully Qualified Domain Names.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type PTRRecordState struct {
	// The FQDN of the DNS PTR Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS PTR Record. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// List of Fully Qualified Domain Names.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (PTRRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*ptrrecordState)(nil)).Elem()
}

type ptrrecordArgs struct {
	// The name of the DNS PTR Record. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// List of Fully Qualified Domain Names.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a PTRRecord resource.
type PTRRecordArgs struct {
	// The name of the DNS PTR Record. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// List of Fully Qualified Domain Names.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (PTRRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ptrrecordArgs)(nil)).Elem()
}

type PTRRecordInput interface {
	pulumi.Input

	ToPTRRecordOutput() PTRRecordOutput
	ToPTRRecordOutputWithContext(ctx context.Context) PTRRecordOutput
}

func (*PTRRecord) ElementType() reflect.Type {
	return reflect.TypeOf((*PTRRecord)(nil))
}

func (i *PTRRecord) ToPTRRecordOutput() PTRRecordOutput {
	return i.ToPTRRecordOutputWithContext(context.Background())
}

func (i *PTRRecord) ToPTRRecordOutputWithContext(ctx context.Context) PTRRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PTRRecordOutput)
}

func (i *PTRRecord) ToPTRRecordPtrOutput() PTRRecordPtrOutput {
	return i.ToPTRRecordPtrOutputWithContext(context.Background())
}

func (i *PTRRecord) ToPTRRecordPtrOutputWithContext(ctx context.Context) PTRRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PTRRecordPtrOutput)
}

type PTRRecordPtrInput interface {
	pulumi.Input

	ToPTRRecordPtrOutput() PTRRecordPtrOutput
	ToPTRRecordPtrOutputWithContext(ctx context.Context) PTRRecordPtrOutput
}

type ptrrecordPtrType PTRRecordArgs

func (*ptrrecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PTRRecord)(nil))
}

func (i *ptrrecordPtrType) ToPTRRecordPtrOutput() PTRRecordPtrOutput {
	return i.ToPTRRecordPtrOutputWithContext(context.Background())
}

func (i *ptrrecordPtrType) ToPTRRecordPtrOutputWithContext(ctx context.Context) PTRRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PTRRecordPtrOutput)
}

// PTRRecordArrayInput is an input type that accepts PTRRecordArray and PTRRecordArrayOutput values.
// You can construct a concrete instance of `PTRRecordArrayInput` via:
//
//          PTRRecordArray{ PTRRecordArgs{...} }
type PTRRecordArrayInput interface {
	pulumi.Input

	ToPTRRecordArrayOutput() PTRRecordArrayOutput
	ToPTRRecordArrayOutputWithContext(context.Context) PTRRecordArrayOutput
}

type PTRRecordArray []PTRRecordInput

func (PTRRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PTRRecord)(nil))
}

func (i PTRRecordArray) ToPTRRecordArrayOutput() PTRRecordArrayOutput {
	return i.ToPTRRecordArrayOutputWithContext(context.Background())
}

func (i PTRRecordArray) ToPTRRecordArrayOutputWithContext(ctx context.Context) PTRRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PTRRecordArrayOutput)
}

// PTRRecordMapInput is an input type that accepts PTRRecordMap and PTRRecordMapOutput values.
// You can construct a concrete instance of `PTRRecordMapInput` via:
//
//          PTRRecordMap{ "key": PTRRecordArgs{...} }
type PTRRecordMapInput interface {
	pulumi.Input

	ToPTRRecordMapOutput() PTRRecordMapOutput
	ToPTRRecordMapOutputWithContext(context.Context) PTRRecordMapOutput
}

type PTRRecordMap map[string]PTRRecordInput

func (PTRRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PTRRecord)(nil))
}

func (i PTRRecordMap) ToPTRRecordMapOutput() PTRRecordMapOutput {
	return i.ToPTRRecordMapOutputWithContext(context.Background())
}

func (i PTRRecordMap) ToPTRRecordMapOutputWithContext(ctx context.Context) PTRRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PTRRecordMapOutput)
}

type PTRRecordOutput struct {
	*pulumi.OutputState
}

func (PTRRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PTRRecord)(nil))
}

func (o PTRRecordOutput) ToPTRRecordOutput() PTRRecordOutput {
	return o
}

func (o PTRRecordOutput) ToPTRRecordOutputWithContext(ctx context.Context) PTRRecordOutput {
	return o
}

func (o PTRRecordOutput) ToPTRRecordPtrOutput() PTRRecordPtrOutput {
	return o.ToPTRRecordPtrOutputWithContext(context.Background())
}

func (o PTRRecordOutput) ToPTRRecordPtrOutputWithContext(ctx context.Context) PTRRecordPtrOutput {
	return o.ApplyT(func(v PTRRecord) *PTRRecord {
		return &v
	}).(PTRRecordPtrOutput)
}

type PTRRecordPtrOutput struct {
	*pulumi.OutputState
}

func (PTRRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PTRRecord)(nil))
}

func (o PTRRecordPtrOutput) ToPTRRecordPtrOutput() PTRRecordPtrOutput {
	return o
}

func (o PTRRecordPtrOutput) ToPTRRecordPtrOutputWithContext(ctx context.Context) PTRRecordPtrOutput {
	return o
}

type PTRRecordArrayOutput struct{ *pulumi.OutputState }

func (PTRRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PTRRecord)(nil))
}

func (o PTRRecordArrayOutput) ToPTRRecordArrayOutput() PTRRecordArrayOutput {
	return o
}

func (o PTRRecordArrayOutput) ToPTRRecordArrayOutputWithContext(ctx context.Context) PTRRecordArrayOutput {
	return o
}

func (o PTRRecordArrayOutput) Index(i pulumi.IntInput) PTRRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PTRRecord {
		return vs[0].([]PTRRecord)[vs[1].(int)]
	}).(PTRRecordOutput)
}

type PTRRecordMapOutput struct{ *pulumi.OutputState }

func (PTRRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PTRRecord)(nil))
}

func (o PTRRecordMapOutput) ToPTRRecordMapOutput() PTRRecordMapOutput {
	return o
}

func (o PTRRecordMapOutput) ToPTRRecordMapOutputWithContext(ctx context.Context) PTRRecordMapOutput {
	return o
}

func (o PTRRecordMapOutput) MapIndex(k pulumi.StringInput) PTRRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PTRRecord {
		return vs[0].(map[string]PTRRecord)[vs[1].(string)]
	}).(PTRRecordOutput)
}

func init() {
	pulumi.RegisterOutputType(PTRRecordOutput{})
	pulumi.RegisterOutputType(PTRRecordPtrOutput{})
	pulumi.RegisterOutputType(PTRRecordArrayOutput{})
	pulumi.RegisterOutputType(PTRRecordMapOutput{})
}
