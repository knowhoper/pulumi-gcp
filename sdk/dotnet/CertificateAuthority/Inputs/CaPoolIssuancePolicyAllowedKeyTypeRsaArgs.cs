// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority.Inputs
{

    public sealed class CaPoolIssuancePolicyAllowedKeyTypeRsaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum allowed RSA modulus size, in bits. If this is not set, or if set to zero, the
        /// service will not enforce an explicit upper bound on RSA modulus sizes.
        /// </summary>
        [Input("maxModulusSize")]
        public Input<string>? MaxModulusSize { get; set; }

        /// <summary>
        /// The minimum allowed RSA modulus size, in bits. If this is not set, or if set to zero, the
        /// service-level min RSA modulus size will continue to apply.
        /// </summary>
        [Input("minModulusSize")]
        public Input<string>? MinModulusSize { get; set; }

        public CaPoolIssuancePolicyAllowedKeyTypeRsaArgs()
        {
        }
    }
}