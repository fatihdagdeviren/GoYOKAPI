package Ozgecmis

//My Service Results
type GetirPatentBilgisiDetayV1ListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirPatentBilgisiDetayListesiV1Response
}
type GetirPatentBilgisiDetayV1ListesiSrvcRequest struct {
	TcKimlikNo string
	PatentId   string
}

//SOAP Service Results
type GetirPatentBilgisiDetayListesiV1Response struct {
	List  []PatentBilgisiDetayListesi `xml:"patentListe"`
	Sonuc Sonuc                       `xml:"Sonuc"`
}

type PatentBilgisiDetayListesi struct {
	PATENT_ID         string `xml:"PATENT_ID"`
	PATENT_NO         string `xml:"PATENT_NO"`
	PATENT_TARIHI     string `xml:"PATENT_TARIHI"`
	BASVURU_SAHIPLERI string `xml:"BASVURU_SAHIPLERI"`
	BULUS_SAHIPLERI   string `xml:"BULUS_SAHIPLERI"`
	KATEGORI          string `xml:"KATEGORI"`
	KATEGORI_ID       string `xml:"KATEGORI_ID"`
	DOSYA_TIPI        string `xml:"DOSYA_TIPI"`
	DOSYA_TIPI_ID     string `xml:"DOSYA_TIPI_ID"`
	KAPSAM            string `xml:"KAPSAM"`
	KAPSAM_ID         string `xml:"KAPSAM_ID"`
	PATENT_SINIF      string `xml:"PATENT_SINIF"`
	PATENT_ADI        string `xml:"PATENT_ADI"`
	GUNCELLEME_TARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIF_PASIF       string `xml:"AKTIF_PASIF"`
	AKTIF_PASIF_AD    string `xml:"AKTIF_PASIF_AD"`
	UNVAN_ID          string `xml:"UNVAN_ID"`
	UNVAN_AD          string `xml:"UNVAN_AD"`
	KURUM_ID          string `xml:"KURUM_ID"`
	KURUM_AD          string `xml:"KURUM_AD"`
	KISI_SAYISI       string `xml:"KISI_SAYISI"`
	TESV_PUAN         string `xml:"TESV_PUAN"`
}
