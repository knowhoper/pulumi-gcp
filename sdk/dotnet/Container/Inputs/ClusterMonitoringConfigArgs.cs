// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class ClusterMonitoringConfigArgs : Pulumi.ResourceArgs
    {
        [Input("enableComponents", required: true)]
        private InputList<string>? _enableComponents;

        /// <summary>
        /// The GKE components exposing logs. Only `SYSTEM_COMPONENTS`
        /// is supported.
        /// </summary>
        public InputList<string> EnableComponents
        {
            get => _enableComponents ?? (_enableComponents = new InputList<string>());
            set => _enableComponents = value;
        }

        public ClusterMonitoringConfigArgs()
        {
        }
    }
}