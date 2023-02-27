package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jbattistella/normative/models"
)

func main() {

	url := "https://economiccalendar.p.rapidapi.com/events/1598072400000/1756771140000"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "9cd8d9ff35mshcd24dd50afc9fc4p12258ejsnd0541ec9480b")
	req.Header.Add("X-RapidAPI-Host", "economiccalendar.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var econEvents models.EconomicEvents

	if err = json.NewDecoder(res.Body).Decode(&econEvents); err != nil {
		log.Print(err)
	}
	for _, v := range econEvents.Events {
		fmt.Println(v.Name)
		fmt.Println(v.Datetime)

	}
}
