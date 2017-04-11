# Open Whisk Library for Go

This is a small library to trigger Bluemix OpenWhisk actions from your Go apps.

## Installation

```
$ go get github.com/JulzDiverse/openwhisk
```

## Example

```
func main(){
  whisk := openwhisk.New(
     "api_endpoint" //"https://openwhisk.ng.bluemix.net/api/v1"
     "token-string",
     "your_bluemix_namespace", //replace '@' with '%40'
  )   


  whisk.TriggerAction("hello", `{"payload":"json"}`)
}
```
