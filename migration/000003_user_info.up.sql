alter table Member
    add column bio         longtext     null,
    add column username    varchar(255) null,
    add column firstName   varchar(255) null,
    add column lastName    varchar(255) null,
    add column phoneNumber varchar(255) null,
    add column type        varchar(255) null
;

alter table ChannelHasMember
    add column joinDate datetime not null
;
