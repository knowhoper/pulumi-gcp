// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./release";
export * from "./ruleset";

// Import resources to register:
import { Release } from "./release";
import { Ruleset } from "./ruleset";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:firebaserules/release:Release":
                return new Release(name, <any>undefined, { urn })
            case "gcp:firebaserules/ruleset:Ruleset":
                return new Ruleset(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "firebaserules/release", _module)
pulumi.runtime.registerResourceModule("gcp", "firebaserules/ruleset", _module)