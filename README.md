# tflint-ruleset-matt-custom

This is a custom ruleset based on the [template repository](https://github.com/terraform-linters/tflint-ruleset-template). For additional reference, See also [Writing Plugins](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/plugins.md).

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|azurerm_storage_account_invalid_account_tier|Rule that checks if the account tier value passed in valid.|ERROR|||
|azurerm_resource_missing_tags|Checks against a list of resources to see if there are tags assigned to it|WARNING|||

## white_list_template.go.tpl

This template file can be used to generate rules that checks a resource against a list of values and throws errors if the values do not match exactly.

The placeholder array `whitelist_name` is the array of values the code checks against for valid data.

This template can also be modified to work as a blacklist as well by ensuring the values checked DO NOT match a value or to see if the string contains a substring that is blacklisted.

## Requirements

- TFLint v0.35+
- Go v1.18

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "template" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/terraform-linters/tflint-ruleset-template"

  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----
  mQINBGCqS2YBEADJ7gHktSV5NgUe08hD/uWWPwY07d5WZ1+F9I9SoiK/mtcNGz4P
  JLrYAIUTMBvrxk3I+kuwhp7MCk7CD/tRVkPRIklONgtKsp8jCke7FB3PuFlP/ptL
  SlbaXx53FCZSOzCJo9puZajVWydoGfnZi5apddd11Zw1FuJma3YElHZ1A1D2YvrF
  ...
  KEY
}
```

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```
