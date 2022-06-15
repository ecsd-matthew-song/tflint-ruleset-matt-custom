package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// Replace RuleTemplateName with the name of the Rule

// Provide a description for what the rule does
type RuleTemplateNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewRuleTemplateNameRule returns new rule with default attributes
func NewRuleTemplateNameRule() *RuleTemplateNameRule {
	return &RuleTemplateNameRule{
		resourceType:  "", // What is the resource being checked
		attributeName: "", // What is the attribute being checked
	}
}

// Name returns the rule name
func (r *RuleTemplateNameRule) Name() string {
	return "rule_template_name"
}

// Enabled returns whether the rule is enabled by default
func (r *RuleTemplateNameRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *RuleTemplateNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link if there is a set of documents
func (r *RuleTemplateNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *RuleTemplateNameRule) Check(runner tflint.Runner) error {
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
			for _, item := range whitelist_name { // This is the whitelist variable that will be checked against to see if the value is valid
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as #ATTRIBUTE_NAME_TO_BE_UPDATED`, val),
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