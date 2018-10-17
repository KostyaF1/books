package query

const AddComment = `
INSERT INTO coments (book_unit_id, body, author)
VALUES ($1, $2, $3) RETURNING id;`
