<template>
  <div class="container mt-5">
    <div class="row">
      <div class="col-md-8 offset-md-2">
        <HomeContent />
      </div>
    </div>
  </div>
</template>

<script>
import HomeContent from '@/components/home/HomeContent.vue';

export default {
  name: "Home",
  components: {
    HomeContent,
  },
};
</script>

<style scoped>
  @import '@/css/homeStyle.css';
</style>