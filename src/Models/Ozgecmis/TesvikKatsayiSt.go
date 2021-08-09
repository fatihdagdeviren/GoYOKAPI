package Ozgecmis

//My Service Results
type GetTesvikKatsayiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetTesvikKatsayiResponse
}
type GetTesvikKatsayiSrvcRequest struct {
	TesvikKuralId string
}

//SOAP Service Results
type GetTesvikKatsayiResponse struct {
	//Name         xml.Name       `xml:"getirProjeListesiResponse"`
	List  []GetTesvikKatsayi `xml:"tesvikKatsayi"`
	Sonuc Sonuc              `xml:"Sonuc"`
}

type GetTesvikKatsayi struct {
	K_ID              string `xml:"K_ID"`
	ESER_TURU         string `xml:"ESER_TURU"`
	KATKI_DUZEYI      string `xml:"KATKI_DUZEYI"`
	ACIKLAMA          string `xml:"ACIKLAMA"`
	KATSAYI           string `xml:"KATSAYI"`
	ALAN_TUR          string `xml:"ALAN_TUR"`
	FAALIYETGENELPUAN string `xml:"FAALIYETGENELPUAN"`
}
