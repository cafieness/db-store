import {
    getRevenue,
    getTopProducts,
    getProductViews,
    getOrdersSummary
} from "./api/api.js";

import { renderCards } from "./components/cards.js";
import { renderTopProductsChart } from "./components/revenueChart.js";
import { renderViewsChart } from "./components/viewsChart.js";

async function init() {

    const [revenue30d, ordersAll] = await Promise.all([
        getRevenue(),
        getOrdersSummary()
    ]);

    renderCards(ordersAll, revenue30d);

    const topProducts = await getTopProducts();
    renderTopProductsChart(topProducts);

    const views = await getProductViews();
    renderViewsChart(views);
}

init();