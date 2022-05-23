package rules

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

const (
	filenameLocals    = "locals.tf"
	filenameData      = "data.tf"
)

// TerraformCustomisedModuleStructureRule adds additional checks on top of Terraform's standard module structure, i.e. That data blocks are located in data.tf and local blocks should be located in locals.tf
type TerraformCustomisedModuleStructureRule struct{}

// NewTerraformCustomisedModuleStructureRule returns a new rule
func NewTerraformCustomisedModuleStructureRule() *TerraformCustomisedModuleStructureRule {
	return &TerraformCustomisedModuleStructureRule{}
}

// Name returns the rule name
func (r *TerraformCustomisedModuleStructureRule) Name() string {
	return "terraform_customised_module_structure"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformCustomisedModuleStructureRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *TerraformCustomisedModuleStructureRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *TerraformCustomisedModuleStructureRule) Link() string {
	return tflint.ReferenceLink(r.Name())
}

// Check emits errors for any missing files and any block types that are included in the wrong file
func (r *TerraformCustomisedModuleStructureRule) Check(runner *tflint.Runner) error {
	if !runner.TFConfig.Path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}

	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	r.checkFiles(runner)
	r.checkData(runner)
	r.checkLocals(runner)

	return nil
}

func (r *TerraformCustomisedModuleStructureRule) checkFiles(runner *tflint.Runner) {
	if r.onlyJSON(runner) {
		return
	}

	f := runner.Files()
	files := make(map[string]*hcl.File, len(f))
	for name, file := range f {
		files[filepath.Base(name)] = file
	}

	log.Printf("[DEBUG] %d files found: %v", len(files), files)

	if files[filenameData] == nil && len(runner.TFConfig.Module.Outputs) == 0 {
		runner.EmitIssue(
			r,
			fmt.Sprintf("Module should include an empty %s file", filenameData),
			hcl.Range{
				Filename: filepath.Join(runner.TFConfig.Module.SourceDir, filenameData),
				Start:    hcl.InitialPos,
			},
		)
	}

	if files[filenameLocals] == nil && len(runner.TFConfig.Module.Locals) == 0 {
		runner.EmitIssue(
			r,
			fmt.Sprintf("Module should include an empty %s file", filenameLocals),
			hcl.Range{
				Filename: filepath.Join(runner.TFConfig.Module.SourceDir, filenameLocals),
				Start:    hcl.InitialPos,
			},
		)
	}
}

func (r *TerraformCustomisedModuleStructureRule) checkData(runner *tflint.Runner) {
	for _, variable := range runner.TFConfig.Module.Outputs {
		if filename := variable.DeclRange.Filename; r.shouldMove(filename, filenameData) {
			runner.EmitIssue(
				r,
				fmt.Sprintf("data %q should be moved from %s to %s", variable.Name, filename, filenameData),
				variable.DeclRange,
			)
		}
	}
}

func (r *TerraformCustomisedModuleStructureRule) checkLocals(runner *tflint.Runner) {
	for _, variable := range runner.TFConfig.Module.Locals {
		if filename := variable.DeclRange.Filename; r.shouldMove(filename, filenameLocals) {
			runner.EmitIssue(
				r,
				fmt.Sprintf("local %q should be moved from %s to %s", variable.Name, filename, filenameLocals),
				variable.DeclRange,
			)
		}
	}
}

func (r *TerraformCustomisedModuleStructureRule) onlyJSON(runner *tflint.Runner) bool {
	files := runner.Files()

	if len(files) == 0 {
		return false
	}

	for filename := range files {
		if filepath.Ext(filename) != ".json" {
			return false
		}
	}

	return true
}

func (r *TerraformCustomisedModuleStructureRule) shouldMove(path string, expected string) bool {
	// json files are likely generated and conventional filenames do not apply
	if filepath.Ext(path) == ".json" {
		return false
	}

	return filepath.Base(path) != expected
}