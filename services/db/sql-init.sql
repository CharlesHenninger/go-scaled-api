CREATE TABLE IF NOT EXISTS domains (
    id SERIAL PRIMARY KEY,
    name text,
    events int,
    bounced boolean,
    UNIQUE(name)
);

INSERT INTO "domains"("name", "events", "bounced")
VALUES
    ('example1.com', 1222, false),
    ('example2.com', 300, false),
    ('example3.com', 1500, true),
    ('example4.com', 400, true)
;

