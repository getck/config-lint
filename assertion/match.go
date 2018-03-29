package assertion

import (
	"regexp"
	"strings"
)

func isNil(data interface{}) bool {
	return data == nil
}

func isString(data interface{}) bool {
	_, ok := data.(string)
	return ok
}

func convertToString(data interface{}) (string, bool) {
	s, ok := data.(string)
	return s, ok
}

func isMatch(data interface{}, op string, value string, valueType string) (bool, error) {
	searchResult, err := JSONStringify(data)
	if err != nil {
		return false, err
	}
	searchResult = unquoted(searchResult)
	switch op {
	case "eq":
		if compare(data, value, valueType) == 0 {
			return true, nil
		}
	case "ne":
		if compare(data, value, valueType) != 0 {
			return true, nil
		}
	case "lt":
		if compare(data, value, valueType) < 0 {
			return true, nil
		}
	case "le":
		if compare(data, value, valueType) <= 0 {
			return true, nil
		}
	case "gt":
		if compare(data, value, valueType) > 0 {
			return true, nil
		}
	case "ge":
		if compare(data, value, valueType) >= 0 {
			return true, nil
		}
	case "in":
		for _, v := range strings.Split(value, ",") {
			if v == searchResult {
				return true, nil
			}
		}
	case "not-in":
		for _, v := range strings.Split(value, ",") {
			if v == searchResult {
				return false, nil
			}
		}
		return true, nil
	case "absent":
		if isAbsent(searchResult) {
			return true, nil
		}
	case "present":
		if isPresent(searchResult) {
			return true, nil
		}
	case "null":
		return isNil(data), nil
	case "not-null":
		return !isNil(data), nil
	case "empty":
		return isEmpty(data), nil
	case "not-empty":
		return !isEmpty(data), nil
	case "intersect":
		if jsonListsIntersect(searchResult, value) {
			return true, nil
		}
	case "contains":
		if c, isSlice := convertToSlice(data); isSlice {
			for _, element := range c {
				if stringElement, isString := element.(string); isString {
					if stringElement == value {
						return true, nil
					}
				}
			}
			return false, nil
		}
		if s, isString := convertToString(data); isString {
			if strings.Contains(s, value) {
				return true, nil
			}
			return false, nil
		}
		return strings.Contains(searchResult, value), nil
	case "regex":
		if regexp.MustCompile(value).MatchString(searchResult) { // FIXME don't use MustCompile
			return true, nil
		}
		return false, nil
	}
	return false, nil
}
