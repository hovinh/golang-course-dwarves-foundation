package main

import (
	"fmt"
	"os"
	"sortermodule/argparser"
	"sortermodule/sorterbytype"
)

func main() {

    args, mode, is_continued := argparser.Parse_parameter_from_terminal()
    if (!is_continued) {
        os.Exit(1)
    }

    switch mode {
        case "int": {
            output_arr := sorterbytype.Sort_string_array_with_integer_mode(args)
            print_int_arr(output_arr)
        }
        case "float": {
            output_arr := sorterbytype.Sort_string_array_with_float_mode(args)
            print_float_arr(output_arr)
        }
        case "string": {
            output_arr := sorterbytype.Sort_string_array_with_string_mode(args)
            print_str_arr(output_arr)
        }
        case "mix": {
            output_arr := sorterbytype.Sort_string_array_with_mix_mode(args)
            print_str_arr(output_arr)
        }
    }

}

func print_str_arr(str_arr []string) {
    for i:= 0; i < len(str_arr); i++ {
        fmt.Print(str_arr[i] + " ")
    }
}

func print_int_arr(int_arr []int) {
    for i:= 0; i < len(int_arr); i++ {
        fmt.Print(int_arr[i], " ")
    }
}
func print_float_arr(float_arr []float64) {
    for i:= 0; i < len(float_arr); i++ {
        fmt.Print(float_arr[i],  " ")
    }
}