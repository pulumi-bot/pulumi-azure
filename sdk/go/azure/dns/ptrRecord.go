// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS PTR Records within Azure DNS.
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
// 		_, err = dns.NewPtrRecord(ctx, "examplePtrRecord", &dns.PtrRecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: pulumi.StringArray{
// 				pulumi.String("yourdomain.com"),
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
// PTR records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:dns/ptrRecord:PtrRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1/PTR/myrecord1
// ```
type PtrRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS PTR Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS PTR Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of Fully Qualified Domain Names.
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewPtrRecord registers a new resource with the given unique name, arguments, and options.
func NewPtrRecord(ctx *pulumi.Context,
	name string, args *PtrRecordArgs, opts ...pulumi.ResourceOption) (*PtrRecord, error) {
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
	var resource PtrRecord
	err := ctx.RegisterResource("azure:dns/ptrRecord:PtrRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPtrRecord gets an existing PtrRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPtrRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PtrRecordState, opts ...pulumi.ResourceOption) (*PtrRecord, error) {
	var resource PtrRecord
	err := ctx.ReadResource("azure:dns/ptrRecord:PtrRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PtrRecord resources.
type ptrRecordState struct {
	// The FQDN of the DNS PTR Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS PTR Record.
	Name *string `pulumi:"name"`
	// List of Fully Qualified Domain Names.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl *int `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type PtrRecordState struct {
	// The FQDN of the DNS PTR Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS PTR Record.
	Name pulumi.StringPtrInput
	// List of Fully Qualified Domain Names.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntPtrInput
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (PtrRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*ptrRecordState)(nil)).Elem()
}

type ptrRecordArgs struct {
	// The name of the DNS PTR Record.
	Name *string `pulumi:"name"`
	// List of Fully Qualified Domain Names.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl int `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a PtrRecord resource.
type PtrRecordArgs struct {
	// The name of the DNS PTR Record.
	Name pulumi.StringPtrInput
	// List of Fully Qualified Domain Names.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntInput
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (PtrRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ptrRecordArgs)(nil)).Elem()
}

type PtrRecordInput interface {
	pulumi.Input

	ToPtrRecordOutput() PtrRecordOutput
	ToPtrRecordOutputWithContext(ctx context.Context) PtrRecordOutput
}

func (PtrRecord) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecord)(nil)).Elem()
}

func (i PtrRecord) ToPtrRecordOutput() PtrRecordOutput {
	return i.ToPtrRecordOutputWithContext(context.Background())
}

func (i PtrRecord) ToPtrRecordOutputWithContext(ctx context.Context) PtrRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtrRecordOutput)
}

type PtrRecordOutput struct {
	*pulumi.OutputState
}

func (PtrRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecordOutput)(nil)).Elem()
}

func (o PtrRecordOutput) ToPtrRecordOutput() PtrRecordOutput {
	return o
}

func (o PtrRecordOutput) ToPtrRecordOutputWithContext(ctx context.Context) PtrRecordOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PtrRecordOutput{})
}
