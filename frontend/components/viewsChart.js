export function renderViewsChart(data) {
    const ctx = document.getElementById("viewsChart");

    new Chart(ctx, {
        type: "bar",
        data: {
            labels: data.map(p => p.name),
            datasets: [{
                label: "Views",
                data: data.map(p => p.total_views),
            }]
        }
    });
}