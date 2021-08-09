package Ozgecmis

//My Service Results
type GetirYazarListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirYazarListesiResponse
}

type GetirYazarListesiSrvcRequest struct {
	TcKimlikNo string
	YazarId    string
}

//SOAP Service Results
type GetirYazarListesiResponse struct {
	List  []YazarlarListesi `xml:"yazarlarListesi"`
	Sonuc Sonuc             `xml:"Sonuc"`
}

type YazarlarListesi struct {
	Y_ID                string `xml:"Y_ID"`
	ARASTIRMACI_ID      string `xml:"ARASTIRMACI_ID"`
	YAZARAD             string `xml:"YAZARAD"`
	YAZARSOYAD          string `xml:"YAZARSOYAD"`
	GUNCELLEME_TARIHI   string `xml:"GUNCELLEME_TARIHI"`
	YAZAR_SIRA          string `xml:"YAZAR_SIRA"`
	YAZAR_FORM_ID       string `xml:"YAZAR_FORM_ID"`
	KULLANICININ_ARS_ID string `xml:"KULLANICININ_ARS_ID"`
	YAZAR_TUR           string `xml:"YAZAR_TUR"`
	YAZAR_TUR_AD        string `xml:"YAZAR_TUR_AD"`
	YAZAR_ORCID         string `xml:"YAZAR_ORCID"`
}
