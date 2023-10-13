<template>
    <div class="product-reviews mt-5">
      <h3 class="review-title">Product Reviews</h3>
      <div v-if="reviews.length === 0" class="no-reviews">No reviews yet.</div>
      <ul v-else class="list-group">
        <li v-for="(review, index) in reviews" :key="index" class="list-group-item">
          <div class="review-author">{{ review.author }}</div>
          <div class="review-rating">
            <span v-for="star in review.rating" :key="star" class="star">&#9733;</span>
          </div>
          <div class="review-text">{{ review.text }}</div>
          <div class="review-date">{{ review.date }}</div>
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      reviews: {
        type: Array,
        default: () => []
      }
    }
  };
  </script>
  
  <style scoped>
  .product-reviews {
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 10px;
    margin-bottom: 4%;
  }
  
  .review-title {
    font-size: 1.5rem;
    font-weight: bold;
    margin-bottom: 15px;
  }
  
  .no-reviews {
    font-size: 1rem;
    color: #777;
  }
  
  .list-group {
    list-style: none;
    padding: 0;
  }
  
  .list-group-item {
    padding: 15px;
    border: 1px solid #ccc;
    border-radius: 10px;
    margin-bottom: 10px;
  }
  
  .review-author {
    font-weight: bold;
    margin-bottom: 5px;
  }
  
  .review-rating {
    margin-bottom: 5px;
  }
  
  .star {
    color: #f1c40f;
  }
  
  .review-text {
    margin-bottom: 5px;
  }
  
  .review-date {
    font-size: 0.8rem;
    color: #777;
  }
  </style>
  