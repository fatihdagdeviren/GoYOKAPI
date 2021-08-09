package Ozgecmis

//My Service Results
type GetirSanatsalFaalv1SrvcResult struct {
	Message string
	Result  bool
	Data    *GetirSanatsalFaalv1Response
}

type GetirSanatsalFaalv1SrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirSanatsalFaalv1Response struct {
	List  []SanatsalFaalListesi `xml:"sanatsalFaalListe"`
	Sonuc Sonuc                 `xml:"Sonuc"`
}

type SanatsalFaalListesi struct {
	S_ID              string `xml:"S_ID"`
	ANA_TUR           string `xml:"ANA_TUR"`
	ANATUR_ADI        string `xml:"ANATUR_ADI"`
	KAPSAM            string `xml:"KAPSAM"`
	KAPSAM_AD         string `xml:"KAPSAM_AD"`
	TIP               string `xml:"TIP"`
	TIP_ADI           string `xml:"TIP_ADI"`
	ETKINLIK_ADI      string `xml:"ETKINLIK_ADI"`
	ETKINLIK_YERI     string `xml:"ETKINLIK_YERI"`
	ETKINLIK_TURU     string `xml:"ETKINLIK_TURU"`
	BAS_TARIH         string `xml:"BAS_TARIH"`
	BIT_TARIH         string `xml:"BIT_TARIH"`
	DUZENLEYENLER     string `xml:"DUZENLEYENLER"`
	ULKE              string `xml:"ULKE"`
	ULKE_ADI          string `xml:"ULKE_ADI"`
	SEHIR             string `xml:"SEHIR"`
	ETKINLIK_DILI     string `xml:"ETKINLIK_DILI"`
	DIL_ADI           string `xml:"DIL_ADI"`
	KISI_SAYISI       string `xml:"KISI_SAYISI"`
	KISI_SIRASI       string `xml:"KISI_SIRASI"`
	ETKINLIK_SURESI   string `xml:"ETKINLIK_SURESI"`
	ETKINLIK_TURU_AD  string `xml:"ETKINLIK_TURU_AD"`
	UNVAN_ID          string `xml:"UNVAN_ID"`
	UNVAN_AD          string `xml:"UNVAN_AD"`
	KURUM_ID          string `xml:"KURUM_ID"`
	KURUM_AD          string `xml:"KURUM_AD"`
	TESV_PUAN         string `xml:"TESV_PUAN"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
}
