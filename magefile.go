//+build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Test - mage run
func Test() error {
	return sh.RunV("go", "test", "-v", "-cover", "./...", "-coverprofile=coverage.out")
}

// Coverage - checking code coverage
func Coverage() error {
	mg.Deps(Test)
	return sh.RunV("go", "tool", "cover", "-html=coverage.out")
}
