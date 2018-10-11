package query

const GetCommByID = `SELECT body || '-author:' || author, array(SELECT body || '-author:' || author FROM coments WHERE answer_id = c.id)
FROM coments AS c
WHERE id = $1;`
