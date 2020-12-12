// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS NS Records within Azure DNS.
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
// 		_, err = dns.NewNsRecord(ctx, "exampleNsRecord", &dns.NsRecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: pulumi.StringArray{
// 				pulumi.String("ns1.contoso.com"),
// 				pulumi.String("ns2.contoso.com"),
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
// NS records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:dns/nsRecord:NsRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1/NS/myrecord1
// ```
type NsRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS NS Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS NS Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of values that make up the NS record.
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

// NewNsRecord registers a new resource with the given unique name, arguments, and options.
func NewNsRecord(ctx *pulumi.Context,
	name string, args *NsRecordArgs, opts ...pulumi.ResourceOption) (*NsRecord, error) {
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
	var resource NsRecord
	err := ctx.RegisterResource("azure:dns/nsRecord:NsRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsRecord gets an existing NsRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsRecordState, opts ...pulumi.ResourceOption) (*NsRecord, error) {
	var resource NsRecord
	err := ctx.ReadResource("azure:dns/nsRecord:NsRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsRecord resources.
type nsRecordState struct {
	// The FQDN of the DNS NS Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS NS Record.
	Name *string `pulumi:"name"`
	// A list of values that make up the NS record.
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

type NsRecordState struct {
	// The FQDN of the DNS NS Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS NS Record.
	Name pulumi.StringPtrInput
	// A list of values that make up the NS record.
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

func (NsRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsRecordState)(nil)).Elem()
}

type nsRecordArgs struct {
	// The name of the DNS NS Record.
	Name *string `pulumi:"name"`
	// A list of values that make up the NS record.
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

// The set of arguments for constructing a NsRecord resource.
type NsRecordArgs struct {
	// The name of the DNS NS Record.
	Name pulumi.StringPtrInput
	// A list of values that make up the NS record.
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

func (NsRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsRecordArgs)(nil)).Elem()
}

type NsRecordInput interface {
	pulumi.Input

	ToNsRecordOutput() NsRecordOutput
	ToNsRecordOutputWithContext(ctx context.Context) NsRecordOutput
}

func (*NsRecord) ElementType() reflect.Type {
	return reflect.TypeOf((*NsRecord)(nil))
}

func (i *NsRecord) ToNsRecordOutput() NsRecordOutput {
	return i.ToNsRecordOutputWithContext(context.Background())
}

func (i *NsRecord) ToNsRecordOutputWithContext(ctx context.Context) NsRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsRecordOutput)
}

func (i *NsRecord) ToNsRecordPtrOutput() NsRecordPtrOutput {
	return i.ToNsRecordPtrOutputWithContext(context.Background())
}

func (i *NsRecord) ToNsRecordPtrOutputWithContext(ctx context.Context) NsRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsRecordPtrOutput)
}

type NsRecordPtrInput interface {
	pulumi.Input

	ToNsRecordPtrOutput() NsRecordPtrOutput
	ToNsRecordPtrOutputWithContext(ctx context.Context) NsRecordPtrOutput
}

type NsRecordOutput struct {
	*pulumi.OutputState
}

func (NsRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsRecord)(nil))
}

func (o NsRecordOutput) ToNsRecordOutput() NsRecordOutput {
	return o
}

func (o NsRecordOutput) ToNsRecordOutputWithContext(ctx context.Context) NsRecordOutput {
	return o
}

type NsRecordPtrOutput struct {
	*pulumi.OutputState
}

func (NsRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsRecord)(nil))
}

func (o NsRecordPtrOutput) ToNsRecordPtrOutput() NsRecordPtrOutput {
	return o
}

func (o NsRecordPtrOutput) ToNsRecordPtrOutputWithContext(ctx context.Context) NsRecordPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NsRecordOutput{})
	pulumi.RegisterOutputType(NsRecordPtrOutput{})
}
