package auth

import "strings"

//VerifyAdminProfile doc
//Summary Verify permissions
//param ([]UserPerm) permissions array
//param (string)     access uri
//param (int)        need permissions
//return (bool)      yes/no
func VerifyAdminProfile(profiles []AdminUserProfile, uri string, need int) bool {
	for _, v := range profiles {
		if strings.ToLower(v.URI) == "all" && (v.Auth&need) != 0 {
			return true
		}

		if strings.Index(uri, v.URI) >= 0 && (v.Auth&need) != 0 {
			return true
		}
	}

	return false
}