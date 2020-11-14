// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Application Insights API key.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
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
// 		exampleInsights, err := appinsights.NewInsights(ctx, "exampleInsights", &appinsights.InsightsArgs{
// 			Location:          pulumi.String("West Europe"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationType:   pulumi.String("web"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		readTelemetry, err := appinsights.NewApiKey(ctx, "readTelemetry", &appinsights.ApiKeyArgs{
// 			ApplicationInsightsId: exampleInsights.ID(),
// 			ReadPermissions: pulumi.StringArray{
// 				pulumi.String("aggregate"),
// 				pulumi.String("api"),
// 				pulumi.String("draft"),
// 				pulumi.String("extendqueries"),
// 				pulumi.String("search"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		writeAnnotations, err := appinsights.NewApiKey(ctx, "writeAnnotations", &appinsights.ApiKeyArgs{
// 			ApplicationInsightsId: exampleInsights.ID(),
// 			WritePermissions: pulumi.StringArray{
// 				pulumi.String("annotations"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		authenticateSdkControlChannelApiKey, err := appinsights.NewApiKey(ctx, "authenticateSdkControlChannelApiKey", &appinsights.ApiKeyArgs{
// 			ApplicationInsightsId: exampleInsights.ID(),
// 			ReadPermissions: pulumi.StringArray{
// 				pulumi.String("agentconfig"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fullPermissions, err := appinsights.NewApiKey(ctx, "fullPermissions", &appinsights.ApiKeyArgs{
// 			ApplicationInsightsId: exampleInsights.ID(),
// 			ReadPermissions: pulumi.StringArray{
// 				pulumi.String("agentconfig"),
// 				pulumi.String("aggregate"),
// 				pulumi.String("api"),
// 				pulumi.String("draft"),
// 				pulumi.String("extendqueries"),
// 				pulumi.String("search"),
// 			},
// 			WritePermissions: pulumi.StringArray{
// 				pulumi.String("annotations"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("readTelemetryApiKey", readTelemetry.ApiKey)
// 		ctx.Export("writeAnnotationsApiKey", writeAnnotations.ApiKey)
// 		ctx.Export("authenticateSdkControlChannel", authenticateSdkControlChannelApiKey.ApiKey)
// 		ctx.Export("fullPermissionsApiKey", fullPermissions.ApiKey)
// 		return nil
// 	})
// }
// ```
type ApiKey struct {
	pulumi.CustomResourceState

	// The API Key secret (Sensitive).
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// The ID of the Application Insights component on which the API key operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringOutput `pulumi:"applicationInsightsId"`
	// Specifies the name of the Application Insights API key. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the list of read permissions granted to the API key. Valid values are `agentconfig`, `aggregate`, `api`, `draft`, `extendqueries`, `search`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	ReadPermissions pulumi.StringArrayOutput `pulumi:"readPermissions"`
	// Specifies the list of write permissions granted to the API key. Valid values are `annotations`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	WritePermissions pulumi.StringArrayOutput `pulumi:"writePermissions"`
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationInsightsId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationInsightsId'")
	}
	var resource ApiKey
	err := ctx.RegisterResource("azure:appinsights/apiKey:ApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiKeyState, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	var resource ApiKey
	err := ctx.ReadResource("azure:appinsights/apiKey:ApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiKey resources.
type apiKeyState struct {
	// The API Key secret (Sensitive).
	ApiKey *string `pulumi:"apiKey"`
	// The ID of the Application Insights component on which the API key operates. Changing this forces a new resource to be created.
	ApplicationInsightsId *string `pulumi:"applicationInsightsId"`
	// Specifies the name of the Application Insights API key. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the list of read permissions granted to the API key. Valid values are `agentconfig`, `aggregate`, `api`, `draft`, `extendqueries`, `search`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	ReadPermissions []string `pulumi:"readPermissions"`
	// Specifies the list of write permissions granted to the API key. Valid values are `annotations`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	WritePermissions []string `pulumi:"writePermissions"`
}

type ApiKeyState struct {
	// The API Key secret (Sensitive).
	ApiKey pulumi.StringPtrInput
	// The ID of the Application Insights component on which the API key operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringPtrInput
	// Specifies the name of the Application Insights API key. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the list of read permissions granted to the API key. Valid values are `agentconfig`, `aggregate`, `api`, `draft`, `extendqueries`, `search`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	ReadPermissions pulumi.StringArrayInput
	// Specifies the list of write permissions granted to the API key. Valid values are `annotations`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	WritePermissions pulumi.StringArrayInput
}

func (ApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyState)(nil)).Elem()
}

type apiKeyArgs struct {
	// The ID of the Application Insights component on which the API key operates. Changing this forces a new resource to be created.
	ApplicationInsightsId string `pulumi:"applicationInsightsId"`
	// Specifies the name of the Application Insights API key. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the list of read permissions granted to the API key. Valid values are `agentconfig`, `aggregate`, `api`, `draft`, `extendqueries`, `search`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	ReadPermissions []string `pulumi:"readPermissions"`
	// Specifies the list of write permissions granted to the API key. Valid values are `annotations`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	WritePermissions []string `pulumi:"writePermissions"`
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	// The ID of the Application Insights component on which the API key operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringInput
	// Specifies the name of the Application Insights API key. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the list of read permissions granted to the API key. Valid values are `agentconfig`, `aggregate`, `api`, `draft`, `extendqueries`, `search`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	ReadPermissions pulumi.StringArrayInput
	// Specifies the list of write permissions granted to the API key. Valid values are `annotations`. Please note these values are case sensitive. Changing this forces a new resource to be created.
	WritePermissions pulumi.StringArrayInput
}

func (ApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyArgs)(nil)).Elem()
}

type ApiKeyInput interface {
	pulumi.Input

	ToApiKeyOutput() ApiKeyOutput
	ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput
}

func (ApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiKey)(nil)).Elem()
}

func (i ApiKey) ToApiKeyOutput() ApiKeyOutput {
	return i.ToApiKeyOutputWithContext(context.Background())
}

func (i ApiKey) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyOutput)
}

type ApiKeyOutput struct {
	*pulumi.OutputState
}

func (ApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiKeyOutput)(nil)).Elem()
}

func (o ApiKeyOutput) ToApiKeyOutput() ApiKeyOutput {
	return o
}

func (o ApiKeyOutput) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiKeyOutput{})
}
