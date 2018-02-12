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
	// TODO .(float)的作用
	return NewRectangle(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		int(data["width"].(float64)),
		int(data["height"].(float64)),
		Attr(data["color"].(float64)),
	)
}

func parseText(data map[string]interface{}) *Text {
	return NewText(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		data["text"].(string),
		Attr(data["fg"].(float64)),
		Attr(data["bg"].(float64)),
	)
}

func parseEntity(data map[string]interface{}) (*Entity, error) {
	filename := data["text"].(string)
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	e := NewEntityFromCanvas(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		CanvasFromString(string(text)),
	)
	bgfile := data["bg"].(string)
	if bgfile != "" {
		e.ApplyCanvas(BackgroundCanvasFromFile(bgfile))
	}
	fgfile := data["fg"].(string)
	if fgfile != "" {
		e.ApplyCanvas(ForegroundCanvasFromFile(fgfile))
	}
	return e, nil
}

func LoadLevelFromMap(jsonMap string, parsers map[string]EntityParser, l Level) error {
	parseMap := []levelMap{}
	if err := json.Unmarshal([]byte(jsonMap), &parseMap); err != nil {
		return err
	}
	for _, lm := range parseMap {
		var entity Drawable
		var err error
		switch lm.Type {
		case "Rectangle":
			entity = parseRectangle(lm.Data)
		case "Text":
			entity = parseText(lm.Data)
		case "Entity":
			entity, err = parseEntity(lm.Data)
			if err != nil {
				return err
			}
		default:
			entity = parsers[lm.Type](lm.Data)
		}
		l.AddEntity(entity)
	}
	return nil
}