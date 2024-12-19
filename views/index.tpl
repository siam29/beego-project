<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Images</title>
    <style>
        img { max-width: 100%; height: auto; }
        .container { text-align: center; padding: 20px; }
        .cat-image { margin-top: 20px; }
    </style>
</head>
<body>

<div class="container">
    <h1>Cat Images</h1>
    <input type="text" id="breed-id" placeholder="Enter Breed ID (e.g. beng)" />
    <button onclick="fetchCatImage()">Fetch Cat Image</button>

    <div class="cat-image" id="cat-image"></div>
</div>

<script>
    function fetchCatImage() {
        const breedID = document.getElementById("breed-id").value || "beng"; // Default breed: beng
        fetch(`/cat/image?breed_id=${breedID}`)
            .then(response => response.json())
            .then(data => {
                const catImageDiv = document.getElementById("cat-image");
                if (data.image_url) {
                    catImageDiv.innerHTML = `<img src="${data.image_url}" alt="Cat Image">`;
                } else {
                    catImageDiv.innerHTML = `<p>Error: ${data.error || "No image found"}</p>`;
                }
            })
            .catch(error => {
                document.getElementById("cat-image").innerHTML = "<p>Failed to fetch image</p>";
            });
    }
</script>

</body>
</html>
