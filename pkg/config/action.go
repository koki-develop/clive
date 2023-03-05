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

var (
	_ Action = (*TypeAction)(nil)
	_ Action = (*KeyAction)(nil)
	_ Action = (*SleepAction)(nil)
	_ Action = (*PauseAction)(nil)
	_ Action = (*CtrlAction)(nil)
	_ Action = (*ScreenshotAction)(nil)
)

type TypeAction struct {
	Type  string `mapstructure:"type"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

func (action *TypeAction) String() string {
	return fmt.Sprintf("Type: %s", action.Type)
}

var typeActionValidFields = []string{"type", "count", "speed"}

type KeyAction struct {
	Key   string `mapstructure:"key"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

var keyActionValidFields = []string{"key", "count", "speed"}

func (action *KeyAction) String() string {
	return fmt.Sprintf("Key: %s", action.Key)
}

type SleepAction struct {
	Sleep int `mapstructure:"sleep"`
}

var sleepActionValidFields = []string{"sleep"}

func (action *SleepAction) String() string {
	return fmt.Sprintf("Sleep: %dms", action.Sleep)
}

type PauseAction struct{}

var pauseActionValidFields = []string{"pause"}

func (*PauseAction) String() string {
	return "Pause: Press enter to continue"
}

type CtrlAction struct {
	Ctrl  string `mapstructure:"ctrl"`
	Count int    `mapstructure:"count"`
	Speed int    `mapstructure:"speed"`
}

var ctrlActionValidFields = []string{"ctrl", "count", "speed"}

func (action *CtrlAction) String() string {
	return fmt.Sprintf("Ctrl+%s", action.Ctrl)
}

type ScreenshotAction struct {
	File string `mapstructure:"file"`
	Dir  string `mapstructure:"dir"`
}

var screenshotActionValidFields = []string{"screenshot", "file", "dir"}

func (action *ScreenshotAction) String() string {
	return "Take a screenshot"
}

var parseFuncMap = map[string]func(*Settings, map[string]interface{}) (Action, error){
	"type":       parseTypeAction,
	"key":        parseKeyAction,
	"ctrl":       parseCtrlAction,
	"sleep":      parseSleepAction,
	"pause":      parsePauseAction,
	"screenshot": parseScreenshotAction,
}

func ParseAction(stgs *Settings, v interface{}) (Action, error) {
	switch v := v.(type) {
	case string:
		switch v {
		case "pause":
			return &PauseAction{}, nil
		case "screenshot":
			return &ScreenshotAction{}, nil
		}
	case map[string]interface{}:
		for _, k := range []string{"type", "key", "ctrl", "sleep", "pause", "screenshot"} {
			if _, ok := v[k]; ok {
				return parseFuncMap[k](stgs, v)
			}
		}
	}

	return nil, NewErrInvalidAction(v)
}

func parseTypeAction(stgs *Settings, m map[string]interface{}) (Action, error) {
	if err := validateActionFields(m, typeActionValidFields); err != nil {
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

func parseKeyAction(settings *Settings, m map[string]interface{}) (Action, error) {
	if err := validateActionFields(m, keyActionValidFields); err != nil {
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

func parseSleepAction(settings *Settings, m map[string]interface{}) (Action, error) {
	if err := validateActionFields(m, sleepActionValidFields); err != nil {
		return nil, err
	}

	var action SleepAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, errors.WithMessage(NewErrInvalidAction(m), err.Error())
	}

	return &action, nil
}

func parsePauseAction(settings *Settings, m map[string]interface{}) (Action, error) {
	if err := validateActionFields(m, pauseActionValidFields); err != nil {
		return nil, err
	}

	return &PauseAction{}, nil
}

func parseCtrlAction(settings *Settings, m map[string]interface{}) (Action, error) {
	if err := validateActionFields(m, ctrlActionValidFields); err != nil {
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

func parseScreenshotAction(settings *Settings, m map[string]interface{}) (Action, error) {
	if err := validateActionFields(m, screenshotActionValidFields); err != nil {
		return nil, err
	}

	return &ScreenshotAction{}, nil
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
