//functions in this package

//func FaToint(string1 string)int
//getting a string in format of arabic/persian number like ۱۲۳۴,۴۴۴ and return 12344

//func MyQuery(body io.ReadCloser, url string) []Book
//example of how use of goquery for scraping

//func Myhtml(body io.ReadCloser, names *string)
//example of how use of net/html package for scrapping

//func Myhttp(url string) (io.ReadCloser, error)
//get a url and return a io.ReadCloser variable for use of in other scrapping funcitons

//save data in sqlite database
//func Sq(Books []book)
// get a slice and save them in database
//other functions is:
//func connect_sqlite3(filename string) (db *sql.DB)
//func create_table(db *sql.DB) (sql.Result, error)
//func get_version(db *sql.DB)
//func insert_record(db *sql.DB, Books []Book) (int64, error)

//save data in json file
//func SaveOnFile(uRl string, price string, filename string) error
//func Encoding_data(uRl string, price string) []byte

//send mail
//func SendMail(extracted_price string, url string, product string) error