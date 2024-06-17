package utils

var icons map[string]string
icons["github.com"] = "g"
icons["instagram.com"] = "i"

func WebsiteIcons(url string) string {
	var parsedUrl string
	parsedUrl = strings.TrimPrefix(url, "https://") // Remove https
	parsedUrl = strings.TrimPrefix(url, "http://") // Remove http
	parsedUrl = strings.split(url, "/")[0] // Remove path
	parsedUrl = strings.split(url, ":")[0] // Remove port

	// check if parsedUrl is present in icons
	var icons string
	if _, ok := ...

	return icons
}