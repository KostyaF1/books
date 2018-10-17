package query

const AddRating = `INSERT INTO rating (store_unit_id, value)
VALUES ($1, $2) RETURNING id;`
