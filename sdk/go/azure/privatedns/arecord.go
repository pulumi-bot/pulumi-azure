// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS A Records within Azure Private DNS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/privatedns"
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
// 		exampleZone, err := privatedns.NewZone(ctx, "exampleZone", &privatedns.ZoneArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatedns.NewARecord(ctx, "exampleARecord", &privatedns.ARecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: pulumi.StringArray{
// 				pulumi.String("10.0.180.17"),
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
// Private DNS A Records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:privatedns/aRecord:ARecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1/A/myrecord1
// ```
type ARecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS A Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS A Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of IPv4 Addresses.
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewARecord registers a new resource with the given unique name, arguments, and options.
func NewARecord(ctx *pulumi.Context,
	name string, args *ARecordArgs, opts ...pulumi.ResourceOption) (*ARecord, error) {
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
	var resource ARecord
	err := ctx.RegisterResource("azure:privatedns/aRecord:ARecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetARecord gets an existing ARecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetARecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ARecordState, opts ...pulumi.ResourceOption) (*ARecord, error) {
	var resource ARecord
	err := ctx.ReadResource("azure:privatedns/aRecord:ARecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ARecord resources.
type arecordState struct {
	// The FQDN of the DNS A Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS A Record.
	Name *string `pulumi:"name"`
	// List of IPv4 Addresses.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type ARecordState struct {
	// The FQDN of the DNS A Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS A Record.
	Name pulumi.StringPtrInput
	// List of IPv4 Addresses.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (ARecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*arecordState)(nil)).Elem()
}

type arecordArgs struct {
	// The name of the DNS A Record.
	Name *string `pulumi:"name"`
	// List of IPv4 Addresses.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a ARecord resource.
type ARecordArgs struct {
	// The name of the DNS A Record.
	Name pulumi.StringPtrInput
	// List of IPv4 Addresses.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (ARecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*arecordArgs)(nil)).Elem()
}

type ARecordInput interface {
	pulumi.Input

	ToARecordOutput() ARecordOutput
	ToARecordOutputWithContext(ctx context.Context) ARecordOutput
}

func (ARecord) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecord)(nil)).Elem()
}

func (i ARecord) ToARecordOutput() ARecordOutput {
	return i.ToARecordOutputWithContext(context.Background())
}

func (i ARecord) ToARecordOutputWithContext(ctx context.Context) ARecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordOutput)
}

type ARecordOutput struct {
	*pulumi.OutputState
}

func (ARecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecordOutput)(nil)).Elem()
}

func (o ARecordOutput) ToARecordOutput() ARecordOutput {
	return o
}

func (o ARecordOutput) ToARecordOutputWithContext(ctx context.Context) ARecordOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ARecordOutput{})
}
