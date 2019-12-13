package docker

import (
	"github.com/crazy-max/diun/internal/model"
	"github.com/crazy-max/diun/internal/provider"
	"github.com/rs/zerolog/log"
)

// Client represents an active docker provider object
type Client struct {
	*provider.Client
	elts []model.PrdDocker
}

// New creates new docker provider instance
func New(elts []model.PrdDocker) *provider.Client {
	return &provider.Client{Handler: &Client{
		elts: elts,
	}}
}

// ListJob returns job list to process
func (c *Client) ListJob() []model.Job {
	if len(c.elts) == 0 {
		return []model.Job{}
	}

	log.Info().Msgf("Found %d docker provider(s) to analyze...", len(c.elts))
	var list []model.Job
	for _, elt := range c.elts {
		// Swarm mode
		if elt.SwarmMode {
			continue
		}

		// Docker
		for _, img := range c.listContainerImage(elt) {
			list = append(list, model.Job{
				Provider: "docker",
				ID:       elt.ID,
				Image:    img,
			})
		}
	}

	return list
}
