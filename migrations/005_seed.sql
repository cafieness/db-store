INSERT INTO users (
        id,
        fullname,
        email,
        phone,
        password_hash,
        address
    )
SELECT i,
    'User ' || i,
    'user' || i || '@mail.com',
    '+755512000' || LPAD(i::text, 4, '0'),
    'hashed_password',
    'Address ' || i
FROM generate_series(1, 10) AS i;

SELECT setval('users_id_seq', 10);

INSERT INTO categories (id, name, parent_id)
VALUES (1, 'Electronics', NULL),
    (2, 'People', NULL),
    (3, 'Bananas', NULL),
    (4, 'Phones', 1),
    (5, 'Cavendish', 3);

SELECT setval('categories_id_seq', 5);

INSERT INTO products (
        id,
        name,
        description,
        price,
        category_id,
        image_url
    )
SELECT i,
    'Product ' || i,
    'Description ' || i,
    i * 1000,
    ((i - 1) % 5) + 1,
    'https://example.com/images/product' || i || '.jpg'
FROM generate_series(1, 10) AS i;

SELECT setval('products_id_seq', 10);

INSERT INTO orders (
        id,
        amount,
        order_status,
        user_id
    )
SELECT i,
    i * 10000,
    (
        CASE
            ((i - 1) % 5)
            WHEN 0 THEN 'delivered'
            WHEN 1 THEN 'received'
            WHEN 2 THEN 'processing'
            WHEN 3 THEN 'canceled'
            ELSE 'pending'
        END
    )::order_type,
    ((i - 1) % 10) + 1
FROM generate_series(1, 20) AS i;

SELECT setval('orders_id_seq', 20);

INSERT INTO order_items (
        product_id,
        order_id,
        unit_price,
        quantity
    )
SELECT i,
    ((i * 2 - 1) % 20) + 1,
    i * 1000,
    i
FROM generate_series(1, 10) AS i;

INSERT INTO product_view (product_id, user_id)
SELECT floor(random() * 10 + 1)::int,
    floor(random() * 10 + 1)::int
FROM generate_series(1, 1000);

INSERT INTO price_history (product_id, price, changed_at)
SELECT i,
    i * 1000,
    NOW() - ((i * 2) || ' days')::INTERVAL
FROM generate_series(1, 10) AS i;