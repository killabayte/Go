package main

import "fmt"

func printServerStatus(serversMap map[string]string) {
	online := 0
	offline := 0
	maintenance := 0
	retired := 0

	fmt.Println("-----------------------------------")
	fmt.Println("Total server number:", len(serversMap))

	for _, v := range serversMap {
		if v == "Online" {
			online++
		} else if v == "Offline" {
			offline++
		} else if v == "Maintenance" {
			maintenance++
		} else if v == "Retired" {
			retired++
		}
	}
	fmt.Println("Online", online)
	fmt.Println("Offline", offline)
	fmt.Println("Maintenance", maintenance)
	fmt.Println("Retired", retired)
	fmt.Println("-----------------------------------")
}

func main() {
	servers := []string{"darkstar", "aiur", "hope", "blackhole"}

	serversMap := make(map[string]string)

	for _, v := range servers {
		serversMap[v] = "Online"
	}
	printServerStatus(serversMap)
	serversMap["darkstar"] = "Retired"
	serversMap["aiur"] = "Offline"
	printServerStatus(serversMap)

	for _, v := range servers {
		serversMap[v] = "Maintenance"
	}
	printServerStatus(serversMap)
}
