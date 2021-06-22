// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionInstanceGroupManagerStatusStatefulPerInstanceConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A bit indicating if all of the group's per-instance configs (listed in the output of a listPerInstanceConfigs API call) have status `EFFECTIVE` or there are no per-instance-configs.
        /// </summary>
        [Input("allEffective")]
        public Input<bool>? AllEffective { get; set; }

        public RegionInstanceGroupManagerStatusStatefulPerInstanceConfigGetArgs()
        {
        }
    }
}