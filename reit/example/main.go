package main

import (
	"financial/percent"
	"fmt"
	"time"
)

const (
	accessToken       string = "6df3f545-343f-47cd-85bc-a86db48a67c7"
	domain            string = "https://mcassessor.maricopa.gov"
	propertyPath      string = "/search/property/"
	rentalPath        string = "/search/rental/"
	parcelDetailsPath string = "/parcel/"
	defaultTime              = 30 * time.Second
)

func main() {
	//req, _ := http.NewRequest(http.MethodGet, domain+rentalPath, nil)
	//q := req.URL.Query()
	//q.Add("q", "85041")
	//req.URL.RawQuery = q.Encode()
	//req.Header.Set("Authorization", accessToken)
	//c := &http.Client{
	//	Timeout: defaultTime,
	//}
	//fmt.Println(req.URL.String())
	//res, _ := c.Do(req)
	//body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))

	//req, _ := http.NewRequest(http.MethodGet, domain+parcelDetailsPath+"10565407", nil)
	//q := req.URL.Query()
	//q.Add("q", "85041")
	//req.URL.RawQuery = q.Encode()
	//req.Header.Set("Authorization", accessToken)
	//c := &http.Client{
	//	Timeout: defaultTime,
	//}
	//fmt.Println(req.URL.String())
	//res, _ := c.Do(req)
	//body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	//d := percent.newDecimal("20%")
	//d2 := decimal.NewFromInt(100)
	//fmt.Println(d.Mul(d2).String())

	p := percent.New("20%")
	fmt.Println(p.String())
	fmt.Println(p.Decimal().String())
}
