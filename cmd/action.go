package cmd

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type action interface {
	String() string
}

type typeAction struct {
	Type  string `mapstructure:"type"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

type keyAction struct {
	Key   string `mapstructure:"key"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

type sleepAction struct {
	Time int `mapstructure:"time"`
}

type pauseAction struct{}

type ctrlAction struct {
	Ctrl  string `mapstructure:"ctrl"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

func (action *typeAction) String() string {
	return fmt.Sprintf("Type: %s", truncateString(action.Type, 37))
}

func (action *keyAction) String() string {
	return fmt.Sprintf("Key: %s", action.Key)
}

func (action *sleepAction) String() string {
	return fmt.Sprintf("Sleep: %dms", action.Time)
}

func (action *pauseAction) String() string {
	return "Press enter to continue"
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
			return parsePauseAction(v)
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
	if _, ok := m["speed"]; !ok {
		m["speed"] = 10
	}

	var action typeAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseKeyAction(m map[string]interface{}) (*keyAction, error) {
	if _, ok := m["count"]; !ok {
		m["count"] = 1
	}
	if _, ok := m["speed"]; !ok {
		m["speed"] = 10
	}

	var action keyAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseSleepAction(m map[string]interface{}) (*sleepAction, error) {
	var action sleepAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parsePauseAction(m map[string]interface{}) (*pauseAction, error) {
	var action pauseAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseCtrlAction(m map[string]interface{}) (*ctrlAction, error) {
	if _, ok := m["count"]; !ok {
		m["count"] = 1
	}
	if _, ok := m["speed"]; !ok {
		m["speed"] = 10
	}

	var action ctrlAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}
