// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get the email address of a project's unique Google Cloud Storage service account.
        /// 
        /// Each Google Cloud project has a unique service account for use with Google Cloud Storage. Only this
        /// special service account can be used to set up `gcp.storage.Notification` resources.
        /// 
        /// For more information see
        /// [the API reference](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/storage_project_service_account.html.markdown.
        /// </summary>
        public static Task<GetProjectServiceAccountResult> GetProjectServiceAccount(GetProjectServiceAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectServiceAccountResult>("gcp:storage/getProjectServiceAccount:getProjectServiceAccount", args, options.WithVersion());
    }

    public sealed class GetProjectServiceAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The project the unique service account was created for. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The project the lookup originates from. This field is used if you are making the request
        /// from a different account than the one you are finding the service account for.
        /// </summary>
        [Input("userProject")]
        public Input<string>? UserProject { get; set; }

        public GetProjectServiceAccountArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetProjectServiceAccountResult
    {
        /// <summary>
        /// The email address of the service account. This value is often used to refer to the service account
        /// in order to grant IAM permissions.
        /// </summary>
        public readonly string EmailAddress;
        public readonly string Project;
        public readonly string? UserProject;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetProjectServiceAccountResult(
            string emailAddress,
            string project,
            string? userProject,
            string id)
        {
            EmailAddress = emailAddress;
            Project = project;
            UserProject = userProject;
            Id = id;
        }
    }
}