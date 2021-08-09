package main

import (
	bapModels "GoYOKAPI/Models/Bap"
	"GoYOKAPI/Models/Ozgecmis"
	myService "GoYOKAPI/Services"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

const (
	user = "fd"
	pass = "fd"
)

///////// ****************** OZGECMIS Services ************************** /////////

func GetProjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//http://localhost:8000/getProjeListesi/tc
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.ProjeListesiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetProjeListesi(request))
}

func GetProjeListesiDetay(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//http://localhost:8000/getProjeListesi/tc
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.ProjeListesiDetaySrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetProjeListesiDetay(request))
}

func GetArastirmaSertifikaBilgisi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo" : "Tc"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.ArastirmaSertifkaBilgisSrvcRequest{}
	json.Unmarshal(body, &request)

	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetArastirmaSertifikaBilgisi(request))
}

func GetBildiriBilgisiv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo" : "Tc"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetBildiriBilgisiV1SrvcRequest{}
	json.Unmarshal(body, &request)

	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetBildiriBilgisi(request))
}

func GetBildiriBilgisiDetayv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo" : "Tc"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetBildiriBilgisiDetayV1SrvcRequest{}
	json.Unmarshal(body, &request)

	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetBildiriBilgisiDetay(request))
}

func GetEditorlukBilgisiv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo" : "Tc"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.EditorlukBilgisiV1SrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetEditorlukBilgisi(request))
}

func GetHakemlikBilgisiv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.HakemlikBilgisiV1SrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetHakemlikBilgisi(request))
}

func GetirAkademikGorevListesiv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirAkademikGorevListesiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetAkademikGorevListesi(request))
}

func GetirAtifSayilari(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//	"Yil":"2019"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	atifSayilariRequest := Ozgecmis.GetirAtifSayilariSrvcRequest{}
	json.Unmarshal(body, &atifSayilariRequest)

	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetAtifSayilari(atifSayilariRequest))
}

func GetirBasvuruDurum(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//	"Yil":"2019"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirBasvuruDurumSrvcRequest{}
	json.Unmarshal(body, &request)

	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetBasvuruDurum(request))
}

func GetirBeyanTesvik(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//	"Yil":"2019"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirBeyanTesvikSrvcRequest{}
	json.Unmarshal(body, &request)

	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetBeyanTesvik(request))
}

func GetDersListesi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirDersListesiSrvcRequest{}
	json.Unmarshal(body, &request)

	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetDersListesi(request))
}

func GetIdariGorevListesi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetIdariGorevListesiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetIdariGorevListesi(request))
}

func GetOgrenimBilgisi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirOgrenimBilgisiListesiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetirOgrenimBilgisi(request))
}

func GetTezDanismanListesi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirTezDanismanListesiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetirTezDanismanListesi(request))
}

func GetUnvDisiDeneyimListesi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirUnvDisiDeneyimListesiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetirUnvDisiDeneyimListesi(request))
}

func GetUyelikListesi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirUyelikListesiResponseSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetirUyelikListesi(request))
}

func GetYabanciDilListesi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirYabanciDilListesiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetirYabanciDilListesi(request))
}

func GetKitapListesi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetKitapBilgisiV1SrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetKitapBilgisiv1(request))
}
func GetKitapListesiDetay(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetKitapBilgisiDetayV1SrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetKitapBilgisiDetay(request))
}

func GetMakaleListesi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetMakaleBilgisiV1SrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetMakaleBilgisiv1(request))
}

func GetMakaleListesiDetay(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetMakaleBilgisiDetayV1SrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetMakaleBilgisiDetay(request))
}

func GetOdulListesi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetOdulListesiV1SrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetOdulListesi(request))
}

func GetPatentBilgisiv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirPatentBilgisiV1ListesiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetPatentBilgisiv1(request))
}

func GetPatentBilgisiDetay(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirPatentBilgisiDetayV1ListesiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetPatentBilgisiDetay(request))
}

func GetPersonelLinkv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirPersonelLinkV1ListesiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetPersonelLinkv1(request))
}

