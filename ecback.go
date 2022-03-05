package ecback

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

type DataError struct {
	ErrorText     string `json:"error_text"`
	StartingPoint string `json:"starting_point"`
	Trace         string `json:"trace"`
}

type ECBack struct {
	SCS           *spew.ConfigState
	DataError     DataError
	JsonString    string
	StopExecution bool
	CallBackUrl   string
	Trace         string
	ShowLog       bool
}

func InitErrCallBack(e ECBack) *ECBack {
	e.SCS = &spew.ConfigState{Indent: "  ", SortKeys: true}

	return &e
}

func (e *ECBack) E(err error, callback func(*ECBack) *ECBack) *error {
	if err == nil {
		return nil
	}

	e.getFileAndLine()
	e.DataError.ErrorText = err.Error()
	e.Trace = e.SCS.Sdump(err)
	e.DataError.Trace = strings.Replace(e.Trace, "\n", "", -1)

	marshal, err := json.Marshal(e.DataError)
	if err != nil {
		return &err
	}

	e.JsonString = string(marshal)

	if e.ShowLog {
		fmt.Println(e.JsonString)
	}

	if callback != nil {
		errCatcher := callback(e)
		if errCatcher != nil {
			if errCatcher.StopExecution {
				syscall.Exit(1)
			}
		}
	}

	return nil
}

func (e *ECBack) getFileAndLine() {
	_, file, line, _ := runtime.Caller(2)
	e.DataError.StartingPoint = fmt.Sprintf("%s:%d", filepath.Base(file), line-1)
}
