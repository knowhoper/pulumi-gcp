// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Outputs
{

    [OutputType]
    public sealed class ClusterAuthenticatorGroupsConfig
    {
        /// <summary>
        /// The name of the RBAC security group for use with Google security groups in Kubernetes RBAC. Group name must be in format `gke-security-groups@yourdomain.com`.
        /// </summary>
        public readonly string SecurityGroup;

        [OutputConstructor]
        private ClusterAuthenticatorGroupsConfig(string securityGroup)
        {
            SecurityGroup = securityGroup;
        }
    }
}