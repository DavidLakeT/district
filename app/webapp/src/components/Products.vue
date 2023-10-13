<template>
    <div class="container mt-5">
      <div class="row">
        <div class="col-md-10 offset-md-1">
          <div class="row">
            <div v-for="product in products" :key="product.id" class="col-md-4 mb-4">
              <router-link :to="'/products/' + product.id" class="card-link">
                <div class="card">
                  <img :src="imageUrl(product.id)" class="card-img-top" :alt="product.name">
                  <div class="card-body">
                    <h5 class="card-title">{{ product.name }}</h5>
                    <p class="card-text">$ {{ product.price.toFixed(2) }}</p>
                    <a href="#" class="btn btn-primary">Add to Cart</a>
                  </div>
                </div>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "Products",
    data() {
      return {
        products: [
          { id: 1, name: "Product 1", price: 10.99 },
          { id: 2, name: "Product 2", price: 15.99}
        ]
      };
    },
    computed: {
        imageUrl() {
            return (productId) => `/images/product${productId}.jpg`;
        }
    }
  };
  </script>
  
  <style scoped>
  .card {
    border: 1px solid #eee;
    border-radius: 10px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s ease;
  }
  
  .card:hover {
    transform: scale(1.05);
  }
  
  .card-title {
    font-size: 1.25rem;
    font-weight: bold;
    margin-bottom: 0.5rem;
  }
  
  .card-text {
    font-size: 1rem;
    margin-bottom: 0.5rem;
  }
  
  .btn-primary {
    background-color: #007BFF;
    border: none;
  }
  
  .btn-primary:hover {
    background-color: #0056b3;
  }
  
  /* Style for the router-link to remove the outline */
  .card-link {
    text-decoration: none;
    outline: none;
  }
  </style>
  
  