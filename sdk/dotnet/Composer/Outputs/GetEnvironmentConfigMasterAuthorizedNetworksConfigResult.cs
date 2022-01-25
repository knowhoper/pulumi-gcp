// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer.Outputs
{

    [OutputType]
    public sealed class GetEnvironmentConfigMasterAuthorizedNetworksConfigResult
    {
        public readonly ImmutableArray<Outputs.GetEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlockResult> CidrBlocks;
        public readonly bool Enabled;

        [OutputConstructor]
        private GetEnvironmentConfigMasterAuthorizedNetworksConfigResult(
            ImmutableArray<Outputs.GetEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlockResult> cidrBlocks,

            bool enabled)
        {
            CidrBlocks = cidrBlocks;
            Enabled = enabled;
        }
    }
}