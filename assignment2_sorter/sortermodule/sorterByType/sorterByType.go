package sorterbytype

import (
	"fmt"
	"sort"
	"strconv"
)

func Sort_string_array_with_integer_mode(str_arr []string) []int{
	int_arr := convert_string_to_int_array(str_arr)
	int_arr = sort_integer_array(int_arr)
	return int_arr
}

func convert_string_to_int_array(str_arr []string) []int {
	int_arr := []int{}
	for i := range str_arr {
		str := str_arr[i]
		int_numb, _ := strconv.Atoi(str)
		int_arr = append(int_arr, int_numb)
	}

	return int_arr
}

func sort_integer_array(int_arr []int) []int {
	sort.Ints(int_arr)
	return int_arr
}

func Sort_string_array_with_float_mode(str_arr []string) []float64{
	float_arr := convert_string_to_float_array(str_arr)
	float_arr = sort_float_array(float_arr)
	return float_arr
}

func convert_string_to_float_array(str_arr []string) []float64 {
	float_arr := []float64{}
	for i := range str_arr {
		str := str_arr[i]
		float_numb, _ := strconv.ParseFloat(str, 64)
		float_arr = append(float_arr, float_numb)
	}
	return float_arr
}

func sort_float_array(float_arr []float64) []float64 {
	sort.Float64s(float_arr)
	return float_arr
}

func Sort_string_array_with_string_mode(str_arr []string) []string{
	sort.Strings(str_arr)
	return str_arr
}

func Sort_string_array_with_mix_mode(str_arr []string) []string{

	sort.SliceStable(str_arr, func(i, j int) bool {
		i_val := str_arr[i]
		j_val := str_arr[j]
		i_type := get_type(i_val)
		j_type := get_type(j_val)

		if ((i_type == "string") || (j_type == "string")) {
			i_as_string := fmt.Sprintf("%v",i_val)
			j_as_string := fmt.Sprintf("%v",j_val)
			return i_as_string < j_as_string
		} else {
			i_as_float, _ := strconv.ParseFloat(i_val, 64) 
			j_as_float, _ := strconv.ParseFloat(j_val, 64) 
			return i_as_float < j_as_float
		}
		
	})

	return str_arr
}

func get_type(str string) string {
	if (check_integer_type(str)){
		return "int"
	} else if (check_float_type(str)) {
		return "float"
	}
	return "string"
}

func check_integer_type(str string) bool {
	_, err := strconv.Atoi(str)
	if (err == nil) {
		return true
	}
	return false
}

func check_float_type(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	if (err == nil) {
		return true
	}
	return false
}
