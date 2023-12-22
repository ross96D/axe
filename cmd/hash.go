package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/ross96D/axe/hash"
	"github.com/spf13/cobra"
)

var inputFile string

var hashCommand = &cobra.Command{
	Use:   "hash",
	Short: "Calculate the crc32 hash of the input file",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(inputFile)
		if err != nil {
			println("error opening the file", fmt.Sprintf("%v", err))
			os.Exit(1)
		}

		hash := hash.Hash(file)

		os.Stdout.WriteString(base64.StdEncoding.EncodeToString(hash))
	},
}

func init() {
	hashCommand.Flags().StringVarP(&inputFile, "file", "f", "", "file to calculate the hash function")
}
