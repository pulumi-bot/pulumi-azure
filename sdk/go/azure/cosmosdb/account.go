// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a CosmosDB (formally DocumentDB) Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/cosmosdb"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.Any(_var.Resource_group_location),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = random.NewRandomInteger(ctx, "ri", &random.RandomIntegerArgs{
// 			Min: pulumi.Int(10000),
// 			Max: pulumi.Int(99999),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cosmosdb.NewAccount(ctx, "db", &cosmosdb.AccountArgs{
// 			Location:                rg.Location,
// 			ResourceGroupName:       rg.Name,
// 			OfferType:               pulumi.String("Standard"),
// 			Kind:                    pulumi.String("GlobalDocumentDB"),
// 			EnableAutomaticFailover: pulumi.Bool(true),
// 			Capabilities: cosmosdb.AccountCapabilityArray{
// 				&cosmosdb.AccountCapabilityArgs{
// 					Name: pulumi.String("EnableAggregationPipeline"),
// 				},
// 				&cosmosdb.AccountCapabilityArgs{
// 					Name: pulumi.String("mongoEnableDocLevelTTL"),
// 				},
// 				&cosmosdb.AccountCapabilityArgs{
// 					Name: pulumi.String("MongoDBv3.4"),
// 				},
// 			},
// 			ConsistencyPolicy: &cosmosdb.AccountConsistencyPolicyArgs{
// 				ConsistencyLevel:     pulumi.String("BoundedStaleness"),
// 				MaxIntervalInSeconds: pulumi.Int(10),
// 				MaxStalenessPrefix:   pulumi.Int(200),
// 			},
// 			GeoLocations: cosmosdb.AccountGeoLocationArray{
// 				&cosmosdb.AccountGeoLocationArgs{
// 					Location:         pulumi.Any(_var.Failover_location),
// 					FailoverPriority: pulumi.Int(1),
// 				},
// 				&cosmosdb.AccountGeoLocationArgs{
// 					Location:         rg.Location,
// 					FailoverPriority: pulumi.Int(0),
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
type Account struct {
	pulumi.CustomResourceState

	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, `EnableServerless`, and `mongoEnableDocLevelTTL`.
	Capabilities AccountCapabilityArrayOutput `pulumi:"capabilities"`
	// A list of connection strings available for this CosmosDB account.
	ConnectionStrings pulumi.StringArrayOutput `pulumi:"connectionStrings"`
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicyOutput `pulumi:"consistencyPolicy"`
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover pulumi.BoolPtrOutput `pulumi:"enableAutomaticFailover"`
	// Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
	EnableFreeTier pulumi.BoolPtrOutput `pulumi:"enableFreeTier"`
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations pulumi.BoolPtrOutput `pulumi:"enableMultipleWriteLocations"`
	// The endpoint used to connect to the CosmosDB account.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations AccountGeoLocationArrayOutput `pulumi:"geoLocations"`
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter pulumi.StringPtrOutput `pulumi:"ipRangeFilter"`
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrOutput `pulumi:"isVirtualNetworkFilterEnabled"`
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType pulumi.StringOutput `pulumi:"offerType"`
	// The Primary master key for the CosmosDB Account.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// Deprecated: This property has been renamed to `primary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	PrimaryMasterKey pulumi.StringOutput `pulumi:"primaryMasterKey"`
	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyKey pulumi.StringOutput `pulumi:"primaryReadonlyKey"`
	// Deprecated: This property has been renamed to `primary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	PrimaryReadonlyMasterKey pulumi.StringOutput `pulumi:"primaryReadonlyMasterKey"`
	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints pulumi.StringArrayOutput `pulumi:"readEndpoints"`
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary master key for the CosmosDB Account.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Deprecated: This property has been renamed to `secondary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	SecondaryMasterKey pulumi.StringOutput `pulumi:"secondaryMasterKey"`
	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyKey pulumi.StringOutput `pulumi:"secondaryReadonlyKey"`
	// Deprecated: This property has been renamed to `secondary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	SecondaryReadonlyMasterKey pulumi.StringOutput `pulumi:"secondaryReadonlyMasterKey"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules AccountVirtualNetworkRuleArrayOutput `pulumi:"virtualNetworkRules"`
	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints pulumi.StringArrayOutput `pulumi:"writeEndpoints"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ConsistencyPolicy == nil {
		return nil, errors.New("invalid value for required argument 'ConsistencyPolicy'")
	}
	if args.GeoLocations == nil {
		return nil, errors.New("invalid value for required argument 'GeoLocations'")
	}
	if args.OfferType == nil {
		return nil, errors.New("invalid value for required argument 'OfferType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Account
	err := ctx.RegisterResource("azure:cosmosdb/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure:cosmosdb/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, `EnableServerless`, and `mongoEnableDocLevelTTL`.
	Capabilities []AccountCapability `pulumi:"capabilities"`
	// A list of connection strings available for this CosmosDB account.
	ConnectionStrings []string `pulumi:"connectionStrings"`
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy *AccountConsistencyPolicy `pulumi:"consistencyPolicy"`
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover *bool `pulumi:"enableAutomaticFailover"`
	// Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
	EnableFreeTier *bool `pulumi:"enableFreeTier"`
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations *bool `pulumi:"enableMultipleWriteLocations"`
	// The endpoint used to connect to the CosmosDB account.
	Endpoint *string `pulumi:"endpoint"`
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations []AccountGeoLocation `pulumi:"geoLocations"`
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter *string `pulumi:"ipRangeFilter"`
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled *bool `pulumi:"isVirtualNetworkFilterEnabled"`
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind *string `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType *string `pulumi:"offerType"`
	// The Primary master key for the CosmosDB Account.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Deprecated: This property has been renamed to `primary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	PrimaryMasterKey *string `pulumi:"primaryMasterKey"`
	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyKey *string `pulumi:"primaryReadonlyKey"`
	// Deprecated: This property has been renamed to `primary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	PrimaryReadonlyMasterKey *string `pulumi:"primaryReadonlyMasterKey"`
	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints []string `pulumi:"readEndpoints"`
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary master key for the CosmosDB Account.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Deprecated: This property has been renamed to `secondary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	SecondaryMasterKey *string `pulumi:"secondaryMasterKey"`
	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyKey *string `pulumi:"secondaryReadonlyKey"`
	// Deprecated: This property has been renamed to `secondary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	SecondaryReadonlyMasterKey *string `pulumi:"secondaryReadonlyMasterKey"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules []AccountVirtualNetworkRule `pulumi:"virtualNetworkRules"`
	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints []string `pulumi:"writeEndpoints"`
}

type AccountState struct {
	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, `EnableServerless`, and `mongoEnableDocLevelTTL`.
	Capabilities AccountCapabilityArrayInput
	// A list of connection strings available for this CosmosDB account.
	ConnectionStrings pulumi.StringArrayInput
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicyPtrInput
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover pulumi.BoolPtrInput
	// Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
	EnableFreeTier pulumi.BoolPtrInput
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations pulumi.BoolPtrInput
	// The endpoint used to connect to the CosmosDB account.
	Endpoint pulumi.StringPtrInput
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations AccountGeoLocationArrayInput
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter pulumi.StringPtrInput
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrInput
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType pulumi.StringPtrInput
	// The Primary master key for the CosmosDB Account.
	PrimaryKey pulumi.StringPtrInput
	// Deprecated: This property has been renamed to `primary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	PrimaryMasterKey pulumi.StringPtrInput
	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyKey pulumi.StringPtrInput
	// Deprecated: This property has been renamed to `primary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	PrimaryReadonlyMasterKey pulumi.StringPtrInput
	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints pulumi.StringArrayInput
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary master key for the CosmosDB Account.
	SecondaryKey pulumi.StringPtrInput
	// Deprecated: This property has been renamed to `secondary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	SecondaryMasterKey pulumi.StringPtrInput
	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyKey pulumi.StringPtrInput
	// Deprecated: This property has been renamed to `secondary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	SecondaryReadonlyMasterKey pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules AccountVirtualNetworkRuleArrayInput
	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints pulumi.StringArrayInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, `EnableServerless`, and `mongoEnableDocLevelTTL`.
	Capabilities []AccountCapability `pulumi:"capabilities"`
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicy `pulumi:"consistencyPolicy"`
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover *bool `pulumi:"enableAutomaticFailover"`
	// Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
	EnableFreeTier *bool `pulumi:"enableFreeTier"`
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations *bool `pulumi:"enableMultipleWriteLocations"`
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations []AccountGeoLocation `pulumi:"geoLocations"`
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter *string `pulumi:"ipRangeFilter"`
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled *bool `pulumi:"isVirtualNetworkFilterEnabled"`
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind *string `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType string `pulumi:"offerType"`
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules []AccountVirtualNetworkRule `pulumi:"virtualNetworkRules"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, `EnableServerless`, and `mongoEnableDocLevelTTL`.
	Capabilities AccountCapabilityArrayInput
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicyInput
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover pulumi.BoolPtrInput
	// Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
	EnableFreeTier pulumi.BoolPtrInput
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations pulumi.BoolPtrInput
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations AccountGeoLocationArrayInput
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter pulumi.StringPtrInput
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrInput
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType pulumi.StringInput
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules AccountVirtualNetworkRuleArrayInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}
