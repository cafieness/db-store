export function renderCards(all, last30) {
    const container = document.getElementById("cards");

    const a = all[0];
    const b = last30[0];

    container.innerHTML = `
        <div class="card">
            <h3>All Orders</h3>
            <p>${a.total_orders}</p>
        </div>

        <div class="card">
            <h3>All Revenue</h3>
            <p>$${a.total_revenue}</p>
        </div>

        <div class="card">
            <h3>30-d Orders</h3>
            <p>${b.total_orders}</p>
        </div>

        <div class="card">
            <h3>30-d Revenue</h3>
            <p>$${b.total_revenue}</p>
        </div>
    `;
}