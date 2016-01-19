package metricsd

import (
	"fmt"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/megamsys/megamd/meta"
	"github.com/megamsys/megamd/metrix"
	"github.com/megamsys/megamd/subd/deployd"
)

const (
	RIAK = "riak"
)

var OUTPUTS = []string{RIAK}

// Service manages the listener and handler for an HTTP endpoint.
type Service struct {
	err     chan error
	Handler *Handler
	stop    chan struct{}
	wg      *sync.WaitGroup
	Meta    *meta.Config
	Deployd *deployd.Config
	Config  *Config
}

// NewService returns a new instance of Service.
func NewService(c *meta.Config, d *deployd.Config, f *Config) *Service {
	s := &Service{
		err:     make(chan error),
		Meta:    c,
		Deployd: d,
		Config:  f,
	}
	s.Handler = NewHandler()
	return s
}

// Open starts the service
func (s *Service) Open() error {
	log.Info("starting metricsd service")
	if s.stop != nil {
		return nil
	}

	s.stop = make(chan struct{})
	s.wg = &sync.WaitGroup{}
	s.wg.Add(1)

	go s.backgroundLoop()
	return nil
}

func (s *Service) backgroundLoop() {
	defer s.wg.Done()
	for {
		select {
		case <-s.stop:
			log.Info("metricd terminating")
		case <-time.After(time.Duration(s.Config.CollectInterval)):
			s.runMetricsCollectors()
		}
	}

}

func (s *Service) runMetricsCollectors() error {
	output := &metrix.OutputHandler{
		RiakAddress: s.Meta.Riak,
	}

	collectors := map[string]metrix.MetricCollector{
		metrix.OPENNEBULA: &metrix.OpenNebula{Url: s.Deployd.OneEndPoint},
	}

	mh := &metrix.MetricHandler{}

	values := make(map[string]string)

	for key, _ := range values {
		if collector, ok := collectors[key]; ok {
			go s.Handler.processCollector(mh, output, collector)
		} else {
			isOutput := false
			for _, out := range OUTPUTS {
				if out == key {
					isOutput = true
					break
				}
			}
			if !isOutput {
				return fmt.Errorf("no collector found for %q\n", key)
			}
		}
	}
	return nil
}

func (s *Service) Close() error {
	if s.stop == nil {
		return nil
	}
	close(s.stop)
	s.wg.Wait()
	s.wg = nil
	s.stop = nil
	return nil

}

// Err returns a channel for fatal errors that occur on the listener.
func (s *Service) Err() <-chan error { return s.err }