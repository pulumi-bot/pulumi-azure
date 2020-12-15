// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS CNAME Records within Azure DNS.
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
// 		_, err = dns.NewCNameRecord(ctx, "exampleCNameRecord", &dns.CNameRecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Record:            pulumi.String("contoso.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Alias Record)
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
// 		target, err := dns.NewCNameRecord(ctx, "target", &dns.CNameRecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Record:            pulumi.String("contoso.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dns.NewCNameRecord(ctx, "exampleCNameRecord", &dns.CNameRecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			TargetResourceId:  target.ID(),
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
// CNAME records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:dns/cNameRecord:CNameRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1/CNAME/myrecord1
// ```
type CNameRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS CName Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS CNAME Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// The target of the CNAME.
	Record pulumi.StringPtrOutput `pulumi:"record"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Azure resource id of the target object. Conflicts with `records`
	TargetResourceId pulumi.StringPtrOutput `pulumi:"targetResourceId"`
	Ttl              pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewCNameRecord registers a new resource with the given unique name, arguments, and options.
func NewCNameRecord(ctx *pulumi.Context,
	name string, args *CNameRecordArgs, opts ...pulumi.ResourceOption) (*CNameRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource CNameRecord
	err := ctx.RegisterResource("azure:dns/cNameRecord:CNameRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCNameRecord gets an existing CNameRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCNameRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CNameRecordState, opts ...pulumi.ResourceOption) (*CNameRecord, error) {
	var resource CNameRecord
	err := ctx.ReadResource("azure:dns/cNameRecord:CNameRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CNameRecord resources.
type cnameRecordState struct {
	// The FQDN of the DNS CName Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS CNAME Record.
	Name *string `pulumi:"name"`
	// The target of the CNAME.
	Record *string `pulumi:"record"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Azure resource id of the target object. Conflicts with `records`
	TargetResourceId *string `pulumi:"targetResourceId"`
	Ttl              *int    `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type CNameRecordState struct {
	// The FQDN of the DNS CName Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS CNAME Record.
	Name pulumi.StringPtrInput
	// The target of the CNAME.
	Record pulumi.StringPtrInput
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Azure resource id of the target object. Conflicts with `records`
	TargetResourceId pulumi.StringPtrInput
	Ttl              pulumi.IntPtrInput
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (CNameRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*cnameRecordState)(nil)).Elem()
}

type cnameRecordArgs struct {
	// The name of the DNS CNAME Record.
	Name *string `pulumi:"name"`
	// The target of the CNAME.
	Record *string `pulumi:"record"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Azure resource id of the target object. Conflicts with `records`
	TargetResourceId *string `pulumi:"targetResourceId"`
	Ttl              int     `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a CNameRecord resource.
type CNameRecordArgs struct {
	// The name of the DNS CNAME Record.
	Name pulumi.StringPtrInput
	// The target of the CNAME.
	Record pulumi.StringPtrInput
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Azure resource id of the target object. Conflicts with `records`
	TargetResourceId pulumi.StringPtrInput
	Ttl              pulumi.IntInput
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (CNameRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cnameRecordArgs)(nil)).Elem()
}

type CNameRecordInput interface {
	pulumi.Input

	ToCNameRecordOutput() CNameRecordOutput
	ToCNameRecordOutputWithContext(ctx context.Context) CNameRecordOutput
}

func (*CNameRecord) ElementType() reflect.Type {
	return reflect.TypeOf((*CNameRecord)(nil))
}

func (i *CNameRecord) ToCNameRecordOutput() CNameRecordOutput {
	return i.ToCNameRecordOutputWithContext(context.Background())
}

func (i *CNameRecord) ToCNameRecordOutputWithContext(ctx context.Context) CNameRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CNameRecordOutput)
}

func (i *CNameRecord) ToCNameRecordPtrOutput() CNameRecordPtrOutput {
	return i.ToCNameRecordPtrOutputWithContext(context.Background())
}

func (i *CNameRecord) ToCNameRecordPtrOutputWithContext(ctx context.Context) CNameRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CNameRecordPtrOutput)
}

type CNameRecordPtrInput interface {
	pulumi.Input

	ToCNameRecordPtrOutput() CNameRecordPtrOutput
	ToCNameRecordPtrOutputWithContext(ctx context.Context) CNameRecordPtrOutput
}

type CNameRecordOutput struct {
	*pulumi.OutputState
}

func (CNameRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CNameRecord)(nil))
}

func (o CNameRecordOutput) ToCNameRecordOutput() CNameRecordOutput {
	return o
}

func (o CNameRecordOutput) ToCNameRecordOutputWithContext(ctx context.Context) CNameRecordOutput {
	return o
}

type CNameRecordPtrOutput struct {
	*pulumi.OutputState
}

func (CNameRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CNameRecord)(nil))
}

func (o CNameRecordPtrOutput) ToCNameRecordPtrOutput() CNameRecordPtrOutput {
	return o
}

func (o CNameRecordPtrOutput) ToCNameRecordPtrOutputWithContext(ctx context.Context) CNameRecordPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CNameRecordOutput{})
	pulumi.RegisterOutputType(CNameRecordPtrOutput{})
}
