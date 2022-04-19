// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudFunctionsV2.Inputs
{

    public sealed class FunctionBuildConfigSourceRepoSourceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Regex matching branches to build.
        /// </summary>
        [Input("branchName")]
        public Input<string>? BranchName { get; set; }

        /// <summary>
        /// Regex matching tags to build.
        /// </summary>
        [Input("commitSha")]
        public Input<string>? CommitSha { get; set; }

        /// <summary>
        /// Directory, relative to the source root, in which to run the build.
        /// </summary>
        [Input("dir")]
        public Input<string>? Dir { get; set; }

        /// <summary>
        /// Only trigger a build if the revision regex does
        /// NOT match the revision regex.
        /// </summary>
        [Input("invertRegex")]
        public Input<bool>? InvertRegex { get; set; }

        /// <summary>
        /// ID of the project that owns the Cloud Source Repository. If omitted, the
        /// project ID requesting the build is assumed.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Name of the Cloud Source Repository.
        /// </summary>
        [Input("repoName")]
        public Input<string>? RepoName { get; set; }

        /// <summary>
        /// Regex matching tags to build.
        /// </summary>
        [Input("tagName")]
        public Input<string>? TagName { get; set; }

        public FunctionBuildConfigSourceRepoSourceGetArgs()
        {
        }
    }
}