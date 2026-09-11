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

var createNamespaceGrackleV1betaCmd = &cobra.Command{
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

var listNamespacesGrackleV1betaCmd = &cobra.Command{
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

var getNamespaceGrackleV1betaCmd = &cobra.Command{
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

var deleteNamespaceGrackleV1betaCmd = &cobra.Command{
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

var updateNamespaceGrackleV1betaCmd = &cobra.Command{
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

var createSemaphoreGrackleV1betaCmd = &cobra.Command{
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

var listSemaphoresGrackleV1betaCmd = &cobra.Command{
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

var getSemaphoreGrackleV1betaCmd = &cobra.Command{
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

var acquireSemaphoreGrackleV1betaCmd = &cobra.Command{
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

var releaseSemaphoreGrackleV1betaCmd = &cobra.Command{
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

var updateSemaphoreGrackleV1betaCmd = &cobra.Command{
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

var deleteSemaphoreGrackleV1betaCmd = &cobra.Command{
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

var createWaitGroupGrackleV1betaCmd = &cobra.Command{
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

var listWaitGroupsGrackleV1betaCmd = &cobra.Command{
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

var getWaitGroupGrackleV1betaCmd = &cobra.Command{
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

var deleteWaitGroupGrackleV1betaCmd = &cobra.Command{
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

var updateWaitGroupGrackleV1betaCmd = &cobra.Command{
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

var waitForWaitGroupGrackleV1betaCmd = &cobra.Command{
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

var completeJobsFromWaitGroupGrackleV1betaCmd = &cobra.Command{
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

var acquireLockGrackleV1betaCmd = &cobra.Command{
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

var releaseLockGrackleV1betaCmd = &cobra.Command{
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

var getLockGrackleV1betaCmd = &cobra.Command{
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

var deleteLockGrackleV1betaCmd = &cobra.Command{
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

var listLocksGrackleV1betaCmd = &cobra.Command{
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

var listSemaphoreHoldersGrackleV1betaCmd = &cobra.Command{
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

var createSemaphoreLeaseGrackleV1betaCmd = &cobra.Command{
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

var revokeSemaphoreLeaseGrackleV1betaCmd = &cobra.Command{
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

var refreshSemaphoreLeaseGrackleV1betaCmd = &cobra.Command{
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

var listSemaphoreLeasesGrackleV1betaCmd = &cobra.Command{
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

var getSemaphoreLeaseGrackleV1betaCmd = &cobra.Command{
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

var listWaitGroupCompletedJobsGrackleV1betaCmd = &cobra.Command{
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

var createLockLeaseGrackleV1betaCmd = &cobra.Command{
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

var revokeLockLeaseGrackleV1betaCmd = &cobra.Command{
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

var refreshLockLeaseGrackleV1betaCmd = &cobra.Command{
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

var listLockLeasesGrackleV1betaCmd = &cobra.Command{
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

var getLockLeaseGrackleV1betaCmd = &cobra.Command{
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

var createBarrierGrackleV1betaCmd = &cobra.Command{
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

var listBarriersGrackleV1betaCmd = &cobra.Command{
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

var getBarrierGrackleV1betaCmd = &cobra.Command{
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

var deleteBarrierGrackleV1betaCmd = &cobra.Command{
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

var updateBarrierGrackleV1betaCmd = &cobra.Command{
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

var arriveAtBarrierGrackleV1betaCmd = &cobra.Command{
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

var waitAtBarrierGrackleV1betaCmd = &cobra.Command{
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

var listBarrierParticipantsGrackleV1betaCmd = &cobra.Command{
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

	grackleV1betaCmd.AddCommand(listSemaphoreHoldersGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(createSemaphoreLeaseGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(revokeSemaphoreLeaseGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(refreshSemaphoreLeaseGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(listSemaphoreLeasesGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(getSemaphoreLeaseGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(listWaitGroupCompletedJobsGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(createLockLeaseGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(revokeLockLeaseGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(refreshLockLeaseGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(listLockLeasesGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(getLockLeaseGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(createBarrierGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(listBarriersGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(getBarrierGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(deleteBarrierGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(updateBarrierGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(arriveAtBarrierGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(waitAtBarrierGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(listBarrierParticipantsGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(createNamespaceGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(listNamespacesGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(getNamespaceGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(deleteNamespaceGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(updateNamespaceGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(createSemaphoreGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(listSemaphoresGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(getSemaphoreGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(acquireSemaphoreGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(releaseSemaphoreGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(updateSemaphoreGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(deleteSemaphoreGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(createWaitGroupGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(listWaitGroupsGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(getWaitGroupGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(deleteWaitGroupGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(updateWaitGroupGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(waitForWaitGroupGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(completeJobsFromWaitGroupGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(acquireLockGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(releaseLockGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(getLockGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(deleteLockGrackleV1betaCmd)
	grackleV1betaCmd.AddCommand(listLocksGrackleV1betaCmd)

	grackleV1betaCmd.PersistentFlags().StringVarP(&grackleV1betaCmdCfg.apiKeyId, "api-key-id", "", "", "API key ID (key_alfa_* or key_bravo_*)")
	grackleV1betaCmd.PersistentFlags().StringVarP(&grackleV1betaCmdCfg.apiSecretKey, "api-secret-key", "", "", "API secret key")

	grackleV1betaCmd.PersistentFlags().StringVarP(&grackleV1betaCmdCfg.endpoint, "endpoint", "", "", "Grackle API address")
	err := grackleV1betaCmd.MarkPersistentFlagRequired("endpoint")
	if err != nil {
		log.Fatal(err)
	}
}

func getGrackleV1betaClient() grackle_v1beta.GrackleApi {
	client, err := grackle_v1beta.NewGrackleGrpcClient(grackleV1betaCmdCfg.endpoint, getSigner(grackleV1betaCmdCfg.apiKeyId, grackleV1betaCmdCfg.apiSecretKey))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
