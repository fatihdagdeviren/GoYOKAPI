package Ozgecmis

//My Service Results
type EditorlukBilgisiV1SrvcResult struct {
	Message string
	Result  bool
	Data    *EditorlukBilgisiV1Response
}
type EditorlukBilgisiV1SrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type EditorlukBilgisiV1Response struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	List  []EditorlukListe `xml:"editorlukListe"`
	Sonuc Sonuc            `xml:"Sonuc"`
}

type EditorlukListe struct {
	YAYINID           string `xml:"YAYIN_ID"`
	KAPSAMID          string `xml:"KAPSAM_ID"`
	KAPSAMAD          string `xml:"KAPSAM_AD"`
	EDITORLUKTUR      string `xml:"EDITORLUK_TUR"`
	EDITORLUKTURAD    string `xml:"EDITORLUK_TUR_AD"`
	EDITORGOREV       string `xml:"EDITOR_GOREV"`
	EDITORGOREVAD     string `xml:"EDITOR_GOREV_AD"`
	YAYINADI          string `xml:"YAYIN_ADI"`
	YAZARADI          string `xml:"YAZAR_ADI"`
	YAZARSAYISI       string `xml:"YAZAR_SAYISI"`
	YAYIN_EVI         string `xml:"YAYIN_EVI"`
	YAYIN_DILI        string `xml:"YAYIN_DILI"`
	YAYIN_DILI_ADI    string `xml:"YAYIN_DILI_ADI"`
	ENDEKS_ID         string `xml:"ENDEKS_ID"`
	ENDEKS            string `xml:"ENDEKS"`
	YIL               string `xml:"YIL"`
	BAS_TARIH         string `xml:"BAS_TARIH"`
	ERISIM_TURU       string `xml:"ERISIM_TURU"`
	ERISIM_TURU_AD    string `xml:"ERISIM_TURU_AD"`
	ERISIM_LINKI      string `xml:"ERISIM_LINKI"`
	ALAN_BILGISI      string `xml:"ALAN_BILGISI"`
	YAZAR_ID          string `xml:"YAZAR_ID"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIFPASIF        string `xml:"AKTIF_PASIF"`
	AKTIFPASIFAD      string `xml:"AKTIF_PASIF_AD"`
}
