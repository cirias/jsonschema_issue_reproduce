package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/jsonschema"
)

type Object struct {
	Embedded `json:",some_tag,does_not_matter"`
}

type Embedded struct {
	F1 string `json:"f1"`
	F2 string `json:"f2"`
}

func main() {
	reflector := jsonschema.Reflector{}
	schema := reflector.Reflect(&Object{})
	jsonBytes, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}
