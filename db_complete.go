package main 

import (
		 pg "github.com/go-pg/pg"
		 orm"github.com/go-pg/pg/orm"
		 "log"
		 "os"
		 "fmt")

type ProductItem struct {
	RefPointer int `sql:"-"`
	
	tableName struct{}`sql:"test_again_and_again"`
	
	ID int `sql:"id,pk"`
	
	Name string `sql:"name,unique"`

}
func CreateProdItemsTable(db *pg.DB) error {
	
	fmt.Println("Code was here ")

	opts:= &orm.CreateTableOptions{
		IfNotExists: true,
	}   
	
	createErr:= db.CreateTable(&ProductItem{},opts)
	
	if createErr != nil{
		log.Printf("Error while creating  table  productItems,Reason:%v\n",createErr)
		return createErr
	}
	
	log.Printf("Table test created successful.\n")

	return nil 
}
//IMPORTANT THIS IS A METHOD OF THE STRUCT 
func (pi *ProductItem)Save(db *pg.DB)error{
	insertErr:= db.Insert(pi)
	if insertErr != nil{
		log.Printf("Error while inserting new item into db ,Reason:%v\n",insertErr)
		return insertErr
	}
	//log.Prinf("ProductItem %s inserted successfully.\n",pi.Name)
	return nil 
}
//if we have to update multiple records 
//purchase history of product 
//purchase history of user 


func connect()*pg.DB{

	opts:=&pg.Options{

		User:	  "postgres",
		Password: "hello2u!",
		Addr:	  "localhost:5432",
		Database:"my_db",
		DialTimeout: 30 *time.Second,
		ReadTimeout: 1 * time.Minute,
		WriteTimeout : 1 *time.Minute,
		IdleTimeout: 30 * time.Minute,// drivers are connected for thirty minutes no response times out 
		MaxAge :	1 * time.Minute,//ideal time out nothing happens in  one minute ping the postgres let postgres know connection is there postgres does not closes .
		poolSize: 200,// number of connections db can access.


	}
	var db *pg.DB = pg.Connect(opts)
	fmt.Println(db)

	if db == nil{

		log.Printf("failed to connect to database ")
		os.Exit(100)

	}else  {
		log.Printf("Coonection to database successful.\n")	
	}
	CreateProdItemsTable(db)
	
	return db
}
//Save product calls this function of struct 
func (pi * ProductItem)SaveMultiple(db *pg.DB,items []*ProductItem)error{
	_,insertErr:=db.Model(items[0],items[1]).Insert()
	if insertErr!= nil{
		fmt.Println("Error while inserting bulk items Reason %v\n",insertErr)
		return insertErr
		
	}
	log.Printf("Bulk insert successful\n")
	return nil
}
func SaveProduct(dbRef * pg.DB){
	newPI:=&ProductItem{
		ID: 30,
		Name: "drake",

	}
	newPI1:=&ProductItem{
		ID: 40,
		Name: "elli",	

	}
	//newPI.Save(dbRef)

	totalItems:=[]*ProductItem{newPI,newPI1 }
	fmt.Printf("code was here save  ")
	//saves an array of struct  to  postgeresql
	newPI.SaveMultiple(dbRef,totalItems)

}
type Params struct {

	Param1 string 
	Param2 string 

}


func PlaceHolderDemo(db * pg.DB)error{
	var value int
	
	params:=Params{
		Param1:"This is param1",
		Param2:"This is param2",
	} 
	
	var query string ="SELECT ? param2"
	
	_,selectErr:=db.Query(pg.Scan(&value),query,params)
	if selectErr != nil{
		log.Printf("Error while running the select query ,Reason: %v\n",selectErr)
		return selectErr   
	}
	log.Printf("Scan successful,Scanned value :%s\n",value)
	return nil 

}
func (pi * ProductItem)DeleteItem (db *pg.DB)error {
		_,deleteErr:=db.Model(pi).Where(	"id = ?id ").Delete()
	if deleteErr != nil {

		log.Printf("Error while deleting %v\n ",deleteErr)
		return deleteErr
	}
	log.Printf("Delete successful,Scanned value :%s\n",pi.ID)
	return nil 
	

}
func DeleteProduct (dbRef *pg.DB){

	newPI:=&ProductItem{
		//Name: "amman  habib",
		ID:30,
		Name:"drake",
	}
	newPI.DeleteItem(dbRef)
}

func (pi * ProductItem)Updatename (db *pg.DB )error {
	_,updateErr:=db.Model(pi).Set("name  = ?name" ).Where("id = ?id ").Update()
	if  updateErr != nil {
		log.Printf("Error while updating name .Reason :%v\n",updateErr)
		return updateErr
	}
		log.Printf("Name  updated successfully for ID %d\n",pi.ID)
	return nil 
}
func UpdateItemName(dbRef * pg.DB){
	newPI:=&ProductItem{
		ID:40,
		Name:"elli",

	}

newPI.Updatename(dbRef)
}
func (pi * ProductItem )GetByID(db *pg.DB )error {
	getErr:= db.Model(pi).Column("id","name").
	Where("id = ? 0",pi.ID).
	Where("name = ? 0",pi.Name).Select()

	//Below code is for the retrieval of a particular struct from the db 
	//getErr:=db.Select(pi)
	if getErr!= nil{
		log.Printf("Error  while getting value  by id ,Reason:%v\n",getErr)
		return getErr
	}
	log.Printf("get by id successfully for  %v\n",*pi )
	return nil 
}	
func GetByID(dbRef * pg.DB){
	newPI:= &ProductItem{

		ID:20,
		Name:"nio",

	}
	newPI.GetByID(dbRef)

}

func main() {


fmt.Println("fmt")
pg_db :=connect()
//SaveProduct(pg_db)
//PlaceHolderDemo(pg_db)
//DeleteProduct(pg_db)

//UpdateItemName(pg_db)
 GetByID(pg_db)
}
