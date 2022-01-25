// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class AwsNodePoolConfigSshConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The name of the EC2 key pair used to login into cluster machines.
        /// </summary>
        [Input("ec2KeyPair", required: true)]
        public Input<string> Ec2KeyPair { get; set; } = null!;

        public AwsNodePoolConfigSshConfigGetArgs()
        {
        }
    }
}