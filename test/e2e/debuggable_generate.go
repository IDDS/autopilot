// use this file to manually invoke code generation (without ap CLI)
// can be useful for debugging
package main

import (
	"log"

	"gitlab.dds-sysu.tech/691729768/autopilot/codegen"
)

func main() {
	if err := codegen.Run("/Users/ilackarms/go/src/autorouter", false, false); err != nil {
		log.Fatal(err)
	}
}
