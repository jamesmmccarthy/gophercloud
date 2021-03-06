package v2

import (
	"fmt"
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/workflow/v2/executions"
	"github.com/gophercloud/gophercloud/openstack/workflow/v2/workflows"
	th "github.com/gophercloud/gophercloud/testhelper"
)

// CreateExecution creates an execution for the given workflow.
func CreateExecution(t *testing.T, client *gophercloud.ServiceClient, workflow *workflows.Workflow) (*executions.Execution, error) {
	executionDescription := tools.RandomString("execution_", 5)

	t.Logf("Attempting to create execution: %s", executionDescription)
	createOpts := executions.CreateOpts{
		ID:                executionDescription,
		WorkflowID:        workflow.ID,
		WorkflowNamespace: workflow.Namespace,
		Description:       executionDescription,
		Input: map[string]interface{}{
			"msg": "Hello World!",
		},
	}
	execution, err := executions.Create(client, createOpts).Extract()
	if err != nil {
		return execution, err
	}

	t.Logf("Execution created: %s", executionDescription)

	th.AssertEquals(t, execution.Description, executionDescription)

	t.Logf("Wait for execution status SUCCESS: %s", executionDescription)
	th.AssertNoErr(t, tools.WaitFor(func() (bool, error) {
		latest, err := executions.Get(client, execution.ID).Extract()
		if err != nil {
			return false, err
		}

		if latest.State == "SUCCESS" {
			execution = latest
			return true, nil
		}

		if latest.State == "ERROR" {
			return false, fmt.Errorf("Execution in ERROR state")
		}

		return false, nil
	}))
	t.Logf("Execution success: %s", executionDescription)

	return execution, nil
}
