package main

import (
    //"fmt"    
    /*"encoding/json"*/
    "database/sql"    
    _ "github.com/lib/pq"
)





type clinfo struct {
    Juserid int
    Jfmusercode string `json:"fmusercode"`
    Jfmusername string `json:"fmusername"`
    Jspaceid    string `json:"spaceid"`
    Jspacename  string `json:"spacename"`
    Jversion    string `json:"version"`
}




func fmsql(usercode string, spcode string) (infos []clinfo){
    db, err := sql.Open("postgres", "user=fengmu password=zj dbname=yq2 sslmode=disable port=5432")
    checkErr(err)

    
    rows, err := db.Query("select t1.id, t1.username as usercode, t1.first_name as username, t2.spaceid, t2.title as spacename, max(t2.version) as version from   auth_user t1,   fmspace t2,  page t3,  kdsphj t4  where t1.id = t2.userid  and t2.spaceindex = t3.spaceindex      and t3.pageindex = t4.hjnum  and t4.spcode=$1  and t1.username =$2 group by t1.id, t1.username, t1.first_name, t2.spaceid, t2.title", spcode, usercode)
    checkErr(err)

    
    for rows.Next() {
        var info clinfo       
        err = rows.Scan(&info.Juserid, &info.Jfmusercode, &info.Jfmusername, &info.Jspaceid, &info.Jspacename, &info.Jversion)
        
        checkErr(err)
        infos = append(infos, info)       
    }    

    db.Close()
    
    return infos
}



type spacedatadtl struct{    
    Jspcode string `json:"spcode"`
    Jspname string `json:"spname"`
    Jhjcode string `json:"hjcode"`
    Jshelf string `json:"shelf"`
    Jdrawer string `json:"drawer"`
    Jbox string `json:"box"`
    Jmd  string `json:"md"`
    Jamount string `json:"amount"`
    Jinstru string `json:"instru"`
    Jobtype string `json:"obtype"`
    Jcltype string `json:"cltype"`
    Jpx string `json:"px"`
    Jpy string `json:"py"`
}





func query_getSpaceData(Usercode string, SpaceId string, Version string) ( result []spacedatadtl){
    db, err := sql.Open("postgres", "user=fengmu password=zj dbname=yq2 sslmode=disable port=5432")
    checkErr(err)

    //fmt.Println("getSpaceData")

    rows, err := db.Query("select t4.spcode, t4.spname, t4.hjcode, t4.shelf, t4.drawer, t4.box, t4.md, t4.amount, t4.instru, t4.obtype, t4.cltype, t4.px, t4.py  from auth_user t1,      fmspace t2,      page t3,      kdsphj t4 where t1.id = t2.userid        and t2.spaceindex = t3.spaceindex       and t3.pagecode = 'chenlie'       and t3.pageindex = t4.hjnum       and t1.username = $1       and t2.spaceid = $2      and t2.version = $3", Usercode, SpaceId, Version)
    checkErr(err)
    
    
    for rows.Next() {
        var line spacedatadtl       
        err = rows.Scan(&line.Jspcode, &line.Jspname, &line.Jhjcode, &line.Jshelf, &line.Jdrawer, &line.Jbox, &line.Jmd, &line.Jamount, &line.Jinstru, &line.Jobtype, &line.Jcltype, &line.Jpx, &line.Jpy)
        
        //fmt.Println("spcode : %s", line.Jspcode)
       
        checkErr(err)
        
        result = append(result, line)
        

        

    }

    db.Close()
    return
}
