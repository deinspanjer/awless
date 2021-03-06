package awsfetch

import "strings"

func sliceOfSlice(in []*string, maxLength int) (res [][]*string) {
	if maxLength <= 0 {
		return
	}
	if len(in) == 0 {
		return
	}
	for i := 0; i < len(in); i += maxLength {
		if i+maxLength < len(in) {
			res = append(res, in[i:i+maxLength])
		} else {
			res = append(res, in[i:])
		}
	}

	return
}

func appendIfNotInSlice(slice []string, s string) []string {
	var found bool
	for _, e := range slice {
		if e == s {
			found = true
		}
	}
	if !found {
		return append(slice, s)
	}
	return slice
}

func arnToName(arn string) string {
	splits := strings.Split(arn, "/")
	if len(splits) > 1 {
		return splits[len(splits)-1]
	}
	return arn
}

func pluralizeIfNeeded(str string, n uint) string {
	if n > 1 {
		return str + "s"
	}
	return str
}
