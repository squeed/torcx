// Copyright 2017 CoreOS Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package torcx

import (
	"path/filepath"
)

const (
	// DefaultRunDir is the default path where torcx unpacks/propagates all runtime assets.
	DefaultRunDir = "/run/torcx/"
	// DefaultBaseDir is the default torcx base directory
	DefaultBaseDir = "/var/lib/torcx/"
	// DefaultConfDir is the default torcx config directory
	DefaultConfDir = "/etc/torcx/"

	// OemStoreDir is the OEM store path
	OemStoreDir = OemDir + "store"
	// OemProfilesDir is the OEM profiles path
	OemProfilesDir = OemDir + "profiles"
	// OemRemotesDir is the OEM remotes path
	OemRemotesDir = OemDir + "remotes"

	// defaultCfgPath is the default path for common torcx config
	defaultCfgPath = DefaultConfDir + "config.json"
)

// VendorRemotesDir is the vendor remotes path
func VendorRemotesDir(usrMountpoint string) string {
	if usrMountpoint == "" {
		usrMountpoint = VendorUsrDir
	}
	return filepath.Join(usrMountpoint, "share", "torcx", "remotes")
}

// VendorProfilesDir is the vendor profiles path
func VendorProfilesDir(usrMountpoint string) string {
	if usrMountpoint == "" {
		usrMountpoint = VendorUsrDir
	}
	return filepath.Join(usrMountpoint, "share", "torcx", "profiles")
}

// VendorStoreDir is the vendor store path
func VendorStoreDir(usrMountpoint string) string {
	if usrMountpoint == "" {
		usrMountpoint = VendorUsrDir
	}
	return filepath.Join(usrMountpoint, "share", "torcx", "store")
}

// RunUnpackDir is the directory where root filesystems are unpacked.
func (cc *CommonConfig) RunUnpackDir() string {
	return filepath.Join(cc.RunDir, "unpack")
}

// RunBinDir is the directory where binaries are symlinked.
func (cc *CommonConfig) RunBinDir() string {
	return filepath.Join(cc.RunDir, "bin")
}

// ProfileDirs are the list of directories where we look for profiles.
func (cc *CommonConfig) ProfileDirs() []string {
	return []string{
		VendorProfilesDir(cc.UsrDir),
		OemProfilesDir,
		cc.UserProfileDir(),
	}
}

// RunProfile is the file where we copy the contents of the applied profile.
func (cc *CommonConfig) RunProfile() string {
	return filepath.Join(cc.RunDir, "profile.json")
}

// UserStorePath is the path where user-fetched archives are written.
// An optional target version can be specified for versioned user store.
func (cc *CommonConfig) UserStorePath(version string) string {
	storePath := filepath.Join(cc.BaseDir, "store")
	if version != "" {
		storePath = filepath.Join(storePath, version)
	}
	return storePath
}

// UserProfileDir is where user profiles are written.
func (cc *CommonConfig) UserProfileDir() string {
	return filepath.Join(cc.ConfDir, "profiles")
}

// NextProfile is the path for the `next-profile` selector configuration file.
func (cc *CommonConfig) NextProfile() string {
	return filepath.Join(cc.ConfDir, "next-profile")
}

// RemotesDirs returns the list of directories where we look for remotes manifests.
func (cc *CommonConfig) RemotesDirs() []string {
	dirs := []string{}
	if cc != nil {
		dirs = append(dirs, VendorRemotesDir(cc.UsrDir))
	}
	dirs = append(dirs, OemRemotesDir)
	if cc != nil {
		dirs = append(dirs, filepath.Join(cc.ConfDir, "remotes"))
	}
	return dirs
}

// VendorOsReleasePath returns the path to vendor os-release file
// for the specific OS partition mounted at `usrMountpoint`.
func VendorOsReleasePath(usrMountpoint string) string {
	if usrMountpoint == "" {
		usrMountpoint = VendorUsrDir
	}
	return filepath.Join(usrMountpoint, "lib", "os-release")
}
