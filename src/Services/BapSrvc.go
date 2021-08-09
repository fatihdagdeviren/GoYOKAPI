package Services

import (
	conf "GoYOKAPI/Configurations"
	"GoYOKAPI/Models/Bap"
	"encoding/xml"
	"github.com/tiaguinho/gosoap"
	"net/http"
	"time"
)

var MyBapService BapSrvc = BapSrvc{url: conf.BapServiceUrl, userId: conf.UserId, password: conf.Password}

type BapSrvc struct {
	url      string
	userId   string
	password string
}

func (s BapSrvc) GetBapProjeBilgisi(request Bap.GetBapProjeBilgiSrvcRequest) Bap.GetBapProjeBilgiSrvcResult {

	resultArr := Bap.GetBapProjeBilgiResponse{}
	serviceResult := Bap.GetBapProjeBilgiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		//log.Fatalf("SoapClient error: %s", err)
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"kontrol": gosoap.Params{
				"P_KULLANICI_ID": s.userId,
				"P_SIFRE":        s.password,
				"P_TC_KIMLIK_NO": request.TcKimlikNo,
			},
			"P_UNV_PROJE_NO": request.ProjeNo,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.BapProjeBilgiName, paramReq)
	if err != nil {
		//log.Fatalf("Call error: %s", err)
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

func (s BapSrvc) GetBapProjeEkipBilgisi(request Bap.GetBapProjeEkipListesiSrvcRequest) Bap.GetBapProjeEkipListesiSrvcResult {

	resultArr := Bap.GetBapProjeEkipListesiResponse{}
	serviceResult := Bap.GetBapProjeEkipListesiSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		//log.Fatalf("SoapClient error: %s", err)
		serviceResult.Message = err.Error()
		return serviceResult
	}
	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"kontrol": gosoap.Params{
				"P_KULLANICI_ID": s.userId,
				"P_SIFRE":        s.password,
				"P_TC_KIMLIK_NO": request.TcKimlikNo,
			},
			"P_YOK_PROJE_ID": request.YOKPROJEID,
		},
	}

	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.BapProjeEkipBilgisiName, paramReq)
	if err != nil {
		//log.Fatalf("Call error: %s", err)
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

func (s BapSrvc) InsertOrUpdateBapProje(request Bap.InsertOrUpdateBapProjeSrvcRequest) Bap.InsertOrUpdateBapProjeSrvcResult {

	resultArr := Bap.InsertOrUpdateBapProjeResponse{}
	serviceResult := Bap.InsertOrUpdateBapProjeSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		//log.Fatalf("SoapClient error: %s", err)
		serviceResult.Message = err.Error()
		return serviceResult
	}

	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"kontrol": gosoap.Params{
				"P_KULLANICI_ID": s.userId,
				"P_SIFRE":        s.password,
				"P_TC_KIMLIK_NO": request.TcKimlikNo,
			},
			"P_YOK_PROJE_ID":  request.YokProjeId,
			"P_UNV_PROJE_NO":  request.ProjeNo,
			"P_PROJE_ADI":     request.ProjeAdi,
			"P_PROJE_ALANI":   request.ProjeAlani,
			"P_PROJE_BAS_TAR": request.ProjeBasTar,
			"P_PROJE_BIT_TAR": request.ProjeBitTar,
			"P_PARA_BIRIMI":   request.ParaBirimi,
			"P_BUTCE":         request.Butce,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.BapInsertOrUpdateRequestName, paramReq)
	if err != nil {
		//log.Fatalf("Call error: %s", err)
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

func (s BapSrvc) DeleteBapProjeOrEkip(request Bap.DeleteProjeSrvcRequest) Bap.DeleteBapProjeSrvcResult {

	resultArr := Bap.InsertOrUpdateBapProjeResponse{}
	serviceResult := Bap.DeleteBapProjeSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		//log.Fatalf("SoapClient error: %s", err)
		serviceResult.Message = err.Error()
		return serviceResult
	}

	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"kontrol": gosoap.Params{
				"P_KULLANICI_ID": s.userId,
				"P_SIFRE":        s.password,
				"P_TC_KIMLIK_NO": request.TcKimlikNo,
			},
			"P_SILINME_YERI":  request.SilinmeYeri,
			"P_UNV_PROJE_NO":  request.ProjeNo,
			"P_PROJE_EKIP_ID": request.ProjeEkipId,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.BapDeleteBapProjeOrEkipName, paramReq)
	if err != nil {
		//log.Fatalf("Call error: %s", err)
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

func (s BapSrvc) InsertOrUpdateBapProjeEkip(request Bap.InsertOrUpdateBapProjeEkipSrvcRequest) Bap.InsertOrUpdateBapProjeEkipSrvcResult {

	resultArr := Bap.InsertOrUpdateBapProjeEkipResponse{}
	serviceResult := Bap.InsertOrUpdateBapProjeEkipSrvcResult{
		Result: false,
		Data:   &resultArr,
	}

	httpClient := &http.Client{
		Timeout: 150000 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(s.url, httpClient)
	if err != nil {
		//log.Fatalf("SoapClient error: %s", err)
		serviceResult.Message = err.Error()
		return serviceResult
	}

	paramReq := gosoap.Params{
		"parametre": gosoap.Params{
			"kontrol": gosoap.Params{
				"P_KULLANICI_ID": s.userId,
				"P_SIFRE":        s.password,
				"P_TC_KIMLIK_NO": request.TcKimlikNo,
			},
			"P_PROJE_EKIP_ID":    request.ProjeEkipId,
			"P_YOKPROJEID":       request.YokProjeId,
			"P_ARASTIRMACI_TIP":  request.ArastirmaciTip,
			"P_ARASTIRMACI_TC":   request.ArastirmaciTc,
			"P_PROJEDEKI_GOREVI": request.ProjedekiGorevi,
			"P_PERSONEL_AD":      request.PersonelAd,
			"P_PERSONEL_SOYAD":   request.PersonelSoyad,
		},
	}
	soap.Username = s.userId
	soap.Password = s.password

	soapResult, err := soap.Call(conf.BapInsertOrUpdateProjeEkipName, paramReq)
	if err != nil {
		//log.Fatalf("Call error: %s", err)
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
