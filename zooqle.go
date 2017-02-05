package main

type ZXML struct {
	Channel ZChannel `xml:"channel"`
}

type ZChannel struct {
	Items []ZItem `xml:"item"`
}

type ZItem struct {
	Title string `xml:"title"`
	Link  ZLink  `xml:"enclosure"`
}

type ZLink struct {
	URL string `xml:"url,attr"`
}
