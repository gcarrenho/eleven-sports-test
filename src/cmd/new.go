package new

/*
import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func main() {
	//bytes := []byte("")
	response, err := get("https://s3-eu-west-1.amazonaws.com/mycujoo-assignments/be-assignment/player.json")
	if err != nil {
		log.Fatal("Error to call get func.")
	}
	log.Println("Get response: ", response)

	statement, err := createStatement(response)
	if err != nil {
		log.Fatal("Error to create statement.")
	}
	log.Println("Statement: ", statement)
}

func createStatement(jsonStatement map[string]interface{}) (string, error) {
	var statement string

	table := jsonStatement["name"].(string)
	reflectFields := reflect.ValueOf(jsonStatement["fields"])
	if reflectFields.IsNil() {
		//log.Println("Error to Unmarshal fields of json statement.")
		return statement, errors.New("Error to Unmarshal fields of json statement.")
	}

	statement += "```sql \nCREATE TABLE " + table + " ("

	for i := 0; i < reflectFields.Len(); i++ {
		column := reflectFields.Index(i).Interface().(map[string]interface{})
		columnName := column["name"].(string)
		columnType := reflect.ValueOf(column["type"])
		var ctype string
		var nulleable bool
		var symbols string

		//Column type will can be string, slice or mapinterface{}
		if columnType.Kind() == reflect.String {
			//String
			ctype = columnType.String()
		} else if columnType.Kind() == reflect.Slice {
			//Slice
			for e := 0; e < columnType.Len(); e++ {
				if columnType.Index(e).Interface().(string) == "null" {
					nulleable = true
				} else {
					ctype = columnType.Index(e).Interface().(string)
				}
			}
		} else if columnType.Kind() == reflect.Map {
			//Map
			ctypeaux := columnType.Interface().(map[string]interface{})
			ctype = ctypeaux["type"].(string)
			elem := reflect.ValueOf(ctypeaux["symbols"])
			for e := 0; e < elem.Len(); e++ {
				if e != 0 {
					symbols += ", "
				}
				symbols += "'" + elem.Index(e).Interface().(string) + "'"
			}
		}

		if ctype == "enum" {
			statement += "\n" + columnName + " " + strings.ToUpper(ctype) + "(" + symbols + ")"
		} else {
			statement += "\n" + columnName + " " + strings.ToUpper(ctype)
			if nulleable == false {
				statement += " NOT NULL"
			}
		}
	}
	statement += "); \n ```"

	return statement, nil
}

func get(url string) (map[string]interface{}, error) {
	response := make(map[string]interface{})

	result, err := http.Get(url)
	if err != nil {
		log.Println("Fail to get ", url, ". Detail: ", err.Error())
		return response, err
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Println("Failt to read get response. Detail: ", err.Error())
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Fail to parse get response. Detail: ", err.Error())
		return response, err
	}

	return response, nil
}
*/
