package gonv_test

import(
  "fmt"
  "github.com/jenazads/gonv"
)

func main(){
  host := gonv.GetStringEnv("HOST", "localhost")
  port := gonv.GetIntEnv("PORT", 3306)
  debug := gonv.GetBoolEnv("DEBUG", false)
  hostname:= gonv.GetEnvOrCreate("USER", "")

  final := fmt.Sprintf("%s:%d debug mode : %t", host, port, debug)
  fmt.Println(final, "hostname:", hostname)
}
