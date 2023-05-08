package leetcode

import "strings"

// https://leetcode.com/problems/unique-email-addresses/
func NumUniqueEmails(emails []string) int {
	all := make(map[string]Void)
	for _, email := range emails {
		var sb strings.Builder
		var dm bool
		var ign bool
		for _, c := range email {
			if !dm {
				if c == '@' {
					dm = true
				}
				if !dm {
					if ign {
						continue
					}
					if c == '.' {
						continue
					}
					if c == '+' {
						ign = true
						continue
					}
				}
			}
			sb.WriteRune(c)
		}
		all[sb.String()] = Void{}
	}
	return len(all)
}
