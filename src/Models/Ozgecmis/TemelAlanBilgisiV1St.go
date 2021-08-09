package Ozgecmis

//My Service Results
type GetTemelAlanBilgisiV1SrvcResult struct {
	Message string
	Result  bool
	Data    *GetTemelAlanBilgisiV1Response
}

type GetTemelAlanBilgisiV1SrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetTemelAlanBilgisiV1Response struct {
	//Name         xml.Name       `xml:"getirProjeListesiResponse"`
	List  []TemelAlanListesi `xml:"temelAlanListe"`
	Sonuc Sonuc              `xml:"Sonuc"`
}

type TemelAlanListesi struct {
	T_UAK_ID          string `xml:"T_UAK_ID"`
	TEMEL_ALAN_ID     string `xml:"TEMEL_ALAN_ID"`
	BILIM_ALAN_ID     string `xml:"BILIM_ALAN_ID"`
	ANAHTARKELIME1_ID string `xml:"ANAHTARKELIME1_ID"`
	ANAHTARKELIME2_ID string `xml:"ANAHTARKELIME2_ID"`
	ANAHTARKELIME3_ID string `xml:"ANAHTARKELIME3_ID"`
	TEMEL_ALAN_AD     string `xml:"TEMEL_ALAN_AD"`
	BILIM_ALAN_AD     string `xml:"BILIM_ALAN_AD"`
	ANAHTARKELIME1_AD string `xml:"ANAHTARKELIME1_AD"`
	ANAHTARKELIME2_AD string `xml:"ANAHTARKELIME2_AD"`
	ANAHTARKELIME3_AD string `xml:"ANAHTARKELIME3_AD"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
}
