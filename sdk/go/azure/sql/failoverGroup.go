// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Create a failover group of databases on a collection of Azure SQL servers.
type FailoverGroup struct {
	pulumi.CustomResourceState

	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayOutput `pulumi:"databases"`
	// the location of the failover group.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServerArrayOutput `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyOutput `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyOutput `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// local replication role of the failover group instance.
	Role pulumi.StringOutput `pulumi:"role"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewFailoverGroup registers a new resource with the given unique name, arguments, and options.
func NewFailoverGroup(ctx *pulumi.Context,
	name string, args *FailoverGroupArgs, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	if args == nil || args.PartnerServers == nil {
		return nil, errors.New("missing required argument 'PartnerServers'")
	}
	if args == nil || args.ReadWriteEndpointFailoverPolicy == nil {
		return nil, errors.New("missing required argument 'ReadWriteEndpointFailoverPolicy'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil {
		args = &FailoverGroupArgs{}
	}
	var resource FailoverGroup
	err := ctx.RegisterResource("azure:sql/failoverGroup:FailoverGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFailoverGroup gets an existing FailoverGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFailoverGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FailoverGroupState, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	var resource FailoverGroup
	err := ctx.ReadResource("azure:sql/failoverGroup:FailoverGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FailoverGroup resources.
type failoverGroupState struct {
	// A list of database ids to add to the failover group
	Databases []string `pulumi:"databases"`
	// the location of the failover group.
	Location *string `pulumi:"location"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers []FailoverGroupPartnerServer `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy *FailoverGroupReadWriteEndpointFailoverPolicy `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy *FailoverGroupReadonlyEndpointFailoverPolicy `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// local replication role of the failover group instance.
	Role *string `pulumi:"role"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type FailoverGroupState struct {
	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayInput
	// the location of the failover group.
	Location pulumi.StringPtrInput
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServerArrayInput
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyPtrInput
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyPtrInput
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringPtrInput
	// local replication role of the failover group instance.
	Role pulumi.StringPtrInput
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (FailoverGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupState)(nil)).Elem()
}

type failoverGroupArgs struct {
	// A list of database ids to add to the failover group
	Databases []string `pulumi:"databases"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers []FailoverGroupPartnerServer `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicy `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy *FailoverGroupReadonlyEndpointFailoverPolicy `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FailoverGroup resource.
type FailoverGroupArgs struct {
	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayInput
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServerArrayInput
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyInput
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyPtrInput
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringInput
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (FailoverGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupArgs)(nil)).Elem()
}
