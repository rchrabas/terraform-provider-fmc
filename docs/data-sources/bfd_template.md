---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_bfd_template Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source can read the BFD Template.
---

# fmc_bfd_template (Data Source)

This data source can read the BFD Template.

## Example Usage

```terraform
data "fmc_bfd_template" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) The name of the FMC domain
- `id` (String) The id of the object
- `name` (String) The name of the bfd template object.

### Read-Only

- `authentication_key_id` (Number) Authentication Key ID
- `authentication_password` (String) Password for BFD Authentication (1-24 characters)
- `authentication_type` (String) Authentication types
- `echo` (String) Enables/disables BFD echo.
- `hop_type` (String) The hop type.
- `interval_time` (String) Interval unit of measurement of time.
- `min_receive` (Number) BFD Minimum Receive unit value in ranges: 50-999 miliseconds, 50000-999000 microseconds
- `min_transmit` (Number) BFD Minimum Transmit unit value.
- `tx_rx_multiplier` (Number) BFD Multipler value.
- `type` (String) Type of the object; this value is always 'BFDTemplate'.