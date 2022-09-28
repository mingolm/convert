package main

import (
	"bufio"
	"context"
	"github.com/mingolm/convert/convert"
	"github.com/spf13/cobra"
	"os"
)

/**
 * convert -f test.xlsx -t csv
 * convert -f test.xlsx -t csv -o test.csv
 * convert -f test.xlsx -t xml
 * convert -f test.xlsx -t json
 * convert -m test1.csv test2.csv test3.csv -o test.csv
 */
func main() {
	var (
		sourceFile string // 输入的源文件
		targetType string // 输出类型,csv/pdf/xml...
		output     string // 输出地址，目录+文件名
		ctx        = context.Background()
	)

	rootCmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert document formats",
		Long:  `A tool for converting document formats`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				sourceExt convert.Ext
				targetExt convert.Ext
			)

			if err := sourceExt.FormString(GetExt(sourceFile)); err != nil {
				return err
			}
			if targetType == "" {
				targetType = GetExt(output)
			}
			if err := targetExt.FormString(targetType); err != nil {
				return err
			}
			f, err := os.Open(sourceFile)
			if err != nil {
				return err
			}

			if err := convert.New(&convert.Config{
				SourceExt:       sourceExt,
				SourceBufReader: bufio.NewReader(f),
				TargetExt:       targetExt,
				Output:          output,
				PrintProcessing: cmd.Flag("v").Value.String() == "true",
			}).Run(ctx); err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.PersistentFlags().Bool("v", false, "print details")
	rootCmd.Flags().StringVarP(
		&sourceFile, "source_file", "f", "",
		"The source file of this execution",
	)
	rootCmd.Flags().StringVarP(
		&output, "output", "o", "",
		"The output of this execution",
	)
	rootCmd.Flags().StringVarP(
		&targetType, "target_type", "t", "",
		"The target type of this execution",
	)
	_ = rootCmd.MarkFlagRequired("source_file")
	_ = rootCmd.Execute()
}
