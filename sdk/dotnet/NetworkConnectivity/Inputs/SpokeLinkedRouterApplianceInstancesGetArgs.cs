// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkConnectivity.Inputs
{

    public sealed class SpokeLinkedRouterApplianceInstancesGetArgs : Pulumi.ResourceArgs
    {
        [Input("instances", required: true)]
        private InputList<Inputs.SpokeLinkedRouterApplianceInstancesInstanceGetArgs>? _instances;

        /// <summary>
        /// The list of router appliance instances
        /// </summary>
        public InputList<Inputs.SpokeLinkedRouterApplianceInstancesInstanceGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.SpokeLinkedRouterApplianceInstancesInstanceGetArgs>());
            set => _instances = value;
        }

        /// <summary>
        /// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
        /// </summary>
        [Input("siteToSiteDataTransfer", required: true)]
        public Input<bool> SiteToSiteDataTransfer { get; set; } = null!;

        public SpokeLinkedRouterApplianceInstancesGetArgs()
        {
        }
    }
}