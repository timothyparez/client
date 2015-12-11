// Copyright 2015 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

// +build windows

package libdokan

import (
	"github.com/keybase/kbfs/dokan"
	"golang.org/x/net/context"
)

// Symlink represents KBFS symlinks.
type Symlink struct {
	// The directory this symlink is in. This should be safe to store
	// here, without fear of Renames etc making it stale, because we
	// never persist a Symlink into Folder.nodes; it has no
	// libkbfs.Node, so that's impossible. This should make FUSE
	// Lookup etc always get new nodes, limiting the lifetime of a
	// single Symlink value.
	parent *Dir
	name   string
	emptyFile
}

// GetFileInformation does stat for dokan.
func (s *Symlink) GetFileInformation(*dokan.FileInfo) (a *dokan.Stat, err error) {
	ctx := context.TODO()
	ctx = NewContextWithOpID(ctx, s.parent.folder.fs.log)
	s.parent.folder.fs.log.CDebugf(ctx, "Symlink Attr")
	defer func() { s.parent.folder.fs.reportErr(ctx, err) }()

	_, _, err = s.parent.folder.fs.config.KBFSOps().Lookup(ctx, s.parent.node, s.name)
	if err != nil {
		return nil, errToDokan(err)
	}

	return defaultSymlinkInformation()
}
