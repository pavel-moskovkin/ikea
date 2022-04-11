CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE sex_type AS ENUM (
    'male',
    'female',
    'other'
    );

CREATE TABLE IF NOT EXISTS users
(
    uuid        uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name  VARCHAR(30)  NOT NULL,
    last_name   VARCHAR(30)  NOT NULL,
    middle_name VARCHAR(30),
    full_name   VARCHAR(255) NOT NULL,
    sex         sex_type     NOT NULL,
    birth_date  timestamp
);


CREATE TABLE IF NOT EXISTS orders
(
    uuid         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    number       int       NOT NULL,
    user_id      uuid,
    item_ids     uuid[]    NOT NULL,
    created_at   timestamp NOT NULL,
    completed_at timestamp,
    deleted_at   timestamp,
    UNIQUE (number)
);

ALTER TABLE orders
    ADD CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES users (uuid)
            ON DELETE CASCADE;


CREATE TABLE IF NOT EXISTS user_orders
(
    id       SERIAL UNIQUE,
    user_id  uuid,
    order_id uuid
);

CREATE UNIQUE INDEX ON user_orders (user_id, order_id);

ALTER TABLE user_orders
    ADD CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES users (uuid);

ALTER TABLE user_orders
    ADD CONSTRAINT fk_order_id
        FOREIGN KEY (order_id)
            REFERENCES orders (uuid);


CREATE TABLE IF NOT EXISTS items
(
    uuid          uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name          VARCHAR(30) NOT NULL,
    description   VARCHAR(255),
    price         float8,
    left_in_stock int,

    CONSTRAINT positive_price check (price > 0),
    CONSTRAINT item_exists check (left_in_stock >= 0)
);

ALTER TABLE orders
    ADD CONSTRAINT fk_order_number
        FOREIGN KEY (number)
            REFERENCES user_orders (id)
            ON DELETE CASCADE;
