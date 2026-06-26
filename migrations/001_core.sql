CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL PRIMARY KEY,
    fullname varchar(255) NOT NULL,
    phone varchar(30) NOT NULL,
    email varchar(255) NOT NULL,
    password_hash varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    created_at timestamp DEFAULT NOW ()
);

CREATE TYPE order_type AS enum (
    'delivered',
    'received',
    'processing',
    'canceled',
    'pending'
);

CREATE TABLE IF NOT EXISTS orders (
    id serial NOT NULL PRIMARY KEY,
    created_at timestamp DEFAULT NOW (),
    amount numeric(10, 2) NOT NULL,
    order_status order_type NOT NULL,
    user_id int NOT NULL REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS categories (
    id serial NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    parent_id int REFERENCES categories (id)
);

CREATE TABLE IF NOT EXISTS products (
    id serial NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    description text NOT NULL,
    price numeric(10, 2) NOT NULL,
    category_id int NOT NULL REFERENCES categories (id),
    image_url text
);

CREATE TABLE IF NOT EXISTS order_items (
    id serial NOT NULL PRIMARY KEY,
    product_id int NOT NULL REFERENCES products (id),
    order_id int NOT NULL REFERENCES orders (id),
    unit_price numeric(10, 2) NOT NULL,
    quantity int NOT NULL
);

CREATE TABLE IF NOT EXISTS product_view (
    id serial NOT NULL PRIMARY KEY,
    product_id int NOT NULL REFERENCES products (id),
    user_id int REFERENCES users (id),
    date timestamp DEFAULT NOW ()
);

CREATE TABLE IF NOT EXISTS price_history (
    id serial NOT NULL PRIMARY KEY,
    product_id int NOT NULL REFERENCES products (id),
    price numeric(10, 2) NOT NULL,
    changed_at timestamp NOT NULL
);