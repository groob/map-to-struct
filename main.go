package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/serenize/snaker"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	var kv map[string]string
	json.Unmarshal(data, &kv)
	for k, _ := range kv {
		fmt.Printf("%s", cond(k))
	}
}

func cond(k string) string {
	const conditional = `
	if v, ok := osMap["%s"]; ok {
		goStruct.%s = v
	}
`
	return fmt.Sprintf(conditional, k, snaker.SnakeToCamel(k))
}
