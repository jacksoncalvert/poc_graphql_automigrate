package psql_orm

import (    
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "fmt"
)

type HellOr struct {
    ID string
    Msg string
}


func MakeGorm() {
    dsn := get_dsn()
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&HellOr{})

    // Create
    db.Create(&HellOr{ID: "D42",Msg: "HellOr Or Team"})

}


func SearchDbHello() (*HellOr) {
    dsn := get_dsn()
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    var hellor HellOr
    db.First(&hellor, "id = ?", "D42")

    return &hellor
}


func get_dsn() (string){
    host := os.Getenv("PSQL_HOST")
    if host == "" {
        host = "host.docker.internal"
    }

    user := os.Getenv("PSQL_USER")
    if user == "" {
        user = "postgres"
    }


    password := os.Getenv("PSQL_PASSWORD")
    if password == "" {
        password = "postgres"
    }


    dbname := os.Getenv("PSQL_DB")
    if dbname == "" {
        dbname = "postgres"
    }

    port := os.Getenv("PSQL_PORT")
    if port == "" {
        port = "5432"
    }

    return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
}
