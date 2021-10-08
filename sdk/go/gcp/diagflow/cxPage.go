// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package diagflow

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Dialogflow CX conversation (session) can be described and visualized as a state machine. The states of a CX session are represented by pages.
//
// To get more information about Page, see:
//
// * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows.pages)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
//
// ## Example Usage
// ### Dialogflowcx Page Full
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/diagflow"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		agent, err := diagflow.NewCxAgent(ctx, "agent", &diagflow.CxAgentArgs{
// 			DisplayName:         pulumi.String("dialogflowcx-agent"),
// 			Location:            pulumi.String("global"),
// 			DefaultLanguageCode: pulumi.String("en"),
// 			SupportedLanguageCodes: pulumi.StringArray{
// 				pulumi.String("fr"),
// 				pulumi.String("de"),
// 				pulumi.String("es"),
// 			},
// 			TimeZone:                 pulumi.String("America/New_York"),
// 			Description:              pulumi.String("Example description."),
// 			AvatarUri:                pulumi.String("https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png"),
// 			EnableStackdriverLogging: pulumi.Bool(true),
// 			EnableSpellCorrection:    pulumi.Bool(true),
// 			SpeechToTextSettings: &diagflow.CxAgentSpeechToTextSettingsArgs{
// 				EnableSpeechAdaptation: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		myPage2, err := diagflow.NewCxPage(ctx, "myPage2", &diagflow.CxPageArgs{
// 			Parent:      agent.StartFlow,
// 			DisplayName: pulumi.String("MyPage2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = diagflow.NewCxPage(ctx, "basicPage", &diagflow.CxPageArgs{
// 			Parent:      agent.StartFlow,
// 			DisplayName: pulumi.String("MyPage"),
// 			EntryFulfillment: &diagflow.CxPageEntryFulfillmentArgs{
// 				Messages: diagflow.CxPageEntryFulfillmentMessageArray{
// 					&diagflow.CxPageEntryFulfillmentMessageArgs{
// 						Text: &diagflow.CxPageEntryFulfillmentMessageTextArgs{
// 							Texts: pulumi.StringArray{
// 								pulumi.String("Welcome to page"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Form: &diagflow.CxPageFormArgs{
// 				Parameters: diagflow.CxPageFormParameterArray{
// 					&diagflow.CxPageFormParameterArgs{
// 						DisplayName: pulumi.String("param1"),
// 						EntityType:  pulumi.String("projects/-/locations/-/agents/-/entityTypes/sys.date"),
// 						FillBehavior: &diagflow.CxPageFormParameterFillBehaviorArgs{
// 							InitialPromptFulfillment: &diagflow.CxPageFormParameterFillBehaviorInitialPromptFulfillmentArgs{
// 								Messages: diagflow.CxPageFormParameterFillBehaviorInitialPromptFulfillmentMessageArray{
// 									&diagflow.CxPageFormParameterFillBehaviorInitialPromptFulfillmentMessageArgs{
// 										Text: &diagflow.CxPageFormParameterFillBehaviorInitialPromptFulfillmentMessageTextArgs{
// 											Texts: pulumi.StringArray{
// 												pulumi.String("Please provide param1"),
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 						Required: pulumi.Bool(true),
// 						Redact:   pulumi.Bool(true),
// 					},
// 				},
// 			},
// 			TransitionRoutes: diagflow.CxPageTransitionRouteArray{
// 				&diagflow.CxPageTransitionRouteArgs{
// 					Condition: pulumi.String(fmt.Sprintf("%v%v", "$", "page.params.status = 'FINAL'")),
// 					TriggerFulfillment: &diagflow.CxPageTransitionRouteTriggerFulfillmentArgs{
// 						Messages: diagflow.CxPageTransitionRouteTriggerFulfillmentMessageArray{
// 							&diagflow.CxPageTransitionRouteTriggerFulfillmentMessageArgs{
// 								Text: &diagflow.CxPageTransitionRouteTriggerFulfillmentMessageTextArgs{
// 									Texts: pulumi.StringArray{
// 										pulumi.String("information completed, navigating to page 2"),
// 									},
// 								},
// 							},
// 						},
// 					},
// 					TargetPage: myPage2.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Page can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:diagflow/cxPage:CxPage default {{parent}}/pages/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:diagflow/cxPage:CxPage default {{parent}}/{{name}}
// ```
type CxPage struct {
	pulumi.CustomResourceState

	// The human-readable name of the parameter, unique within the form.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	// Structure is documented below.
	EntryFulfillment CxPageEntryFulfillmentPtrOutput `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	// Structure is documented below.
	EventHandlers CxPageEventHandlerArrayOutput `pulumi:"eventHandlers"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	// Structure is documented below.
	Form CxPageFormPtrOutput `pulumi:"form"`
	// The language of the following fields in page:
	// Page.entry_fulfillment.messages
	// Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages
	// Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages
	// Page.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrOutput `pulumi:"languageCode"`
	// -
	// The unique identifier of this event handler.
	Name pulumi.StringOutput `pulumi:"name"`
	// The flow to create a page for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.
	// If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route > page's transition route group > flow's transition routes.
	// If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups pulumi.StringArrayOutput `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	// When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in flow with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in the page with only condition specified.
	// TransitionRoutes defined in the transition route groups with only condition specified.
	// Structure is documented below.
	TransitionRoutes CxPageTransitionRouteArrayOutput `pulumi:"transitionRoutes"`
}

