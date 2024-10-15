### DB の各要素を COLATE にすべきか

大文字小文字を判断することはないので COLATE でいいと考えている。

### DB 選択理由

DB 設計時に、MySQL 以外の DB にするか悩んだが、MySQL にすることにした。
なぜなら、Postgres は ACID 原則に準拠するものの、遅くなる傾向があるから(https://aws.amazon.com/jp/compare/the-difference-between-mysql-vs-postgresql/)
MariaDB の方も速いらしいが、MySQL8.0 で MySQL の方が速くなったり、そもそも JSON が格納できない(https://aws.amazon.com/jp/compare/the-difference-between-mariadb-vs-mysql/)AWSのAuroraにも対応していない。
他の DB も一応特徴を見たが、アプリ設計において MySQL を上回る明確な長所が見当たらない。
NoSQL は将来 JOIN などでデータ分析するため、使用しない。

### 小説タグ設定

タグはテーブルを分けることにする。一つの小説につき、タグを十個以上つける予定なので、一つのカラムには収まらない。
