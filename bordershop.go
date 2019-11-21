package bordershop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CategoryResponse struct {
	Products []struct {
		IsCheapest       bool `json:"isCheapest"`
		LeftBottomSplash bool `json:"leftBottomSplash"`
		Discount         struct {
			ShowPriceForOne bool   `json:"showPriceForOne"`
			IsSmileOffer    bool   `json:"isSmileOffer"`
			DiscountText    string `json:"discountText"`
			TakeMoreText    string `json:"takeMoreText"`
		} `json:"discount,omitempty"`
		LeftSplash struct {
			Type int `json:"type"`
		} `json:"leftSplash"`
		Uom                  string `json:"uom"`
		QtyPrUom             string `json:"qtyPrUom"`
		UnitPriceText1       string `json:"unitPriceText1"`
		UnitPriceText2       string `json:"unitPriceText2"`
		ID                   string `json:"id"`
		Brand                string `json:"brand"`
		ProductClickTracking string `json:"productClickTracking"`
		AddToBasket          struct {
			MinimumQuantityText string `json:"minimumQuantityText"`
			MinimumQuantity     int    `json:"minimumQuantity"`
			ID                  string `json:"id"`
			InitialQuantity     int    `json:"initialQuantity"`
			ListName            string `json:"listName"`
			ListPosition        int    `json:"listPosition"`
			IsShopOnly          bool   `json:"isShopOnly"`
			IsSoldOut           bool   `json:"isSoldOut"`
			ProductID           string `json:"productId"`
		} `json:"addToBasket"`
		Price struct {
			AmountAsDecimal float64 `json:"amountAsDecimal"`
			Amount          string  `json:"amount"`
			Major           string  `json:"major"`
			Minor           string  `json:"minor"`
		} `json:"price"`
		DisplayName string `json:"displayName"`
		Image       string `json:"image"`
		URL         string `json:"url"`
	} `json:"products"`
	Facets []struct {
		DisplayName string `json:"displayName"`
		Key         string `json:"key"`
		Values      []struct {
			Selected    bool   `json:"selected"`
			Key         string `json:"key"`
			Count       int    `json:"count"`
			DisplayName string `json:"displayName"`
		} `json:"values"`
		Priority int `json:"priority"`
	} `json:"facets"`
	ChildCategories []interface{} `json:"childCategories"`
	Total           int           `json:"total"`
}

func GetCategory(CategoryID int64) (CategoryResponse, error) {

	cg := strconv.FormatInt(CategoryID, 10)
	cgParsed := url.QueryEscape(cg)

	url := fmt.Sprintf("https://www.bordershop.com/se/bordershop/api/catalogsearchapi/productsearch?categoryId=%s", cgParsed)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CategoryResponse{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return CategoryResponse{}, err
	}

	defer resp.Body.Close()

	var category CategoryResponse

	if err := json.NewDecoder(resp.Body).Decode(&category); err != nil {
		return CategoryResponse{}, err
	}

	return category, nil
}

func GetMostPopular(Count int64) (CategoryResponse, error) {

	cg := strconv.FormatInt(Count, 10)
	cgParsed := url.QueryEscape(cg)

	url := fmt.Sprintf("https://www.bordershop.com/se/bordershop/api/recommendationapi/getmostpopularrightnow?count=%s", cgParsed)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CategoryResponse{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return CategoryResponse{}, err
	}

	defer resp.Body.Close()

	var category CategoryResponse

	if err := json.NewDecoder(resp.Body).Decode(&category); err != nil {
		return CategoryResponse{}, err
	}

	return category, nil
}
