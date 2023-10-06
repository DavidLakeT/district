const products = [
    { name: "Product 1", price: 10.99, image: "images/product1.jpg" },
    { name: "Product 2", price: 15.99, image: "images/product2.jpg" },
    // Add more products as needed
];

export default function generateProductsHTML() {
    let productsHTML = '';

    products.forEach(product => {
        productsHTML += `
            <div class="col-md-4 mb-4">
                <div class="card">
                    <img src="${product.image}" class="card-img-top" alt="Product Image">
                    <div class="card-body">
                        <h5 class="card-title">${product.name}</h5>
                        <p class="card-text">$${product.price.toFixed(2)}</p>
                        <a href="#" class="btn btn-primary">Add to Cart</a>
                    </div>
                </div>
            </div>
        `;
    });

    return `
        <div class="container mt-5">
            <div class="row" id="product-container">
                ${productsHTML}
            </div>
        </div>
    `;
}
