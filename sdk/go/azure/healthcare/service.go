// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Healthcare Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewService(ctx, "example", &healthcare.ServiceArgs{
// 			AccessPolicyObjectIds: pulumi.StringArray{
// 				pulumi.String("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
// 			},
// 			AuthenticationConfiguration: &healthcare.ServiceAuthenticationConfigurationArgs{
// 				Audience:          pulumi.String("https://azurehealthcareapis.com/"),
// 				Authority:         pulumi.String(fmt.Sprintf("%v%v%v%v%v%v", "https://login.microsoftonline.com/", "$", "%", "7Bdata.azurerm_client_config.current.tenant_id", "%", "7D")),
// 				SmartProxyEnabled: pulumi.Bool(true),
// 			},
// 			CorsConfiguration: &healthcare.ServiceCorsConfigurationArgs{
// 				AllowCredentials: pulumi.Bool(true),
// 				AllowedHeaders: pulumi.StringArray{
// 					pulumi.String("x-tempo-*"),
// 					pulumi.String("x-tempo2-*"),
// 				},
// 				AllowedMethods: pulumi.StringArray{
// 					pulumi.String("GET"),
// 					pulumi.String("PUT"),
// 				},
// 				AllowedOrigins: pulumi.StringArray{
// 					pulumi.String("http://www.example.com"),
// 					pulumi.String("http://www.example2.com"),
// 				},
// 				MaxAgeInSeconds: pulumi.Int(500),
// 			},
// 			CosmosdbThroughput: pulumi.Int(2000),
// 			Kind:               pulumi.String("fhir-R4"),
// 			Location:           pulumi.String("westus2"),
// 			ResourceGroupName:  pulumi.String("sample-resource-group"),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("testenv"),
// 				"purpose":     pulumi.String("AcceptanceTests"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Service struct {
	pulumi.CustomResourceState

	AccessPolicyObjectIds pulumi.StringArrayOutput `pulumi:"accessPolicyObjectIds"`
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfiguration ServiceAuthenticationConfigurationOutput `pulumi:"authenticationConfiguration"`
	// A `corsConfiguration` block as defined below.
	CorsConfiguration ServiceCorsConfigurationOutput `pulumi:"corsConfiguration"`
	// The provisioned throughput for the backing database. Range of `400`-`1000`. Defaults to `400`.
	CosmosdbThroughput pulumi.IntPtrOutput `pulumi:"cosmosdbThroughput"`
	// The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Specifies the supported Azure Region where the Service should be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the service instance. Used for service endpoint, must be unique within the audience.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which to create the Service.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Service
	err := ctx.RegisterResource("azure:healthcare/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure:healthcare/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	AccessPolicyObjectIds []string `pulumi:"accessPolicyObjectIds"`
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfiguration *ServiceAuthenticationConfiguration `pulumi:"authenticationConfiguration"`
	// A `corsConfiguration` block as defined below.
	CorsConfiguration *ServiceCorsConfiguration `pulumi:"corsConfiguration"`
	// The provisioned throughput for the backing database. Range of `400`-`1000`. Defaults to `400`.
	CosmosdbThroughput *int `pulumi:"cosmosdbThroughput"`
	// The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
	Kind *string `pulumi:"kind"`
	// Specifies the supported Azure Region where the Service should be created.
	Location *string `pulumi:"location"`
	// The name of the service instance. Used for service endpoint, must be unique within the audience.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which to create the Service.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ServiceState struct {
	AccessPolicyObjectIds pulumi.StringArrayInput
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfiguration ServiceAuthenticationConfigurationPtrInput
	// A `corsConfiguration` block as defined below.
	CorsConfiguration ServiceCorsConfigurationPtrInput
	// The provisioned throughput for the backing database. Range of `400`-`1000`. Defaults to `400`.
	CosmosdbThroughput pulumi.IntPtrInput
	// The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
	Kind pulumi.StringPtrInput
	// Specifies the supported Azure Region where the Service should be created.
	Location pulumi.StringPtrInput
	// The name of the service instance. Used for service endpoint, must be unique within the audience.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which to create the Service.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	AccessPolicyObjectIds []string `pulumi:"accessPolicyObjectIds"`
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfiguration *ServiceAuthenticationConfiguration `pulumi:"authenticationConfiguration"`
	// A `corsConfiguration` block as defined below.
	CorsConfiguration *ServiceCorsConfiguration `pulumi:"corsConfiguration"`
	// The provisioned throughput for the backing database. Range of `400`-`1000`. Defaults to `400`.
	CosmosdbThroughput *int `pulumi:"cosmosdbThroughput"`
	// The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
	Kind *string `pulumi:"kind"`
	// Specifies the supported Azure Region where the Service should be created.
	Location *string `pulumi:"location"`
	// The name of the service instance. Used for service endpoint, must be unique within the audience.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which to create the Service.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	AccessPolicyObjectIds pulumi.StringArrayInput
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfiguration ServiceAuthenticationConfigurationPtrInput
	// A `corsConfiguration` block as defined below.
	CorsConfiguration ServiceCorsConfigurationPtrInput
	// The provisioned throughput for the backing database. Range of `400`-`1000`. Defaults to `400`.
	CosmosdbThroughput pulumi.IntPtrInput
	// The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
	Kind pulumi.StringPtrInput
	// Specifies the supported Azure Region where the Service should be created.
	Location pulumi.StringPtrInput
	// The name of the service instance. Used for service endpoint, must be unique within the audience.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which to create the Service.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (Service) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil)).Elem()
}

func (i Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct {
	*pulumi.OutputState
}

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceOutput)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
