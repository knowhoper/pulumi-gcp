// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Gets an existing bucket in Google Cloud Storage service (GCS).
 * See [the official documentation](https://cloud.google.com/storage/docs/key-terms#buckets)
 * and
 * [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my_bucket = pulumi.output(gcp.storage.getBucket({
 *     name: "my-bucket",
 * }));
 * ```
 */
export function getBucket(args: GetBucketArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:storage/getBucket:getBucket", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getBucket.
 */
export interface GetBucketArgs {
    /**
     * The name of the bucket.
     */
    name: string;
}

/**
 * A collection of values returned by getBucket.
 */
export interface GetBucketResult {
    readonly cors: outputs.storage.GetBucketCor[];
    readonly defaultEventBasedHold: boolean;
    readonly encryptions: outputs.storage.GetBucketEncryption[];
    readonly forceDestroy: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: {[key: string]: string};
    readonly lifecycleRules: outputs.storage.GetBucketLifecycleRule[];
    readonly location: string;
    readonly loggings: outputs.storage.GetBucketLogging[];
    readonly name: string;
    readonly project: string;
    readonly requesterPays: boolean;
    readonly retentionPolicies: outputs.storage.GetBucketRetentionPolicy[];
    readonly selfLink: string;
    readonly storageClass: string;
    readonly uniformBucketLevelAccess: boolean;
    readonly url: string;
    readonly versionings: outputs.storage.GetBucketVersioning[];
    readonly websites: outputs.storage.GetBucketWebsite[];
}

export function getBucketOutput(args: GetBucketOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBucketResult> {
    return pulumi.output(args).apply(a => getBucket(a, opts))
}

/**
 * A collection of arguments for invoking getBucket.
 */
export interface GetBucketOutputArgs {
    /**
     * The name of the bucket.
     */
    name: pulumi.Input<string>;
}