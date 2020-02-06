// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DatabaseVulnerabilityAssessmentRuleBaselineBaselineResult struct {
	Results []string `pulumi:"results"`
}

type DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultInput interface {
	pulumi.Input

	ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput() DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput
	ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutputWithContext(context.Context) DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput
}

type DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArgs struct {
	Results pulumi.StringArrayInput `pulumi:"results"`
}

func (DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineBaselineResult)(nil)).Elem()
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArgs) ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput() DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput {
	return i.ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutputWithContext(context.Background())
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArgs) ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayInput interface {
	pulumi.Input

	ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput
	ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutputWithContext(context.Context) DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput
}

type DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArray []DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultInput

func (DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVulnerabilityAssessmentRuleBaselineBaselineResult)(nil)).Elem()
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArray) ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput {
	return i.ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutputWithContext(context.Background())
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArray) ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput struct { *pulumi.OutputState }

func (DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineBaselineResult)(nil)).Elem()
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput() DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput) Results() pulumi.StringArrayOutput {
	return o.ApplyT(func (v DatabaseVulnerabilityAssessmentRuleBaselineBaselineResult) []string { return v.Results }).(pulumi.StringArrayOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput struct { *pulumi.OutputState}

func (DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVulnerabilityAssessmentRuleBaselineBaselineResult)(nil)).Elem()
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput) Index(i pulumi.IntInput) DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) DatabaseVulnerabilityAssessmentRuleBaselineBaselineResult {
		return vs[0].([]DatabaseVulnerabilityAssessmentRuleBaselineBaselineResult)[vs[1].(int)]
	}).(DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput)
}

type ElasticPoolElasticPoolProperties struct {
	CreationDate *string `pulumi:"creationDate"`
	LicenseType *string `pulumi:"licenseType"`
	// The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
	MaxSizeBytes *int `pulumi:"maxSizeBytes"`
	State *string `pulumi:"state"`
	// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type ElasticPoolElasticPoolPropertiesInput interface {
	pulumi.Input

	ToElasticPoolElasticPoolPropertiesOutput() ElasticPoolElasticPoolPropertiesOutput
	ToElasticPoolElasticPoolPropertiesOutputWithContext(context.Context) ElasticPoolElasticPoolPropertiesOutput
}

type ElasticPoolElasticPoolPropertiesArgs struct {
	CreationDate pulumi.StringPtrInput `pulumi:"creationDate"`
	LicenseType pulumi.StringPtrInput `pulumi:"licenseType"`
	// The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
	MaxSizeBytes pulumi.IntPtrInput `pulumi:"maxSizeBytes"`
	State pulumi.StringPtrInput `pulumi:"state"`
	// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
	ZoneRedundant pulumi.BoolPtrInput `pulumi:"zoneRedundant"`
}

func (ElasticPoolElasticPoolPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolElasticPoolProperties)(nil)).Elem()
}

func (i ElasticPoolElasticPoolPropertiesArgs) ToElasticPoolElasticPoolPropertiesOutput() ElasticPoolElasticPoolPropertiesOutput {
	return i.ToElasticPoolElasticPoolPropertiesOutputWithContext(context.Background())
}

func (i ElasticPoolElasticPoolPropertiesArgs) ToElasticPoolElasticPoolPropertiesOutputWithContext(ctx context.Context) ElasticPoolElasticPoolPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolElasticPoolPropertiesOutput)
}

func (i ElasticPoolElasticPoolPropertiesArgs) ToElasticPoolElasticPoolPropertiesPtrOutput() ElasticPoolElasticPoolPropertiesPtrOutput {
	return i.ToElasticPoolElasticPoolPropertiesPtrOutputWithContext(context.Background())
}

func (i ElasticPoolElasticPoolPropertiesArgs) ToElasticPoolElasticPoolPropertiesPtrOutputWithContext(ctx context.Context) ElasticPoolElasticPoolPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolElasticPoolPropertiesOutput).ToElasticPoolElasticPoolPropertiesPtrOutputWithContext(ctx)
}

