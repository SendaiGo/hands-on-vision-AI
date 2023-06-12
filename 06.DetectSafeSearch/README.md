# 06. Detect Safe Search

## Run

```bash
$ go run 06.detectSafeSearch/main.go xxx.jpg
```


## Result

```bash
Adult: VERY_UNLIKELY
Medical: VERY_UNLIKELY
Racy: VERY_UNLIKELY
Spoofed: VERY_UNLIKELY
Violence: VERY_UNLIKELY
```

## セーフサーチ注釈

| タイプ | 説明 |
| ----- | ----- |
| adult | 画像のアダルト コンテンツの可能性を表します。アダルト コンテンツには、ヌード、ポルノ画像や漫画、性的行為などの要素が含まれる場合があります。 |
| spoof | 画像のなりすましの可能性を表します。画像の正規バージョンが面白く、不快に見えるように変更が加えられた可能性があります。 |
| medical | これは医療画像である可能性があります。 |
| violence | この画像には暴力的なコンテンツが含まれている可能性があります。 |
| racy | リクエスト画像に際どい内容が含まれている可能性。きわどいコンテンツには、露出度の高いまたは薄手の衣服、戦略的に覆われたヌード、卑猥または挑発的なポーズ、敏感な身体部分のクローズアップなどが含まれる場合がありますが、これらに限定されません。 |

## 可能性

| タイプ | 説明 |
| ----- | ----- |
| UNKNOWN | 可能性は不明 |
| VERY_UNLIKELY | その可能性は非常に低いです。 |
| UNLIKELY | それはありそうにありません。 |
| POSSIBLE | 可能です。 |
| LIKELY | それは可能性があります。 |
| VERY_LIKELY | その可能性は非常に高いです。 |
