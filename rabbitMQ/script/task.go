package script

import (
	"context"

	"go-cloud-disk/rabbitMQ/task"
	"go-cloud-disk/utils/logger"
)

func SendConfirmEmailSync(ctx context.Context) {
	err := task.RunSendConfirmEmail(ctx)
	if err != nil {
		logger.Log().Error("SendConfirmEmailSync error: ", err)
	}
}
