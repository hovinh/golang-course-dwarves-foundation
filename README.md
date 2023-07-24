# golang-course-dwarves-foundation

## Assignment 1

- Objective: Create a program that reorders names based on the specified country's format.
- Inputs: First name, last name, middle name (optional), and country code.
- Outputs: Reordered name based on the country's format.
- Description: Parse command line arguments to extract the first name, last name, middle name (if provided), and the country code. Determine the name format based on the country code. Reorder the name according to the determined format. Output the reordered name.
- Examples:

```bash
$ go run reorder_names.go John Smith VN
Output: Smith John
$ go run reorder_names.go Emily Rose Watson US
Output: Emily Rose Watson
```

- Program:

```bash
$ cd assignment1_name_reordering
$ go run reorder_names.go John Smith VN
```

- Unit test:

```bash
$ cd assignment1_name_reordering
$ go test -v
```

## Assignment 2

- Objective: A CLI tool to sort input provided by user
- Inputs: Number (integer or float) array, string array
- Outputs: Sorted result based on the provided input type.
- Description: create a CLI tool to parse the input from the command line. Determine the type of input (integer arrays, float arrays, string arrays, or mixed). Utilize the corresponding sorting function from the package to sort the elements. Output the sorted result.
- Examples:

```bash
$ go run sorter.go -int 5 2 10 1
Output: 1 2 5 10
$ go run sorter.go -string apple orange banana
Output: apple banana orange
$ go run sorter.go -mix 5.5 apple 2.7 orange 3 banana
Output: 2.7 3 5.5 apple banana orange
```
