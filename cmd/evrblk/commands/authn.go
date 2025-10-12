package commands

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/evrblk/evrblk-go/authn"
	"github.com/spf13/cobra"
)

// authnCmd represents the base command for working with API keys
var authnCmd = &cobra.Command{
	Use:   "authn",
	Short: "Work with API keys",
	Long:  "",
}

var generateAlfaKeyCmd = &cobra.Command{
	Use:   "generate-alfa-key",
	Short: "Generate Alfa API Key",
	Run: func(cmd *cobra.Command, args []string) {
		privatePem, publicPem, err := authn.GenerateAlfaKeys()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Public Key:\n")
		fmt.Println(publicPem)

		fmt.Println("Private Key:\n")
		fmt.Println(privatePem)

		fmt.Println("Example Key ID (not for cloud use):\n")
		fmt.Printf("key_alfa_%x\n", rand.Uint64())
	},
}

var generateBravoKeyCmd = &cobra.Command{
	Use:   "generate-bravo-key",
	Short: "Generate Bravo API Key",
	Run: func(cmd *cobra.Command, args []string) {
		secret := authn.GenerateBravoSecret()

		fmt.Println("Bravo secret:\n")
		fmt.Println(secret)
		fmt.Println()

		fmt.Println("Example Key ID (not for cloud use):\n")
		fmt.Printf("key_bravo_%x\n", rand.Uint64())
	},
}

func init() {
	rootCmd.AddCommand(authnCmd)

	authnCmd.AddCommand(generateAlfaKeyCmd)
	authnCmd.AddCommand(generateBravoKeyCmd)
}
