package Ozgecmis

//My Service Results
type GetKitapBilgisiDetayV1SrvcResult struct {
	Message string
	Result  bool
	Data    *GetKitapBilgisiDetayV1Response
}

type GetKitapBilgisiDetayV1SrvcRequest struct {
	TcKimlikNo string
	KitapId    string
}

//SOAP Service Results
type GetKitapBilgisiDetayV1Response struct {
	List  []KitapListeDetay `xml:"kitapListe"`
	Sonuc Sonuc             `xml:"Sonuc"`
}

type KitapListeDetay struct {
	YAYIN_ID          string `xml:"YAYIN_ID"`
	KAPSAM_ID         string `xml:"KAPSAM_ID"`
	KAPSAM_AD         string `xml:"KAPSAM_AD"`
	KITAP_TUR_ID      string `xml:"KITAP_TUR_ID"`
	KITAP_TUR         string `xml:"KITAP_TUR"`
	KITAP_ADI         string `xml:"KITAP_ADI"`
	KATKI_DUZEYI      string `xml:"KATKI_DUZEYI"`
	KATKI_DUZEYI_AD   string `xml:"KATKI_DUZEYI_AD"`
	BOLUM_ADI         string `xml:"BOLUM_ADI"`
	YAZAR_ADI         string `xml:"YAZAR_ADI"`
	YAZAR_SAYISI      string `xml:"YAZAR_SAYISI"`
	EDITOR_ADI        string `xml:"EDITOR_ADI"`
	ULKE              string `xml:"ULKE"`
	ULKE_ADI          string `xml:"ULKE_ADI"`
	SEHIR             string `xml:"SEHIR"`
	YAYIN_EVI         string `xml:"YAYIN_EVI"`
	YAYIN_DILI        string `xml:"YAYIN_DILI"`
	YAYIN_DILI_ADI    string `xml:"YAYIN_DILI_ADI"`
	YIL               string `xml:"YIL"`
	ISBN              string `xml:"ISBN"`
	KACINCI_BASIM     string `xml:"KACINCI_BASIM"`
	SAYFA_SAYISI      string `xml:"SAYFA_SAYISI"`
	BASIM_TURU        string `xml:"BASIM_TURU"`
	BASIM_TURU_AD     string `xml:"BASIM_TURU_AD"`
	ERISIM_LINKI      string `xml:"ERISIM_LINKI"`
	ALAN_BILGISI      string `xml:"ALAN_BILGISI"`
	YAZAR_ID          string `xml:"YAZAR_ID"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
}
