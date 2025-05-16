USE stackexchenege;



LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_21.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_22.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_23.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_24.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_25.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_26.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_27.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_28.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_29.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;
                     
LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_30.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_31.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_32.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_33.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_posthistory/PostHistory_Part_34.xml'
INTO TABLE post_history (
	@Id, @PostHistoryTypeId, @PostId, @RevisionGUID,
    @CreationDate, @UserId, @Text, @ContentLicense
)
SET id = @Id, 
post_history_type_id = @PostHistoryTypeId,
post_id = @PostId, 
revision_guid = @RevisionGUID,
creation_date = @CreationDate, 
user_id = @UserId, 
post_text = SUBSTRING(@Text, 1, 10000),
content_license=@ContentLicense;

commit;
flush tables;