CREATE TABLE IF NOT EXISTS lib (
    group_name VARCHAR(200),
    song_name VARCHAR(200),
    release_date DATE,
    text TEXT,
    link VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS verses (
    group_name VARCHAR(200),
    song_name VARCHAR(200),
    text TEXT,
);
