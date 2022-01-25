// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class AzureClusterControlPlaneRootVolumeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
        /// </summary>
        [Input("sizeGib")]
        public Input<int>? SizeGib { get; set; }

        public AzureClusterControlPlaneRootVolumeGetArgs()
        {
        }
    }
}