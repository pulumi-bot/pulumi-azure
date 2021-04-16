// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables you to manage DNS MX Records within Azure DNS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/dns"
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
// 		exampleZone, err := dns.NewZone(ctx, "exampleZone", &dns.ZoneArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dns.NewMxRecord(ctx, "exampleMxRecord", &dns.MxRecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: dns.MxRecordRecordArray{
// 				&dns.MxRecordRecordArgs{
// 					Preference: pulumi.String("10"),
// 					Exchange:   pulumi.String("mail1.contoso.com"),
// 				},
// 				&dns.MxRecordRecordArgs{
// 					Preference: pulumi.String("20"),
// 					Exchange:   pulumi.String("mail2.contoso.com"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("Production"),
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
// MX records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:dns/mxRecord:MxRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1/MX/myrecord1
// ```
type MxRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS MX Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS MX Record. Defaults to `@` (root). Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of values that make up the MX record. Each `record` block supports fields documented below.
	Records MxRecordRecordArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewMxRecord registers a new resource with the given unique name, arguments, and options.
func NewMxRecord(ctx *pulumi.Context,
	name string, args *MxRecordArgs, opts ...pulumi.ResourceOption) (*MxRecord, error) {
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
	var resource MxRecord
	err := ctx.RegisterResource("azure:dns/mxRecord:MxRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMxRecord gets an existing MxRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMxRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MxRecordState, opts ...pulumi.ResourceOption) (*MxRecord, error) {
	var resource MxRecord
	err := ctx.ReadResource("azure:dns/mxRecord:MxRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MxRecord resources.
type mxRecordState struct {
	// The FQDN of the DNS MX Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS MX Record. Defaults to `@` (root). Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of values that make up the MX record. Each `record` block supports fields documented below.
	Records []MxRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl *int `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type MxRecordState struct {
	// The FQDN of the DNS MX Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS MX Record. Defaults to `@` (root). Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of values that make up the MX record. Each `record` block supports fields documented below.
	Records MxRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntPtrInput
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (MxRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*mxRecordState)(nil)).Elem()
}

type mxRecordArgs struct {
	// The name of the DNS MX Record. Defaults to `@` (root). Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of values that make up the MX record. Each `record` block supports fields documented below.
	Records []MxRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl int `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a MxRecord resource.
type MxRecordArgs struct {
	// The name of the DNS MX Record. Defaults to `@` (root). Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of values that make up the MX record. Each `record` block supports fields documented below.
	Records MxRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntInput
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (MxRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mxRecordArgs)(nil)).Elem()
}

type MxRecordInput interface {
	pulumi.Input

	ToMxRecordOutput() MxRecordOutput
	ToMxRecordOutputWithContext(ctx context.Context) MxRecordOutput
}

func (*MxRecord) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecord)(nil))
}

func (i *MxRecord) ToMxRecordOutput() MxRecordOutput {
	return i.ToMxRecordOutputWithContext(context.Background())
}

func (i *MxRecord) ToMxRecordOutputWithContext(ctx context.Context) MxRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordOutput)
}

func (i *MxRecord) ToMxRecordPtrOutput() MxRecordPtrOutput {
	return i.ToMxRecordPtrOutputWithContext(context.Background())
}

func (i *MxRecord) ToMxRecordPtrOutputWithContext(ctx context.Context) MxRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordPtrOutput)
}

type MxRecordPtrInput interface {
	pulumi.Input

	ToMxRecordPtrOutput() MxRecordPtrOutput
	ToMxRecordPtrOutputWithContext(ctx context.Context) MxRecordPtrOutput
}

type mxRecordPtrType MxRecordArgs

func (*mxRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MxRecord)(nil))
}

func (i *mxRecordPtrType) ToMxRecordPtrOutput() MxRecordPtrOutput {
	return i.ToMxRecordPtrOutputWithContext(context.Background())
}

func (i *mxRecordPtrType) ToMxRecordPtrOutputWithContext(ctx context.Context) MxRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordPtrOutput)
}

