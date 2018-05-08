package main

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)
type Users struct {
        Name string
        Email string
        Phone string
}
func main() {
	http.HandleFunc("/",handleRequest)
	http.HandleFunc("/saveDetails",saveDetailsApi)
	http.HandleFunc("/viewDetails", getUserDetail)
	http.ListenAndServe("localhost:8900",nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request){
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers","*")
	}
	if r.Method == "OPTIONS" {
            return
    }
    //http.ServeFile(w, r, r.URL.Path[1:])
}

func connectToDb() *mgo.Collection{
	session, err := mgo.Dial("mongodb://localhost:27017/testingMongo")
    if err != nil {
        panic(err)
    }
    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    database:="testingMongo"
    collection:="users"
    
    c := session.DB(database).C(collection)
    return c
}

func saveDetailsApi(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers","*")
	}
	if r.Method == "OPTIONS" {
            return
    }
    name := r.FormValue("name")
    email := r.FormValue("email")
    phone := r.FormValue("phone")
    
    c := connectToDb();
   	
    err := c.Insert(&Users{name,email,phone})
    if err != nil {
            log.Fatal(err)
    }else{
    	fmt.Println("Data saved successfully!")
    }
}

func getUserDetail(w http.ResponseWriter, r *http.Request) {
    if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers","*")
	}
	if r.Method == "OPTIONS" {
            return
    }
    email := r.FormValue("email")
    response,err:=getData(email)
	if err != nil{
		panic(err)
	}

	//fmt.Println("it is ",string(response))
	fmt.Fprintf(w,string(response))
    
}

func getData(email string) ([]byte,error){
    c := connectToDb()
	var result Users
    err := c.Find(bson.M{"email": email}).One(&result)
    if err != nil {
            log.Fatal(err)
    }

    details := make(map [string] string)
    details["Name"] = result.Name
    details["Phone"] = result.Phone

    return json.MarshalIndent(details,""," ")
}