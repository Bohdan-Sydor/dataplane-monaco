package runtask

import (
	"context"
	"dataplane/mainapp/database/models"
	modelmain "dataplane/mainapp/database/models"
	"dataplane/workers/config"
	"dataplane/workers/messageq"
	"log"
	"os"
	"syscall"
	"time"
)

type TaskResponse struct {
	R string
	M string
}

func ListenTasks() {

	// Responding to a task request
	messageq.NATSencoded.Subscribe("task."+os.Getenv("worker_group")+"."+config.WorkerID, func(subj, reply string, msg models.WorkerTaskSend) {
		// log.Println(msg)

		response := "ok"
		message := "ok"
		// msg.EnvironmentID
		if config.EnvID != msg.EnvironmentID {
			response = "failed"
			message = "Incorrect environment"
			if config.Debug == "true" {
				log.Println("response", response, message)
			}

			TaskFinal := modelmain.WorkerTasks{
				TaskID:        msg.TaskID,
				EnvironmentID: config.EnvID,
				RunID:         msg.RunID,
				WorkerID:      config.WorkerID,
				NodeID:        msg.NodeID,
				PipelineID:    msg.PipelineID,
				Status:        "Fail",
				Reason:        message,
				EndDT:         time.Now().UTC(),
			}
			UpdateWorkerTasks(TaskFinal)
		}

		x := TaskResponse{R: response, M: message}
		messageq.NATSencoded.Publish(reply, x)

		if x.R == "ok" {
			TaskID := msg.TaskID
			ctx, cancel := context.WithCancel(context.Background())
			var task Task

			task.ID = TaskID
			task.Context = ctx
			task.Cancel = cancel

			// Tasks[task.ID] = task
			Tasks.Set(task.ID, task)
			// command := `for((i=1;i<=10000; i+=1)); do echo "Welcome $i times"; sleep 1; done`
			// command := `find . | sed -e "s/[^ ][^\/]*\// |/g" -e "s/|\([^ ]\)/| \1/"`
			go worker(ctx, msg)
		}
	})
	if config.Debug == "true" {
		log.Println("Listening for tasks on subject:", "task."+os.Getenv("worker_group")+"."+config.WorkerID)
	}

	messageq.NATSencoded.Subscribe("taskcancel."+os.Getenv("worker_group")+"."+config.WorkerID, func(subj, reply string, msg models.WorkerTaskSend) {
		// Respond to cancelling a task
		id := msg.TaskID

		var TasksRun Task

		if tmp, ok := Tasks.Get(id); ok {
			TasksRun = tmp.(Task)
		}

		if TasksRun.PID != 0 {
			_ = syscall.Kill(-TasksRun.PID, syscall.SIGKILL)
		}
		TasksRun.Cancel()
		// TasksStatus[id] = "cancel"
		TasksStatus.Set(id, "cancel")

		response := "ok"
		message := "ok"
		x := TaskResponse{R: response, M: message}
		messageq.NATSencoded.Publish(reply, x)

	})

}
