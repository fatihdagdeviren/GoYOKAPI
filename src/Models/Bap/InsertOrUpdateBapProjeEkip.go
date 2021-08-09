package Bap

//My Service Results
type InsertOrUpdateBapProjeEkipSrvcResult struct {
	Message string
	Result  bool
	Data    *InsertOrUpdateBapProjeEkipResponse
}

type InsertOrUpdateBapProjeEkipSrvcRequest struct {
	TcKimlikNo      string
	ProjeEkipId     string
	YokProjeId      string
	ArastirmaciTip  string
	ArastirmaciTc   string
	ProjedekiGorevi string
	PersonelAd      string
	PersonelSoyad   string
}

//SOAP Service Results
type InsertOrUpdateBapProjeEkipResponse struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	PROJE_EKIP_ID string   `xml:"PROJE_EKIP_ID"`
	Sonuc         SonucBap `xml:"sonuc"`
}
