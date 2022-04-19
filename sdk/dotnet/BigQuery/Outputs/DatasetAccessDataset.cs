// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery.Outputs
{

    [OutputType]
    public sealed class DatasetAccessDataset
    {
        /// <summary>
        /// The dataset this entry applies to
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.DatasetAccessDatasetDataset Dataset;
        /// <summary>
        /// Which resources in the dataset this entry applies to. Currently, only views are supported,
        /// but additional target types may be added in the future. Possible values: VIEWS
        /// </summary>
        public readonly ImmutableArray<string> TargetTypes;

        [OutputConstructor]
        private DatasetAccessDataset(
            Outputs.DatasetAccessDatasetDataset dataset,

            ImmutableArray<string> targetTypes)
        {
            Dataset = dataset;
            TargetTypes = targetTypes;
        }
    }
}