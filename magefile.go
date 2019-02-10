//+build mage

package main

import "github.com/magefile/mage/sh"

// Runs dep ensure and then installs the binary.
func Build() error {
	if err := sh.Run("dep", "ensure"); err != nil {
		return err
	}
	return sh.Run("go", "install", "./...")
}
