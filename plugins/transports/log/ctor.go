package log

import (
	"github.com/cybozu-go/kkok"
	"github.com/cybozu-go/kkok/util"
	"github.com/cybozu-go/log"
	"github.com/pkg/errors"
)

var severityMap = map[string]int{
	"critical": log.LvCritical,
	"error":    log.LvError,
	"warn":     log.LvWarn,
	"info":     log.LvInfo,
	"debug":    log.LvDebug,
}

func ctor(params map[string]interface{}) (kkok.Transport, error) {
	tr := newTransport()

	label, err := util.GetString("label", params)
	if err != nil {
		return nil, errors.Wrap(err, "log transport: label")
	}
	tr.label = label

	all, err := util.GetBool("all", params)
	if err != nil && !util.IsNotFound(err) {
		return nil, errors.Wrap(err, "log transport: all")
	}
	tr.all = all

	severity, err := util.GetString("severity", params)
	if err != nil && !util.IsNotFound(err) {
		return nil, errors.Wrap(err, "log transport: severity")
	}
	if v, ok := severityMap[severity]; ok {
		tr.severity = v
	}
	return tr, nil
}

func init() {
	kkok.RegisterTransport(transportType, ctor)
}
