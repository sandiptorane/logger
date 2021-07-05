package logger

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AppLogger struct {
	logging *logrus.Logger
	reqId   string
	userId  *string
}

type ReqLogger struct {
	logging *logrus.Logger
	reqId   string
}

func NewDefaultLogger(userId *string) *AppLogger {
	return &AppLogger{
		logging: logrus.New(),
		reqId:   uuid.NewString(),
		userId:  userId,
	}
}

func (logger *AppLogger) SetUserId(userId string) {
	logger.userId = &userId
}

func NewDefaultWithNewReqId(reqId string, userId *string) *AppLogger {
	return &AppLogger{
		logging: logrus.New(),
		reqId:   reqId,
		userId:  userId,
	}
}

func NewRequestLogger() *ReqLogger {
	return &ReqLogger{
		logging: logrus.New(),
		reqId:   uuid.NewString(),
	}
}

func NewRequestLoggerWithNewReqId(reqId string) *ReqLogger {
	return &ReqLogger{
		logging: logrus.New(),
		reqId:   reqId,
	}
}

func getReqFields(reqId string, input map[string]interface{}) map[string]interface{} {
	logrusFields := map[string]interface{}{
		"request_id":          reqId,
		"tag":                 input["tag"],
		"duration":            input["duration"],
		"request_ip":          input["request_ip"],
		"request_remote_addr": input["request_remote_addr"],
		"request_uri":         input["request_uri"],
		"request_pattern":     input["request_pattern"],
		"request_user_agent":  input["request_user_agent"],
		"response_code":       input["response_code"],
		"response_bytes":      input["response_bytes"],
		"actor_id":            input["user_id"],
	}
	return logrusFields
}

func getAppFields(reqId, tag string, userId *string) map[string]interface{} {
	logrusFields := map[string]interface{}{
		"request_id": reqId,
		"tag":        tag,
		"actor_id":   userId,
	}
	return logrusFields
}

func (l *ReqLogger) Infof() {
	l.logging.Info("haha")
}
func (l *ReqLogger) Info(input map[string]interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getReqFields(l.reqId, input)
	l.logging.WithFields(k).Info(input["message"])

}

func (l *ReqLogger) Debug(input map[string]interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getReqFields(l.reqId, input)
	l.logging.WithFields(k).Debug(input["message"])
}

func (l *ReqLogger) Trace(input map[string]interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getReqFields(l.reqId, input)
	l.logging.WithFields(k).Trace(input["message"])
}

func (l *ReqLogger) Warn(input map[string]interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getReqFields(l.reqId, input)
	l.logging.WithFields(k).Warn(input["message"])
}

func (l *ReqLogger) Error(input map[string]interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getReqFields(l.reqId, input)
	l.logging.WithFields(k).Error(input["message"])
}

func (l *ReqLogger) Fatal(input map[string]interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getReqFields(l.reqId, input)
	l.logging.WithFields(k).Fatal(input["message"])
}

func (l *ReqLogger) Panic(input map[string]interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getReqFields(l.reqId, input)
	l.logging.WithFields(k).Panic(input["message"])
}

/////////AppLogger Methods/////

func (l *AppLogger) Info(tag string, message ...interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getAppFields(l.reqId, tag, l.userId)
	l.logging.WithFields(k).Info(message...)
}

func (l *AppLogger) Debug(tag string, message ...interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getAppFields(l.reqId, tag, l.userId)
	l.logging.WithFields(k).Debug(message...)
}

func (l *AppLogger) Trace(tag string, message ...interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getAppFields(l.reqId, tag, l.userId)
	l.logging.WithFields(k).Trace(message...)
}

func (l *AppLogger) Warn(tag string, message ...interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getAppFields(l.reqId, tag, l.userId)
	l.logging.WithFields(k).Warn(message...)
}

func (l *AppLogger) Error(tag string, message ...interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getAppFields(l.reqId, tag, l.userId)
	l.logging.WithFields(k).Error(message...)
}

func (l *AppLogger) Fatal(tag string, message ...interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getAppFields(l.reqId, tag, l.userId)
	l.logging.WithFields(k).Fatal(message...)
}

func (l *AppLogger) Panic(tag string, message ...interface{}) {
	l.logging.SetFormatter(&logrus.JSONFormatter{})
	k := getAppFields(l.reqId, tag, l.userId)
	l.logging.WithFields(k).Panic(message...)
}
