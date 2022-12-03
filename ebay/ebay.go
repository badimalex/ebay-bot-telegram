package ebay

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHtml(url string) *http.Response {
  response, error := http.Get(url)
	if error != nil {
		fmt.Println(error)
	}

	if response.StatusCode > 400 {
		fmt.Println("Status code: ", response.StatusCode)
	}

	return response;
}

func scrapePageData(doc *goquery.Document) {
	doc.Find("ul.srp-results>li.s-item").Each(func(index int, item *goquery.Selection) {
		a := item.Find("a.s-item__link")
		img := item.Find("img.s-item__image-img")
		div := item.Find("div.s-item__title")

		title := strings.TrimSpace(div.Text())
		url, _ := a.Attr("href")
		cover, _ := img.Attr("src")

		price_span := item.Find("span.s-item__price").Text();

		fmt.Println(cover, price_span, title, url);
	})
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func InitSearch() {
	url := "https://www.ebay.it/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw=hohner+verdi&_sacat=0"

	response := getHtml(url)

	defer response.Body.Close()

	doc, erra := goquery.NewDocumentFromReader(response.Body);

	check(erra)

	scrapePageData(doc)
}
