// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a HDInsight Kafka Cluster.
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
// 		_, err = hdinsight.NewKafkaCluster(ctx, "exampleKafkaCluster", &hdinsight.KafkaClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ClusterVersion:    pulumi.String("4.0"),
// 			Tier:              pulumi.String("Standard"),
// 			ComponentVersion: &hdinsight.KafkaClusterComponentVersionArgs{
// 				Kafka: pulumi.String("2.1"),
// 			},
// 			Gateway: &hdinsight.KafkaClusterGatewayArgs{
// 				Enabled:  pulumi.Bool(true),
// 				Username: pulumi.String("acctestusrgw"),
// 				Password: pulumi.String("Password123!"),
// 			},
// 			StorageAccounts: hdinsight.KafkaClusterStorageAccountArray{
// 				&hdinsight.KafkaClusterStorageAccountArgs{
// 					StorageContainerId: exampleContainer.ID(),
// 					StorageAccountKey:  exampleAccount.PrimaryAccessKey,
// 					IsDefault:          pulumi.Bool(true),
// 				},
// 			},
// 			Roles: &hdinsight.KafkaClusterRolesArgs{
// 				HeadNode: &hdinsight.KafkaClusterRolesHeadNodeArgs{
// 					VmSize:   pulumi.String("Standard_D3_V2"),
// 					Username: pulumi.String("acctestusrvm"),
// 					Password: pulumi.String("AccTestvdSC4daf986!"),
// 				},
// 				WorkerNode: &hdinsight.KafkaClusterRolesWorkerNodeArgs{
// 					VmSize:               pulumi.String("Standard_D3_V2"),
// 					Username:             pulumi.String("acctestusrvm"),
// 					Password:             pulumi.String("AccTestvdSC4daf986!"),
// 					NumberOfDisksPerNode: pulumi.Int(3),
// 					TargetInstanceCount:  pulumi.Int(3),
// 				},
// 				ZookeeperNode: &hdinsight.KafkaClusterRolesZookeeperNodeArgs{
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
type KafkaCluster struct {
	pulumi.CustomResourceState

	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion KafkaClusterComponentVersionOutput `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway KafkaClusterGatewayOutput `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight Kafka Cluster.
	HttpsEndpoint pulumi.StringOutput `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores KafkaClusterMetastoresPtrOutput `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor KafkaClusterMonitorPtrOutput `pulumi:"monitor"`
	// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles KafkaClusterRolesOutput `pulumi:"roles"`
	// The SSH Connectivity Endpoint for this HDInsight Kafka Cluster.
	SshEndpoint pulumi.StringOutput `pulumi:"sshEndpoint"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 KafkaClusterStorageAccountGen2PtrOutput `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts KafkaClusterStorageAccountArrayOutput `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringOutput    `pulumi:"tier"`
	TlsMinVersion pulumi.StringPtrOutput `pulumi:"tlsMinVersion"`
}

// NewKafkaCluster registers a new resource with the given unique name, arguments, and options.
func NewKafkaCluster(ctx *pulumi.Context,
	name string, args *KafkaClusterArgs, opts ...pulumi.ResourceOption) (*KafkaCluster, error) {
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
	var resource KafkaCluster
	err := ctx.RegisterResource("azure:hdinsight/kafkaCluster:KafkaCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaCluster gets an existing KafkaCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaClusterState, opts ...pulumi.ResourceOption) (*KafkaCluster, error) {
	var resource KafkaCluster
	err := ctx.ReadResource("azure:hdinsight/kafkaCluster:KafkaCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaCluster resources.
type kafkaClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion *string `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion *KafkaClusterComponentVersion `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway *KafkaClusterGateway `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight Kafka Cluster.
	HttpsEndpoint *string `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores *KafkaClusterMetastores `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor *KafkaClusterMonitor `pulumi:"monitor"`
	// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles *KafkaClusterRoles `pulumi:"roles"`
	// The SSH Connectivity Endpoint for this HDInsight Kafka Cluster.
	SshEndpoint *string `pulumi:"sshEndpoint"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 *KafkaClusterStorageAccountGen2 `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []KafkaClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          *string `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

type KafkaClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringPtrInput
	// A `componentVersion` block as defined below.
	ComponentVersion KafkaClusterComponentVersionPtrInput
	// A `gateway` block as defined below.
	Gateway KafkaClusterGatewayPtrInput
	// The HTTPS Connectivity Endpoint for this HDInsight Kafka Cluster.
	HttpsEndpoint pulumi.StringPtrInput
	// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `metastores` block as defined below.
	Metastores KafkaClusterMetastoresPtrInput
	// A `monitor` block as defined below.
	Monitor KafkaClusterMonitorPtrInput
	// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `roles` block as defined below.
	Roles KafkaClusterRolesPtrInput
	// The SSH Connectivity Endpoint for this HDInsight Kafka Cluster.
	SshEndpoint pulumi.StringPtrInput
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 KafkaClusterStorageAccountGen2PtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts KafkaClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringPtrInput
	TlsMinVersion pulumi.StringPtrInput
}

func (KafkaClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaClusterState)(nil)).Elem()
}

type kafkaClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion string `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion KafkaClusterComponentVersion `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway KafkaClusterGateway `pulumi:"gateway"`
	// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores *KafkaClusterMetastores `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor *KafkaClusterMonitor `pulumi:"monitor"`
	// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles KafkaClusterRoles `pulumi:"roles"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 *KafkaClusterStorageAccountGen2 `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []KafkaClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          string  `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

// The set of arguments for constructing a KafkaCluster resource.
type KafkaClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringInput
	// A `componentVersion` block as defined below.
	ComponentVersion KafkaClusterComponentVersionInput
	// A `gateway` block as defined below.
	Gateway KafkaClusterGatewayInput
	// Specifies the Azure Region which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `metastores` block as defined below.
	Metastores KafkaClusterMetastoresPtrInput
	// A `monitor` block as defined below.
	Monitor KafkaClusterMonitorPtrInput
	// Specifies the name for this HDInsight Kafka Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight Kafka Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `roles` block as defined below.
	Roles KafkaClusterRolesInput
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 KafkaClusterStorageAccountGen2PtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts KafkaClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight Kafka Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight Kafka Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringInput
	TlsMinVersion pulumi.StringPtrInput
}

func (KafkaClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaClusterArgs)(nil)).Elem()
}

type KafkaClusterInput interface {
	pulumi.Input

	ToKafkaClusterOutput() KafkaClusterOutput
	ToKafkaClusterOutputWithContext(ctx context.Context) KafkaClusterOutput
}

func (KafkaCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*KafkaCluster)(nil)).Elem()
}

func (i KafkaCluster) ToKafkaClusterOutput() KafkaClusterOutput {
	return i.ToKafkaClusterOutputWithContext(context.Background())
}

func (i KafkaCluster) ToKafkaClusterOutputWithContext(ctx context.Context) KafkaClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaClusterOutput)
}

type KafkaClusterOutput struct {
	*pulumi.OutputState
}

func (KafkaClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KafkaClusterOutput)(nil)).Elem()
}

func (o KafkaClusterOutput) ToKafkaClusterOutput() KafkaClusterOutput {
	return o
}

func (o KafkaClusterOutput) ToKafkaClusterOutputWithContext(ctx context.Context) KafkaClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KafkaClusterOutput{})
}
