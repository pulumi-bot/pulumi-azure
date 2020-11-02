// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    /// <summary>
    /// Manages an API Management API Assignment to a Product.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleService = Output.Create(Azure.ApiManagement.GetService.InvokeAsync(new Azure.ApiManagement.GetServiceArgs
    ///         {
    ///             Name = "example-api",
    ///             ResourceGroupName = "example-resources",
    ///         }));
    ///         var exampleApi = Output.Tuple(exampleService, exampleService).Apply(values =&gt;
    ///         {
    ///             var exampleService = values.Item1;
    ///             var exampleService1 = values.Item2;
    ///             return Output.Create(Azure.ApiManagement.GetApi.InvokeAsync(new Azure.ApiManagement.GetApiArgs
    ///             {
    ///                 Name = "search-api",
    ///                 ApiManagementName = exampleService.Name,
    ///                 ResourceGroupName = exampleService1.ResourceGroupName,
    ///                 Revision = "2",
    ///             }));
    ///         });
    ///         var exampleProduct = Output.Tuple(exampleService, exampleService).Apply(values =&gt;
    ///         {
    ///             var exampleService = values.Item1;
    ///             var exampleService1 = values.Item2;
    ///             return Output.Create(Azure.ApiManagement.GetProduct.InvokeAsync(new Azure.ApiManagement.GetProductArgs
    ///             {
    ///                 ProductId = "my-product",
    ///                 ApiManagementName = exampleService.Name,
    ///                 ResourceGroupName = exampleService1.ResourceGroupName,
    ///             }));
    ///         });
    ///         var exampleProductApi = new Azure.ApiManagement.ProductApi("exampleProductApi", new Azure.ApiManagement.ProductApiArgs
    ///         {
    ///             ApiName = exampleApi.Apply(exampleApi =&gt; exampleApi.Name),
    ///             ProductId = exampleProduct.Apply(exampleProduct =&gt; exampleProduct.ProductId),
    ///             ApiManagementName = exampleService.Apply(exampleService =&gt; exampleService.Name),
    ///             ResourceGroupName = exampleService.Apply(exampleService =&gt; exampleService.ResourceGroupName),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// API Management Product API's can be imported using the `resource id`, e.g.
    /// </summary>
    public partial class ProductApi : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiManagementName")]
        public Output<string> ApiManagementName { get; private set; } = null!;

        /// <summary>
        /// The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiName")]
        public Output<string> ApiName { get; private set; } = null!;

        /// <summary>
        /// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("productId")]
        public Output<string> ProductId { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a ProductApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProductApi(string name, ProductApiArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/productApi:ProductApi", name, args ?? new ProductApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProductApi(string name, Input<string> id, ProductApiState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/productApi:ProductApi", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProductApi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProductApi Get(string name, Input<string> id, ProductApiState? state = null, CustomResourceOptions? options = null)
        {
            return new ProductApi(name, id, state, options);
        }
    }

    public sealed class ProductApiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiName", required: true)]
        public Input<string> ApiName { get; set; } = null!;

        /// <summary>
        /// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ProductApiArgs()
        {
        }
    }

    public sealed class ProductApiState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName")]
        public Input<string>? ApiManagementName { get; set; }

        /// <summary>
        /// The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiName")]
        public Input<string>? ApiName { get; set; }

        /// <summary>
        /// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public ProductApiState()
        {
        }
    }
}
