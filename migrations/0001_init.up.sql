CREATE TABLE IF NOT EXISTS teams (
    name TEXT PRIMARY KEY,
    strength REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    week INT NOT NULL,
    home_team TEXT REFERENCES teams(name),
    away_team TEXT REFERENCES teams(name),
    home_goals INT,
    away_goals INT,
    played BOOLEAN NOT NULL DEFAULT false
);
