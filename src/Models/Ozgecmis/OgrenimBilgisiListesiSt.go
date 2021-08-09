package Ozgecmis

//My Service Results
type GetirOgrenimBilgisiListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirOgrenimBilgisiListesiResponse
}

type GetirOgrenimBilgisiListesiSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirOgrenimBilgisiListesiResponse struct {
	List  []OgrenimBilgisiListesi `xml:"OgrenimBilgisiListesi"`
	Sonuc Sonuc                   `xml:"Sonuc"`
}

type OgrenimBilgisiListesi struct {
	ID                 string `xml:"ID"`
	YER_ID             string `xml:"YER_ID"`
	YER_AD             string `xml:"YER_AD"`
	BASTAR1            string `xml:"BASTAR1"`
	BITTAR1            string `xml:"BITTAR1"`
	ULKE_ID            string `xml:"ULKE_ID"`
	BIRIM_YOK          string `xml:"BIRIM_YOK"`
	ULKE_AD            string `xml:"ULKE_AD"`
	PROGRAM_ID         string `xml:"PROGRAM_ID"`
	PROGRAM_ADI        string `xml:"PROGRAM_ADI"`
	TEZ_ASAMASI        string `xml:"TEZ_ASAMASI"`
	TEZ_ASAMASI_AD     string `xml:"TEZ_ASAMASI_AD"`
	UNIV_ID            string `xml:"UNIV_ID"`
	UNV_BIRIM_ADI      string `xml:"UNV_BIRIM_ADI"`
	AKADEMIK_BIRIM_ID  string `xml:"AKADEMIK_BIRIM_ID"`
	AKADEMIK_BIRIM_ADI string `xml:"AKADEMIK_BIRIM_ADI"`
	FAKULTEBILGISI     string `xml:"FAKULTEBILGISI"`
	BOLUMBILGISI       string `xml:"BOLUMBILGISI"`
	TEZ_ADI            string `xml:"TEZ_ADI"`
	TEZ_BAS_TAR        string `xml:"TEZ_BAS_TAR"`
	DANISMAN_AD_SOYAD  string `xml:"DANISMAN_AD_SOYAD"`
	AKTIF_PASIF        string `xml:"AKTIF_PASIF"`
	GUNCELLEME_TARIHI  string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF_AD     string `xml:"AKTIF_PASIF_AD"`
}
