package Ozgecmis

//My Service Results
type GetirAkademikGorevListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirAkademikGorevListesiResponse
}
type GetirAkademikGorevListesiSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirAkademikGorevListesiResponse struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	List  []AkademikGorevListesi `xml:"AkademikGorevListesi"`
	Sonuc Sonuc                  `xml:"Sonuc"`
}

type AkademikGorevListesi struct {
	GOREVID          string `xml:"GOREV_ID"`
	YERID            string `xml:"YER_ID"`
	YERAD            string `xml:"YER_AD"`
	BASTAR1          string `xml:"BASTAR1"`
	BILIMALANADI     string `xml:"BILIMALAN_ADI"`
	UZMANLIKALANI    string `xml:"UZMANLIK_ALANI"`
	UZMANLIKALANIAD  string `xml:"UZMANLIK_ALANI_AD"`
	UNIVID           string `xml:"UNIV_ID"`
	UNIVBIRIMADI     string `xml:"UNIV_BIRIM_ADI"`
	BIRIMID          string `xml:"BIRIM_ID"`
	FAKULTEBILGISI   string `xml:"FAKULTEBILGISI"`
	BOLUMBILGISI     string `xml:"BOLUMBILGISI"`
	KADROUNVANID     string `xml:"KADRO_UNVAN_ID"`
	KADROUNVANADI    string `xml:"KADRO_UNVAN_ADI"`
	AKADEMIKDURUM    string `xml:"AKADEMIK_DURUM"`
	AKADEMIKDURUMADI string `xml:"AKADEMIK_DURUM_ADI"`
	AKADEMIKBIRIMADI string `xml:"AKADEMIK_BIRIM_ADI"`
	GUNCELLEMETARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIFPASIF       string `xml:"AKTIF_PASIF"`
	AKTIFPASIFAD     string `xml:"AKTIF_PASIF_AD"`
}
