package argparser

import (
	"flag"
	"fmt"
)

func Parse_parameter_from_terminal() ([]string, string, bool){
    
	mode, args := parse_mode_and_array()
	has_zero_args := check_zero_arguments(args)
	is_continued := true
	if (has_zero_args) {
		print_program_usage_instruction()
		is_continued = false
	}
	return args, mode, is_continued
}

func parse_mode_and_array() (string, []string) {
	var is_integer bool
	var is_float bool
	var is_string bool
	var is_mix bool
    flag.BoolVar(&is_integer, "int", false, "followed by an integer array")
    flag.BoolVar(&is_float, "float", false, "followed by a float array")
    flag.BoolVar(&is_string, "string", false, "followed by a string array")
    flag.BoolVar(&is_mix, "mix", false, "followed by a mixed array")
    flag.Parse()
	
	mode := parse_mode(is_integer, is_float, is_string, is_mix)
	args := flag.Args()
    
	return mode, args
}

func parse_mode(is_integer, is_float, is_string, is_mix bool) string {
	var mode string
	if (is_integer) {
        mode = "int"
    } else if (is_float) {
        mode = "float"
    } else if (is_string) {
        mode = "string"
    } else {
        mode = "mix"
    }
	return mode
}

func check_zero_arguments(args []string) bool {
	if len(args) == 0 {
		return true
	}
	return false
}

func print_program_usage_instruction() {
	fmt.Println("Usage: sorter.go [-float | -int | -mix | -string] objects...")
	flag.PrintDefaults()
}