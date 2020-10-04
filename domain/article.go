package domain

type Article struct {
	ID            int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	CreatedBy     string
	ModifiedBy    string
	State         int
}
