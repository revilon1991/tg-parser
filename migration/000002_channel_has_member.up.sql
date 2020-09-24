create table ChannelHasMember
(
    id        int      not null auto_increment primary key,
    channelId int      not null,
    memberId  int      not null,
    createdAt datetime not null,
    updatedAt datetime not null
);
CREATE UNIQUE INDEX uniqChannelHasMember ON ChannelHasMember (channelId, memberId);
