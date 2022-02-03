## デプロイメントフロー

このリポジトリのCI/CDフローは、github-flowをベースにしています。

ここではgithub-flowについて簡単に紹介します。

1. mainをベースとしたfeatureブランチを作成する。
2. featureブランチからmainブランチへのプルリクエストを作成する。
3. mainにマージされたら、すぐにProductionにデプロイする。

mainブランチはいつでもデプロイできる状態にしておくのが理想的です。

このリポジトリは、リリースの公開がトリガーとなります。

1. 機能へのプッシュ時に <image:hash> を作成します。
2. mainにpushしたときに<image:hash>を作成して<image:latest>を作成します。
3. mainにpushしたときにリリースが作成または更新されます。
4. 3.の後、本番環境に<image:latest>をデプロイする。
5. 5.リリース公開時に<image:release-version>を作成し、<image:latest>を作成する。
6. 6. <image:release-version>を、5.以降に本番環境に導入する。

## リリースノート

リリースノートには、checkedとnot yetの項目があります。
対象のリポジトリにcheckedラベルがあればcheckedセクションに表示され、なければnot yetセクションに表示されます。

## テストについて

テストや静的解析などは、機能ブランチにプッシュしたときだけ行われます。
以前は、ブランチが壊れていないことを確認するために、テストはマージの後に行われていました。
しかし、Protected Branch で "Require branches to be up to date before merging" を有効にすると、機能ブランチをマージする際にベースである main に従うようにすることができるようになります。

## 問題

- 誰でもリリースを作成することができます。
- 誰でもいつでも公開できます。

www.DeepL.com/Translator（無料版）で翻訳しました。