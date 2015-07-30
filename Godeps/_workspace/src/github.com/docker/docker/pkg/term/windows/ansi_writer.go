// +build windows

package windows

import (
	"io/ioutil"
	"os"

	. "github.com/Azure/go-ansiterm"
	. "github.com/Azure/go-ansiterm/winterm"
	"github.com/Sirupsen/logrus"
)

var logger *logrus.Logger

// ansiWriter wraps a standard output file (e.g., os.Stdout) providing ANSI sequence translation.
type ansiWriter struct {
	file           *os.File
	fd             uintptr
	infoReset      *CONSOLE_SCREEN_BUFFER_INFO
	command        []byte
	escapeSequence []byte
	inAnsiSequence bool
	parser         *AnsiParser
}

func newAnsiWriter(nFile int) *ansiWriter {
	logFile := ioutil.Discard

	if isDebugEnv := os.Getenv(LogEnv); isDebugEnv == "1" {
		logFile, _ = os.Create("ansiReaderWriter.log")
	}

	logger = &logrus.Logger{
		Out:       logFile,
		Formatter: new(logrus.TextFormatter),
		Level:     logrus.DebugLevel,
	}

	file, fd := GetStdFile(nFile)
	info, err := GetConsoleScreenBufferInfo(fd)
	if err != nil {
		return nil
	}

	parser := CreateParser("Ground", CreateWinEventHandler(fd, file))
	logger.Infof("newAnsiWriter: parser %p", parser)

	aw := &ansiWriter{
		file:           file,
		fd:             fd,
		infoReset:      info,
		command:        make([]byte, 0, ANSI_MAX_CMD_LENGTH),
		escapeSequence: []byte(KEY_ESC_CSI),
		parser:         parser,
	}

	logger.Infof("newAnsiWriter: aw.parser %p", aw.parser)
	logger.Infof("newAnsiWriter: %v", aw)
	return aw
}

func (aw *ansiWriter) Fd() uintptr {
	return aw.fd
}

// Write writes len(p) bytes from p to the underlying data stream.
func (aw *ansiWriter) Write(p []byte) (total int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	logger.Infof("Write: % x", p)
	logger.Infof("Write: %s", string(p))
	return aw.parser.Parse(p)
}
