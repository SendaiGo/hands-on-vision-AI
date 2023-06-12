# 02.DetectFaces

## Run

```bash
$ go run 02.DetectFaces/main.go xxx.jpg
```

## Result

```bash
[bounding_poly:{vertices:{x:493 y:64} vertices:{x:751 y:64} vertices:{x:751 y:363} vertices:{x:493 y:363}} fd_bounding_poly:{vertices:{x:513 y:133} vertices:{x:724 y:133} vertices:{x:724 y:345} vertices:{x:513 y:345}} landmarks:{type:LEFT_EYE position:{x:549.52026 y:212.12364 z:-3.0994415e-05}} landmarks:{type:RIGHT_EYE position:{x:615.4475 y:184.65811 z:-31.79933}} landmarks:{type:LEFT_OF_LEFT_EYEBROW position:{x:527.77136 y:197.8679 z:24.209892}} landmarks:{type:RIGHT_OF_LEFT_EYEBROW position:{x:555.8008 y:175.86327 z:-13.240559}} landmarks:{type:LEFT_OF_RIGHT_EYEBROW position:{x:585.65594 y:165.95679 z:-28.863201}} landmarks:{type:RIGHT_OF_RIGHT_EYEBROW position:{x:638.7073 y:156.81773 z:-30.58471}} landmarks:{type:MIDPOINT_BETWEEN_EYES position:{x:574.0105 y:189.96881 z:-28.240627}} landmarks:{type:NOSE_TIP position:{x:571.21295 y:229.3549 z:-61.743618}} landmarks:{type:UPPER_LIP position:{x:587.5421 y:261.27563 z:-56.251694}} landmarks:{type:LOWER_LIP position:{x:595.48566 y:282.74283 z:-58.63859}} landmarks:{type:MOUTH_LEFT position:{x:572.4152 y:282.47763 z:-27.901333}} landmarks:{type:MOUTH_RIGHT position:{x:624.3226 y:265.46658 z:-54.703693}} landmarks:{type:MOUTH_CENTER position:{x:593.15106 y:273.073 z:-54.94792}} landmarks:{type:NOSE_BOTTOM_RIGHT position:{x:605.2475 y:234.26892 z:-48.252926}} landmarks:{type:NOSE_BOTTOM_LEFT position:{x:568.808 y:246.15688 z:-29.794645}} landmarks:{type:NOSE_BOTTOM_CENTER position:{x:583.3907 y:245.1505 z:-51.22753}} landmarks:{type:LEFT_EYE_TOP_BOUNDARY position:{x:546.2267 y:205.50565 z:-1.6250787}} landmarks:{type:LEFT_EYE_RIGHT_CORNER position:{x:563.82965 y:206.70546 z:-6.9059787}} landmarks:{type:LEFT_EYE_BOTTOM_BOUNDARY position:{x:550.2179 y:219.52277 z:-2.4921508}} landmarks:{type:LEFT_EYE_LEFT_CORNER position:{x:537.2231 y:217.619 z:13.544509}} landmarks:{type:RIGHT_EYE_TOP_BOUNDARY position:{x:613.6303 y:177.52704 z:-34.03508}} landmarks:{type:RIGHT_EYE_RIGHT_CORNER position:{x:633.8228 y:181.10614 z:-32.807674}} landmarks:{type:RIGHT_EYE_BOTTOM_BOUNDARY position:{x:617.3309 y:191.18805 z:-35.104168}} landmarks:{type:RIGHT_EYE_LEFT_CORNER position:{x:600.2542 y:189.16919 z:-25.04515}} landmarks:{type:LEFT_EYEBROW_UPPER_MIDPOINT position:{x:538.32477 y:181.20029 z:2.756199}} landmarks:{type:RIGHT_EYEBROW_UPPER_MIDPOINT position:{x:609.68054 y:154.26794 z:-32.10845}} landmarks:{type:LEFT_EAR_TRAGION position:{x:553.97345 y:253.37596 z:100.08048}} landmarks:{type:RIGHT_EAR_TRAGION position:{x:714.0595 y:225.42667 z:19.677036}} landmarks:{type:FOREHEAD_GLABELLA position:{x:569.61224 y:171.95905 z:-23.73064}} landmarks:{type:CHIN_GNATHION position:{x:607.37604 y:324.48624 z:-61.622322}} landmarks:{type:CHIN_LEFT_GONION position:{x:563.10736 y:320.58163 z:40.61389}} landmarks:{type:CHIN_RIGHT_GONION position:{x:698.907 y:274.9891 z:-19.0029}} landmarks:{type:LEFT_CHEEK_CENTER position:{x:549.84436 y:264.77277 z:-2.7229156}} landmarks:{type:RIGHT_CHEEK_CENTER position:{x:642.69617 y:230.4593 z:-47.845352}} roll_angle:-12.220201 pan_angle:-25.237288 tilt_angle:16.877129 detection_confidence:0.96875 landmarking_confidence:0.43151987 joy_likelihood:VERY_UNLIKELY sorrow_likelihood:VERY_UNLIKELY anger_likelihood:VERY_UNLIKELY surprise_likelihood:VERY_UNLIKELY under_exposed_likelihood:VERY_UNLIKELY blurred_likelihood:VERY_UNLIKELY headwear_likelihood:VERY_UNLIKELY bounding_poly:{vertices:{x:124 y:123} vertices:{x:373 y:123} vertices:{x:373 y:413} vertices:{x:124 y:413}} fd_bounding_poly:{vertices:{x:147 y:194} vertices:{x:346 y:194} vertices:{x:346 y:394} vertices:{x:147 y:394}} landmarks:{type:LEFT_EYE position:{x:205.96619 y:243.74507 z:0.001707077}} landmarks:{type:RIGHT_EYE position:{x:271.38803 y:253.67548 z:-17.02018}} landmarks:{type:LEFT_OF_LEFT_EYEBROW position:{x:187.85399 y:219.83246 z:17.942162}} landmarks:{type:RIGHT_OF_LEFT_EYEBROW position:{x:222.633 y:221.1972 z:-11.707433}} landmarks:{type:LEFT_OF_RIGHT_EYEBROW position:{x:259.28052 y:226.08652 z:-20.155453}} landmarks:{type:RIGHT_OF_RIGHT_EYEBROW position:{x:307.46234 y:237.87735 z:-11.46135}} landmarks:{type:MIDPOINT_BETWEEN_EYES position:{x:235.1293 y:243.03249 z:-22.174685}} landmarks:{type:NOSE_TIP position:{x:223.10556 y:274.7159 z:-52.229553}} landmarks:{type:UPPER_LIP position:{x:220.97012 y:305.59232 z:-43.985443}} landmarks:{type:LOWER_LIP position:{x:219.49635 y:328.38168 z:-45.052876}} landmarks:{type:MOUTH_LEFT position:{x:193.02719 y:315.52478 z:-20.868435}} landmarks:{type:MOUTH_RIGHT position:{x:254.13298 y:325.59985 z:-35.46117}} landmarks:{type:MOUTH_CENTER position:{x:219.7203 y:317.60834 z:-41.78154}} landmarks:{type:NOSE_BOTTOM_RIGHT position:{x:245.07803 y:291.9621 z:-33.725464}} landmarks:{type:NOSE_BOTTOM_LEFT position:{x:209.5132 y:282.8627 z:-23.157837}} landmarks:{type:NOSE_BOTTOM_CENTER position:{x:224.98653 y:290.97873 z:-40.344765}} landmarks:{type:LEFT_EYE_TOP_BOUNDARY position:{x:205.76907 y:236.71962 z:-2.1858006}} landmarks:{type:LEFT_EYE_RIGHT_CORNER position:{x:218.97418 y:245.07816 z:-3.5706005}} landmarks:{type:LEFT_EYE_BOTTOM_BOUNDARY position:{x:204.45676 y:248.63248 z:-1.8048601}} landmarks:{type:LEFT_EYE_LEFT_CORNER position:{x:190.9444 y:242.83325 z:10.112448}} landmarks:{type:RIGHT_EYE_TOP_BOUNDARY position:{x:272.7985 y:246.79872 z:-19.627985}} landmarks:{type:RIGHT_EYE_RIGHT_CORNER position:{x:288.20068 y:257.73187 z:-14.572316}} landmarks:{type:RIGHT_EYE_BOTTOM_BOUNDARY position:{x:270.4733 y:259.13004 z:-19.4219}} landmarks:{type:RIGHT_EYE_LEFT_CORNER position:{x:257.33392 y:251.28273 z:-13.646151}} landmarks:{type:LEFT_EYEBROW_UPPER_MIDPOINT position:{x:204.1974 y:213.44092 z:0.50525093}} landmarks:{type:RIGHT_EYEBROW_UPPER_MIDPOINT position:{x:284.3752 y:224.61887 z:-18.450003}} landmarks:{type:LEFT_EAR_TRAGION position:{x:166.68912 y:268.76965 z:92.56256}} landmarks:{type:RIGHT_EAR_TRAGION position:{x:340.15695 y:321.55084 z:46.749336}} landmarks:{type:FOREHEAD_GLABELLA position:{x:239.13635 y:224.17644 z:-18.60495}} landmarks:{type:CHIN_GNATHION position:{x:210.66524 y:364.67874 z:-44.344406}} landmarks:{type:CHIN_LEFT_GONION position:{x:164.07071 y:325.13388 z:42.761024}} landmarks:{type:CHIN_RIGHT_GONION position:{x:296.01172 y:367.85547 z:8.923088}} landmarks:{type:LEFT_CHEEK_CENTER position:{x:182.5281 y:286.99353 z:-1.8351536}} landmarks:{type:RIGHT_CHEEK_CENTER position:{x:278.66486 y:301.93497 z:-25.960789}} roll_angle:12.567263 pan_angle:-14.094562 tilt_angle:13.821742 detection_confidence:0.9765625 landmarking_confidence:0.7429294 joy_likelihood:UNLIKELY sorrow_likelihood:VERY_UNLIKELY anger_likelihood:VERY_UNLIKELY surprise_likelihood:VERY_UNLIKELY under_exposed_likelihood:VERY_UNLIKELY blurred_likelihood:VERY_UNLIKELY headwear_likelihood:VERY_UNLIKELY]
```


