<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Search</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    <div class="container">
        <h1>Search for Cat Breeds</h1>

        <!-- Dropdown for Cat Breeds -->
        <select id="catBreedSelect" class="dropdown" onchange="fetchCatImage()">
            <option value="" selected disabled>Select Cat Breed</option>
            <option value="abys">Abyssinian</option>
            <option value="aege">Aegean</option>
            <option value="bobs">American Bobtail</option>
            <option value="curl">American Curl</option>
            <option value="sh">American Shorthair</option>
            <option value="abys">American Wirehair</option>
            <option value="abys">Arabian Mau</option>
            <option value="abys">Australian Mist</option>
            <option value="abys">Balinese</option>
            <option value="abys">Bambino</option>
            <option value="abys">Bengal</option>
            <option value="abys">Birman</option>
            <option value="abys">Bombay</option>
            <option value="abys">British Longhair</option>
            <option value="abys">British Shorthair</option>
            <option value="abys">Burmese</option>
            <option value="abys">Burmilla</option>
            <option value="abys">California Spangled</option>
            <option value="abys">Chantilly-Tiffany</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
            <option value="abys">Birman</option>
        </select>

        <!-- Display Cat Image -->
        <div id="catResult" class="result">
            <img id="catImage" src="" alt="" style="display: none;">
            <p id="catDescription" style="display: none;"></p>
        </div>
    </div>

    <script src="/static/js/script.js"></script> 
</body>
</html>
