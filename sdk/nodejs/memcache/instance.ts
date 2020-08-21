// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:memcache/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The full name of the GCE network to connect the instance to.  If not provided,
     * 'default' will be used.
     */
    public readonly authorizedNetwork!: pulumi.Output<string | undefined>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A user-visible name for the instance.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The full version of memcached server running on this instance.
     */
    public /*out*/ readonly memcacheFullVersion!: pulumi.Output<string>;
    /**
     * Additional information about the instance state, if available.
     */
    public /*out*/ readonly memcacheNodes!: pulumi.Output<outputs.memcache.InstanceMemcacheNode[]>;
    /**
     * User-specified parameters for this memcache instance.
     * Structure is documented below.
     */
    public readonly memcacheParameters!: pulumi.Output<outputs.memcache.InstanceMemcacheParameters | undefined>;
    /**
     * The major version of Memcached software. If not provided, latest supported version will be used.
     * Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
     * determined by our system based on the latest supported minor version.
     * Default value is `MEMCACHE_1_5`.
     * Possible values are `MEMCACHE_1_5`.
     */
    public readonly memcacheVersion!: pulumi.Output<string | undefined>;
    /**
     * The resource name of the instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration for memcache nodes.
     * Structure is documented below.
     */
    public readonly nodeConfig!: pulumi.Output<outputs.memcache.InstanceNodeConfig>;
    /**
     * Number of nodes in the memcache instance.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the Memcache region of the instance.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Zones where memcache nodes should be provisioned.  If not
     * provided, all zones will be used.
     */
    public readonly zones!: pulumi.Output<string[]>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["authorizedNetwork"] = state ? state.authorizedNetwork : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["memcacheFullVersion"] = state ? state.memcacheFullVersion : undefined;
            inputs["memcacheNodes"] = state ? state.memcacheNodes : undefined;
            inputs["memcacheParameters"] = state ? state.memcacheParameters : undefined;
            inputs["memcacheVersion"] = state ? state.memcacheVersion : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeConfig"] = state ? state.nodeConfig : undefined;
            inputs["nodeCount"] = state ? state.nodeCount : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.nodeConfig === undefined) {
                throw new Error("Missing required property 'nodeConfig'");
            }
            if (!args || args.nodeCount === undefined) {
                throw new Error("Missing required property 'nodeCount'");
            }
            if (!args || args.region === undefined) {
                throw new Error("Missing required property 'region'");
            }
            inputs["authorizedNetwork"] = args ? args.authorizedNetwork : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["memcacheParameters"] = args ? args.memcacheParameters : undefined;
            inputs["memcacheVersion"] = args ? args.memcacheVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeConfig"] = args ? args.nodeConfig : undefined;
            inputs["nodeCount"] = args ? args.nodeCount : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["memcacheFullVersion"] = undefined /*out*/;
            inputs["memcacheNodes"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The full name of the GCE network to connect the instance to.  If not provided,
     * 'default' will be used.
     */
    readonly authorizedNetwork?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * A user-visible name for the instance.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The full version of memcached server running on this instance.
     */
    readonly memcacheFullVersion?: pulumi.Input<string>;
    /**
     * Additional information about the instance state, if available.
     */
    readonly memcacheNodes?: pulumi.Input<pulumi.Input<inputs.memcache.InstanceMemcacheNode>[]>;
    /**
     * User-specified parameters for this memcache instance.
     * Structure is documented below.
     */
    readonly memcacheParameters?: pulumi.Input<inputs.memcache.InstanceMemcacheParameters>;
    /**
     * The major version of Memcached software. If not provided, latest supported version will be used.
     * Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
     * determined by our system based on the latest supported minor version.
     * Default value is `MEMCACHE_1_5`.
     * Possible values are `MEMCACHE_1_5`.
     */
    readonly memcacheVersion?: pulumi.Input<string>;
    /**
     * The resource name of the instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration for memcache nodes.
     * Structure is documented below.
     */
    readonly nodeConfig?: pulumi.Input<inputs.memcache.InstanceNodeConfig>;
    /**
     * Number of nodes in the memcache instance.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The name of the Memcache region of the instance.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Zones where memcache nodes should be provisioned.  If not
     * provided, all zones will be used.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The full name of the GCE network to connect the instance to.  If not provided,
     * 'default' will be used.
     */
    readonly authorizedNetwork?: pulumi.Input<string>;
    /**
     * A user-visible name for the instance.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * User-specified parameters for this memcache instance.
     * Structure is documented below.
     */
    readonly memcacheParameters?: pulumi.Input<inputs.memcache.InstanceMemcacheParameters>;
    /**
     * The major version of Memcached software. If not provided, latest supported version will be used.
     * Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
     * determined by our system based on the latest supported minor version.
     * Default value is `MEMCACHE_1_5`.
     * Possible values are `MEMCACHE_1_5`.
     */
    readonly memcacheVersion?: pulumi.Input<string>;
    /**
     * The resource name of the instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration for memcache nodes.
     * Structure is documented below.
     */
    readonly nodeConfig: pulumi.Input<inputs.memcache.InstanceNodeConfig>;
    /**
     * Number of nodes in the memcache instance.
     */
    readonly nodeCount: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The name of the Memcache region of the instance.
     */
    readonly region: pulumi.Input<string>;
    /**
     * Zones where memcache nodes should be provisioned.  If not
     * provided, all zones will be used.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}