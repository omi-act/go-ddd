-- ユーザーテーブルの作成
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    age INTEGER CHECK (age >= 0 AND age <= 150),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- インデックスの作成
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- テストデータの挿入
INSERT INTO users (name, email, age) VALUES 
    ('田中太郎', 'tanaka@example.com', 32),
    ('佐藤花子', 'sato@example.com', 28),
    ('鈴木一郎', 'suzuki@example.com', 45),
    ('高橋美咲', 'takahashi@example.com', 26),
    ('伊藤健太', 'ito@example.com', 35),
    ('山田由美', 'yamada@example.com', 29),
    ('中村智也', 'nakamura@example.com', 41),
    ('小林あい', 'kobayashi@example.com', 24),
    ('加藤隆', 'kato@example.com', 38),
    ('清水麻衣', 'shimizu@example.com', 33),
    ('林大輔', 'hayashi@example.com', 27),
    ('松本真理', 'matsumoto@example.com', 36),
    ('井上直樹', 'inoue@example.com', 42),
    ('木村恵子', 'kimura@example.com', 31),
    ('斎藤翔太', 'saito@example.com', 25);
