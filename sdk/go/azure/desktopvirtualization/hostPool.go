// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Virtual Desktop Host Pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/desktopvirtualization"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = desktopvirtualization.NewHostPool(ctx, "exampleHostPool", &desktopvirtualization.HostPoolArgs{
// 			Location:               exampleResourceGroup.Location,
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			FriendlyName:           pulumi.String("pooleddepthfirst"),
// 			ValidateEnvironment:    pulumi.Bool(true),
// 			Description:            pulumi.String("Acceptance Test: A pooled host pool - pooleddepthfirst"),
// 			Type:                   pulumi.String("Pooled"),
// 			MaximumSessionsAllowed: pulumi.Int(50),
// 			LoadBalancerType:       pulumi.String("DepthFirst"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type HostPool struct {
	pulumi.CustomResourceState

	// A description for the Virtual Desktop Host Pool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A friendly name for the Virtual Desktop Host Pool.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// `Breadthfirst` load balancing distributes new user sessions across all available session hosts in the host pool.
	// `Depthfirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
	// `Persistent` should be used if the host pool type is `Personal`
	LoadBalancerType pulumi.StringOutput `pulumi:"loadBalancerType"`
	// The location/region where the Virtual Desktop Host Pool is
	// located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
	// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
	MaximumSessionsAllowed pulumi.IntPtrOutput `pulumi:"maximumSessionsAllowed"`
	// The name of the Virtual Desktop Host Pool. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// `Automatic` assignment – The service will select an available host and assign it to an user.
	// `Direct` Assignment – Admin selects a specific host to assign to an user.
	PersonalDesktopAssignmentType pulumi.StringPtrOutput `pulumi:"personalDesktopAssignmentType"`
	// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
	// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
	PreferredAppGroupType pulumi.StringPtrOutput `pulumi:"preferredAppGroupType"`
	// A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
	RegistrationInfo HostPoolRegistrationInfoPtrOutput `pulumi:"registrationInfo"`
	// The name of the resource group in which to
	// create the Virtual Desktop Host Pool. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the Virtual Desktop Host Pool. Valid options are
	// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
	Type                pulumi.StringOutput  `pulumi:"type"`
	ValidateEnvironment pulumi.BoolPtrOutput `pulumi:"validateEnvironment"`
}

