package whitelist

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/processwhitelist"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stackrox/rox/pkg/sync"
)

// Evaluator encapsulates the interface to the baseline evaluator
type Evaluator interface {
	RemoveDeployment(id string)
	AddWhitelist(whitelist *storage.ProcessWhitelist)
	IsOutsideLockedWhitelist(pi *storage.ProcessIndicator) bool
}

type whitelistEvaluator struct {
	// deployment -> container name -> exec file paths within baseline
	whitelists    map[string]map[string]set.StringSet
	whitelistLock sync.RWMutex
}

// NewWhitelistEvaluator creates a new baseline evaluator
func NewWhitelistEvaluator() Evaluator {
	return &whitelistEvaluator{
		whitelists: make(map[string]map[string]set.StringSet),
	}
}

// RemoveDeployment removes the baselines for this specific deployment
func (w *whitelistEvaluator) RemoveDeployment(id string) {
	w.whitelistLock.Lock()
	defer w.whitelistLock.Unlock()

	delete(w.whitelists, id)
}

// AddWhitelist adds a baseline to the store
// If the baseline is unlocked, then we remove the baseline references because for the purposes
// of this package, an unlocked baseline has no impact. Locked baselines will have all of the processes
// added to a map
func (w *whitelistEvaluator) AddWhitelist(whitelist *storage.ProcessWhitelist) {
	// We'll get this msg with an unlocked baseline if a user unlocks a baseline
	// so we need to purge it from the baseline
	if whitelist.GetUserLockedTimestamp() == nil {
		w.whitelistLock.Lock()
		defer w.whitelistLock.Unlock()

		delete(w.whitelists[whitelist.GetKey().GetDeploymentId()], whitelist.GetKey().GetContainerName())
		return
	}

	// Create the baseline and overwrite the value in the map
	// We'll receive this message for all user locked baselines
	whitelistSet := set.NewStringSet()
	for _, elem := range whitelist.GetElements() {
		if process := elem.GetElement().GetProcessName(); process != "" {
			whitelistSet.Add(process)
		}
	}

	w.whitelistLock.Lock()
	defer w.whitelistLock.Unlock()

	containerNameMap := w.whitelists[whitelist.GetKey().GetDeploymentId()]
	if containerNameMap == nil {
		containerNameMap = make(map[string]set.StringSet)
		w.whitelists[whitelist.GetKey().GetDeploymentId()] = containerNameMap
	}
	containerNameMap[whitelist.GetKey().GetContainerName()] = whitelistSet
}

// IsInWhitelist checks if the process indicator is within a locked baseline
// If the baseline does not exist, then we return true
func (w *whitelistEvaluator) IsOutsideLockedWhitelist(pi *storage.ProcessIndicator) bool {
	w.whitelistLock.RLock()
	defer w.whitelistLock.RUnlock()

	whitelist := w.whitelists[pi.GetDeploymentId()][pi.GetContainerName()]
	// If there is no baseline, then we are counting it as if it's within the baseline
	return whitelist != nil && !whitelist.Contains(processwhitelist.WhitelistItemFromProcess(pi))
}
