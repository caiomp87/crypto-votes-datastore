CREATE TABLE IF NOT EXISTS cryptos(
    id SERIAL PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    network VARCHAR (50) NOT NULL,
    upVotes INTEGER DEFAULT 0,
    downVotes INTEGER DEFAULT 0,
    createdat TIMESTAMP DEFAULT NOW(),
    updatedat TIMESTAMP DEFAULT NOW()
)
