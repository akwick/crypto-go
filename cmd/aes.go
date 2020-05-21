// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/akwick/crypto-go/aes"
	"github.com/spf13/cobra"
)

// aesCmd represents the aes command
var aesCmd = &cobra.Command{
	Use:   "aes",
	Short: "A usage example for encryption",
	Long: `This command will show you an example for an AES-256-GCM encryption.
	Each time you call this command a new random key will be generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("### Usage Example AES-256-GCM encryption ###")
		key := aes.GetKey()
		msg := []byte("Hello go get -u community!")
		fmt.Printf("Encrypt message: %s\n", msg)

		cipher, err := aes.Encrypt(msg, &key)
		if err != nil {
			fmt.Printf("encryption failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Encrypted message: %s\n", cipher)

		plaintext, err := aes.Decrypt(cipher, &key)
		if err != nil {
			fmt.Printf("decryption failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Decrypted message: %s\n", plaintext)
	},
}

func init() {
	rootCmd.AddCommand(aesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
