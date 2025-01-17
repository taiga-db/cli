// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package access_control

import (
	"fmt"

	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/cli/libs/flags"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/spf13/cobra"
)

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var cmdOverrides []func(*cobra.Command)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "access-control",
		Short: `These APIs manage access rules on resources in an account.`,
		Long: `These APIs manage access rules on resources in an account. Currently, only
  grant rules are supported. A grant rule specifies a role assigned to a set of
  principals. A list of rules attached to a resource is called a rule set.`,
		GroupID: "iam",
		Annotations: map[string]string{
			"package": "iam",
		},
	}

	// Apply optional overrides to this command.
	for _, fn := range cmdOverrides {
		fn(cmd)
	}

	return cmd
}

// start get-assignable-roles-for-resource command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var getAssignableRolesForResourceOverrides []func(
	*cobra.Command,
	*iam.GetAssignableRolesForResourceRequest,
)

func newGetAssignableRolesForResource() *cobra.Command {
	cmd := &cobra.Command{}

	var getAssignableRolesForResourceReq iam.GetAssignableRolesForResourceRequest

	// TODO: short flags

	cmd.Use = "get-assignable-roles-for-resource RESOURCE"
	cmd.Short = `Get assignable roles for a resource.`
	cmd.Long = `Get assignable roles for a resource.
  
  Gets all the roles that can be granted on an account level resource. A role is
  grantable if the rule set on the resource can contain an access rule of the
  role.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustAccountClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		a := root.AccountClient(ctx)

		getAssignableRolesForResourceReq.Resource = args[0]

		response, err := a.AccessControl.GetAssignableRolesForResource(ctx, getAssignableRolesForResourceReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range getAssignableRolesForResourceOverrides {
		fn(cmd, &getAssignableRolesForResourceReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newGetAssignableRolesForResource())
	})
}

// start get-rule-set command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var getRuleSetOverrides []func(
	*cobra.Command,
	*iam.GetRuleSetRequest,
)

func newGetRuleSet() *cobra.Command {
	cmd := &cobra.Command{}

	var getRuleSetReq iam.GetRuleSetRequest

	// TODO: short flags

	cmd.Use = "get-rule-set NAME ETAG"
	cmd.Short = `Get a rule set.`
	cmd.Long = `Get a rule set.
  
  Get a rule set by its name. A rule set is always attached to a resource and
  contains a list of access rules on the said resource. Currently only a default
  rule set for each resource is supported.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(2)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustAccountClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		a := root.AccountClient(ctx)

		getRuleSetReq.Name = args[0]
		getRuleSetReq.Etag = args[1]

		response, err := a.AccessControl.GetRuleSet(ctx, getRuleSetReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range getRuleSetOverrides {
		fn(cmd, &getRuleSetReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newGetRuleSet())
	})
}

// start update-rule-set command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var updateRuleSetOverrides []func(
	*cobra.Command,
	*iam.UpdateRuleSetRequest,
)

func newUpdateRuleSet() *cobra.Command {
	cmd := &cobra.Command{}

	var updateRuleSetReq iam.UpdateRuleSetRequest
	var updateRuleSetJson flags.JsonFlag

	// TODO: short flags
	cmd.Flags().Var(&updateRuleSetJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	cmd.Use = "update-rule-set"
	cmd.Short = `Update a rule set.`
	cmd.Long = `Update a rule set.
  
  Replace the rules of a rule set. First, use get to read the current version of
  the rule set before modifying it. This pattern helps prevent conflicts between
  concurrent updates.`

	cmd.Annotations = make(map[string]string)

	cmd.PreRunE = root.MustAccountClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		a := root.AccountClient(ctx)

		if cmd.Flags().Changed("json") {
			err = updateRuleSetJson.Unmarshal(&updateRuleSetReq)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("please provide command input in JSON format by specifying the --json flag")
		}

		response, err := a.AccessControl.UpdateRuleSet(ctx, updateRuleSetReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range updateRuleSetOverrides {
		fn(cmd, &updateRuleSetReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newUpdateRuleSet())
	})
}

// end service AccountAccessControl
