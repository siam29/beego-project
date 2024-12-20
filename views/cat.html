<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Viewer</title>
</head>
<body>
    <h1>Cat Viewer</h1>
    <form id="cat-form">
        <label for="breed">Select Cat Breed:</label>
        <select name="breed" id="breed">
            {{ range .Breeds }}
            <option value="{{ .ID }}">{{ .Name }}</option>
            {{ end }}
        </select>
        <button type="submit">Show</button>
    </form>

    <h2>Cat Image</h2>
    <img id="cat-image" src="" alt="Cat Image" style="max-width: 500px; display: block;">

    <h2>Description</h2>
    <p id="cat-description">Select a breed to see its description.</p>

    <script>
        const form = document.getElementById("cat-form");
        const image = document.getElementById("cat-image");
        const description = document.getElementById("cat-description");

        form.addEventListener("submit", (event) => {
            event.preventDefault();

            const breed = document.getElementById("breed").value;

            // Close existing EventSource if any
            if (window.eventSource) {
                window.eventSource.close();
            }

            // Open a new EventSource
            window.eventSource = new EventSource(`/stream?breed=${breed}`);

            // Listen for image updates
            window.eventSource.addEventListener("image", (event) => {
                image.src = event.data;
            });

            // Listen for description updates
            window.eventSource.addEventListener("description", (event) => {
                description.textContent = event.data;
            });

            // Handle errors
            window.eventSource.onerror = () => {
                console.error("EventSource failed.");
                window.eventSource.close();
            };
        });
    </script>
</body>
</html>
