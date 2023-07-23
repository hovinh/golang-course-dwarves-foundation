package main

import (
	"fmt"
	"os"
)

const N_VALID_ARGS_WITHOUT_MIDDLE_NAME = 3
const N_VALID_ARGS_WITH_MIDDLE_NAME = 4

func main() {
	arg_list := os.Args
	is_valid_input := verify_input_args(arg_list)
	if (is_valid_input == false) {
		return;
	}

	output_name := process_input_args(arg_list)
	
	fmt.Println(output_name)
}

func verify_input_args(arg_list []string) bool {
	if (verify_number_arguments(arg_list) == false) {
		return false
	}
	_, _, _, country_code := parse_arguments_to_name_and_country(arg_list)
	if (verify_country_code(country_code) == false) {
		return false
	}
	return true
}

func verify_number_arguments(arg_list []string) bool {
	n_input_args := get_number_input_args(arg_list)
	
	if (n_input_args != N_VALID_ARGS_WITHOUT_MIDDLE_NAME && n_input_args != N_VALID_ARGS_WITH_MIDDLE_NAME) {
		print_invalid_number_arguments(n_input_args)
		return false
	}
	return true
}

func get_number_input_args(arg_list []string) int {
	return len(arg_list) - 1
}

func print_invalid_number_arguments(n_input_args int) {
	fmt.Printf("Number of input argument = %d is invalid. Please use one of the following formats:\n", n_input_args)
	fmt.Println("Option 1: go run reorder_names.go first-name last-name country-code")
	fmt.Println("Option 2: go run reorder_names.go first-name last-name middle name country-code")
}

func parse_arguments_to_name_and_country(arg_list []string) (string, string, string, string) {
	n_input_args := get_number_input_args(arg_list)
	var first_name, last_name, middle_name, country_code string
	if n_input_args == N_VALID_ARGS_WITHOUT_MIDDLE_NAME {
		first_name, last_name, country_code = arg_list[1], arg_list[2], arg_list[3]
		middle_name = ""
	} else {
		first_name, last_name, middle_name, country_code = arg_list[1], arg_list[2], arg_list[3], arg_list[4]
	}

	return first_name, last_name, middle_name, country_code
}

func verify_country_code(country_code string) bool {
	switch country_code {
	case "VN", "US": return true
	default: {
		print_invalid_country_code(country_code)
		return false		
	}
}
}

func print_invalid_country_code(country_code string) {
	fmt.Printf("country-code = '%s' is invalid. Acceptable values are ['VN', 'US']", country_code)
}

func process_input_args(arg_list []string) string {
	first_name, last_name, middle_name, country_code := parse_arguments_to_name_and_country(arg_list)
	var output_name string

	if (middle_name != ""){
		switch country_code {
		case "VN": output_name = fmt.Sprintf("%s %s %s", last_name, middle_name, first_name)
		case "US": output_name = fmt.Sprintf("%s %s %s", first_name, middle_name, last_name)
		}
	} else if (middle_name == ""){
		switch country_code {
			case "VN": output_name = fmt.Sprintf("%s %s", last_name, first_name)
			case "US": output_name = fmt.Sprintf("%s %s", first_name, last_name)
		}
	}
	return output_name
}


