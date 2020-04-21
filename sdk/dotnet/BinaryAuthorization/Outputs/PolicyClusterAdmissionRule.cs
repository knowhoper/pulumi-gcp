// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BinaryAuthorization.Outputs
{

    [OutputType]
    public sealed class PolicyClusterAdmissionRule
    {
        /// <summary>
        /// The identifier for this object. Format specified above.
        /// </summary>
        public readonly string Cluster;
        public readonly string EnforcementMode;
        public readonly string EvaluationMode;
        public readonly ImmutableArray<string> RequireAttestationsBies;

        [OutputConstructor]
        private PolicyClusterAdmissionRule(
            string cluster,

            string enforcementMode,

            string evaluationMode,

            ImmutableArray<string> requireAttestationsBies)
        {
            Cluster = cluster;
            EnforcementMode = enforcementMode;
            EvaluationMode = evaluationMode;
            RequireAttestationsBies = requireAttestationsBies;
        }
    }
}