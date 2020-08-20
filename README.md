### Telegram parser
This repository golang wrapper for [Telegram Database Library](https://core.telegram.org/tdlib).
Contains docker build, web server and REST api interface for call C functions of tdlib.

### Installation
You must prepend install [Docker Desktop](https://www.docker.com/get-started).
```bash
git clone https://github.com/revilon1991/tg-parser.git
cd tg-parser
cp .env.dist .env
```
Create [Telegram Application](https://core.telegram.org/api/obtaining_api_id) and get
`TELEGRAM_APP_ID` `TELEGRAM_APP_HASH` `TELEGRAM_BOT_TOKEN`

After create [Telegram bot](https://core.telegram.org/bots#3-how-do-i-create-a-bot) and get `TELEGRAM_BOT_TOKEN`
set it to `.env` and run
```bash
docker-compose up -d
docker-compose exec go go get -u github.com/joho/godotenv
```
> INFO
> Build docker image it is long process. Wait ~40 minutes. You can see progress bar from terminal.

### Usage
```bash
docker-compose exec go go run src/main.go src/api.go src/init.go src/structure.go
```
You see updates from telegram from terminal.
***
Use browser for view information your bot [localhost:8080/getMe](http://localhost:8080/v1/getMe)
 