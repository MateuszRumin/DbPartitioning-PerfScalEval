USE stackexchenege;


LOAD XML
INFILE 'C:/datafill//PostLinks.xml'
INTO TABLE post_links (
	@Id, @CreationDate, @PostId, @RelatedPostId, @LinkTypeId
)
SET id = @Id, 
creation_date = @CreationDate, 
post_id = @PostId, 
related_post_id = @RelatedPostId,
link_type_id = @LinkTypeId;
                     