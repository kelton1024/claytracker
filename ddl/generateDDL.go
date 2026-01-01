package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
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

func createDDL(schemaContent *Schema) (string, error) {
	var strBuilder strings.Builder
	strBuilder.WriteString(fmt.Sprintf("---- Table: %v\n", schemaContent.Id))
	strBuilder.WriteString(fmt.Sprintf("DROP TABLE IF EXISTS %v CASCADE;\n", schemaContent.Id))
	strBuilder.WriteString(fmt.Sprintf("CREATE TABLE %v\n(\n", schemaContent.Id))

	var pk, reg, fk []string
	for key, value := range schemaContent.Properties {
		sqlType, ok := typeMappings[value.Type]
		if !ok {
			return "", fmt.Errorf("encountered unknown type '%v'", value.Type)
		}

		if value.Comment != "" {
			sqlType = value.Comment
		}

		columnData := fmt.Sprintf("\t%v %v,\n", key, sqlType)
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

	// Since the order of maps are not deterministic, we will sort so we are consistent after each run
	slices.Sort(pk)
	slices.Sort(reg)
	slices.Sort(fk)
	combinedSQL := slices.Concat(pk, reg, fk)
	for i, column := range combinedSQL {
		if i == len(combinedSQL)-1 {
			column = strings.ReplaceAll(column, ",", "")
		}
		strBuilder.WriteString(column)
	}

	strBuilder.WriteString(");\n\n")
	return strBuilder.String(), nil
}

func generateDDL() ([]string, error) {
	var ddl []string
	ddl = append(ddl, fmt.Sprintf("----Drop Database\nDROP DATABASE IF EXISTS %v;\n", databaseName))
	ddl = append(ddl, fmt.Sprintf("CREATE DATABASE %v;\n\n", databaseName))

	schemaFiles, err := gatherFiles()
	if err != nil {
		return nil, err
	}

	for _, file := range schemaFiles {
		fullPath := path.Join(jsonDirectory, file)
		fileContent, err := os.ReadFile(fullPath)
		if err != nil {
			return nil, err
		}

		schemaContent := Schema{}
		err = json.Unmarshal(fileContent, &schemaContent)
		if err != nil {
			return nil, err
		}

		generatedSQL, err := createDDL(&schemaContent)
		if err != nil {
			return nil, err
		}
		ddl = append(ddl, generatedSQL)
	}
	return ddl, nil
}
