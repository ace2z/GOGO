package CUSTOM_GO_MODULE

import (    
	"os"
	//. "local/_CORE"    
    //. "local/_DATABASE"
    . "github.com/ace2z/GOGO/Gadgets"
    //. "github.com/ace2z/GOGO/Gadgets/Mongolian"

    "github.com/PuerkitoBio/goquery"
)


func SCRAPE_from_FILE(filename string) (bool, *goquery.Document) {
	// Read from FILE

	var doc *goquery.Document

	f, err := os.Open(filename)
	if err != nil {
		M.Println("Error in SCRAPE_from_FILE: ")
		W.Println(err)
		return false, doc
	}	
	defer f.Close()


	doc, err = goquery.NewDocumentFromReader(f)
	if err != nil {
		M.Println("Error in SCRAPE_from_FILE DOCREADER: ")
		W.Println(err)
		return false, doc
	}

	return true, doc
}