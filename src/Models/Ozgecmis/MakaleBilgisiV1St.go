package Ozgecmis

//My Service Results
type GetMakaleBilgisiV1SrvcResult struct {
	Message string
	Result  bool
	Data    *GetMakaleBilgisiV1Response
}
type GetMakaleBilgisiV1SrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetMakaleBilgisiV1Response struct {
	List  []MakaleListe `xml:"makaleListe"`
	Sonuc Sonuc         `xml:"Sonuc"`
}

type MakaleListe struct {
	YAYIN_ID          string `xml:"YAYIN_ID"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
}
