package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	response, err := doRequest()
	if err != nil {
		log.Fatal(err)
	}

	table := buildCommandCreateDBTable(response)
	showCommandCreateTable(table)

}

func doRequest() (map[string]interface{}, error) {
	respose, err := http.Get("https://s3-eu-west-1.amazonaws.com/mycujoo-assignments/be-assignment/player.json")
	if err != nil {
		log.Fatal(err)
	}

	responsByte, err := ioutil.ReadAll(respose.Body)
	if err != nil {
		log.Fatal(err)
	}

	players := make(map[string]interface{})
	err = json.Unmarshal(responsByte, &players)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return players, nil
}
func showCommandCreateTable(table string) {
	fmt.Println("================================")
	fmt.Println(table)
	fmt.Println("================================")
}

func buildCommandCreateDBTable(players map[string]interface{}) string {
	fields := players["fields"]
	sliceFields := fields.([]interface{})
	var column map[string]interface{}
	createTableCommand := "CREATE TABLE " + players["name"].(string) + "("

	for _, f := range sliceFields {
		column = f.(map[string]interface{})
		columName := column["name"].(string)
		columType := checkType(column["type"])
		createTableCommand += columName + " " + columType + ",\n"
	}

	createTableCommandLength := len(createTableCommand)

	return createTableCommand[0:createTableCommandLength-2] + ");"
}

func checkType(t interface{}) string {

	if c, ok := t.(string); ok {
		return translateString(c)
	}

	if l, ok := t.([]interface{}); ok {
		return translateArray(l)
	}

	if o, ok := t.(map[string]interface{}); ok {
		return translateEnum(o)
	}

	return ""
}

func translateString(c string) string {
	if c == "int" {
		return "INT"
	}
	if c == "string" {
		return "TEXT"
	}
	if c == "boolean" {
		return "BOOLEAN"
	}
	return ""
}

func translateArray(a []interface{}) string {
	if a[0] == "int" || a[1] == "int" {
		return "INT NULL"
	}
	if a[0] == "string" || a[1] == "string" {
		return "TEX NULL"
	}
	return ""
}
func translateEnum(o map[string]interface{}) string {
	enum := "ENUM("
	if o["type"] == "enum" {
		symbols := o["symbols"]
		r := symbols.([]interface{})
		for _, s := range r {
			enum += "'" + s.(string) + "'" + ","
		}
		enumLength := len(enum)
		return enum[0:enumLength-1] + ")"
	}
	return ""
}
