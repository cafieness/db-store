CREATE VIEW product_view_summary_30ds AS
SELECT p.id,
    p.name,
    COUNT(v.product_id) AS view_count
FROM product_view AS v
    JOIN products AS p ON v.product_id = p.id
WHERE v.date BETWEEN CURRENT_DATE - INTERVAL '30 days'
    AND CURRENT_DATE
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
    JOIN products p ON o.product_id = p.id
WHERE o.order_status = 'delivered'
    AND o.created_at BETWEEN CURRENT_DATE - INTERVAL '30 days'
    AND CURRENT_DATE
GROUP BY p.id,
    p.name
ORDER BY total_sold DESC
LIMIT 10;

CREATE MATERIALIZED VIEW revenue_snapshots_30ds AS
SELECT (sum(oi.unit_price * oi.quantity)) AS total_revenue,
    sum(oi.quantity) AS total_products_sold,
    sum(o.amount) AS total_orders
FROM order_items oi
    JOIN orders o ON o.id = oi.order_id
WHERE o.order_status = 'delivered'
    AND o.created_at BETWEEN CURRENT_DATE - INTERVAL '30 days'
    AND CURRENT_DATE;