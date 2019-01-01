package main

import (
    "fmt"
    "syscall"
    "time"

    "github.com/c0re100/go-tdlib"
    "golang.org/x/crypto/ssh/terminal"
)

type Client struct {
    client    *tdlib.Client
    UID       int32
}

func main() {
    tdlib.SetLogVerbosityLevel(1)
    tdlib.SetFilePath("./errors.txt")

    f9 := &Client{
        client: tdlib.NewClient(tdlib.Config{
            APIID:               "132712",
            APIHash:             "e82c07ad653399a37baca8d1e498e472",
            SystemLanguageCode:  "en",
            DeviceModel:         "F9TelegramUtils",
            SystemVersion:       "1.0",
            ApplicationVersion:  "1.0",
            UseMessageDatabase:  true,
            UseFileDatabase:     true,
            UseChatInfoDatabase: true,
            UseTestDataCenter:   false,
            DatabaseDirectory:   "./tdlib-db",
            FileDirectory:       "./tdlib-files",
            IgnoreFileNames:     false,
        }),
    }

    for {
        currentState, _ := f9.client.Authorize()
        if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {
            fmt.Print("Enter phone: ")
            var number string
            fmt.Scanln(&number)
            _, err := f9.client.SendPhoneNumber(number)
            if err != nil {
                fmt.Printf("Error sending phone number: %v", err)
            }
        } else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
            fmt.Print("Enter code: ")
            var code string
            fmt.Scanln(&code)
            _, err := f9.client.SendAuthCode(code)
            if err != nil {
                fmt.Printf("Error sending auth code : %v", err)
            }
        } else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPasswordType {
            fmt.Print("Enter Password: ")
            bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
            if err != nil {
                fmt.Println(err)
            }
            _, err = f9.client.SendAuthPassword(string(bytePassword))
            if err != nil {
                fmt.Printf("Error sending auth password: %v", err)
            }
        } else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
            me, err := f9.client.GetMe()
            if err != nil {
                fmt.Println(err)
                return
            }
            f9.UID = me.Id
            fmt.Println("Hello!", me.FirstName, me.LastName, "("+me.Username+")")
            break
        }
    }

    go f9.StatusHook()
    go f9.MessageHook()

    fmt.Println("You're Always Online now.")
    for {
        f9.AlwaysOnline()
        time.Sleep(30 * time.Second)
    }
}
