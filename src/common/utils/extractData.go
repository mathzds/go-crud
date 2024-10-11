package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	models "api/src/common/types"
)

func ExtractData(html string, dataType string) (interface{}, error) {
	var jsonString string

	if splitData := strings.Split(html, "__NEXT_DATA__"); len(splitData) > 1 {
		if jsonPart := strings.Split(splitData[1], `type="application/json">`); len(jsonPart) > 1 {
			jsonString = strings.Split(jsonPart[1], "</script>")[0]
		} else {
			return nil, fmt.Errorf("nextdata not found")
		}
	} else {
		return nil, fmt.Errorf("nextdata not found")
	}

	switch dataType {
	case "release":
		var data models.Data
		if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
			return nil, err
		}
		return data.Props.PageProps.Data.DataReleases, nil
	case "anime":
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
			return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
		}
		return data, nil
	default:
		return nil, fmt.Errorf("unknown data type: %s", dataType)
	}
}
