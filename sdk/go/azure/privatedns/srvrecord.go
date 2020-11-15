// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS SRV Records within Azure Private DNS.
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
// 		_, err := core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testZone, err := privatedns.NewZone(ctx, "testZone", &privatedns.ZoneArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_resource_group.Test.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatedns.NewSRVRecord(ctx, "testSRVRecord", &privatedns.SRVRecordArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_resource_group.Test.Name),
// 			ZoneName:          testZone.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: privatedns.SRVRecordRecordArray{
// 				&privatedns.SRVRecordRecordArgs{
// 					Priority: pulumi.Int(1),
// 					Weight:   pulumi.Int(5),
// 					Port:     pulumi.Int(8080),
// 					Target:   pulumi.String("target1.contoso.com"),
// 				},
// 				&privatedns.SRVRecordRecordArgs{
// 					Priority: pulumi.Int(10),
// 					Weight:   pulumi.Int(10),
// 					Port:     pulumi.Int(8080),
// 					Target:   pulumi.String("target2.contoso.com"),
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
// Private DNS SRV Records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:privatedns/sRVRecord:SRVRecord test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/contoso.com/SRV/test
// ```
type SRVRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS SRV Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS SRV Record. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records SRVRecordRecordArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewSRVRecord registers a new resource with the given unique name, arguments, and options.
func NewSRVRecord(ctx *pulumi.Context,
	name string, args *SRVRecordArgs, opts ...pulumi.ResourceOption) (*SRVRecord, error) {
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
		args = &SRVRecordArgs{}
	}
	var resource SRVRecord
	err := ctx.RegisterResource("azure:privatedns/sRVRecord:SRVRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSRVRecord gets an existing SRVRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSRVRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SRVRecordState, opts ...pulumi.ResourceOption) (*SRVRecord, error) {
	var resource SRVRecord
	err := ctx.ReadResource("azure:privatedns/sRVRecord:SRVRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SRVRecord resources.
type srvrecordState struct {
	// The FQDN of the DNS SRV Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS SRV Record. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records []SRVRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type SRVRecordState struct {
	// The FQDN of the DNS SRV Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS SRV Record. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `record` blocks as defined below.
	Records SRVRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (SRVRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*srvrecordState)(nil)).Elem()
}

type srvrecordArgs struct {
	// The name of the DNS SRV Record. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records []SRVRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a SRVRecord resource.
type SRVRecordArgs struct {
	// The name of the DNS SRV Record. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `record` blocks as defined below.
	Records SRVRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (SRVRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*srvrecordArgs)(nil)).Elem()
}

type SRVRecordInput interface {
	pulumi.Input

	ToSRVRecordOutput() SRVRecordOutput
	ToSRVRecordOutputWithContext(ctx context.Context) SRVRecordOutput
}

func (SRVRecord) ElementType() reflect.Type {
	return reflect.TypeOf((*SRVRecord)(nil)).Elem()
}

func (i SRVRecord) ToSRVRecordOutput() SRVRecordOutput {
	return i.ToSRVRecordOutputWithContext(context.Background())
}

func (i SRVRecord) ToSRVRecordOutputWithContext(ctx context.Context) SRVRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SRVRecordOutput)
}

type SRVRecordOutput struct {
	*pulumi.OutputState
}

func (SRVRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SRVRecordOutput)(nil)).Elem()
}

func (o SRVRecordOutput) ToSRVRecordOutput() SRVRecordOutput {
	return o
}

func (o SRVRecordOutput) ToSRVRecordOutputWithContext(ctx context.Context) SRVRecordOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SRVRecordOutput{})
}
