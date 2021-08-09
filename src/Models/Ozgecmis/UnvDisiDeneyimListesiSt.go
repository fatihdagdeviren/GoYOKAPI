package Ozgecmis

//My Service Results
type GetirUnvDisiDeneyimListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirUnvDisiDeneyimListesiResponse
}

type GetirUnvDisiDeneyimListesiSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirUnvDisiDeneyimListesiResponse struct {
	//Name         xml.Name       `xml:"getirProjeListesiResponse"`
	List  []DeneyimListesi `xml:"deneyimListesi"`
	Sonuc Sonuc            `xml:"Sonuc"`
}

type DeneyimListesi struct {
	DENEYIM_ID        string `xml:"DENEYIM_ID"`
	KURULUS_ADI       string `xml:"KURULUS_ADI"`
	GOREV_ADI         string `xml:"GOREV_ADI"`
	BAS_TAR           string `xml:"BAS_TAR"`
	BIT_TAR           string `xml:"BIT_TAR"`
	IS_TANIMI         string `xml:"IS_TANIMI"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
}
