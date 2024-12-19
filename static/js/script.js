function fetchCatImage() {
    var breed = document.getElementById("catBreedSelect").value;

    if (breed) {
        fetch(`/catimage?breed=${breed}`)
            .then(response => response.json())
            .then(data => {
                // Handle the response data and display the cat image
                if (data.ImageURL) {
                    document.getElementById("catImage").style.display = "block";
                    document.getElementById("catImage").src = data.ImageURL;
                    document.getElementById("catDescription").style.display = "block";
                    document.getElementById("catDescription").textContent = "Breed: " + data.Breed;
                } else {
                    document.getElementById("catImage").style.display = "none";
                    document.getElementById("catDescription").style.display = "none";
                }
            })
            .catch(error => console.error("Error fetching cat image:", error));
    }
}
