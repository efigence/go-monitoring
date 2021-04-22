package monitoring

import (
	"time"
)

const (
	// 0 is invalid state because that is initial value of int variable which can mean that someone didn't bother to actually set the state
	StatusInvalid = iota
	// nagios-compatible block
	// to get compatible state take nagios one and add +1
	// Service is ok
	StatusOk
	// service is in warning state
	// should be only used if service is *actually* working but have some problems that need to be resolved
	// like "disk getting full" or "worker queue is 90% busy"
	StatusWarning
	// Service is in critical state and is not performing its function
	StatusCritical
	// check failed to get status of service (for reason other than "service is not working)
	// i.e. check itself crashed before providing any useful information about service
	StatusUnknown
)

const (
	// 0 is invalid state because that is initial value of int variable which can mean that someone didn't bother to actually set the state
	HostInvalid = iota
	// host (as in "unit running service checks") is up
	HostUp
	// host is directly unavailable
	HostDown
	// host is down because its parent is down (it is impossble to check because device that connects to the host is unavailable
	HostUnreachable
)

type Common struct {
	// Whether upstream acknowledged state (eg. started doing something about it)
	Acknowledged bool `json:"acknowledged"`
	// Whether check is flapping between states
	Flapping bool `json:"flapping"`
	// Whether it is in downtime
	Downtime bool `json:"downtime"`
	// numeric state
	State uint8 `json:"state"`
	PreviousState uint8 `json:"previous_state"`
	StateHard bool `json:"state_hard"`
	// timestamp of the check
	Timestamp time.Time `json:"ts"`
	// Last state change
	LastStateChange time.Time `json:"last_state_change,omitempty"`
	LastHardStateChange time.Time `json:"last_hard_state_change,omitempty"`
	// URL service/host can be looked at
	URL string `json:"url,omitempty"`
	// message returned from healthcheck
	CheckMessage string `json:"message"`
	// display name set for the service/host
	DisplayName string `json:"display_name"`
}


type Service struct {
	Common
	// name of the host/metahost service is running on
	Host string `json:"host"`
	// name of service
	Service string `json:"service"`
	// sub-service state
	// if service (say web app) have multiple internal components (for example DB backend, video transcoder etc) that allows it to send state of them to the upstream without multiplying amount of service checks
	// Note that status of them **HAVE** to be aggregated into parent's State
	Components []Service `json:"components,omitempty"`
}

type Host struct {
	Common
	Host string `json:"host"`
}