// NewCxPage registers a new resource with the given unique name, arguments, and options.
func NewCxPage(ctx *pulumi.Context,
	name string, args *CxPageArgs, opts ...pulumi.ResourceOption) (*CxPage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource CxPage
	err := ctx.RegisterResource("gcp:diagflow/cxPage:CxPage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCxPage gets an existing CxPage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCxPage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CxPageState, opts ...pulumi.ResourceOption) (*CxPage, error) {
	var resource CxPage
	err := ctx.ReadResource("gcp:diagflow/cxPage:CxPage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CxPage resources.
type cxPageState struct {
	// The human-readable name of the parameter, unique within the form.
	DisplayName *string `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	// Structure is documented below.
	EntryFulfillment *CxPageEntryFulfillment `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	// Structure is documented below.
	EventHandlers []CxPageEventHandler `pulumi:"eventHandlers"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	// Structure is documented below.
	Form *CxPageForm `pulumi:"form"`
	// The language of the following fields in page:
	// Page.entry_fulfillment.messages
	// Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages
	// Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages
	// Page.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	// -
	// The unique identifier of this event handler.
	Name *string `pulumi:"name"`
	// The flow to create a page for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent *string `pulumi:"parent"`
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.
	// If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route > page's transition route group > flow's transition routes.
	// If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	// When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in flow with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in the page with only condition specified.
	// TransitionRoutes defined in the transition route groups with only condition specified.
	// Structure is documented below.
	TransitionRoutes []CxPageTransitionRoute `pulumi:"transitionRoutes"`
}

type CxPageState struct {
	// The human-readable name of the parameter, unique within the form.
	DisplayName pulumi.StringPtrInput
	// The fulfillment to call when the session is entering the page.
	// Structure is documented below.
	EntryFulfillment CxPageEntryFulfillmentPtrInput
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	// Structure is documented below.
	EventHandlers CxPageEventHandlerArrayInput
	// The form associated with the page, used for collecting parameters relevant to the page.
	// Structure is documented below.
	Form CxPageFormPtrInput
	// The language of the following fields in page:
	// Page.entry_fulfillment.messages
	// Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages
	// Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages
	// Page.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	// -
	// The unique identifier of this event handler.
	Name pulumi.StringPtrInput
	// The flow to create a page for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent pulumi.StringPtrInput
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.
	// If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route > page's transition route group > flow's transition routes.
	// If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups pulumi.StringArrayInput
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	// When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in flow with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in the page with only condition specified.
	// TransitionRoutes defined in the transition route groups with only condition specified.
	// Structure is documented below.
	TransitionRoutes CxPageTransitionRouteArrayInput
}

func (CxPageState) ElementType() reflect.Type {
	return reflect.TypeOf((*cxPageState)(nil)).Elem()
}

type cxPageArgs struct {
	// The human-readable name of the parameter, unique within the form.
	DisplayName string `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	// Structure is documented below.
	EntryFulfillment *CxPageEntryFulfillment `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	// Structure is documented below.
	EventHandlers []CxPageEventHandler `pulumi:"eventHandlers"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	// Structure is documented below.
	Form *CxPageForm `pulumi:"form"`
	// The language of the following fields in page:
	// Page.entry_fulfillment.messages
	// Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages
	// Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages
	// Page.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	// The flow to create a page for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent *string `pulumi:"parent"`
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.
	// If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route > page's transition route group > flow's transition routes.
	// If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	// When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in flow with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in the page with only condition specified.
	// TransitionRoutes defined in the transition route groups with only condition specified.
	// Structure is documented below.
	TransitionRoutes []CxPageTransitionRoute `pulumi:"transitionRoutes"`
}

// The set of arguments for constructing a CxPage resource.
type CxPageArgs struct {
	// The human-readable name of the parameter, unique within the form.
	DisplayName pulumi.StringInput
	// The fulfillment to call when the session is entering the page.
	// Structure is documented below.
	EntryFulfillment CxPageEntryFulfillmentPtrInput
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	// Structure is documented below.
	EventHandlers CxPageEventHandlerArrayInput
	// The form associated with the page, used for collecting parameters relevant to the page.
	// Structure is documented below.
	Form CxPageFormPtrInput
	// The language of the following fields in page:
	// Page.entry_fulfillment.messages
	// Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages
	// Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages
	// Page.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	// The flow to create a page for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent pulumi.StringPtrInput
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.
	// If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route > page's transition route group > flow's transition routes.
	// If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups pulumi.StringArrayInput
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	// When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in flow with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in the page with only condition specified.
	// TransitionRoutes defined in the transition route groups with only condition specified.
	// Structure is documented below.
	TransitionRoutes CxPageTransitionRouteArrayInput
}

func (CxPageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cxPageArgs)(nil)).Elem()
}

type CxPageInput interface {
	pulumi.Input

	ToCxPageOutput() CxPageOutput
	ToCxPageOutputWithContext(ctx context.Context) CxPageOutput
}

func (*CxPage) ElementType() reflect.Type {
	return reflect.TypeOf((*CxPage)(nil))
}

func (i *CxPage) ToCxPageOutput() CxPageOutput {
	return i.ToCxPageOutputWithContext(context.Background())
}

func (i *CxPage) ToCxPageOutputWithContext(ctx context.Context) CxPageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxPageOutput)
}

