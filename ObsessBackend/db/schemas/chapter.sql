チャプターテーブル（chapters）
カラム名	型	説明
id	INT (AUTO_INCREMENT)	主キー
novel_id	INT	小説ID（外部キー）
title	VARCHAR(255)	章タイトル
content	TEXT	章の内容
chapter_number	INT	章番号
created_at	DATETIME	投稿日時
updated_at	DATETIME	更新日時