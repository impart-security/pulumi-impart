// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage a log binding.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as impart from "@impart-security/pulumi-impart";
 *
 * // Create a new log binding
 * const example = new impart.LogBinding("example", {
 *     name: "log_binding_example",
 *     patternType: "grok",
 *     pattern: "<pattern>\n",
 *     logstreamId: "logstream_id",
 *     specId: resource.impart_spec.example.id,
 * });
 * ```
 */
export class LogBinding extends pulumi.CustomResource {
    /**
     * Get an existing LogBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogBindingState, opts?: pulumi.CustomResourceOptions): LogBinding {
        return new LogBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'impart:index/logBinding:LogBinding';

    /**
     * Returns true if the given object is an instance of LogBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogBinding.__pulumiType;
    }

    /**
     * The logstream id for this log binding.
     */
    public readonly logstreamId!: pulumi.Output<string | undefined>;
    /**
     * The name for this log binding.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The grok/json pattern for this log binding.
     */
    public readonly pattern!: pulumi.Output<string>;
    /**
     * The pattern type for this log binding. Accepted values: grok, json
     */
    public readonly patternType!: pulumi.Output<string>;
    /**
     * The specification id.
     */
    public readonly specId!: pulumi.Output<string>;

    /**
     * Create a LogBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogBindingArgs | LogBindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogBindingState | undefined;
            resourceInputs["logstreamId"] = state ? state.logstreamId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pattern"] = state ? state.pattern : undefined;
            resourceInputs["patternType"] = state ? state.patternType : undefined;
            resourceInputs["specId"] = state ? state.specId : undefined;
        } else {
            const args = argsOrState as LogBindingArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.pattern === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pattern'");
            }
            if ((!args || args.patternType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'patternType'");
            }
            if ((!args || args.specId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'specId'");
            }
            resourceInputs["logstreamId"] = args ? args.logstreamId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pattern"] = args ? args.pattern : undefined;
            resourceInputs["patternType"] = args ? args.patternType : undefined;
            resourceInputs["specId"] = args ? args.specId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogBinding resources.
 */
export interface LogBindingState {
    /**
     * The logstream id for this log binding.
     */
    logstreamId?: pulumi.Input<string>;
    /**
     * The name for this log binding.
     */
    name?: pulumi.Input<string>;
    /**
     * The grok/json pattern for this log binding.
     */
    pattern?: pulumi.Input<string>;
    /**
     * The pattern type for this log binding. Accepted values: grok, json
     */
    patternType?: pulumi.Input<string>;
    /**
     * The specification id.
     */
    specId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogBinding resource.
 */
export interface LogBindingArgs {
    /**
     * The logstream id for this log binding.
     */
    logstreamId?: pulumi.Input<string>;
    /**
     * The name for this log binding.
     */
    name: pulumi.Input<string>;
    /**
     * The grok/json pattern for this log binding.
     */
    pattern: pulumi.Input<string>;
    /**
     * The pattern type for this log binding. Accepted values: grok, json
     */
    patternType: pulumi.Input<string>;
    /**
     * The specification id.
     */
    specId: pulumi.Input<string>;
}
