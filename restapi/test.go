package main

// import (
//     "database/sql"
//     _ "github.com/go-sql-driver/mysql"
//     "log"
// )


// type Tag struct {
//   ID   int    `json:"id"`
//   Name string `json:"name"`
// }
// func main4() {
//   // Open up our database connection.
//   db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/testdb")

//   // if there is an error opening the connection, handle it
//   if err != nil {
//       log.Print(err.Error())
//   }
//   defer db.Close()

//   // Execute the query
// //   creation of table manually
// //   results, err := db.Query("CREATE TABLE Employees ( id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, name VARCHAR(20) NOT NULL, post VARCHAR(20) NOT NULL )")

//     // results, err := db.Query("")


//   if err != nil {
//       panic(err.Error()) // proper error handling instead of panic in your app
//   }

//   for results.Next() {
//       var tag Tag
//       // for each row, scan the result into our tag composite object
//       err = results.Scan(&tag.ID, &tag.Name)
//       if err != nil {
//           panic(err.Error()) // proper error handling instead of panic in your app
//       }
//               // and then print out the tag's Name attribute
//       log.Printf(tag.Name)
//   }

// }