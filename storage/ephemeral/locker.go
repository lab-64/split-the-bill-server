package ephemeral

import (
	"sync"
)

// Locker holds mutexes for the managed resources and avoids deadlocks by ordering the locks, locking in order and
// unlocking in reverse order
type Locker struct {
	locks []sync.Mutex
}

func NewLocker() *Locker {
	locks := make([]sync.Mutex, NumResources)
	return &Locker{locks: locks}
}

// Lock locks the mutexes for the given resources and returns a combined resource that can be used to unlock.
// The res parameter is mandatory to avoid misuse of the function (e.g. locking without specifying any resource).
// Other than that the function is indifferent to the order of the parameters.
//
// Usage example:
// r := l.Lock(RUsers, RNameIndex)
// defer l.Unlock(r)
func (l *Locker) Lock(res Resource, resources ...Resource) Resource {
	r := combine(res, resources...)
	for i := range l.locks {
		if r&(1<<uint(i)) != 0 {
			l.locks[i].Lock()
		}
	}
	return r
}

// Unlock unlocks the mutexes for the given resource(s). Passing at least one resource is mandatory to avoid misuse
// (e.g. unlocking without specifying any resource). Other than that the function is indifferent to the order of the
// parameters.
func (l *Locker) Unlock(res Resource, resources ...Resource) {
	r := combine(res, resources...)
	// unlock in reverse order
	for i := len(l.locks) - 1; i >= 0; i-- {
		if r&(1<<uint(i)) != 0 {
			l.locks[i].Unlock()
		}
	}
}

func combine(r Resource, resources ...Resource) Resource {
	var res Resource = 0
	for _, resource := range resources {
		res |= resource
	}
	return res | r
}
