// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeHub.Inputs
{

    public sealed class MembershipAuthorityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://` and // be a valid
        /// with length &lt;2000 characters.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        public MembershipAuthorityArgs()
        {
        }
    }
}