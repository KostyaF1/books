package query

const (
	DeleteComment = `
DELETE
FROM comments
WHERE id = $1;
`
	DeleteAnswer = `
DELETE
FROM comments
WHERE answer_id = $1;
`
)
