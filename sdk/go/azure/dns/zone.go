// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
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
	// A mapping of tags to assign to the resource.
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
	// A mapping of tags to assign to the resource.
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
	// A mapping of tags to assign to the resource.
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
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// The name of the DNS Zone. Must be a valid domain name.
	Name pulumi.StringPtrInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}
