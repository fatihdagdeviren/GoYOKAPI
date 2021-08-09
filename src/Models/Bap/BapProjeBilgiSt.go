package Bap

//My Service Results
type GetBapProjeBilgiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetBapProjeBilgiResponse
}
type GetBapProjeBilgiSrvcRequest struct {
	TcKimlikNo string
	ProjeNo    string
}

//SOAP Service Results
type GetBapProjeBilgiResponse struct {
	//Name         xml.Name       `xml:"getBildiriBilgisiV1Response"`
	AkademikGorevListesi []BapProjeBilgisi `xml:"bapProjeBilgisi"`
	Sonuc                SonucBap          `xml:"sonuc"`
}

type BapProjeBilgisi struct {
	YOK_PROJE_ID      string `xml:"YOK_PROJE_ID"`
	UNV_PROJE_NO      string `xml:"UNV_PROJE_NO"`
	PROJE_ADI         string `xml:"PROJE_ADI"`
	PROJE_ALANI       string `xml:"PROJE_ALANI"`
	PROJE_TURU        string `xml:"PROJE_TURU"`
	PROJE_BAS_TAR     string `xml:"PROJE_BAS_TAR"`
	PROJE_BIT_TAR     string `xml:"PROJE_BIT_TAR"`
	OZET              string `xml:"OZET"`
	ANAHTAR_KELIME    string `xml:"ANAHTAR_KELIME"`
	BIRIM_ID          string `xml:"BIRIM_ID"`
	PARA_BIRIMI       string `xml:"PARA_BIRIMI"`
	BUTCE             string `xml:"BUTCE"`
	EKLEYEN           string `xml:"EKLEYEN"`
	GUNCELLEYEN       string `xml:"GUNCELLEYEN"`
	EKLEME_TARIHI     string `xml:"EKLEME_TARIHI"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
}
