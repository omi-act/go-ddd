-- テスト用テーブルの作成
CREATE TABLE IF NOT EXISTS test_table (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- テストデータ10件の挿入
INSERT INTO test_table (name) VALUES 
    ('test_name_1'),
    ('test_name_2'),
    ('test_name_3'),
    ('test_name_4'),
    ('test_name_5'),
    ('test_name_6'),
    ('test_name_7'),
    ('test_name_8'),
    ('test_name_9'),
    ('test_name_10');
