package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strings"
)

const (
	jsonDirectory = "../json/"
	ddlDirectory  = "./"
	databaseName  = "range_tracker"
)

var typeMappings = map[string]string{
	"number": "int",
	"string": "text",
}

type Schema struct {
	Schema               string                `json:"$schema,omitempty"`
	Id                   string                `json:"$id,omitempty"`
	Title                string                `json:"title,omitempty"`
	Description          string                `json:"description,omitempty"`
	Type                 string                `json:"type,omitempty"`
	Properties           map[string]PropValues `json:"properties,omitempty"`
	Required             []string              `json:"required,omitempty"`
	AdditionalProperties bool                  `json:"additionalProperties,omitempty"`
}

type PropValues struct {
	Type    string `json:"type"`
	Comment string `json:"$comment"`
}

// TODO: Pass directory instead of using globals variables
func gatherFiles() ([]string, error) {
	directory, err := os.ReadDir(jsonDirectory)
	if err != nil {
		return nil, err
	}

	schemaFiles := []string{}
	for _, file := range directory {
		if strings.Contains(file.Name(), "example") {
			continue
		}
		schemaFiles = append(schemaFiles, file.Name())
	}
	return schemaFiles, nil
}

func createDDL(schemaContent *Schema) ([]byte, error) {
	var strBuilder bytes.Buffer
	strBuilder.WriteString(fmt.Sprintf("---- Table: %v\n", schemaContent.Id))
	strBuilder.WriteString(fmt.Sprintf("DROP TABLE IF EXISTS %v CASCADE;\n", schemaContent.Id))
	strBuilder.WriteString(fmt.Sprintf("CREATE TABLE %v\n(\n", schemaContent.Id))

	var pk, reg, fk []string
	for key, value := range schemaContent.Properties {
		sqlType, ok := typeMappings[value.Type]
		if !ok {
			return nil, fmt.Errorf("encountered unknown type '%v'", value.Type)
		}

		if value.Comment != "" {
			sqlType = value.Comment
		}

		columnData := fmt.Sprintf("\t%v %v\n", key, sqlType)
		if strings.Contains(sqlType, "primary key") {
			pk = append(pk, columnData)
			continue
		}
		if strings.Contains(sqlType, "references") {
			fk = append(fk, columnData)
			continue
		}
		reg = append(reg, columnData)

	}

	combinedSQL := slices.Concat(pk, reg, fk)
	for _, column := range combinedSQL {
		strBuilder.WriteString(column)
	}

	strBuilder.WriteString(");\n\n")
	return strBuilder.Bytes(), nil
}

func main() {
	schemaFiles, err := gatherFiles()
	if err != nil {
		log.Fatalf("failed to read directory '%v' with the following error %v", jsonDirectory, err)
	}

	outputFile, err := os.Create("gen_range_tracker_relational_model.sql")
	if err != nil {
		log.Fatalf("failed to create the output file")
	}
	dropDDL := fmt.Sprintf("DROP DATABASE IF EXISTS %v CASCADE;\nCREATE DATABASE %v\n\n", databaseName, databaseName)
	outputFile.Write([]byte(dropDDL))

	for _, file := range schemaFiles {
		fullPath := path.Join(jsonDirectory, file)
		fileContent, err := os.ReadFile(fullPath)
		if err != nil {
			fmt.Printf("failed to read file contents for file '%v' with error %v\n", file, err)
			continue
		}

		schemaContent := Schema{}
		err = json.Unmarshal(fileContent, &schemaContent)
		if err != nil {
			fmt.Printf("failed to unmarhsal file content for file '%v' with error %v\n", file, err)
			continue
		}

		generatedSQL, err := createDDL(&schemaContent)
		if err != nil {
			fmt.Printf("failed to create DDL using file '%v' with error %v\n", file, err)
		}
		outputFile.Write(generatedSQL)
	}
}
