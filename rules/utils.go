package rules

// Used for the Storage Account
var validAccountTier = []string{
	"Standard",
	"Premium",
}

// Used for checking tags
// It may be possible to generate the resources list dynamically by further investigating the code here: https://github.com/terraform-linters/tflint-ruleset-aws/tree/master/rules/tags
var Resources = []string{
	"azurerm_resource_group",
	"azurerm_key_vault",
}