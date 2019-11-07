// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_target_https_proxy.html.markdown.
    /// </summary>
    public partial class RegionTargetHttpsProxy : Pulumi.CustomResource
    {
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("proxyId")]
        public Output<int> ProxyId { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        [Output("sslCertificates")]
        public Output<ImmutableArray<string>> SslCertificates { get; private set; } = null!;

        [Output("urlMap")]
        public Output<string> UrlMap { get; private set; } = null!;


        /// <summary>
        /// Create a RegionTargetHttpsProxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionTargetHttpsProxy(string name, RegionTargetHttpsProxyArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy", name, args, MakeResourceOptions(options, ""))
        {
        }

        private RegionTargetHttpsProxy(string name, Input<string> id, RegionTargetHttpsProxyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RegionTargetHttpsProxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionTargetHttpsProxy Get(string name, Input<string> id, RegionTargetHttpsProxyState? state = null, CustomResourceOptions? options = null)
        {
            return new RegionTargetHttpsProxy(name, id, state, options);
        }
    }

    public sealed class RegionTargetHttpsProxyArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sslCertificates", required: true)]
        private InputList<string>? _sslCertificates;
        public InputList<string> SslCertificates
        {
            get => _sslCertificates ?? (_sslCertificates = new InputList<string>());
            set => _sslCertificates = value;
        }

        [Input("urlMap", required: true)]
        public Input<string> UrlMap { get; set; } = null!;

        public RegionTargetHttpsProxyArgs()
        {
        }
    }

    public sealed class RegionTargetHttpsProxyState : Pulumi.ResourceArgs
    {
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("proxyId")]
        public Input<int>? ProxyId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("sslCertificates")]
        private InputList<string>? _sslCertificates;
        public InputList<string> SslCertificates
        {
            get => _sslCertificates ?? (_sslCertificates = new InputList<string>());
            set => _sslCertificates = value;
        }

        [Input("urlMap")]
        public Input<string>? UrlMap { get; set; }

        public RegionTargetHttpsProxyState()
        {
        }
    }
}