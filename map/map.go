package _map

func Map() {
	//map is a collection of key-value pairs
	websites := map[string]string{
		"facebook": "https://facebook.com",
		"google":   "https://google.com",
		"twitter":  "https://twitter.com",
	}

	//add element
	websites["instagram"] = "https://instagram.com"

	//delete element
	delete(websites, "facebook")

	//map vs struct
	//map is a collection of key-value pairs
	//struct is a collection of fields
	//use map when you need a collection of key-value pairs
	//use struct when you need a collection of fields
	//map is flexible, struct is rigid

}
