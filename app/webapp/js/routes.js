import loadHomeContent from '../components/homeContent.js';
import loadProductsCatalog from '../components/productCatalog.js';

const routes = {
    home: loadHomeContent,
    products: loadProductsCatalog,
    //cart: 'path/to/shoppingCart.js',
    // Add more routes as needed
};

async function loadContent(route) {
    const contentElement = document.getElementById('content');
    const componentPath = routes[route];

    if (componentPath) {
        try {
            const componentHTML = await loadComponent(componentPath);
            contentElement.innerHTML = componentHTML;
        } catch (error) {
            contentElement.innerHTML = `<p>Error loading component: ${error.message}</p>`;
        }
    } else {
        contentElement.innerHTML = '<p>Route not recognized.</p>';
    }
}
