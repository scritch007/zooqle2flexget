package main

//ZXML RSS xml parsing
type ZXML struct {
	Channel ZChannel `xml:"channel"`
}

//ZChannel channel parsing
type ZChannel struct {
	Items []ZItem `xml:"item"`
}

//ZItem item Parsing
type ZItem struct {
	Title string `xml:"title"`
	GUID  ZGUID  `xml:"enclosure"`
	Link  string `xml:"link"`
}

//ZGUID URL of the GUID
type ZGUID struct {
	URL string `xml:"url,attr"`
}
