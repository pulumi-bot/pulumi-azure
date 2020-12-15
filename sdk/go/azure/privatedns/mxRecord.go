// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS MX Records within Azure Private DNS.
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
// 		_, err = privatedns.NewMxRecord(ctx, "exampleMxRecord", &privatedns.MxRecordArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ZoneName:          exampleZone.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: privatedns.MxRecordRecordArray{
// 				&privatedns.MxRecordRecordArgs{
// 					Preference: pulumi.Int(10),
// 					Exchange:   pulumi.String("mx1.contoso.com"),
// 				},
// 				&privatedns.MxRecordRecordArgs{
// 					Preference: pulumi.Int(20),
// 					Exchange:   pulumi.String("backupmx.contoso.com"),
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
// Private DNS MX Records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:privatedns/mxRecord:MxRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/contoso.com/MX/@
// ```
type MxRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS MX Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS MX Record. Changing this forces a new resource to be created. Default to '@' for root zone entry.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records MxRecordRecordArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
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
	err := ctx.RegisterResource("azure:privatedns/mxRecord:MxRecord", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:privatedns/mxRecord:MxRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MxRecord resources.
type mxRecordState struct {
	// The FQDN of the DNS MX Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS MX Record. Changing this forces a new resource to be created. Default to '@' for root zone entry.
	Name *string `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records []MxRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type MxRecordState struct {
	// The FQDN of the DNS MX Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS MX Record. Changing this forces a new resource to be created. Default to '@' for root zone entry.
	Name pulumi.StringPtrInput
	// One or more `record` blocks as defined below.
	Records MxRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (MxRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*mxRecordState)(nil)).Elem()
}

type mxRecordArgs struct {
	// The name of the DNS MX Record. Changing this forces a new resource to be created. Default to '@' for root zone entry.
	Name *string `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records []MxRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a MxRecord resource.
type MxRecordArgs struct {
	// The name of the DNS MX Record. Changing this forces a new resource to be created. Default to '@' for root zone entry.
	Name pulumi.StringPtrInput
	// One or more `record` blocks as defined below.
	Records MxRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
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

func init() {
	pulumi.RegisterOutputType(MxRecordOutput{})
	pulumi.RegisterOutputType(MxRecordPtrOutput{})
}
