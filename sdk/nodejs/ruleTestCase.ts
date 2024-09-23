// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export class RuleTestCase extends pulumi.CustomResource {
    /**
     * Get an existing RuleTestCase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleTestCaseState, opts?: pulumi.CustomResourceOptions): RuleTestCase {
        return new RuleTestCase(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'impart:index/ruleTestCase:RuleTestCase';

    /**
     * Returns true if the given object is an instance of RuleTestCase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleTestCase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleTestCase.__pulumiType;
    }

    /**
     * The assertions of the test case.
     */
    public readonly assertions!: pulumi.Output<outputs.RuleTestCaseAssertion[] | undefined>;
    /**
     * The description of the test case.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The messages of the test case.
     */
    public readonly messages!: pulumi.Output<outputs.RuleTestCaseMessage[]>;
    /**
     * The name of the test case.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Sets if test case required to pass on update.
     */
    public readonly required!: pulumi.Output<boolean | undefined>;

    /**
     * Create a RuleTestCase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleTestCaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleTestCaseArgs | RuleTestCaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleTestCaseState | undefined;
            resourceInputs["assertions"] = state ? state.assertions : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["messages"] = state ? state.messages : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["required"] = state ? state.required : undefined;
        } else {
            const args = argsOrState as RuleTestCaseArgs | undefined;
            if ((!args || args.messages === undefined) && !opts.urn) {
                throw new Error("Missing required property 'messages'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["assertions"] = args ? args.assertions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["messages"] = args ? args.messages : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["required"] = args ? args.required : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RuleTestCase.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleTestCase resources.
 */
export interface RuleTestCaseState {
    /**
     * The assertions of the test case.
     */
    assertions?: pulumi.Input<pulumi.Input<inputs.RuleTestCaseAssertion>[]>;
    /**
     * The description of the test case.
     */
    description?: pulumi.Input<string>;
    /**
     * The messages of the test case.
     */
    messages?: pulumi.Input<pulumi.Input<inputs.RuleTestCaseMessage>[]>;
    /**
     * The name of the test case.
     */
    name?: pulumi.Input<string>;
    /**
     * Sets if test case required to pass on update.
     */
    required?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RuleTestCase resource.
 */
export interface RuleTestCaseArgs {
    /**
     * The assertions of the test case.
     */
    assertions?: pulumi.Input<pulumi.Input<inputs.RuleTestCaseAssertion>[]>;
    /**
     * The description of the test case.
     */
    description?: pulumi.Input<string>;
    /**
     * The messages of the test case.
     */
    messages: pulumi.Input<pulumi.Input<inputs.RuleTestCaseMessage>[]>;
    /**
     * The name of the test case.
     */
    name: pulumi.Input<string>;
    /**
     * Sets if test case required to pass on update.
     */
    required?: pulumi.Input<boolean>;
}
