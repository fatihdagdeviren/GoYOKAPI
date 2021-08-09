package Ozgecmis

//My Service Results
type GetirUyelikListesiResponseSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirUyelikListesiResponse
}
type GetirUyelikListesiResponseSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirUyelikListesiResponse struct {
	//Name         xml.Name       `xml:"getirProjeListesiResponse"`
	List  []UyelikListesi `xml:"uyelikListesi"`
	Sonuc Sonuc           `xml:"Sonuc"`
}

type UyelikListesi struct {
	UYELIK_ID         string `xml:"UYELIK_ID"`
	KURULUS_ADI       string `xml:"KURULUS_ADI"`
	BAS_TAR           string `xml:"BAS_TAR"`
	BIT_TAR           string `xml:"BIT_TAR"`
	UYELIK_DURUMU     string `xml:"UYELIK_DURUMU"`
	UYELIK_DURUMU_AD  string `xml:"UYELIK_DURUMU_AD"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
}
