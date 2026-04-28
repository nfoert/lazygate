package plugin

import (
	"github.com/nfoert/lazygate/pkg/provider"
	"github.com/nfoert/lazygate/pkg/provider/docker"
	"github.com/nfoert/lazygate/pkg/provider/nomad"
	"github.com/nfoert/lazygate/pkg/provider/pufferpanel"
	"github.com/nfoert/lazygate/pkg/queue"
	"github.com/nfoert/lazygate/pkg/queue/kick"
	"github.com/nfoert/lazygate/pkg/queue/send"
	"github.com/nfoert/lazygate/pkg/queue/wait"
)

// Options contains customizable plugin options.
type Options struct {
	Queues    []queue.Queue       // Registered queues.
	Providers []provider.Provider // Registered providers.
}

// DefaultOptions returns options object with default parameters.
func DefaultOptions() *Options {
	return &Options{
		Queues: []queue.Queue{
			&kick.Queue{},
			&send.Queue{},
			&wait.Queue{},
		},
		Providers: []provider.Provider{
			&docker.Provider{},
			&nomad.Provider{},
			&pufferpanel.Provider{},
		},
	}
}
