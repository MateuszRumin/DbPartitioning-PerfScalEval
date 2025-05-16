USE stackexchenege;


LOAD XML
INFILE 'D:/split_files_votes/Votes_Part_1.xml'
INTO TABLE votes (
	@Id, @PostId, @VoteTypeId, @CreationDate
)
SET id = @Id, 
post_id = @PostId, 
vote_type_id = @VoteTypeId, 
creation_date = @CreationDate;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_votes/Votes_Part_2.xml'
INTO TABLE votes (
	@Id, @PostId, @VoteTypeId, @CreationDate
)
SET id = @Id, 
post_id = @PostId, 
vote_type_id = @VoteTypeId, 
creation_date = @CreationDate;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_votes/Votes_Part_3.xml'
INTO TABLE votes (
	@Id, @PostId, @VoteTypeId, @CreationDate
)
SET id = @Id, 
post_id = @PostId, 
vote_type_id = @VoteTypeId, 
creation_date = @CreationDate;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_votes/Votes_Part_4.xml'
INTO TABLE votes (
	@Id, @PostId, @VoteTypeId, @CreationDate
)
SET id = @Id, 
post_id = @PostId, 
vote_type_id = @VoteTypeId, 
creation_date = @CreationDate;

commit;
flush tables;

LOAD XML
INFILE 'D:/split_files_votes/Votes_Part_5.xml'
INTO TABLE votes (
	@Id, @PostId, @VoteTypeId, @CreationDate
)
SET id = @Id, 
post_id = @PostId, 
vote_type_id = @VoteTypeId, 
creation_date = @CreationDate;

commit;
flush tables;