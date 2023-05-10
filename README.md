## vertixAIハンズオン with Go

### 事前準備

### Goをインストールする
- https://golang.org/doc/install

### Gitをインストールする
- https://git-scm.com/downloads

### Githubからソースコードをダウンロードする

```
git clone https://github.com/SendaiGo/hands-on-vision-AI
```

### 必要なライブラリをインストールする

```
cd hands-on-vision-AI
go mod tidy
```

### サンプルコードを実行する

```
go run xxx/main.go xxx.jpg
```
## サンプルコードの説明

## step1 detectImageProperties
- 画像の色を調べる

## step2 detectFaces
- 画像の中の顔を調べる

## step3 detectLabels
- 画像の中の物体を調べる

## step4 detectText
- 画像の中の文字を調べる

## step5 productSearch
- 画像の中の商品を調べる

## step6 objectLocalization
- 画像の中の物体の位置を調べる

## step7 detectWebEntities
- 画像の中のWeb上の物体を調べる

## step8 detectCropHints
- 画像の中の切り取り方を調べる

## step9 detectSafeSearch
- 画像の中の安全性を調べる

## step10 detectWebEntities
- 画像の中のWeb上の物体を調べる


### Vertix AI の モデルを作成する
時間があったら

```
Vertaによるモデル管理
クライアントライブラリのインストール: go get github.com/VertaAI/verta-go
プロジェクト・データセット・モデルの作成: client.CreateProject(), client.CreateDataset(), client.CreateModel()
ノートブックの実行と結果の保存
Jupyterノートブック上から、vertaパッケージをインストール
ノートブックから、Run()メソッドを呼び出して処理を実行： runID, err := client.CreateRun(ctx, &CreateRunRequest{ExperimentID: experiment.ID})
モデルのトレーニングと評価
データセットの取得: dataset, err := client.FindDatasetByName(datasetName)
モデルにデータセットを割り当て: model.LogDataset(dataset)
モデルのトレーニングと評価: model.Fit(data)、model.Evaluate(testData)
モデルのデプロイ
モデルの保存: model.Save()
モデルのデプロイ: endpt, err := model.Deploy()
```
