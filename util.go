package common

import (
	"fmt"
	"os"
	"path"
)

func CheckWritable(dir string) error {
	_, err := os.Stat(dir)
	if err == nil {
		// dir exists, make sure we can write to it
		testfile := path.Join(dir, "test")
		fi, err := os.Create(testfile)
		if err != nil {
			if os.IsPermission(err) {
				return fmt.Errorf("%s is not writeable by the current user", dir)
			}
			return fmt.Errorf("unexpected error while checking writeablility of repo root: %s", err)
		}
		_ = fi.Close()
		return os.Remove(testfile)
	}

	if os.IsNotExist(err) {
		// dir doesn't exist, check that we can create it
		return os.Mkdir(dir, 0775)
	}

	if os.IsPermission(err) {
		return fmt.Errorf("cannot write to %s, incorrect permissions", err)
	}

	return err
}

func CheckPermissions(path string) (bool, error) {
	_, err := os.Open(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if os.IsPermission(err) {
		return false, fmt.Errorf("error opening directory at %s: permission denied", path)
	}

	return true, nil
}
