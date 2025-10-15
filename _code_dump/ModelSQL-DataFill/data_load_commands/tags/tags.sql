USE stackexchenege;


LOAD XML
INFILE 'C:/datafill/Tags.xml'
INTO TABLE tags (
	@Id, @TagName, @Count, @ExcerptPostId, @WikiPostId
)
SET id = @Id, 
tag_name = @TagName, 
tag_count = @Count, 
except_post_id = @ExcerptPostId,
wiki_post_id = @WikiPostId;
                     