type ElasticPoolElasticPoolPropertiesPtrInput interface {
	pulumi.Input

	ToElasticPoolElasticPoolPropertiesPtrOutput() ElasticPoolElasticPoolPropertiesPtrOutput
	ToElasticPoolElasticPoolPropertiesPtrOutputWithContext(context.Context) ElasticPoolElasticPoolPropertiesPtrOutput
}

type elasticPoolElasticPoolPropertiesPtrType ElasticPoolElasticPoolPropertiesArgs

func ElasticPoolElasticPoolPropertiesPtr(v *ElasticPoolElasticPoolPropertiesArgs) ElasticPoolElasticPoolPropertiesPtrInput {	return (*elasticPoolElasticPoolPropertiesPtrType)(v)
}

func (*elasticPoolElasticPoolPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolElasticPoolProperties)(nil)).Elem()
}

func (i *elasticPoolElasticPoolPropertiesPtrType) ToElasticPoolElasticPoolPropertiesPtrOutput() ElasticPoolElasticPoolPropertiesPtrOutput {
	return i.ToElasticPoolElasticPoolPropertiesPtrOutputWithContext(context.Background())
}

func (i *elasticPoolElasticPoolPropertiesPtrType) ToElasticPoolElasticPoolPropertiesPtrOutputWithContext(ctx context.Context) ElasticPoolElasticPoolPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolElasticPoolPropertiesPtrOutput)
}

type ElasticPoolElasticPoolPropertiesOutput struct { *pulumi.OutputState }

func (ElasticPoolElasticPoolPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolElasticPoolProperties)(nil)).Elem()
}

func (o ElasticPoolElasticPoolPropertiesOutput) ToElasticPoolElasticPoolPropertiesOutput() ElasticPoolElasticPoolPropertiesOutput {
	return o
}

func (o ElasticPoolElasticPoolPropertiesOutput) ToElasticPoolElasticPoolPropertiesOutputWithContext(ctx context.Context) ElasticPoolElasticPoolPropertiesOutput {
	return o
}

func (o ElasticPoolElasticPoolPropertiesOutput) ToElasticPoolElasticPoolPropertiesPtrOutput() ElasticPoolElasticPoolPropertiesPtrOutput {
	return o.ToElasticPoolElasticPoolPropertiesPtrOutputWithContext(context.Background())
}

func (o ElasticPoolElasticPoolPropertiesOutput) ToElasticPoolElasticPoolPropertiesPtrOutputWithContext(ctx context.Context) ElasticPoolElasticPoolPropertiesPtrOutput {
	return o.ApplyT(func(v ElasticPoolElasticPoolProperties) *ElasticPoolElasticPoolProperties {
		return &v
	}).(ElasticPoolElasticPoolPropertiesPtrOutput)
}
func (o ElasticPoolElasticPoolPropertiesOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ElasticPoolElasticPoolProperties) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o ElasticPoolElasticPoolPropertiesOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ElasticPoolElasticPoolProperties) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
func (o ElasticPoolElasticPoolPropertiesOutput) MaxSizeBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ElasticPoolElasticPoolProperties) *int { return v.MaxSizeBytes }).(pulumi.IntPtrOutput)
}

func (o ElasticPoolElasticPoolPropertiesOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ElasticPoolElasticPoolProperties) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
func (o ElasticPoolElasticPoolPropertiesOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ElasticPoolElasticPoolProperties) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

type ElasticPoolElasticPoolPropertiesPtrOutput struct { *pulumi.OutputState}

func (ElasticPoolElasticPoolPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolElasticPoolProperties)(nil)).Elem()
}

func (o ElasticPoolElasticPoolPropertiesPtrOutput) ToElasticPoolElasticPoolPropertiesPtrOutput() ElasticPoolElasticPoolPropertiesPtrOutput {
	return o
}

func (o ElasticPoolElasticPoolPropertiesPtrOutput) ToElasticPoolElasticPoolPropertiesPtrOutputWithContext(ctx context.Context) ElasticPoolElasticPoolPropertiesPtrOutput {
	return o
}

