package gui

import "fmt"

//go:generate go run github.com/vektra/mockery/cmd/mockery -name terminalOutputter -inpkg --filename terminal_outputter_mock.go
type terminalOutputter interface {
	Printf(format string, a ...interface{}) error
}

type terminalOutput struct{}

func (*terminalOutput) Printf(format string, a ...interface{}) error {
	_, err := fmt.Printf(format, a...)
	return err
}