func (i *CxPage) ToCxPagePtrOutput() CxPagePtrOutput {
	return i.ToCxPagePtrOutputWithContext(context.Background())
}

func (i *CxPage) ToCxPagePtrOutputWithContext(ctx context.Context) CxPagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxPagePtrOutput)
}

type CxPagePtrInput interface {
	pulumi.Input

	ToCxPagePtrOutput() CxPagePtrOutput
	ToCxPagePtrOutputWithContext(ctx context.Context) CxPagePtrOutput
}

type cxPagePtrType CxPageArgs

func (*cxPagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CxPage)(nil))
}

func (i *cxPagePtrType) ToCxPagePtrOutput() CxPagePtrOutput {
	return i.ToCxPagePtrOutputWithContext(context.Background())
}

func (i *cxPagePtrType) ToCxPagePtrOutputWithContext(ctx context.Context) CxPagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxPagePtrOutput)
}

// CxPageArrayInput is an input type that accepts CxPageArray and CxPageArrayOutput values.
// You can construct a concrete instance of `CxPageArrayInput` via:
//
//          CxPageArray{ CxPageArgs{...} }
type CxPageArrayInput interface {
	pulumi.Input

	ToCxPageArrayOutput() CxPageArrayOutput
	ToCxPageArrayOutputWithContext(context.Context) CxPageArrayOutput
}

