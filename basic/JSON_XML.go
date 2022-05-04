package basic

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
)

// JavaScrip Object Notation
func JSONExample() {
	// json unmarshal example
	unmarshalPerson()
	// json marshal example
	marshalPerson()
	// json marshal with "omitempty" and "-"(skip) directive
	marshalPersonWithoutAge()
	// func to parse struct tags
	parseStructTags()
}

// struct tags are string literals which help to docode JSON correctly
type MyJson struct {
	Person `json:"person"`
}
type Person struct {
	Name string `json:"name"`
	Age uint `json:"age"`
}

func unmarshalPerson() {
	myJson := []byte(`{"person":{"name":"Jared","age":40}}`)
	c := MyJson{}
	err := json.Unmarshal(myJson, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Person.Name)
	fmt.Println(c.Person.Age)
}

type Person2 struct {
	ID uint `json:"id"`
	Name string	`json:"name"`
	Age uint `json:"age"`
	AdInf AdditionalInfo `json:"additionalInfo"`
}

type AdditionalInfo struct {
	Gender string `json:"gender"`
	Occupation string `json:"occupation"`
}

func marshalPerson() {
	p := Person2{ID: 50, Name: "Rick", Age: 60, AdInf: AdditionalInfo{Gender: "male", Occupation: "scientist"}}
	b, err := json.MarshalIndent(p, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

type Person3 struct {
	Name string `json:"name"`
	Age uint `json:"age,omitempty"`
	AdInfo AdditionalInfo `json:"-"`
}

func marshalPersonWithoutAge() {
	p := Person3{Name: "Morty"}
	pM, err := json.MarshalIndent(p, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pM))
}

func parseStructTags() {
	p := Person3{Name: "Dude"}
	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("field Name %s\n", t.Field(i).Name)
		fmt.Printf("field Tag %s\n", t.Field(i).Tag)
	}
}

// Xtensible Markup Language
func XMLExample() {
	// unmarshal xml
	unmarshalPersonXML()
	// marshal xml
	marshalPersonXML()
	// marshal xml with special tags
	marshalXMLWithAttr()
}

type MyXML struct {
	Person4 `xml:"person"`
}

type Person4 struct {
	// special xml.Name field
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age uint `xml:"age"`
}

func unmarshalPersonXML() {
	myXML := []byte(`<person>
	<name>Dude</name>
	<age>30</age>
	</person>`)
	p := MyXML{}
	err := xml.Unmarshal(myXML, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Person4.Name)
	fmt.Println(p.Person4.Age)
}

func marshalPersonXML() {
	p := Person4{Name: "Brad", Age: 50}
	pI, err := xml.MarshalIndent(p, "", "	")
	if err != nil {
		panic(err)
	}
	// it's required to add xml header
	// <?xml version="1.0" encoding="UTF-8"?> using xml.Header
	xmlWithHeader := xml.Header + string(pI)
	fmt.Println(xmlWithHeader)
}

type Price struct {
	// xml attributes
	Text string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
}

type Product struct {
	Comment string `xml:",comment"`
	Price Price `xml:"price"`
	Name string `xml:"wrap>name"`
	MoreXML MoreXML `xml:"moreXML"`
}

type MoreXML struct {
	Text string `xml:",cdata"`
}

func marshalXMLWithAttr() {
	prod := Product{
		Comment: "my comment",
		Price: Price{
			Text: "100",
			Currency: "USD",
		},
		Name: "some name",
		MoreXML: MoreXML{
			Text: "123",
		},
	}
	prodM, err := xml.MarshalIndent(prod, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prodM)) 
}