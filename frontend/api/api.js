const BASE_URL = "http://localhost:8080";

export async function getRevenue() {
    const res = await fetch(`${BASE_URL}/analytics/revenue`);
    return res.json();
}

export async function getTopProducts() {
    const res = await fetch(`${BASE_URL}/analytics/top-products`);
    return res.json();
}

export async function getProductViews() {
    const res = await fetch(`${BASE_URL}/analytics/productview`);
    return res.json();
}

export async function getOrdersSummary() {
    const res = await fetch("http://localhost:8080/analytics/orders-summary");
    return res.json();
}