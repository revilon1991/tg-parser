alter table Member
    drop column bio,
    drop column username,
    drop column firstName,
    drop column lastName,
    drop column phoneNumber,
    drop column type
;

alter table ChannelHasMember
    drop column joinDate
;
