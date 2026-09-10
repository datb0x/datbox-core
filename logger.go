package datboxcore

import (
	"errors"
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/log/v2"
)

type logger struct {
	loggers       map[string]*Logger
	prefixPadding int
}

type Logger struct {
	Name string
}

var (
	globalLogger  logger
	defaultLogger *Logger
	prefixStyle   = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)

	ErrLoggerExists = errors.New("Logger already exists")
)

func init() {
	globalLogger.loggers = make(map[string]*Logger)
	defaultLogger, _ = NewLogger("main")
}

func DefaultLogger() *Logger {
	return defaultLogger
}

func NewLogger(name string) (*Logger, error) {
	if _, exists := globalLogger.loggers[name]; exists {
		return nil, ErrLoggerExists
	}

	newLogger := Logger{Name: name}
	globalLogger.loggers[name] = &newLogger
	if len(name) > globalLogger.prefixPadding {
		globalLogger.prefixPadding = len(name)
	}
	return &newLogger, nil
}

func (logger *Logger) getPrefix() string {
	return prefixStyle.Render(fmt.Sprintf(fmt.Sprintf("%%%ds ", globalLogger.prefixPadding), logger.Name))
}

func (logger *Logger) Debugf(format string, a ...any) {
	log.Debug(logger.getPrefix() + fmt.Sprintf(format, a...))
}

func (logger *Logger) Debug(a ...any) {
	log.Debug(logger.getPrefix() + fmt.Sprint(a...))
}

func (logger *Logger) Infof(format string, a ...any) {
	log.Info(logger.getPrefix() + fmt.Sprintf(format, a...))
}

func (logger *Logger) Info(a ...any) {
	log.Info(logger.getPrefix() + fmt.Sprint(a...))
}

func (logger *Logger) Warnf(format string, a ...any) {
	log.Warn(logger.getPrefix() + fmt.Sprintf(format, a...))
}

func (logger *Logger) Warn(a ...any) {
	log.Warn(logger.getPrefix() + fmt.Sprint(a...))
}

func (logger *Logger) Errorf(format string, a ...any) {
	log.Error(logger.getPrefix() + fmt.Sprintf(format, a...))
}

func (logger *Logger) Error(a ...any) {
	log.Error(logger.getPrefix() + fmt.Sprint(a...))
}

func (logger *Logger) Fatalf(format string, a ...any) {
	log.Fatal(logger.getPrefix() + fmt.Sprintf(format, a...))
}

func (logger *Logger) Fatal(a ...any) {
	log.Fatal(logger.getPrefix() + fmt.Sprint(a...))
}

func (logger *Logger) Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func (logger *Logger) Print(a ...any) {
	fmt.Print(a...)
}

func (logger *Logger) Println(a ...any) {
	fmt.Println(a...)
}
