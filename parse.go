package sourcify

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ParseMetadata - try to parse metadata.json
func ParseMetadata(data string) (*Metadata, error) {
	var metadata Metadata
	err := json.UnmarshalFromString(data, &metadata)
	return &metadata, err
}
