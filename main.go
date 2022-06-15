package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/ecsd-matthew-song/tflint-ruleset-matt-custom/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "matt-custom",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewAzurermResourceMissingTagsRule(),
				rules.NewAzurermStorageAccountInvalidAccountTierRule(),
			},
		},
	})
}