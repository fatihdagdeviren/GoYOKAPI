package Ozgecmis

//My Service Results
type GetirBasvuruDurumSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirBasvuruDurumResponse
}

type GetirBasvuruDurumSrvcRequest struct {
	TcKimlikNo string
	Yil        string
}

//SOAP Service Results
type GetirBasvuruDurumResponse struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	List  []BasvuruDurum `xml:"basvuruDurum"`
	Sonuc Sonuc          `xml:"Sonuc"`
}

type BasvuruDurum struct {
	BASVURU_ID       string `xml:"BASVURU_ID"`
	DONEM_ID         string `xml:"DONEM_ID"`
	DONEM_AD         string `xml:"DONEM_AD"`
	BASVURU_TARIHI   string `xml:"BASVURU_TARIHI"`
	SON_ISLEM_TARIHI string `xml:"SON_ISLEM_TARIHI"`
}
