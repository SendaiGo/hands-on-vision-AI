# 07.Detect Web Entities 

## Run

```bash
$ go run 07.DetectWebEntities/main.go xxx.jpg
```

## Result

```bash
url:"https://unsplash.com/photos/1GvSPe52xx4" page_title:"Multicolored dome castle photo – Free Building Image on Unsplash" full_matching_images:{url:"https://images.unsplash.com/photo-1513326738677-b964603b136d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxjb2xsZWN0aW9uLXRodW1ibmFpbHx8ODY4ODY2N3x8ZW58MHx8fHx8&auto=format&fit=crop&w=420&q=60"}
url:"https://unsplash.com/photos/nav82EFdekY" page_title:"A <b>red</b> traffic light on a city street photo - Unsplash" full_matching_images:{url:"https://images.unsplash.com/photo-1513326738677-b964603b136d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxjb2xsZWN0aW9uLXRodW1ibmFpbHx8d0xjNEJJMUlTc018fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=420&q=60"}
```

## ウェブエンティティの検出

| プロパティ | 説明 |
| ----- | ----- |
| entity_id | ウェブエンティティの識別子。 |
| score | ウェブエンティティの信頼度。 |
| description | ウェブエンティティの説明。 |
| topicality | ウェブエンティティのトピカリティ。 |
| bounding_poly | ウェブエンティティの境界ボックス。 |
| page_title | ウェブエンティティのページタイトル。 |
| partial_matching_images | ウェブエンティティの部分一致画像。 |
| full_matching_images | ウェブエンティティの完全一致画像。 |
| pages_with_matching_images | ウェブエンティティの一致画像を含むページ。 |

