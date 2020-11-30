    package main
    import(
        "fmt"
        "time"
        "math/rand"
        "crypto/md5"
        "encoding/base64"
        "net/http"
        "io/ioutil"
        "log"
        "encoding/json"
        "bytes"
        
       // "math/rand"
    )
    func genAuthenKey(username string, password string, method string) string{
        timestamp:=time.Now().Unix()
        idran := makeId()
        signatureRawData := method + fmt.Sprint(timestamp) + idran
        md5hash:= md5.Sum([]byte(signatureRawData))
        signature := base64.StdEncoding.EncodeToString(md5hash[:])
        authenvalue := signature + ":" + idran + ":" + fmt.Sprint(timestamp) + ":" + username + ":" + password
        return authenvalue

    }
    
        

   func getHeaders(username string, password string, method string) map[string]string {
        var headers map[string]string
        headers=make(map[string]string)
        headers["Content-Type"] = "application/json"
        headers["Accept"] = "application/json"
        headers["Authentication"] = genAuthenKey(username, password, method)
        return headers
    }
        
    func makeId()string{
        strsequence := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
        idran := ""
        for i:=0;i<10;i++{
            idx:= rand.Intn(len(strsequence))
           
            idran += string(strsequence[idx])
        }
        return idran

    }

    type importBody struct {
    Pattern, Serial, XmlData string
    IkeyEmail                map[string]string
}

func post() {
    postBody, _ := json.Marshal(importBody{
        Pattern:   "01GTKT0/001",
        Serial:    "ST/20E",
        XmlData:   "<Invoices> <Inv> <Invoice> <Ikey>A000116038ABC</Ikey> <CusCode>071DC</CusCode> <Buyer/> <CusName><![CDATA[Phan Phuc Doan]]></CusName> <CusAddress>446 Võ Văn Kiệt</CusAddress> <CusBankName/> <CusPhone/> <CusBankNo/> <CusTaxCode/> <PaymentMethod>TM/CK</PaymentMethod> <CurrencyUnit>VND</CurrencyUnit> <ExchangeRate>1</ExchangeRate> <PaymentStatus>1</PaymentStatus> <Email>test@softdreams.vn</Email> <EmailCC>leminhhoho@gmail.com,test2@softdreams.vn</EmailCC> <Products> <Product> <Code/> <ProdName>Bàn học</ProdName> <ProdUnit>Cái</ProdUnit> <ProdQuantity>2</ProdQuantity> <ProdPrice>1050000</ProdPrice> <Total>584320</Total> <VATRate>10</VATRate> <VATAmount>58432</VATAmount> <Amount>642752</Amount> <Extra></Extra> </Product> </Products> <ArisingDate>06/03/2017</ArisingDate> <Total>1145268</Total> <VATRate>10</VATRate> <VATAmount>114527</VATAmount> <Amount>1259795</Amount> <AmountInWords>Một triệu hai trăm năm mươi chín nghìn bảy trăm chín mươi lăm đồng.</AmountInWords> </Invoice> </Inv> </Invoices>",
        IkeyEmail: map[string]string{"2018001": ""},
    })
    responseBody := bytes.NewBuffer(postBody)
    client := &http.Client{}

    req, err := http.NewRequest("POST","http://0316054160.softdreams.vn/api/publish/importInvoice", responseBody)
    req.Header.Set("Content-Type","application/json")
    req.Header.Set("Accept","application/json")
    req.Header.Set("Authentication",genAuthenKey("admin","pass!@#$%","POST"))
    resp, err := client.Do(req)

    if err != nil {
        log.Fatalf("An Error Occured %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    sb := string(body)
    log.Printf(sb)
}
    func main(){
        post()

    }
        

   
        
