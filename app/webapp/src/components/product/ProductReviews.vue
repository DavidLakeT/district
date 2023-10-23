<template>
  <div class="product-reviews mt-5">
    <h3 class="review-title">Product Reviews</h3>
    <div v-if="reviews.length === 0" class="no-reviews">No reviews yet.</div>
    <ul v-else class="list-group">
      <li v-for="(review, index) in reviews" :key="index" class="list-group-item">
        <div class="review-author">{{ review.user_email }}</div>
        <div class="review-rating">
          <span v-for="star in review.rating" :key="star" class="star">&#9733;</span>
        </div>
        <div class="review-text">{{ review.content }}</div>
        <div class="review-date">{{ review.date }}</div>
      </li>
    </ul>

    <!-- Form to add a new review -->
    <form @submit.prevent="submitReview" v-if="showReviewForm" class="mt-3">
      <div class="form-group">
        <!--<label for="author">Your Name:</label>-->
        <input hidden type="text" v-model="newReview.author" class="form-control" id="author">
      </div>
      <div class="form-group">
        <label for="rating">Rating:</label>
        <select v-model="newReview.rating" class="form-control" id="rating" required>
          <option value="1">1 star</option>
          <option value="2">2 stars</option>
          <option value="3">3 stars</option>
          <option value="4">4 stars</option>
          <option value="5">5 stars</option>
        </select>
      </div>
      <div class="form-group">
        <label for="text">Review:</label>
        <textarea v-model="newReview.content" class="form-control" id="content" rows="3" required></textarea>
      </div>
      <button type="submit" class="btn btn-primary">Submit Review</button>
    </form>
    <button @click="toggleReviewForm" v-else class="btn btn-primary mt-3">Add a Review</button>
  </div>
</template>

<script>
import { submitReviewForm } from '@/js/review.js'
export default {
  props: {
    product: Object,
    reviews: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      showReviewForm: false,
      newReview: {
        user_id: "",
        rating: 1,
        content: "",
        product_id: parseInt(this.$route.params.id, 10),
      },
    };
  },
  methods: {
    toggleReviewForm() {
      this.showReviewForm = !this.showReviewForm;
    },
    async submitReview() {
      try {
        const response = await submitReviewForm(this.newReview);
        if(response){
          window.location.reload();
          this.toggleReviewForm();
        }
      } catch (error) {
        console.error('Error submitting review:', error);
      }
    }
  },
};
</script>

<style scoped>
@import '@/css/reviewStyle.css';
</style>
  