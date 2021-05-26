// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing HDInsight Cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/hdinsight"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := hdinsight.GetCluster(ctx, &hdinsight.GetClusterArgs{
// 			Name:              "example",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("httpsEndpoint", example.HttpsEndpoint)
// 		return nil
// 	})
// }
// ```
func GetCluster(ctx *pulumi.Context, args *GetClusterArgs, opts ...pulumi.InvokeOption) (*GetClusterResult, error) {
	var rv GetClusterResult
	err := ctx.Invoke("azure:hdinsight/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type GetClusterArgs struct {
	// Specifies the name of this HDInsight Cluster.
	Name string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Cluster exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getCluster.
type GetClusterResult struct {
	// The version of HDInsights which is used on this HDInsight Cluster.
	ClusterVersion string `pulumi:"clusterVersion"`
	// A map of versions of software used on this HDInsights Cluster.
	ComponentVersions map[string]string `pulumi:"componentVersions"`
	// The SSH Endpoint of the Edge Node for this HDInsight Cluster, if an Edge Node exists.
	EdgeSshEndpoint string `pulumi:"edgeSshEndpoint"`
	// A `gateway` block as defined below.
	Gateways []GetClusterGateway `pulumi:"gateways"`
	// The HTTPS Endpoint for this HDInsight Cluster.
	HttpsEndpoint string `pulumi:"httpsEndpoint"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Kafka Rest Proxy Endpoint for this HDInsight Cluster.
	KafkaRestProxyEndpoint string `pulumi:"kafkaRestProxyEndpoint"`
	// The kind of HDInsight Cluster this is, such as a Spark or Storm cluster.
	Kind string `pulumi:"kind"`
	// The Azure Region in which this HDInsight Cluster exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SSH Endpoint for this HDInsight Cluster.
	SshEndpoint string `pulumi:"sshEndpoint"`
	// A map of tags assigned to the HDInsight Cluster.
	Tags map[string]string `pulumi:"tags"`
	// The SKU / Tier of this HDInsight Cluster.
	Tier string `pulumi:"tier"`
	// The minimal supported tls version.
	TlsMinVersion string `pulumi:"tlsMinVersion"`
}

func GetClusterApply(ctx *pulumi.Context, args GetClusterApplyInput, opts ...pulumi.InvokeOption) GetClusterResultOutput {
	return args.ToGetClusterApplyOutput().ApplyT(func(v GetClusterArgs) (GetClusterResult, error) {
		r, err := GetCluster(ctx, &v, opts...)
		return *r, err

	}).(GetClusterResultOutput)
}

// GetClusterApplyInput is an input type that accepts GetClusterApplyArgs and GetClusterApplyOutput values.
// You can construct a concrete instance of `GetClusterApplyInput` via:
//
//          GetClusterApplyArgs{...}
type GetClusterApplyInput interface {
	pulumi.Input

	ToGetClusterApplyOutput() GetClusterApplyOutput
	ToGetClusterApplyOutputWithContext(context.Context) GetClusterApplyOutput
}

// A collection of arguments for invoking getCluster.
type GetClusterApplyArgs struct {
	// Specifies the name of this HDInsight Cluster.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Cluster exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetClusterApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterArgs)(nil)).Elem()
}

func (i GetClusterApplyArgs) ToGetClusterApplyOutput() GetClusterApplyOutput {
	return i.ToGetClusterApplyOutputWithContext(context.Background())
}

func (i GetClusterApplyArgs) ToGetClusterApplyOutputWithContext(ctx context.Context) GetClusterApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClusterApplyOutput)
}

// A collection of arguments for invoking getCluster.
type GetClusterApplyOutput struct{ *pulumi.OutputState }

func (GetClusterApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterArgs)(nil)).Elem()
}

func (o GetClusterApplyOutput) ToGetClusterApplyOutput() GetClusterApplyOutput {
	return o
}

func (o GetClusterApplyOutput) ToGetClusterApplyOutputWithContext(ctx context.Context) GetClusterApplyOutput {
	return o
}

// Specifies the name of this HDInsight Cluster.
func (o GetClusterApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterArgs) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the name of the Resource Group in which this HDInsight Cluster exists.
func (o GetClusterApplyOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterArgs) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A collection of values returned by getCluster.
type GetClusterResultOutput struct{ *pulumi.OutputState }

func (GetClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterResult)(nil)).Elem()
}

func (o GetClusterResultOutput) ToGetClusterResultOutput() GetClusterResultOutput {
	return o
}

func (o GetClusterResultOutput) ToGetClusterResultOutputWithContext(ctx context.Context) GetClusterResultOutput {
	return o
}

// The version of HDInsights which is used on this HDInsight Cluster.
func (o GetClusterResultOutput) ClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.ClusterVersion }).(pulumi.StringOutput)
}

// A map of versions of software used on this HDInsights Cluster.
func (o GetClusterResultOutput) ComponentVersions() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetClusterResult) map[string]string { return v.ComponentVersions }).(pulumi.StringMapOutput)
}

// The SSH Endpoint of the Edge Node for this HDInsight Cluster, if an Edge Node exists.
func (o GetClusterResultOutput) EdgeSshEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.EdgeSshEndpoint }).(pulumi.StringOutput)
}

// A `gateway` block as defined below.
func (o GetClusterResultOutput) Gateways() GetClusterGatewayArrayOutput {
	return o.ApplyT(func(v GetClusterResult) []GetClusterGateway { return v.Gateways }).(GetClusterGatewayArrayOutput)
}

// The HTTPS Endpoint for this HDInsight Cluster.
func (o GetClusterResultOutput) HttpsEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.HttpsEndpoint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Kafka Rest Proxy Endpoint for this HDInsight Cluster.
func (o GetClusterResultOutput) KafkaRestProxyEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.KafkaRestProxyEndpoint }).(pulumi.StringOutput)
}

// The kind of HDInsight Cluster this is, such as a Spark or Storm cluster.
func (o GetClusterResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The Azure Region in which this HDInsight Cluster exists.
func (o GetClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetClusterResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The SSH Endpoint for this HDInsight Cluster.
func (o GetClusterResultOutput) SshEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.SshEndpoint }).(pulumi.StringOutput)
}

// A map of tags assigned to the HDInsight Cluster.
func (o GetClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The SKU / Tier of this HDInsight Cluster.
func (o GetClusterResultOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.Tier }).(pulumi.StringOutput)
}

// The minimal supported tls version.
func (o GetClusterResultOutput) TlsMinVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterResult) string { return v.TlsMinVersion }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterApplyOutput{})
	pulumi.RegisterOutputType(GetClusterResultOutput{})
}
