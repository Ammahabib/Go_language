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
func connect()*pg.DB{

	opts:=&pg.Options{

		User:	  "postgres",
		Password: "hello2u!",
		Addr:	  "localhost:5432",
		Database:"my_db",
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
func SaveProduct(dbRef * pg.DB){
	newPI:=&ProductItem{
		ID: 20,
		Name: "amman",
	}
	newPI.Save(dbRef)

}

func main() {


fmt.Println("fmt")
pg_db:=connect()
SaveProduct(pg_db)

}
