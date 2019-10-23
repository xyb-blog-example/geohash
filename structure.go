package geohash

import (
   "sync"
)

type GeoMap struct {
   gm *geoMap
}

type geoMap struct {
   lock     sync.RWMutex
   maxX     float64   //x轴最大值，单位：毫米
   maxY     float64   //y轴最大值，单位：毫米
   minX     float64   //x轴最小值，单位：毫米
   minY     float64   //y轴最小值，单位：毫米
   cutCount int32   //切割次数
   hashMap  map[string][]string   //map[哈希值][]nodeID
}

type Node struct {
    X   float64
    Y   float64
}

var base32Map = []string{
    "0","1","2","3","4","5","6","7","8","9","b","c","d","e","f","g","h","j","k","m","n","p","q","r","s","t","u","v","w","x","y","z",
}

var base32DecodeMap = map[string]uint8{
    "0":0,"1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"b":10,"c":11,"d":12,"e":13,"f":14,"g":15,"h":16,
    "j":17,"k":18,"m":19,"n":20,"p":21,"q":22,"r":23,"s":24,"t":25,"u":26,"v":27,"w":28,"x":29,"y":30,"z":31,
}