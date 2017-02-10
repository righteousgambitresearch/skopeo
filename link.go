package buildah

import (
	"fmt"
	"os"
	"path/filepath"
)

func (b *Builder) Link(location string) error {
	if b.MountPoint == "" {
		return fmt.Errorf("build container is not mounted")
	}
	path, err := filepath.Abs(location)
	if err != nil {
		return fmt.Errorf("error resolving symbolic link %q to an absolute path: %v", location, err)
	}
	err = os.Symlink(b.MountPoint, path)
	if err != nil {
		return fmt.Errorf("error creating symbolic link %q: %v", path, err)
	}
	present := false
	for _, l := range b.Links {
		if l == path {
			present = true
			break
		}
	}
	if !present {
		b.Links = append(b.Links, path)
	}
	return b.Save()
}
