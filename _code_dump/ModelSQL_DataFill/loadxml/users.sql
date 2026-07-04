USE testdb;


LOAD XML
INFILE '/var/lib/mysqlfiles/Users.xml'
INTO TABLE users (
	@Id, @Reputation, @CreationDate, @DisplayName, 
    @LastAccessDate, @WebsiteUrl, @Location, @AboutMe, 
    @Views, @UpVotes, @DownVotes, @AccountId
)
SET id = @Id, 
reputation = @Reputation, 
creation_date = @CreationDate, 
display_name = @DisplayName,
last_access_date = @LastAccessDate,
website_url = @WebsiteUrl,
location = @Location,
about_me = @AboutMe,
views = @Views,
upvotes = @UpVotes,
downvotes = @DownVotes,
account_id = @AccountId;