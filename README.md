[![Build Status](https://travis-ci.org/wawandco/whatsgo.svg?branch=master)](https://travis-ci.org/wawandco/whatsgo)

# Whatsgo
A package that allows you send Whatsapp messages using [Whatools](http://wha.tools) API.

## Overview
Whatsgo is a **Go** package that acts as a wrapper to communicate with [Whatools](http://wha.tools) API. With **Whatsgo** you can:
- Subscribe to your Whatools API account.
- Send messages to another Whatsapp users.
- Set and Get your Whatsapp nickname.
- Set and Get the status message for your account.

## Usage
### Installing
Using Whatsgo  is easy. First use `go get`command to install the latest version of the package.

```bash
$ go get github.com/wawandco/whatsgo
```
Next include Whatsgo in yoour application.
```go
import "github.com/wawandco/whatsgo"
```

### Subscribe to your Whatools API account
Before use any of the operations available in Whatsgo you have to subscribe against [Whatools API](https://api.wha.tools/v2). **Whatsgo** provides a function for you to get subscription ready.

```go
wgo := whatsgo.Subscribe("MYWHATOOLSAPIKEY") // e.g abcd1234-abb0-123c-1fe1-abcdef012345
if wgo != nil {
  // do something cool with wgo.
}
```
If subscription was successful `Subscribe` function will return a `WhatsGo` type struct so you can deal with all operations available in the package.

#### WhatsGo struct
Represents a successful data response sent by **Wha.tools** subscribe API endpoint.
```go
type WhatsGo struct {
	Key         string    // API Key provided by whatools
	Status      string    // User account status: active, existing, new
	CountryCode string    // self explanatory, Country code
	PhoneNumber string    // self explanatory, phone number
}
```
### Send messages to another Whatsapp users
Once you get a `WhatsGo` reference you can use its `SendMessage` method to send messages to another Whatsapp users.
```go
msg = &whatsgo.Message{"573001234567", "Hello, from whatsgo", true}
wgo.SendMessage(msg)
```
Note that we are passing a reference of `Message` type. This struct represents a message to send:

```go
type Message struct {
	To    string //phone number with country code associated, e.g 573001231232
	Body  string //text to send
	Honor bool   //(true/false) This flag prevents the phone number from being formatted again if it is already in international form.
}
```
You can send a group of messages in a one method call:
```go
msg1 = &whatsgo.Message{"573001234567", "Hello, from whatsgo", true}
msg2 = &whatsgo.Message{"3001232345", "Hello, from whatsgo", false}
wgo.SendMessage(msg1, msg2, msg1, msg2)
```

### Set and Get your Whatsapp nickname
You can get your _Whatsapp_ nickname by calling `GetNickname()` method.
```go
  wgo.GetNickname() // e.g Wawandco, Whatsgo, or Jhon Doe.
```
You can set your _Whatsapp_ nickname by calling `SetNickname(string)` method.
```go
  wgo.SetNickname("My cool nickname")
```

### Set and Get the status message for your account.
You can get the status message for your account by calling `GetStatusMessage()` method.
```go
  wgo.GetStatusMessage() // e.g I think therefore I am.
```

You can set the status message for your account by calling `SetStatusMessage(string)` method.
```go
  wgo.SetStatusMessage("My new status")
```

### Unsubscribe after you're done.
When you are done sending any messages, updating your nickname, or changing your status message, you can turn off client access to whatools API by calling `Unsubscribe` method. This method will disable any request to API endpoint with your existing API key.
**Note:** You have to access to wha.tools dashboard to enable your API key if you use this method before. This is used for security reasons, so use it wisely.
```go
wgo.Unsubscribe()
```

## Contributing
1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request
