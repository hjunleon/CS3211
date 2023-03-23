package main

import (
	"fmt"
	"os"
)

var DEBUG_LEVEL int = 1

var DEBUG_OB = "ob"
var DEBUG_EG = "eg"
var DEBUG_OR = "or"

var DEBUG_PART = map[string]bool{
	(DEBUG_EG): true,
	(DEBUG_OB): true,
	(DEBUG_OR): true,
}

func log_debug(debug_level int, debug_part string, to_print ...string) {
	if DEBUG_PART[debug_part] == false {
		return
	}
	if debug_level == 0 {
		debug_level = 1
	}
	if debug_level < DEBUG_LEVEL {
		return
	}
	fmt.Fprintf(os.Stderr, "[DEBUG_%v]: ", debug_level)
	for _, el := range to_print {
		fmt.Fprintf(os.Stderr, el)
	}
	// fmt.Println("")
	fmt.Fprintf(os.Stderr, "\n")
	// fmt.Fprintf(os.Stderr, "%v\n", to_print)
}
