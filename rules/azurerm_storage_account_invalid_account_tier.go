package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermStorageAccountInvalidAccountTierRule checks the pattern is valid
type AzurermStorageAccountInvalidAccountTierRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewAzurermStorageAccountInvalidAccountTierRule returns new rule with default attributes
func NewAzurermStorageAccountInvalidAccountTierRule() *AzurermStorageAccountInvalidAccountTierRule {
	return &AzurermStorageAccountInvalidAccountTierRule{
		resourceType:  "azurerm_storage_account",
		attributeName: "account_tier",
	}
}

// Name returns the rule name
func (r *AzurermStorageAccountInvalidAccountTierRule) Name() string {
	return "azurerm_storage_account_invalid_account_tier"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermStorageAccountInvalidAccountTierRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *AzurermStorageAccountInvalidAccountTierRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermStorageAccountInvalidAccountTierRule) Link() string {
	//return project.ReferenceLink(r.Name())
	return ""
}

// Check checks the pattern is valid
func (r *AzurermStorageAccountInvalidAccountTierRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range validAccountTier {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as Account Tier`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}