// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a HDInsight Hadoop Cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/hdinsight"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName:  exampleAccount.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = hdinsight.NewHadoopCluster(ctx, "exampleHadoopCluster", &hdinsight.HadoopClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ClusterVersion:    pulumi.String("3.6"),
// 			Tier:              pulumi.String("Standard"),
// 			ComponentVersion: &hdinsight.HadoopClusterComponentVersionArgs{
// 				Hadoop: pulumi.String("2.7"),
// 			},
// 			Gateway: &hdinsight.HadoopClusterGatewayArgs{
// 				Enabled:  pulumi.Bool(true),
// 				Username: pulumi.String("acctestusrgw"),
// 				Password: pulumi.String("PAssword123!"),
// 			},
// 			StorageAccounts: hdinsight.HadoopClusterStorageAccountArray{
// 				&hdinsight.HadoopClusterStorageAccountArgs{
// 					StorageContainerId: exampleContainer.ID(),
// 					StorageAccountKey:  exampleAccount.PrimaryAccessKey,
// 					IsDefault:          pulumi.Bool(true),
// 				},
// 			},
// 			Roles: &hdinsight.HadoopClusterRolesArgs{
// 				HeadNode: &hdinsight.HadoopClusterRolesHeadNodeArgs{
// 					VmSize:   pulumi.String("Standard_D3_V2"),
// 					Username: pulumi.String("acctestusrvm"),
// 					Password: pulumi.String("AccTestvdSC4daf986!"),
// 				},
// 				WorkerNode: &hdinsight.HadoopClusterRolesWorkerNodeArgs{
// 					VmSize:              pulumi.String("Standard_D4_V2"),
// 					Username:            pulumi.String("acctestusrvm"),
// 					Password:            pulumi.String("AccTestvdSC4daf986!"),
// 					TargetInstanceCount: pulumi.Int(3),
// 				},
// 				ZookeeperNode: &hdinsight.HadoopClusterRolesZookeeperNodeArgs{
// 					VmSize:   pulumi.String("Standard_D3_V2"),
// 					Username: pulumi.String("acctestusrvm"),
// 					Password: pulumi.String("AccTestvdSC4daf986!"),
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
//
// ## Import
//
// HDInsight Hadoop Clusters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:hdinsight/hadoopCluster:HadoopCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
// ```
type HadoopCluster struct {
	pulumi.CustomResourceState

	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion HadoopClusterComponentVersionOutput `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway HadoopClusterGatewayOutput `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight Hadoop Cluster.
	HttpsEndpoint pulumi.StringOutput `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores HadoopClusterMetastoresPtrOutput `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor HadoopClusterMonitorPtrOutput `pulumi:"monitor"`
	// Specifies the name for this HDInsight Hadoop Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles HadoopClusterRolesOutput `pulumi:"roles"`
	// The SSH Connectivity Endpoint for this HDInsight Hadoop Cluster.
	SshEndpoint pulumi.StringOutput `pulumi:"sshEndpoint"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 HadoopClusterStorageAccountGen2PtrOutput `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts HadoopClusterStorageAccountArrayOutput `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Hadoop Cluster.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Hadoop Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringOutput    `pulumi:"tier"`
	TlsMinVersion pulumi.StringPtrOutput `pulumi:"tlsMinVersion"`
}

// NewHadoopCluster registers a new resource with the given unique name, arguments, and options.
func NewHadoopCluster(ctx *pulumi.Context,
	name string, args *HadoopClusterArgs, opts ...pulumi.ResourceOption) (*HadoopCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterVersion == nil {
		return nil, errors.New("invalid value for required argument 'ClusterVersion'")
	}
	if args.ComponentVersion == nil {
		return nil, errors.New("invalid value for required argument 'ComponentVersion'")
	}
	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	var resource HadoopCluster
	err := ctx.RegisterResource("azure:hdinsight/hadoopCluster:HadoopCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHadoopCluster gets an existing HadoopCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHadoopCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HadoopClusterState, opts ...pulumi.ResourceOption) (*HadoopCluster, error) {
	var resource HadoopCluster
	err := ctx.ReadResource("azure:hdinsight/hadoopCluster:HadoopCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HadoopCluster resources.
type hadoopClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion *string `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion *HadoopClusterComponentVersion `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway *HadoopClusterGateway `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight Hadoop Cluster.
	HttpsEndpoint *string `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores *HadoopClusterMetastores `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor *HadoopClusterMonitor `pulumi:"monitor"`
	// Specifies the name for this HDInsight Hadoop Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles *HadoopClusterRoles `pulumi:"roles"`
	// The SSH Connectivity Endpoint for this HDInsight Hadoop Cluster.
	SshEndpoint *string `pulumi:"sshEndpoint"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 *HadoopClusterStorageAccountGen2 `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []HadoopClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Hadoop Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Hadoop Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          *string `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

type HadoopClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringPtrInput
	// A `componentVersion` block as defined below.
	ComponentVersion HadoopClusterComponentVersionPtrInput
	// A `gateway` block as defined below.
	Gateway HadoopClusterGatewayPtrInput
	// The HTTPS Connectivity Endpoint for this HDInsight Hadoop Cluster.
	HttpsEndpoint pulumi.StringPtrInput
	// Specifies the Azure Region which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `metastores` block as defined below.
	Metastores HadoopClusterMetastoresPtrInput
	// A `monitor` block as defined below.
	Monitor HadoopClusterMonitorPtrInput
	// Specifies the name for this HDInsight Hadoop Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `roles` block as defined below.
	Roles HadoopClusterRolesPtrInput
	// The SSH Connectivity Endpoint for this HDInsight Hadoop Cluster.
	SshEndpoint pulumi.StringPtrInput
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 HadoopClusterStorageAccountGen2PtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts HadoopClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight Hadoop Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight Hadoop Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringPtrInput
	TlsMinVersion pulumi.StringPtrInput
}

func (HadoopClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*hadoopClusterState)(nil)).Elem()
}

type hadoopClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion string `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion HadoopClusterComponentVersion `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway HadoopClusterGateway `pulumi:"gateway"`
	// Specifies the Azure Region which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores *HadoopClusterMetastores `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor *HadoopClusterMonitor `pulumi:"monitor"`
	// Specifies the name for this HDInsight Hadoop Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles HadoopClusterRoles `pulumi:"roles"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 *HadoopClusterStorageAccountGen2 `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []HadoopClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Hadoop Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Hadoop Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          string  `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

// The set of arguments for constructing a HadoopCluster resource.
type HadoopClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringInput
	// A `componentVersion` block as defined below.
	ComponentVersion HadoopClusterComponentVersionInput
	// A `gateway` block as defined below.
	Gateway HadoopClusterGatewayInput
	// Specifies the Azure Region which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `metastores` block as defined below.
	Metastores HadoopClusterMetastoresPtrInput
	// A `monitor` block as defined below.
	Monitor HadoopClusterMonitorPtrInput
	// Specifies the name for this HDInsight Hadoop Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight Hadoop Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `roles` block as defined below.
	Roles HadoopClusterRolesInput
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 HadoopClusterStorageAccountGen2PtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts HadoopClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight Hadoop Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight Hadoop Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringInput
	TlsMinVersion pulumi.StringPtrInput
}

func (HadoopClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hadoopClusterArgs)(nil)).Elem()
}

type HadoopClusterInput interface {
	pulumi.Input

	ToHadoopClusterOutput() HadoopClusterOutput
	ToHadoopClusterOutputWithContext(ctx context.Context) HadoopClusterOutput
}

func (*HadoopCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*HadoopCluster)(nil))
}

func (i *HadoopCluster) ToHadoopClusterOutput() HadoopClusterOutput {
	return i.ToHadoopClusterOutputWithContext(context.Background())
}

func (i *HadoopCluster) ToHadoopClusterOutputWithContext(ctx context.Context) HadoopClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HadoopClusterOutput)
}

type HadoopClusterOutput struct {
	*pulumi.OutputState
}

func (HadoopClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HadoopCluster)(nil))
}

func (o HadoopClusterOutput) ToHadoopClusterOutput() HadoopClusterOutput {
	return o
}

func (o HadoopClusterOutput) ToHadoopClusterOutputWithContext(ctx context.Context) HadoopClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HadoopClusterOutput{})
}
