## [UNRELEASED]
### Added
- Added end-point `/v1/storage/getChannelList`
- Added end-point `/v1/storage/getMemberList`

## [1.1.0] - 2020-10-05
### Added
- Added migration `000002_channel_has_member`.
- Added migration `000003_user_info`.
- Added store member info `username`, `firstName`, `lastName`, `phoneNumber`, `type`, `bio` when running `fetch-members` command.
- Added end-point `/proxy` for full proxy request to tdlib.
- Added migration `000004_user_photo`.
- Now member photo links store to database.
### Changed
- Now command `fetch-members` can store relation between channels and members
- Refactoring tab to spaces
- Response end-point `/v1/getUser`. New keys for `photo_list` remote uniq id and remote id for.

## [1.0.0] - 2020-09-25
### Added
- Add `mysql` docker container
- Add package `migration` to docker container
- Add migration `000001_init_schema`
- Add port to go docker container for dlv debug
- Add script to dockerfile mysql for run migration when up docker-compose
- Add command `fetch-members` for parse channel members and save to storage
- Add handle update `updateSupergroup` for save `supergroup_id` to storage
- Add end-point `/v1/getChannelInfo?channel_id={id}`
- Add end-point `/v1/getChannel?channel_id={id}`
- Add end-point `/v1/getMembers?channel_id={id}`
- Add end-point `/v1/getPhoto?photo_id={id}`
- Add end-point `/v1/getUser?user_id={id}`
- Add wrapper mysql connect
- Add command `run-server` for run api server
### Fixed
- Fixed eternal subscription for method `PUT /receipt`.
- Fix handle `SIGINT` signal for safe close tdlib client
### Changed
- Refactoring
- Send `device_model` and `system_version` when create tdlib client instance

## [0.1.0] - 2020-08-25
### Added
- Added skeleton application
- Added API end-point `/getMe`
- Added show updates
