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
    public sealed class AwsClusterAuthorizationAdminUser
    {
        /// <summary>
        /// Required. The name of the user, e.g. `my-gcp-id@gmail.com`.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private AwsClusterAuthorizationAdminUser(string username)
        {
            Username = username;
        }
    }
}