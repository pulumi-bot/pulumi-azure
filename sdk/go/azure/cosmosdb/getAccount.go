// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing CosmosDB (formally DocumentDB) Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/cosmosdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cosmosdb.LookupAccount(ctx, &cosmosdb.LookupAccountArgs{
// 			Name:              "tfex-cosmosdb-account",
// 			ResourceGroupName: "tfex-cosmosdb-account-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cosmosdbAccountEndpoint", data.Azurerm_cosmosdb_account.Jobs.Endpoint)
// 		return nil
// 	})
// }
// ```
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure:cosmosdb/getAccount:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccount.
type LookupAccountArgs struct {
	// Specifies the name of the CosmosDB Account.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group in which the CosmosDB Account resides.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getAccount.
type LookupAccountResult struct {
	// Capabilities enabled on this Cosmos DB account.
	Capabilities        []GetAccountCapability        `pulumi:"capabilities"`
	ConsistencyPolicies []GetAccountConsistencyPolicy `pulumi:"consistencyPolicies"`
	// If automatic failover is enabled for this CosmosDB Account.
	EnableAutomaticFailover bool `pulumi:"enableAutomaticFailover"`
	// If Free Tier pricing option is enabled for this CosmosDB Account.
	EnableFreeTier bool `pulumi:"enableFreeTier"`
	// If multi-master is enabled for this Cosmos DB account.
	EnableMultipleWriteLocations bool `pulumi:"enableMultipleWriteLocations"`
	// The endpoint used to connect to the CosmosDB account.
	Endpoint     string                  `pulumi:"endpoint"`
	GeoLocations []GetAccountGeoLocation `pulumi:"geoLocations"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current IP Filter for this CosmosDB account
	IpRangeFilter string `pulumi:"ipRangeFilter"`
	// If virtual network filtering is enabled for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled bool `pulumi:"isVirtualNetworkFilterEnabled"`
	// The Key Vault key URI for CMK encryption.
	KeyVaultKeyId string `pulumi:"keyVaultKeyId"`
	// The Kind of the CosmosDB account.
	Kind string `pulumi:"kind"`
	// The name of the Azure region hosting replicated data.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// The Offer Type to used by this CosmosDB Account.
	OfferType string `pulumi:"offerType"`
	// The Primary master key for the CosmosDB Account.
	PrimaryKey string `pulumi:"primaryKey"`
	// Deprecated: This property has been renamed to `primary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	PrimaryMasterKey string `pulumi:"primaryMasterKey"`
	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyKey string `pulumi:"primaryReadonlyKey"`
	// Deprecated: This property has been renamed to `primary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	PrimaryReadonlyMasterKey string `pulumi:"primaryReadonlyMasterKey"`
	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints     []string `pulumi:"readEndpoints"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	// The Secondary master key for the CosmosDB Account.
	SecondaryKey string `pulumi:"secondaryKey"`
	// Deprecated: This property has been renamed to `secondary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	SecondaryMasterKey string `pulumi:"secondaryMasterKey"`
	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyKey string `pulumi:"secondaryReadonlyKey"`
	// Deprecated: This property has been renamed to `secondary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	SecondaryReadonlyMasterKey string `pulumi:"secondaryReadonlyMasterKey"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Subnets that are allowed to access this CosmosDB account.
	VirtualNetworkRules []GetAccountVirtualNetworkRule `pulumi:"virtualNetworkRules"`
	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints []string `pulumi:"writeEndpoints"`
}

func LookupAccountApply(ctx *pulumi.Context, args LookupAccountApplyInput, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return args.ToLookupAccountApplyOutput().ApplyT(func(v LookupAccountArgs) (LookupAccountResult, error) {
		r, err := LookupAccount(ctx, &v, opts...)
		return *r, err

	}).(LookupAccountResultOutput)
}

// LookupAccountApplyInput is an input type that accepts LookupAccountApplyArgs and LookupAccountApplyOutput values.
// You can construct a concrete instance of `LookupAccountApplyInput` via:
//
//          LookupAccountApplyArgs{...}
type LookupAccountApplyInput interface {
	pulumi.Input

	ToLookupAccountApplyOutput() LookupAccountApplyOutput
	ToLookupAccountApplyOutputWithContext(context.Context) LookupAccountApplyOutput
}

// A collection of arguments for invoking getAccount.
type LookupAccountApplyArgs struct {
	// Specifies the name of the CosmosDB Account.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the resource group in which the CosmosDB Account resides.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccountApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}

func (i LookupAccountApplyArgs) ToLookupAccountApplyOutput() LookupAccountApplyOutput {
	return i.ToLookupAccountApplyOutputWithContext(context.Background())
}

func (i LookupAccountApplyArgs) ToLookupAccountApplyOutputWithContext(ctx context.Context) LookupAccountApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupAccountApplyOutput)
}

// A collection of arguments for invoking getAccount.
type LookupAccountApplyOutput struct{ *pulumi.OutputState }

