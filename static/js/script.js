const form = document.getElementById("cat-form");
      const image = document.getElementById("cat-image");
      const description = document.getElementById("cat-description");
      const name = document.getElementById("cat-name");
      const origin = document.getElementById("cat-origin");

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

      let currentImages = [];
      let currentImageIndex = 0;
      let currentEventSource = null;
      let breedData = [];

      // Show the voting section initially
      showSection("voting");

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

        if (section === "voting") {
          loadRandomImages();
        } else if (section === "favs") {
          loadFavorites();
        }
      }

      // Load random images for voting
      async function loadRandomImages() {
        const response = await fetch("/random");
        currentImages = await response.json();
        currentImageIndex = 0;
        showCurrentImage();
      }

      function showCurrentImage() {
        if (currentImageIndex >= currentImages.length) {
          loadRandomImages();
          return;
        }
        document.getElementById("voting-image").src =
          currentImages[currentImageIndex].url;
      }

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

      // Load favorite images
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

      // Load breed data and populate the dropdown
      async function loadBreeds() {
        const response = await fetch("/breeds");
        breedData = await response.json();

        const breedSelect = document.getElementById("breed-select");
        breedData.forEach((breed) => {
          const option = document.createElement("option");
          option.value = breed.name;
          option.textContent = breed.name;
          breedSelect.appendChild(option);
        });

        const breedSelect1 = document.getElementById("breed-select");
        breedData.forEach((breed) => {
          const option = document.createElement("option");
          option.value = breed.origin;
          option.textContent = breed.origin;
          breedSelect.appendChild(option);
        });

        breedSelect.addEventListener("change", function () {
          const selectedBreed = breedData.find(
            (breed) => breed.name === breedSelect.value
          );
          if (selectedBreed) {
            document.getElementById("breed-image").src = selectedBreed.image;
            document.getElementById("breed-description").textContent =
              selectedBreed.description;
            document.getElementById("breed-origin").textContent =
              selectedBreed.description;
          } else {
            document.getElementById("breed-image").src = "";
            document.getElementById("breed-description").textContent =
              "Select a breed to see its description.";
          }
        });
      }

      // Initial load of breeds when the page loads
      loadBreeds();