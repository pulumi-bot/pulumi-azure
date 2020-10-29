// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Redis Linked Server (ie Geo Location)
type LinkedServer struct {
	pulumi.CustomResourceState

	// The ID of the linked Redis cache. Changing this forces a new Redis to be created.
	LinkedRedisCacheId pulumi.StringOutput `pulumi:"linkedRedisCacheId"`
	// The location of the linked Redis cache. Changing this forces a new Redis to be created.
	LinkedRedisCacheLocation pulumi.StringOutput `pulumi:"linkedRedisCacheLocation"`
	// The name of the linked server.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Redis caches exists. Changing this forces a new Redis to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The role of the linked Redis cache (eg "Secondary"). Changing this forces a new Redis to be created.
	ServerRole pulumi.StringOutput `pulumi:"serverRole"`
	// The name of Redis cache to link with. Changing this forces a new Redis to be created. (eg The primary role)
	TargetRedisCacheName pulumi.StringOutput `pulumi:"targetRedisCacheName"`
}

// NewLinkedServer registers a new resource with the given unique name, arguments, and options.
func NewLinkedServer(ctx *pulumi.Context,
	name string, args *LinkedServerArgs, opts ...pulumi.ResourceOption) (*LinkedServer, error) {
	if args == nil || args.LinkedRedisCacheId == nil {
		return nil, errors.New("missing required argument 'LinkedRedisCacheId'")
	}
	if args == nil || args.LinkedRedisCacheLocation == nil {
		return nil, errors.New("missing required argument 'LinkedRedisCacheLocation'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerRole == nil {
		return nil, errors.New("missing required argument 'ServerRole'")
	}
	if args == nil || args.TargetRedisCacheName == nil {
		return nil, errors.New("missing required argument 'TargetRedisCacheName'")
	}
	if args == nil {
		args = &LinkedServerArgs{}
	}
	var resource LinkedServer
	err := ctx.RegisterResource("azure:redis/linkedServer:LinkedServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServer gets an existing LinkedServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServerState, opts ...pulumi.ResourceOption) (*LinkedServer, error) {
	var resource LinkedServer
	err := ctx.ReadResource("azure:redis/linkedServer:LinkedServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServer resources.
type linkedServerState struct {
	// The ID of the linked Redis cache. Changing this forces a new Redis to be created.
	LinkedRedisCacheId *string `pulumi:"linkedRedisCacheId"`
	// The location of the linked Redis cache. Changing this forces a new Redis to be created.
	LinkedRedisCacheLocation *string `pulumi:"linkedRedisCacheLocation"`
	// The name of the linked server.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Redis caches exists. Changing this forces a new Redis to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The role of the linked Redis cache (eg "Secondary"). Changing this forces a new Redis to be created.
	ServerRole *string `pulumi:"serverRole"`
	// The name of Redis cache to link with. Changing this forces a new Redis to be created. (eg The primary role)
	TargetRedisCacheName *string `pulumi:"targetRedisCacheName"`
}

type LinkedServerState struct {
	// The ID of the linked Redis cache. Changing this forces a new Redis to be created.
	LinkedRedisCacheId pulumi.StringPtrInput
	// The location of the linked Redis cache. Changing this forces a new Redis to be created.
	LinkedRedisCacheLocation pulumi.StringPtrInput
	// The name of the linked server.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Redis caches exists. Changing this forces a new Redis to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The role of the linked Redis cache (eg "Secondary"). Changing this forces a new Redis to be created.
	ServerRole pulumi.StringPtrInput
	// The name of Redis cache to link with. Changing this forces a new Redis to be created. (eg The primary role)
	TargetRedisCacheName pulumi.StringPtrInput
}

func (LinkedServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServerState)(nil)).Elem()
}

type linkedServerArgs struct {
	// The ID of the linked Redis cache. Changing this forces a new Redis to be created.
	LinkedRedisCacheId string `pulumi:"linkedRedisCacheId"`
	// The location of the linked Redis cache. Changing this forces a new Redis to be created.
	LinkedRedisCacheLocation string `pulumi:"linkedRedisCacheLocation"`
	// The name of the Resource Group where the Redis caches exists. Changing this forces a new Redis to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The role of the linked Redis cache (eg "Secondary"). Changing this forces a new Redis to be created.
	ServerRole string `pulumi:"serverRole"`
	// The name of Redis cache to link with. Changing this forces a new Redis to be created. (eg The primary role)
	TargetRedisCacheName string `pulumi:"targetRedisCacheName"`
}

// The set of arguments for constructing a LinkedServer resource.
type LinkedServerArgs struct {
	// The ID of the linked Redis cache. Changing this forces a new Redis to be created.
	LinkedRedisCacheId pulumi.StringInput
	// The location of the linked Redis cache. Changing this forces a new Redis to be created.
	LinkedRedisCacheLocation pulumi.StringInput
	// The name of the Resource Group where the Redis caches exists. Changing this forces a new Redis to be created.
	ResourceGroupName pulumi.StringInput
	// The role of the linked Redis cache (eg "Secondary"). Changing this forces a new Redis to be created.
	ServerRole pulumi.StringInput
	// The name of Redis cache to link with. Changing this forces a new Redis to be created. (eg The primary role)
	TargetRedisCacheName pulumi.StringInput
}

func (LinkedServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServerArgs)(nil)).Elem()
}
