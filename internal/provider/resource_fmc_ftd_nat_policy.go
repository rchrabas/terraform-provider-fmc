// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &FTDNATPolicyResource{}
	_ resource.ResourceWithImportState = &FTDNATPolicyResource{}
)

func NewFTDNATPolicyResource() resource.Resource {
	return &FTDNATPolicyResource{}
}

type FTDNATPolicyResource struct {
	client *fmc.Client
}

func (r *FTDNATPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftd_nat_policy"
}

func (r *FTDNATPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a FTD NAT Policy.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the ftd network address translation (nat) policy.").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description").String,
				Optional:            true,
			},
			"manual_nat_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ordered list of manual NAT rules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Identifier of the manual nat rule.").String,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("User-specified unique string.").String,
							Required:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("My manual nat rule 1").String,
							Optional:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates if the rule is enabled.").String,
							Optional:            true,
						},
						"section": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("To which section the rule belongs.").AddStringEnumDescription("before_auto", "after_auto").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("before_auto", "after_auto"),
							},
						},
						"fall_through": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"nat_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of the rule").AddStringEnumDescription("STATIC", "DYNAMIC").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("STATIC", "DYNAMIC"),
							},
						},
						"interface_in_original_destination": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"interface_in_translated_source": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"ipv6": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"net_to_net": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"no_proxy_arp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"unidirectional": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the rule is unidirectional").String,
							Optional:            true,
						},
						"source_interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of source security zone").String,
							Optional:            true,
						},
						"original_source_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of original source network object").String,
							Optional:            true,
						},
						"original_source_port_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of original source port object").String,
							Optional:            true,
						},
						"original_destination_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of original destination network object").String,
							Optional:            true,
						},
						"original_destination_port_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of original destination port object").String,
							Optional:            true,
						},
						"route_lookup": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"destination_interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of destination security zone").String,
							Optional:            true,
						},
						"translated_source_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of translated source network object").String,
							Optional:            true,
						},
						"translated_source_port_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of translated source port object").String,
							Optional:            true,
						},
						"translate_dns": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Perform translation of addresses in DNS packets").String,
							Optional:            true,
						},
						"translated_destination_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of translated destination network object").String,
							Optional:            true,
						},
						"translated_destination_port_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of translated destination port object").String,
							Optional:            true,
						},
					},
				},
			},
			"auto_nat_rules": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ordered list of manual NAT rules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Identifier of the auto nat rule.").String,
							Computed:            true,
						},
						"nat_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of the rule").AddStringEnumDescription("STATIC", "DYNAMIC").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("STATIC", "DYNAMIC"),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("My auto nat rule 1").String,
							Optional:            true,
						},
						"destination_interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of destination security zone").String,
							Optional:            true,
						},
						"fall_through": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"ipv6": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"net_to_net": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"no_proxy_arp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"original_network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of original network object").String,
							Optional:            true,
						},
						"original_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Original port number").String,
							Optional:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocl of the service").AddStringEnumDescription("TCP", "UDP").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("TCP", "UDP"),
							},
						},
						"perform_route_lookup": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"source_interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of source security zone").String,
							Optional:            true,
						},
						"translate_dns": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Perform address translation in DNS packets").String,
							Optional:            true,
						},
						"translated_network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of translated network object").String,
							Optional:            true,
						},
						"translated_network_is_destination_interface": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TBD???????????").String,
							Optional:            true,
						},
						"translated_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translated port number").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *FTDNATPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *FTDNATPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FTDNATPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, FTDNATPolicy{})
	res, err := r.client.Post(plan.getPath(), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	res, err = r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *FTDNATPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FTDNATPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	urlPath := state.getPath() + "/" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get(urlPath, reqMods...)

	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *FTDNATPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FTDNATPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !plan.Domain.IsNull() && plan.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(plan.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx, state)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	res, err = r.client.Get(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *FTDNATPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FTDNATPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !state.Domain.IsNull() && state.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(state.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import

func (r *FTDNATPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin createSubresources

// End of section. //template:end createSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin deleteSubresources

// End of section. //template:end deleteSubresources

// Section below is generated&owned by "gen/generator.go". //template:begin updateSubresources

// End of section. //template:end updateSubresources
