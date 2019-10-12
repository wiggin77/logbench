package logbench

import (
	"io"
	"time"

	"github.com/wiggin77/logr"
	"github.com/wiggin77/logr/format"
	"github.com/wiggin77/logr/target"
)

func init() {
	tests["Logr"] = logrTester{}
}

type logrTester struct {
	l   *logr.Logger
	lgr *logr.Logr
}

var (
	_ logTesterArray = (*logrTester)(nil)
)

func (logrTester) newLogger(out io.Writer, disabled bool) logTester {
	lvl := logr.Debug
	if disabled {
		lvl = logr.Panic
	}

	lgr := &logr.Logr{MaxQueueSize: -1}
	formatter := &format.JSON{DisableTimestamp: true, KeyMsg: "message"}
	filter := logr.StdFilter{Lvl: lvl}
	target := target.NewWriterTarget(filter, formatter, out, 0)
	lgr.AddTarget(target)

	logger := lgr.NewLogger()

	return logrTester{
		l:   logger,
		lgr: lgr,
	}
}

func (t logrTester) logMsg(msg string) {
	t.l.Info(msg)
}

func (t logrTester) logFormat(format string, v ...interface{}) bool {
	t.l.Infof(format, v...)
	return true
}

func (t logrTester) withContext(context map[string]interface{}) (logTester, bool) {
	return logrTester{t.l.WithFields(context), t.lgr}, true
}

func (t logrTester) logBool(msg, key string, value bool) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}

func (t logrTester) logInt(msg, key string, value int) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}

func (t logrTester) logFloat32(msg, key string, value float32) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}

func (t logrTester) logFloat64(msg, key string, value float64) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}

func (t logrTester) logTime(msg, key string, value time.Time) bool {
	return false
}

func (t logrTester) logDuration(msg, key string, value time.Duration) bool {
	return false
}

func (t logrTester) logError(msg, key string, value error) bool {
	return false
}

func (t logrTester) logString(msg, key string, value string) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}

func (t logrTester) logObject(msg, key string, value *obj) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}

func (t logrTester) logBools(msg, key string, value []bool) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}

func (t logrTester) logInts(msg, key string, value []int) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}

func (t logrTester) logFloats32(msg, key string, value []float32) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}

func (t logrTester) logFloats64(msg, key string, value []float64) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}

func (t logrTester) logTimes(msg, key string, value []time.Time) bool {
	return false
}

func (t logrTester) logDurations(msg, key string, value []time.Duration) bool {
	return false
}

func (t logrTester) logErrors(msg, key string, value []error) bool {
	return false
}

func (t logrTester) logStrings(msg, key string, value []string) bool {
	t.l.WithFields(logr.Fields{key: value}).Info(msg)
	return true
}
