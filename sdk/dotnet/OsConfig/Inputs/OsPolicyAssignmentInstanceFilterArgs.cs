// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsConfig.Inputs
{

    public sealed class OsPolicyAssignmentInstanceFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target all VMs in the project. If true, no other criteria is permitted.
        /// </summary>
        [Input("all")]
        public Input<bool>? All { get; set; }

        [Input("exclusionLabels")]
        private InputList<Inputs.OsPolicyAssignmentInstanceFilterExclusionLabelArgs>? _exclusionLabels;

        /// <summary>
        /// List of label sets used for VM exclusion. If the list has more than one label set, the VM is excluded if any of the label sets are applicable for the VM.
        /// </summary>
        public InputList<Inputs.OsPolicyAssignmentInstanceFilterExclusionLabelArgs> ExclusionLabels
        {
            get => _exclusionLabels ?? (_exclusionLabels = new InputList<Inputs.OsPolicyAssignmentInstanceFilterExclusionLabelArgs>());
            set => _exclusionLabels = value;
        }

        [Input("inclusionLabels")]
        private InputList<Inputs.OsPolicyAssignmentInstanceFilterInclusionLabelArgs>? _inclusionLabels;

        /// <summary>
        /// List of label sets used for VM inclusion. If the list has more than one `LabelSet`, the VM is included if any of the label sets are applicable for the VM.
        /// </summary>
        public InputList<Inputs.OsPolicyAssignmentInstanceFilterInclusionLabelArgs> InclusionLabels
        {
            get => _inclusionLabels ?? (_inclusionLabels = new InputList<Inputs.OsPolicyAssignmentInstanceFilterInclusionLabelArgs>());
            set => _inclusionLabels = value;
        }

        [Input("inventories")]
        private InputList<Inputs.OsPolicyAssignmentInstanceFilterInventoryArgs>? _inventories;

        /// <summary>
        /// List of inventories to select VMs. A VM is selected if its inventory data matches at least one of the following inventories.
        /// </summary>
        public InputList<Inputs.OsPolicyAssignmentInstanceFilterInventoryArgs> Inventories
        {
            get => _inventories ?? (_inventories = new InputList<Inputs.OsPolicyAssignmentInstanceFilterInventoryArgs>());
            set => _inventories = value;
        }

        public OsPolicyAssignmentInstanceFilterArgs()
        {
        }
    }
}