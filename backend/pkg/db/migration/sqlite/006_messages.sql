-- +migrate Up
CREATE TABLE IF NOT EXISTS messages (
    "message_id" VARCHAR(255) not null,
    "sender_id" VARCHAR(255) not null,
    "receiver_id" DATETIME not null default CURRENT_TIMESTAMP,
    "type" VARCHAR(255) not null,
    "created_at" VARCHAR(255) not null,
    "content" VARCHAR(255) not null,
    primary key ("message_id")
);

-- +migrate Down
DROP TABLE messages;