### Telegram parser
This repository golang wrapper for [Telegram Database Library](https://core.telegram.org/tdlib).
Contains docker build, web server and REST api interface for call C functions of tdlib.

### Installation
Download latest [release](https://github.com/revilon1991/tg-parser/releases) for your OS

Create [Telegram Application](https://core.telegram.org/api/obtaining_api_id) and get
`TELEGRAM_APP_ID` `TELEGRAM_APP_HASH`

After create [Telegram bot](https://core.telegram.org/bots#3-how-do-i-create-a-bot) and get `TELEGRAM_BOT_TOKEN`
set it to [.env](./.env.dist).
Locate it beside with binary

### Installation with docker from source
You must prepend install [Docker Desktop](https://www.docker.com/get-started).
```bash
git clone https://github.com/revilon1991/tg-parser.git
cd tg-parser
cp .env.dist .env
```
Create [Telegram Application](https://core.telegram.org/api/obtaining_api_id) and get
`TELEGRAM_APP_ID` `TELEGRAM_APP_HASH`

After create [Telegram bot](https://core.telegram.org/bots#3-how-do-i-create-a-bot) and get `TELEGRAM_BOT_TOKEN`
set it to `.env` and run
```bash
docker-compose up -d
docker-compose exec go make
docker-compose exec go make install
```
> INFO
> Build docker image it is long process. Wait ~40 minutes. You can see progress bar from terminal.

### Usage
```bash
./tg-parser [command]

# or if installation with docker from source
docker-compose exec go go run cmd/tg-parser/main.go [command]
```
##### Commands:
```bash
# API server
# You can see updates telegram from terminal.
./tg-parser run-server
```
```bash
# Fetch members id from channels
./tg-parser fetch-members
```

##### API end-points:
* [/getMe](http://localhost:8080/v1/getMe) Information your bot
* [/getChannelInfo](http://localhost:8080/v1/getChannelInfo?channel_id={id}) Information channel
* [/getChannel](http://localhost:8080/v1/getChannel?channel_id={id}) Short information by channel id
* [/getMembers](http://localhost:8080/v1/getMembers?channel_id={id}) Members by channel id
* [/getUser](http://localhost:8080/v1/getUser?user_id={id}) Get user photo by photo id
* [/getPhoto](http://localhost:8080/v1/getPhoto?photo_id={id}) Get user photo by photo id

> INFO  
> When you call end-point, api server must be running   
> `channel_id` - get from database after invite bot to channel  
> `photo_id` - get from `/getUser` end-point
  