<template>
  <div class="card">
    <div class="card-body">
      <h2 class="card-title mb-4">Welcome to District VulnWeb</h2>
      <p class="card-text">
        District VulnWeb is a deliberately vulnerable web application designed for educational purposes.
        It contains various security vulnerabilities that allow users to practice identifying and exploiting security flaws in a safe environment.
        By exploring and interacting with this application, users can enhance their understanding of cybersecurity concepts and techniques.
      </p>

      <h3 class="mt-4">Key Features</h3>
      <HomeFeatures />

      <h3 class="mt-4">About District VulnWeb</h3>
      <p>
        District VulnWeb was created as a part of an educational project to help individuals learn and practice cybersecurity skills.
        This application is intentionally vulnerable, showcasing common security weaknesses found in web applications.
        Users are encouraged to explore the application, identify vulnerabilities, and learn how to mitigate them effectively.
      </p>

      <p class="text-muted mt-4">
        Note: This application is deliberately vulnerable and should only be used for educational purposes.
        Do not attempt to exploit or misuse any vulnerabilities in a real-world scenario.
      </p>
    </div>
  </div>
</template>

<script>
import HomeFeatures from '@/components/home/HomeFeatures.vue';

export default {
  name: "HomeContent",
  components: {
    HomeFeatures,
  },
};
</script>