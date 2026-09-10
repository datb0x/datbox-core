package datboxcore

import (
	"testing"
)

// TestLogLevels prints logs at all available levels
func TestLogLevels(t *testing.T) {
	logger := DefaultLogger()
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	logger.Print("print")
	logger.Println()
}

// TestMultiLogger prints logs with different loggers
func TestMultiLogger(t *testing.T) {
	defaultLogger := DefaultLogger()
	shortLogger, err := NewLogger("s")
	if err != nil {
		t.Errorf("Couldn't create short logger: %v", err)
	}
	longLogger, err := NewLogger("loooooooong")
	if err != nil {
		t.Errorf("Couldn't create long logger: %v", err)
	}
	defaultLogger.Info("default")
	shortLogger.Info("short")
	longLogger.Info("long")
}

func TestDuplicateLogger(t *testing.T) {
	_, err := NewLogger("main")
	if err == nil {
		t.Error("Created duplicate 'main' without error, which is wrong")
	}
	_, err = NewLogger("l1")
	if err != nil {
		t.Errorf("Couldn't create l1 logger: %v", err)
	}
	_, err = NewLogger("l1")
	if err == nil {
		t.Error("Created duplicate 'l1' without error, which is wrong")
	}
}
