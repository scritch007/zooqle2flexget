package main

import "encoding/xml"

//FItem export Item information. Only the minimum subset is exported
type FItem struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	GUID  string `xml:"guid"`
}

//FChannel Channel xml parsing
type FChannel struct {
	Items []FItem `xml:"item"`
}

//FXML RSS xml parsing
type FXML struct {
	XMLName xml.Name `xml:"rss"`
	Channel FChannel `xml:"channel"`
}