// MxRecordArrayInput is an input type that accepts MxRecordArray and MxRecordArrayOutput values.
// You can construct a concrete instance of `MxRecordArrayInput` via:
//
//          MxRecordArray{ MxRecordArgs{...} }
type MxRecordArrayInput interface {
	pulumi.Input

	ToMxRecordArrayOutput() MxRecordArrayOutput
	ToMxRecordArrayOutputWithContext(context.Context) MxRecordArrayOutput
}

type MxRecordArray []MxRecordInput

func (MxRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MxRecord)(nil))
}

func (i MxRecordArray) ToMxRecordArrayOutput() MxRecordArrayOutput {
	return i.ToMxRecordArrayOutputWithContext(context.Background())
}

func (i MxRecordArray) ToMxRecordArrayOutputWithContext(ctx context.Context) MxRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordArrayOutput)
}

// MxRecordMapInput is an input type that accepts MxRecordMap and MxRecordMapOutput values.
// You can construct a concrete instance of `MxRecordMapInput` via:
//
//          MxRecordMap{ "key": MxRecordArgs{...} }
type MxRecordMapInput interface {
	pulumi.Input

	ToMxRecordMapOutput() MxRecordMapOutput
	ToMxRecordMapOutputWithContext(context.Context) MxRecordMapOutput
}

type MxRecordMap map[string]MxRecordInput

func (MxRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MxRecord)(nil))
}

func (i MxRecordMap) ToMxRecordMapOutput() MxRecordMapOutput {
	return i.ToMxRecordMapOutputWithContext(context.Background())
}

func (i MxRecordMap) ToMxRecordMapOutputWithContext(ctx context.Context) MxRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordMapOutput)
}

type MxRecordOutput struct {
	*pulumi.OutputState
}

func (MxRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecord)(nil))
}

func (o MxRecordOutput) ToMxRecordOutput() MxRecordOutput {
	return o
}

func (o MxRecordOutput) ToMxRecordOutputWithContext(ctx context.Context) MxRecordOutput {
	return o
}

func (o MxRecordOutput) ToMxRecordPtrOutput() MxRecordPtrOutput {
	return o.ToMxRecordPtrOutputWithContext(context.Background())
}

func (o MxRecordOutput) ToMxRecordPtrOutputWithContext(ctx context.Context) MxRecordPtrOutput {
	return o.ApplyT(func(v MxRecord) *MxRecord {
		return &v
	}).(MxRecordPtrOutput)
}

type MxRecordPtrOutput struct {
	*pulumi.OutputState
}

func (MxRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MxRecord)(nil))
}

func (o MxRecordPtrOutput) ToMxRecordPtrOutput() MxRecordPtrOutput {
	return o
}

func (o MxRecordPtrOutput) ToMxRecordPtrOutputWithContext(ctx context.Context) MxRecordPtrOutput {
	return o
}

type MxRecordArrayOutput struct{ *pulumi.OutputState }

func (MxRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecord)(nil))
}

func (o MxRecordArrayOutput) ToMxRecordArrayOutput() MxRecordArrayOutput {
	return o
}

func (o MxRecordArrayOutput) ToMxRecordArrayOutputWithContext(ctx context.Context) MxRecordArrayOutput {
	return o
}

func (o MxRecordArrayOutput) Index(i pulumi.IntInput) MxRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MxRecord {
		return vs[0].([]MxRecord)[vs[1].(int)]
	}).(MxRecordOutput)
}

type MxRecordMapOutput struct{ *pulumi.OutputState }

func (MxRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MxRecord)(nil))
}

func (o MxRecordMapOutput) ToMxRecordMapOutput() MxRecordMapOutput {
	return o
}

func (o MxRecordMapOutput) ToMxRecordMapOutputWithContext(ctx context.Context) MxRecordMapOutput {
	return o
}

func (o MxRecordMapOutput) MapIndex(k pulumi.StringInput) MxRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MxRecord {
		return vs[0].(map[string]MxRecord)[vs[1].(string)]
	}).(MxRecordOutput)
}

func init() {
	pulumi.RegisterOutputType(MxRecordOutput{})
	pulumi.RegisterOutputType(MxRecordPtrOutput{})
	pulumi.RegisterOutputType(MxRecordArrayOutput{})
	pulumi.RegisterOutputType(MxRecordMapOutput{})
}
