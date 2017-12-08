// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestExamples(t *testing.T) {
	environ := os.Getenv("AZURE_ENVIRONMENT")
	if environ == "" {
		t.Skipf("Skipping test due to missing AWS_REGION environment variable")
	}
	cwd, err := os.Getwd()
	if !assert.NoError(t, err, "expected a valid working directory: %v", err) {
		return
	}

	// base options shared amongst all tests.
	base := integration.ProgramTestOptions{
		Config: map[string]string{
			"azurerm:config:environment": environ,
		},
		Dependencies: []string{
			"pulumi",
			"@pulumi/azurerm",
		},
	}

	examples := []integration.ProgramTestOptions{
		base.With(integration.ProgramTestOptions{Dir: path.Join(cwd, "minimal")}),
	}
	if !testing.Short() {
		examples = append(examples, []integration.ProgramTestOptions{
			base.With(integration.ProgramTestOptions{Dir: path.Join(cwd, "webserver")}),
		}...)
	}

	for _, ex := range examples {
		example := ex
		t.Run(example.Dir, func(t *testing.T) {
			integration.ProgramTest(t, example)
		})
	}
}

func createEditDir(dir string) integration.EditDir {
	return integration.EditDir{Dir: dir, ExtraRuntimeValidation: nil}
}
