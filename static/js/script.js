let currentImages = [];
let currentImageIndex = 0;

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

function vote(type) {
  if (currentImageIndex < currentImages.length) {
    console.log(`Voted ${type} for image:`, currentImages[currentImageIndex]);
    currentImageIndex++; // Move to the next image
    showCurrentImage(); // Display the next image
  } else {
    console.log("End of current images. Fetching more...");
    loadRandomImages(); // Fetch new images if needed
  }
}

// Attach button click handlers
document.querySelector(".vote-buttons button:nth-child(1)").addEventListener("click", function () {
  vote("up"); // Handle "Upvote" action
});
document.querySelector(".vote-buttons button:nth-child(2)").addEventListener("click", function () {
  vote("favorite"); // Handle "Favorite" action
});
document.querySelector(".vote-buttons button:nth-child(3)").addEventListener("click", function () {
  vote("down"); // Handle "Downvote" action
});

// Load initial images when the voting section is activated
function showSection(section) {
  document
    .querySelectorAll(".section")
    .forEach((el) => el.classList.remove("active"));
  document.getElementById(`${section}-section`).classList.add("active");

  // Update navigation button styles
  document
    .querySelectorAll(".nav-buttons button")
    .forEach((btn) => btn.classList.remove("button-active"));
  document.getElementById(`${section}-btn`).classList.add("button-active");

  if (section === "voting") {
    loadRandomImages(); // Load a batch of images when the voting section is shown
  } else if (section === "favs") {
    loadFavorites();
  }
}

showSection("voting");