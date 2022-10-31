package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod/lib/input"
	"gopkg.in/yaml.v3"
)

type action interface {
	IsAction()
}

type typeAction struct {
	Type string
	Time time.Duration
}

func (action *typeAction) IsAction() {}

type keyAction struct {
	Key   input.Key
	Count int
	Time  time.Duration
}

func (action *keyAction) IsAction() {}

type sleepAction struct {
	Time time.Duration
}

func (action sleepAction) IsAction() {}

type pauseAction struct{}

func (action pauseAction) IsAction() {}

func loadConfig(p string) (*config, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var y configYaml
	if err := yaml.NewDecoder(f).Decode(&y); err != nil {
		return nil, err
	}

	var actions []action
	for _, a := range y.Actions {
		action, err := parseAction(a)
		if err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}

	return &config{
		Actions: actions,
	}, nil
}

func parseAction(v interface{}) (action, error) {
	switch v := v.(type) {
	case string:
		switch v {
		case "pause":
			return &pauseAction{}, nil
		}
	case map[string]interface{}:
		for k := range v {
			switch k {
			case "pause":
				return &pauseAction{}, nil
			case "type":
				return parseTypeAction(v)
			case "key":
				return parseKeyAction(v)
			case "sleep":
				return parseSleepAction(v)
			}
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

	k, ok := map[string]input.Key{
		"enter": input.Enter,
	}[m["key"].(string)]
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
