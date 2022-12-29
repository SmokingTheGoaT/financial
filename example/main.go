package main

import (
	"fmt"
	"math"
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
	//fmt.Println(req.URL.Percent())
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
	//fmt.Println(req.URL.Percent())
	//res, _ := c.Do(req)
	//body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	//d := percent.newDecimal("20%")
	//d2 := decimal.NewFromInt(100)
	//fmt.Println(d.Mul(d2).Percent())

	//p := percent.New("20%")
	//fmt.Println(p.Percent())
	//fmt.Println(p.Decimal().Percent())

	//fmt.Println(currency.USD.Amount(200).Currency().Percent())

	//m := map[string]string{
	//	"$": "",
	//}

	//c := currency.New("$20", unit.USD, utils.RemoveStrings(m))
	//fmt.Println(c.Percent())
	//c2 := currency.New("$25", unit.USD, utils.RemoveStrings(m))
	//c = c.Add(c2)
	//fmt.Println(c.Percent())

	//d := tvm.fvif(percent.New("6%"), types.Yearly.Term(2))
	//fmt.Println(d.Percent())

	//pmt := tvm.Pmt(percent.New(0.003125, 1.0), types.Monthly.Term(300),
	//	currency.New("$56000", currencyUnit.USD, utils.RemoveStrings(m)).Decimal(),
	//	decimal.NewFromInt(0), types.EndOfPeriod)

	//fmt.Println(pmt.Percent())
	//
	//pmt, err := financial.PMT(types.New(0.003125, 1.0), decimal.NewFromInt(300),
	//	decimal.NewFromInt(56000), decimal.NewFromInt(0), types.EndOfPeriod)
	//fmt.Println(err)
	//fmt.Println(pmt)

	init := 3
	end := 10
	count := int(math.Abs(float64(end) + 1))
	re := []int{}
	for i := init; i < count; i++ {
		switch {
		case i == init:
			re = append(re, init)
		case i == end:
			re = append(re, end)
		default:
			re = append(re, i)
		}
	}
	fmt.Println(re)
}
