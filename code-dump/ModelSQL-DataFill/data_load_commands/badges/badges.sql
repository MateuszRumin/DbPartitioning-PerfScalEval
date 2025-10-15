USE stackexchenege;


LOAD XML
INFILE 'C:/datafill//Badges.xml'
INTO TABLE badges (
	@Id, @UserID, @Name, @Date, @Class, @TagBased
)
SET id = @Id,
user_id = @UserID,
badge_name = @Name,
badge_date = @Date,
class=@Class,
tag_based=@TagBased;

                     