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

var waitForWaitGroupPreviewCmd = &cobra.Command{
	Use:   "wait-for-wait-group",
	Short: "WaitForWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.WaitForWaitGroupRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.WaitForWaitGroup(context.Background(), req)
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

var listSemaphoreHoldersCmd = &cobra.Command{
	Use:   "list-semaphore-holders",
	Short: "ListSemaphoreHolders",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ListSemaphoreHoldersRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ListSemaphoreHolders(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var createSemaphoreLeaseCmd = &cobra.Command{
	Use:   "create-semaphore-lease",
	Short: "CreateSemaphoreLease",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.CreateSemaphoreLeaseRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.CreateSemaphoreLease(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var revokeSemaphoreLeaseCmd = &cobra.Command{
	Use:   "revoke-semaphore-lease",
	Short: "RevokeSemaphoreLease",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.RevokeSemaphoreLeaseRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.RevokeSemaphoreLease(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var refreshSemaphoreLeaseCmd = &cobra.Command{
	Use:   "refresh-semaphore-lease",
	Short: "RefreshSemaphoreLease",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.RefreshSemaphoreLeaseRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.RefreshSemaphoreLease(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listSemaphoreLeasesCmd = &cobra.Command{
	Use:   "list-semaphore-leases",
	Short: "ListSemaphoreLeases",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ListSemaphoreLeasesRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ListSemaphoreLeases(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getSemaphoreLeaseCmd = &cobra.Command{
	Use:   "get-semaphore-lease",
	Short: "GetSemaphoreLease",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.GetSemaphoreLeaseRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.GetSemaphoreLease(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listWaitGroupJobsCmd = &cobra.Command{
	Use:   "list-wait-group-jobs",
	Short: "ListWaitGroupJobs",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ListWaitGroupJobsRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ListWaitGroupJobs(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var createLockLeaseCmd = &cobra.Command{
	Use:   "create-lock-lease",
	Short: "CreateLockLease",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.CreateLockLeaseRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.CreateLockLease(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var revokeLockLeaseCmd = &cobra.Command{
	Use:   "revoke-lock-lease",
	Short: "RevokeLockLease",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.RevokeLockLeaseRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.RevokeLockLease(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var refreshLockLeaseCmd = &cobra.Command{
	Use:   "refresh-lock-lease",
	Short: "RefreshLockLease",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.RefreshLockLeaseRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.RefreshLockLease(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listLockLeasesCmd = &cobra.Command{
	Use:   "list-lock-leases",
	Short: "ListLockLeases",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ListLockLeasesRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ListLockLeases(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getLockLeaseCmd = &cobra.Command{
	Use:   "get-lock-lease",
	Short: "GetLockLease",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.GetLockLeaseRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.GetLockLease(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var createBarrierCmd = &cobra.Command{
	Use:   "create-barrier",
	Short: "CreateBarrier",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.CreateBarrierRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.CreateBarrier(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listBarriersCmd = &cobra.Command{
	Use:   "list-barriers",
	Short: "ListBarriers",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ListBarriersRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ListBarriers(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getBarrierCmd = &cobra.Command{
	Use:   "get-barrier",
	Short: "GetBarrier",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.GetBarrierRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.GetBarrier(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteBarrierCmd = &cobra.Command{
	Use:   "delete-barrier",
	Short: "DeleteBarrier",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.DeleteBarrierRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.DeleteBarrier(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var updateBarrierCmd = &cobra.Command{
	Use:   "update-barrier",
	Short: "UpdateBarrier",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.UpdateBarrierRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.UpdateBarrier(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var arriveAtBarrierCmd = &cobra.Command{
	Use:   "arrive-at-barrier",
	Short: "ArriveAtBarrier",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ArriveAtBarrierRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ArriveAtBarrier(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var waitAtBarrierCmd = &cobra.Command{
	Use:   "wait-at-barrier",
	Short: "WaitAtBarrier",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.WaitAtBarrierRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.WaitAtBarrier(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listBarrierParticipantsCmd = &cobra.Command{
	Use:   "list-barrier-participants",
	Short: "ListBarrierParticipants",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_preview.ListBarrierParticipantsRequest{}
		readRequest(req)

		client := getGracklePreviewClient()

		resp, err := client.ListBarrierParticipants(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

func init() {
	rootCmd.AddCommand(gracklePreviewCmd)

	gracklePreviewCmd.AddCommand(listSemaphoreHoldersCmd)
	gracklePreviewCmd.AddCommand(createSemaphoreLeaseCmd)
	gracklePreviewCmd.AddCommand(revokeSemaphoreLeaseCmd)
	gracklePreviewCmd.AddCommand(refreshSemaphoreLeaseCmd)
	gracklePreviewCmd.AddCommand(listSemaphoreLeasesCmd)
	gracklePreviewCmd.AddCommand(getSemaphoreLeaseCmd)
	gracklePreviewCmd.AddCommand(listWaitGroupJobsCmd)
	gracklePreviewCmd.AddCommand(createLockLeaseCmd)
	gracklePreviewCmd.AddCommand(revokeLockLeaseCmd)
	gracklePreviewCmd.AddCommand(refreshLockLeaseCmd)
	gracklePreviewCmd.AddCommand(listLockLeasesCmd)
	gracklePreviewCmd.AddCommand(getLockLeaseCmd)
	gracklePreviewCmd.AddCommand(createBarrierCmd)
	gracklePreviewCmd.AddCommand(listBarriersCmd)
	gracklePreviewCmd.AddCommand(getBarrierCmd)
	gracklePreviewCmd.AddCommand(deleteBarrierCmd)
	gracklePreviewCmd.AddCommand(updateBarrierCmd)
	gracklePreviewCmd.AddCommand(arriveAtBarrierCmd)
	gracklePreviewCmd.AddCommand(waitAtBarrierCmd)
	gracklePreviewCmd.AddCommand(listBarrierParticipantsCmd)
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
	gracklePreviewCmd.AddCommand(waitForWaitGroupPreviewCmd)
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
