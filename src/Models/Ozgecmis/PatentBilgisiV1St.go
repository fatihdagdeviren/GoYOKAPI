package Ozgecmis

//My Service Results
type GetirPatentBilgisiV1ListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirPatentBilgisiListesiV1Response
}
type GetirPatentBilgisiV1ListesiSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirPatentBilgisiListesiV1Response struct {
	List  []PatentBilgisiListesi `xml:"patentListe"`
	Sonuc Sonuc                  `xml:"Sonuc"`
}

type PatentBilgisiListesi struct {
	PATENT_ID         string `xml:"PATENT_ID"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
}
