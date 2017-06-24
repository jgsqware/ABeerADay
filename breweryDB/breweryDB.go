package breweryDB

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Beer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	NameDisplay string `json:"nameDisplay"`
	Description string `json:"description"`
	Abv         string `json:"abv"`
	GlasswareID int    `json:"glasswareId"`
	AvailableID int    `json:"availableId"`
	StyleID     int    `json:"styleId"`
	IsOrganic   string `json:"isOrganic"`
	Labels      struct {
		Icon   string `json:"icon"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"labels"`
	Status        string `json:"status"`
	StatusDisplay string `json:"statusDisplay"`
	CreateDate    string `json:"createDate"`
	UpdateDate    string `json:"updateDate"`
	Glass         struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		CreateDate string `json:"createDate"`
	} `json:"glass"`
	Available struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"available"`
	Style struct {
		ID         int `json:"id"`
		CategoryID int `json:"categoryId"`
		Category   struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			CreateDate string `json:"createDate"`
		} `json:"category"`
		Name        string `json:"name"`
		ShortName   string `json:"shortName"`
		Description string `json:"description"`
		IbuMin      string `json:"ibuMin"`
		IbuMax      string `json:"ibuMax"`
		AbvMin      string `json:"abvMin"`
		AbvMax      string `json:"abvMax"`
		SrmMin      string `json:"srmMin"`
		SrmMax      string `json:"srmMax"`
		OgMin       string `json:"ogMin"`
		FgMin       string `json:"fgMin"`
		FgMax       string `json:"fgMax"`
		CreateDate  string `json:"createDate"`
		UpdateDate  string `json:"updateDate"`
	} `json:"style"`
}

type Request struct {
	CurrentPage   int    `json:"currentPage"`
	NumberOfPages int    `json:"numberOfPages"`
	TotalResults  int    `json:"totalResults"`
	Data          []Beer `json:"data"`
	Status        string `json:"status"`
}

type Client struct {
	key string
}

func NewClient(key string) Client {
	return Client{key}
}

func (c Client) RandomBeer() (Beer, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.brewerydb.com/v2/beers?abv=8&order=random&randomCount=1&key=%s&format=json", c.key))

	if err != nil {
		return Beer{}, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Beer{}, err
	}

	var req Request
	err = json.Unmarshal(b, &req)

	if err != nil {
		return Beer{}, err
	}
	return req.Data[0], nil
}
