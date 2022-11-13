package config

import (
	"fmt"

	"github.com/koki-develop/clive/pkg/util"
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