// NewHostPool registers a new resource with the given unique name, arguments, and options.
func NewHostPool(ctx *pulumi.Context,
	name string, args *HostPoolArgs, opts ...pulumi.ResourceOption) (*HostPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.LoadBalancerType == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource HostPool
	err := ctx.RegisterResource("azure:desktopvirtualization/hostPool:HostPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostPool gets an existing HostPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostPoolState, opts ...pulumi.ResourceOption) (*HostPool, error) {
	var resource HostPool
	err := ctx.ReadResource("azure:desktopvirtualization/hostPool:HostPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostPool resources.
type hostPoolState struct {
	// A description for the Virtual Desktop Host Pool.
	Description *string `pulumi:"description"`
	// A friendly name for the Virtual Desktop Host Pool.
	FriendlyName *string `pulumi:"friendlyName"`
	// `Breadthfirst` load balancing distributes new user sessions across all available session hosts in the host pool.
	// `Depthfirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
	// `Persistent` should be used if the host pool type is `Personal`
	LoadBalancerType *string `pulumi:"loadBalancerType"`
	// The location/region where the Virtual Desktop Host Pool is
	// located. Changing the location/region forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
	// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
	MaximumSessionsAllowed *int `pulumi:"maximumSessionsAllowed"`
	// The name of the Virtual Desktop Host Pool. Changing the name
	// forces a new resource to be created.
	Name *string `pulumi:"name"`
	// `Automatic` assignment – The service will select an available host and assign it to an user.
	// `Direct` Assignment – Admin selects a specific host to assign to an user.
	PersonalDesktopAssignmentType *string `pulumi:"personalDesktopAssignmentType"`
	// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
	// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
	PreferredAppGroupType *string `pulumi:"preferredAppGroupType"`
	// A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
	RegistrationInfo *HostPoolRegistrationInfo `pulumi:"registrationInfo"`
	// The name of the resource group in which to
	// create the Virtual Desktop Host Pool. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the Virtual Desktop Host Pool. Valid options are
	// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
	Type                *string `pulumi:"type"`
	ValidateEnvironment *bool   `pulumi:"validateEnvironment"`
}

type HostPoolState struct {
	// A description for the Virtual Desktop Host Pool.
	Description pulumi.StringPtrInput
	// A friendly name for the Virtual Desktop Host Pool.
	FriendlyName pulumi.StringPtrInput
	// `Breadthfirst` load balancing distributes new user sessions across all available session hosts in the host pool.
	// `Depthfirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
	// `Persistent` should be used if the host pool type is `Personal`
	LoadBalancerType pulumi.StringPtrInput
	// The location/region where the Virtual Desktop Host Pool is
	// located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
	// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
	MaximumSessionsAllowed pulumi.IntPtrInput
	// The name of the Virtual Desktop Host Pool. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringPtrInput
	// `Automatic` assignment – The service will select an available host and assign it to an user.
	// `Direct` Assignment – Admin selects a specific host to assign to an user.
	PersonalDesktopAssignmentType pulumi.StringPtrInput
	// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
	// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
	PreferredAppGroupType pulumi.StringPtrInput
	// A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
	RegistrationInfo HostPoolRegistrationInfoPtrInput
	// The name of the resource group in which to
	// create the Virtual Desktop Host Pool. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of the Virtual Desktop Host Pool. Valid options are
	// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
	Type                pulumi.StringPtrInput
	ValidateEnvironment pulumi.BoolPtrInput
}

func (HostPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPoolState)(nil)).Elem()
}

type hostPoolArgs struct {
	// A description for the Virtual Desktop Host Pool.
	Description *string `pulumi:"description"`
	// A friendly name for the Virtual Desktop Host Pool.
	FriendlyName *string `pulumi:"friendlyName"`
	// `Breadthfirst` load balancing distributes new user sessions across all available session hosts in the host pool.
	// `Depthfirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
	// `Persistent` should be used if the host pool type is `Personal`
	LoadBalancerType string `pulumi:"loadBalancerType"`
	// The location/region where the Virtual Desktop Host Pool is
	// located. Changing the location/region forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
	// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
	MaximumSessionsAllowed *int `pulumi:"maximumSessionsAllowed"`
	// The name of the Virtual Desktop Host Pool. Changing the name
	// forces a new resource to be created.
	Name *string `pulumi:"name"`
	// `Automatic` assignment – The service will select an available host and assign it to an user.
	// `Direct` Assignment – Admin selects a specific host to assign to an user.
	PersonalDesktopAssignmentType *string `pulumi:"personalDesktopAssignmentType"`
	// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
	// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
	PreferredAppGroupType *string `pulumi:"preferredAppGroupType"`
	// A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
	RegistrationInfo *HostPoolRegistrationInfo `pulumi:"registrationInfo"`
	// The name of the resource group in which to
	// create the Virtual Desktop Host Pool. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the Virtual Desktop Host Pool. Valid options are
	// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
	Type                string `pulumi:"type"`
	ValidateEnvironment *bool  `pulumi:"validateEnvironment"`
}

// The set of arguments for constructing a HostPool resource.
type HostPoolArgs struct {
	// A description for the Virtual Desktop Host Pool.
	Description pulumi.StringPtrInput
	// A friendly name for the Virtual Desktop Host Pool.
	FriendlyName pulumi.StringPtrInput
	// `Breadthfirst` load balancing distributes new user sessions across all available session hosts in the host pool.
	// `Depthfirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
	// `Persistent` should be used if the host pool type is `Personal`
	LoadBalancerType pulumi.StringInput
	// The location/region where the Virtual Desktop Host Pool is
	// located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
	// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
	MaximumSessionsAllowed pulumi.IntPtrInput
	// The name of the Virtual Desktop Host Pool. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringPtrInput
	// `Automatic` assignment – The service will select an available host and assign it to an user.
	// `Direct` Assignment – Admin selects a specific host to assign to an user.
	PersonalDesktopAssignmentType pulumi.StringPtrInput
	// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
	// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
	PreferredAppGroupType pulumi.StringPtrInput
	// A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
	RegistrationInfo HostPoolRegistrationInfoPtrInput
	// The name of the resource group in which to
	// create the Virtual Desktop Host Pool. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of the Virtual Desktop Host Pool. Valid options are
	// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
	Type                pulumi.StringInput
	ValidateEnvironment pulumi.BoolPtrInput
}

func (HostPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPoolArgs)(nil)).Elem()
}
