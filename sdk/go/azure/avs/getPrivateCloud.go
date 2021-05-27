// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateCloud(ctx *pulumi.Context, args *LookupPrivateCloudArgs, opts ...pulumi.InvokeOption) (*LookupPrivateCloudResult, error) {
	var rv LookupPrivateCloudResult
	err := ctx.Invoke("azure:avs/getPrivateCloud:getPrivateCloud", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateCloud.
type LookupPrivateCloudArgs struct {
	// The name of this Vmware Private Cloud.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Vmware Private Cloud exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getPrivateCloud.
type LookupPrivateCloudResult struct {
	// A `circuit` block as defined below.
	Circuits []GetPrivateCloudCircuit `pulumi:"circuits"`
	// The endpoint for the HCX Cloud Manager.
	HcxCloudManagerEndpoint string `pulumi:"hcxCloudManagerEndpoint"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Is the Vmware Private Cluster connected to the internet?
	InternetConnectionEnabled bool `pulumi:"internetConnectionEnabled"`
	// The Azure Region where the Vmware Private Cloud exists.
	Location string `pulumi:"location"`
	// A `managementCluster` block as defined below.
	ManagementClusters []GetPrivateCloudManagementCluster `pulumi:"managementClusters"`
	// The network used to access vCenter Server and NSX-T Manager.
	ManagementSubnetCidr string `pulumi:"managementSubnetCidr"`
	Name                 string `pulumi:"name"`
	// The subnet cidr of the Vmware Private Cloud.
	NetworkSubnetCidr string `pulumi:"networkSubnetCidr"`
	// The thumbprint of the NSX-T Manager SSL certificate.
	NsxtCertificateThumbprint string `pulumi:"nsxtCertificateThumbprint"`
	// The endpoint for the NSX-T Data Center manager.
	NsxtManagerEndpoint string `pulumi:"nsxtManagerEndpoint"`
	// The network which isused for virtual machine cold migration, cloning, and snapshot migration.
	ProvisioningSubnetCidr string `pulumi:"provisioningSubnetCidr"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	// The Name of the SKU used for this Private Cloud.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags assigned to the Vmware Private Cloud.
	Tags map[string]string `pulumi:"tags"`
	// The thumbprint of the vCenter Server SSL certificate.
	VcenterCertificateThumbprint string `pulumi:"vcenterCertificateThumbprint"`
	// The endpoint for Virtual Center Server Appliance.
	VcsaEndpoint string `pulumi:"vcsaEndpoint"`
	// The network which is used for live migration of virtual machines.
	VmotionSubnetCidr string `pulumi:"vmotionSubnetCidr"`
}

func LookupPrivateCloudOutput(ctx *pulumi.Context, args LookupPrivateCloudOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateCloudResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateCloudResult, error) {
			args := v.(LookupPrivateCloudArgs)
			r, err := LookupPrivateCloud(ctx, &args, opts...)
			return *r, err
		}).(LookupPrivateCloudResultOutput)
}

// A collection of arguments for invoking getPrivateCloud.
type LookupPrivateCloudOutputArgs struct {
	// The name of this Vmware Private Cloud.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group where the Vmware Private Cloud exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateCloudOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateCloudArgs)(nil)).Elem()
}

// A collection of values returned by getPrivateCloud.
type LookupPrivateCloudResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateCloudResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateCloudResult)(nil)).Elem()
}

func (o LookupPrivateCloudResultOutput) ToLookupPrivateCloudResultOutput() LookupPrivateCloudResultOutput {
	return o
}

func (o LookupPrivateCloudResultOutput) ToLookupPrivateCloudResultOutputWithContext(ctx context.Context) LookupPrivateCloudResultOutput {
	return o
}

// A `circuit` block as defined below.
func (o LookupPrivateCloudResultOutput) Circuits() GetPrivateCloudCircuitArrayOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) []GetPrivateCloudCircuit { return v.Circuits }).(GetPrivateCloudCircuitArrayOutput)
}

// The endpoint for the HCX Cloud Manager.
func (o LookupPrivateCloudResultOutput) HcxCloudManagerEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.HcxCloudManagerEndpoint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrivateCloudResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.Id }).(pulumi.StringOutput)
}

// Is the Vmware Private Cluster connected to the internet?
func (o LookupPrivateCloudResultOutput) InternetConnectionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) bool { return v.InternetConnectionEnabled }).(pulumi.BoolOutput)
}

// The Azure Region where the Vmware Private Cloud exists.
func (o LookupPrivateCloudResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.Location }).(pulumi.StringOutput)
}

// A `managementCluster` block as defined below.
func (o LookupPrivateCloudResultOutput) ManagementClusters() GetPrivateCloudManagementClusterArrayOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) []GetPrivateCloudManagementCluster { return v.ManagementClusters }).(GetPrivateCloudManagementClusterArrayOutput)
}

// The network used to access vCenter Server and NSX-T Manager.
func (o LookupPrivateCloudResultOutput) ManagementSubnetCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.ManagementSubnetCidr }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.Name }).(pulumi.StringOutput)
}

// The subnet cidr of the Vmware Private Cloud.
func (o LookupPrivateCloudResultOutput) NetworkSubnetCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.NetworkSubnetCidr }).(pulumi.StringOutput)
}

// The thumbprint of the NSX-T Manager SSL certificate.
func (o LookupPrivateCloudResultOutput) NsxtCertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.NsxtCertificateThumbprint }).(pulumi.StringOutput)
}

// The endpoint for the NSX-T Data Center manager.
func (o LookupPrivateCloudResultOutput) NsxtManagerEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.NsxtManagerEndpoint }).(pulumi.StringOutput)
}

// The network which isused for virtual machine cold migration, cloning, and snapshot migration.
func (o LookupPrivateCloudResultOutput) ProvisioningSubnetCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.ProvisioningSubnetCidr }).(pulumi.StringOutput)
}

func (o LookupPrivateCloudResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The Name of the SKU used for this Private Cloud.
func (o LookupPrivateCloudResultOutput) SkuName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.SkuName }).(pulumi.StringOutput)
}

// A mapping of tags assigned to the Vmware Private Cloud.
func (o LookupPrivateCloudResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The thumbprint of the vCenter Server SSL certificate.
func (o LookupPrivateCloudResultOutput) VcenterCertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.VcenterCertificateThumbprint }).(pulumi.StringOutput)
}

// The endpoint for Virtual Center Server Appliance.
func (o LookupPrivateCloudResultOutput) VcsaEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.VcsaEndpoint }).(pulumi.StringOutput)
}

// The network which is used for live migration of virtual machines.
func (o LookupPrivateCloudResultOutput) VmotionSubnetCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateCloudResult) string { return v.VmotionSubnetCidr }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateCloudResultOutput{})
}