func GetSanatsalFaalv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirSanatsalFaalv1SrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetSanatsalFaalv1(request))
}

func GetTasarimBilgisiv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetTasarimBilgisiV1SrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetTasarimBilgisiv1(request))
}

func GetTemelAlanBilgisiv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetTemelAlanBilgisiV1SrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetTemelAlanBilgisiv1(request))
}

func GetTesvikKatsayi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TesvikKuralId":"TC",
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetTesvikKatsayiSrvcRequest{}
	json.Unmarshal(body, &request)
	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetTesvikKatSayi(request))
}

func GetYazarListesiv1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"TC",
	//	"YazarId":"YazarId" -> Kitap, Makaleden alınacak.
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	request := Ozgecmis.GetirYazarListesiSrvcRequest{}
	json.Unmarshal(body, &request)

	json.NewEncoder(w).Encode(myService.MyOzgecmisService.GetYazarListesi(request))
}

// ********************************************************************** ///////////

///////// ****************** BAP Services ************************** /////////

func GetBapProjeBilgisi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TCKimlikNo":"TC",
	//	"ProjeNo":"14-TIP-069"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	bapProjeBilgiRequest := bapModels.GetBapProjeBilgiSrvcRequest{}
	json.Unmarshal(body, &bapProjeBilgiRequest)
	json.NewEncoder(w).Encode(myService.MyBapService.GetBapProjeBilgisi(bapProjeBilgiRequest))
}

func GetBapProjeEkipListesi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TCKimlikNo":"TC",
	//	"YOKPROJEID":"34137"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	bapProjeEkipBilgiRequest := bapModels.GetBapProjeEkipListesiSrvcRequest{}
	json.Unmarshal(body, &bapProjeEkipBilgiRequest)
	json.NewEncoder(w).Encode(myService.MyBapService.GetBapProjeEkipBilgisi(bapProjeEkipBilgiRequest))
}

func InsertOrUpdateBapProjeBilgisi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo": "TC",
	//	"ProjeNo" : "12345-FTH",
	//	"ProjeAdi" : "Test Proje Adı",
	//	"ProjeAlani" : "2",
	//	"ProjeBasTar" :"04/12/2020",
	//	"ProjeBitTar" : "04/12/2021",
	//	"ParaBirimi" : "2",
	//	"Butce"	: "1560"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	bapProjeBilgiRequest := bapModels.InsertOrUpdateBapProjeSrvcRequest{}
	json.Unmarshal(body, &bapProjeBilgiRequest)
	json.NewEncoder(w).Encode(myService.MyBapService.InsertOrUpdateBapProje(bapProjeBilgiRequest))
}

func DeleteBapProjeOrEkip(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"",
	//	"SilinmeYeri":"1",
	//	"ProjeNo":"12345-FTH",
	//	"ProjeEkipId":""
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	bapProjeBilgiRequest := bapModels.DeleteProjeSrvcRequest{}
	json.Unmarshal(body, &bapProjeBilgiRequest)
	json.NewEncoder(w).Encode(myService.MyBapService.DeleteBapProjeOrEkip(bapProjeBilgiRequest))
}

func InsertOrUpdateBapProjeEkipBilgisi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//{
	//	"TcKimlikNo":"",
	//	"ProjeEkipId":"",
	//	"YokProjeId":"60910",
	//	"ArastirmaciTip":"2",
	//	"ArastirmaciTc":"",
	//	"ProjedekiGorevi":"12",
	//	"PersonelAd":"Fatih",
	//	"PersonelSoyad":"DAĞDEVİREN"
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	bapProjeBilgiRequest := bapModels.InsertOrUpdateBapProjeEkipSrvcRequest{}
	json.Unmarshal(body, &bapProjeBilgiRequest)
	json.NewEncoder(w).Encode(myService.MyBapService.InsertOrUpdateBapProjeEkip(bapProjeBilgiRequest))
}

// ********************************************************************** ///////////

func BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func main() {
	r := httprouter.New()
	r.POST("/getArastirmaSertifikaBilgisi", BasicAuth(GetArastirmaSertifikaBilgisi, user, pass))
	r.POST("/getBildiriBilgisiv1", BasicAuth(GetBildiriBilgisiv1, user, pass))
	r.POST("/getBildiriBilgisiDetayv1", BasicAuth(GetBildiriBilgisiDetayv1, user, pass))
	r.POST("/getEditorlukBilgisiv1", BasicAuth(GetEditorlukBilgisiv1, user, pass))
	r.POST("/getHakemlikBilgisiv1", BasicAuth(GetHakemlikBilgisiv1, user, pass))
	r.POST("/getAkademikGorevListesiv1", BasicAuth(GetirAkademikGorevListesiv1, user, pass))
	r.POST("/getAtifSayilari", BasicAuth(GetirAtifSayilari, user, pass))
	r.POST("/getBasvuruDurum", BasicAuth(GetirBasvuruDurum, user, pass))
	r.POST("/getBeyanTesvik", BasicAuth(GetirBeyanTesvik, user, pass))
	r.POST("/getDersListesi", BasicAuth(GetDersListesi, user, pass))
	r.POST("/getIdariGorevListesi", BasicAuth(GetIdariGorevListesi, user, pass))
	r.POST("/getOgrenimBilgisi", BasicAuth(GetOgrenimBilgisi, user, pass))
	r.POST("/getProjeListesi", BasicAuth(GetProjects, user, pass))
	r.POST("/getProjeListesiDetay", BasicAuth(GetProjeListesiDetay, user, pass))
	r.POST("/getTezDanismanListesi", BasicAuth(GetTezDanismanListesi, user, pass))
	r.POST("/getUnvDisiDeneyimListesi", BasicAuth(GetUnvDisiDeneyimListesi, user, pass))
	r.POST("/getUyelikListesi", BasicAuth(GetUyelikListesi, user, pass))
	r.POST("/getYabanciDilListesi", BasicAuth(GetYabanciDilListesi, user, pass))
	r.POST("/getKitapListesi", BasicAuth(GetKitapListesi, user, pass))
	r.POST("/getKitapListesiDetay", BasicAuth(GetKitapListesiDetay, user, pass))
	r.POST("/getMakaleListesi", BasicAuth(GetMakaleListesi, user, pass))
	r.POST("/getMakaleListesiDetay", BasicAuth(GetMakaleListesiDetay, user, pass))
	r.POST("/getOdulListesi", BasicAuth(GetOdulListesi, user, pass))
	r.POST("/getPatentBilgisiv1", BasicAuth(GetPatentBilgisiv1, user, pass))
	r.POST("/getPatentBilgisiDetay", BasicAuth(GetPatentBilgisiDetay, user, pass))
	r.POST("/getPersonelLinkv1", BasicAuth(GetPersonelLinkv1, user, pass))
	r.POST("/getSanatsalFaalv1", BasicAuth(GetSanatsalFaalv1, user, pass))
	r.POST("/getTasarimBilgisiv1", BasicAuth(GetTasarimBilgisiv1, user, pass))
	r.POST("/getTemelAlanBilgisiv1", BasicAuth(GetTemelAlanBilgisiv1, user, pass))
	r.POST("/getTesvikKatsayi", BasicAuth(GetTesvikKatsayi, user, pass))
	r.POST("/getYazarListesiv1", BasicAuth(GetYazarListesiv1, user, pass))

	//BAP
	r.POST("/getBapProjeBilgisi", BasicAuth(GetBapProjeBilgisi, user, pass))
	r.POST("/getBapProjeEkipListesi", BasicAuth(GetBapProjeEkipListesi, user, pass))
	r.POST("/insertOrUpdateBapProjeBilgisi", BasicAuth(InsertOrUpdateBapProjeBilgisi, user, pass))
	r.POST("/deleteBapProjeOrEkip", BasicAuth(DeleteBapProjeOrEkip, user, pass))
	r.POST("/insertOrUpdateBapProjeEkipBilgisi", BasicAuth(InsertOrUpdateBapProjeEkipBilgisi, user, pass))

	http.ListenAndServe(":8000", r)
}
