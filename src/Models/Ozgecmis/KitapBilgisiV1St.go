package Ozgecmis

//My Service Results
type GetKitapBilgisiV1SrvcResult struct {
	Message string
	Result  bool
	Data    *GetKitapBilgisiV1Response
}
type GetKitapBilgisiV1SrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetKitapBilgisiV1Response struct {
	List  []KitapListe `xml:"kitapListe"`
	Sonuc Sonuc        `xml:"Sonuc"`
}

type KitapListe struct {
	YAYIN_ID          string `xml:"YAYIN_ID"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
}
