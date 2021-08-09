package Ozgecmis

//My Service Results
type ProjeListesiDetaySrvcResult struct {
	Message string
	Result  bool
	Data    *GetirProjeListesiDetayResponse
}
type ProjeListesiDetaySrvcRequest struct {
	TcKimlikNo string
	ProjeId    string
}

//SOAP Service Results
type GetirProjeListesiDetayResponse struct {
	//Name         xml.Name       `xml:"getirProjeListesiResponse"`
	List  []ProjeListesiDetay `xml:"projeListe"`
	Sonuc Sonuc               `xml:"Sonuc"`
}

type ProjeListesiDetay struct {
	PROJE_ID          string `xml:"PROJE_ID"`
	PROJE_AD          string `xml:"PROJE_AD"`
	PROJE_DURUMU_ID   string `xml:"PROJE_DURUMU_ID"`
	PROJE_DURUMU_AD   string `xml:"PROJE_DURUMU_AD"`
	BAS_TAR           string `xml:"BAS_TAR"`
	BIT_TAR           string `xml:"BIT_TAR"`
	BUTCE             string `xml:"BUTCE"`
	PROJE_KONUMU_ID   string `xml:"PROJE_KONUMU_ID"`
	PROJE_KONUMU_AD   string `xml:"PROJE_KONUMU_AD"`
	PROJE_TURU_ID     string `xml:"PROJE_TURU_ID"`
	PROJE_TURU_AD     string `xml:"PROJE_TURU_AD"`
	PARA_BIRIMI_ID    string `xml:"PARA_BIRIMI_ID"`
	PARA_BIRIMI_AD    string `xml:"PARA_BIRIMI_AD"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
	KAPSAM            string `xml:"KAPSAM"`
	KAPSAM_AD         string `xml:"KAPSAM_AD"`
	UNVAN_ID          string `xml:"UNVAN_ID"`
	UNVAN_AD          string `xml:"UNVAN_AD"`
	KURUM_ID          string `xml:"KURUM_ID"`
	KURUM_AD          string `xml:"KURUM_AD"`
	TESV_PUAN         string `xml:"TESV_PUAN"`
}
