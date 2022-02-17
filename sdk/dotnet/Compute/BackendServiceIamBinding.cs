// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    [GcpResourceType("gcp:compute/backendServiceIamBinding:BackendServiceIamBinding")]
    public partial class BackendServiceIamBinding : Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.BackendServiceIamBindingCondition?> Condition { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a BackendServiceIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackendServiceIamBinding(string name, BackendServiceIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/backendServiceIamBinding:BackendServiceIamBinding", name, args ?? new BackendServiceIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackendServiceIamBinding(string name, Input<string> id, BackendServiceIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/backendServiceIamBinding:BackendServiceIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackendServiceIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackendServiceIamBinding Get(string name, Input<string> id, BackendServiceIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new BackendServiceIamBinding(name, id, state, options);
        }
    }

    public sealed class BackendServiceIamBindingArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.BackendServiceIamBindingConditionArgs>? Condition { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public BackendServiceIamBindingArgs()
        {
        }
    }

    public sealed class BackendServiceIamBindingState : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.BackendServiceIamBindingConditionGetArgs>? Condition { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        public BackendServiceIamBindingState()
        {
        }
    }
}