// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Allows associating hierarchical firewall policies with the target where they are applied. This allows creating policies and rules in a different location than they are applied.
    /// 
    /// For more information on applying hierarchical firewall policies see the [official documentation](https://cloud.google.com/vpc/docs/firewall-policies#managing_hierarchical_firewall_policy_resources)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var defaultFirewallPolicy = new Gcp.Compute.FirewallPolicy("defaultFirewallPolicy", new Gcp.Compute.FirewallPolicyArgs
    ///         {
    ///             Parent = "organizations/12345",
    ///             ShortName = "my-policy",
    ///             Description = "Example Resource",
    ///         });
    ///         var defaultFirewallPolicyAssociation = new Gcp.Compute.FirewallPolicyAssociation("defaultFirewallPolicyAssociation", new Gcp.Compute.FirewallPolicyAssociationArgs
    ///         {
    ///             FirewallPolicy = defaultFirewallPolicy.Id,
    ///             AttachmentTarget = google_folder.Folder.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// FirewallPolicyAssociation can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default {{firewall_policy}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation")]
    public partial class FirewallPolicyAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// The target that the firewall policy is attached to.
        /// </summary>
        [Output("attachmentTarget")]
        public Output<string> AttachmentTarget { get; private set; } = null!;

        /// <summary>
        /// The firewall policy ID of the association.
        /// </summary>
        [Output("firewallPolicy")]
        public Output<string> FirewallPolicy { get; private set; } = null!;

        /// <summary>
        /// The name for an association.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The short name of the firewall policy of the association.
        /// </summary>
        [Output("shortName")]
        public Output<string> ShortName { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallPolicyAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallPolicyAssociation(string name, FirewallPolicyAssociationArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation", name, args ?? new FirewallPolicyAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallPolicyAssociation(string name, Input<string> id, FirewallPolicyAssociationState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallPolicyAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallPolicyAssociation Get(string name, Input<string> id, FirewallPolicyAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallPolicyAssociation(name, id, state, options);
        }
    }

    public sealed class FirewallPolicyAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target that the firewall policy is attached to.
        /// </summary>
        [Input("attachmentTarget", required: true)]
        public Input<string> AttachmentTarget { get; set; } = null!;

        /// <summary>
        /// The firewall policy ID of the association.
        /// </summary>
        [Input("firewallPolicy", required: true)]
        public Input<string> FirewallPolicy { get; set; } = null!;

        /// <summary>
        /// The name for an association.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallPolicyAssociationArgs()
        {
        }
    }

    public sealed class FirewallPolicyAssociationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target that the firewall policy is attached to.
        /// </summary>
        [Input("attachmentTarget")]
        public Input<string>? AttachmentTarget { get; set; }

        /// <summary>
        /// The firewall policy ID of the association.
        /// </summary>
        [Input("firewallPolicy")]
        public Input<string>? FirewallPolicy { get; set; }

        /// <summary>
        /// The name for an association.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The short name of the firewall policy of the association.
        /// </summary>
        [Input("shortName")]
        public Input<string>? ShortName { get; set; }

        public FirewallPolicyAssociationState()
        {
        }
    }
}