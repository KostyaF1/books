package query

const AddAnswer = `
INSERT INTO coments (book_unit_id, body, author, answer_id)
VALUES ($1, $2, $3, $4) RETURNING id;`
