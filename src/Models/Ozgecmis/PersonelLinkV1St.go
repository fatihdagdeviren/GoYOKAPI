package Ozgecmis

//My Service Results
type GetirPersonelLinkV1ListesiSrvcResult struct {
	Message string
	Result  bool
	Data    *GetirPersonelLinkListesiV1Response
}
type GetirPersonelLinkV1ListesiSrvcRequest struct {
	TcKimlikNo string
}

//SOAP Service Results
type GetirPersonelLinkListesiV1Response struct {
	List  []PersonelLinkListesi `xml:"personelLinkListe"`
	Sonuc Sonuc                 `xml:"Sonuc"`
}

type PersonelLinkListesi struct {
	ARASTIRMACI_ID   string `xml:"ARASTIRMACI_ID"`
	KADRO_UNVAN_ADI  string `xml:"KADRO_UNVAN_ADI"`
	PERSONEL_ADI     string `xml:"PERSONEL_ADI"`
	PERSONEL_SOYADI  string `xml:"PERSONEL_SOYADI"`
	KADRO_YERI       string `xml:"KADRO_YERI"`
	YOKAKADEMIK_LINK string `xml:"YOKAKADEMIK_LINK"`
	ORCID            string `xml:"ORCID"`
	RESEARCHER_ID    string `xml:"RESEARCHER_ID"`
}
