# golang-course-dwarves-foundation

## Assignment 1

- Name: Reordering names based on country code
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
