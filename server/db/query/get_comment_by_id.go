package query

const GetCommByID = `SELECT body || '-author:' || author AS comment, array(SELECT body || '-author:' || author 
FROM coments WHERE answer_id = c.id) AS answer
FROM coments AS c
WHERE id = $1;`
