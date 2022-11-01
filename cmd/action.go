package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod/lib/input"
)

type action interface {
	String() string
}

type typeAction struct {
	Type string
	Time time.Duration
}

func (action *typeAction) String() string {
	return fmt.Sprintf("Type: %s", truncateString(action.Type, 37))
}

type keyAction struct {
	Key   input.Key
	Count int
	Time  time.Duration
}

func (action *keyAction) String() string {
	txt := ""
	for k, v := range specialkeymap {
		if v == action.Key {
			txt = k
		}
	}

	return fmt.Sprintf("Key: %s", txt)
}

type sleepAction struct {
	Time time.Duration
}

func (action *sleepAction) String() string {
	return fmt.Sprintf("Sleep: %dms", action.Time.Milliseconds())
}

type pauseAction struct{}

func (action *pauseAction) String() string {
	return "Press enter to continue"
}

type ctrlAction struct {
	Ctrl  string
	Count int
	Time  time.Duration
}

func (action *ctrlAction) String() string {
	return fmt.Sprintf("Ctrl+%s", action.Ctrl)
}

func parseAction(v interface{}) (action, error) {
	switch v := v.(type) {
	case string:
		switch v {
		case "pause":
			return &pauseAction{}, nil
		}
	case map[string]interface{}:
		if _, ok := v["pause"]; ok {
			return &pauseAction{}, nil
		}
		if _, ok := v["type"]; ok {
			return parseTypeAction(v)
		}
		if _, ok := v["key"]; ok {
			return parseKeyAction(v)
		}
		if _, ok := v["sleep"]; ok {
			return parseSleepAction(v)
		}
		if _, ok := v["ctrl"]; ok {
			return parseCtrlAction(v)
		}
	}

	return nil, fmt.Errorf("invalid action: %#v", v)
}

func parseTypeAction(m map[string]interface{}) (*typeAction, error) {
	for k := range m {
		switch k {
		case "type", "time":
		default:
			return nil, fmt.Errorf("invalid action: %#v", m)
		}
	}

	var t time.Duration
	if v, ok := m["time"]; ok {
		t = time.Duration(v.(int)) * time.Millisecond
	} else {
		t = 10 * time.Millisecond
	}

	return &typeAction{
		Type: m["type"].(string),
		Time: t,
	}, nil
}

func parseKeyAction(m map[string]interface{}) (*keyAction, error) {
	for k := range m {
		switch k {
		case "key", "count", "time":
		default:
			return nil, fmt.Errorf("invalid action: %#v", m)
		}
	}

	c := 1
	if v, ok := m["count"]; ok {
		c = v.(int)
	}

	t := 10 * time.Millisecond
	if v, ok := m["time"]; ok {
		t = time.Duration(v.(int)) * time.Millisecond
	}

	k, ok := specialkeymap[strings.ToLower(m["key"].(string))]
	if !ok {
		return nil, fmt.Errorf("invalid action: %#v", m)
	}

	return &keyAction{
		Key:   k,
		Count: c,
		Time:  t,
	}, nil
}

func parseSleepAction(m map[string]interface{}) (*sleepAction, error) {
	for k := range m {
		switch k {
		case "sleep":
		default:
			return nil, fmt.Errorf("invalid action: %#v", m)
		}
	}

	return &sleepAction{
		Time: time.Duration(m["sleep"].(int)) * time.Millisecond,
	}, nil
}

func parseCtrlAction(m map[string]interface{}) (*ctrlAction, error) {
	for k := range m {
		switch k {
		case "ctrl", "time", "count":
		default:
			return nil, fmt.Errorf("invalid action: %#v", m)
		}
	}

	ctrl := m["ctrl"].(string)

	c := 1
	if v, ok := m["count"]; ok {
		c = v.(int)
	}

	t := 10 * time.Millisecond
	if v, ok := m["time"]; ok {
		t = time.Duration(v.(int)) * time.Millisecond
	}

	return &ctrlAction{
		Ctrl:  ctrl,
		Count: c,
		Time:  t,
	}, nil
}
