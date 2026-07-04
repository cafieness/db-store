export function renderTopProductsChart(data) {
    const ctx = document.getElementById("topProductsChart");

    new Chart(ctx, {
        type: "bar",
        data: {
            labels: data.map(p => p.name),
            datasets: [{
                label: "Units Sold",
                data: data.map(p => p.total_sold),
            }]
        }
    });
}