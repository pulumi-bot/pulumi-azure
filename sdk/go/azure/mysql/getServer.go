// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing MySQL Server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mysql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := mysql.LookupServer(ctx, &mysql.LookupServerArgs{
// 			Name:              "existingMySqlServer",
// 			ResourceGroupName: "existingResGroup",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure:mysql/getServer:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServer.
type LookupServerArgs struct {
	// Specifies the name of the MySQL Server.
	Name string `pulumi:"name"`
	// The name of the resource group for the MySQL Server.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getServer.
type LookupServerResult struct {
	// The Administrator Login for the MySQL Server.
	AdministratorLogin string `pulumi:"administratorLogin"`
	// The auto grow setting for this MySQL Server.
	AutoGrowEnabled bool `pulumi:"autoGrowEnabled"`
	// The backup retention days for this MySQL server.
	BackupRetentionDays int `pulumi:"backupRetentionDays"`
	// The FQDN of the MySQL Server.
	Fqdn string `pulumi:"fqdn"`
	// The geo redundant backup setting for this MySQL Server.
	GeoRedundantBackupEnabled bool `pulumi:"geoRedundantBackupEnabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An `identity` block for this MySQL server as defined below.
	Identities []GetServerIdentity `pulumi:"identities"`
	// Whether or not infrastructure is encrypted for this MySQL Server.
	InfrastructureEncryptionEnabled bool `pulumi:"infrastructureEncryptionEnabled"`
	// The Azure location where the resource exists.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// Whether or not public network access is allowed for this MySQL Server.
	PublicNetworkAccessEnabled bool   `pulumi:"publicNetworkAccessEnabled"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	RestorePointInTime         string `pulumi:"restorePointInTime"`
	// The SKU Name for this MySQL Server.
	SkuName string `pulumi:"skuName"`
	// Specifies if SSL should be enforced on connections for this MySQL Server.
	SslEnforcementEnabled bool `pulumi:"sslEnforcementEnabled"`
	// The minimum TLS version to support for this MySQL Server.
	SslMinimalTlsVersionEnforced string `pulumi:"sslMinimalTlsVersionEnforced"`
	// Max storage allowed for this MySQL Server.
	StorageMb int `pulumi:"storageMb"`
	// A mapping of tags to assign to the resource.
	// ---
	Tags map[string]string `pulumi:"tags"`
	// Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threatDetectionPolicy` block exports fields documented below.
	ThreatDetectionPolicies []GetServerThreatDetectionPolicy `pulumi:"threatDetectionPolicies"`
	// The version of this MySQL Server.
	Version string `pulumi:"version"`
}

func LookupServerOutput(ctx *pulumi.Context, args LookupServerOutputArgs, opts ...pulumi.InvokeOption) LookupServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerResult, error) {
			args := v.(LookupServerArgs)
			r, err := LookupServer(ctx, &args, opts...)
			return *r, err
		}).(LookupServerResultOutput)
}

// A collection of arguments for invoking getServer.
type LookupServerOutputArgs struct {
	// Specifies the name of the MySQL Server.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group for the MySQL Server.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}

// A collection of values returned by getServer.
type LookupServerResultOutput struct{ *pulumi.OutputState }

func (LookupServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerResult)(nil)).Elem()
}

func (o LookupServerResultOutput) ToLookupServerResultOutput() LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToLookupServerResultOutputWithContext(ctx context.Context) LookupServerResultOutput {
	return o
}

// The Administrator Login for the MySQL Server.
func (o LookupServerResultOutput) AdministratorLogin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.AdministratorLogin }).(pulumi.StringOutput)
}

// The auto grow setting for this MySQL Server.
func (o LookupServerResultOutput) AutoGrowEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerResult) bool { return v.AutoGrowEnabled }).(pulumi.BoolOutput)
}

// The backup retention days for this MySQL server.
func (o LookupServerResultOutput) BackupRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServerResult) int { return v.BackupRetentionDays }).(pulumi.IntOutput)
}

// The FQDN of the MySQL Server.
func (o LookupServerResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// The geo redundant backup setting for this MySQL Server.
func (o LookupServerResultOutput) GeoRedundantBackupEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerResult) bool { return v.GeoRedundantBackupEnabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// An `identity` block for this MySQL server as defined below.
func (o LookupServerResultOutput) Identities() GetServerIdentityArrayOutput {
	return o.ApplyT(func(v LookupServerResult) []GetServerIdentity { return v.Identities }).(GetServerIdentityArrayOutput)
}

// Whether or not infrastructure is encrypted for this MySQL Server.
func (o LookupServerResultOutput) InfrastructureEncryptionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerResult) bool { return v.InfrastructureEncryptionEnabled }).(pulumi.BoolOutput)
}

// The Azure location where the resource exists.
func (o LookupServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether or not public network access is allowed for this MySQL Server.
func (o LookupServerResultOutput) PublicNetworkAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerResult) bool { return v.PublicNetworkAccessEnabled }).(pulumi.BoolOutput)
}

func (o LookupServerResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) RestorePointInTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.RestorePointInTime }).(pulumi.StringOutput)
}

// The SKU Name for this MySQL Server.
func (o LookupServerResultOutput) SkuName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.SkuName }).(pulumi.StringOutput)
}

// Specifies if SSL should be enforced on connections for this MySQL Server.
func (o LookupServerResultOutput) SslEnforcementEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerResult) bool { return v.SslEnforcementEnabled }).(pulumi.BoolOutput)
}

// The minimum TLS version to support for this MySQL Server.
func (o LookupServerResultOutput) SslMinimalTlsVersionEnforced() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.SslMinimalTlsVersionEnforced }).(pulumi.StringOutput)
}

// Max storage allowed for this MySQL Server.
func (o LookupServerResultOutput) StorageMb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServerResult) int { return v.StorageMb }).(pulumi.IntOutput)
}

// A mapping of tags to assign to the resource.
// ---
func (o LookupServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Threat detection policy configuration, known in the API as Server Security Alerts Policy. The `threatDetectionPolicy` block exports fields documented below.
func (o LookupServerResultOutput) ThreatDetectionPolicies() GetServerThreatDetectionPolicyArrayOutput {
	return o.ApplyT(func(v LookupServerResult) []GetServerThreatDetectionPolicy { return v.ThreatDetectionPolicies }).(GetServerThreatDetectionPolicyArrayOutput)
}

// The version of this MySQL Server.
func (o LookupServerResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
