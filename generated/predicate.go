/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

// PredicateFunc defines data-type for validation function
type PredicateFunc func(*volumeRollback) bool

// IsSetshouldDestroy method check if the shouldDestroy field of volumeRollback object is set.
func IsshouldDestroySet() PredicateFunc {
	return func(v *volumeRollback) bool {
		return v.shouldDestroy == true
	}
}

// IsSetforceunmount method check if the forceunmount field of volumeRollback object is set.
func IsforceunmountSet() PredicateFunc {
	return func(v *volumeRollback) bool {
		return v.forceunmount == true
	}
}

// IsSetshouldDestroySnap method check if the shouldDestroySnap field of volumeRollback object is set.
func IsshouldDestroySnapSet() PredicateFunc {
	return func(v *volumeRollback) bool {
		return v.shouldDestroySnap == true
	}
}

// IsSetsnapshot method check if the snapshot field of volumeRollback object is set.
func IssnapshotSet() PredicateFunc {
	return func(v *volumeRollback) bool {
		return len(v.snapshot) == 0
	}
}
