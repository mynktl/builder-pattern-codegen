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

// SetshouldDestroy method set the shouldDestroy field of volumeRollback object.
func (v *volumeRollback) SetshouldDestroy(shouldDestroy bool) {
	v.shouldDestroy = shouldDestroy
}

// Setforceunmount method set the forceunmount field of volumeRollback object.
func (v *volumeRollback) Setforceunmount(forceunmount bool) {
	v.forceunmount = forceunmount
}

// SetshouldDestroySnap method set the shouldDestroySnap field of volumeRollback object.
func (v *volumeRollback) SetshouldDestroySnap(shouldDestroySnap bool) {
	v.shouldDestroySnap = shouldDestroySnap
}

// Setsnapshot method set the snapshot field of volumeRollback object.
func (v *volumeRollback) Setsnapshot(snapshot string) {
	v.snapshot = snapshot
}

// GetshouldDestroy method get the shouldDestroy field of volumeRollback object.
func (v *volumeRollback) GetshouldDestroy() bool {
	return v.shouldDestroy
}

// Getforceunmount method get the forceunmount field of volumeRollback object.
func (v *volumeRollback) Getforceunmount() bool {
	return v.forceunmount
}

// GetshouldDestroySnap method get the shouldDestroySnap field of volumeRollback object.
func (v *volumeRollback) GetshouldDestroySnap() bool {
	return v.shouldDestroySnap
}

// Getsnapshot method get the snapshot field of volumeRollback object.
func (v *volumeRollback) Getsnapshot() string {
	return v.snapshot
}
