package Ozgecmis

//My Service Results
type GetOdulListesiV1SrvcResult struct {
	Message string
	Result  bool
	Data    *GetOdulListesiV1Response
}
type GetOdulListesiV1SrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetOdulListesiV1Response struct {
	List  []OdulListesi `xml:"odulListesi"`
	Sonuc Sonuc         `xml:"Sonuc"`
}

type OdulListesi struct {
	ODUL_ID           string `xml:"ODUL_ID"`
	ODUL_ADI          string `xml:"ODUL_ADI"`
	ODUL_TARIH        string `xml:"ODUL_TARIH"`
	FAAL_DETAY_ID     string `xml:"FAAL_DETAY_ID"`
	FAAL_DETAY_ADI    string `xml:"FAAL_DETAY_ADI"`
	ODUL_TUR_ID       string `xml:"ODUL_TUR_ID"`
	ODUL_TURU         string `xml:"ODUL_TURU"`
	KURULUS_ADI       string `xml:"KURULUS_ADI"`
	P_UNVAN_ID        string `xml:"P_UNVAN_ID"`
	P_UNVAN_AD        string `xml:"P_UNVAN_AD"`
	P_KURUM_ID        string `xml:"P_KURUM_ID"`
	P_KURUM_AD        string `xml:"P_KURUM_AD"`
	ISYERI_TURU_ID    string `xml:"ISYERI_TURU_ID"`
	ISYERI_TURU_ADI   string `xml:"ISYERI_TURU_ADI"`
	ULKE_ID           string `xml:"ULKE_ID"`
	ULKE_AD           string `xml:"ULKE_AD"`
	KISI_SAYISI       string `xml:"KISI_SAYISI"`
	ODUL_KISI_SIRA    string `xml:"ODUL_KISI_SIRA"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
}
