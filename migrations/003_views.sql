CREATE VIEW product_view_summary_30ds AS
SELECT p.id,
    p.name,
    COUNT(v.product_id) AS view_count
FROM product_view v
    JOIN products p ON p.id = v.product_id
WHERE v.date >= NOW() - INTERVAL '30 days'
GROUP BY p.id,
    p.name;

CREATE VIEW category_tree AS WITH recursive tree AS (
    SELECT c2.*
    FROM categories c2
    WHERE c2.parent_id IS NULL
    UNION ALL
    SELECT c.*
    FROM categories c
        JOIN tree t ON t.id = c.parent_id
)
SELECT *
FROM tree;

CREATE MATERIALIZED VIEW top_products_30ds AS
SELECT p.id,
    p.name,
    sum(oi.quantity) AS total_sold
FROM order_items oi
    JOIN orders o ON o.id = oi.order_id
    JOIN products p ON oi.product_id = p.id
WHERE o.order_status = 'delivered'
    AND o.created_at >= NOW() - INTERVAL '30 days'
GROUP BY p.id,
    p.name;

CREATE MATERIALIZED VIEW revenue_snapshots_30ds AS
SELECT COALESCE(
        SUM(
            COALESCE(oi.unit_price, 0) * COALESCE(oi.quantity, 0)
        ),
        0
    ) AS total_revenue,
    COALESCE(SUM(COALESCE(oi.quantity, 0)), 0) AS total_products_sold,
    COUNT(DISTINCT o.id) AS total_orders
FROM order_items oi
    JOIN orders o ON o.id = oi.order_id
WHERE o.order_status = 'delivered'
    AND o.created_at >= NOW() - INTERVAL '30 days';