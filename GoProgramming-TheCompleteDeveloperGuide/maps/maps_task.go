package main

import "fmt"

func printServerStatus(serversMap map[string]string) {
	stats := make(map[string]int)

	fmt.Println("-----------------------------------")
	fmt.Println("Total server number:", len(serversMap))

	for _, v := range serversMap {
		if v == "Online" {
			stats["Online"]++
		} else if v == "Offline" {
			stats["Offline"]++
		} else if v == "Maintenance" {
			stats["Maintenance"]++
		} else if v == "Retired" {
			stats["Retired"]++
		}
	}
	fmt.Println("Online", stats["Online"])
	fmt.Println("Offline", stats["Offline"])
	fmt.Println("Maintenance", stats["Maintenance"])
	fmt.Println("Retired", stats["Retired"])
	fmt.Println("-----------------------------------")
}

func main() {
	servers := []string{"darkstar", "aiur", "hope", "blackhole", "r2d2"}

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
