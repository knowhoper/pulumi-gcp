// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority.Inputs
{

    public sealed class CertificateConfigX509ConfigAdditionalExtensionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether or not this extension is critical (i.e., if the client does not know how to
        /// handle this extension, the client should consider this to be an error).
        /// </summary>
        [Input("critical", required: true)]
        public Input<bool> Critical { get; set; } = null!;

        /// <summary>
        /// Describes values that are relevant in a CA certificate.
        /// Structure is documented below.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<Inputs.CertificateConfigX509ConfigAdditionalExtensionObjectIdArgs> ObjectId { get; set; } = null!;

        /// <summary>
        /// The value of this X.509 extension. A base64-encoded string.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public CertificateConfigX509ConfigAdditionalExtensionArgs()
        {
        }
    }
}