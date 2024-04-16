/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cipher_cli/pkg/bacon"
	"cipher_cli/pkg/caesar"
	"cipher_cli/pkg/validation"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const CAESAR = "caesar"

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt [string to decode] --algorithm=[algorithm]",
	Short: "Decode a cryptic message",
	Long: `for example: 
	cipher_cli decrypt AABAAABBABABAABABBBABBAAA --algorithm=bacon
	cipher_cli decrypt fff.jkl.gh --algorithm=caesar --key=87
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ciphertext := strings.Join(args, " ")
		var plaintext = ""
		algorithm := cmd.Flags().Lookup("algorithm").Value.String()
		key := cmd.Flags().Lookup("key").Value.String()
		if strings.ToLower(algorithm) == CAESAR {
			rotations, err := validation.GetRotationsFromKey(key)
			if err != nil {
				fmt.Printf("Error: %s", err)
				os.Exit(1)
			}
			plaintext = caesar.Decrypt(ciphertext, rotations)
		} else {
			plaintext = bacon.Decrypt(ciphertext)
		}
		fmt.Printf("CipherText: %s\nPlaintext: %s\n", ciphertext, plaintext)
	},
}

func init() {
	// declare local flag
	decryptCmd.Flags().StringP("algorithm", "a", "", "algorithm to use for this action")
	_ = decryptCmd.MarkFlagRequired("algorithm")
	rootCmd.AddCommand(decryptCmd)
}
