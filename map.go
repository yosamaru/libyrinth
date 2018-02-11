package labyrinth

import (
	"encoding/json"
	"io/ioutil"
)

// 创建不同等级对应的地图
type levelMap struct {
	Type string
	Data map[string]interface{}
}

type EntityParser func(map[string]interface{}) Drawable

func parseRectangle(data map[string]interface{}) *Rectangle {
	return nil
}

func parseText(data map[string]interface{}) *Text {
	return nil
}

func parseEntity(data map[string]interface{}) (*Entity, error) {
	return nil, nil
}

func LoadLevelFromMap(jsonMap string, parsers map[string]EntityParser, l Level) error {
	return nil
}