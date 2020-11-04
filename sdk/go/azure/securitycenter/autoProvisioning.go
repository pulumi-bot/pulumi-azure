// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securitycenter

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AutoProvisioning struct {
	pulumi.CustomResourceState

	AutoProvision pulumi.StringOutput `pulumi:"autoProvision"`
}

// NewAutoProvisioning registers a new resource with the given unique name, arguments, and options.
func NewAutoProvisioning(ctx *pulumi.Context,
	name string, args *AutoProvisioningArgs, opts ...pulumi.ResourceOption) (*AutoProvisioning, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AutoProvision == nil {
		return nil, errors.New("invalid value for required argument 'AutoProvision'")
	}
	var resource AutoProvisioning
	err := ctx.RegisterResource("azure:securitycenter/autoProvisioning:AutoProvisioning", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoProvisioning gets an existing AutoProvisioning resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoProvisioning(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoProvisioningState, opts ...pulumi.ResourceOption) (*AutoProvisioning, error) {
	var resource AutoProvisioning
	err := ctx.ReadResource("azure:securitycenter/autoProvisioning:AutoProvisioning", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoProvisioning resources.
type autoProvisioningState struct {
	AutoProvision *string `pulumi:"autoProvision"`
}

type AutoProvisioningState struct {
	AutoProvision pulumi.StringPtrInput
}

func (AutoProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoProvisioningState)(nil)).Elem()
}

type autoProvisioningArgs struct {
	AutoProvision string `pulumi:"autoProvision"`
}

// The set of arguments for constructing a AutoProvisioning resource.
type AutoProvisioningArgs struct {
	AutoProvision pulumi.StringInput
}

func (AutoProvisioningArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoProvisioningArgs)(nil)).Elem()
}
