package gonv

import(
  "github.com/jenazads/goutils";
)

// Gonv object
type Gonv struct {
  enVar map[string]interface{}
}

// Create New Gonv Object
func NewGonv() (*Gonv){
  mv := make(map[string]interface{})
  return &Gonv{
    enVar: mv}
}

// Set new environmental variable
func (gnv *Gonv) SetEnv(tag string, value interface{}) error {
  if val, err := gnv.GetEnv(tag); val == nil || err != nil {
    gnv.enVar[tag] = value
    return nil
  }else{
    return GonvErrAlreadyExist
  }
}

// Get environmental variable
func (gnv *Gonv) GetEnv(tag string) (interface{}, error) {
  if val := gnv.enVar[tag]; val == nil{
    return nil, GonvErrNotFound
  }else{
    return val, nil
  }
}

// Update environmental variable
func (gnv *Gonv) UpdateEnv(tag string, value interface{}) error {
  if val, err := gnv.GetEnv(tag); val == nil || err != nil {
    return GonvErrNotFound
  }else{
    gnv.enVar[tag] = value
    return nil
  }
}

// Get environmental variable and transforms to string
func (gnv *Gonv) GetStringEnv(tag string) (string, error){
  value, err := gnv.GetEnv(tag)
  return goutils.ToString(value), err
}

// Get environmental variable and transforms to int
func (gnv *Gonv) GetIntEnv(tag string) (int, error){
  value, err := gnv.GetEnv(tag)
  return goutils.ToInt(value), err
}

// Get environmental variable and transforms to bool
func (gnv *Gonv) GetBoolEnv(tag string) (bool, error){
  value, err := gnv.GetEnv(tag)
  return goutils.ToBool(value), err
}
