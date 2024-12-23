document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("cat-form");
    const breedSelect = document.getElementById("breed");
    const image = document.getElementById("cat-image");
    const description = document.getElementById("cat-description");
    const name = document.getElementById("cat-name");
    const origin = document.getElementById("cat-origin");
    const favoritesGrid = document.getElementById("favorites-grid");
  
    let currentEventSource = null;
    let currentImages = [];
    let currentImageIndex = 0;
    let breedData = [];
    let favoriteImages = []; // Local favorites array (for UI)
  
    // Load breed data and automatically show Abyssinian
    async function loadBreeds() {
      try {
        const response = await fetch("/breeds");
        if (!response.ok) throw new Error("Failed to fetch breeds.");
        
        breedData = await response.json();
        breedSelect.innerHTML = "";
  
        // Populate breed dropdown
        breedData.forEach((breed) => {
          const option = document.createElement("option");
          option.value = breed.id;
          option.textContent = `${breed.name} (${breed.origin})`;
          breedSelect.appendChild(option);
        });
  
        // Load breed data for Abyssinian by default
        loadBreedData("abys");
      } catch (error) {
        console.error("Error loading breeds:", error.message);
      }
    }
  
    // Load breed data using EventSource
    function loadBreedData(breedID) {
      if (currentEventSource) currentEventSource.close();
  
      currentEventSource = new EventSource(`/stream-breed?breed=${breedID}`);
  
      currentEventSource.addEventListener("description", (event) => {
        description.textContent = event.data || "No description available.";
      });
  
      currentEventSource.addEventListener("name", (event) => {
        name.textContent = event.data || "No name available.";
      });
  
      currentEventSource.addEventListener("origin", (event) => {
        origin.textContent = event.data || "No origin available.";
      });
  
      currentEventSource.addEventListener("image", (event) => {
        image.src = event.data || "placeholder.png";
      });
  
      currentEventSource.onerror = () => {
        console.error("EventSource error occurred.");
        currentEventSource.close();
      };
    }
  
    // Show the correct section when a button is clicked
    function showSection(section) {
      document.querySelectorAll(".section").forEach((el) => el.classList.remove("active"));
      document.getElementById(`${section}-section`).classList.add("active");
      document.querySelectorAll(".nav-buttons button").forEach((btn) => btn.classList.remove("button-active"));
      document.getElementById(`${section}-btn`).classList.add("button-active");
  
      if (section === "voting") {
        loadVotingSection();
      } else if (section === "favs") {
        loadFavoritesSection();
      }
    }
  
    // Load random images for voting
    async function loadVotingSection() {
      const response = await fetch("/random");
      currentImages = await response.json();
      currentImageIndex = 0;
      showCurrentImage();
    }
  
    function showCurrentImage() {
      if (currentImageIndex >= currentImages.length) {
        loadVotingSection();
        return;
      }
      document.getElementById("voting-image").src = currentImages[currentImageIndex].url;
    }
  
    // Handle vote (up or down) and add to favorites
    async function vote(voteType) {
      const image = currentImages[currentImageIndex];
      await fetch("/vote", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          image_id: image.id,
          vote_type: voteType,
        }),
      });
  
      // Add image to favorites if "up" (like) vote is chosen
      if (voteType === "up") {
        addToFavorites(image);
      }
  
      currentImageIndex++;
      showCurrentImage();
    }
  
    // Add image to favorites and update UI
    function addToFavorites(image) {
      // Prevent duplicate favorites
      if (!favoriteImages.find((fav) => fav.id === image.id)) {
        favoriteImages.push(image);
      }
  
      loadFavoritesSection(); // Refresh the favorites section
    }
  
    // Load and display the favorites section
    function loadFavoritesSection() {
      favoritesGrid.innerHTML = ""; // Clear current favorites
  
      // Display each favorite image
      favoriteImages.forEach((image) => {
        const img = document.createElement("img");
        img.src = image.url;
        img.alt = "Favorite Cat";
        favoritesGrid.appendChild(img);
      });
    }
  
    // Handle breed form submission
    form.addEventListener("submit", (event) => {
        
      event.preventDefault();
      const breedID = breedSelect.value;
      loadBreedData(breedID); // Load the selected breed's data
    });
  
    // Button click listeners for switching sections
    document.getElementById("breeds-btn").addEventListener("click", () => {
      showSection("breeds");
      loadBreedData("abys");
    });
  
    document.getElementById("voting-btn").addEventListener("click", () => {
      showSection("voting");
    });
  
    document.getElementById("favs-btn").addEventListener("click", () => {
      showSection("favs");
    });
  
    // Set up voting buttons
    document.getElementById("up-vote-btn").addEventListener("click", () => {
      vote("up");
    });
  
    document.getElementById("down-vote-btn").addEventListener("click", () => {
      vote("down");
    });
  
    // Initial setup: load breeds and show the voting section
    loadBreeds();
    showSection("voting");
  });
  