func (LookupAccountApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}

func (o LookupAccountApplyOutput) ToLookupAccountApplyOutput() LookupAccountApplyOutput {
	return o
}

func (o LookupAccountApplyOutput) ToLookupAccountApplyOutputWithContext(ctx context.Context) LookupAccountApplyOutput {
	return o
}

// Specifies the name of the CosmosDB Account.
func (o LookupAccountApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountArgs) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the name of the resource group in which the CosmosDB Account resides.
func (o LookupAccountApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A collection of values returned by getAccount.
type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

// Capabilities enabled on this Cosmos DB account.
func (o LookupAccountResultOutput) Capabilities() GetAccountCapabilityArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []GetAccountCapability { return v.Capabilities }).(GetAccountCapabilityArrayOutput)
}

func (o LookupAccountResultOutput) ConsistencyPolicies() GetAccountConsistencyPolicyArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []GetAccountConsistencyPolicy { return v.ConsistencyPolicies }).(GetAccountConsistencyPolicyArrayOutput)
}

// If automatic failover is enabled for this CosmosDB Account.
func (o LookupAccountResultOutput) EnableAutomaticFailover() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccountResult) bool { return v.EnableAutomaticFailover }).(pulumi.BoolOutput)
}

// If Free Tier pricing option is enabled for this CosmosDB Account.
func (o LookupAccountResultOutput) EnableFreeTier() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccountResult) bool { return v.EnableFreeTier }).(pulumi.BoolOutput)
}

// If multi-master is enabled for this Cosmos DB account.
func (o LookupAccountResultOutput) EnableMultipleWriteLocations() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccountResult) bool { return v.EnableMultipleWriteLocations }).(pulumi.BoolOutput)
}

// The endpoint used to connect to the CosmosDB account.
func (o LookupAccountResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) GeoLocations() GetAccountGeoLocationArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []GetAccountGeoLocation { return v.GeoLocations }).(GetAccountGeoLocationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The current IP Filter for this CosmosDB account
func (o LookupAccountResultOutput) IpRangeFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.IpRangeFilter }).(pulumi.StringOutput)
}

// If virtual network filtering is enabled for this Cosmos DB account.
func (o LookupAccountResultOutput) IsVirtualNetworkFilterEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccountResult) bool { return v.IsVirtualNetworkFilterEnabled }).(pulumi.BoolOutput)
}

// The Key Vault key URI for CMK encryption.
func (o LookupAccountResultOutput) KeyVaultKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.KeyVaultKeyId }).(pulumi.StringOutput)
}

// The Kind of the CosmosDB account.
func (o LookupAccountResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the Azure region hosting replicated data.
func (o LookupAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Offer Type to used by this CosmosDB Account.
func (o LookupAccountResultOutput) OfferType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.OfferType }).(pulumi.StringOutput)
}

// The Primary master key for the CosmosDB Account.
func (o LookupAccountResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// Deprecated: This property has been renamed to `primary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
func (o LookupAccountResultOutput) PrimaryMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.PrimaryMasterKey }).(pulumi.StringOutput)
}

// The Primary read-only master Key for the CosmosDB Account.
func (o LookupAccountResultOutput) PrimaryReadonlyKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.PrimaryReadonlyKey }).(pulumi.StringOutput)
}

// Deprecated: This property has been renamed to `primary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
func (o LookupAccountResultOutput) PrimaryReadonlyMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.PrimaryReadonlyMasterKey }).(pulumi.StringOutput)
}

// A list of read endpoints available for this CosmosDB account.
func (o LookupAccountResultOutput) ReadEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []string { return v.ReadEndpoints }).(pulumi.StringArrayOutput)
}

func (o LookupAccountResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The Secondary master key for the CosmosDB Account.
func (o LookupAccountResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

// Deprecated: This property has been renamed to `secondary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
func (o LookupAccountResultOutput) SecondaryMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.SecondaryMasterKey }).(pulumi.StringOutput)
}

// The Secondary read-only master key for the CosmosDB Account.
func (o LookupAccountResultOutput) SecondaryReadonlyKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.SecondaryReadonlyKey }).(pulumi.StringOutput)
}

// Deprecated: This property has been renamed to `secondary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
func (o LookupAccountResultOutput) SecondaryReadonlyMasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.SecondaryReadonlyMasterKey }).(pulumi.StringOutput)
}

// A mapping of tags assigned to the resource.
func (o LookupAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Subnets that are allowed to access this CosmosDB account.
func (o LookupAccountResultOutput) VirtualNetworkRules() GetAccountVirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []GetAccountVirtualNetworkRule { return v.VirtualNetworkRules }).(GetAccountVirtualNetworkRuleArrayOutput)
}

// A list of write endpoints available for this CosmosDB account.
func (o LookupAccountResultOutput) WriteEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []string { return v.WriteEndpoints }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountApplyOutput{})
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
