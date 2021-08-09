package Ozgecmis

//My Service Results
type GetirYabanciDilListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirYabanciDilListesiResponse
}
type GetirYabanciDilListesiSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirYabanciDilListesiResponse struct {
	//Name         xml.Name       `xml:"getirProjeListesiResponse"`
	List  []YabanciDilListesi `xml:"yabanciDilListesi"`
	Sonuc Sonuc               `xml:"Sonuc"`
}

type YabanciDilListesi struct {
	Y_ID              string `xml:"Y_ID"`
	DIL_ID            string `xml:"DIL_ID"`
	DIL_AD            string `xml:"DIL_AD"`
	DIL_SINAV_ID      string `xml:"DIL_SINAV_ID"`
	DIL_SINAV_AD      string `xml:"DIL_SINAV_AD"`
	PUAN              string `xml:"PUAN"`
	YIL               string `xml:"YIL"`
	DONEM_ID          string `xml:"DONEM_ID"`
	DONEM_AD          string `xml:"DONEM_AD"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
	SINAVBILGISIYOK   string `xml:"SINAVBILGISIYOK"`
}
