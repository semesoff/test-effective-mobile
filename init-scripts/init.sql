DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NUll CHECK (name <> ''),
    surname    VARCHAR(255) NOT NULL CHECK (surname <> ''),
    patronymic VARCHAR(255),
    age        INT          NOT NULL CHECK ( age >= 0 AND age <= 130 ),
    gender     VARCHAR(10)  NOT NULL CHECK (gender IN ('male', 'female')),
    nation     VARCHAR(30)  NOT NULL
);