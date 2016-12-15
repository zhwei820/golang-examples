package main

import (
	"fmt"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func main() {
	m := pcre.MustCompile("abc", pcre.CASELESS).MatcherString("Abc", 0)
	fmt.Printf("pcre 1: %v\n", m.Matches())

	m = pcre.MustCompile("abc", 0).MatcherString("Abc", 0)
	fmt.Printf("pcre 2: %v\n", m.Matches())

	m = pcre.MustCompile("<([A-Z][A-Z0-9]*)\\b[^>]*>.*?</\\1>", pcre.CASELESS).MatcherString("<H1>foobar</H1>", 0)
	fmt.Printf("pcre 3: %v\n", m.Matches())

	m = pcre.MustCompile("<([A-Z][A-Z0-9]*)\\b[^>]*>.*?</\\1>", pcre.CASELESS).MatcherString("<H1>foobar</H2>", 0)
	fmt.Printf("pcre 4: %v\n", m.Matches())

	m = pcre.MustCompile("abc", 0).MatcherString("Abc", 0)
	fmt.Printf("pcre 5: %v\n", m.Matches())

	re := pcre.MustCompile("foo", 0)
	result := re.ReplaceAll([]byte("I like foods."), []byte("car"), 0)
	fmt.Printf("pcre 6: %v\n", string(result))

	result = re.ReplaceAll([]byte("food fight fools foo"), []byte("car"), 0)
	fmt.Printf("pcre 7: %v\n", string(result))

	fmt.Println("")
	m = pcre.MustCompile("(?<Word>\\w+)", pcre.CASELESS).MatcherString("aabbcccdddd", 0)
	fmt.Printf("pcre 5: %v\n", m.Matches())

	fmt.Println("")
	m = pcre.MustCompile("(?<ISO8601>(?<year>\\d{4})(\\-W((?<week>\\d{1,2})\\-(?<weekday>\\d)?)|\\-(?<month>\\d{2})\\-(?<day>\\d{2})(T(?<hour>\\d{2}):(?<min>\\d{2})(:(?<sec>\\d{2})(\\+\\d{2}:\\d{2}Z?)?)?)?|\\-(?<yearday>\\d{1,3})))", pcre.CASELESS).MatcherString("aabbcccdddd", 0)
	fmt.Printf("pcre 5: %v\n", m.Matches())

	// r := pcre.MustCompile("(?<ISO8601>(?<year>\\d{4})(\\-W((?<week>\\d{1,2})\\-(?<weekday>\\d)?)|\\-(?<month>\\d{2})\\-(?<day>\\d{2})(T(?<hour>\\d{2}):(?<min>\\d{2})(:(?<sec>\\d{2})(\\+\\d{2}:\\d{2}Z?)?)?)?|\\-(?<yearday>\\d{1,3})))", pcre.CASELESS)

}
