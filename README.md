# bdd-apitest-godog
CucumberをベースとしたBDDフレームワーク[godog](https://github.com/cucumber/godog)を用いたAPIのテスト。   

## BDD
BDD（振る舞い駆動開発）はTDD（テスト駆動開発）を拡張した開発手法で、ユーザーが実際に製品・サービスを使用するシナリオに基づいてテストを実装し、それらのテストを満たすような機能を実装していくという開発手法。  

##  How to test
ルートディレクトリにて以下のコマンドを実行することで、全てのフィーチャに対するAPI単位の挙動テストが可能。  
`godog run`
  
個別のフィーチャに対してテストを行う場合は、以下のようにファイルを指定してコマンドを実行。  
`godog run feature/hoge.feature`
