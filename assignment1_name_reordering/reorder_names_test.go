package main

import (
	"fmt"
	"testing"
)

func Test_Valid_Number_Arguments(t *testing.T) {
	input_args := []string{"program_name", "Ho", "Vinh", "Xuan", "VN"}
	is_valid := verify_number_arguments(input_args)
	if is_valid != true {
		t.Fatalf("Program must accept 4 arguments (excluding program name).")
	}

	input_args = []string{"program_name", "Ho", "Vinh", "VN"}
	is_valid = verify_number_arguments(input_args)
	if is_valid != true {
		t.Fatalf("Program must accept 3 arguments (excluding program name).")
	}
}

func Test_Valid_Country_Code(t *testing.T) {
	country_code_list := []string{"VN", "US"}
	var is_valid bool
	for _, country_code := range country_code_list {
		is_valid = verify_country_code((country_code))
		if (is_valid == false) {
			t.Fatalf("Program must accept 'VN' and 'US' country code.")
		}	
	}
}

func Test_US_Name_Without_Middle(t *testing.T) {
	input_args := []string{"program_name", "Vinh", "Ho", "US"}
	expected := fmt.Sprintf("Vinh Ho")
	output_name := process_input_args(input_args)
	if (expected != output_name) {
		err_msg := fmt.Sprintf("Expected: %s | Output: %s", expected, output_name)
		t.Fatalf(err_msg)
	}
}

func Test_US_Name_With_Middle(t *testing.T) {
	input_args := []string{"program_name", "Vinh", "Ho", "Xuan", "US"}
	expected := fmt.Sprintf("Vinh Xuan Ho")
	output_name := process_input_args(input_args)
	if (expected != output_name) {
		err_msg := fmt.Sprintf("Expected: %s | Output: %s", expected, output_name)
		t.Fatalf(err_msg)
	}
}

func Test_VN_Name_Without_Middle(t *testing.T) {
	input_args := []string{"program_name", "Vinh", "Ho", "VN"}
	expected := fmt.Sprintf("Ho Vinh")
	output_name := process_input_args(input_args)
	if (expected != output_name) {
		err_msg := fmt.Sprintf("Expected: %s | Output: %s", expected, output_name)
		t.Fatalf(err_msg)
	}
}

func Test_VN_Name_With_Middle(t *testing.T) {
	input_args := []string{"program_name", "Vinh", "Ho", "Xuan", "VN"}
	expected := fmt.Sprintf("Ho Xuan Vinh")
	output_name := process_input_args(input_args)
	if (expected != output_name) {
		err_msg := fmt.Sprintf("Expected: %s | Output: %s", expected, output_name)
		t.Fatalf(err_msg)
	}
}