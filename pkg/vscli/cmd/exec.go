///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	fnrunner "gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/client/runner"
	models "gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/models"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/vscli/i18n"
)

var (
	execLong = i18n.T(`Execute a serverless function.`)

	// TODO: Add examples
	execExample = i18n.T(``)

	execWait  = false
	execInput = "{}"
)

// NewCmdExec creates a command to execute a serverless function.
func NewCmdExec(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "exec [--wait] [--input JSON] FUNCTION_NAME",
		Short:   i18n.T("Execute a serverless function"),
		Long:    execLong,
		Example: execExample,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := runExec(out, errOut, cmd, args)
			CheckErr(err)
		},
		PreRunE: validateFnExecFunc(errOut),
	}
	cmd.Flags().BoolVar(&execWait, "wait", false, "Wait for the function to complete execution.")
	cmd.Flags().StringVar(&execInput, "input", "{}", "Function input JSON object")
	return cmd
}

func validateFnExecFunc(errOut io.Writer) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return nil
	}
}

func runExec(out, errOut io.Writer, cmd *cobra.Command, args []string) error {
	functionName := args[0]
	var input map[string]interface{}
	err := json.Unmarshal([]byte(execInput), &input)
	if err != nil {
		fmt.Fprintf(errOut, "Error when parsing function parameters %s\n", execInput)
		return err
	}
	run := &models.Run{
		Blocking: execWait,
		Input:    input,
	}

	params := &fnrunner.RunFunctionParams{
		Body:         run,
		Context:      context.Background(),
		FunctionName: functionName,
	}
	client := functionManagerClient()
	executed, executing, err := client.Runner.RunFunction(params)
	if err != nil {
		fmt.Fprintf(errOut, "Error when running a function %s\n", functionName)
		return err
	}
	if executed != nil {
		fmt.Fprintf(out, "Function %s finished successfully.\n", functionName)
		bytes, _ := json.Marshal(executed.Payload.Output)
		fmt.Println(string(bytes))
	}
	if executing != nil {
		fmt.Fprintf(out, "Function %s started\n", functionName)
	}

	return nil
}