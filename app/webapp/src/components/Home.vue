<template>
    <div class="container mt-5">
      <div class="row">
        <div class="col-md-8 offset-md-2">
          <div class="card">
            <div class="card-body">
              <h2 class="card-title mb-4">Welcome to District VulnWeb</h2>
              <p class="card-text">
                District VulnWeb is a deliberately vulnerable web application designed for educational purposes.
                It contains various security vulnerabilities that allow users to practice identifying and exploiting security flaws in a safe environment.
                By exploring and interacting with this application, users can enhance their understanding of cybersecurity concepts and techniques.
              </p>
  
              <h3 class="mt-4">Key Features</h3>
              <ul>
                <li>Realistic simulated vulnerabilities for hands-on learning</li>
                <li>User-friendly interface for easy navigation and interaction</li>
                <li>Comprehensive documentation on identified vulnerabilities</li>
                <li>Safe environment for experimenting with cybersecurity concepts</li>
                <li>Application designed and made with modern technologies to make it more realistic.</li>
              </ul>
  
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
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "Home"
  }
  </script>
  
  <style scoped>
  /* Add component-specific styles here */
  .card {
    border: 1px solid #ccc;
    border-radius: 10px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    margin-bottom: 7%;
  }
  
  .card-title {
    color: #333;
  }
  
  .card-text {
    color: #555;
  }
  
  /* Additional styles as needed */
  </style>
  