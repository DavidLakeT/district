const products = [
    { name: "Product 1", price: 10.99, image: "images/product1.jpg" },
    { name: "Product 2", price: 15.99, image: "images/product2.jpg" },
    // Add more products as needed
];

document.addEventListener("DOMContentLoaded", () => {
    const productContainer = document.getElementById("product-container");

    // Generate product cards
    products.forEach(product => {
        const productCard = document.createElement("div");
        productCard.classList.add("col-md-4", "mb-4");

        productCard.innerHTML = `
            <div class="card">
                <img src="${product.image}" class="card-img-top" alt="Product Image">
                <div class="card-body">
                    <h5 class="card-title">${product.name}</h5>
                    <p class="card-text">$${product.price.toFixed(2)}</p>
                    <a href="#" class="btn btn-primary">Add to Cart</a>
                </div>
            </div>
        `;

        productContainer.appendChild(productCard);
    });
});
