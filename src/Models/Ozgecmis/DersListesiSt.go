package Ozgecmis

//My Service Results
type GetirDersListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirDersListesiResponse
}

type GetirDersListesiSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirDersListesiResponse struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	List  []DersListesi `xml:"dersListesi"`
	Sonuc Sonuc         `xml:"Sonuc"`
}

type DersListesi struct {
	DERS_ID           string `xml:"DERS_ID"`
	DERS_ADI          string `xml:"DERS_ADI"`
	OGRENIM_ID        string `xml:"OGRENIM_ID"`
	OGRENIM_ADI       string `xml:"OGRENIM_ADI"`
	AKADEMIK_YIL_ID   string `xml:"AKADEMIK_YIL_ID"`
	AKADEMIK_YIL      string `xml:"AKADEMIK_YIL"`
	DIL_ID            string `xml:"DIL_ID"`
	DIL_ADI           string `xml:"DIL_ADI"`
	DERS_SAATI        string `xml:"DERS_SAATI"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	EKLEME_TARIHI     string `xml:"EKLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
}
