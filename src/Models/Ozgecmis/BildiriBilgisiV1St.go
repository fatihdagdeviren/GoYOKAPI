package Ozgecmis

//My Service Results
type GetBildiriBilgisiV1SrvcResult struct {
	Message string
	Result  bool
	Data    *GetBildiriBilgisiV1Response
}
type GetBildiriBilgisiV1SrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetBildiriBilgisiV1Response struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	List  []BildiriBilgisi `xml:"bildiriListe"`
	Sonuc Sonuc            `xml:"Sonuc"`
}

type BildiriBilgisi struct {
	YAYINID          string `xml:"YAYIN_ID"`
	GUNCELLEMETARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIFPASIF       string `xml:"AKTIF_PASIF"`
	AKTIFPASIFAD     string `xml:"AKTIF_PASIF_AD"`
}
