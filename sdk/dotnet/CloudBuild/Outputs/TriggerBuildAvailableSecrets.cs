// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Outputs
{

    [OutputType]
    public sealed class TriggerBuildAvailableSecrets
    {
        /// <summary>
        /// Pairs a secret environment variable with a SecretVersion in Secret Manager.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.TriggerBuildAvailableSecretsSecretManager> SecretManagers;

        [OutputConstructor]
        private TriggerBuildAvailableSecrets(ImmutableArray<Outputs.TriggerBuildAvailableSecretsSecretManager> secretManagers)
        {
            SecretManagers = secretManagers;
        }
    }
}