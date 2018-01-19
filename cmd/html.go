package cmd

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/Feltix/feltixhtml"
	"github.com/Feltix/feltixparser"
	"github.com/spf13/cobra"
)

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	Use:              "html [path to input file]",
	Short:            "Export file as an HTML webpage",
	Args:             cobra.ExactArgs(1),
	TraverseChildren: true,
	Long:             longDescription,
	Run:              htmlRun,
}

var (
	htmlEmbedableFlag      bool
	htmlNoscenenumbersFlag bool
)

func init() {
	htmlCmd.Flags().BoolVarP(&htmlEmbedableFlag, "embedable", "e", false, "only output the play itself")
	htmlCmd.Flags().BoolVarP(&htmlNoscenenumbersFlag, "noscenenumbers", "s", false, "remove scenenumbers from output")

	FeltixCmd.AddCommand(htmlCmd)
}

func htmlRun(cmd *cobra.Command, args []string) {
	pathToFile := args[0]

	if isFeltixFile(pathToFile) {
		feltixparser.UseFeltixExtensions = true
	}

	startTime := time.Now()
	script, err := feltixparser.ParseFile(pathToFile)
	handle(err)

	// Get the filepath to use during export.
	if outFlag != "" {
		pathToFile = outFlag
	} else {
		extension := filepath.Ext(pathToFile)
		pathToFile = strings.TrimSuffix(pathToFile, extension) + ".html"
	}

	startExportTime := time.Now()

	if htmlNoscenenumbersFlag {
		feltixhtml.AddSceneNumbers = false
	}

	if htmlEmbedableFlag {
		html := feltixhtml.ToHTML(script)
		err = ioutil.WriteFile(pathToFile, []byte(html), 0664)

	} else {
		err = feltixhtml.WriteHTMLPage(script, pathToFile)
	}

	handle(err)

	endTime := time.Now()

	if benchmarkFlag {
		printBenchmarks(startTime, startExportTime, endTime)
	}
}