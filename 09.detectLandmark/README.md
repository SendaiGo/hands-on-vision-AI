# 09.DetectLandmarks

## Run

```bash
$ go run 10.DetectLandmarks/main.go xxx.jpg
```

## Result

```bash
[mid:"/m/0hm_7"  description:"Red Square"  score:0.8557956  bounding_poly:{vertices:{}  vertices:{x:503}  vertices:{x:503  y:650}  vertices:{y:650}}  locations:{lat_lng:{latitude:55.75393029999999  longitude:37.620794999999994}}]
```

```
https://www.google.com/maps/@55.7556213,37.620526,16z?hl=ja&entry=ttu
```


##  ランドマークの検出

| プロパティ | 説明 |
| ----- | ----- |
| mid | ランドマークの識別子。 |
| description | ランドマークの説明。 |
| score | ランドマークの信頼度。 |
| bounding_poly | ランドマークの境界ボックス。 |
| locations | ランドマークの位置情報。 |
