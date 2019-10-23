package geohash

func CreateGeoMap(maxX, maxY, minX, minY float64, cutCount int32, nodeList map[string]Node) (gm *GeoMap) {
    gm = &GeoMap{
        gm: &geoMap{
            minX:       minX,
            minY:       minY,
            maxX:       maxX,
            maxY:       maxY,
            cutCount:   cutCount,
            hashMap:    make(map[string][]string),
        },
    }

    //1 遍历所有点列表，生成geoHash
    for nodeID, node := range nodeList {
        hashKey := gm.gm.createHashKey(node)
        gm.gm.hashMap[hashKey] = append(gm.gm.hashMap[hashKey], nodeID)
    }
    return gm
}

func(gm *geoMap) createHashKey(node Node) string {
    var xResult, yResult []uint8
    //1 计算X轴的哈希Key
    minX, maxX  := float64(gm.minX), float64(gm.maxX)
    for i := int32(0); i < gm.cutCount; i++ {
        result, newMinX, newMaxX  := checkBelongTo(minX, maxX, node.X)
        xResult = append(xResult, result)
        minX, maxX = newMinX, newMaxX
    }

    //2 计算Y轴的哈希Key
    minY, maxY  := float64(gm.minY), float64(gm.maxY)
    for i := int32(0); i < gm.cutCount; i++ {
        result, newMinY, newMaxY  := checkBelongTo(minY, maxY, node.Y)
        yResult = append(yResult, result)
        minY, maxY = newMinY, newMaxY
    }

    //3 合并XY轴的哈希值
    var hashKeyCharList []uint8
    for i := int32(0); i < gm.cutCount; i++ {
        hashKeyCharList = append(hashKeyCharList, xResult[i])
        hashKeyCharList = append(hashKeyCharList, yResult[i])
    }

    //4 计算哈希值
    var hashKey string
    var sum int
    for i, j := 0, 0; i < len(hashKeyCharList); i++ {
        value := hashKeyCharList[i]
        if value == 1 {
            sum += 1 << uint(4 - j)
        }
        if j == 4 || i == len(hashKeyCharList) - 1{
            value := base32Map[sum]
            hashKey += value
            j = 0
            sum = 0
            continue
        }
        j++
    }
    return hashKey
}

func checkBelongTo(min, max, target float64) (result uint8, nextMin, nextMax float64) {
    half := float64(min + max) / 2
    if float64(target) >= half {
        return 1, half, max
    } else {
        return 0, min, half
    }
}