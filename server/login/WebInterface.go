package login

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostToWebServer(addr string, w http.ResponseWriter, r *http.Request) {
	reqData, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqData))
	client := &http.Client{}
	//data := make(map[string]interface{})
	//data["uuid"] = "zhaofan"
	//data["age"] = "23"
	//bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", CONF.PvpWeb.Endpoints[0]+"/login", bytes.NewReader(reqData) /*bytes.NewReader(bytesData)*/)
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func Login(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("login", w, r)
}
func createPlayer(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("createPlayer", w, r)
}
func addHero(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("addHero", w, r)
}
func bindHeroEquip(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("bindHeroEquip", w, r)
}
func tdownHeroEquip(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("tdownHeroEquip", w, r)
}
func addItem(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("addItem", w, r)
}
func addEquip(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("addEquip", w, r)
}
func getGoodsByReduce(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("getGoodsByReduce", w, r)
}
func openBox(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("openBox", w, r)
}
func getToken(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("getToken", w, r)
}
func getNonce(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("getNonce", w, r)
}
func singVerify(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("singVerify", w, r)
}
func getLeaderboard(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("getLeaderboard", w, r)
}
func refreshUserLeaderboard(w http.ResponseWriter, r *http.Request) {
	PostToWebServer("refreshUserLeaderboard", w, r)
}