## 顔の注釈

| プロパティ | 説明 |
| ----- | ----- |
| bounding_poly | 顔の境界ボックス。 |
| fd_bounding_poly | 顔の境界ボックス。 |
| landmarks | 顔のランドマーク。 |
| roll_angle | 顔の回転角度。 |
| pan_angle | 顔の回転角度。 |
| tilt_angle | 顔の回転角度。 |
| detection_confidence | 顔の検出信頼度。 |
| landmarking_confidence | 顔のランドマーク信頼度。 |
| joy_likelihood | 顔の喜び度。 |
| sorrow_likelihood | 顔の悲しみ度。 |
| anger_likelihood | 顔の怒り度。 |
| surprise_likelihood | 顔の驚き度。 |
| under_exposed_likelihood | 顔の露出度。 |
| blurred_likelihood | 顔のぼやけ度。 |
| headwear_likelihood | 顔の帽子度。 |

## 可能性

| タイプ | 説明 |
| ----- | ----- |
| UNKNOWN | 可能性は不明 |
| VERY_UNLIKELY | その可能性は非常に低いです。 |
| UNLIKELY | それはありそうにありません。 |
| POSSIBLE | 可能です。 |
| LIKELY | それは可能性があります。 |
| VERY_LIKELY | その可能性は非常に高いです。 |

## 応用 
顔の抽出

```bash
$ go run 02.DetectFaces/main.go xxx.jpg
```