func (o ElasticPoolElasticPoolPropertiesPtrOutput) Elem() ElasticPoolElasticPoolPropertiesOutput {
	return o.ApplyT(func (v *ElasticPoolElasticPoolProperties) ElasticPoolElasticPoolProperties { return *v }).(ElasticPoolElasticPoolPropertiesOutput)
}

func (o ElasticPoolElasticPoolPropertiesPtrOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ElasticPoolElasticPoolProperties) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o ElasticPoolElasticPoolPropertiesPtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ElasticPoolElasticPoolProperties) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// The max data size of the elastic pool in bytes. Conflicts with `maxSizeGb`.
func (o ElasticPoolElasticPoolPropertiesPtrOutput) MaxSizeBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ElasticPoolElasticPoolProperties) *int { return v.MaxSizeBytes }).(pulumi.IntPtrOutput)
}

func (o ElasticPoolElasticPoolPropertiesPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ElasticPoolElasticPoolProperties) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Whether or not this elastic pool is zone redundant. `tier` needs to be `Premium` for `DTU` based  or `BusinessCritical` for `vCore` based `sku`. Defaults to `false`.
func (o ElasticPoolElasticPoolPropertiesPtrOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ElasticPoolElasticPoolProperties) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

type ElasticPoolPerDatabaseSettings struct {
	// The maximum capacity any one database can consume.
	MaxCapacity float64 `pulumi:"maxCapacity"`
	// The minimum capacity all databases are guaranteed.
	MinCapacity float64 `pulumi:"minCapacity"`
}

type ElasticPoolPerDatabaseSettingsInput interface {
	pulumi.Input

	ToElasticPoolPerDatabaseSettingsOutput() ElasticPoolPerDatabaseSettingsOutput
	ToElasticPoolPerDatabaseSettingsOutputWithContext(context.Context) ElasticPoolPerDatabaseSettingsOutput
}

type ElasticPoolPerDatabaseSettingsArgs struct {
	// The maximum capacity any one database can consume.
	MaxCapacity pulumi.Float64Input `pulumi:"maxCapacity"`
	// The minimum capacity all databases are guaranteed.
	MinCapacity pulumi.Float64Input `pulumi:"minCapacity"`
}

func (ElasticPoolPerDatabaseSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolPerDatabaseSettings)(nil)).Elem()
}

func (i ElasticPoolPerDatabaseSettingsArgs) ToElasticPoolPerDatabaseSettingsOutput() ElasticPoolPerDatabaseSettingsOutput {
	return i.ToElasticPoolPerDatabaseSettingsOutputWithContext(context.Background())
}

func (i ElasticPoolPerDatabaseSettingsArgs) ToElasticPoolPerDatabaseSettingsOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPerDatabaseSettingsOutput)
}

func (i ElasticPoolPerDatabaseSettingsArgs) ToElasticPoolPerDatabaseSettingsPtrOutput() ElasticPoolPerDatabaseSettingsPtrOutput {
	return i.ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(context.Background())
}

func (i ElasticPoolPerDatabaseSettingsArgs) ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPerDatabaseSettingsOutput).ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(ctx)
}

type ElasticPoolPerDatabaseSettingsPtrInput interface {
	pulumi.Input

	ToElasticPoolPerDatabaseSettingsPtrOutput() ElasticPoolPerDatabaseSettingsPtrOutput
	ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(context.Context) ElasticPoolPerDatabaseSettingsPtrOutput
}

type elasticPoolPerDatabaseSettingsPtrType ElasticPoolPerDatabaseSettingsArgs

func ElasticPoolPerDatabaseSettingsPtr(v *ElasticPoolPerDatabaseSettingsArgs) ElasticPoolPerDatabaseSettingsPtrInput {	return (*elasticPoolPerDatabaseSettingsPtrType)(v)
}

func (*elasticPoolPerDatabaseSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolPerDatabaseSettings)(nil)).Elem()
}

func (i *elasticPoolPerDatabaseSettingsPtrType) ToElasticPoolPerDatabaseSettingsPtrOutput() ElasticPoolPerDatabaseSettingsPtrOutput {
	return i.ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(context.Background())
}

