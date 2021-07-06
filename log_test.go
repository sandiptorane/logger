package logger

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	TAG        = "Log Test"
	testUserId = "1234"
)

var Input = map[string]interface{}{
	"tag":                 "tag_value",
	"message":             "ok",
	"duration":            4.902,
	"request_ip":          "36.71.235.223",
	"request_remote_addr": "130.211.2.191:62695",
	"request_uri":         "https://erp.hoge.com/users/1234",
	"request_pattern":     "users/:id",
	"request_user_agent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36",
	"response_code":       200,
	"response_bytes":      1234,
	"user_id":             1234,
}

func LogAndAssertJSON(t *testing.T, log func(*logrus.Logger), assertions func(fields logrus.Fields)) {
	var buffer bytes.Buffer
	var fields logrus.Fields

	logger := logrus.New()
	logger.Out = &buffer
	logger.Formatter = new(logrus.JSONFormatter)

	log(logger)

	err := json.Unmarshal(buffer.Bytes(), &fields)
	require.Nil(t, err)

	assertions(fields)
}

func TestNewDefaultLogger(t *testing.T) {
	userId := "test123"
	testAppLogger := NewDefaultLogger(&userId)
	assert.NotNil(t, testAppLogger)
	assert.NotNil(t, testAppLogger.logging)
	assert.Equal(t, *testAppLogger.userId, "test123")
	assert.Equal(t, len(testAppLogger.reqId), 36)
}

func TestNewDefaultWithNewReqId(t *testing.T) {
	userId := "test123"
	reqId := "1237-12321-543252-2423-123"
	testLogger := NewDefaultWithNewReqId(reqId, &userId)
	assert.NotNil(t, testLogger)
	assert.NotNil(t, testLogger.logging)
	assert.Equal(t, *testLogger.userId, "test123")
	assert.Equal(t, testLogger.reqId, "1237-12321-543252-2423-123")
}

func TestNewRequestLogger(t *testing.T) {
	testLogger := NewRequestLogger()
	assert.NotNil(t, testLogger)
	assert.NotNil(t, testLogger.logging)
	assert.Equal(t, len(testLogger.reqId), 36)
}

func TestNewRequestLoggerWithNewReqId(t *testing.T) {
	reqId := "1237-12321-543252-2423-123"
	testLogger := NewRequestLoggerWithNewReqId(reqId)
	assert.NotNil(t, testLogger)
	assert.NotNil(t, testLogger.logging)
	assert.Equal(t, testLogger.reqId, "1237-12321-543252-2423-123")
}

