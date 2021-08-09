package Ozgecmis

//My Service Results
type GetTasarimBilgisiV1SrvcResult struct {
	Message string
	Result  bool
	Data    *GetTasarimBilgisiV1Response
}
type GetTasarimBilgisiV1SrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetTasarimBilgisiV1Response struct {
	//Name         xml.Name       `xml:"getirProjeListesiResponse"`
	List  []TasarimListe `xml:"tasarimListe"`
	Sonuc Sonuc          `xml:"Sonuc"`
}

type TasarimListe struct {
	P_TASARIM_ID          string `xml:"P_TASARIM_ID"`
	TASARIM_SAHIPLERI     string `xml:"TASARIM_SAHIPLERI"`
	TASARIM_TURU          string `xml:"TASARIM_TURU"`
	TASARIM_TURU_DETAY    string `xml:"TASARIM_TURU_DETAY"`
	TASARIM_TURU_DETAY_AD string `xml:"TASARIM_TURU_DETAY_AD"`
	KAPSAM                string `xml:"KAPSAM"`
	KAPSAM_AD             string `xml:"KAPSAM_AD"`
	TASARIM_ADI           string `xml:"TASARIM_ADI"`
	TASARIM_OZETI         string `xml:"TASARIM_OZETI"`
	BAS_TARIHI            string `xml:"BAS_TARIHI"`
	BITIS_TARIHI          string `xml:"BITIS_TARIHI"`
	UNVAN_ID              string `xml:"UNVAN_ID"`
	UNVAN_AD              string `xml:"UNVAN_AD"`
	KURUM_ID              string `xml:"KURUM_ID"`
	KURUM_AD              string `xml:"KURUM_AD"`
	TESV_PUAN             string `xml:"TESV_PUAN"`
	KISI_SAYISI           string `xml:"KISI_SAYISI"`
	GUNCELLEME_TARIHI     string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF           string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD        string `xml:"AKTIF_PASIF_AD"`
}
