// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudTasks.Inputs
{

    public sealed class QueueStackdriverLoggingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the fraction of operations to write to Stackdriver Logging.
        /// This field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the
        /// default and means that no operations are logged.
        /// </summary>
        [Input("samplingRatio", required: true)]
        public Input<double> SamplingRatio { get; set; } = null!;

        public QueueStackdriverLoggingConfigArgs()
        {
        }
    }
}