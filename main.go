package main

import (
    "fmt"
    "github.com/drone/routes"
    "net/http" 
    "encoding/json"
    /*"database/sql"    
    _ "github.com/lib/pq"*/
)







func getclinfo(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    userid := params.Get(":userid")
    spcode := params.Get(":spcode")


    
    fmt.Println("********************"+userid)
    infos := fmsql(userid, spcode)
    fmt.Println("********************"+spcode)

    //fmt.Println("usercode:", infos[0].Jfmusercode)

    b, err := json.Marshal(infos)

    if err != nil {
        fmt.Println("json err:", err)
    }

    fmt.Fprintf(w, "%s", b)
}

func getSpaceData(w http.ResponseWriter, r *http.Request){
    params := r.URL.Query()
    username:= params.Get(":username")
    spaceid := params.Get(":spaceid")
    version := params.Get(":version")
    
    fmt.Println("*********************"+spaceid)
    spacedata := query_getSpaceData(username, spaceid, version)
    fmt.Println("*********************"+version)
    //fmt.Println("spcode", spacedata[0].Jspcode)

    b, err := json.Marshal(spacedata)

    if err != nil {
        fmt.Println("json err:", err)
    }

    fmt.Fprintf(w, "%s", b)

}

func wx(w http.ResponseWriter, r *http.Request){    
    r.ParseForm()
    echostr := r.Form["echostr"][0]
    fmt.Println(r.Form)
    fmt.Println(r.URL.Path)
    fmt.Println(echostr)
    fmt.Fprintf(w, "%s", echostr)
}





func checkErr(err error) {    
    if err != nil {
        panic(err)
    }
}

func main() {
    //query_getSpaceData("fengmu","2014071211424522","20140712164102")
    mux := routes.New()
    //mux.Get("/getclinfo/:userid/:spcode", getclinfo)
    //mux.Get("/getSpaceData/:username/:spaceid/:version", getSpaceData)
    mux.Get("/wx", wx)
    http.Handle("/", mux)
    http.ListenAndServe(":20900", nil)
}