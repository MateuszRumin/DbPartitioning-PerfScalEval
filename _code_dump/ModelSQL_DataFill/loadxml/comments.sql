USE testdb;



LOAD XML
INFILE '/var/lib/mysqlfiles/Comments.xml'
INTO TABLE comments (
	@Id, @PostId, @Score, @Text, 
    @CreationDate, @UserId, @ContentLicense
)
SET id = @Id, 
post_id = @PostId, 
score = @Score, 
comment_text = @Text, 
creation_date=@CreationDate, 
user_id=@UserId, 
content_license=@ContentLicense;
                     