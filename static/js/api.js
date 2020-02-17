
export function getColormapping(){
  return m.request("api/colormapping");
}

export function getMapblock(posx, posy, posz){
  return m.request("api/viewblock/"+posx+"/"+posy+"/"+posz);
}
