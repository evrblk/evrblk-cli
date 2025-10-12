package commands

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/evrblk/evrblk-go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func getSigner(apiKeyId string, apiSecretKey string) evrblk.RequestSigner {
	if apiKeyId != "" {
		if strings.HasPrefix(apiKeyId, "key_alfa_") {
			signer, err := evrblk.NewAlfaRequestSigner(apiKeyId, apiSecretKey)
			if err != nil {
				log.Fatal(err)
			}
			return signer
		} else if strings.HasPrefix(apiKeyId, "key_bravo_") {
			signer, err := evrblk.NewBravoRequestSigner(apiKeyId, apiSecretKey)
			if err != nil {
				log.Fatal(err)
			}
			return signer
		} else {
			log.Fatalf("Unknown API key type %s", apiKeyId)
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
