package main

import (
    "fmt"
    "github.com/fffilimonov/OCIP_go"
    "os"
)

func main() {
    larg:=len(os.Args)
    if larg < 2 {
        ocip.LogErr(nil,"no args")
        os.Exit (1)
    }
    var file string = os.Args[1]
    Config := ocip.ReadConfig(file)
    var USERID,MODE,DST string
    if larg > 2 {
        MODE = os.Args[2]
    }
    if larg > 3 {
        USERID = ocip.ConcatStr("",os.Args[3],"@spb.swisstok.ru")
    }
    if larg > 4 {
        DST = os.Args[4]
    }
    var resp string
    if MODE == "status" {
        resp=ocip.OCIPsend(Config,"UserCallForwardingAlwaysGetRequest",ocip.ConcatStr("","userId=",USERID))
    }
    if MODE == "false" {
        resp=ocip.OCIPsend(Config,"UserCallForwardingAlwaysModifyRequest",ocip.ConcatStr("","userId=",USERID),"isActive=false")
    }
    if MODE == "true" {
        resp=ocip.OCIPsend(Config,"UserCallForwardingAlwaysModifyRequest",ocip.ConcatStr("","userId=",USERID),"isActive=true",ocip.ConcatStr("","forwardToPhoneNumber=",DST))
    }
    fmt.Fprint(os.Stdout,resp)
}
