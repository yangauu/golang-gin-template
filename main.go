package main

import "go-template/app"

func main() {
	app.RunApp()
}

// 检测值是否为空
// param := ""
// if len(u.Username) > 0 {
// 	param = param + "username=?,"
// }
// if u.Age > 0 {
// 	param = param + "age=?,"
// }
// if len(u.Sex) > 0 {
// 	param = param + "sex=?,"
// }
// if len(u.Address) > 0 {
// 	param = param + "address=?"
// }
// check := param[len(param)-1 : len(param)-0]
// if check == "," {
// 	param = param[0 : len(param)-1]
// }
// log.Println(param)
