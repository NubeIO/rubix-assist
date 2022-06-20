package assitcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/iancoleman/strcase"
	"github.com/stretchr/objx"
	"github.com/tidwall/gjson"
	"reflect"
	"strings"
	"testing"
	"unicode"
)

type Out struct {
	Title string      `json:"title"`
	Key   string      `json:"key"`
	Index string      `json:"dataIndex"`
	Data  interface{} `json:"data"`
}

func ArrayOfMaps(data []byte) *[]Out {
	value := gjson.ParseBytes(data)
	val := reflect.ValueOf(value.Value())
	var res []Out
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			if val.Index(i).Kind() == reflect.Interface {
				maps := val.Index(i).Elem()
				if maps.Kind() == reflect.Map {
					for _, e := range maps.MapKeys() {
						v := maps.MapIndex(e)
						switch t := v.Interface().(type) {
						case int:
						case string:
							res = append(res, Out{Key: e.String(), Title: toTitleCase(e.String()), Data: t})
						case bool:
						default:

						}
					}
				}
			}
		}

	}
	return &res
}

func isWhitespace(r rune) bool {
	result := false
	switch r {
	case
		'\u0009', // horizontal tab
		'\u000A', // line feed
		'\u000B', // vertical tab
		'\u000C', // form feed
		'\u000D', // carriage return
		'\u0020', // space
		'\u0085', // next line
		'\u00A0', // no-break space
		'\u1680', // ogham space mark
		'\u180E', // mongolian vowel separator
		'\u2000', // en quad
		'\u2001', // em quad
		'\u2002', // en space
		'\u2003', // em space
		'\u2004', // three-per-em space
		'\u2005', // four-per-em space
		'\u2006', // six-per-em space
		'\u2007', // figure space
		'\u2008', // punctuation space
		'\u2009', // thin space
		'\u200A', // hair space
		'\u2028', // line separator
		'\u2029', // paragraph separator
		'\u202F', // narrow no-break space
		'\u205F', // medium mathematical space
		'\u3000': // ideographic space
		result = true
	default:
		result = false
	}
	return result
}

func toTitleCase(s string) string {
	prev := ' '
	s = strings.Replace(strcase.ToSnake(s), "_", " ", -1)
	result := strings.Map(
		func(r rune) rune {
			if isWhitespace(prev) || '_' == prev || '-' == prev {
				prev = r
				return unicode.ToTitle(r)
			} else {
				prev = r
				return unicode.ToLower(r)
			}
		},
		s)
	return result
}

func Map(data []byte) *[]Out {
	value := gjson.ParseBytes(data)
	val := reflect.ValueOf(value.Value())
	var res []Out
	//fmt.Println(string(val.Bytes()))
	//var inInterface []model.Location
	//mapstructure.Decode(data, &inInterface)
	//for i, location := range inInterface {
	//	fmt.Println(i, location)
	//
	//}

	if val.Kind() == reflect.Map {
		for _, e := range val.MapKeys() {
			//v := val.MapIndex(e)
			fmt.Println(e.String())
			//fmt.Println(e.MapKeys())
			//fmt.Println(v.Type())

			//rv := reflect.ValueOf(&mapped).Elem()

			//i := v.(*model.Location)
			//var inInterface []model.Location
			//err := json.Unmarshal(v.Bytes(), &inInterface)
			//fmt.Println(err)
			//if err != nil {
			//	return nil
			//}
			//
			//for i, location := range inInterface {
			//	fmt.Println(i, location)
			//}

			//fmt.Println(v)
			//for i := 0; i < e.Len(); i++ {
			//	fmt.Println(e.String())
			//
			//	//maps := e.Index(i).Elem()
			//	//fmt.Println(maps)
			//}

			//fmt.Println(v)
			//val := reflect.ValueOf(v)
			//fmt.Println(val)

			//if v.Kind() == reflect.Map {
			//	fmt.Println(v.MapKeys())
			//	for _, f := range v.MapKeys() {
			//		fmt.Println(f.Type())
			//		fmt.Println(f.Kind())
			//	}
			//
			//}

			//switch t := v.Interface().(type) {
			//case int:
			//case string:
			//	res = append(res, Out{Key: e.String(), Title: toTitleCase(e.String()), Data: t})
			//case bool:
			//default:
			//
			//}
		}
	}
	return &res
}

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}))
		default:
			fmt.Println(key, ":", concreteVal)
		}
	}
}

func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}))
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}

type TableSchema struct {
	Title string `json:"title"`
	Key   string `json:"key"`
	Index string `json:"dataIndex"`
}

func BuildTableSchema(data interface{}) []TableSchema {
	var schema []TableSchema
	for i, value := range objx.New(data) {
		v := objx.New(value)
		title := ""
		key := i
		index := key
		for k, v := range v {
			if k == "title" {
				title = v.(string)
			}
		}
		schema = append(schema, TableSchema{Title: title, Key: key, Index: index})
	}
	return schema
}

func TestClient_InstallApp(t *testing.T) {

	client := New("0.0.0.0", 1662)
	d, _ := client.GetLocationSchema()

	pprint.PrintJOSN(BuildTableSchema(d))

}
