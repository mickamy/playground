CREATE TABLE users
(
    id         BINARY(16)   DEFAULT (UUID_TO_BIN(UUID())) NOT NULL PRIMARY KEY,
    slug       VARCHAR(50)                                NULL,
    created_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6)  NOT NULL,
    updated_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON
        UPDATE CURRENT_TIMESTAMP(6)                       NOT NULL
);

CREATE UNIQUE INDEX idx_users_on_slug ON users (slug);

CREATE TABLE user_accounts
(
    user_id    BINARY(16)                                                               NOT NULL PRIMARY KEY,
    provider   ENUM ('google')                                                          NOT NULL,
    uid        VARCHAR(50)                                                              NOT NULL,
    email      VARCHAR(256)                                                             NOT NULL,
    created_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6)                                NOT NULL,
    updated_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) NOT NULL,
    CONSTRAINT fk_user_accounts_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE UNIQUE INDEX idx_user_accounts_on_email
    ON user_accounts (email);

CREATE UNIQUE INDEX idx_user_accounts_on_provider_and_uid
    ON user_accounts (provider, uid);


CREATE TABLE user_profiles
(
    user_id    BINARY(16)                                                               NOT NULL PRIMARY KEY,
    name       VARCHAR(50)                                                              NOT NULL,
    bio        TEXT                                                                     NOT NULL,
    created_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6)                                NOT NULL,
    updated_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) NOT NULL,
    CONSTRAINT fk_user_profiles_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);
