// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class WorkflowTemplatePlacementManagedClusterConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Autoscaling config for the policy associated with the cluster. Cluster does not autoscale if this field is unset.
        /// </summary>
        [Input("autoscalingConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigArgs>? AutoscalingConfig { get; set; }

        /// <summary>
        /// Optional. Encryption settings for the cluster.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfigArgs>? EncryptionConfig { get; set; }

        /// <summary>
        /// Optional. Port/endpoint configuration for this cluster
        /// </summary>
        [Input("endpointConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigEndpointConfigArgs>? EndpointConfig { get; set; }

        /// <summary>
        /// Optional. The shared Compute Engine config settings for all instances in a cluster.
        /// </summary>
        [Input("gceClusterConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigArgs>? GceClusterConfig { get; set; }

        /// <summary>
        /// Optional. The Kubernetes Engine config for Dataproc clusters deployed to Kubernetes. Setting this is considered mutually exclusive with Compute Engine-based options such as `gce_cluster_config`, `master_config`, `worker_config`, `secondary_worker_config`, and `autoscaling_config`.
        /// </summary>
        [Input("gkeClusterConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigArgs>? GkeClusterConfig { get; set; }

        [Input("initializationActions")]
        private InputList<Inputs.WorkflowTemplatePlacementManagedClusterConfigInitializationActionArgs>? _initializationActions;

        /// <summary>
        /// Optional. Commands to execute on each node after config is completed. By default, executables are run on master and all worker nodes. You can test a node's `role` metadata to run an executable on a master or worker node, as shown below using `curl` (you can also use `wget`): ROLE=$(curl -H Metadata-Flavor:Google http://metadata/computeMetadata/v1/instance/attributes/dataproc-role) if ; then ... master specific actions ... else ... worker specific actions ... fi
        /// </summary>
        public InputList<Inputs.WorkflowTemplatePlacementManagedClusterConfigInitializationActionArgs> InitializationActions
        {
            get => _initializationActions ?? (_initializationActions = new InputList<Inputs.WorkflowTemplatePlacementManagedClusterConfigInitializationActionArgs>());
            set => _initializationActions = value;
        }

        /// <summary>
        /// Optional. Lifecycle setting for the cluster.
        /// </summary>
        [Input("lifecycleConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfigArgs>? LifecycleConfig { get; set; }

        /// <summary>
        /// Optional. The Compute Engine config settings for additional worker instances in a cluster.
        /// </summary>
        [Input("masterConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigMasterConfigArgs>? MasterConfig { get; set; }

        /// <summary>
        /// Optional. Metastore configuration.
        /// </summary>
        [Input("metastoreConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigMetastoreConfigArgs>? MetastoreConfig { get; set; }

        /// <summary>
        /// Optional. The Compute Engine config settings for additional worker instances in a cluster.
        /// </summary>
        [Input("secondaryWorkerConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigArgs>? SecondaryWorkerConfig { get; set; }

        /// <summary>
        /// Optional. Security settings for the cluster.
        /// </summary>
        [Input("securityConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigArgs>? SecurityConfig { get; set; }

        /// <summary>
        /// Optional. The config settings for software inside the cluster.
        /// </summary>
        [Input("softwareConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigArgs>? SoftwareConfig { get; set; }

        /// <summary>
        /// Optional. A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see (https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)).
        /// </summary>
        [Input("stagingBucket")]
        public Input<string>? StagingBucket { get; set; }

        /// <summary>
        /// Optional. A Cloud Storage bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. If you do not specify a temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's temp bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket. The default bucket has a TTL of 90 days, but you can use any TTL (or none) if you specify a bucket.
        /// </summary>
        [Input("tempBucket")]
        public Input<string>? TempBucket { get; set; }

        /// <summary>
        /// Optional. The Compute Engine config settings for additional worker instances in a cluster.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigArgs>? WorkerConfig { get; set; }

        public WorkflowTemplatePlacementManagedClusterConfigArgs()
        {
        }
    }
}