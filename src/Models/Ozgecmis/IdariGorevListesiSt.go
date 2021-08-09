package Ozgecmis

//My Service Results
type GetIdariGorevListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *IdariGorevListesiResponse
}

type GetIdariGorevListesiSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type IdariGorevListesiResponse struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	List  []IdariGorevListesi `xml:"idariGorevListesi"`
	Sonuc Sonuc               `xml:"Sonuc"`
}

type IdariGorevListesi struct {
	IDGOR_ID          string `xml:"IDGOR_ID"`
	YER_ID            string `xml:"YER_ID"`
	YER_ADI           string `xml:"YER_ADI"`
	ULKE_ID           string `xml:"ULKE_ID"`
	ULKE_ADI          string `xml:"ULKE_ADI"`
	GOREV_ID          string `xml:"GOREV_ID"`
	GOREV_ADI         string `xml:"GOREV_ADI"`
	BAS_TAR           string `xml:"BAS_TAR"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
}
