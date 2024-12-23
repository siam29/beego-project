
    const form = document.getElementById("cat-form");
    const breedSelect = document.getElementById("breed");
    const image = document.getElementById("cat-image");
    const description = document.getElementById("cat-description");
    const name = document.getElementById("cat-name");
    const origin = document.getElementById("cat-origin");
    const favoritesGrid = document.getElementById("favorites-grid");
  
    let currentImages = [];
    let currentImageIndex = 0;
    let currentEventSource = null;
    let breedData = [];

  
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
      document
          .querySelectorAll(".section")
          .forEach((el) => el.classList.remove("active"));
      document.getElementById(`${section}-section`).classList.add("active");
  
      // Update navigation button styles
      document
          .querySelectorAll(".nav-buttons button")
          .forEach((btn) => btn.classList.remove("button-active"));
      document
          .getElementById(`${section}-btn`)
          .classList.add("button-active");
  
      // Only call specific loaders where necessary
      if (section === "favs") {
          loadFavorites();
      }
  }
  

      async function loadRandomImages() {
        const response = await fetch("/random");
        currentImages = await response.json();
        currentImageIndex = 0;
        showCurrentImage();
      }
  
    // Load random images for voting
    async function loadVotingSection() {
      const response = await fetch("/random");
      currentImages = await response.json();
      currentImageIndex = 0;
      showCurrentImage();
    }
  
    function showCurrentImage() {
      if (currentImages.length === 0 || currentImageIndex >= currentImages.length) {
          console.error("No images available. Loading more images...");
          loadRandomImages();
          return;
      }
  
      document.getElementById("voting-image").src =
          currentImages[currentImageIndex].url;
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
    async function loadFavorites() {
        const response = await fetch("/favorites");
        const favorites = await response.json();
        const grid = document.getElementById("favorites-grid");
        grid.innerHTML = "";
        favorites.forEach((image) => {
          const img = document.createElement("img");
          img.src = image.url;
          img.alt = "Favorite Cat";
          grid.appendChild(img);
        });
      }

      
  
    // Handle breed form submission
    form.addEventListener("submit", (event) => {
        event.preventDefault();

        const breed = document.getElementById("breed").value;

        // Close existing EventSource if any
        if (window.eventSource) {
          window.eventSource.close();
        }

        // Open a new EventSource with the correct breed parameter
        window.eventSource = new EventSource(`/stream-breed?breed=${breed}`);

        // Listen for description updates
        window.eventSource.addEventListener("description", (event) => {
          console.log("Description received:", event.data); // Debugging log
          description.textContent = event.data || "No description available.";
        });

        // Listen for name updates
        window.eventSource.addEventListener("name", (event) => {
          console.log("Name received:", event.data); // Debugging log
          name.textContent = event.data || "No name available.";
        });

        window.eventSource.addEventListener("origin", (event) => {
          console.log("Origin received:", event.data); // Debugging log
          origin.textContent = event.data || "No origin available.";
        });

        // Listen for image updates
        window.eventSource.addEventListener("image", (event) => {
          console.log("Image URL received:", event.data); // Debugging log
          image.src = event.data;
        });

        // Handle errors
        window.eventSource.onerror = () => {
          console.error("EventSource failed.");
          window.eventSource.close();
        };
      });

  
    // Button click listeners for switching sections
    document.getElementById("breeds-btn").addEventListener("click", () => {
      showSection("breeds");
      loadBreedData("abys");
    });
  
    document.getElementById("voting-btn").addEventListener("click", () => {
      showSection("voting");
      loadRandomImages(); // Ensure images are fetched
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
  