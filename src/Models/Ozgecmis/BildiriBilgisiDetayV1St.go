package Ozgecmis

//My Service Results
type GetBildiriBilgisiDetayV1SrvcResult struct {
	Message string
	Result  bool
	Data    *GetBildiriBilgisiDetayV1Response
}
type GetBildiriBilgisiDetayV1SrvcRequest struct {
	TcKimlikNo string
	EserId     string
}

//SOAP Service Results
type GetBildiriBilgisiDetayV1Response struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	List  []BildiriBilgisiDetay `xml:"bildiriListe"`
	Sonuc Sonuc                 `xml:"Sonuc"`
}

type BildiriBilgisiDetay struct {
	YAYINID               string `xml:"YAYIN_ID"`
	KAPSAM_ID             string `xml:"KAPSAM_ID"`
	KAPSAM_AD             string `xml:"KAPSAM_AD"`
	BILDIRI_TUR_ID        string `xml:"BILDIRI_TUR_ID"`
	BILDIRI_TUR           string `xml:"BILDIRI_TUR"`
	BILDIRI_ADI           string `xml:"BILDIRI_ADI"`
	YAZAR_ADI             string `xml:"YAZAR_ADI"`
	YAZAR_SAYISI          string `xml:"YAZAR_SAYISI"`
	ETKINLIK_ADI          string `xml:"ETKINLIK_ADI"`
	YAYIN_DILI            string `xml:"YAYIN_DILI"`
	ETKINLIK_BAS_TARIHI   string `xml:"ETKINLIK_BAS_TARIHI"`
	ETKINLIK_BIT_TARIHI   string `xml:"ETKINLIK_BIT_TARIHI"`
	YAYIN_DURUMU          string `xml:"YAYIN_DURUMU"`
	YAYIN_DURUMU_AD       string `xml:"YAYIN_DURUMU_AD"`
	BASIM_TARIHI          string `xml:"BASIM_TARIHI"`
	BASIM_TURU            string `xml:"BASIM_TURU"`
	BASIM_TURU_AD         string `xml:"BASIM_TURU_AD"`
	ALAN_BILGISI          string `xml:"ALAN_BILGISI"`
	OZEL_SAYI             string `xml:"OZEL_SAYI"`
	YAZAR_ID              string `xml:"YAZAR_ID"`
	GUNCELLEME_TARIHI     string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF           string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD        string `xml:"AKTIF_PASIF_AD"`
	BILDIRI_SUNUM_TURU    string `xml:"BILDIRI_SUNUM_TURU"`
	BILDIRI_SUNUM_TURU_AD string `xml:"BILDIRI_SUNUM_TURU_AD"`
}
