CREATE TABLE users
(
    id         BINARY(16)   DEFAULT (BIN_TO_UUID(UUID())) NOT NULL PRIMARY KEY,
    slug       VARCHAR(50)                                NOT NULL,
    created_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6)  NOT NULL,
    updated_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON
        UPDATE CURRENT_TIMESTAMP(6)                       NOT NULL
);

CREATE UNIQUE INDEX index_users_on_slug ON users (slug);
