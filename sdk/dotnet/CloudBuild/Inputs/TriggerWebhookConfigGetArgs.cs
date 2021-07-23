// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Inputs
{

    public sealed class TriggerWebhookConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Secrets to decrypt using Cloud Key Management Service.
        /// Structure is documented below.
        /// </summary>
        [Input("secret", required: true)]
        public Input<string> Secret { get; set; } = null!;

        /// <summary>
        /// -
        /// Potential issues with the underlying Pub/Sub subscription configuration.
        /// Only populated on get requests.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public TriggerWebhookConfigGetArgs()
        {
        }
    }
}