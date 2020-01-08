// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"fmt"
	"io"

	"github.com/aws/amazon-ecs-cli-v2/internal/pkg/describe"
	"github.com/aws/amazon-ecs-cli-v2/internal/pkg/store"
	"github.com/aws/amazon-ecs-cli-v2/internal/pkg/term/color"
	"github.com/aws/amazon-ecs-cli-v2/internal/pkg/term/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ShowProjectOpts contains the fields to collect for showing a project.
type ShowProjectOpts struct {
	shouldOutputJSON bool

	storeSvc storeReader

	w io.Writer

	*GlobalOpts
}

// Validate returns an error if the values provided by the user are invalid.
func (o *ShowProjectOpts) Validate() error {
	if o.ProjectName() != "" {
		_, err := o.storeSvc.GetProject(o.ProjectName())
		if err != nil {
			return err
		}
	}

	return nil
}

// Ask asks for fields that are required but not passed in.
func (o *ShowProjectOpts) Ask() error {
	if err := o.askProject(); err != nil {
		return err
	}

	return nil
}

// Execute shows the applications through the prompt.
func (o *ShowProjectOpts) Execute() error {
	projectToSerialize, err := o.retrieveData()
	if err != nil {
		return err
	}
	if o.shouldOutputJSON {
		data, err := projectToSerialize.JSONString()
		if err != nil {
			return err
		}
		fmt.Fprintf(o.w, data)
	} else {
		fmt.Fprintf(o.w, projectToSerialize.HumanString())
	}

	return nil
}

func (o *ShowProjectOpts) retrieveData() (*describe.Project, error) {
	proj, err := o.storeSvc.GetProject(o.ProjectName())
	if err != nil {
		return nil, fmt.Errorf("get project: %w", err)
	}
	envs, err := o.storeSvc.ListEnvironments(o.ProjectName())
	if err != nil {
		return nil, fmt.Errorf("list environment: %w", err)
	}
	apps, err := o.storeSvc.ListApplications(o.ProjectName())
	if err != nil {
		return nil, fmt.Errorf("list application: %w", err)
	}
	var envsToSerialize []*describe.Environment
	for _, env := range envs {
		envsToSerialize = append(envsToSerialize, &describe.Environment{
			Name:      env.Name,
			AccountID: env.AccountID,
			Region:    env.Region,
		})
	}
	var appsToSerialize []*describe.Application
	for _, app := range apps {
		appsToSerialize = append(appsToSerialize, &describe.Application{
			Name: app.Name,
			Type: app.Type,
		})
	}
	return &describe.Project{
		Name: proj.Name,
		URI:  proj.Domain,
		Envs: envsToSerialize,
		Apps: appsToSerialize,
	}, nil
}

func (o *ShowProjectOpts) askProject() error {
	if o.ProjectName() != "" {
		return nil
	}
	projNames, err := o.retrieveProjects()
	if err != nil {
		return err
	}
	if len(projNames) == 0 {
		return fmt.Errorf("no project found: run %s to set one up, or %s into your workspace please", color.HighlightCode("project init"), color.HighlightCode("cd"))
	}
	proj, err := o.prompt.SelectOne(
		applicationShowProjectNamePrompt,
		applicationShowProjectNameHelpPrompt,
		projNames,
	)
	if err != nil {
		return fmt.Errorf("select project: %w", err)
	}
	o.projectName = proj

	return nil
}

func (o *ShowProjectOpts) retrieveProjects() ([]string, error) {
	projs, err := o.storeSvc.ListProjects()
	if err != nil {
		return nil, fmt.Errorf("list project: %w", err)
	}
	projNames := make([]string, len(projs))
	for ind, proj := range projs {
		projNames[ind] = proj.Name
	}
	return projNames, nil
}

// BuildProjectShowCmd builds the command for showing details of a project.
func BuildProjectShowCmd() *cobra.Command {
	opts := ShowProjectOpts{
		w:          log.OutputWriter,
		GlobalOpts: NewGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Shows info about a project.",
		Long:  "Shows configuration, environments and applications for a project.",
		Example: `
  Shows info about the project "my-project"
  /code $ ecs-preview project show -p my-project`,
		PreRunE: runCmdE(func(cmd *cobra.Command, args []string) error {
			ssmStore, err := store.New()
			if err != nil {
				return fmt.Errorf("connect to environment datastore: %w", err)
			}
			opts.storeSvc = ssmStore

			return nil
		}),
		RunE: runCmdE(func(cmd *cobra.Command, args []string) error {
			if err := opts.Validate(); err != nil {
				return err
			}
			if err := opts.Ask(); err != nil {
				return err
			}
			if err := opts.Execute(); err != nil {
				return err
			}

			return nil
		}),
	}
	// The flags bound by viper are available to all sub-commands through viper.GetString({flagName})
	cmd.Flags().BoolVar(&opts.shouldOutputJSON, jsonFlag, false, jsonFlagDescription)
	cmd.Flags().StringP(projectFlag, projectFlagShort, "" /* default */, projectFlagDescription)
	viper.BindPFlag(projectFlag, cmd.Flags().Lookup(projectFlag))
	return cmd
}
