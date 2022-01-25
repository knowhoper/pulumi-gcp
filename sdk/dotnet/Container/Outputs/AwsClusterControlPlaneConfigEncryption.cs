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
    public sealed class AwsClusterControlPlaneConfigEncryption
    {
        /// <summary>
        /// Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
        /// </summary>
        public readonly string KmsKeyArn;

        [OutputConstructor]
        private AwsClusterControlPlaneConfigEncryption(string kmsKeyArn)
        {
            KmsKeyArn = kmsKeyArn;
        }
    }
}