func (i *elasticPoolPerDatabaseSettingsPtrType) ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPerDatabaseSettingsPtrOutput)
}

type ElasticPoolPerDatabaseSettingsOutput struct { *pulumi.OutputState }

func (ElasticPoolPerDatabaseSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolPerDatabaseSettings)(nil)).Elem()
}

func (o ElasticPoolPerDatabaseSettingsOutput) ToElasticPoolPerDatabaseSettingsOutput() ElasticPoolPerDatabaseSettingsOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsOutput) ToElasticPoolPerDatabaseSettingsOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsOutput) ToElasticPoolPerDatabaseSettingsPtrOutput() ElasticPoolPerDatabaseSettingsPtrOutput {
	return o.ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(context.Background())
}

func (o ElasticPoolPerDatabaseSettingsOutput) ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsPtrOutput {
	return o.ApplyT(func(v ElasticPoolPerDatabaseSettings) *ElasticPoolPerDatabaseSettings {
		return &v
	}).(ElasticPoolPerDatabaseSettingsPtrOutput)
}
// The maximum capacity any one database can consume.
func (o ElasticPoolPerDatabaseSettingsOutput) MaxCapacity() pulumi.Float64Output {
	return o.ApplyT(func (v ElasticPoolPerDatabaseSettings) float64 { return v.MaxCapacity }).(pulumi.Float64Output)
}

// The minimum capacity all databases are guaranteed.
func (o ElasticPoolPerDatabaseSettingsOutput) MinCapacity() pulumi.Float64Output {
	return o.ApplyT(func (v ElasticPoolPerDatabaseSettings) float64 { return v.MinCapacity }).(pulumi.Float64Output)
}

type ElasticPoolPerDatabaseSettingsPtrOutput struct { *pulumi.OutputState}

func (ElasticPoolPerDatabaseSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolPerDatabaseSettings)(nil)).Elem()
}

func (o ElasticPoolPerDatabaseSettingsPtrOutput) ToElasticPoolPerDatabaseSettingsPtrOutput() ElasticPoolPerDatabaseSettingsPtrOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsPtrOutput) ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsPtrOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsPtrOutput) Elem() ElasticPoolPerDatabaseSettingsOutput {
	return o.ApplyT(func (v *ElasticPoolPerDatabaseSettings) ElasticPoolPerDatabaseSettings { return *v }).(ElasticPoolPerDatabaseSettingsOutput)
}

// The maximum capacity any one database can consume.
func (o ElasticPoolPerDatabaseSettingsPtrOutput) MaxCapacity() pulumi.Float64Output {
	return o.ApplyT(func (v ElasticPoolPerDatabaseSettings) float64 { return v.MaxCapacity }).(pulumi.Float64Output)
}

// The minimum capacity all databases are guaranteed.
func (o ElasticPoolPerDatabaseSettingsPtrOutput) MinCapacity() pulumi.Float64Output {
	return o.ApplyT(func (v ElasticPoolPerDatabaseSettings) float64 { return v.MinCapacity }).(pulumi.Float64Output)
}

type ElasticPoolSku struct {
	// The scale up/out capacity, representing server's compute units. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
	Capacity int `pulumi:"capacity"`
	// The `family` of hardware `Gen4` or `Gen5`.
	Family *string `pulumi:"family"`
	// Specifies the SKU Name for this Elasticpool. The name of the SKU, will be either `vCore` based `tier` + `family` pattern (e.g. GP_Gen4, BC_Gen5) or the `DTU` based `BasicPool`, `StandardPool`, or `PremiumPool` pattern.
	Name string `pulumi:"name"`
	// The tier of the particular SKU. Possible values are `GeneralPurpose`, `BusinessCritical`, `Basic`, `Standard`, or `Premium`. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
	Tier string `pulumi:"tier"`
}

type ElasticPoolSkuInput interface {
	pulumi.Input

	ToElasticPoolSkuOutput() ElasticPoolSkuOutput
	ToElasticPoolSkuOutputWithContext(context.Context) ElasticPoolSkuOutput
}

