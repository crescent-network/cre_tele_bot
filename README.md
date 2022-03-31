# cre_tele_bot

This is to answer to the FAQ automatically in telegram chat room with buttons and commands.

This bot is based on https://github.com/tucnak/telebot

where the API documentations are found in https://pkg.go.dev/gopkg.in/telebot.v3

Before running, please install the package as below.

```bash
go get -u gopkg.in/telebot.v3
```

In order to use telegram bot, a new telegram bot is needed. Or, the existing telegram bot can be used. The telegram bot can be created as explained in https://core.telegram.org/bots

or you can add `@BotFather` in telegram to create the bot. When a new bot is created via `@BotFather`, a `new key` for the bot is generated. 
The `new key` value needs to be written in `msg/config.go` as `TelegramToken_Address`. Please be careful not to expose the key value outside of your group.

There are two types of the user inputs.

- Button: when at least one user in the chat room inputs `/faq`, the buttons start to be displayed at the bottom of the chat room. Then, any user in the chat room can push the buttons so that they can see the response from the pushed button. The button names and the arrangement can be adjusted in `button/btnHandler.go`.

- Command: A user can push `/` in the chat room, then the command list will be seen. The user can push a command in the list, or directly input the command, e.g., `/teaser`. The command names and the order can be adjusted in `command/command.go`.

The response message can be modified at `msg/config.go`.
