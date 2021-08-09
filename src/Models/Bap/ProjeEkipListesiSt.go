package Bap

//My Service Results
type GetBapProjeEkipListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetBapProjeEkipListesiResponse
}
type GetBapProjeEkipListesiSrvcRequest struct {
	TcKimlikNo string
	YOKPROJEID string
}

//SOAP Service Results
type GetBapProjeEkipListesiResponse struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	ProjeEkipListesi []ProjeEkipListesi `xml:"bapProjeEkipListesi"`
	Sonuc            SonucBap           `xml:"sonuc"`
}

type ProjeEkipListesi struct {
	PROJE_EKIP_ID     string `xml:"PROJE_EKIP_ID"`
	YOKPROJEID        string `xml:"YOKPROJEID"`
	ARASTIRMACI_TIP   string `xml:"ARASTIRMACI_TIP"`
	ARASTIRMACI_TC    string `xml:"ARASTIRMACI_TC"`
	PROJEDEKI_GOREVI  string `xml:"PROJEDEKI_GOREVI"`
	EKLEYEN           string `xml:"EKLEYEN"`
	EKLEME_TARIHI     string `xml:"EKLEME_TARIHI"`
	GUNCELLEYEN       string `xml:"GUNCELLEYEN"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	PERSONEL_AD       string `xml:"PERSONEL_AD"`
	PERSONEL_SOYAD    string `xml:"PERSONEL_SOYAD"`
}
