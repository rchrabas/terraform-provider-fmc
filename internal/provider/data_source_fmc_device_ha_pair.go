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

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceHAPairDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceHAPairDataSource{}
)

func NewDeviceHAPairDataSource() datasource.DataSource {
	return &DeviceHAPairDataSource{}
}

type DeviceHAPairDataSource struct {
	client *fmc.Client
}

func (d *DeviceHAPairDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ha_pair"
}

func (d *DeviceHAPairDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Device HA Pair.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the High Availability (HA) Pair.",
				Optional:            true,
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the resource; This is always `DeviceHAPair`.",
				Computed:            true,
			},
			"primary_device_id": schema.StringAttribute{
				MarkdownDescription: "ID of primary FTD in the HA Pair.",
				Computed:            true,
			},
			"secondary_device_id": schema.StringAttribute{
				MarkdownDescription: "ID of secondary FTD in the HA Pair.",
				Computed:            true,
			},
			"ha_link_interface_id": schema.StringAttribute{
				MarkdownDescription: "ID of High Availability Link interface.",
				Computed:            true,
			},
			"ha_link_interface_name": schema.StringAttribute{
				MarkdownDescription: "Name of High Availability Link interface.",
				Computed:            true,
			},
			"ha_link_interface_type": schema.StringAttribute{
				MarkdownDescription: "Type of High Availability Link interface.",
				Computed:            true,
			},
			"ha_link_logical_name": schema.StringAttribute{
				MarkdownDescription: "The logical name of failover interface.",
				Computed:            true,
			},
			"ha_link_use_ipv6": schema.BoolAttribute{
				MarkdownDescription: "Use IPv6 addressing for HA communication.",
				Computed:            true,
			},
			"ha_link_primary_ip": schema.StringAttribute{
				MarkdownDescription: "The IP of primary node interface.",
				Computed:            true,
			},
			"ha_link_secondary_ip": schema.StringAttribute{
				MarkdownDescription: "The IP of secondary node interface.",
				Computed:            true,
			},
			"ha_link_netmask": schema.StringAttribute{
				MarkdownDescription: "Subnet mask for HA link.",
				Computed:            true,
			},
			"state_link_use_same_as_ha": schema.BoolAttribute{
				MarkdownDescription: "Use the same link for state and HA.",
				Computed:            true,
			},
			"state_link_interface_id": schema.StringAttribute{
				MarkdownDescription: "ID of physical interface.",
				Computed:            true,
			},
			"state_link_interface_name": schema.StringAttribute{
				MarkdownDescription: "Name of state link interface.",
				Computed:            true,
			},
			"state_link_interface_type": schema.StringAttribute{
				MarkdownDescription: "Type of state link interface.",
				Computed:            true,
			},
			"state_link_logical_name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"state_link_use_ipv6": schema.BoolAttribute{
				MarkdownDescription: "Use IPv6 addressing for HA communication.",
				Computed:            true,
			},
			"state_link_primary_ip": schema.StringAttribute{
				MarkdownDescription: "The IP of primary node interface.",
				Computed:            true,
			},
			"state_link_secondary_ip": schema.StringAttribute{
				MarkdownDescription: "The IP of secondary node interface.",
				Computed:            true,
			},
			"state_link_netmask": schema.StringAttribute{
				MarkdownDescription: "Subnet mask for state link.",
				Computed:            true,
			},
			"encryption_enabled": schema.BoolAttribute{
				MarkdownDescription: "Use encryption for communication.",
				Computed:            true,
			},
			"encryption_key_generation_scheme": schema.StringAttribute{
				MarkdownDescription: "Select the encyption key generation scheme.",
				Computed:            true,
			},
			"encryption_key": schema.StringAttribute{
				MarkdownDescription: "Pass shared key for encryption if CUSTOM key geneeration scheme is selected.",
				Computed:            true,
			},
			"failed_interfaces_percent": schema.Int64Attribute{
				MarkdownDescription: "Percentage of Failed Interfaces that triggers failover.",
				Computed:            true,
			},
			"failed_interfaces_limit": schema.Int64Attribute{
				MarkdownDescription: "Number of Failed Interfaces that triggers failover.",
				Computed:            true,
			},
			"peer_poll_time": schema.Int64Attribute{
				MarkdownDescription: "Peer Pool Time (1-15 SEC or 200-999 MSEC)",
				Computed:            true,
			},
			"peer_poll_time_unit": schema.StringAttribute{
				MarkdownDescription: "Peer Pool Time Unit",
				Computed:            true,
			},
			"peer_hold_time": schema.Int64Attribute{
				MarkdownDescription: "Peer Hold Time (3-45 SEC or 800-999 MSEC)",
				Computed:            true,
			},
			"peer_hold_time_unit": schema.StringAttribute{
				MarkdownDescription: "Peer Hold Time Unit",
				Computed:            true,
			},
			"interface_poll_time": schema.Int64Attribute{
				MarkdownDescription: "Peer Pool Time (1-15 SEC or 500-999 MSEC)",
				Computed:            true,
			},
			"interface_poll_time_unit": schema.StringAttribute{
				MarkdownDescription: "Peer Pool Time Unit",
				Computed:            true,
			},
			"interface_hold_time": schema.Int64Attribute{
				MarkdownDescription: "Interface Hold Time in seconds",
				Computed:            true,
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "FTD HA PUT operation action. Specifically used for manual switch. HA Break will be triggered when you run terraform destroy",
				Computed:            true,
			},
		},
	}
}
func (d *DeviceHAPairDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *DeviceHAPairDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceHAPairDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceHAPair

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !config.Domain.IsNull() && config.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(config.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d&expanded=true", limit, offset)
			res, err := d.client.Get(config.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
						return false
					}
					return true
				})
			}
			if !config.Id.IsNull() || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}
	urlPath := config.getPath() + "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read