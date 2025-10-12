package commands

import (
	"context"
	"log"

	grackle_preview "github.com/evrblk/evrblk-go/grackle/preview"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var gracklePreviewCmdCfg struct {
	apiKeyId     string
	apiSecretKey string
	endpoint     string
}

// gracklePreviewCmd represents the base command for calling Grackle Preview APIs
var gracklePreviewCmd = &cobra.Command{
	Use:   "grackle-preview",
	Short: "Call Grackle Preview API methods",
	Long:  "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetEnvPrefix("EVRBLK")
		viper.AutomaticEnv()

		if viper.IsSet("api_key_id") {
			gracklePreviewCmdCfg.apiKeyId = viper.GetString("api_key_id")
		}
		if viper.IsSet("api_secret_key") {
			gracklePreviewCmdCfg.apiSecretKey = viper.GetString("api_secret_key")
		}

		return nil
	},
}

var createNamespacePreviewCmd = &cobra.Command{
	Use:   "create-namespace",
	Short: "CreateNamespace",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.CreateNamespaceRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.CreateNamespace(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listNamespacesPreviewCmd = &cobra.Command{
	Use:   "list-namespaces",
	Short: "ListNamespaces",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ListNamespacesRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ListNamespaces(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getNamespacePreviewCmd = &cobra.Command{
	Use:   "get-namespace",
	Short: "GetNamespace",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.GetNamespaceRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.GetNamespace(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteNamespacePreviewCmd = &cobra.Command{
	Use:   "delete-namespace",
	Short: "DeleteNamespace",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.DeleteNamespaceRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.DeleteNamespace(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var updateNamespacePreviewCmd = &cobra.Command{
	Use:   "update-namespace",
	Short: "UpdateNamespace",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.UpdateNamespaceRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.UpdateNamespace(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var createSemaphorePreviewCmd = &cobra.Command{
	Use:   "create-semaphore",
	Short: "CreateSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.CreateSemaphoreRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.CreateSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listSemaphoresPreviewCmd = &cobra.Command{
	Use:   "list-semaphores",
	Short: "ListSemaphores",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ListSemaphoresRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ListSemaphores(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getSemaphorePreviewCmd = &cobra.Command{
	Use:   "get-semaphore",
	Short: "GetSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.GetSemaphoreRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.GetSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var acquireSemaphorePreviewCmd = &cobra.Command{
	Use:   "acquire-semaphore",
	Short: "AcquireSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.AcquireSemaphoreRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.AcquireSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var releaseSemaphorePreviewCmd = &cobra.Command{
	Use:   "release-semaphore",
	Short: "ReleaseSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ReleaseSemaphoreRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ReleaseSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var updateSemaphorePreviewCmd = &cobra.Command{
	Use:   "update-semaphore",
	Short: "UpdateSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.UpdateSemaphoreRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.UpdateSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteSemaphorePreviewCmd = &cobra.Command{
	Use:   "delete-semaphore",
	Short: "DeleteSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.DeleteSemaphoreRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.DeleteSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var createWaitGroupPreviewCmd = &cobra.Command{
	Use:   "create-wait-group",
	Short: "CreateWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.CreateWaitGroupRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.CreateWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listWaitGroupsPreviewCmd = &cobra.Command{
	Use:   "list-wait-groups",
	Short: "ListWaitGroups",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ListWaitGroupsRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ListWaitGroups(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getWaitGroupPreviewCmd = &cobra.Command{
	Use:   "get-wait-group",
	Short: "GetWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.GetWaitGroupRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.GetWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteWaitGroupPreviewCmd = &cobra.Command{
	Use:   "delete-wait-group",
	Short: "DeleteWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.DeleteWaitGroupRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.DeleteWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var addJobsToWaitGroupPreviewCmd = &cobra.Command{
	Use:   "add-jobs-to-wait-group",
	Short: "AddJobsToWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.AddJobsToWaitGroupRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.AddJobsToWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var completeJobsFromWaitGroupPreviewCmd = &cobra.Command{
	Use:   "complete-jobs-from-wait-group",
	Short: "CompleteJobsFromWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.CompleteJobsFromWaitGroupRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.CompleteJobsFromWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var acquireLockPreviewCmd = &cobra.Command{
	Use:   "acquire-lock",
	Short: "AcquireLock",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.AcquireLockRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.AcquireLock(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var releaseLockPreviewCmd = &cobra.Command{
	Use:   "release-lock",
	Short: "ReleaseLock",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ReleaseLockRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ReleaseLock(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getLockPreviewCmd = &cobra.Command{
	Use:   "get-lock",
	Short: "GetLock",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.GetLockRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.GetLock(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteLockPreviewCmd = &cobra.Command{
	Use:   "delete-lock",
	Short: "DeleteLock",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.DeleteLockRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.DeleteLock(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listLocksPreviewCmd = &cobra.Command{
	Use:   "list-locks",
	Short: "ListLocks",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ListLocksRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ListLocks(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

func init() {
	rootCmd.AddCommand(gracklePreviewCmd)

	gracklePreviewCmd.AddCommand(createNamespacePreviewCmd)
	gracklePreviewCmd.AddCommand(listNamespacesPreviewCmd)
	gracklePreviewCmd.AddCommand(getNamespacePreviewCmd)
	gracklePreviewCmd.AddCommand(deleteNamespacePreviewCmd)
	gracklePreviewCmd.AddCommand(updateNamespacePreviewCmd)
	gracklePreviewCmd.AddCommand(createSemaphorePreviewCmd)
	gracklePreviewCmd.AddCommand(listSemaphoresPreviewCmd)
	gracklePreviewCmd.AddCommand(getSemaphorePreviewCmd)
	gracklePreviewCmd.AddCommand(acquireSemaphorePreviewCmd)
	gracklePreviewCmd.AddCommand(releaseSemaphorePreviewCmd)
	gracklePreviewCmd.AddCommand(updateSemaphorePreviewCmd)
	gracklePreviewCmd.AddCommand(deleteSemaphorePreviewCmd)
	gracklePreviewCmd.AddCommand(createWaitGroupPreviewCmd)
	gracklePreviewCmd.AddCommand(listWaitGroupsPreviewCmd)
	gracklePreviewCmd.AddCommand(getWaitGroupPreviewCmd)
	gracklePreviewCmd.AddCommand(deleteWaitGroupPreviewCmd)
	gracklePreviewCmd.AddCommand(addJobsToWaitGroupPreviewCmd)
	gracklePreviewCmd.AddCommand(completeJobsFromWaitGroupPreviewCmd)
	gracklePreviewCmd.AddCommand(acquireLockPreviewCmd)
	gracklePreviewCmd.AddCommand(releaseLockPreviewCmd)
	gracklePreviewCmd.AddCommand(getLockPreviewCmd)
	gracklePreviewCmd.AddCommand(deleteLockPreviewCmd)
	gracklePreviewCmd.AddCommand(listLocksPreviewCmd)

	gracklePreviewCmd.PersistentFlags().StringVarP(&gracklePreviewCmdCfg.apiKeyId, "api-key-id", "", "", "API key ID (key_alfa_* or key_bravo_*)")
	gracklePreviewCmd.PersistentFlags().StringVarP(&gracklePreviewCmdCfg.apiSecretKey, "api-secret-key", "", "", "API secret key")

	gracklePreviewCmd.PersistentFlags().StringVarP(&gracklePreviewCmdCfg.endpoint, "endpoint", "", "", "Grackle API address")
	err := gracklePreviewCmd.MarkPersistentFlagRequired("endpoint")
	if err != nil {
		log.Fatal(err)
	}
}

func getGracklePreviewClient() grackle_preview.GrackleApi {
	return grackle_preview.NewGrackleGrpcClient(gracklePreviewCmdCfg.endpoint, getSigner(gracklePreviewCmdCfg.apiKeyId, gracklePreviewCmdCfg.apiSecretKey))
}
