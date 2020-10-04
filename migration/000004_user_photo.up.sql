create table MemberPhoto
(
    id        int          not null auto_increment primary key,
    memberId  int          not null,
    link      varchar(255) not null,
    createdAt datetime     not null,
    updatedAt datetime     not null
);
