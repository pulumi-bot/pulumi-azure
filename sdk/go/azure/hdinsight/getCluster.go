// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing HDInsight Cluster.
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
