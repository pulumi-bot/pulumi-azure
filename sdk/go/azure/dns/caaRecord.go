// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS CAA Records within Azure DNS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/dns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
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
// 		_, err = dns.NewCaaRecord(ctx, "exampleCaaRecord", &dns.CaaRecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: dns.CaaRecordRecordArray{
// 				&dns.CaaRecordRecordArgs{
// 					Flags: pulumi.Int(0),
// 					Tag:   pulumi.String("issue"),
// 					Value: pulumi.String("example.com"),
// 				},
// 				&dns.CaaRecordRecordArgs{
// 					Flags: pulumi.Int(0),
// 					Tag:   pulumi.String("issue"),
// 					Value: pulumi.String("example.net"),
// 				},
// 				&dns.CaaRecordRecordArgs{
// 					Flags: pulumi.Int(0),
// 					Tag:   pulumi.String("issuewild"),
// 					Value: pulumi.String(";"),
// 				},
// 				&dns.CaaRecordRecordArgs{
// 					Flags: pulumi.Int(0),
// 					Tag:   pulumi.String("iodef"),
// 					Value: pulumi.String("mailto:user@nonexisting.tld"),
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
type CaaRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS CAA Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS CAA Record. If you are creating the record in the apex of the zone use `"@"` as the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of values that make up the CAA record. Each `record` block supports fields documented below.
	Records CaaRecordRecordArrayOutput `pulumi:"records"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewCaaRecord registers a new resource with the given unique name, arguments, and options.
func NewCaaRecord(ctx *pulumi.Context,
	name string, args *CaaRecordArgs, opts ...pulumi.ResourceOption) (*CaaRecord, error) {
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
	var resource CaaRecord
	err := ctx.RegisterResource("azure:dns/caaRecord:CaaRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaaRecord gets an existing CaaRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaaRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaaRecordState, opts ...pulumi.ResourceOption) (*CaaRecord, error) {
	var resource CaaRecord
	err := ctx.ReadResource("azure:dns/caaRecord:CaaRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaaRecord resources.
type caaRecordState struct {
	// The FQDN of the DNS CAA Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS CAA Record. If you are creating the record in the apex of the zone use `"@"` as the name.
	Name *string `pulumi:"name"`
	// A list of values that make up the CAA record. Each `record` block supports fields documented below.
	Records []CaaRecordRecord `pulumi:"records"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl *int `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type CaaRecordState struct {
	// The FQDN of the DNS CAA Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS CAA Record. If you are creating the record in the apex of the zone use `"@"` as the name.
	Name pulumi.StringPtrInput
	// A list of values that make up the CAA record. Each `record` block supports fields documented below.
	Records CaaRecordRecordArrayInput
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntPtrInput
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (CaaRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*caaRecordState)(nil)).Elem()
}

type caaRecordArgs struct {
	// The name of the DNS CAA Record. If you are creating the record in the apex of the zone use `"@"` as the name.
	Name *string `pulumi:"name"`
	// A list of values that make up the CAA record. Each `record` block supports fields documented below.
	Records []CaaRecordRecord `pulumi:"records"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl int `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a CaaRecord resource.
type CaaRecordArgs struct {
	// The name of the DNS CAA Record. If you are creating the record in the apex of the zone use `"@"` as the name.
	Name pulumi.StringPtrInput
	// A list of values that make up the CAA record. Each `record` block supports fields documented below.
	Records CaaRecordRecordArrayInput
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntInput
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (CaaRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caaRecordArgs)(nil)).Elem()
}
