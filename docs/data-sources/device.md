---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_device Data Source - terraform-provider-fmc"
subcategory: "Devices"
description: |-
  This data source can read the Device.
---

# fmc_device (Data Source)

This data source can read the Device.

## Example Usage

```terraform
data "fmc_device" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) The name of the FMC domain
- `id` (String) The id of the object
- `name` (String) User-specified name, must be unique. Example: 'Device 01 - 192.168.0.152'

### Read-Only

- `access_policy_id` (String) The UUID of the assigned access control policy. For example `fmc_access_control_policy.example.id`.
- `host_name` (String) Hostname or IP address of the device. Either the host_name or nat_id must be present.
- `license_capabilities` (Set of String) Array of strings representing the license capabilities on the managed device. For registering FTD, the allowed values are: BASE (mandatory), THREAT, URLFilter, MALWARE, APEX, PLUS, VPNOnly. For Firepower ASA or NGIPSv devices, allowed values are: BASE, THREAT, PROTECT, CONTROL, URLFilter, MALWARE, VPN, SSL.
- `nat_id` (String) (used for device registration behind NAT) If the device to be registered and the Firepower Management Center are separated by network address translation (NAT), set a unique string identifier.
- `nat_policy_id` (String) The UUID of the assigned NAT policy.
- `performance_tier` (String) Performance tier for the managed device, applicable only to vFTD devices >=6.8.0.
- `prohibit_packet_transfer` (Boolean) Value true prohibits the device from sending packet data with events to the Firepower Management Center. Value false allows the transfer when a certain event is triggered. Not all traffic data is sent; connection events do not include a payload, only connection metadata.
- `registration_key` (String) Registration Key identical to the one previously configured on the device (`configure manager`).
- `type` (String) Type of the device; this value is always 'Device'.
