package query

const GetCommByID = `SELECT body || '-author:' || author AS comment, array(SELECT body || '-author:' || author 
FROM comments WHERE answer_id = c.id) AS answer
FROM comments AS c
WHERE id = $1;`
