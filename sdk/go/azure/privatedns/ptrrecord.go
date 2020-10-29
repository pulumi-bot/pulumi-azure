// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS PTR Records within Azure Private DNS.
type PTRRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS PTR Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS PTR Record. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of Fully Qualified Domain Names.
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewPTRRecord registers a new resource with the given unique name, arguments, and options.
func NewPTRRecord(ctx *pulumi.Context,
	name string, args *PTRRecordArgs, opts ...pulumi.ResourceOption) (*PTRRecord, error) {
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
		args = &PTRRecordArgs{}
	}
	var resource PTRRecord
	err := ctx.RegisterResource("azure:privatedns/pTRRecord:PTRRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPTRRecord gets an existing PTRRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPTRRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PTRRecordState, opts ...pulumi.ResourceOption) (*PTRRecord, error) {
	var resource PTRRecord
	err := ctx.ReadResource("azure:privatedns/pTRRecord:PTRRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PTRRecord resources.
type ptrrecordState struct {
	// The FQDN of the DNS PTR Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS PTR Record. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// List of Fully Qualified Domain Names.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type PTRRecordState struct {
	// The FQDN of the DNS PTR Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS PTR Record. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// List of Fully Qualified Domain Names.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (PTRRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*ptrrecordState)(nil)).Elem()
}

type ptrrecordArgs struct {
	// The name of the DNS PTR Record. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// List of Fully Qualified Domain Names.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a PTRRecord resource.
type PTRRecordArgs struct {
	// The name of the DNS PTR Record. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// List of Fully Qualified Domain Names.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (PTRRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ptrrecordArgs)(nil)).Elem()
}
