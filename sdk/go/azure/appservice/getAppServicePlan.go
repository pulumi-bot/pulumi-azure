// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing App Service Plan (formerly known as a `Server Farm`).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := appservice.GetAppServicePlan(ctx, "azure:appservice:getAppServicePlan", &appservice.GetAppServicePlanArgs{
// 			Name:              "search-app-service-plan",
// 			ResourceGroupName: "search-service",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("appServicePlanId", example.Id)
// 		return nil
// 	})
// }
// ```
func GetAppServicePlan(ctx *pulumi.Context, args *GetAppServicePlanArgs, opts ...pulumi.InvokeOption) (*GetAppServicePlanResult, error) {
	var rv GetAppServicePlanResult
	err := ctx.Invoke("azure:appservice/getAppServicePlan:getAppServicePlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppServicePlan.
type GetAppServicePlanArgs struct {
	// The name of the App Service Plan.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the App Service Plan exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getAppServicePlan.
type GetAppServicePlanResult struct {
	// The ID of the App Service Environment where the App Service Plan is located.
	AppServiceEnvironmentId string `pulumi:"appServiceEnvironmentId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A flag that indicates if it's a xenon plan (support for Windows Container)
	IsXenon bool `pulumi:"isXenon"`
	// The Operating System type of the App Service Plan
	Kind string `pulumi:"kind"`
	// The Azure location where the App Service Plan exists
	Location string `pulumi:"location"`
	// The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
	MaximumElasticWorkerCount int `pulumi:"maximumElasticWorkerCount"`
	// The maximum number of workers supported with the App Service Plan's sku.
	MaximumNumberOfWorkers int    `pulumi:"maximumNumberOfWorkers"`
	Name                   string `pulumi:"name"`
	// Can Apps assigned to this App Service Plan be scaled independently?
	PerSiteScaling bool `pulumi:"perSiteScaling"`
	// Is this App Service Plan `Reserved`?
	Reserved          bool   `pulumi:"reserved"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `sku` block as documented below.
	Sku GetAppServicePlanSku `pulumi:"sku"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}
