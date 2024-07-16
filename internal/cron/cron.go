package cron

import "github.com/updevru/go-micro-kit/server"

func NewCron(cleaner *Cleaner) []server.CronTask {
	tasks := make([]server.CronTask, 0)

	tasks = append(tasks, server.CronTask{
		Name: "Cleaner",
		Cron: "* * * * *",
		Fn:   cleaner.Clean,
	})

	return tasks
}
