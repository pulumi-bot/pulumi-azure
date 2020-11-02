// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS TXT Records within Azure DNS.
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
// 		_, err = dns.NewTxtRecord(ctx, "exampleTxtRecord", &dns.TxtRecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: dns.TxtRecordRecordArray{
// 				&dns.TxtRecordRecordArgs{
// 					Value: pulumi.String("google-site-authenticator"),
// 				},
// 				&dns.TxtRecordRecordArgs{
// 					Value: pulumi.String("more site information here"),
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
// TXT records can be imported using the `resource id`, e.g.
type TxtRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS TXT Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS TXT Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of values that make up the txt record. Each `record` block supports fields documented below.
	Records TxtRecordRecordArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewTxtRecord registers a new resource with the given unique name, arguments, and options.
func NewTxtRecord(ctx *pulumi.Context,
	name string, args *TxtRecordArgs, opts ...pulumi.ResourceOption) (*TxtRecord, error) {
	if args == nil || args.Records == nil {
		return nil, errors.New("missing required argument 'Records'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Ttl == nil {
		return nil, errors.New("missing required argument 'Ttl'")
	}
	if args == nil || args.ZoneName == nil {
		return nil, errors.New("missing required argument 'ZoneName'")
	}
	if args == nil {
		args = &TxtRecordArgs{}
	}
	var resource TxtRecord
	err := ctx.RegisterResource("azure:dns/txtRecord:TxtRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTxtRecord gets an existing TxtRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTxtRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TxtRecordState, opts ...pulumi.ResourceOption) (*TxtRecord, error) {
	var resource TxtRecord
	err := ctx.ReadResource("azure:dns/txtRecord:TxtRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TxtRecord resources.
type txtRecordState struct {
	// The FQDN of the DNS TXT Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS TXT Record.
	Name *string `pulumi:"name"`
	// A list of values that make up the txt record. Each `record` block supports fields documented below.
	Records []TxtRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl *int `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type TxtRecordState struct {
	// The FQDN of the DNS TXT Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS TXT Record.
	Name pulumi.StringPtrInput
	// A list of values that make up the txt record. Each `record` block supports fields documented below.
	Records TxtRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntPtrInput
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (TxtRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*txtRecordState)(nil)).Elem()
}

type txtRecordArgs struct {
	// The name of the DNS TXT Record.
	Name *string `pulumi:"name"`
	// A list of values that make up the txt record. Each `record` block supports fields documented below.
	Records []TxtRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl int `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a TxtRecord resource.
type TxtRecordArgs struct {
	// The name of the DNS TXT Record.
	Name pulumi.StringPtrInput
	// A list of values that make up the txt record. Each `record` block supports fields documented below.
	Records TxtRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntInput
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (TxtRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*txtRecordArgs)(nil)).Elem()
}
