// +build linux

/*
   Copyright The containerd Authors.

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

package mount

import (
	"github.com/moby/sys/mountinfo"
)

// Self retrieves a list of mounts for the current running process.
func Self() ([]Info, error) {
	m, err := mountinfo.GetMounts(nil)
	return fromMountinfoSlice(m), err
}

// PID collects the mounts for a specific process ID. If the process
// ID is unknown, it is better to use `Self` which will inspect
// "/proc/self/mountinfo" instead.
func PID(pid int) ([]Info, error) {
	m, err := mountinfo.PidMountInfo(pid)
	return fromMountinfoSlice(m), err
}
