package main

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
	Channel FChannel `xml:"channel"`
}
