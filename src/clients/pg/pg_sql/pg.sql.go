package pg_sql

var SeedPosts = struct {
	TruncatePosts string
	ResetPostsId  string
	InsertPosts   string
}{
	TruncatePosts: `truncate posts;`,
	ResetPostsId:  `alter sequence posts_id_seq restart;`,
	InsertPosts: `
	insert into posts
	(title, text)
	values
	('first post', 'my first post'),
	('second post', 'my second post'),
	('third post', 'my third post');
`,
}

var Posts = struct {
	GetAll     string
	InsertPost string
}{
	GetAll: `select * from posts`,
	InsertPost: `
	insert into posts
	(title, text)
	values
	($1, $2)
	returning *;
	`,
}
