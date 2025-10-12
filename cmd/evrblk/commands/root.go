package commands

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/evrblk/evrblk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var rootCmdCfg struct {
	apiKeyId     string
	apiSecretKey string
	endpoint     string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "evrblk",
	Short: "Run EVRBLK CLI",
	Long:  "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetEnvPrefix("EVRBLK")
		viper.AutomaticEnv()

		if viper.IsSet("api_key_id") {
			rootCmdCfg.apiKeyId = viper.GetString("api_key_id")
		}
		if viper.IsSet("api_secret_key") {
			rootCmdCfg.apiSecretKey = viper.GetString("api_secret_key")
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&rootCmdCfg.endpoint, "endpoint", "", "", "Grackle API address")
	err := rootCmd.MarkPersistentFlagRequired("endpoint")
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().StringVarP(&rootCmdCfg.apiKeyId, "api-key-id", "", "", "API key ID (key_alfa_* or key_bravo_*)")
	rootCmd.PersistentFlags().StringVarP(&rootCmdCfg.apiSecretKey, "api-secret-key", "", "", "API secret key")
}

func getSigner() evrblk.RequestSigner {
	if rootCmdCfg.apiKeyId != "" {
		if strings.HasPrefix(rootCmdCfg.apiKeyId, "key_alfa_") {
			signer, err := evrblk.NewAlfaRequestSigner(rootCmdCfg.apiKeyId, rootCmdCfg.apiSecretKey)
			if err != nil {
				log.Fatal(err)
			}
			return signer
		} else if strings.HasPrefix(rootCmdCfg.apiKeyId, "key_bravo_") {
			signer, err := evrblk.NewBravoRequestSigner(rootCmdCfg.apiKeyId, rootCmdCfg.apiSecretKey)
			if err != nil {
				log.Fatal(err)
			}
			return signer
		} else {
			log.Fatalf("Unknown API key type %s", rootCmdCfg.apiKeyId)
		}
	}

	return evrblk.NewNoOpSigner()
}

func printResponse(resp proto.Message) {
	respJson, err := protojson.MarshalOptions{Indent: "  "}.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respJson))
}

func readRequest(req proto.Message) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return // no pipe
	}

	reqJson, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	err = protojson.Unmarshal(reqJson, req)
	if err != nil {
		log.Fatal(err)
	}
}