type ElasticPoolSkuArgs struct {
	// The scale up/out capacity, representing server's compute units. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
	Capacity pulumi.IntInput `pulumi:"capacity"`
	// The `family` of hardware `Gen4` or `Gen5`.
	Family pulumi.StringPtrInput `pulumi:"family"`
	// Specifies the SKU Name for this Elasticpool. The name of the SKU, will be either `vCore` based `tier` + `family` pattern (e.g. GP_Gen4, BC_Gen5) or the `DTU` based `BasicPool`, `StandardPool`, or `PremiumPool` pattern.
	Name pulumi.StringInput `pulumi:"name"`
	// The tier of the particular SKU. Possible values are `GeneralPurpose`, `BusinessCritical`, `Basic`, `Standard`, or `Premium`. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (ElasticPoolSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolSku)(nil)).Elem()
}

func (i ElasticPoolSkuArgs) ToElasticPoolSkuOutput() ElasticPoolSkuOutput {
	return i.ToElasticPoolSkuOutputWithContext(context.Background())
}

func (i ElasticPoolSkuArgs) ToElasticPoolSkuOutputWithContext(ctx context.Context) ElasticPoolSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolSkuOutput)
}

func (i ElasticPoolSkuArgs) ToElasticPoolSkuPtrOutput() ElasticPoolSkuPtrOutput {
	return i.ToElasticPoolSkuPtrOutputWithContext(context.Background())
}

func (i ElasticPoolSkuArgs) ToElasticPoolSkuPtrOutputWithContext(ctx context.Context) ElasticPoolSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolSkuOutput).ToElasticPoolSkuPtrOutputWithContext(ctx)
}

type ElasticPoolSkuPtrInput interface {
	pulumi.Input

	ToElasticPoolSkuPtrOutput() ElasticPoolSkuPtrOutput
	ToElasticPoolSkuPtrOutputWithContext(context.Context) ElasticPoolSkuPtrOutput
}

type elasticPoolSkuPtrType ElasticPoolSkuArgs

func ElasticPoolSkuPtr(v *ElasticPoolSkuArgs) ElasticPoolSkuPtrInput {	return (*elasticPoolSkuPtrType)(v)
}

func (*elasticPoolSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolSku)(nil)).Elem()
}

func (i *elasticPoolSkuPtrType) ToElasticPoolSkuPtrOutput() ElasticPoolSkuPtrOutput {
	return i.ToElasticPoolSkuPtrOutputWithContext(context.Background())
}

func (i *elasticPoolSkuPtrType) ToElasticPoolSkuPtrOutputWithContext(ctx context.Context) ElasticPoolSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolSkuPtrOutput)
}

type ElasticPoolSkuOutput struct { *pulumi.OutputState }

func (ElasticPoolSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolSku)(nil)).Elem()
}

func (o ElasticPoolSkuOutput) ToElasticPoolSkuOutput() ElasticPoolSkuOutput {
	return o
}

func (o ElasticPoolSkuOutput) ToElasticPoolSkuOutputWithContext(ctx context.Context) ElasticPoolSkuOutput {
	return o
}

func (o ElasticPoolSkuOutput) ToElasticPoolSkuPtrOutput() ElasticPoolSkuPtrOutput {
	return o.ToElasticPoolSkuPtrOutputWithContext(context.Background())
}

func (o ElasticPoolSkuOutput) ToElasticPoolSkuPtrOutputWithContext(ctx context.Context) ElasticPoolSkuPtrOutput {
	return o.ApplyT(func(v ElasticPoolSku) *ElasticPoolSku {
		return &v
	}).(ElasticPoolSkuPtrOutput)
}
// The scale up/out capacity, representing server's compute units. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
func (o ElasticPoolSkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func (v ElasticPoolSku) int { return v.Capacity }).(pulumi.IntOutput)
}

