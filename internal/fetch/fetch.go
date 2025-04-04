package fetch

import (
	"fmt"
	"io"
	"net/http"
)

func Fetch() {
	url := "https://api.weatherapi.com/v1/current.json?key=c295957f530146078c2125310250404&q=Tokyo"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}
	fmt.Println(string(body))
}
