package main

import str "strings"
import . "fmt"

func isInterleave(s1, s2, s3 string) bool {
	if len(s3) == 0 && len(s1) == 0 && len(s2) == 0 {
		return true
	} else if len(s1) == 0 {
		if str.Compare(s2, s3) == 0 {
			return true
		} else {
			return false
		}
	} else if len(s2) == 0 {
		if str.Compare(s1, s3) == 0 {
			return true
		} else {
			return false
		}
	} else {
		s1b := []byte(s1)
		s2b := []byte(s2)
		s3b := []byte(s3)
		if s1b[0] == s3b[0] {
			if s2b[0] == s3b[0] {
				return isInterleave(string(s1b[:]), string(s2b[1:]), string(s3b[1:])) || isInterleave(string(s1b[1:]), string(s2b[:]), string(s3b[1:]))
			} else {
				return isInterleave(string(s1b[1:]), string(s2b[:]), string(s3b[1:]))
			}
		} else if s2b[0] == s3b[0] {
			return isInterleave(string(s1b[:]), string(s2b[1:]), string(s3b[1:]))
		} else {
			return false
		}
	}
}

func main() {
	str1 := "aabccc"
	str2 := "dbbca"
	str3 := "aadbbcbcac"
	result := isInterleave(str1, str2, str3)
	Println(result)
}