// The `family` of hardware `Gen4` or `Gen5`.
func (o ElasticPoolSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ElasticPoolSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

// Specifies the SKU Name for this Elasticpool. The name of the SKU, will be either `vCore` based `tier` + `family` pattern (e.g. GP_Gen4, BC_Gen5) or the `DTU` based `BasicPool`, `StandardPool`, or `PremiumPool` pattern.
func (o ElasticPoolSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v ElasticPoolSku) string { return v.Name }).(pulumi.StringOutput)
}

// The tier of the particular SKU. Possible values are `GeneralPurpose`, `BusinessCritical`, `Basic`, `Standard`, or `Premium`. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
func (o ElasticPoolSkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func (v ElasticPoolSku) string { return v.Tier }).(pulumi.StringOutput)
}

type ElasticPoolSkuPtrOutput struct { *pulumi.OutputState}

func (ElasticPoolSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolSku)(nil)).Elem()
}

func (o ElasticPoolSkuPtrOutput) ToElasticPoolSkuPtrOutput() ElasticPoolSkuPtrOutput {
	return o
}

func (o ElasticPoolSkuPtrOutput) ToElasticPoolSkuPtrOutputWithContext(ctx context.Context) ElasticPoolSkuPtrOutput {
	return o
}

func (o ElasticPoolSkuPtrOutput) Elem() ElasticPoolSkuOutput {
	return o.ApplyT(func (v *ElasticPoolSku) ElasticPoolSku { return *v }).(ElasticPoolSkuOutput)
}

// The scale up/out capacity, representing server's compute units. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
func (o ElasticPoolSkuPtrOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func (v ElasticPoolSku) int { return v.Capacity }).(pulumi.IntOutput)
}

// The `family` of hardware `Gen4` or `Gen5`.
func (o ElasticPoolSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ElasticPoolSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

// Specifies the SKU Name for this Elasticpool. The name of the SKU, will be either `vCore` based `tier` + `family` pattern (e.g. GP_Gen4, BC_Gen5) or the `DTU` based `BasicPool`, `StandardPool`, or `PremiumPool` pattern.
func (o ElasticPoolSkuPtrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v ElasticPoolSku) string { return v.Name }).(pulumi.StringOutput)
}

// The tier of the particular SKU. Possible values are `GeneralPurpose`, `BusinessCritical`, `Basic`, `Standard`, or `Premium`. For more information see the documentation for your Elasticpool configuration: [vCore-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-vcore-resource-limits-elastic-pools) or [DTU-based](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-dtu-resource-limits-elastic-pools).
func (o ElasticPoolSkuPtrOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func (v ElasticPoolSku) string { return v.Tier }).(pulumi.StringOutput)
}

type ServerVulnerabilityAssessmentRecurringScans struct {
	// Boolean flag which specifies if the schedule scan notification will be sent to the subscription administrators. Defaults to `false`.
	EmailSubscriptionAdmins *bool `pulumi:"emailSubscriptionAdmins"`
	// Specifies an array of e-mail addresses to which the scan notification is sent.
	Emails []string `pulumi:"emails"`
	// Boolean flag which specifies if recurring scans is enabled or disabled. Defaults to `false`.
	Enabled *bool `pulumi:"enabled"`
}

type ServerVulnerabilityAssessmentRecurringScansInput interface {
	pulumi.Input

	ToServerVulnerabilityAssessmentRecurringScansOutput() ServerVulnerabilityAssessmentRecurringScansOutput
	ToServerVulnerabilityAssessmentRecurringScansOutputWithContext(context.Context) ServerVulnerabilityAssessmentRecurringScansOutput
}

type ServerVulnerabilityAssessmentRecurringScansArgs struct {
	// Boolean flag which specifies if the schedule scan notification will be sent to the subscription administrators. Defaults to `false`.
	EmailSubscriptionAdmins pulumi.BoolPtrInput `pulumi:"emailSubscriptionAdmins"`
	// Specifies an array of e-mail addresses to which the scan notification is sent.
	Emails pulumi.StringArrayInput `pulumi:"emails"`
	// Boolean flag which specifies if recurring scans is enabled or disabled. Defaults to `false`.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (ServerVulnerabilityAssessmentRecurringScansArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerVulnerabilityAssessmentRecurringScans)(nil)).Elem()
}

