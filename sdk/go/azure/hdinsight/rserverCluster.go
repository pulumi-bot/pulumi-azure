// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a HDInsight RServer Cluster.
//
// !> **Note:** [HDInsight 3.6 is deprecated and will be retired on 2020-12-31 - HDInsight 4.0 no longer supports RServer Clusters](https://docs.microsoft.com/en-us/azure/hdinsight/hdinsight-component-versioning#available-versions) - as such this resource is deprecated and will be removed in the next major version of the Provider.
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
// 		_, err = hdinsight.NewRServerCluster(ctx, "exampleRServerCluster", &hdinsight.RServerClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ClusterVersion:    pulumi.String("3.6"),
// 			Tier:              pulumi.String("Standard"),
// 			Rstudio:           pulumi.Bool(true),
// 			Gateway: &hdinsight.RServerClusterGatewayArgs{
// 				Enabled:  pulumi.Bool(true),
// 				Username: pulumi.String("acctestusrgw"),
// 				Password: pulumi.String("Password123!"),
// 			},
// 			StorageAccounts: hdinsight.RServerClusterStorageAccountArray{
// 				&hdinsight.RServerClusterStorageAccountArgs{
// 					StorageContainerId: exampleContainer.ID(),
// 					StorageAccountKey:  exampleAccount.PrimaryAccessKey,
// 					IsDefault:          pulumi.Bool(true),
// 				},
// 			},
// 			Roles: &hdinsight.RServerClusterRolesArgs{
// 				HeadNode: &hdinsight.RServerClusterRolesHeadNodeArgs{
// 					VmSize:   pulumi.String("Standard_D3_v2"),
// 					Username: pulumi.String("acctestusrvm"),
// 					Password: pulumi.String("AccTestvdSC4daf986!"),
// 				},
// 				WorkerNode: &hdinsight.RServerClusterRolesWorkerNodeArgs{
// 					VmSize:              pulumi.String("Standard_D4_V2"),
// 					Username:            pulumi.String("acctestusrvm"),
// 					Password:            pulumi.String("AccTestvdSC4daf986!"),
// 					TargetInstanceCount: pulumi.Int(3),
// 				},
// 				ZookeeperNode: &hdinsight.RServerClusterRolesZookeeperNodeArgs{
// 					VmSize:   pulumi.String("Standard_D3_v2"),
// 					Username: pulumi.String("acctestusrvm"),
// 					Password: pulumi.String("AccTestvdSC4daf986!"),
// 				},
// 				EdgeNode: &hdinsight.RServerClusterRolesEdgeNodeArgs{
// 					VmSize:   pulumi.String("Standard_D3_v2"),
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
type RServerCluster struct {
	pulumi.CustomResourceState

	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`
	// The SSH Connectivity Endpoint for the Edge Node of the HDInsight RServer Cluster.
	EdgeSshEndpoint pulumi.StringOutput `pulumi:"edgeSshEndpoint"`
	// A `gateway` block as defined below.
	Gateway RServerClusterGatewayOutput `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight RServer Cluster.
	HttpsEndpoint pulumi.StringOutput `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name for this HDInsight RServer Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles RServerClusterRolesOutput `pulumi:"roles"`
	// Should R Studio community edition for RServer be installed? Changing this forces a new resource to be created.
	Rstudio pulumi.BoolOutput `pulumi:"rstudio"`
	// The SSH Connectivity Endpoint for this HDInsight RServer Cluster.
	SshEndpoint pulumi.StringOutput `pulumi:"sshEndpoint"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts RServerClusterStorageAccountArrayOutput `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight RServer Cluster.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight RServer Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringOutput    `pulumi:"tier"`
	TlsMinVersion pulumi.StringPtrOutput `pulumi:"tlsMinVersion"`
}

// NewRServerCluster registers a new resource with the given unique name, arguments, and options.
func NewRServerCluster(ctx *pulumi.Context,
	name string, args *RServerClusterArgs, opts ...pulumi.ResourceOption) (*RServerCluster, error) {
	if args == nil || args.ClusterVersion == nil {
		return nil, errors.New("missing required argument 'ClusterVersion'")
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
	if args == nil || args.Rstudio == nil {
		return nil, errors.New("missing required argument 'Rstudio'")
	}
	if args == nil || args.Tier == nil {
		return nil, errors.New("missing required argument 'Tier'")
	}
	if args == nil {
		args = &RServerClusterArgs{}
	}
	var resource RServerCluster
	err := ctx.RegisterResource("azure:hdinsight/rServerCluster:RServerCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRServerCluster gets an existing RServerCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRServerCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RServerClusterState, opts ...pulumi.ResourceOption) (*RServerCluster, error) {
	var resource RServerCluster
	err := ctx.ReadResource("azure:hdinsight/rServerCluster:RServerCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RServerCluster resources.
type rserverClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion *string `pulumi:"clusterVersion"`
	// The SSH Connectivity Endpoint for the Edge Node of the HDInsight RServer Cluster.
	EdgeSshEndpoint *string `pulumi:"edgeSshEndpoint"`
	// A `gateway` block as defined below.
	Gateway *RServerClusterGateway `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight RServer Cluster.
	HttpsEndpoint *string `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name for this HDInsight RServer Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles *RServerClusterRoles `pulumi:"roles"`
	// Should R Studio community edition for RServer be installed? Changing this forces a new resource to be created.
	Rstudio *bool `pulumi:"rstudio"`
	// The SSH Connectivity Endpoint for this HDInsight RServer Cluster.
	SshEndpoint *string `pulumi:"sshEndpoint"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []RServerClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight RServer Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight RServer Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          *string `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

type RServerClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringPtrInput
	// The SSH Connectivity Endpoint for the Edge Node of the HDInsight RServer Cluster.
	EdgeSshEndpoint pulumi.StringPtrInput
	// A `gateway` block as defined below.
	Gateway RServerClusterGatewayPtrInput
	// The HTTPS Connectivity Endpoint for this HDInsight RServer Cluster.
	HttpsEndpoint pulumi.StringPtrInput
	// Specifies the Azure Region which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name for this HDInsight RServer Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `roles` block as defined below.
	Roles RServerClusterRolesPtrInput
	// Should R Studio community edition for RServer be installed? Changing this forces a new resource to be created.
	Rstudio pulumi.BoolPtrInput
	// The SSH Connectivity Endpoint for this HDInsight RServer Cluster.
	SshEndpoint pulumi.StringPtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts RServerClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight RServer Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight RServer Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringPtrInput
	TlsMinVersion pulumi.StringPtrInput
}

func (RServerClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*rserverClusterState)(nil)).Elem()
}

type rserverClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion string `pulumi:"clusterVersion"`
	// A `gateway` block as defined below.
	Gateway RServerClusterGateway `pulumi:"gateway"`
	// Specifies the Azure Region which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name for this HDInsight RServer Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles RServerClusterRoles `pulumi:"roles"`
	// Should R Studio community edition for RServer be installed? Changing this forces a new resource to be created.
	Rstudio bool `pulumi:"rstudio"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []RServerClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight RServer Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight RServer Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          string  `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

// The set of arguments for constructing a RServerCluster resource.
type RServerClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringInput
	// A `gateway` block as defined below.
	Gateway RServerClusterGatewayInput
	// Specifies the Azure Region which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name for this HDInsight RServer Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `roles` block as defined below.
	Roles RServerClusterRolesInput
	// Should R Studio community edition for RServer be installed? Changing this forces a new resource to be created.
	Rstudio pulumi.BoolInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts RServerClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight RServer Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight RServer Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringInput
	TlsMinVersion pulumi.StringPtrInput
}

func (RServerClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rserverClusterArgs)(nil)).Elem()
}

type RServerClusterInput interface {
	pulumi.Input

	ToRServerClusterOutput() RServerClusterOutput
	ToRServerClusterOutputWithContext(ctx context.Context) RServerClusterOutput
}

func (RServerCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*RServerCluster)(nil)).Elem()
}

func (i RServerCluster) ToRServerClusterOutput() RServerClusterOutput {
	return i.ToRServerClusterOutputWithContext(context.Background())
}

func (i RServerCluster) ToRServerClusterOutputWithContext(ctx context.Context) RServerClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RServerClusterOutput)
}

type RServerClusterOutput struct {
	*pulumi.OutputState
}

func (RServerClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RServerClusterOutput)(nil)).Elem()
}

func (o RServerClusterOutput) ToRServerClusterOutput() RServerClusterOutput {
	return o
}

func (o RServerClusterOutput) ToRServerClusterOutputWithContext(ctx context.Context) RServerClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RServerClusterOutput{})
}
