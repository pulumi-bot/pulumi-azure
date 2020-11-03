// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Application Insights WebTest.
type WebTest struct {
	pulumi.CustomResourceState

	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringOutput `pulumi:"applicationInsightsId"`
	// An XML configuration specification for a WebTest.
	Configuration pulumi.StringOutput `pulumi:"configuration"`
	// Purpose/user defined descriptive test for this WebTest.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Interval in seconds between test runs for this WebTest. Default is `300`.
	Frequency pulumi.IntPtrOutput `pulumi:"frequency"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations pulumi.StringArrayOutput `pulumi:"geoLocations"`
	// = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The location of the resource group.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Application Insights WebTest. Changing this forces a
	// new resource to be created.
	Name              pulumi.StringOutput `pulumi:"name"`
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Allow for retries should this WebTest fail.
	RetryEnabled       pulumi.BoolPtrOutput `pulumi:"retryEnabled"`
	SyntheticMonitorId pulumi.StringOutput  `pulumi:"syntheticMonitorId"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default is `30`.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewWebTest registers a new resource with the given unique name, arguments, and options.
func NewWebTest(ctx *pulumi.Context,
	name string, args *WebTestArgs, opts ...pulumi.ResourceOption) (*WebTest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ApplicationInsightsId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationInsightsId'")
	}
	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.GeoLocations == nil {
		return nil, errors.New("invalid value for required argument 'GeoLocations'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource WebTest
	err := ctx.RegisterResource("azure:appinsights/webTest:WebTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTest gets an existing WebTest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTestState, opts ...pulumi.ResourceOption) (*WebTest, error) {
	var resource WebTest
	err := ctx.ReadResource("azure:appinsights/webTest:WebTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTest resources.
type webTestState struct {
	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsId *string `pulumi:"applicationInsightsId"`
	// An XML configuration specification for a WebTest.
	Configuration *string `pulumi:"configuration"`
	// Purpose/user defined descriptive test for this WebTest.
	Description *string `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled *bool `pulumi:"enabled"`
	// Interval in seconds between test runs for this WebTest. Default is `300`.
	Frequency *int `pulumi:"frequency"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations []string `pulumi:"geoLocations"`
	// = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
	Kind *string `pulumi:"kind"`
	// The location of the resource group.
	Location *string `pulumi:"location"`
	// Specifies the name of the Application Insights WebTest. Changing this forces a
	// new resource to be created.
	Name              *string `pulumi:"name"`
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Allow for retries should this WebTest fail.
	RetryEnabled       *bool   `pulumi:"retryEnabled"`
	SyntheticMonitorId *string `pulumi:"syntheticMonitorId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default is `30`.
	Timeout *int `pulumi:"timeout"`
}

type WebTestState struct {
	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringPtrInput
	// An XML configuration specification for a WebTest.
	Configuration pulumi.StringPtrInput
	// Purpose/user defined descriptive test for this WebTest.
	Description pulumi.StringPtrInput
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrInput
	// Interval in seconds between test runs for this WebTest. Default is `300`.
	Frequency pulumi.IntPtrInput
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations pulumi.StringArrayInput
	// = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
	Kind pulumi.StringPtrInput
	// The location of the resource group.
	Location pulumi.StringPtrInput
	// Specifies the name of the Application Insights WebTest. Changing this forces a
	// new resource to be created.
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringPtrInput
	// Allow for retries should this WebTest fail.
	RetryEnabled       pulumi.BoolPtrInput
	SyntheticMonitorId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Seconds until this WebTest will timeout and fail. Default is `30`.
	Timeout pulumi.IntPtrInput
}

func (WebTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestState)(nil)).Elem()
}

type webTestArgs struct {
	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsId string `pulumi:"applicationInsightsId"`
	// An XML configuration specification for a WebTest.
	Configuration string `pulumi:"configuration"`
	// Purpose/user defined descriptive test for this WebTest.
	Description *string `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled *bool `pulumi:"enabled"`
	// Interval in seconds between test runs for this WebTest. Default is `300`.
	Frequency *int `pulumi:"frequency"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations []string `pulumi:"geoLocations"`
	// = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
	Kind string `pulumi:"kind"`
	// The location of the resource group.
	Location *string `pulumi:"location"`
	// Specifies the name of the Application Insights WebTest. Changing this forces a
	// new resource to be created.
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	// Allow for retries should this WebTest fail.
	RetryEnabled *bool `pulumi:"retryEnabled"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default is `30`.
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a WebTest resource.
type WebTestArgs struct {
	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringInput
	// An XML configuration specification for a WebTest.
	Configuration pulumi.StringInput
	// Purpose/user defined descriptive test for this WebTest.
	Description pulumi.StringPtrInput
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrInput
	// Interval in seconds between test runs for this WebTest. Default is `300`.
	Frequency pulumi.IntPtrInput
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations pulumi.StringArrayInput
	// = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`.
	Kind pulumi.StringInput
	// The location of the resource group.
	Location pulumi.StringPtrInput
	// Specifies the name of the Application Insights WebTest. Changing this forces a
	// new resource to be created.
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	// Allow for retries should this WebTest fail.
	RetryEnabled pulumi.BoolPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Seconds until this WebTest will timeout and fail. Default is `30`.
	Timeout pulumi.IntPtrInput
}

func (WebTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestArgs)(nil)).Elem()
}
