package config

import (
	"fmt"
	"sort"

	"github.com/koki-develop/clive/pkg/util"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type Action interface {
	String() string
}

type TypeAction struct {
	Type  string `mapstructure:"type"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

func (action *TypeAction) String() string {
	return fmt.Sprintf("Type: %s", util.TruncateString(action.Type, 37))
}

var TypeActionValidFields = []string{"type", "count", "speed"}

type KeyAction struct {
	Key   string `mapstructure:"key"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

var KeyActionValidFields = []string{"key", "count", "speed"}

func (action *KeyAction) String() string {
	return fmt.Sprintf("Key: %s", action.Key)
}

type SleepAction struct {
	Sleep int `mapstructure:"sleep"`
}

var SleepActionValidFields = []string{"sleep"}

func (action *SleepAction) String() string {
	return fmt.Sprintf("Sleep: %dms", action.Sleep)
}

type PauseAction struct{}

var PauseActionValidFields = []string{"pause"}

func (action *PauseAction) String() string {
	return "Pause: Press enter to continue"
}

type CtrlAction struct {
	Ctrl  string `mapstructure:"ctrl"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

var CtrlActionValidFields = []string{"ctrl", "count", "speed"}

func (action *CtrlAction) String() string {
	return fmt.Sprintf("Ctrl+%s", action.Ctrl)
}

func ParseAction(stgs *Settings, v interface{}) (Action, error) {
	switch v := v.(type) {
	case string:
		switch v {
		case "pause":
			return &PauseAction{}, nil
		}
	case map[string]interface{}:
		if _, ok := v["type"]; ok {
			return parseTypeAction(stgs, v)
		}
		if _, ok := v["key"]; ok {
			return parseKeyAction(stgs, v)
		}
		if _, ok := v["ctrl"]; ok {
			return parseCtrlAction(stgs, v)
		}
		if _, ok := v["sleep"]; ok {
			return parseSleepAction(stgs, v)
		}
		if _, ok := v["pause"]; ok {
			return parsePauseAction(stgs, v)
		}
	}

	return nil, NewErrInvalidAction(v)
}

func parseTypeAction(stgs *Settings, m map[string]interface{}) (*TypeAction, error) {
	if err := validateActionFields(m, TypeActionValidFields); err != nil {
		return nil, err
	}

	action := TypeAction{
		Count: 1,
		Speed: stgs.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, errors.WithMessage(NewErrInvalidAction(m), err.Error())
	}

	return &action, nil
}

func parseKeyAction(settings *Settings, m map[string]interface{}) (*KeyAction, error) {
	if err := validateActionFields(m, KeyActionValidFields); err != nil {
		return nil, err
	}

	action := KeyAction{
		Count: 1,
		Speed: settings.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, errors.WithMessage(NewErrInvalidAction(m), err.Error())
	}

	if _, ok := SpecialKeyMap[action.Key]; !ok {
		validKeys := []string{}
		for k := range SpecialKeyMap {
			validKeys = append(validKeys, k)
		}
		sort.Strings(validKeys)
		return nil, errors.WithMessagef(NewErrInvalidAction(m), "valid keys %s", validKeys)
	}

	return &action, nil
}

func parseSleepAction(settings *Settings, m map[string]interface{}) (*SleepAction, error) {
	if err := validateActionFields(m, SleepActionValidFields); err != nil {
		return nil, err
	}

	var action SleepAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, errors.WithMessage(NewErrInvalidAction(m), err.Error())
	}

	return &action, nil
}

func parsePauseAction(settings *Settings, m map[string]interface{}) (*PauseAction, error) {
	if err := validateActionFields(m, PauseActionValidFields); err != nil {
		return nil, err
	}

	var action PauseAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, errors.WithMessage(NewErrInvalidAction(m), err.Error())
	}

	return &action, nil
}

func parseCtrlAction(settings *Settings, m map[string]interface{}) (*CtrlAction, error) {
	if err := validateActionFields(m, CtrlActionValidFields); err != nil {
		return nil, err
	}

	action := CtrlAction{
		Count: 1,
		Speed: settings.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, errors.WithMessage(NewErrInvalidAction(m), err.Error())
	}

	return &action, nil
}

func validateActionFields(m map[string]interface{}, valid []string) error {
	invalidFields := []string{}
	for k := range m {
		if !util.Contains(valid, k) {
			invalidFields = append(invalidFields, k)
		}
	}
	if len(invalidFields) > 0 {
		return errors.WithMessagef(NewErrInvalidAction(m), "unknown fields %s", invalidFields)
	}

	return nil
}
