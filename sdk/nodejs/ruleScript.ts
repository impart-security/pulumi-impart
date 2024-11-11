// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage a rule script.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as impart from "@impart-security/pulumi-impart";
 *
 * // Create a new rule script
 * const example = new impart.RuleScript("example", {
 *     name: "example",
 *     disabled: false,
 *     description: "Rule description",
 *     sourceFile: `${path.module}/rule.js`,
 *     sourceHash: "<sha256 hash for the source_file content>",
 *     blockingEffect: "block",
 * });
 * // Create a new rule script with content
 * const exampleContent = new impart.RuleScript("exampleContent", {
 *     name: "example",
 *     disabled: false,
 *     description: "Rule description",
 *     content: fs.readFileSync(`${path.module}/rule.js`, "utf8"),
 *     blockingEffect: "block",
 * });
 * ```
 */
export class RuleScript extends pulumi.CustomResource {
    /**
     * Get an existing RuleScript resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleScriptState, opts?: pulumi.CustomResourceOptions): RuleScript {
        return new RuleScript(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'impart:index/ruleScript:RuleScript';

    /**
     * Returns true if the given object is an instance of RuleScript.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleScript {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleScript.__pulumiType;
    }

    /**
     * The rule blocking effect. Allowed values: block, simulate. If not set effect will be block.
     */
    public readonly blockingEffect!: pulumi.Output<string | undefined>;
    /**
     * The rule body content.
     */
    public readonly content!: pulumi.Output<string | undefined>;
    /**
     * The description for this rule script.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Set true to disable the rule script.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * The applied labels.
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * The name for this rule script.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The rule source file.
     */
    public readonly sourceFile!: pulumi.Output<string | undefined>;
    /**
     * The rule source hash.
     */
    public readonly sourceHash!: pulumi.Output<string | undefined>;

    /**
     * Create a RuleScript resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleScriptArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleScriptArgs | RuleScriptState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleScriptState | undefined;
            resourceInputs["blockingEffect"] = state ? state.blockingEffect : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sourceFile"] = state ? state.sourceFile : undefined;
            resourceInputs["sourceHash"] = state ? state.sourceHash : undefined;
        } else {
            const args = argsOrState as RuleScriptArgs | undefined;
            if ((!args || args.disabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'disabled'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["blockingEffect"] = args ? args.blockingEffect : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sourceFile"] = args ? args.sourceFile : undefined;
            resourceInputs["sourceHash"] = args ? args.sourceHash : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RuleScript.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleScript resources.
 */
export interface RuleScriptState {
    /**
     * The rule blocking effect. Allowed values: block, simulate. If not set effect will be block.
     */
    blockingEffect?: pulumi.Input<string>;
    /**
     * The rule body content.
     */
    content?: pulumi.Input<string>;
    /**
     * The description for this rule script.
     */
    description?: pulumi.Input<string>;
    /**
     * Set true to disable the rule script.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The applied labels.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name for this rule script.
     */
    name?: pulumi.Input<string>;
    /**
     * The rule source file.
     */
    sourceFile?: pulumi.Input<string>;
    /**
     * The rule source hash.
     */
    sourceHash?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RuleScript resource.
 */
export interface RuleScriptArgs {
    /**
     * The rule blocking effect. Allowed values: block, simulate. If not set effect will be block.
     */
    blockingEffect?: pulumi.Input<string>;
    /**
     * The rule body content.
     */
    content?: pulumi.Input<string>;
    /**
     * The description for this rule script.
     */
    description?: pulumi.Input<string>;
    /**
     * Set true to disable the rule script.
     */
    disabled: pulumi.Input<boolean>;
    /**
     * The applied labels.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name for this rule script.
     */
    name: pulumi.Input<string>;
    /**
     * The rule source file.
     */
    sourceFile?: pulumi.Input<string>;
    /**
     * The rule source hash.
     */
    sourceHash?: pulumi.Input<string>;
}
