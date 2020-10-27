package snake

import "strings"

// ToSnakeCase - converts string <s> formatted as camelCase or PascalCase to snake_case
func ToSnakeCase(s string) string {
	// for each underscore that is inserted to the slice,
	// it grows in size and hence the index becomes inconsistent
	// with input string. Offset makes up for such inconsistence.
	var offset int
	slice := strings.Split(s, "")

	for i := 1; i < len(s)-1; i++ {
		r := s[i]
		if string(r) >= "A" && string(r) <= "Z" {
			slice = append(slice[:i+1+offset], slice[i+offset:]...)
			slice[i+offset] = "_"
			offset++
		}
	}

	result := strings.Join(slice, "")
	result = strings.ToLower(result)

	return result
}

// ToKebabCase - converts string <s> formatted as camelCase or PascalCase to kebab-case
func ToKebabCase(s string) string {
	// for each underscore that is inserted to the slice,
	// it grows in size and hence the index becomes inconsistent
	// with input string. Offset makes up for such inconsistence.
	var offset int
	slice := strings.Split(s, "")

	for i := 1; i < len(s)-1; i++ {
		r := s[i]
		if string(r) >= "A" && string(r) <= "Z" {
			slice = append(slice[:i+1+offset], slice[i+offset:]...)
			slice[i+offset] = "-"
			offset++
		}
	}

	result := strings.Join(slice, "")
	result = strings.ToLower(result)

	return result
}
