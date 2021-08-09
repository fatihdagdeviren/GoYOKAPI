package Ozgecmis

//My Service Results
type GetMakaleBilgisiDetayV1SrvcResult struct {
	Message string
	Result  bool
	Data    *GetMakaleBilgisiDetayV1Response
}

type GetMakaleBilgisiDetayV1SrvcRequest struct {
	TcKimlikNo string
	MakaleId   string
}

//SOAP Service Results
type GetMakaleBilgisiDetayV1Response struct {
	List  []MakaleListeDetay `xml:"makaleListesiDetay"`
	Sonuc Sonuc              `xml:"sonuc"`
}

type MakaleListeDetay struct {
	YAYIN_ID          string `xml:"YAYIN_ID"`
	KAPSAM_ID         string `xml:"KAPSAM_ID"`
	KAPSAM_AD         string `xml:"KAPSAM_AD"`
	HAKEM_TUR         string `xml:"HAKEM_TUR"`
	HAKEM_TUR_AD      string `xml:"HAKEM_TUR_AD"`
	ENDEKS_ID         string `xml:"ENDEKS_ID"`
	ENDEKS            string `xml:"ENDEKS"`
	MAKALE_ADI        string `xml:"MAKALE_ADI"`
	YAZAR_ADI         string `xml:"YAZAR_ADI"`
	YAZAR_SAYISI      string `xml:"YAZAR_SAYISI"`
	DERGI_ADI         string `xml:"DERGI_ADI"`
	YAYIN_DILI        string `xml:"YAYIN_DILI"`
	YAYIN_DILI_ADI    string `xml:"YAYIN_DILI_ADI"`
	AY                string `xml:"AY"`
	YIL               string `xml:"YIL"`
	ISSN              string `xml:"ISSN"`
	ERISIM_TURU       string `xml:"ERISIM_TURU"`
	ERISIM_TURU_AD    string `xml:"ERISIM_TURU_AD"`
	ERISIM_LINKI      string `xml:"ERISIM_LINKI"`
	ALAN_BILGISI      string `xml:"ALAN_BILGISI"`
	OZEL_SAYI         string `xml:"OZEL_SAYI"`
	YAZAR_ID          string `xml:"YAZAR_ID"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
	MAKALE_TURU_ID    string `xml:"MAKALE_TURU_ID"`
	MAKALE_TURU_AD    string `xml:"MAKALE_TURU_AD"`
}
