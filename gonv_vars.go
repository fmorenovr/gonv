package gonv

import(
  "os";
  "github.com/jenazads/goutils";
)

// Set new environmental variable
func SetEnv(tag string, value interface{}) error {
  if val := os.Getenv(tag); val == ""{
    err := os.Setenv(tag, goutils.ToString(value))
    return err
  }else{
    return GonvErrAlreadyExist
  }
}

// Get environmental variable
func GetEnv(tag string) (string, error) {
  if val := os.Getenv(tag); val == ""{
    return "", GonvErrNotFound
  }else{
    return val, nil
  }
}

// Update environmental variable
func UpdateEnv(tag string, value interface{}) error {
  if val := os.Getenv(tag); val == ""{
    return GonvErrNotFound
  }else{
    err:=os.Setenv(tag, goutils.ToString(value))
    return err
  }
}

// Get environmental variable and transforms to string
func GetStringEnv(tag string) (string, error){
  value, err := GetEnv(tag)
  return goutils.ToString(value), err
}

// Get environmental variable and transforms to int
func GetIntEnv(tag string) (int, error){
  value, err := GetEnv(tag)
  return goutils.ToInt(value), err
}

// Get environmental variable and transforms to bool
func GetBoolEnv(tag string) (bool, error){
  value, err := GetEnv(tag)
  return goutils.ToBool(value), err
}
