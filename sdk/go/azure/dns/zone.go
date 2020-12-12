// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS zones within Azure DNS. These zones are hosted on Azure's name servers to which you can delegate the zone from the parent domain.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/dns"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/privatedns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dns.NewZone(ctx, "example_public", &dns.ZoneArgs{
// 			ResourceGroupName: example.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatedns.NewZone(ctx, "example_private", &privatedns.ZoneArgs{
// 			ResourceGroupName: example.Name,
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
// DNS Zones can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:dns/zone:Zone zone1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1
// ```
type Zone struct {
	pulumi.CustomResourceState

	// (Optional) Maximum number of Records in the zone. Defaults to `1000`.
	MaxNumberOfRecordSets pulumi.IntOutput `pulumi:"maxNumberOfRecordSets"`
	// The name of the DNS Zone. Must be a valid domain name.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Optional) A list of values that make up the NS record for the zone.
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// (Optional) The number of records already in the zone.
	NumberOfRecordSets pulumi.IntOutput `pulumi:"numberOfRecordSets"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// An `soaRecord` block as defined below. Changing this forces a new resource to be created.
	SoaRecord ZoneSoaRecordOutput `pulumi:"soaRecord"`
	// A mapping of tags to assign to the Record Set.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Zone
	err := ctx.RegisterResource("azure:dns/zone:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("azure:dns/zone:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
	// (Optional) Maximum number of Records in the zone. Defaults to `1000`.
	MaxNumberOfRecordSets *int `pulumi:"maxNumberOfRecordSets"`
	// The name of the DNS Zone. Must be a valid domain name.
	Name *string `pulumi:"name"`
	// (Optional) A list of values that make up the NS record for the zone.
	NameServers []string `pulumi:"nameServers"`
	// (Optional) The number of records already in the zone.
	NumberOfRecordSets *int `pulumi:"numberOfRecordSets"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// An `soaRecord` block as defined below. Changing this forces a new resource to be created.
	SoaRecord *ZoneSoaRecord `pulumi:"soaRecord"`
	// A mapping of tags to assign to the Record Set.
	Tags map[string]string `pulumi:"tags"`
}

type ZoneState struct {
	// (Optional) Maximum number of Records in the zone. Defaults to `1000`.
	MaxNumberOfRecordSets pulumi.IntPtrInput
	// The name of the DNS Zone. Must be a valid domain name.
	Name pulumi.StringPtrInput
	// (Optional) A list of values that make up the NS record for the zone.
	NameServers pulumi.StringArrayInput
	// (Optional) The number of records already in the zone.
	NumberOfRecordSets pulumi.IntPtrInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// An `soaRecord` block as defined below. Changing this forces a new resource to be created.
	SoaRecord ZoneSoaRecordPtrInput
	// A mapping of tags to assign to the Record Set.
	Tags pulumi.StringMapInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// The name of the DNS Zone. Must be a valid domain name.
	Name *string `pulumi:"name"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// An `soaRecord` block as defined below. Changing this forces a new resource to be created.
	SoaRecord *ZoneSoaRecord `pulumi:"soaRecord"`
	// A mapping of tags to assign to the Record Set.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// The name of the DNS Zone. Must be a valid domain name.
	Name pulumi.StringPtrInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// An `soaRecord` block as defined below. Changing this forces a new resource to be created.
	SoaRecord ZoneSoaRecordPtrInput
	// A mapping of tags to assign to the Record Set.
	Tags pulumi.StringMapInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((*Zone)(nil))
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

func (i *Zone) ToZonePtrOutput() ZonePtrOutput {
	return i.ToZonePtrOutputWithContext(context.Background())
}

func (i *Zone) ToZonePtrOutputWithContext(ctx context.Context) ZonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZonePtrOutput)
}

type ZonePtrInput interface {
	pulumi.Input

	ToZonePtrOutput() ZonePtrOutput
	ToZonePtrOutputWithContext(ctx context.Context) ZonePtrOutput
}

type ZoneOutput struct {
	*pulumi.OutputState
}

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Zone)(nil))
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

type ZonePtrOutput struct {
	*pulumi.OutputState
}

func (ZonePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil))
}

func (o ZonePtrOutput) ToZonePtrOutput() ZonePtrOutput {
	return o
}

func (o ZonePtrOutput) ToZonePtrOutputWithContext(ctx context.Context) ZonePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ZoneOutput{})
	pulumi.RegisterOutputType(ZonePtrOutput{})
}
