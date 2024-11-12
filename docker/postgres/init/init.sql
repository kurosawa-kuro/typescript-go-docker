-- データベースの文字コード設定
ALTER DATABASE ${DATABASE} SET client_encoding TO 'UTF8';

-- タイムゾーンの設定
ALTER DATABASE ${DATABASE} SET timezone TO 'Asia/Tokyo'; 