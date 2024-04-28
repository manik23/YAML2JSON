/*
Copyright © 2024 manik.x.mahajan@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type info struct {
	filePath string
	outPath  string
}

func parseYaml(file, outfile string) string {

	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return err.Error()
	}

	if info.IsDir() {
		return fmt.Sprintf("%s : %s", file, "is a directory")
	}

	b, e := os.ReadFile(file)
	if e != nil {
		return e.Error()
	}

	output := make(map[string]any)

	_ = yaml.Unmarshal(b, &output)

	content, err := json.MarshalIndent(output, "", " ")

	if err != nil {
		return "err : unable to decode to json"
	}

	_ = os.WriteFile(outfile, content, 0644)

	return string(content)
}

func parse() *cobra.Command {
	i := info{}
	cmd := &cobra.Command{
		Use:   "parse",
		Short: "YAML to JSON",
		Long:  "parse the YAML file and convert it to JSON",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(parseYaml(i.filePath, i.outPath))
		},
	}

	cmd.Flags().StringVar(&i.filePath, "path", "", "use to locate the yaml file")
	cmd.Flags().StringVar(&i.outPath, "o", "value.json", "name of desired json file")
	cmd.MarkFlagRequired("path")
	return cmd
}
