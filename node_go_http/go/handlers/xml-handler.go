package handlers

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

func XmlHandler(name string) func(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("../common/xml/" + name + ".xml")
	if err != nil {
		panic("cant open xml file")
	}

	return pickHandler(name, file)
}

func pickHandler(name string, file []byte) func(w http.ResponseWriter, r *http.Request) {
	if name == "tiny" {
		var obj TinyXml
		return func(w http.ResponseWriter, r *http.Request) {
			xml.Unmarshal(file, &obj)
			fmt.Fprint(w, obj)
		}
	}

	if name == "small" {
		var obj SmallXml
		return func(w http.ResponseWriter, r *http.Request) {
			xml.Unmarshal(file, &obj)
			fmt.Fprint(w, obj)
		}
	}

	if name == "medium" {
		var obj MediumXml
		return func(w http.ResponseWriter, r *http.Request) {
			xml.Unmarshal(file, &obj)
			fmt.Fprint(w, obj)
		}
	}

	var obj LargeXml
	return func(w http.ResponseWriter, r *http.Request) {
		xml.Unmarshal(file, &obj)
		fmt.Fprint(w, obj)
	}
}

type TinyXml struct {
	XMLName xml.Name `xml:"note"`
	To      string   `xml:"to"`
	From    string   `xml:"from"`
	Heading string   `xml:"heading"`
	Body    string   `xml:"body"`
}

type SmallXml struct {
	XMLName xml.Name `xml:"breakfast_menu"`
	Food    []Food   `xml:"food"`
}

type Food struct {
	Name        string `xml:"name"`
	Price       string `xml:"price"`
	Description string `xml:"description"`
	Calories    string `xml:"calories"`
}

type MediumXml struct {
	XMLName xml.Name `xml:"CATALOG"`
	Plants  []Plant  `xml:"PLANT"`
}

type Plant struct {
	COMMON       string `xml:"COMMON"`
	BOTANICAL    string `xml:"BOTANICAL"`
	ZONE         string `xml:"ZONE"`
	LIGHT        string `xml:"LIGHT"`
	PRICE        string `xml:"PRICE"`
	AVAILABILITY string `xml:"AVAILABILITY"`
}

type LargeXml struct {
	XMLName  xml.Name `xml:"LARGE"`
	Users    Users    `xml:"USERS"`
	Catalogs Mediums  `xml:"CATALOGS"`
}

type Mediums struct {
	XMLName xml.Name    `xml:"CATALOGS"`
	Mediums []MediumXml `xml:"CATALOG"`
}
type Users struct {
	XMLName xml.Name `xml:"USERS"`
	Users   []User   `xml:"USER"`
}
type User struct {
	XMLName xml.Name `xml:"USER"`
	Name    string   `xml:"NAME"`
	Surname string   `xml:"SURNAME"`
	Info    Info     `xml:"INFO"`
}

type Info struct {
	XMLName  xml.Name `xml:"INFO"`
	Personal Personal `xml:"PERSONAL"`
	Public   Public   `xml:"PUBLIC"`
}

type Personal struct {
	XMLName xml.Name `xml:"PERSONAL"`
	WHO     string   `xml:"WHO"`
	WHAT    string   `xml:"WHAT"`
	WHEN    string   `xml:"WHEN"`
}
type Public struct {
	XMLName xml.Name `xml:"PUBLIC"`
	WHO     string   `xml:"WHO"`
	WHAT    string   `xml:"WHAT"`
	WHEN    string   `xml:"WHEN"`
}
