// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql.Outputs
{

    [OutputType]
    public sealed class GetDatabaseInstanceRestoreBackupContextResult
    {
        public readonly int BackupRunId;
        public readonly string InstanceId;
        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// </summary>
        public readonly string Project;

        [OutputConstructor]
        private GetDatabaseInstanceRestoreBackupContextResult(
            int backupRunId,

            string instanceId,

            string project)
        {
            BackupRunId = backupRunId;
            InstanceId = instanceId;
            Project = project;
        }
    }
}