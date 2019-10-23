package geohash

import (
    "git-pd.megvii-inc.com/srgbase/gin"
)

func(gm *GeoMap) QueryNeighborNodeList(c *gin.Context, warehouseID, mapID string, x, y, radius float64) (neighborList []string) {
    nowHashKey      := gm.gm.createHashKey(c, Node{X:x,Y:y})
    neighborList    = append(neighborList, gm.gm.hashMap[nowHashKey]...)

    if y + radius < gm.gm.maxY {
        //1 北方邻居
        northHashKey    := gm.gm.createHashKey(c, Node{X:x,Y:y + radius})
        if nowHashKey != northHashKey {
            neighborList    = append(neighborList, gm.gm.hashMap[northHashKey]...)
        }
        //2 东北方邻居
        if x + radius < gm.gm.maxX {
            eastHashKey     := gm.gm.createHashKey(c, Node{X:x + radius,Y:y + radius})
            if nowHashKey != eastHashKey {
                neighborList    = append(neighborList, gm.gm.hashMap[eastHashKey]...)
            }
        }
        //3 西北方邻居
        if x - radius > gm.gm.minX {
            westHashKey     := gm.gm.createHashKey(c, Node{X:x - radius,Y:y + radius})
            if nowHashKey != westHashKey {
                neighborList    = append(neighborList, gm.gm.hashMap[westHashKey]...)
            }
        }

        if y + radius * 2 < gm.gm.maxY {
            northHashKey    = gm.gm.createHashKey(c, Node{X:x,Y:y + radius * 2})
            if nowHashKey != northHashKey {
                neighborList    = append(neighborList, gm.gm.hashMap[northHashKey]...)
            }
        }
    }

    //4 东方邻居
    if x + radius < gm.gm.maxX {
        eastHashKey     := gm.gm.createHashKey(c, Node{X:x + radius,Y:y})
        if nowHashKey != eastHashKey {
            neighborList    = append(neighborList, gm.gm.hashMap[eastHashKey]...)
        }
        if x + radius * 2 < gm.gm.maxX {
            eastHashKey     = gm.gm.createHashKey(c, Node{X:x + radius * 2,Y:y})
            if nowHashKey != eastHashKey {
                neighborList    = append(neighborList, gm.gm.hashMap[eastHashKey]...)
            }
        }
    }

    if y - radius > gm.gm.minY {
        //5 南方邻居
        southHashKey     := gm.gm.createHashKey(c, Node{X:x,Y:y - radius})
        if nowHashKey != southHashKey {
            neighborList    = append(neighborList, gm.gm.hashMap[southHashKey]...)
        }
        if y - radius * 2 > gm.gm.minY {
            southHashKey     = gm.gm.createHashKey(c, Node{X:x,Y:y - radius * 2})
            if nowHashKey != southHashKey {
                neighborList    = append(neighborList, gm.gm.hashMap[southHashKey]...)
            }
        }
        //6 东南方邻居
        if x + radius < gm.gm.maxX {
            eastHashKey     := gm.gm.createHashKey(c, Node{X:x + radius,Y:y - radius})
            if nowHashKey != eastHashKey {
                neighborList    = append(neighborList, gm.gm.hashMap[eastHashKey]...)
            }
        }
        //7 西南方邻居
        if x - radius > gm.gm.minX {
            westHashKey     := gm.gm.createHashKey(c, Node{X:x - radius,Y:y - radius})
            if nowHashKey != westHashKey {
                neighborList    = append(neighborList, gm.gm.hashMap[westHashKey]...)
            }
        }
    }

    //4 西方邻居
    if x - radius > gm.gm.minX {
        westHashKey     := gm.gm.createHashKey(c, Node{X:x - radius,Y:y})
        if nowHashKey != westHashKey {
            neighborList    = append(neighborList, gm.gm.hashMap[westHashKey]...)
        }
        if x - radius * 2 > gm.gm.minX {
            westHashKey     := gm.gm.createHashKey(c, Node{X:x - radius * 2,Y:y})
            if nowHashKey != westHashKey {
                neighborList    = append(neighborList, gm.gm.hashMap[westHashKey]...)
            }
        }
    }
    return neighborList
}
