package Ozgecmis

//My Service Results
type ProjeListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirProjeListesiResponse
}
type ProjeListesiSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirProjeListesiResponse struct {
	//Name         xml.Name       `xml:"getirProjeListesiResponse"`
	List  []ProjeListesi `xml:"projeListe"`
	Sonuc Sonuc          `xml:"Sonuc"`
}

type ProjeListesi struct {
	ProjeID          string `xml:"PROJE_ID"`
	GuncellemeTarihi string `xml:"GUNCELLEME_TARIHI"`
	AktifPasif       string `xml:"AKTIF_PASIF"`
	AktifPasifAd     string `xml:"AKTIF_PASIF_AD"`
}
