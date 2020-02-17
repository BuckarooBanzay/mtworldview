import { getColor } from './colormapping.js';

var materialCache = {};

export function getMaterial(nodeName){
  var material = materialCache[nodeName];
  if (!material) {
    var colorObj = getColor(nodeName);

    if (!colorObj){
      return;
    }

    var color = new THREE.Color( colorObj.r/256, colorObj.g/256, colorObj.b/256 );
    material = new THREE.MeshBasicMaterial( { color: color } );

		if (nodeName == "default:water_source"){
			material.transparent = true;
			material.opacity = 0.5;
		}

    materialCache[nodeName] = material;
  }

  return material;
}
