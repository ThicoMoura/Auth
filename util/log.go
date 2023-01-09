package util

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	Linfo    = " |* INFO *| "
	Lwarning = " |* WARNING *| "
	Lerror   = " |* ERROR *| "
	reset    = "\033[0m"
	red      = "\033[31m"
	blue     = "\033[36m"
	yellow   = "\033[33m"
)

type Log struct {
	out    io.Writer
	prefix string
	loc    *time.Location
}

func (log Log) formatHeader(buf *[]byte, t time.Time) {
	*buf = append(*buf, t.In(log.loc).Format("02/01/2006 15:04:05")...)
	if log.prefix != "" {
		switch log.prefix {
		case Lerror:
			*buf = append(*buf, reset+red+log.prefix+reset...)
		case Lwarning:
			*buf = append(*buf, reset+yellow+log.prefix+reset...)
		case Linfo:
			*buf = append(*buf, reset+blue+log.prefix+reset...)
		default:
			*buf = append(*buf, log.prefix...)
		}
	}
}

func (log Log) Output(s string) {
	var buf []byte
	log.formatHeader(&buf, time.Now())
	buf = append(buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}
	log.out.Write(buf)
}

func (log Log) Fatal(v ...any) {
	log.Output(fmt.Sprint(v...))
	os.Exit(1)
}

func (log Log) Fatalf(format string, v ...any) {
	log.Output(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (log Log) Printf(format string, v ...any) {
	log.Output(fmt.Sprintf(format, v...))
}

func (log Log) Println(v ...any) {
	log.Output(fmt.Sprint(v...))
}

func NewLogs(out io.Writer, prefix string, loc *time.Location) *Log {
	log := &Log{
		prefix: prefix,
		out:    out,
		loc:    loc,
	}
	return log
}
