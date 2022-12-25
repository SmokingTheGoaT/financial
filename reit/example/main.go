package main

import (
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

	//p := percent.New("20%")
	//fmt.Println(p.String())
	//fmt.Println(p.Decimal().String())

	//fmt.Println(currency.USD.Amount(200).Currency().String())

	//m := map[string]string{
	//	"$": "",
	//}
	//c := currency.New("$20", unit.USD, utils.RemoveStrings(m))
	//fmt.Println(c.String())
	//c2 := currency.New("$25", unit.USD, utils.RemoveStrings(m))
	//c = c.Add(c2)
	//fmt.Println(c.String())

	//d := tvm.fvif(percent.New("6%"), types.Yearly.Term(2))
	//fmt.Println(d.String())

	//pmt := tvm.Pmt(percent.New(0.003125, 1.0), types.Monthly.Term(300),
	//	currency.New("$56000", currencyUnit.USD, utils.RemoveStrings(m)).Decimal(),
	//	decimal.NewFromInt(0), types.EndOfPeriod)
	//
	//fmt.Println(pmt.String())
}
