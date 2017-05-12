# API仕様書

# /answer-data(解答の提出)
解答の提出に使うAPI

~~~json
{
    "QuestionID": "",
    "AnswerData": "",
    "Lang": ""
}
~~~

見本

~~~json
{
    "QuestionID": "1",
    "AnswerData": "puts (1..gets.to_i).to_a.reduce(:+)",
    "Lang": "rb"
}
~~~

|Key|Value|Type|
|:--|:--|:--|
|QuestionID|1~|Integer|
|AnswerData|-|String|
|Lang|rb,py|String|

### QuestionID
問題のIDと関連するIDを指定する問題と異なるIDを指定すると  
採点時に解答が有っていても異なる結果が帰ってくる  

### AnswerData
解答を入れる

### Lang
使用した言語の「拡張子」を入力する。  

現在対応してる言語  

|ProgrammingLang|extension|
|:--|:--|
|Ruby|rb|
|Python|py|

# /create-answer(解答を作る)
***基本的にこのAPIはRailsからの使用になるのでClientが使うことはない***

## POST形式(提出する場合の形式)

~~~json
{
    "Action": "",
    "QuestionID": "",
    "AnswerData": "",
}
~~~

見本

~~~json
{
    "Action": "Create",
    "QuestionID": "1",
    "AnswerData": "1,1\n2,3\n3,6\n4,10\n5,15",
}
~~~

|Key|Value|Type|
|:--|:--|:--|
|Action|Create,Update,Delete|String|
|QuestionID|1~|Integer|
|AnswerData|[,]区切りでinputとoutput|String|

### Action
解答のCreate,Update,Deleteのどれを行うか指定する

### QuestionID
問題のIDと関連するIDを指定する問題と異なるIDを指定すると  
採点時に解答が有っていても異なる結果が帰ってくる  

### AnswerData
[,]区切りで左に入力(case)、右に出力結果（答え）を入力する  
Caseの数に上限はない  
例:

~~~
1,false
0,true
~~~

## Resule形式

~~~json
{
    "Result": ""
}
~~~

見本

~~~json
{
    "Result": "true"
}
~~~

|Key|Value|Type|
|:--|:--|:--|
|Result|true,false|boolean|

### result 
解答ファイルの作成が成功したときはtrue作成できなければfalse
