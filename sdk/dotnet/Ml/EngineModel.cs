// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Ml
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/ml_engine_model.html.markdown.
    /// </summary>
    public partial class EngineModel : Pulumi.CustomResource
    {
        [Output("defaultVersion")]
        public Output<Outputs.EngineModelDefaultVersion?> DefaultVersion { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("onlinePredictionConsoleLogging")]
        public Output<bool?> OnlinePredictionConsoleLogging { get; private set; } = null!;

        [Output("onlinePredictionLogging")]
        public Output<bool?> OnlinePredictionLogging { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("regions")]
        public Output<string?> Regions { get; private set; } = null!;


        /// <summary>
        /// Create a EngineModel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EngineModel(string name, EngineModelArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:ml/engineModel:EngineModel", name, args, MakeResourceOptions(options, ""))
        {
        }

        private EngineModel(string name, Input<string> id, EngineModelState? state = null, CustomResourceOptions? options = null)
            : base("gcp:ml/engineModel:EngineModel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EngineModel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EngineModel Get(string name, Input<string> id, EngineModelState? state = null, CustomResourceOptions? options = null)
        {
            return new EngineModel(name, id, state, options);
        }
    }

    public sealed class EngineModelArgs : Pulumi.ResourceArgs
    {
        [Input("defaultVersion")]
        public Input<Inputs.EngineModelDefaultVersionArgs>? DefaultVersion { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("onlinePredictionConsoleLogging")]
        public Input<bool>? OnlinePredictionConsoleLogging { get; set; }

        [Input("onlinePredictionLogging")]
        public Input<bool>? OnlinePredictionLogging { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("regions")]
        public Input<string>? Regions { get; set; }

        public EngineModelArgs()
        {
        }
    }

    public sealed class EngineModelState : Pulumi.ResourceArgs
    {
        [Input("defaultVersion")]
        public Input<Inputs.EngineModelDefaultVersionGetArgs>? DefaultVersion { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("onlinePredictionConsoleLogging")]
        public Input<bool>? OnlinePredictionConsoleLogging { get; set; }

        [Input("onlinePredictionLogging")]
        public Input<bool>? OnlinePredictionLogging { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("regions")]
        public Input<string>? Regions { get; set; }

        public EngineModelState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class EngineModelDefaultVersionArgs : Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public EngineModelDefaultVersionArgs()
        {
        }
    }

    public sealed class EngineModelDefaultVersionGetArgs : Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public EngineModelDefaultVersionGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class EngineModelDefaultVersion
    {
        public readonly string? Name;

        [OutputConstructor]
        private EngineModelDefaultVersion(string? name)
        {
            Name = name;
        }
    }
    }
}