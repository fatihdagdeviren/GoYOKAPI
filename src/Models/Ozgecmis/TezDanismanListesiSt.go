package Ozgecmis

//My Service Results
type GetirTezDanismanListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirTezDanismanListesiResponse
}
type GetirTezDanismanListesiSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirTezDanismanListesiResponse struct {
	//Name         xml.Name       `xml:"getirProjeListesiResponse"`
	List  []TezDanismanListesi `xml:"tezDanismanListesi"`
	Sonuc Sonuc                `xml:"Sonuc"`
}

type TezDanismanListesi struct {
	VERIKAYNAK        string `xml:"VERIKAYNAK"`
	YER_ID            string `xml:"YER_ID"`
	YER_ADI           string `xml:"YER_ADI"`
	ULKE              string `xml:"ULKE"`
	KAYIT_ID          string `xml:"KAYIT_ID"`
	YIL               string `xml:"YIL"`
	TUR_ID            string `xml:"TUR_ID"`
	TUR_ADI           string `xml:"TUR_ADI"`
	TEZ_ADI           string `xml:"TEZ_ADI"`
	DURUM_ADI         string `xml:"DURUM_ADI"`
	YAZAR_ADI         string `xml:"YAZAR_ADI"`
	YAZAR_SOYADI      string `xml:"YAZAR_SOYADI"`
	UNIVERSITE_AD     string `xml:"UNIVERSITE_AD"`
	ENSTITU_AD        string `xml:"ENSTITU_AD"`
	ABD_AD            string `xml:"ABD_AD"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
	DANISMAN_SIRASI   string `xml:"DANISMAN_SIRASI"`
}
