// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS TXT Records within Azure Private DNS.
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
// 		_, err = core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testZone, err := privatedns.NewZone(ctx, "testZone", &privatedns.ZoneArgs{
// 			ResourceGroupName: pulumi.String(azurerm_resource_group.Test.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatedns.NewTxtRecord(ctx, "testTxtRecord", &privatedns.TxtRecordArgs{
// 			ResourceGroupName: pulumi.String(azurerm_resource_group.Test.Name),
// 			ZoneName:          testZone.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: privatedns.TxtRecordRecordArray{
// 				&privatedns.TxtRecordRecordArgs{
// 					Value: pulumi.String("v=spf1 mx ~all"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type TxtRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS TXT Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS TXT Record. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records TxtRecordRecordArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
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
	err := ctx.RegisterResource("azure:privatedns/txtRecord:TxtRecord", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:privatedns/txtRecord:TxtRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TxtRecord resources.
type txtRecordState struct {
	// The FQDN of the DNS TXT Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS TXT Record. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records []TxtRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type TxtRecordState struct {
	// The FQDN of the DNS TXT Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS TXT Record. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `record` blocks as defined below.
	Records TxtRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (TxtRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*txtRecordState)(nil)).Elem()
}

type txtRecordArgs struct {
	// The name of the DNS TXT Record. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records []TxtRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a TxtRecord resource.
type TxtRecordArgs struct {
	// The name of the DNS TXT Record. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `record` blocks as defined below.
	Records TxtRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (TxtRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*txtRecordArgs)(nil)).Elem()
}
