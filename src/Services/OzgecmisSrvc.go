package Services

import (
	conf "GoYOKAPI/Configurations"
	"GoYOKAPI/Models/Ozgecmis"
	"encoding/xml"
	"github.com/tiaguinho/gosoap"
	"log"
	"net/http"
	"time"
)

var MyOzgecmisService OzgecmisSrvc = OzgecmisSrvc{url: conf.OzgecmisServiceUrl, userId: conf.UserId, password: conf.Password}

type OzgecmisSrvc struct {
	url      string
	userId   string
	password string
}

func (s OzgecmisSrvc) GetProjeListesi(request Ozgecmis.ProjeListesiSrvcRequest) Ozgecmis.ProjeListesiSrvcResult {

	resultArr := Ozgecmis.GetirProjeListesiResponse{}
	serviceResult := Ozgecmis.ProjeListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}

	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.GetirProjeListesiName, paramReq)
	if err != nil {
		log.Fatalf("Call error: %s", err)
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}
func (s OzgecmisSrvc) GetProjeListesiDetay(request Ozgecmis.ProjeListesiDetaySrvcRequest) Ozgecmis.ProjeListesiDetaySrvcResult {

	resultArr := Ozgecmis.GetirProjeListesiDetayResponse{}
	serviceResult := Ozgecmis.ProjeListesiDetaySrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}

	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
			"P_ESER_ID":      request.ProjeId,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.GetirProjeListesiDetayName, paramReq)
	if err != nil {
		log.Fatalf("Call error: %s", err)
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetArastirmaSertifikaBilgisi(request Ozgecmis.ArastirmaSertifkaBilgisSrvcRequest) Ozgecmis.ArastirmaSertifkaBilgisSrvcResult {

	resultArr := Ozgecmis.GetArastirmaSertifkaBilgisiV1Response{}
	serviceResult := Ozgecmis.ArastirmaSertifkaBilgisSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.GetArastirmaSertifikaBilgisiv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetBildiriBilgisi(request Ozgecmis.GetBildiriBilgisiV1SrvcRequest) Ozgecmis.GetBildiriBilgisiV1SrvcResult {

	resultArr := Ozgecmis.GetBildiriBilgisiV1Response{}
	serviceResult := Ozgecmis.GetBildiriBilgisiV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.GetBildiriBilgisiv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetBildiriBilgisiDetay(request Ozgecmis.GetBildiriBilgisiDetayV1SrvcRequest) Ozgecmis.GetBildiriBilgisiDetayV1SrvcResult {

	resultArr := Ozgecmis.GetBildiriBilgisiDetayV1Response{}
	serviceResult := Ozgecmis.GetBildiriBilgisiDetayV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
			"P_ESER_ID":      request.EserId,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.GetBildiriBilgisiDetayv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetEditorlukBilgisi(request Ozgecmis.EditorlukBilgisiV1SrvcRequest) Ozgecmis.EditorlukBilgisiV1SrvcResult {

	resultArr := Ozgecmis.EditorlukBilgisiV1Response{}
	serviceResult := Ozgecmis.EditorlukBilgisiV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.GetEditorlukBilgisiv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetHakemlikBilgisi(request Ozgecmis.HakemlikBilgisiV1SrvcRequest) Ozgecmis.HakemlikBilgisiV1SrvcResult {

	resultArr := Ozgecmis.HakemlikBilgisiV1Response{}
	serviceResult := Ozgecmis.HakemlikBilgisiV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.GetHakemlikBilgisiV1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetAkademikGorevListesi(request Ozgecmis.GetirAkademikGorevListesiSrvcRequest) Ozgecmis.GetirAkademikGorevListesiSrvcResult {

	resultArr := Ozgecmis.GetirAkademikGorevListesiResponse{}
	serviceResult := Ozgecmis.GetirAkademikGorevListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.GetirAkademikGorevListesiName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetAtifSayilari(request Ozgecmis.GetirAtifSayilariSrvcRequest) Ozgecmis.GetirAtifSayilariSrvcResult {

	resultArr := Ozgecmis.GetirAtifSayilariResponse{}
	serviceResult := Ozgecmis.GetirAtifSayilariSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
			"P_DONEM":        request.Yil,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetirAtifSayilariName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetBasvuruDurum(request Ozgecmis.GetirBasvuruDurumSrvcRequest) Ozgecmis.GetirBasvuruDurumSrvcResult {

	resultArr := Ozgecmis.GetirBasvuruDurumResponse{}
	serviceResult := Ozgecmis.GetirBasvuruDurumSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
			"P_DONEM":        request.Yil,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetirBasvuruDurumName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetBeyanTesvik(request Ozgecmis.GetirBeyanTesvikSrvcRequest) Ozgecmis.GetirBeyanTesvikSrvcResult {

	resultArr := Ozgecmis.GetirBeyanTesvikResponse{}
	serviceResult := Ozgecmis.GetirBeyanTesvikSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
			"P_DONEM":        request.Yil,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetirBeyanTesvikName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetDersListesi(request Ozgecmis.GetirDersListesiSrvcRequest) Ozgecmis.GetirDersListesiSrvcResult {

	resultArr := Ozgecmis.GetirDersListesiResponse{}
	serviceResult := Ozgecmis.GetirDersListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetirDersListesiName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetIdariGorevListesi(request Ozgecmis.GetIdariGorevListesiSrvcRequest) Ozgecmis.GetIdariGorevListesiSrvcResult {

	resultArr := Ozgecmis.IdariGorevListesiResponse{}
	serviceResult := Ozgecmis.GetIdariGorevListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetirIdariGorevListesiName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetirOgrenimBilgisi(request Ozgecmis.GetirOgrenimBilgisiListesiSrvcRequest) Ozgecmis.GetirOgrenimBilgisiListesiSrvcResult {

	resultArr := Ozgecmis.GetirOgrenimBilgisiListesiResponse{}
	serviceResult := Ozgecmis.GetirOgrenimBilgisiListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetirOgrenimBilgisiListesiName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetirTezDanismanListesi(request Ozgecmis.GetirTezDanismanListesiSrvcRequest) Ozgecmis.GetirTezDanismanListesiSrvcResult {

	resultArr := Ozgecmis.GetirTezDanismanListesiResponse{}
	serviceResult := Ozgecmis.GetirTezDanismanListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetirTezDanismanListesiName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetirUnvDisiDeneyimListesi(request Ozgecmis.GetirUnvDisiDeneyimListesiSrvcRequest) Ozgecmis.GetirUnvDisiDeneyimListesiSrvcResult {

	resultArr := Ozgecmis.GetirUnvDisiDeneyimListesiResponse{}
	serviceResult := Ozgecmis.GetirUnvDisiDeneyimListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetirUnvDisiDeneyimListesiName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetirUyelikListesi(request Ozgecmis.GetirUyelikListesiResponseSrvcRequest) Ozgecmis.GetirUyelikListesiResponseSrvcResult {

	resultArr := Ozgecmis.GetirUyelikListesiResponse{}
	serviceResult := Ozgecmis.GetirUyelikListesiResponseSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetirUyelikListesiName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetirYabanciDilListesi(request Ozgecmis.GetirYabanciDilListesiSrvcRequest) Ozgecmis.GetirYabanciDilListesiSrvcResult {

	resultArr := Ozgecmis.GetirYabanciDilListesiResponse{}
	serviceResult := Ozgecmis.GetirYabanciDilListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetirYabanciDilListesi, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetKitapBilgisiv1(request Ozgecmis.GetKitapBilgisiV1SrvcRequest) Ozgecmis.GetKitapBilgisiV1SrvcResult {

	resultArr := Ozgecmis.GetKitapBilgisiV1Response{}
	serviceResult := Ozgecmis.GetKitapBilgisiV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetKitapBilgisiv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetKitapBilgisiDetay(request Ozgecmis.GetKitapBilgisiDetayV1SrvcRequest) Ozgecmis.GetKitapBilgisiDetayV1SrvcResult {

	resultArr := Ozgecmis.GetKitapBilgisiDetayV1Response{}
	serviceResult := Ozgecmis.GetKitapBilgisiDetayV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
			"P_ESER_ID":      request.KitapId,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetKitapBilgisiDetayName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetMakaleBilgisiDetay(request Ozgecmis.GetMakaleBilgisiDetayV1SrvcRequest) Ozgecmis.GetMakaleBilgisiDetayV1SrvcResult {

	resultArr := Ozgecmis.GetMakaleBilgisiDetayV1Response{}
	serviceResult := Ozgecmis.GetMakaleBilgisiDetayV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
			"P_ESER_ID":      request.MakaleId,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetMakaleBilgisiDetayv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetMakaleBilgisiv1(request Ozgecmis.GetMakaleBilgisiV1SrvcRequest) Ozgecmis.GetMakaleBilgisiV1SrvcResult {

	resultArr := Ozgecmis.GetMakaleBilgisiV1Response{}
	serviceResult := Ozgecmis.GetMakaleBilgisiV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetMakaleBilgisiv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetOdulListesi(request Ozgecmis.GetOdulListesiV1SrvcRequest) Ozgecmis.GetOdulListesiV1SrvcResult {

	resultArr := Ozgecmis.GetOdulListesiV1Response{}
	serviceResult := Ozgecmis.GetOdulListesiV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetOdulListesiv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetPatentBilgisiv1(request Ozgecmis.GetirPatentBilgisiV1ListesiSrvcRequest) Ozgecmis.GetirPatentBilgisiV1ListesiSrvcResult {

	resultArr := Ozgecmis.GetirPatentBilgisiListesiV1Response{}
	serviceResult := Ozgecmis.GetirPatentBilgisiV1ListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetPatentBilgisiv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetPatentBilgisiDetay(request Ozgecmis.GetirPatentBilgisiDetayV1ListesiSrvcRequest) Ozgecmis.GetirPatentBilgisiDetayV1ListesiSrvcResult {

	resultArr := Ozgecmis.GetirPatentBilgisiDetayListesiV1Response{}
	serviceResult := Ozgecmis.GetirPatentBilgisiDetayV1ListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
			"P_ESER_ID":      request.PatentId,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetPatentBilgisiDetayv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetPersonelLinkv1(request Ozgecmis.GetirPersonelLinkV1ListesiSrvcRequest) Ozgecmis.GetirPersonelLinkV1ListesiSrvcResult {

	resultArr := Ozgecmis.GetirPersonelLinkListesiV1Response{}
	serviceResult := Ozgecmis.GetirPersonelLinkV1ListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetPersonelLinkv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetSanatsalFaalv1(request Ozgecmis.GetirSanatsalFaalv1SrvcRequest) Ozgecmis.GetirSanatsalFaalv1SrvcResult {

	resultArr := Ozgecmis.GetirSanatsalFaalv1Response{}
	serviceResult := Ozgecmis.GetirSanatsalFaalv1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetSanatsalFaalv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetTasarimBilgisiv1(request Ozgecmis.GetTasarimBilgisiV1SrvcRequest) Ozgecmis.GetTasarimBilgisiV1SrvcResult {

	resultArr := Ozgecmis.GetTasarimBilgisiV1Response{}
	serviceResult := Ozgecmis.GetTasarimBilgisiV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetTasarimBilgisiv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetTemelAlanBilgisiv1(request Ozgecmis.GetTemelAlanBilgisiV1SrvcRequest) Ozgecmis.GetTemelAlanBilgisiV1SrvcResult {

	resultArr := Ozgecmis.GetTemelAlanBilgisiV1Response{}
	serviceResult := Ozgecmis.GetTemelAlanBilgisiV1SrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.GetTemelAlanBilgisiv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetTesvikKatSayi(request Ozgecmis.GetTesvikKatsayiSrvcRequest) Ozgecmis.GetTesvikKatsayiSrvcResult {

	resultArr := Ozgecmis.GetTesvikKatsayiResponse{}
	serviceResult := Ozgecmis.GetTesvikKatsayiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID":  s.userId,
			"P_SIFRE":         s.password,
			"TESVIK_KURAL_ID": request.TesvikKuralId,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetTesvikKatsayisiName, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}

func (s OzgecmisSrvc) GetYazarListesi(request Ozgecmis.GetirYazarListesiSrvcRequest) Ozgecmis.GetirYazarListesiSrvcResult {

	resultArr := Ozgecmis.GetirYazarListesiResponse{}
	serviceResult := Ozgecmis.GetirYazarListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"P_KULLANICI_ID": s.userId,
			"P_SIFRE":        s.password,
			"P_TC_KIMLIK_NO": request.TcKimlikNo,
			"P_YAZAR_ID":     request.YazarId,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password
	soapResult, err := soap.Call(conf.GetYazarListesiv1Name, paramReq)
	if err != nil {
		serviceResult.Message = err.Error()
		return serviceResult
	}
	err = xml.Unmarshal(soapResult.Body, &resultArr)
	if err != nil {
		return serviceResult
	}

	serviceResult.Message = conf.ServiceResultSuccessName
	serviceResult.Result = true
	return serviceResult
}
