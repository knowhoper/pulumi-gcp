// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BinaryAuthorization.Outputs
{

    [OutputType]
    public sealed class AttestorAttestationAuthorityNotePublicKey
    {
        public readonly string? AsciiArmoredPgpPublicKey;
        public readonly string? Comment;
        /// <summary>
        /// an identifier for the resource with format `projects/{{project}}/attestors/{{name}}`
        /// </summary>
        public readonly string? Id;
        public readonly Outputs.AttestorAttestationAuthorityNotePublicKeyPkixPublicKey? PkixPublicKey;

        [OutputConstructor]
        private AttestorAttestationAuthorityNotePublicKey(
            string? asciiArmoredPgpPublicKey,

            string? comment,

            string? id,

            Outputs.AttestorAttestationAuthorityNotePublicKeyPkixPublicKey? pkixPublicKey)
        {
            AsciiArmoredPgpPublicKey = asciiArmoredPgpPublicKey;
            Comment = comment;
            Id = id;
            PkixPublicKey = pkixPublicKey;
        }
    }
}