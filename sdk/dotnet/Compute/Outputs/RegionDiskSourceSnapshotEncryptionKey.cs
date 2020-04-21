// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class RegionDiskSourceSnapshotEncryptionKey
    {
        public readonly string? KmsKeyName;
        public readonly string? RawKey;
        public readonly string? Sha256;

        [OutputConstructor]
        private RegionDiskSourceSnapshotEncryptionKey(
            string? kmsKeyName,

            string? rawKey,

            string? sha256)
        {
            KmsKeyName = kmsKeyName;
            RawKey = rawKey;
            Sha256 = sha256;
        }
    }
}