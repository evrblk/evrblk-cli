package commands

import (
	"context"
	"log"

	grackle_v1beta "github.com/evrblk/evrblk-go/grackle/v1beta"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var grackleV1betaCmdCfg struct {
	apiKeyId     string
	apiSecretKey string
	endpoint     string
}

// grackleV1betaCmd represents the base command for calling Grackle V1 Beta APIs
var grackleV1betaCmd = &cobra.Command{
	Use:   "grackle-v1beta",
	Short: "Call Grackle V1 Beta API methods",
	Long:  "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetEnvPrefix("EVRBLK")
		viper.AutomaticEnv()

		if viper.IsSet("api_key_id") {
			grackleV1betaCmdCfg.apiKeyId = viper.GetString("api_key_id")
		}
		if viper.IsSet("api_secret_key") {
			grackleV1betaCmdCfg.apiSecretKey = viper.GetString("api_secret_key")
		}

		return nil
	},
}

var createNamespaceV1betaCmd = &cobra.Command{
	Use:   "create-namespace",
	Short: "CreateNamespace",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.CreateNamespaceRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.CreateNamespace(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listNamespacesV1betaCmd = &cobra.Command{
	Use:   "list-namespaces",
	Short: "ListNamespaces",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.ListNamespacesRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.ListNamespaces(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getNamespaceV1betaCmd = &cobra.Command{
	Use:   "get-namespace",
	Short: "GetNamespace",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.GetNamespaceRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.GetNamespace(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteNamespaceV1betaCmd = &cobra.Command{
	Use:   "delete-namespace",
	Short: "DeleteNamespace",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.DeleteNamespaceRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.DeleteNamespace(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var updateNamespaceV1betaCmd = &cobra.Command{
	Use:   "update-namespace",
	Short: "UpdateNamespace",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.UpdateNamespaceRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.UpdateNamespace(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var createSemaphoreV1betaCmd = &cobra.Command{
	Use:   "create-semaphore",
	Short: "CreateSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.CreateSemaphoreRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.CreateSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listSemaphoresV1betaCmd = &cobra.Command{
	Use:   "list-semaphores",
	Short: "ListSemaphores",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.ListSemaphoresRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.ListSemaphores(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getSemaphoreV1betaCmd = &cobra.Command{
	Use:   "get-semaphore",
	Short: "GetSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.GetSemaphoreRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.GetSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var acquireSemaphoreV1betaCmd = &cobra.Command{
	Use:   "acquire-semaphore",
	Short: "AcquireSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.AcquireSemaphoreRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.AcquireSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var releaseSemaphoreV1betaCmd = &cobra.Command{
	Use:   "release-semaphore",
	Short: "ReleaseSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.ReleaseSemaphoreRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.ReleaseSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var updateSemaphoreV1betaCmd = &cobra.Command{
	Use:   "update-semaphore",
	Short: "UpdateSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.UpdateSemaphoreRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.UpdateSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteSemaphoreV1betaCmd = &cobra.Command{
	Use:   "delete-semaphore",
	Short: "DeleteSemaphore",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.DeleteSemaphoreRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.DeleteSemaphore(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var createWaitGroupV1betaCmd = &cobra.Command{
	Use:   "create-wait-group",
	Short: "CreateWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.CreateWaitGroupRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.CreateWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listWaitGroupsV1betaCmd = &cobra.Command{
	Use:   "list-wait-groups",
	Short: "ListWaitGroups",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.ListWaitGroupsRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.ListWaitGroups(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getWaitGroupV1betaCmd = &cobra.Command{
	Use:   "get-wait-group",
	Short: "GetWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.GetWaitGroupRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.GetWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteWaitGroupV1betaCmd = &cobra.Command{
	Use:   "delete-wait-group",
	Short: "DeleteWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.DeleteWaitGroupRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.DeleteWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var updateWaitGroupV1betaCmd = &cobra.Command{
	Use:   "update-wait-group",
	Short: "UpdateWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.UpdateWaitGroupRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.UpdateWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var waitForWaitGroupV1betaCmd = &cobra.Command{
	Use:   "wait-for-wait-group",
	Short: "WaitForWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.WaitForWaitGroupRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.WaitForWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var completeJobsFromWaitGroupV1betaCmd = &cobra.Command{
	Use:   "complete-jobs-from-wait-group",
	Short: "CompleteJobsFromWaitGroup",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.CompleteJobsFromWaitGroupRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.CompleteJobsFromWaitGroup(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var acquireLockV1betaCmd = &cobra.Command{
	Use:   "acquire-lock",
	Short: "AcquireLock",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.AcquireLockRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.AcquireLock(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var releaseLockV1betaCmd = &cobra.Command{
	Use:   "release-lock",
	Short: "ReleaseLock",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.ReleaseLockRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.ReleaseLock(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getLockV1betaCmd = &cobra.Command{
	Use:   "get-lock",
	Short: "GetLock",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.GetLockRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.GetLock(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteLockV1betaCmd = &cobra.Command{
	Use:   "delete-lock",
	Short: "DeleteLock",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.DeleteLockRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.DeleteLock(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listLocksV1betaCmd = &cobra.Command{
	Use:   "list-locks",
	Short: "ListLocks",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.ListLocksRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.ListSemaphoreHoldersRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.CreateSemaphoreLeaseRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.RevokeSemaphoreLeaseRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.RefreshSemaphoreLeaseRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.ListSemaphoreLeasesRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.GetSemaphoreLeaseRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.GetSemaphoreLease(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listWaitGroupCompletedJobsCmd = &cobra.Command{
	Use:   "list-wait-group-completed-jobs",
	Short: "ListWaitGroupCompletedJobs",
	Run: func(cmd *cobra.Command, args []string) {
		req := &grackle_v1beta.ListWaitGroupCompletedJobsRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.ListWaitGroupCompletedJobs(context.Background(), req)
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
		req := &grackle_v1beta.CreateLockLeaseRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.RevokeLockLeaseRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.RefreshLockLeaseRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.ListLockLeasesRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.GetLockLeaseRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.CreateBarrierRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.ListBarriersRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.GetBarrierRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.DeleteBarrierRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.UpdateBarrierRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.ArriveAtBarrierRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.WaitAtBarrierRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

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
		req := &grackle_v1beta.ListBarrierParticipantsRequest{}
		readRequest(req)

		client := getGrackleV1betaClient()

		resp, err := client.ListBarrierParticipants(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

func init() {
	rootCmd.AddCommand(grackleV1betaCmd)

	grackleV1betaCmd.AddCommand(listSemaphoreHoldersCmd)
	grackleV1betaCmd.AddCommand(createSemaphoreLeaseCmd)
	grackleV1betaCmd.AddCommand(revokeSemaphoreLeaseCmd)
	grackleV1betaCmd.AddCommand(refreshSemaphoreLeaseCmd)
	grackleV1betaCmd.AddCommand(listSemaphoreLeasesCmd)
	grackleV1betaCmd.AddCommand(getSemaphoreLeaseCmd)
	grackleV1betaCmd.AddCommand(listWaitGroupCompletedJobsCmd)
	grackleV1betaCmd.AddCommand(createLockLeaseCmd)
	grackleV1betaCmd.AddCommand(revokeLockLeaseCmd)
	grackleV1betaCmd.AddCommand(refreshLockLeaseCmd)
	grackleV1betaCmd.AddCommand(listLockLeasesCmd)
	grackleV1betaCmd.AddCommand(getLockLeaseCmd)
	grackleV1betaCmd.AddCommand(createBarrierCmd)
	grackleV1betaCmd.AddCommand(listBarriersCmd)
	grackleV1betaCmd.AddCommand(getBarrierCmd)
	grackleV1betaCmd.AddCommand(deleteBarrierCmd)
	grackleV1betaCmd.AddCommand(updateBarrierCmd)
	grackleV1betaCmd.AddCommand(arriveAtBarrierCmd)
	grackleV1betaCmd.AddCommand(waitAtBarrierCmd)
	grackleV1betaCmd.AddCommand(listBarrierParticipantsCmd)
	grackleV1betaCmd.AddCommand(createNamespaceV1betaCmd)
	grackleV1betaCmd.AddCommand(listNamespacesV1betaCmd)
	grackleV1betaCmd.AddCommand(getNamespaceV1betaCmd)
	grackleV1betaCmd.AddCommand(deleteNamespaceV1betaCmd)
	grackleV1betaCmd.AddCommand(updateNamespaceV1betaCmd)
	grackleV1betaCmd.AddCommand(createSemaphoreV1betaCmd)
	grackleV1betaCmd.AddCommand(listSemaphoresV1betaCmd)
	grackleV1betaCmd.AddCommand(getSemaphoreV1betaCmd)
	grackleV1betaCmd.AddCommand(acquireSemaphoreV1betaCmd)
	grackleV1betaCmd.AddCommand(releaseSemaphoreV1betaCmd)
	grackleV1betaCmd.AddCommand(updateSemaphoreV1betaCmd)
	grackleV1betaCmd.AddCommand(deleteSemaphoreV1betaCmd)
	grackleV1betaCmd.AddCommand(createWaitGroupV1betaCmd)
	grackleV1betaCmd.AddCommand(listWaitGroupsV1betaCmd)
	grackleV1betaCmd.AddCommand(getWaitGroupV1betaCmd)
	grackleV1betaCmd.AddCommand(deleteWaitGroupV1betaCmd)
	grackleV1betaCmd.AddCommand(updateWaitGroupV1betaCmd)
	grackleV1betaCmd.AddCommand(waitForWaitGroupV1betaCmd)
	grackleV1betaCmd.AddCommand(completeJobsFromWaitGroupV1betaCmd)
	grackleV1betaCmd.AddCommand(acquireLockV1betaCmd)
	grackleV1betaCmd.AddCommand(releaseLockV1betaCmd)
	grackleV1betaCmd.AddCommand(getLockV1betaCmd)
	grackleV1betaCmd.AddCommand(deleteLockV1betaCmd)
	grackleV1betaCmd.AddCommand(listLocksV1betaCmd)

	grackleV1betaCmd.PersistentFlags().StringVarP(&grackleV1betaCmdCfg.apiKeyId, "api-key-id", "", "", "API key ID (key_alfa_* or key_bravo_*)")
	grackleV1betaCmd.PersistentFlags().StringVarP(&grackleV1betaCmdCfg.apiSecretKey, "api-secret-key", "", "", "API secret key")

	grackleV1betaCmd.PersistentFlags().StringVarP(&grackleV1betaCmdCfg.endpoint, "endpoint", "", "", "Grackle API address")
	err := grackleV1betaCmd.MarkPersistentFlagRequired("endpoint")
	if err != nil {
		log.Fatal(err)
	}
}

func getGrackleV1betaClient() grackle_v1beta.GrackleApi {
	return grackle_v1beta.NewGrackleGrpcClient(grackleV1betaCmdCfg.endpoint, getSigner(grackleV1betaCmdCfg.apiKeyId, grackleV1betaCmdCfg.apiSecretKey))
}
