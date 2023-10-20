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
        <textarea v-model="newReview.text" class="form-control" id="text" rows="3" required></textarea>
      </div>
      <button type="submit" class="btn btn-primary">Submit Review</button>
    </form>
    <button @click="toggleReviewForm" v-else class="btn btn-primary mt-3">Add a Review</button>
  </div>
</template>

<script>
export default {
  props: {
    reviews: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      showReviewForm: false,
      newReview: {
        author: "",
        rating: 1,
        text: "",
      },
    };
  },
  methods: {
    submitReview() {
      // Implement submitting the review (you can send it to a server or handle it locally)
      console.log("New review submitted:", this.newReview);

      // Reset the newReview object and hide the form
      this.newReview = { author: "", rating: 1, text: "" };
      this.showReviewForm = false;
    },
    toggleReviewForm() {
      this.showReviewForm = !this.showReviewForm;
    },
  },
};
</script>

<style scoped>
@import '@/css/reviewStyle.css';
</style>
  