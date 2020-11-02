// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package relay

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Relay Hybrid Connection.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/relay"
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
// 		exampleNamespace, err := relay.NewNamespace(ctx, "exampleNamespace", &relay.NamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("Standard"),
// 			Tags: pulumi.StringMap{
// 				"source": pulumi.String("managed"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = relay.NewHybridConnection(ctx, "exampleHybridConnection", &relay.HybridConnectionArgs{
// 			ResourceGroupName:           exampleResourceGroup.Name,
// 			RelayNamespaceName:          exampleNamespace.Name,
// 			RequiresClientAuthorization: pulumi.Bool(false),
// 			UserMetadata:                pulumi.String("testmetadata"),
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
// Relay Hybrid Connection's can be imported using the `resource id`, e.g.
type HybridConnection struct {
	pulumi.CustomResourceState

	// Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	RelayNamespaceName pulumi.StringOutput `pulumi:"relayNamespaceName"`
	// Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
	RequiresClientAuthorization pulumi.BoolPtrOutput `pulumi:"requiresClientAuthorization"`
	// The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrOutput `pulumi:"userMetadata"`
}

// NewHybridConnection registers a new resource with the given unique name, arguments, and options.
func NewHybridConnection(ctx *pulumi.Context,
	name string, args *HybridConnectionArgs, opts ...pulumi.ResourceOption) (*HybridConnection, error) {
	if args == nil || args.RelayNamespaceName == nil {
		return nil, errors.New("missing required argument 'RelayNamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &HybridConnectionArgs{}
	}
	var resource HybridConnection
	err := ctx.RegisterResource("azure:relay/hybridConnection:HybridConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridConnection gets an existing HybridConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridConnectionState, opts ...pulumi.ResourceOption) (*HybridConnection, error) {
	var resource HybridConnection
	err := ctx.ReadResource("azure:relay/hybridConnection:HybridConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridConnection resources.
type hybridConnectionState struct {
	// Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	RelayNamespaceName *string `pulumi:"relayNamespaceName"`
	// Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
	RequiresClientAuthorization *bool `pulumi:"requiresClientAuthorization"`
	// The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata *string `pulumi:"userMetadata"`
}

type HybridConnectionState struct {
	// Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	RelayNamespaceName pulumi.StringPtrInput
	// Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
	RequiresClientAuthorization pulumi.BoolPtrInput
	// The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrInput
}

func (HybridConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionState)(nil)).Elem()
}

type hybridConnectionArgs struct {
	// Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	RelayNamespaceName string `pulumi:"relayNamespaceName"`
	// Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
	RequiresClientAuthorization *bool `pulumi:"requiresClientAuthorization"`
	// The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata *string `pulumi:"userMetadata"`
}

// The set of arguments for constructing a HybridConnection resource.
type HybridConnectionArgs struct {
	// Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	RelayNamespaceName pulumi.StringInput
	// Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created.
	RequiresClientAuthorization pulumi.BoolPtrInput
	// The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrInput
}

func (HybridConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionArgs)(nil)).Elem()
}
