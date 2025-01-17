// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage.Outputs
{

    [OutputType]
    public sealed class TransferJobTransferSpecAzureBlobStorageDataSource
    {
        /// <summary>
        /// Credentials used to authenticate API requests to Azure block.
        /// </summary>
        public readonly Outputs.TransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials AzureCredentials;
        /// <summary>
        /// The container to transfer from the Azure Storage account.`
        /// </summary>
        public readonly string Container;
        /// <summary>
        /// Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The name of the Azure Storage account.
        /// </summary>
        public readonly string StorageAccount;

        [OutputConstructor]
        private TransferJobTransferSpecAzureBlobStorageDataSource(
            Outputs.TransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials azureCredentials,

            string container,

            string? path,

            string storageAccount)
        {
            AzureCredentials = azureCredentials;
            Container = container;
            Path = path;
            StorageAccount = storageAccount;
        }
    }
}