type CxPageArray []CxPageInput

func (CxPageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CxPage)(nil)).Elem()
}

func (i CxPageArray) ToCxPageArrayOutput() CxPageArrayOutput {
	return i.ToCxPageArrayOutputWithContext(context.Background())
}

func (i CxPageArray) ToCxPageArrayOutputWithContext(ctx context.Context) CxPageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxPageArrayOutput)
}

// CxPageMapInput is an input type that accepts CxPageMap and CxPageMapOutput values.
// You can construct a concrete instance of `CxPageMapInput` via:
//
//          CxPageMap{ "key": CxPageArgs{...} }
type CxPageMapInput interface {
	pulumi.Input

	ToCxPageMapOutput() CxPageMapOutput
	ToCxPageMapOutputWithContext(context.Context) CxPageMapOutput
}

type CxPageMap map[string]CxPageInput

func (CxPageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CxPage)(nil)).Elem()
}

func (i CxPageMap) ToCxPageMapOutput() CxPageMapOutput {
	return i.ToCxPageMapOutputWithContext(context.Background())
}

func (i CxPageMap) ToCxPageMapOutputWithContext(ctx context.Context) CxPageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxPageMapOutput)
}

type CxPageOutput struct{ *pulumi.OutputState }

func (CxPageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CxPage)(nil))
}

func (o CxPageOutput) ToCxPageOutput() CxPageOutput {
	return o
}

func (o CxPageOutput) ToCxPageOutputWithContext(ctx context.Context) CxPageOutput {
	return o
}

func (o CxPageOutput) ToCxPagePtrOutput() CxPagePtrOutput {
	return o.ToCxPagePtrOutputWithContext(context.Background())
}

func (o CxPageOutput) ToCxPagePtrOutputWithContext(ctx context.Context) CxPagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CxPage) *CxPage {
		return &v
	}).(CxPagePtrOutput)
}

type CxPagePtrOutput struct{ *pulumi.OutputState }

func (CxPagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CxPage)(nil))
}

func (o CxPagePtrOutput) ToCxPagePtrOutput() CxPagePtrOutput {
	return o
}

func (o CxPagePtrOutput) ToCxPagePtrOutputWithContext(ctx context.Context) CxPagePtrOutput {
	return o
}

func (o CxPagePtrOutput) Elem() CxPageOutput {
	return o.ApplyT(func(v *CxPage) CxPage {
		if v != nil {
			return *v
		}
		var ret CxPage
		return ret
	}).(CxPageOutput)
}

type CxPageArrayOutput struct{ *pulumi.OutputState }

func (CxPageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CxPage)(nil))
}

func (o CxPageArrayOutput) ToCxPageArrayOutput() CxPageArrayOutput {
	return o
}

func (o CxPageArrayOutput) ToCxPageArrayOutputWithContext(ctx context.Context) CxPageArrayOutput {
	return o
}

func (o CxPageArrayOutput) Index(i pulumi.IntInput) CxPageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CxPage {
		return vs[0].([]CxPage)[vs[1].(int)]
	}).(CxPageOutput)
}

type CxPageMapOutput struct{ *pulumi.OutputState }

func (CxPageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CxPage)(nil))
}

func (o CxPageMapOutput) ToCxPageMapOutput() CxPageMapOutput {
	return o
}

func (o CxPageMapOutput) ToCxPageMapOutputWithContext(ctx context.Context) CxPageMapOutput {
	return o
}

func (o CxPageMapOutput) MapIndex(k pulumi.StringInput) CxPageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CxPage {
		return vs[0].(map[string]CxPage)[vs[1].(string)]
	}).(CxPageOutput)
}

func init() {
	pulumi.RegisterOutputType(CxPageOutput{})
	pulumi.RegisterOutputType(CxPagePtrOutput{})
	pulumi.RegisterOutputType(CxPageArrayOutput{})
	pulumi.RegisterOutputType(CxPageMapOutput{})
}