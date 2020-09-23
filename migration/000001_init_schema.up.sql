create table Member
(
    id        int      not null auto_increment primary key,
    userId    int      not null,
    createdAt datetime not null,
    updatedAt datetime not null
);

CREATE UNIQUE INDEX uniqUserId ON Member (userId);

create table Channel
(
    id          int          not null auto_increment primary key,
    channelId   int          not null,
    username    varchar(255) null,
    description longtext     null,
    memberCount int          not null default 0,
    createdAt   datetime     not null,
    updatedAt   datetime     not null
);

CREATE UNIQUE INDEX uniqChannelId ON Channel (channelId);
