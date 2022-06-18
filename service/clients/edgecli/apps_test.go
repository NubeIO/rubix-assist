package edgecli

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"reflect"
	"strings"
	"testing"
	"unicode"
)

func IsWhitespace(r rune) bool {
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

func ToTitleCase(s string) string {
	prev := ' '
	s = strings.Replace(s, "_", " ", -1)
	result := strings.Map(
		func(r rune) rune {
			if IsWhitespace(prev) || '_' == prev || '-' == prev {
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

type out struct {
	key  string
	name string
	val  interface{}
}

func TestHost(*testing.T) {

	cli := New("", 0)
	time, msg, err := cli.GetTime()
	fmt.Println(err)
	if msg != nil {
		return
	}

	ee, _ := json.Marshal(&time)

	value := gjson.ParseBytes(ee)

	val := reflect.ValueOf(value.Value())

	var aa []out

	if val.Kind() == reflect.Map {
		for _, e := range val.MapKeys() {
			v := val.MapIndex(e)
			switch t := v.Interface().(type) {
			case int:
			case string:
				aa = append(aa, out{key: e.String(), name: ToTitleCase(e.String()), val: t})
			case bool:
			default:

			}
		}
	}
	fmt.Println(aa)

}
