package sqlite

const getBlockQuery = `
select pos,data from blocks b where b.pos = ?
`