func (i ServerVulnerabilityAssessmentRecurringScansArgs) ToServerVulnerabilityAssessmentRecurringScansOutput() ServerVulnerabilityAssessmentRecurringScansOutput {
	return i.ToServerVulnerabilityAssessmentRecurringScansOutputWithContext(context.Background())
}

func (i ServerVulnerabilityAssessmentRecurringScansArgs) ToServerVulnerabilityAssessmentRecurringScansOutputWithContext(ctx context.Context) ServerVulnerabilityAssessmentRecurringScansOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerVulnerabilityAssessmentRecurringScansOutput)
}

func (i ServerVulnerabilityAssessmentRecurringScansArgs) ToServerVulnerabilityAssessmentRecurringScansPtrOutput() ServerVulnerabilityAssessmentRecurringScansPtrOutput {
	return i.ToServerVulnerabilityAssessmentRecurringScansPtrOutputWithContext(context.Background())
}

func (i ServerVulnerabilityAssessmentRecurringScansArgs) ToServerVulnerabilityAssessmentRecurringScansPtrOutputWithContext(ctx context.Context) ServerVulnerabilityAssessmentRecurringScansPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerVulnerabilityAssessmentRecurringScansOutput).ToServerVulnerabilityAssessmentRecurringScansPtrOutputWithContext(ctx)
}

type ServerVulnerabilityAssessmentRecurringScansPtrInput interface {
	pulumi.Input

	ToServerVulnerabilityAssessmentRecurringScansPtrOutput() ServerVulnerabilityAssessmentRecurringScansPtrOutput
	ToServerVulnerabilityAssessmentRecurringScansPtrOutputWithContext(context.Context) ServerVulnerabilityAssessmentRecurringScansPtrOutput
}

type serverVulnerabilityAssessmentRecurringScansPtrType ServerVulnerabilityAssessmentRecurringScansArgs

func ServerVulnerabilityAssessmentRecurringScansPtr(v *ServerVulnerabilityAssessmentRecurringScansArgs) ServerVulnerabilityAssessmentRecurringScansPtrInput {	return (*serverVulnerabilityAssessmentRecurringScansPtrType)(v)
}

func (*serverVulnerabilityAssessmentRecurringScansPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerVulnerabilityAssessmentRecurringScans)(nil)).Elem()
}

func (i *serverVulnerabilityAssessmentRecurringScansPtrType) ToServerVulnerabilityAssessmentRecurringScansPtrOutput() ServerVulnerabilityAssessmentRecurringScansPtrOutput {
	return i.ToServerVulnerabilityAssessmentRecurringScansPtrOutputWithContext(context.Background())
}

func (i *serverVulnerabilityAssessmentRecurringScansPtrType) ToServerVulnerabilityAssessmentRecurringScansPtrOutputWithContext(ctx context.Context) ServerVulnerabilityAssessmentRecurringScansPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerVulnerabilityAssessmentRecurringScansPtrOutput)
}

type ServerVulnerabilityAssessmentRecurringScansOutput struct { *pulumi.OutputState }

func (ServerVulnerabilityAssessmentRecurringScansOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerVulnerabilityAssessmentRecurringScans)(nil)).Elem()
}

func (o ServerVulnerabilityAssessmentRecurringScansOutput) ToServerVulnerabilityAssessmentRecurringScansOutput() ServerVulnerabilityAssessmentRecurringScansOutput {
	return o
}

func (o ServerVulnerabilityAssessmentRecurringScansOutput) ToServerVulnerabilityAssessmentRecurringScansOutputWithContext(ctx context.Context) ServerVulnerabilityAssessmentRecurringScansOutput {
	return o
}

func (o ServerVulnerabilityAssessmentRecurringScansOutput) ToServerVulnerabilityAssessmentRecurringScansPtrOutput() ServerVulnerabilityAssessmentRecurringScansPtrOutput {
	return o.ToServerVulnerabilityAssessmentRecurringScansPtrOutputWithContext(context.Background())
}

