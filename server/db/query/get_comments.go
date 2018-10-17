package query

const GetComments = `SELECT body || '-author:' || author, array(SELECT body || '-author:' || author 
FROM comments WHERE answer_id = c.id)
FROM comments AS c
WHERE book_unit_id = $1;`
