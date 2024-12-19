function fetchCatImages() {
    const breed = document.getElementById('catBreedSelect').value;
    if (!breed) {
        alert("Please select a breed.");
        return;
    }

    fetch(`/catimages?breed=${breed}`)
        .then(response => response.json())
        .then(data => {
            const catResult = document.getElementById('catResult');
            catResult.innerHTML = ''; // Clear any existing content

            if (data.ImageURLs && data.ImageURLs.length > 0) {
                data.ImageURLs.forEach(url => {
                    const img = document.createElement('img');
                    img.src = url;
                    img.alt = data.Breed;
                    img.style.maxWidth = '80%';
                    img.style.marginBottom = '10px';
                    catResult.appendChild(img);
                });
            } else {
                catResult.innerHTML = '<p>No images found for this breed.</p>';
            }
        })
        .catch(error => {
            console.error('Error fetching cat images:', error);
            alert('Failed to fetch cat images.');
        });
}
