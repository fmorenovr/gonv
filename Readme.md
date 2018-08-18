# Golang + Environment = gonv

gonv (Golang eNVironment) is a Golang implementation to manage environment variables in system.  
You can see an extended doc in [godocs](https://godoc.org/github.com/Jenazads/gonv).

Install it writing in terminal:

    go get github.com/jenazads/gonv

Example:

  ```go
    // supose that this is your struct
    type MyDataType struct{
      ID int
      name string
    }

    // your comparator, you cna use this function as TypeComparator type function
    func MyComparator(a, b interface{}) int {
      AaAsserted := a.(MyDataType)
      bAsserted := b.(MyDataType)
      switch {
        case aAsserted.ID > bAsserted.ID:
          return 1
        case aAsserted.ID < bAsserted.ID:
          return -1
        default:
          return 0
      }
    }
  ```