func TestgetReqFields(t *testing.T) {
	input := map[string]interface{}{
		"tag":                 "tag_value",
		"message":             "ok",
		"duration":            4.902,
		"request_ip":          "36.71.235.223",
		"request_remote_addr": "130.211.2.191:62695",
		"request_uri":         "https://erp.hoge.com/users/1234",
		"request_pattern":     "users/:id",
		"request_user_agent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36",
		"response_code":       200,
		"response_bytes":      1234,
		"user_id":             1234,
	}
	logFields := getReqFields("123-123-455-1134", input)
	assert.NotNil(t, logFields)
	assert.Equal(t, logFields["tag"], "tag_value")
	assert.Equal(t, logFields["message"], "ok")
	assert.Equal(t, logFields["duration"], 4.902)
	assert.Equal(t, logFields["request_ip"], "36.71.235.223")
	assert.Equal(t, logFields["request_remote_addr"], "130.211.2.191:62695")
	assert.Equal(t, logFields["request_uri"], "https://erp.hoge.com/users/1234")
	assert.Equal(t, logFields["request_pattern"], "users/:id")
	assert.Equal(t, logFields["request_user_agent"], "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	assert.Equal(t, logFields["response_code"], 200)
	assert.Equal(t, logFields["response_bytes"], 1234)
	assert.Equal(t, logFields["user_id"], 1234)

}

func TestgetAppFields(t *testing.T) {
	req_id := "123-123-455-1134"
	tag := "tag_value"
	user_id := "testuser"
	input := getAppFields(req_id, tag, &user_id)
	assert.NotNil(t, input)
}

func TestReqInfo(t *testing.T) {
	loggy := NewRequestLogger()
	var buf bytes.Buffer
	loggy.logging.SetOutput(&buf)
	defer func() {
		loggy.logging.SetOutput(os.Stderr)
	}()
	loggy.Info(Input)
	t.Log(buf.String())
	assert.Equal(t, len(buf.String()), len(` {"actor_id":1234,"duration":4.902,"level":"info","msg":"ok","request_id":"66558efb-a4f1-4d33-b44c-55c7c2d900c5","request_ip":"36.71.235.223","request_pattern":"users/:id","request_remote_addr":"130.211.2.191:62695","request_uri":"https://erp.hoge.com/users/1234","request_user_agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36","response_bytes":1234,"response_code":200,"tag":"tag_value","time":"2021-02-26T09:58:43+05:30"}`))
}

func TestReqDebug(t *testing.T) {
	loggy := NewRequestLogger()
	var buf bytes.Buffer
	loggy.logging.SetOutput(&buf)
	defer func() {
		loggy.logging.SetOutput(os.Stderr)
	}()
	loggy.Debug(Input)
	t.Log(buf.String())
	assert.NotNil(t, buf.String())
}

func TestReqTrace(t *testing.T) {
	loggy := NewRequestLogger()
	var buf bytes.Buffer
	loggy.logging.SetOutput(&buf)
	defer func() {
		loggy.logging.SetOutput(os.Stderr)
	}()
	loggy.Trace(Input)
	t.Log(buf.String())
	assert.NotNil(t, buf.String())
}

func TestReqWarn(t *testing.T) {
	loggy := NewRequestLogger()
	var buf bytes.Buffer
	loggy.logging.SetOutput(&buf)
	defer func() {
		loggy.logging.SetOutput(os.Stderr)
	}()
	loggy.Warn(Input)
	t.Log(buf.String())
	assert.Equal(t, len(buf.String()), len(` {"actor_id":1234,"duration":4.902,"level":"warning","msg":"ok","request_id":"439cebce-4a79-466f-af39-61dd6f956642","request_ip":"36.71.235.223","request_pattern":"users/:id","request_remote_addr":"130.211.2.191:62695","request_uri":"https://erp.hoge.com/users/1234","request_user_agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36","response_bytes":1234,"response_code":200,"tag":"tag_value","time":"2021-02-26T10:04:04+05:30"}`))
}

func TestReqError(t *testing.T) {
	loggy := NewRequestLogger()
	var buf bytes.Buffer
	loggy.logging.SetOutput(&buf)
	defer func() {
		loggy.logging.SetOutput(os.Stderr)
	}()
	loggy.Error(Input)
	t.Log(buf.String())
	assert.Equal(t, len(buf.String()), len(` {"actor_id":1234,"duration":4.902,"level":"error","msg":"ok","request_id":"439cebce-4a79-466f-af39-61dd6f956642","request_ip":"36.71.235.223","request_pattern":"users/:id","request_remote_addr":"130.211.2.191:62695","request_uri":"https://erp.hoge.com/users/1234","request_user_agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36","response_bytes":1234,"response_code":200,"tag":"tag_value","time":"2021-02-26T10:04:04+05:30"}`))
}

func TestAppLoggerInfo(t *testing.T) {
	userId := testUserId
	loggy := NewDefaultLogger(&userId)
	var buf bytes.Buffer
	loggy.logging.SetOutput(&buf)
	defer func() {
		loggy.logging.SetOutput(os.Stderr)
	}()
	loggy.Info(TAG, "log message")
	t.Log(buf.String())
	assert.Equal(t, len(buf.String()), len(` "actor_id":"1234","level":"info","msg":"log message","request_id":"05aa15fb-c85a-46b4-97fe-f5298deb6c7a","tag":"Log Test","time":"2021-07-05T16:54:10+05:30"} `))
}

func TestAppLoggerDebug(t *testing.T) {
	userId := testUserId
	loggy := NewDefaultLogger(&userId)
	var buf bytes.Buffer
	loggy.logging.SetOutput(&buf)
	defer func() {
		loggy.logging.SetOutput(os.Stderr)
	}()
	loggy.Debug(TAG, Input)
	t.Log(buf.String())
	assert.NotNil(t, buf.String())
}

func TestAppLoggerTrace(t *testing.T) {
	userId := testUserId
	loggy := NewDefaultLogger(&userId)
	var buf bytes.Buffer
	loggy.logging.SetOutput(&buf)
	defer func() {
		loggy.logging.SetOutput(os.Stderr)
	}()
	loggy.Trace(TAG, Input)
	t.Log(buf.String())
	assert.NotNil(t, buf.String())
}

func TestAppLoggerWarn(t *testing.T) {
	userId := testUserId
	loggy := NewDefaultLogger(&userId)
	var buf bytes.Buffer
	loggy.logging.SetOutput(&buf)
	defer func() {
		loggy.logging.SetOutput(os.Stderr)
	}()
	loggy.Warn(TAG, "log message")
	t.Log(buf.String())
	assert.Equal(t, len(buf.String()), len(` {"actor_id":"1234","level":"warning","msg":"log message","request_id":"2903dac6-4880-4d66-9186-977ac3950ebd","tag":"Log Test","time":"2021-07-05T17:01:18+05:30"}`))
}

func TestAppLoggerError(t *testing.T) {
	userId := testUserId
	loggy := NewDefaultLogger(&userId)
	var buf bytes.Buffer
	loggy.logging.SetOutput(&buf)
	defer func() {
		loggy.logging.SetOutput(os.Stderr)
	}()
	loggy.Error(TAG, "log message")
	t.Log(buf.String())
	assert.Equal(t, len(buf.String()), len(` {"actor_id":"1234","level":"error","msg":"log message","request_id":"b42ac180-0202-4a46-9e40-a7bac16e316f","tag":"Log Test","time":"2021-07-05T17:04:02+05:30"}`))
}
