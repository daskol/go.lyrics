package logging

import "testing"

func TestNew(t *testing.T) {
	logger := New()
	logger.Infof("test message: ans=%d", 42)
}

func TestDefault(t *testing.T) {
	t.Run("InfoMessage", func(t *testing.T) {
		logger := Default()
		logger.Debugf("debug message: val=%d", 43)
	})
	t.Run("SetupVerbosity", func(t *testing.T) {
		logger := Default()
		logger.SetLevel(Debug)
	})
	t.Run("DebugMessage", func(t *testing.T) {
		logger := Default()
		logger.Debugf("debug message: val=%d", 44)
	})
}
