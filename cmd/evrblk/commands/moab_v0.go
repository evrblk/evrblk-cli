package commands

import (
	"context"
	"log"

	moab_v0 "github.com/evrblk/evrblk-go/moab/v0"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var moabV0CmdCfg struct {
	apiKeyId     string
	apiSecretKey string
	endpoint     string
}

// moabV0Cmd represents the base command for calling Moab V0 APIs
var moabV0Cmd = &cobra.Command{
	Use:   "moab-v0",
	Short: "Call Moab V0 API methods",
	Long:  "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetEnvPrefix("EVRBLK")
		viper.AutomaticEnv()

		if viper.IsSet("api_key_id") {
			moabV0CmdCfg.apiKeyId = viper.GetString("api_key_id")
		}
		if viper.IsSet("api_secret_key") {
			moabV0CmdCfg.apiSecretKey = viper.GetString("api_secret_key")
		}

		return nil
	},
}

var createQueueMoabV0Cmd = &cobra.Command{
	Use:   "create-queue",
	Short: "CreateQueue",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.CreateQueueRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.CreateQueue(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getQueueMoabV0Cmd = &cobra.Command{
	Use:   "get-queue",
	Short: "GetQueue",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.GetQueueRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.GetQueue(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var updateQueueMoabV0Cmd = &cobra.Command{
	Use:   "update-queue",
	Short: "UpdateQueue",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.UpdateQueueRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.UpdateQueue(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteQueueMoabV0Cmd = &cobra.Command{
	Use:   "delete-queue",
	Short: "DeleteQueue",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.DeleteQueueRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.DeleteQueue(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listQueuesMoabV0Cmd = &cobra.Command{
	Use:   "list-queues",
	Short: "ListQueues",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.ListQueuesRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.ListQueues(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getTaskMoabV0Cmd = &cobra.Command{
	Use:   "get-task",
	Short: "GetTask",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.GetTaskRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.GetTask(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listTasksMoabV0Cmd = &cobra.Command{
	Use:   "list-tasks",
	Short: "ListTasks",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.ListTasksRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.ListTasks(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var enqueueMoabV0Cmd = &cobra.Command{
	Use:   "enqueue",
	Short: "Enqueue",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.EnqueueRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.Enqueue(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var dequeueMoabV0Cmd = &cobra.Command{
	Use:   "dequeue",
	Short: "Dequeue",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.DequeueRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.Dequeue(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var reportStatusMoabV0Cmd = &cobra.Command{
	Use:   "report-status",
	Short: "ReportStatus",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.ReportStatusRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.ReportStatus(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteTasksMoabV0Cmd = &cobra.Command{
	Use:   "delete-tasks",
	Short: "DeleteTasks",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.DeleteTasksRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.DeleteTasks(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var restartTasksMoabV0Cmd = &cobra.Command{
	Use:   "restart-tasks",
	Short: "RestartTasks",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.RestartTasksRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.RestartTasks(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var purgeQueueMoabV0Cmd = &cobra.Command{
	Use:   "purge-queue",
	Short: "PurgeQueue",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.PurgeQueueRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.PurgeQueue(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var createScheduleMoabV0Cmd = &cobra.Command{
	Use:   "create-schedule",
	Short: "CreateSchedule",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.CreateScheduleRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.CreateSchedule(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var getScheduleMoabV0Cmd = &cobra.Command{
	Use:   "get-schedule",
	Short: "GetSchedule",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.GetScheduleRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.GetSchedule(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var updateScheduleMoabV0Cmd = &cobra.Command{
	Use:   "update-schedule",
	Short: "UpdateSchedule",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.UpdateScheduleRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.UpdateSchedule(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var deleteScheduleMoabV0Cmd = &cobra.Command{
	Use:   "delete-schedule",
	Short: "DeleteSchedule",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.DeleteScheduleRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.DeleteSchedule(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

var listSchedulesMoabV0Cmd = &cobra.Command{
	Use:   "list-schedules",
	Short: "ListSchedules",
	Run: func(cmd *cobra.Command, args []string) {
		req := &moab_v0.ListSchedulesRequest{}
		readRequest(req)

		client := getMoabV0Client()

		resp, err := client.ListSchedules(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(resp)
	},
}

func init() {
	rootCmd.AddCommand(moabV0Cmd)

	moabV0Cmd.AddCommand(createQueueMoabV0Cmd)
	moabV0Cmd.AddCommand(getQueueMoabV0Cmd)
	moabV0Cmd.AddCommand(updateQueueMoabV0Cmd)
	moabV0Cmd.AddCommand(deleteQueueMoabV0Cmd)
	moabV0Cmd.AddCommand(listQueuesMoabV0Cmd)
	moabV0Cmd.AddCommand(getTaskMoabV0Cmd)
	moabV0Cmd.AddCommand(listTasksMoabV0Cmd)
	moabV0Cmd.AddCommand(enqueueMoabV0Cmd)
	moabV0Cmd.AddCommand(dequeueMoabV0Cmd)
	moabV0Cmd.AddCommand(reportStatusMoabV0Cmd)
	moabV0Cmd.AddCommand(deleteTasksMoabV0Cmd)
	moabV0Cmd.AddCommand(restartTasksMoabV0Cmd)
	moabV0Cmd.AddCommand(purgeQueueMoabV0Cmd)
	moabV0Cmd.AddCommand(createScheduleMoabV0Cmd)
	moabV0Cmd.AddCommand(getScheduleMoabV0Cmd)
	moabV0Cmd.AddCommand(updateScheduleMoabV0Cmd)
	moabV0Cmd.AddCommand(deleteScheduleMoabV0Cmd)
	moabV0Cmd.AddCommand(listSchedulesMoabV0Cmd)

	moabV0Cmd.PersistentFlags().StringVarP(&moabV0CmdCfg.apiKeyId, "api-key-id", "", "", "API key ID (key_alfa_* or key_bravo_*)")
	moabV0Cmd.PersistentFlags().StringVarP(&moabV0CmdCfg.apiSecretKey, "api-secret-key", "", "", "API secret key")

	moabV0Cmd.PersistentFlags().StringVarP(&moabV0CmdCfg.endpoint, "endpoint", "", "", "Moab API address")
	err := moabV0Cmd.MarkPersistentFlagRequired("endpoint")
	if err != nil {
		log.Fatal(err)
	}
}

func getMoabV0Client() moab_v0.MoabApi {
	client, err := moab_v0.NewMoabGrpcClient(moabV0CmdCfg.endpoint, getSigner(moabV0CmdCfg.apiKeyId, moabV0CmdCfg.apiSecretKey))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
