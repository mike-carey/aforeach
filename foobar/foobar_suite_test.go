package foobar_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFoobar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Foobar Suite")
}
