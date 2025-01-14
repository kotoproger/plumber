package worker

import (
	"github.com/nexus-rpc/sdk-go/nexus"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type WorkerDecorator struct {
	inner worker.Worker
}

func (d WorkerDecorator) RegisterWorkflow(w interface{}) {
	d.inner.RegisterWorkflow(w)
}

func (d WorkerDecorator) RegisterWorkflowWithOptions(w interface{}, options workflow.RegisterOptions) {
	d.inner.RegisterWorkflowWithOptions(w, options)
}

func (d WorkerDecorator) RegisterActivity(a interface{}) {
	d.inner.RegisterActivity(a)
}

func (d WorkerDecorator) RegisterActivityWithOptions(a interface{}, options activity.RegisterOptions) {
	d.inner.RegisterActivityWithOptions(a, options)
}

func (d WorkerDecorator) RegisterNexusService(service *nexus.Service) {
	d.inner.RegisterNexusService(service)
}

func (d WorkerDecorator) Start() error {
	return d.inner.Start()
}

func (d WorkerDecorator) Run(interruptCh <-chan interface{}) error {
	return d.inner.Run(interruptCh)
}

func (d WorkerDecorator) Stop() {
	d.inner.Stop()
}

func NewPlumberWorker(w worker.Worker) WorkerDecorator {
	return WorkerDecorator{
		inner: w,
	}
}
