package Ozgecmis

//My Service Results
type ArastirmaSertifkaBilgisSrvcResult struct {
	Message string
	Result  bool
	Data    *GetArastirmaSertifkaBilgisiV1Response
}

type ArastirmaSertifkaBilgisSrvcRequest struct {
	TcKimlikNo string
}

type GetArastirmaSertifkaBilgisiV1Response struct {
	List  []ArastirmaListesi `xml:"arastirmaListe"`
	Sonuc Sonuc              `xml:"Sonuc"`
}

type ArastirmaListesi struct {
	SID              string `xml:"S_ID"`
	TURID            string `xml:"TUR_ID"`
	TURADI           string `xml:"TUR_ADI"`
	ADI              string `xml:"ADI"`
	ICERIK           string `xml:"ICERIK"`
	YER              string `xml:"YER"`
	KAPSAM           string `xml:"KAPSAM"`
	KAPSAMAD         string `xml:"KAPSAM_AD"`
	BASTAR           string `xml:"BASTAR"`
	BITTAR           string `xml:"BITTAR"`
	UNVANID          string `xml:"UNVAN_ID"`
	UNVANAD          string `xml:"UNVAN_AD"`
	KURUMAD          string `xml:"KURUM_AD"`
	KURUMID          string `xml:"KURUM_ID"`
	KISISAYISI       string `xml:"KISI_SAYISI"`
	ULKESEHIR        string `xml:"ULKE_SEHIR"`
	GUNCELLEMETARIHI string `xml:"GUNCELLEME_TARIHI"`
	AKTIFPASIF       string `xml:"AKTIF_PASIF"`
	AKTIFPASIFAD     string `xml:"AKTIF_PASIF_AD"`
}
