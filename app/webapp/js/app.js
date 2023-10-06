import generateHomeContent from '../components/home.js';
import generateProductsHTML from '../components/productCatalog.js';

const routes = {
    home: generateHomeContent,
    products: generateProductsHTML,
    // Add more routes as needed
};

window.loadContent = async function(route = 'home') {
    const contentElement = document.getElementById('content');
    const loadFunction = routes[route];

    if (loadFunction) {
        try {
            const componentHTML = loadFunction();
            contentElement.innerHTML = componentHTML;
        } catch (error) {
            contentElement.innerHTML = `<p>Error loading component: ${error.message}</p>`;
        }
    } else {
        contentElement.innerHTML = '<p>Route not recognized.</p>';
    }
};

const handleNavigation = (route) => {
    loadContent(route);
};

const navLinks = document.querySelectorAll('.nav-link');
navLinks.forEach(link => {
    link.addEventListener('click', (event) => {
        event.preventDefault();
        const route = event.target.getAttribute('data-route');
        handleNavigation(route);
    });
});

loadContent();