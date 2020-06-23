// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS CNAME Records within Azure Private DNS.
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
// 		_, err = privatedns.NewCnameRecord(ctx, "exampleCnameRecord", &privatedns.CnameRecordArgs{
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
type CnameRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS CNAME Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS CNAME Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// The target of the CNAME.
	Record pulumi.StringOutput `pulumi:"record"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewCnameRecord registers a new resource with the given unique name, arguments, and options.
func NewCnameRecord(ctx *pulumi.Context,
	name string, args *CnameRecordArgs, opts ...pulumi.ResourceOption) (*CnameRecord, error) {
	if args == nil || args.Record == nil {
		return nil, errors.New("missing required argument 'Record'")
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
		args = &CnameRecordArgs{}
	}
	var resource CnameRecord
	err := ctx.RegisterResource("azure:privatedns/cnameRecord:CnameRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCnameRecord gets an existing CnameRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCnameRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CnameRecordState, opts ...pulumi.ResourceOption) (*CnameRecord, error) {
	var resource CnameRecord
	err := ctx.ReadResource("azure:privatedns/cnameRecord:CnameRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CnameRecord resources.
type cnameRecordState struct {
	// The FQDN of the DNS CNAME Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS CNAME Record.
	Name *string `pulumi:"name"`
	// The target of the CNAME.
	Record *string `pulumi:"record"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type CnameRecordState struct {
	// The FQDN of the DNS CNAME Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS CNAME Record.
	Name pulumi.StringPtrInput
	// The target of the CNAME.
	Record pulumi.StringPtrInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (CnameRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*cnameRecordState)(nil)).Elem()
}

type cnameRecordArgs struct {
	// The name of the DNS CNAME Record.
	Name *string `pulumi:"name"`
	// The target of the CNAME.
	Record string `pulumi:"record"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a CnameRecord resource.
type CnameRecordArgs struct {
	// The name of the DNS CNAME Record.
	Name pulumi.StringPtrInput
	// The target of the CNAME.
	Record pulumi.StringInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (CnameRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cnameRecordArgs)(nil)).Elem()
}
