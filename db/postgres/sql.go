package postgres

const getBlockQuery = `
select posx,posy,posz,data from blocks b
where b.posx = $1
and b.posy = $2
and b.posz = $3
`
