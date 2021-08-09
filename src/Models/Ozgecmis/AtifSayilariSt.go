package Ozgecmis

//My Service Results
type GetirAtifSayilariSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirAtifSayilariResponse
}

type GetirAtifSayilariSrvcRequest struct {
	TcKimlikNo string
	Yil        string
}

//SOAP Service Results
type GetirAtifSayilariResponse struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	List  []AtifSayilari `xml:"docAtifSayilari"`
	Sonuc Sonuc          `xml:"Sonuc"`
}

type AtifSayilari struct {
	A_ID              string `xml:"A_ID"`
	DONEM             string `xml:"DONEM"`
	TUR               string `xml:"TUR"`
	TUR_AD            string `xml:"TUR_AD"`
	ESER_TURU         string `xml:"ESER_TURU"`
	ESER_ID           string `xml:"ESER_ID"`
	SSCI_INDEKS_ATF   string `xml:"SSCI_INDEKS_ATF"`
	TESV_PUAN         string `xml:"TESV_PUAN"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
}
