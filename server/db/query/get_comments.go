package query

const GetComments = `SELECT body || '-author:' || author, array(SELECT body || '-author:' || author FROM coments WHERE answer_id = c.id)
FROM coments AS c
WHERE book_unit_id = $1;`
