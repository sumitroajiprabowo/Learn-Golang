package main

import "fmt"

type listBlacklist func(string) bool

func registerUser(name string, blacklist listBlacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

func blacklistUser(name string) bool {
	slice := []string{"root", "admin", "admin2"}
	for _, value := range slice {
		if name == value {
			return true
		}
	}
	return false
}

func main() {
	// blackList := func(name string) bool {
	// 	return name == "root"
	// }

	blacklistName := blacklistUser

	registerUser("root", blacklistName)
	registerUser("Bowo", blacklistName)

	registerUser("Danu", func(name string) bool {
		return name == "admin"
	})

	registerUser("admin", func(name string) bool {
		return name == "admin"
	})

}
