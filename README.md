# Open Whisk Library for Go 

This is a small library to trigger Bluemix OpenWhisk actions from your Go apps. 

## Installation

``` 
$ go get github.com/JulzDiverse/openwhisk 
```

## Example 

```
func main(){
  ow := openwhisk.New(
     "api_endpoint" //"https://openwhisk.ng.bluemix.net/api/v1"
     "token-string",
     "your_bluemix_namespace",
  )   
 

  ow.TriggerAction("hello", `{"payload":"json"}`)
}
```



