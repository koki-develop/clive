package cmd

import (
	"fmt"
	"sort"

	"github.com/koki-develop/clive/pkg/config"
	"github.com/koki-develop/clive/pkg/util"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

func parseAction(settings *config.Settings, v interface{}) (config.Action, error) {
	switch v := v.(type) {
	case string:
		switch v {
		case "pause":
			return &config.PauseAction{}, nil
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

func parseTypeAction(settings *config.Settings, m map[string]interface{}) (*config.TypeAction, error) {
	if err := validateActionFields(m, config.TypeActionValidFields); err != nil {
		return nil, err
	}

	action := config.TypeAction{
		Count: 1,
		Speed: settings.DefaultSpeed,
	}
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseKeyAction(settings *config.Settings, m map[string]interface{}) (*config.KeyAction, error) {
	if err := validateActionFields(m, config.KeyActionValidFields); err != nil {
		return nil, err
	}

	action := config.KeyAction{
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

func parseSleepAction(settings *config.Settings, m map[string]interface{}) (*config.SleepAction, error) {
	if err := validateActionFields(m, config.SleepActionValidFields); err != nil {
		return nil, err
	}

	var action config.SleepAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parsePauseAction(settings *config.Settings, m map[string]interface{}) (*config.PauseAction, error) {
	if err := validateActionFields(m, config.PauseActionValidFields); err != nil {
		return nil, err
	}

	var action config.PauseAction
	if err := mapstructure.Decode(m, &action); err != nil {
		return nil, err
	}

	return &action, nil
}

func parseCtrlAction(settings *config.Settings, m map[string]interface{}) (*config.CtrlAction, error) {
	if err := validateActionFields(m, config.CtrlActionValidFields); err != nil {
		return nil, err
	}

	action := config.CtrlAction{
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
		if !util.Contains(validFields, k) {
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
