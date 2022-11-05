package cmd

import (
	"fmt"
	"strings"

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

var typeActionValidKeys = []string{"type", "count", "speed"}

type keyAction struct {
	Key   string `mapstructure:"key"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

var keyActionValidKeys = []string{"key", "count", "speed"}

type sleepAction struct {
	Sleep int `mapstructure:"sleep"`
}

var sleepActionValidKeys = []string{"sleep"}

type pauseAction struct{}

var pauseActionValidKeys = []string{"pause"}

type ctrlAction struct {
	Ctrl  string `mapstructure:"ctrl"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

var ctrlActionValidKeys = []string{"ctrl", "count", "speed"}

func (action *typeAction) String() string {
	return fmt.Sprintf("Type: %s", truncateString(action.Type, 37))
}

func (action *keyAction) String() string {
	return fmt.Sprintf("Key: %s", action.Key)
}

func (action *sleepAction) String() string {
	return fmt.Sprintf("Sleep: %dms", action.Sleep)
}

func (action *pauseAction) String() string {
	return "Pause: Press enter to continue"
}

func (action *ctrlAction) String() string {
	return fmt.Sprintf("Ctrl+%s", action.Ctrl)
}

func parseAction(settings *settings, v interface{}) (action, error) {
	switch v := v.(type) {
	case string:
		switch v {
		case "pause":
			return &pauseAction{}, nil
		}
	case map[string]interface{}:
		if _, ok := v["type"]; ok {
			return parseTypeAction(settings, v)
		}
		if _, ok := v["key"]; ok {
			return parseKeyAction(settings, v)
		}
		if _, ok := v["ctrl"]; ok {
			return parseCtrlAction(settings, v)
		}
		if _, ok := v["sleep"]; ok {
			return parseSleepAction(settings, v)
		}
		if _, ok := v["pause"]; ok {
			return parsePauseAction(settings, v)
		}
	}

	return nil, newInvalidActionError(v)
}

func parseTypeAction(settings *settings, m map[string]interface{}) (*typeAction, error) {
	if !validateKeys(m, typeActionValidKeys) {
		return nil, newInvalidActionError(m)
	}

	action := typeAction{
		Count: 1,
		Speed: settings.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseKeyAction(settings *settings, m map[string]interface{}) (*keyAction, error) {
	if !validateKeys(m, keyActionValidKeys) {
		return nil, newInvalidActionError(m)
	}

	action := keyAction{
		Count: 1,
		Speed: settings.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	if _, ok := specialkeymap[strings.ToLower(action.Key)]; !ok {
		return nil, newInvalidActionError(m)
	}

	return &action, nil
}

func parseSleepAction(settings *settings, m map[string]interface{}) (*sleepAction, error) {
	if !validateKeys(m, sleepActionValidKeys) {
		return nil, newInvalidActionError(m)
	}

	var action sleepAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parsePauseAction(settings *settings, m map[string]interface{}) (*pauseAction, error) {
	if !validateKeys(m, pauseActionValidKeys) {
		return nil, newInvalidActionError(m)
	}

	var action pauseAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseCtrlAction(settings *settings, m map[string]interface{}) (*ctrlAction, error) {
	if !validateKeys(m, ctrlActionValidKeys) {
		return nil, newInvalidActionError(m)
	}

	action := ctrlAction{
		Count: 1,
		Speed: settings.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func validateKeys(m map[string]interface{}, validKeys []string) bool {
	for k := range m {
		if !contains(validKeys, k) {
			return false
		}
	}
	return true
}
