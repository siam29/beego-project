
const form = document.getElementById("cat-form");
const breedSelect = document.getElementById("breed");
const image = document.getElementById("cat-image");
const description = document.getElementById("cat-description");
const name = document.getElementById("cat-name");
const origin = document.getElementById("cat-origin");
const favoritesGrid = document.getElementById("favorites-grid");

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

  window.eventSource.addEventListener("wikipedia", (event) => {
    const wikipediaElement = document.getElementById("cat-wikipedia");
    if (event.data) {
      wikipediaElement.innerHTML = `<a href="${event.data}" target="_blank" rel="noopener noreferrer">Wikipedia</a>`;
    } else {
      wikipediaElement.textContent = "Select a breed to see its Wikipedia link.";
    }
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

showSection("voting");


// Load breed data and automatically show Abyssinian
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

    // Automatically select the Abyssinian breed and show its Wikipedia link
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

  currentEventSource.addEventListener("wikipedia", (event) => {
    const wikipediaElement = document.getElementById("cat-wikipedia");
    if (event.data) {
      wikipediaElement.innerHTML = `<a href="${event.data}" target="_blank" rel="noopener noreferrer">Wikipedia</a>`;
    } else {
      wikipediaElement.textContent = "Select a breed to see its Wikipedia link.";
    }
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

  if (section === "voting") {
    loadRandomImages();
  } else if (section === "favs") {
    loadFavorites();
  }
}



async function loadRandomImages() {
  const voteButtons = document.querySelectorAll(".vote-buttons button");
  voteButtons.forEach((btn) => (btn.disabled = true)); // Disable all vote buttons while fetching

  try {
    const response = await fetch("/random");
    currentImages = await response.json();
    currentImageIndex = 0;
    showCurrentImage();
  } catch (error) {
    console.error("Failed to load images:", error);
  } finally {
    voteButtons.forEach((btn) => (btn.disabled = false)); // Re-enable vote buttons after fetching
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
  if (currentImages.length === 0) {
    console.error("No images available.");
    document.getElementById("voting-image").alt = "No images available.";
    return;
  }

  if (currentImageIndex < currentImages.length) {
    const image = currentImages[currentImageIndex];
    document.getElementById("voting-image").src = image.url;
    document.getElementById("voting-image").alt = "Cat Image";
  } else {
    console.log("Reloading images...");
    document.getElementById("voting-image").alt = "Loading new images...";
    loadRandomImages();
  }
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


