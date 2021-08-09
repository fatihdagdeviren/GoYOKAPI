package Bap

//My Service Results
type InsertOrUpdateBapProjeSrvcResult struct {
	Message string
	Result  bool
	Data    *InsertOrUpdateBapProjeResponse
}

type InsertOrUpdateBapProjeSrvcRequest struct {
	TcKimlikNo  string
	YokProjeId  string
	ProjeNo     string
	ProjeAdi    string
	ProjeAlani  string
	ProjeBasTar string
	ProjeBitTar string
	ParaBirimi  string
	Butce       string
}

//SOAP Service Results
type InsertOrUpdateBapProjeResponse struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	PROJE_ID string   `xml:"PROJE_ID"`
	Sonuc    SonucBap `xml:"sonuc"`
}