func (o ServerVulnerabilityAssessmentRecurringScansOutput) ToServerVulnerabilityAssessmentRecurringScansPtrOutputWithContext(ctx context.Context) ServerVulnerabilityAssessmentRecurringScansPtrOutput {
	return o.ApplyT(func(v ServerVulnerabilityAssessmentRecurringScans) *ServerVulnerabilityAssessmentRecurringScans {
		return &v
	}).(ServerVulnerabilityAssessmentRecurringScansPtrOutput)
}
// Boolean flag which specifies if the schedule scan notification will be sent to the subscription administrators. Defaults to `false`.
func (o ServerVulnerabilityAssessmentRecurringScansOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ServerVulnerabilityAssessmentRecurringScans) *bool { return v.EmailSubscriptionAdmins }).(pulumi.BoolPtrOutput)
}

// Specifies an array of e-mail addresses to which the scan notification is sent.
func (o ServerVulnerabilityAssessmentRecurringScansOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ServerVulnerabilityAssessmentRecurringScans) []string { return v.Emails }).(pulumi.StringArrayOutput)
}

// Boolean flag which specifies if recurring scans is enabled or disabled. Defaults to `false`.
func (o ServerVulnerabilityAssessmentRecurringScansOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ServerVulnerabilityAssessmentRecurringScans) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type ServerVulnerabilityAssessmentRecurringScansPtrOutput struct { *pulumi.OutputState}

func (ServerVulnerabilityAssessmentRecurringScansPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerVulnerabilityAssessmentRecurringScans)(nil)).Elem()
}

func (o ServerVulnerabilityAssessmentRecurringScansPtrOutput) ToServerVulnerabilityAssessmentRecurringScansPtrOutput() ServerVulnerabilityAssessmentRecurringScansPtrOutput {
	return o
}

func (o ServerVulnerabilityAssessmentRecurringScansPtrOutput) ToServerVulnerabilityAssessmentRecurringScansPtrOutputWithContext(ctx context.Context) ServerVulnerabilityAssessmentRecurringScansPtrOutput {
	return o
}

func (o ServerVulnerabilityAssessmentRecurringScansPtrOutput) Elem() ServerVulnerabilityAssessmentRecurringScansOutput {
	return o.ApplyT(func (v *ServerVulnerabilityAssessmentRecurringScans) ServerVulnerabilityAssessmentRecurringScans { return *v }).(ServerVulnerabilityAssessmentRecurringScansOutput)
}

// Boolean flag which specifies if the schedule scan notification will be sent to the subscription administrators. Defaults to `false`.
func (o ServerVulnerabilityAssessmentRecurringScansPtrOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ServerVulnerabilityAssessmentRecurringScans) *bool { return v.EmailSubscriptionAdmins }).(pulumi.BoolPtrOutput)
}

// Specifies an array of e-mail addresses to which the scan notification is sent.
func (o ServerVulnerabilityAssessmentRecurringScansPtrOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ServerVulnerabilityAssessmentRecurringScans) []string { return v.Emails }).(pulumi.StringArrayOutput)
}

// Boolean flag which specifies if recurring scans is enabled or disabled. Defaults to `false`.
func (o ServerVulnerabilityAssessmentRecurringScansPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ServerVulnerabilityAssessmentRecurringScans) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultOutput{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineBaselineResultArrayOutput{})
	pulumi.RegisterOutputType(ElasticPoolElasticPoolPropertiesOutput{})
	pulumi.RegisterOutputType(ElasticPoolElasticPoolPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ElasticPoolPerDatabaseSettingsOutput{})
	pulumi.RegisterOutputType(ElasticPoolPerDatabaseSettingsPtrOutput{})
	pulumi.RegisterOutputType(ElasticPoolSkuOutput{})
	pulumi.RegisterOutputType(ElasticPoolSkuPtrOutput{})
	pulumi.RegisterOutputType(ServerVulnerabilityAssessmentRecurringScansOutput{})
	pulumi.RegisterOutputType(ServerVulnerabilityAssessmentRecurringScansPtrOutput{})
}