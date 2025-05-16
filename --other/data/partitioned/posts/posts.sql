USE stackexchenege;




LOAD XML
INFILE 'C:/posts_split/Posts_Part_1.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_2.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_3.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_4.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_5.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_6.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_7.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_8.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_9.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_10.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_11.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_12.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_13.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_14.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_15.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_16.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_17.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_18.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_19.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;

LOAD XML
INFILE 'C:/posts_split/Posts_Part_20.xml'
INTO TABLE posts (
	@Id, @PostTypeId, @CreationDate, @ParentID, @AcceptedAnswerId, 
	@Score, @ViewCount, @Body, @OwnerUserId, @LastEditorUserId,
    @LastEditDate, @LastActivityDate, @Title, @Tags, @AnswerCount,
    @CommentCount, @ContentLicense
)
SET id = @Id, 
post_type_id = @PostTypeId, 
creation_date = @CreationDate, 
parent_id = @ParentID,
accepted_answer_id = @AcceptedAnswerId,
score = @Score,
view_count = @ViewCount,
post_body = SUBSTRING(@Body, 1, 10000),
owner_user_id = @OwnerUserId,
last_editor_user_id = @LastEditorUserId,
last_edit_date = @LastEditDate,
last_activity_date = @LastActivityDate,
post_title = @Title,
tags = @Tags,
answer_count = @AnswerCount,
comment_count = @CommentCount,
content_license = @ContentLicense;

commit;
flush tables;