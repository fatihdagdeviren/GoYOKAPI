package Ozgecmis

//My Service Results
type HakemlikBilgisiV1SrvcResult struct {
	Message string
	Result  bool
	Data    *HakemlikBilgisiV1Response
}

type HakemlikBilgisiV1SrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type HakemlikBilgisiV1Response struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	List  []HakemlikListe `xml:"hakemlikListe"`
	Sonuc Sonuc           `xml:"Sonuc"`
}

type HakemlikListe struct {
	Text             string `xml:",chardata"`
	YAYINID          string `xml:"YAYIN_ID"`
	KAPSAMID         string `xml:"KAPSAM_ID"`
	KAPSAMAD         string `xml:"KAPSAM_AD"`
	HAKEMLIKTURU     string `xml:"HAKEMLIK_TURU"`
	HAKEMLIKTURUAD   string `xml:"HAKEMLIK_TURU_AD"`
	YAYINYERI        string `xml:"YAYIN_YERI"`
	HAKEMLIKSAYISI   string `xml:"HAKEMLIK_SAYISI"`
	YAYINDILI        string `xml:"YAYIN_DILI"`
	YAYINDILIADI     string `xml:"YAYIN_DILI_ADI"`
	ENDEKSID         string `xml:"ENDEKS_ID"`
	ENDEKS           string `xml:"ENDEKS"`
	ALANBILGISI      string `xml:"ALAN_BILGISI"`
	GUNCELLEMETARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIFPASIF       string `xml:"AKTIF_PASIF"`
	AKTIFPASIFAD     string `xml:"AKTIF_PASIF_AD"`
	YIL              string `xml:"YIL"`
	TESVPUAN         string `xml:"TESV_PUAN"`
}
