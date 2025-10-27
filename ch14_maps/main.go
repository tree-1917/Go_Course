package main

import "fmt"

func main() {
	// ===================== //
	// 			Maps 		 //
	// ===================== //
	countries := map[string]string{
		"eg": "Egypt",
		"us": "United States",
		"uk": "United Kingdom",
		"fr": "France",
		"de": "Germany",
		"it": "Italy",
		"sa": "Saudi Arabia",
		"ae": "United Arab Emirates",
		"jp": "Japan",
		"cn": "China",
		"in": "India",
		"br": "Brazil",
		"ca": "Canada",
		"au": "Australia",
		"ru": "Russia",
	}

	for k, v := range countries {
		if strLen := len(v); strLen%2 == 1 {
			fmt.Printf("[%s] => %s ~> %d\n", k, v, strLen)
		}
	}

	if country, ok := countries["eg"]; ok {
		fmt.Println("[Eg] =>", country)
	} else {
		fmt.Println("Does not Exist !")
	}

	delete(countries, "eg")
	for k, v := range countries {
		if strLen := len(v); strLen%2 == 0 {
			fmt.Printf("[%s] => %s ~> %d\n", k, v, strLen)
		}
	}
}
