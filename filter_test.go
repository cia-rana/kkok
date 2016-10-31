package kkok

import (
	"os/exec"
	"testing"
	"time"
)

func newBaseFilter(id string, dynamic bool, params map[string]interface{}) (*BaseFilter, error) {
	b := new(BaseFilter)
	err := b.Init(id, dynamic, params)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func testBaseFilterAll(t *testing.T) {
	t.Parallel()
	params := map[string]interface{}{
		"disabled": true,
		"all":      true,
		"if":       "alerts.length > 1",
	}

	b, err := newBaseFilter("base", false, params)
	if err != nil {
		t.Fatal(err)
	}

	if b.Dynamic() {
		t.Error(`b.Dynamic()`)
	}

	if b.ID() != "base" {
		t.Error(`b.ID() != "base"`)
	}
	if !b.Disabled() {
		t.Error(`!b.Disabled()`)
	}
	b.Enable(true)
	if b.Disabled() {
		t.Error(`b.Disabled()`)
	}
	if !b.All() {
		t.Error(`!b.All()`)
	}

	ok, err := b.EvalAllAlerts(nil)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Error("condition shoult not be met")
	}

	ok, err = b.EvalAllAlerts([]*Alert{&Alert{}, &Alert{}})
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Error("condition shoult be met")
	}
}

func testBaseFilterOne(t *testing.T) {
	t.Parallel()
	params := map[string]interface{}{
		"all": false,
		"if":  "alert.From == 'hoge'",
	}

	b, err := newBaseFilter("base", true, params)
	if err != nil {
		t.Fatal(err)
	}

	if !b.Dynamic() {
		t.Error(`!b.Dynamic()`)
	}

	if b.Disabled() {
		t.Error(`b.Disabled()`)
	}
	if b.All() {
		t.Error(`b.All()`)
	}

	ok, err := b.EvalAlert(&Alert{})
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Error("condition shoult not be met")
	}

	ok, err = b.EvalAlert(&Alert{From: "hoge"})
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Error("condition shoult be met")
	}
}

func testBaseFilterParseError(t *testing.T) {
	t.Parallel()
	params := map[string]interface{}{
		"all": false,
		"if":  "alert.From =",
	}

	_, err := newBaseFilter("id", true, params)
	if err == nil {
		t.Fatal("if must cause a parse error")
	}
	t.Log(err)
}

func testBaseFilterCommand(t *testing.T) {
	jq, err := exec.LookPath("jq")
	if err != nil {
		t.Skip(err)
	}

	t.Parallel()
	params := map[string]interface{}{
		"all": false,
		"if":  []interface{}{jq, "-e", `.From == "hoge"`},
	}

	b, err := newBaseFilter("base", false, params)
	if err != nil {
		t.Fatal(err)
	}

	ok, err := b.EvalAlert(&Alert{})
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Error("condition shoult not be met")
	}

	ok, err = b.EvalAlert(&Alert{From: "hoge"})
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Error("condition shoult be met")
	}
}

func testBaseFilterExpire(t *testing.T) {
	t.Parallel()

	params := map[string]interface{}{
		"expire": "hoge",
	}

	_, err := newBaseFilter("id", true, params)
	if err == nil {
		t.Fatal("expire must cause a parse error")
	}
	t.Log(err)

	now := time.Now().UTC()

	params["expire"] = now.Add(-1 * time.Hour).Format(time.RFC3339)
	f, err := newBaseFilter("id", false, params)
	if err != nil {
		t.Fatal(err)
	}
	if !f.expire.IsZero() {
		t.Error(`!f.expire.IsZero()`)
	}

	f, err = newBaseFilter("id", true, params)
	if err != nil {
		t.Fatal(err)
	}
	if f.expire.IsZero() {
		t.Error(`f.expire.IsZero()`)
	}
	if !f.Expired() {
		t.Error(`!f.Expired()`)
	}

	params["expire"] = now.Add(1 * time.Hour).Format(time.RFC3339)
	f, err = newBaseFilter("id", true, params)
	if err != nil {
		t.Fatal(err)
	}
	if f.expire.IsZero() {
		t.Error(`f.expire.IsZero()`)
	}
	if f.Expired() {
		t.Error(`f.Expired()`)
	}
}

func testBaseFilterAddParams(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	params := map[string]interface{}{
		"disabled": true,
		"all":      true,
		"if":       "alerts.length > 1",
		"expire":   now.Format(time.RFC3339Nano),
	}

	f, err := newBaseFilter("id", true, params)
	if err != nil {
		t.Fatal(err)
	}

	m := make(map[string]interface{})
	f.AddParams(m)

	if !m["disabled"].(bool) {
		t.Error(`!m["disabled"].(bool)`)
	}

	if !m["all"].(bool) {
		t.Error(`!m["all"].(bool)`)
	}

	if m["if"].(string) != "alerts.length > 1" {
		t.Error(`m["if"].(string) != "alerts.length > 1"`)
	}

	if !now.Equal(m["expire"].(time.Time)) {
		t.Error(`!now.Equal(m["expire"].(time.Time))`)
	}
}

func TestBaseFilter(t *testing.T) {
	t.Run("All", testBaseFilterAll)
	t.Run("One", testBaseFilterOne)
	t.Run("ParseError", testBaseFilterParseError)
	t.Run("Command", testBaseFilterCommand)
	t.Run("Expire", testBaseFilterExpire)
	t.Run("AddParams", testBaseFilterAddParams)
}
