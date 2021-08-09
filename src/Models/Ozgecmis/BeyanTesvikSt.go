package Ozgecmis

//My Service Results
type GetirBeyanTesvikSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirBeyanTesvikResponse
}

type GetirBeyanTesvikSrvcRequest struct {
	TcKimlikNo string
	Yil        string
}

//SOAP Service Results
type GetirBeyanTesvikResponse struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	List  []BeyanTesvik `xml:"docTesvikBeyan"`
	Sonuc Sonuc         `xml:"Sonuc"`
}

type BeyanTesvik struct {
	FB_ID             string `xml:"FB_ID"`
	TUR               string `xml:"TUR"`
	TUR_AD            string `xml:"TUR_AD"`
	ESER_TUR          string `xml:"ESER_TUR"`
	ESER_TUR_AD       string `xml:"ESER_TUR_AD"`
	ESER_ID           string `xml:"ESER_ID"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	TESVIK_KURAL_ID   string `xml:"TESVIK_KURAL_ID"`
}
