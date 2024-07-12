// tests/bootstrap_test.go
package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/uzzalhcse/go-fiber-kit/bootstrap"
	"sync"
	"testing"
)

func TestAppSingleton(t *testing.T) {
	// Reset the singleton instance before the test
	resetSingleton()

	// First call to App function
	appInstance1 := bootstrap.App()

	// Second call to App function
	appInstance2 := bootstrap.App()

	// Ensure that both instances are the same
	assert.Same(t, appInstance1, appInstance2, "App function does not return a singleton instance")
}

func resetSingleton() {
	// Reset the singleton instance before the test
	bootstrap.Once = sync.Once{}
	bootstrap.AppInstance = nil
}
