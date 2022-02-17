// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    [GcpResourceType("gcp:compute/backendServiceIamMember:BackendServiceIamMember")]
    public partial class BackendServiceIamMember : Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.BackendServiceIamMemberCondition?> Condition { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a BackendServiceIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackendServiceIamMember(string name, BackendServiceIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/backendServiceIamMember:BackendServiceIamMember", name, args ?? new BackendServiceIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackendServiceIamMember(string name, Input<string> id, BackendServiceIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/backendServiceIamMember:BackendServiceIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackendServiceIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackendServiceIamMember Get(string name, Input<string> id, BackendServiceIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new BackendServiceIamMember(name, id, state, options);
        }
    }

    public sealed class BackendServiceIamMemberArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.BackendServiceIamMemberConditionArgs>? Condition { get; set; }

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public BackendServiceIamMemberArgs()
        {
        }
    }

    public sealed class BackendServiceIamMemberState : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.BackendServiceIamMemberConditionGetArgs>? Condition { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        public BackendServiceIamMemberState()
        {
        }
    }
}