// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a HDInsight HBase Cluster.
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
// 		_, err = hdinsight.NewHBaseCluster(ctx, "exampleHBaseCluster", &hdinsight.HBaseClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ClusterVersion:    pulumi.String("3.6"),
// 			Tier:              pulumi.String("Standard"),
// 			ComponentVersion: &hdinsight.HBaseClusterComponentVersionArgs{
// 				Hbase: pulumi.String("1.1"),
// 			},
// 			Gateway: &hdinsight.HBaseClusterGatewayArgs{
// 				Enabled:  pulumi.Bool(true),
// 				Username: pulumi.String("acctestusrgw"),
// 				Password: pulumi.String("Password123!"),
// 			},
// 			StorageAccounts: hdinsight.HBaseClusterStorageAccountArray{
// 				&hdinsight.HBaseClusterStorageAccountArgs{
// 					StorageContainerId: exampleContainer.ID(),
// 					StorageAccountKey:  exampleAccount.PrimaryAccessKey,
// 					IsDefault:          pulumi.Bool(true),
// 				},
// 			},
// 			Roles: &hdinsight.HBaseClusterRolesArgs{
// 				HeadNode: &hdinsight.HBaseClusterRolesHeadNodeArgs{
// 					VmSize:   pulumi.String("Standard_D3_V2"),
// 					Username: pulumi.String("acctestusrvm"),
// 					Password: pulumi.String("AccTestvdSC4daf986!"),
// 				},
// 				WorkerNode: &hdinsight.HBaseClusterRolesWorkerNodeArgs{
// 					VmSize:              pulumi.String("Standard_D3_V2"),
// 					Username:            pulumi.String("acctestusrvm"),
// 					Password:            pulumi.String("AccTestvdSC4daf986!"),
// 					TargetInstanceCount: pulumi.Int(3),
// 				},
// 				ZookeeperNode: &hdinsight.HBaseClusterRolesZookeeperNodeArgs{
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
type HBaseCluster struct {
	pulumi.CustomResourceState

	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion HBaseClusterComponentVersionOutput `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway HBaseClusterGatewayOutput `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight HBase Cluster.
	HttpsEndpoint pulumi.StringOutput `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores HBaseClusterMetastoresPtrOutput `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor HBaseClusterMonitorPtrOutput `pulumi:"monitor"`
	// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles HBaseClusterRolesOutput `pulumi:"roles"`
	// The SSH Connectivity Endpoint for this HDInsight HBase Cluster.
	SshEndpoint pulumi.StringOutput `pulumi:"sshEndpoint"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 HBaseClusterStorageAccountGen2PtrOutput `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts HBaseClusterStorageAccountArrayOutput `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight HBase Cluster.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringOutput    `pulumi:"tier"`
	TlsMinVersion pulumi.StringPtrOutput `pulumi:"tlsMinVersion"`
}

// NewHBaseCluster registers a new resource with the given unique name, arguments, and options.
func NewHBaseCluster(ctx *pulumi.Context,
	name string, args *HBaseClusterArgs, opts ...pulumi.ResourceOption) (*HBaseCluster, error) {
	if args == nil || args.ClusterVersion == nil {
		return nil, errors.New("missing required argument 'ClusterVersion'")
	}
	if args == nil || args.ComponentVersion == nil {
		return nil, errors.New("missing required argument 'ComponentVersion'")
	}
	if args == nil || args.Gateway == nil {
		return nil, errors.New("missing required argument 'Gateway'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Roles == nil {
		return nil, errors.New("missing required argument 'Roles'")
	}
	if args == nil || args.Tier == nil {
		return nil, errors.New("missing required argument 'Tier'")
	}
	if args == nil {
		args = &HBaseClusterArgs{}
	}
	var resource HBaseCluster
	err := ctx.RegisterResource("azure:hdinsight/hBaseCluster:HBaseCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHBaseCluster gets an existing HBaseCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHBaseCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HBaseClusterState, opts ...pulumi.ResourceOption) (*HBaseCluster, error) {
	var resource HBaseCluster
	err := ctx.ReadResource("azure:hdinsight/hBaseCluster:HBaseCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HBaseCluster resources.
type hbaseClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion *string `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion *HBaseClusterComponentVersion `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway *HBaseClusterGateway `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight HBase Cluster.
	HttpsEndpoint *string `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores *HBaseClusterMetastores `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor *HBaseClusterMonitor `pulumi:"monitor"`
	// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles *HBaseClusterRoles `pulumi:"roles"`
	// The SSH Connectivity Endpoint for this HDInsight HBase Cluster.
	SshEndpoint *string `pulumi:"sshEndpoint"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 *HBaseClusterStorageAccountGen2 `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []HBaseClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight HBase Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          *string `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

type HBaseClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringPtrInput
	// A `componentVersion` block as defined below.
	ComponentVersion HBaseClusterComponentVersionPtrInput
	// A `gateway` block as defined below.
	Gateway HBaseClusterGatewayPtrInput
	// The HTTPS Connectivity Endpoint for this HDInsight HBase Cluster.
	HttpsEndpoint pulumi.StringPtrInput
	// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `metastores` block as defined below.
	Metastores HBaseClusterMetastoresPtrInput
	// A `monitor` block as defined below.
	Monitor HBaseClusterMonitorPtrInput
	// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `roles` block as defined below.
	Roles HBaseClusterRolesPtrInput
	// The SSH Connectivity Endpoint for this HDInsight HBase Cluster.
	SshEndpoint pulumi.StringPtrInput
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 HBaseClusterStorageAccountGen2PtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts HBaseClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight HBase Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringPtrInput
	TlsMinVersion pulumi.StringPtrInput
}

func (HBaseClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*hbaseClusterState)(nil)).Elem()
}

type hbaseClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion string `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion HBaseClusterComponentVersion `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway HBaseClusterGateway `pulumi:"gateway"`
	// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores *HBaseClusterMetastores `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor *HBaseClusterMonitor `pulumi:"monitor"`
	// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles HBaseClusterRoles `pulumi:"roles"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 *HBaseClusterStorageAccountGen2 `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []HBaseClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight HBase Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          string  `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

// The set of arguments for constructing a HBaseCluster resource.
type HBaseClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringInput
	// A `componentVersion` block as defined below.
	ComponentVersion HBaseClusterComponentVersionInput
	// A `gateway` block as defined below.
	Gateway HBaseClusterGatewayInput
	// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `metastores` block as defined below.
	Metastores HBaseClusterMetastoresPtrInput
	// A `monitor` block as defined below.
	Monitor HBaseClusterMonitorPtrInput
	// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `roles` block as defined below.
	Roles HBaseClusterRolesInput
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 HBaseClusterStorageAccountGen2PtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts HBaseClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight HBase Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringInput
	TlsMinVersion pulumi.StringPtrInput
}

func (HBaseClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hbaseClusterArgs)(nil)).Elem()
}

type HBaseClusterInput interface {
	pulumi.Input

	ToHBaseClusterOutput() HBaseClusterOutput
	ToHBaseClusterOutputWithContext(ctx context.Context) HBaseClusterOutput
}

func (HBaseCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*HBaseCluster)(nil)).Elem()
}

func (i HBaseCluster) ToHBaseClusterOutput() HBaseClusterOutput {
	return i.ToHBaseClusterOutputWithContext(context.Background())
}

func (i HBaseCluster) ToHBaseClusterOutputWithContext(ctx context.Context) HBaseClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HBaseClusterOutput)
}

type HBaseClusterOutput struct {
	*pulumi.OutputState
}

func (HBaseClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HBaseClusterOutput)(nil)).Elem()
}

func (o HBaseClusterOutput) ToHBaseClusterOutput() HBaseClusterOutput {
	return o
}

func (o HBaseClusterOutput) ToHBaseClusterOutputWithContext(ctx context.Context) HBaseClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HBaseClusterOutput{})
}
