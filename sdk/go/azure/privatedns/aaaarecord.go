// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS AAAA Records within Azure Private DNS.
//
// ## Example Usage
//
//
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
// 		testResourceGroup, err := core.NewResourceGroup(ctx, "testResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testZone, err := privatedns.NewZone(ctx, "testZone", &privatedns.ZoneArgs{
// 			ResourceGroupName: testResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatedns.NewAAAARecord(ctx, "testAAAARecord", &privatedns.AAAARecordArgs{
// 			ZoneName:          testZone.Name,
// 			ResourceGroupName: testResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: pulumi.StringArray{
// 				pulumi.String("fd5d:70bc:930e:d008:0000:0000:0000:7334"),
// 				pulumi.String("fd5d:70bc:930e:d008::7335"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AAAARecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS AAAA Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS A Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of IPv6 Addresses.
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewAAAARecord registers a new resource with the given unique name, arguments, and options.
func NewAAAARecord(ctx *pulumi.Context,
	name string, args *AAAARecordArgs, opts ...pulumi.ResourceOption) (*AAAARecord, error) {
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
		args = &AAAARecordArgs{}
	}
	var resource AAAARecord
	err := ctx.RegisterResource("azure:privatedns/aAAARecord:AAAARecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAAAARecord gets an existing AAAARecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAAAARecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AAAARecordState, opts ...pulumi.ResourceOption) (*AAAARecord, error) {
	var resource AAAARecord
	err := ctx.ReadResource("azure:privatedns/aAAARecord:AAAARecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AAAARecord resources.
type aaaarecordState struct {
	// The FQDN of the DNS AAAA Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS A Record.
	Name *string `pulumi:"name"`
	// A list of IPv6 Addresses.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type AAAARecordState struct {
	// The FQDN of the DNS AAAA Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS A Record.
	Name pulumi.StringPtrInput
	// A list of IPv6 Addresses.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (AAAARecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*aaaarecordState)(nil)).Elem()
}

type aaaarecordArgs struct {
	// The name of the DNS A Record.
	Name *string `pulumi:"name"`
	// A list of IPv6 Addresses.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a AAAARecord resource.
type AAAARecordArgs struct {
	// The name of the DNS A Record.
	Name pulumi.StringPtrInput
	// A list of IPv6 Addresses.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (AAAARecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aaaarecordArgs)(nil)).Elem()
}
