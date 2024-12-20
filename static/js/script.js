document.getElementById("breed-select").addEventListener("change", function() {
    const breedID = this.value;
    if (breedID) {
        fetchBreedData(breedID);
    }
});

function fetchBreedData(breedID) {
    fetch(`/cat/getBreedData?breed=${breedID}`)
        .then(response => response.json())
        .then(data => {
            displayBreedData(data);
        })
        .catch(error => console.error("Error fetching breed data:", error));
}

function displayBreedData(data) {
    const descriptionElement = document.getElementById("breed-description");
    const imagesElement = document.getElementById("cat-images");

    // Clear previous content
    descriptionElement.innerHTML = "";
    imagesElement.innerHTML = "";

    // Display breed description
    descriptionElement.innerHTML = `<h2>Description</h2><p>${data.Description}</p>`;

    // Display breed images
    let imageIndex = 0;
    const images = data.Images;
    const imageInterval = setInterval(() => {
        if (imageIndex < images.length) {
            const imgElement = document.createElement("img");
            imgElement.src = images[imageIndex];
            imgElement.alt = "Cat Image";
            imagesElement.appendChild(imgElement);
            imageIndex++;
        } else {
            clearInterval(imageInterval);  // Stop after displaying 5 images
        }
    }, 3000);  // Show images every 3000ms
}
