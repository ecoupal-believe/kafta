package schema

import (
	"os"

	"github.com/electric-saw/kafta/internal/pkg/configuration"
	"github.com/electric-saw/kafta/internal/pkg/schema"
	cmdutil "github.com/electric-saw/kafta/pkg/cmd/util"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

type subjectVersion struct {
	config      *configuration.Configuration
	subjectName string
}

func NewCmdSubjectVersion(config *configuration.Configuration) *cobra.Command {
	options := &subjectVersion{config: config}
	cmd := &cobra.Command{
		Use:   "subjects-version",
		Short: "subjects-version",
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(options.complete(cmd))
			cmdutil.CheckErr(options.run())
		},
	}
	return cmd
}

func (o *subjectVersion) complete(cmd *cobra.Command) error {
	args := cmd.Flags().Args()
	if len(args) > 1 {
		return cmdutil.HelpErrorf(cmd, "Unexpected args: %v", args)
	}
	if len(args) == 1 {
		o.subjectName = args[0]
	}
	return nil
}

func (o *subjectVersion) run() error {
	versions := schema.NewSubjecVersion(o.config, o.subjectName)

	t := table.NewWriter()
	t.SetStyle(table.StyleDefault)

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Versions"})

	t.AppendRow(table.Row{versions})

	t.AppendSeparator()
	t.Render()

	return nil
}