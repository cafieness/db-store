CREATE INDEX idx_orders_user_id ON orders(user_id);

CREATE INDEX idx_categories_parent_id ON categories(parent_id);

CREATE INDEX idx_products_category_id ON products(category_id);

CREATE INDEX idx_order_items_product_id ON order_items(product_id);

CREATE INDEX idx_order_items_order_id ON order_items(order_id);

CREATE INDEX idx_product_product_id ON product_view(product_id);

CREATE INDEX idx_product_user_id ON product_view(user_id);

CREATE INDEX idx_price_history_product_id ON price_history(product_id);

CREATE INDEX idx_orders_created_at ON orders(created_at);

CREATE INDEX idx_orders_status ON orders(order_status);