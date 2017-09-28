package log

import (
	"encoding/json"

	"github.com/cybozu-go/kkok"
	"github.com/cybozu-go/log"
	"github.com/pkg/errors"
)

const (
	transportType   = "log"
	defaultSeverity = log.LvInfo
)

type transport struct {
	label    string
	all      bool
	severity int
}

func newTransport() *transport {
	return &transport{
		severity: defaultSeverity,
	}
}

func (t *transport) String() string {
	if len(t.label) > 0 {
		return t.label
	}
	return transportType
}

func (t *transport) Params() kkok.PluginParams {
	m := map[string]interface{}{
		"severity": t.severity,
	}

	if len(t.label) > 0 {
		m["label"] = t.label
	}

	m["all"] = t.all

	return kkok.PluginParams{
		Type:   transportType,
		Params: m,
	}
}

func (t *transport) log(data []byte) error {
	defaultLogger := log.DefaultLogger()
	if defaultLogger == nil {
		return errors.New("log transport: defaultLogger is nil")
	}
	return defaultLogger.Log(t.severity, "[kkok] logged alerts", map[string]interface{}{"alerts": string(data)})
}

func (t *transport) Deliver(alerts []*kkok.Alert) error {
	if t.all {
		data, err := json.Marshal(alerts)
		if err != nil {
			return errors.Wrap(err, transportType)
		}
		err = t.log(data)
		if err != nil {
			return errors.Wrap(err, transportType)
		}
		return nil
	}

	for _, a := range alerts {
		data, err := json.Marshal(a)
		if err != nil {
			return errors.Wrap(err, transportType)
		}
		err = t.log(data)
		if err != nil {
			return errors.Wrap(err, transportType)
		}
	}

	return nil
}
