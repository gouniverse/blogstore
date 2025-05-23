package blogstore

type StoreInterface interface {
	AutoMigrate() error
	EnableDebug(debug bool) StoreInterface

	PostCount(options PostQueryOptions) (int64, error)
	PostCreate(post *Post) error
	PostDelete(post *Post) error
	PostDeleteByID(postID string) error
	PostFindByID(id string) (*Post, error)
	PostList(options PostQueryOptions) ([]Post, error)
	PostSoftDelete(post *Post) error
	PostSoftDeleteByID(postID string) error
	PostTrash(post *Post) error
	PostUpdate(post *Post) error
}
