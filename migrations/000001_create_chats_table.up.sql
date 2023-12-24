CREATE TABLE IF NOT EXISTS chats
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR NOT NULL DEFAULT '',
    type       SMALLINT NOT NULL DEFAULT 1,
    image_url  VARCHAR NOT NULL DEFAULT '',
    created_by BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);