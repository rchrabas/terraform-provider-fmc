---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_device_bgp_general_settings Resource - terraform-provider-fmc"
subcategory: "Devices"
description: |-
  This resource can manage a Device BGP General Settings.
---

# fmc_device_bgp_general_settings (Resource)

This resource can manage a Device BGP General Settings.

## Example Usage

```terraform
resource "fmc_device_bgp_general_settings" "example" {
  device_id                   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  as_number                   = "65535"
  router_id                   = "AUTOMATIC"
  scanning_interval           = 60
  as_number_in_path_attribute = 50
  tcp_path_mtu_discovery      = true
  reset_session_upon_failover = true
  enforce_first_peer_as       = true
  use_dot_notation            = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `as_number` (String) Autonomous System (AS) number in asplain or asdot format
- `device_id` (String) UUID of the parent device (fmc_device.example.id).

### Optional

- `aggregate_timer` (Number) Interval at which BGP routes will be aggregated or to disable timer-based router aggregation (in seconds).
  - Range: `6`-`60`
- `as_number_in_path_attribute` (Number) Range to discard routes that have as-path segments that exceed a specified value.
  - Range: `1`-`254`
- `compare_med_from_different_neighbors` (Boolean) Allow comparing MED from different neighbors
- `compare_router_id_in_path` (Boolean) Compare Router ID for identical EBGP paths
- `default_local_preference` (Number) Default local preference
  - Range: `0`-`4294967295`
- `domain` (String) The name of the FMC domain
- `enforce_first_peer_as` (Boolean) Discard updates received from an external BGP (eBGP) peers that do not list their autonomous system (AS) number.
- `graceful_restart` (Boolean) Enable graceful restart
- `graceful_restart_restart_time` (Number) Graceful Restart Time in seconds
  - Range: `1`-`3600`
- `graceful_restart_stale_path_time` (Number) Stalepath Time in seconds
  - Range: `1`-`3600`
- `hold_time` (Number) Hold time in seconds
  - Range: `0`-`65535`
- `keepalive_interval` (Number) Keepalive interval in seconds
  - Range: `0`-`65535`
- `log_neighbor_changes` (Boolean) Enable logging of BGP neighbor status changes.
- `min_hold_time` (Number) Minimum hold time (0 or 3-65535 seconds)
  - Range: `0`-`65535`
- `missing_med_as_best` (Boolean) Treat missing MED as the best preferred path
- `next_hop_address_tracking` (Boolean) Enable next hop address tracking
- `next_hop_delay_interval` (Number) Next hop delay interval in seconds
  - Range: `0`-`100`
- `pick_best_med` (Boolean) Pick the best-MED path among paths advertised by neighbor AS
- `reset_session_upon_failover` (Boolean) Reset session upon failover
- `router_id` (String) String value for the routerID. Possible values can be 'AUTOMATIC' or valid ipv4 address
- `scanning_interval` (Number) Scanning interval of BGP routers for next hop validation in Seconds.
  - Range: `5`-`60`
- `tcp_path_mtu_discovery` (Boolean) Use TCP path MTU discovery.
- `use_dot_notation` (Boolean) Change format of BGP 4-byte autonomous system numbers from asplain (decimal values) to dot notation.

### Read-Only

- `id` (String) The id of the object
- `name` (String) Object name; Always set to 'AsaBGPGeneralTable'

## Import

Import is supported using the following syntax:

```shell
terraform import fmc_device_bgp_general_settings.example "<device_id>,<id>"
```