package Bap

//My Service Results
type DeleteBapProjeSrvcResult struct {
	Message string
	Result  bool
	Data    *InsertOrUpdateBapProjeResponse
}

type DeleteProjeSrvcRequest struct {
	TcKimlikNo  string
	SilinmeYeri string // "1"-> ProjeNo, "2" -> Çalışan
	ProjeNo     string
	ProjeEkipId string
}

//SOAP Service Results
type DeleteBapProjeResponse struct {
	Sonuc SonucBap `xml:"sonuc"`
}
