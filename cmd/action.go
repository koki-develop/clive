package cmd

import (
	"fmt"
	"sort"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type action interface {
	String() string
}

type typeAction struct {
	Type  string `mapstructure:"type"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

var typeActionValidFields = []string{"type", "count", "speed"}

type keyAction struct {
	Key   string `mapstructure:"key"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

var keyActionValidFields = []string{"key", "count", "speed"}

type sleepAction struct {
	Sleep int `mapstructure:"sleep"`
}

var sleepActionValidFields = []string{"sleep"}

type pauseAction struct{}

var pauseActionValidFields = []string{"pause"}

type ctrlAction struct {
	Ctrl  string `mapstructure:"ctrl"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

var ctrlActionValidFields = []string{"ctrl", "count", "speed"}

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
	if err := validateActionFields(m, typeActionValidFields); err != nil {
		return nil, err
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
	if err := validateActionFields(m, keyActionValidFields); err != nil {
		return nil, err
	}

	action := keyAction{
		Count: 1,
		Speed: settings.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	if _, ok := specialkeymap[action.Key]; !ok {
		validKeys := []string{}
		for k := range specialkeymap {
			validKeys = append(validKeys, k)
		}
		sort.Strings(validKeys)
		return nil, errors.WithMessagef(fmt.Errorf("valid keys %s", validKeys), newInvalidActionError(m).Error())
	}

	return &action, nil
}

func parseSleepAction(settings *settings, m map[string]interface{}) (*sleepAction, error) {
	if err := validateActionFields(m, sleepActionValidFields); err != nil {
		return nil, err
	}

	var action sleepAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parsePauseAction(settings *settings, m map[string]interface{}) (*pauseAction, error) {
	if err := validateActionFields(m, pauseActionValidFields); err != nil {
		return nil, err
	}

	var action pauseAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseCtrlAction(settings *settings, m map[string]interface{}) (*ctrlAction, error) {
	if err := validateActionFields(m, ctrlActionValidFields); err != nil {
		return nil, err
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

func validateFields(m map[string]interface{}, validFields []string) error {
	invalidFields := []string{}
	for k := range m {
		if !contains(validFields, k) {
			invalidFields = append(invalidFields, k)
		}
	}

	if len(invalidFields) > 0 {
		return fmt.Errorf("unknown fields %s", invalidFields)
	}

	return nil
}

func validateActionFields(m map[string]interface{}, validFields []string) error {
	if err := validateFields(m, validFields); err != nil {
		return errors.Wrap(err, newInvalidActionError(m).Error())
	}

	return nil
}
