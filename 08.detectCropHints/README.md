# 08. Detect Crop Hints 

## Run

```bash
$ go run 08.detectCropHints/main.go xxx.jpg
```

## Result

```bash
crop_hints:{bounding_poly:{vertices:{x:75} vertices:{x:439} vertices:{x:439 y:649} vertices:{x:75 y:649}} confidence:0.6153631 importance_fraction:1}
```

## Crop Hints

| プロパティ | 説明 |
| ----- | ----- |
| bounding_poly | 画像の切り取りヒントの境界ボックス。 |
| confidence | 画像の切り取りヒントの信頼度。 |
| importance_fraction | 画像の切り取りヒントの重要性の割合。 |
