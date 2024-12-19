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
            <option value="abys">Chartreux </option>
            <option value="abys">Chausie </option>
            <option value="abys">Cheetoh </option>
            <option value="abys">Colorpoint Shorthair</option>
            <option value="abys">Cornish Rex</option>
            <option value="abys">Cymric</option>
            <option value="abys">Cyprus</option>
            <option value="abys">Devon Rex</option>
            <option value="abys">Donskoy</option>
            <option value="abys">Dragon Li</option>
            <option value="abys">Egyptian Mau</option>
            <option value="abys">European Burmese</option>
            <option value="abys">Exotic Shorthair</option>
            <option value="abys">Havana Brown</option>
            <option value="abys">Himalayan </option>
            <option value="abys">Japanese Bobtail</option>
            <option value="abys">Javanese</option>
            <option value="abys">Khao Manee</option>
            <option value="abys">Korat </option>

            <option value="abys">Kurilian</option>

            <option value="abys">LaPerm</option>

            <option value="abys">Maine Coon</option>

            <option value="abys">Malayan</option>

            <option value="abys">Manx </option>

            <option value="abys">Munchkin</option>

            <option value="abys">Nebelung</option>

            <option value="abys">Norwegian Forest Cat</option>

            <option value="abys">Ocicat</option>

            <option value="abys">Oriental</option>

            <option value="abys">Persian</option>

            <option value="abys">Pixie-bob</option>

            <option value="abys">Ragamuffin</option>

            <option value="abys">Ragdoll</option>
            <option value="abys">Russian Blue</option>
            <option value="abys">Savannah</option>
            <option value="abys">Scottish Fold</option>
            <option value="abys">Selkirk Rex</option>
            <option value="abys">Siamese</option>
            <option value="abys">Siberian</option>
            <option value="abys">Singapura</option>
            <option value="abys">Snowshoe</option>
            <option value="abys">Somali </option>
            <option value="abys">Sphynx</option>
            <option value="abys">Tonkinese</option>
            <option value="abys">Toyge</option>
            <option value="abys">Turkish Angora</option>
            <option value="abys">Turkish Van</option>
            <option value="abys">York Chocolate</option>

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
