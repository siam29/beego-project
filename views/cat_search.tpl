<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Search</title>
    <style>
        #cat-images {
            max-width: 600px;
            margin: 20px auto;
            text-align: center;
        }
        img {
            width: 300px;
            height: 250px;
            display: none; /* Hide all images initially */
        }
    </style>
</head>
<body>
    <h1>Cat Images</h1>
    <form action="/" method="GET">
        <select name="breed">
    <<option value="abys" {{ if eq .SelectedBreed "abys" }}selected{{ end }}>Abyssinian</option>
    <option value="aege" {{ if eq .SelectedBreed "aege" }}selected{{ end }}>Aegean</option>
    <option value="abob" {{ if eq .SelectedBreed "bobs" }}selected{{ end }}>American Bobtail</option>
    <option value="acur" {{ if eq .SelectedBreed "curl" }}selected{{ end }}>American Curl</option>
    <option value="asho" {{ if eq .SelectedBreed "sh" }}selected{{ end }}>American Shorthair</option>
    <option value="awir" {{ if eq .SelectedBreed "wiri" }}selected{{ end }}>American Wirehair</option>
    <option value="amau" {{ if eq .SelectedBreed "amau" }}selected{{ end }}>Arabian Mau</option>
    <option value="amis" {{ if eq .SelectedBreed "ausm" }}selected{{ end }}>Australian Mist</option>
    <option value="bali" {{ if eq .SelectedBreed "bali" }}selected{{ end }}>Balinese</option> <!-- Added Balinese -->
    <option value="bamb" {{ if eq .SelectedBreed "bamb" }}selected{{ end }}>Bambino</option> <!-- Added Bambino -->
    <option value="beng" {{ if eq .SelectedBreed "beng" }}selected{{ end }}>Bengal</option>
    <option value="birm" {{ if eq .SelectedBreed "birman" }}selected{{ end }}>Birman</option>
    <option value="bomb" {{ if eq .SelectedBreed "bomb" }}selected{{ end }}>Bombay</option>
    <option value="bslo" {{ if eq .SelectedBreed "blh" }}selected{{ end }}>British Longhair</option>
    <option value="bsho" {{ if eq .SelectedBreed "bsh" }}selected{{ end }}>British Shorthair</option>
    <option value="bure" {{ if eq .SelectedBreed "burm" }}selected{{ end }}>Burmese</option>
    <option value="buri" {{ if eq .SelectedBreed "burm" }}selected{{ end }}>Burmilla</option>
    <option value="cspa" {{ if eq .SelectedBreed "casp" }}selected{{ end }}>California Spangled</option>
    <option value="ctif" {{ if eq .SelectedBreed "ct" }}selected{{ end }}>Chantilly-Tiffany</option>
    <option value="char" {{ if eq .SelectedBreed "chtr" }}selected{{ end }}>Chartreux</option>
    <option value="chau" {{ if eq .SelectedBreed "chaus" }}selected{{ end }}>Chausie</option>
    <option value="chee" {{ if eq .SelectedBreed "cheet" }}selected{{ end }}>Cheetoh</option>
    <option value="csho" {{ if eq .SelectedBreed "cpsh" }}selected{{ end }}>Colorpoint Shorthair</option>
    <option value="crex" {{ if eq .SelectedBreed "corn" }}selected{{ end }}>Cornish Rex</option>
    <option value="cymr" {{ if eq .SelectedBreed "cymr" }}selected{{ end }}>Cymric</option>
    <option value="cypr" {{ if eq .SelectedBreed "cypr" }}selected{{ end }}>Cyprus</option>
    <option value="drex" {{ if eq .SelectedBreed "dev" }}selected{{ end }}>Devon Rex</option>
    <option value="dons" {{ if eq .SelectedBreed "don" }}selected{{ end }}>Donskoy</option>
    <option value="lihu" {{ if eq .SelectedBreed "dragon" }}selected{{ end }}>Dragon Li</option>
    <option value="emau" {{ if eq .SelectedBreed "emau" }}selected{{ end }}>Egyptian Mau</option>
    <option value="ebur" {{ if eq .SelectedBreed "ebur" }}selected{{ end }}>European Burmese</option>
    <option value="esho" {{ if eq .SelectedBreed "esh" }}selected{{ end }}>Exotic Shorthair</option>
    <option value="hbro" {{ if eq .SelectedBreed "havn" }}selected{{ end }}>Havana Brown</option>
    <option value="hima" {{ if eq .SelectedBreed "him" }}selected{{ end }}>Himalayan</option>
    <option value="jbob" {{ if eq .SelectedBreed "jbb" }}selected{{ end }}>Japanese Bobtail</option>
    <option value="java" {{ if eq .SelectedBreed "java" }}selected{{ end }}>Javanese</option>
    <option value="khao" {{ if eq .SelectedBreed "kmanee" }}selected{{ end }}>Khao Manee</option>
    <option value="kora" {{ if eq .SelectedBreed "korat" }}selected{{ end }}>Korat</option>
    <option value="kuri" {{ if eq .SelectedBreed "kuril" }}selected{{ end }}>Kurilian</option>
    <option value="lape" {{ if eq .SelectedBreed "lperm" }}selected{{ end }}>LaPerm</option>
    <option value="mcoo" {{ if eq .SelectedBreed "coon" }}selected{{ end }}>Maine Coon</option>
    <option value="mala" {{ if eq .SelectedBreed "mal" }}selected{{ end }}>Malayan</option>
    <option value="manx" {{ if eq .SelectedBreed "manx" }}selected{{ end }}>Manx</option>
    <option value="munc" {{ if eq .SelectedBreed "munch" }}selected{{ end }}>Munchkin</option>
    <option value="nebe" {{ if eq .SelectedBreed "nebel" }}selected{{ end }}>Nebelung</option>
    <option value="norw" {{ if eq .SelectedBreed "nfc" }}selected{{ end }}>Norwegian Forest Cat</option>
    <option value="ocic" {{ if eq .SelectedBreed "ocicat" }}selected{{ end }}>Ocicat</option>
    <option value="orie" {{ if eq .SelectedBreed "orient" }}selected{{ end }}>Oriental</option>
    <option value="pers" {{ if eq .SelectedBreed "pers" }}selected{{ end }}>Persian</option>
    <option value="pixi" {{ if eq .SelectedBreed "pixiebob" }}selected{{ end }}>Pixie-bob</option>
    <option value="raga" {{ if eq .SelectedBreed "rag" }}selected{{ end }}>Ragamuffin</option>
    <option value="ragd" {{ if eq .SelectedBreed "ragd" }}selected{{ end }}>Ragdoll</option>
    <option value="rblu" {{ if eq .SelectedBreed "rblu" }}selected{{ end }}>Russian Blue</option>
    <option value="sava" {{ if eq .SelectedBreed "sav" }}selected{{ end }}>Savannah</option>
    <option value="sfol" {{ if eq .SelectedBreed "scof" }}selected{{ end }}>Scottish Fold</option>
    <option value="srex" {{ if eq .SelectedBreed "selk" }}selected{{ end }}>Selkirk Rex</option>
    <option value="siam" {{ if eq .SelectedBreed "siam" }}selected{{ end }}>Siamese</option>
    <option value="sibe" {{ if eq .SelectedBreed "sibe" }}selected{{ end }}>Siberian</option>
    <option value="sing" {{ if eq .SelectedBreed "sing" }}selected{{ end }}>Singapura</option>
    <option value="snow" {{ if eq .SelectedBreed "snow" }}selected{{ end }}>Snowshoe</option>
    <option value="soma" {{ if eq .SelectedBreed "somali" }}selected{{ end }}>Somali</option>
    <option value="sphy" {{ if eq .SelectedBreed "sph" }}selected{{ end }}>Sphynx</option>
    <option value="tonk" {{ if eq .SelectedBreed "tonk" }}selected{{ end }}>Tonkinese</option>
    <option value="toyg" {{ if eq .SelectedBreed "toyge" }}selected{{ end }}>Toyger</option>
    <option value="tang" {{ if eq .SelectedBreed "tang" }}selected{{ end }}>Turkish Angora</option>
    <option value="tvan" {{ if eq .SelectedBreed "van" }}selected{{ end }}>Turkish Van</option>
    <option value="ycho" {{ if eq .SelectedBreed "york" }}selected{{ end }}>York Chocolate</option>
        </select>
        <button type="submit">Search</button>
    </form>

    <div id="cat-images">
        {{ if .Images }}
            {{ range .Images }}
                <img src="{{ .URL }}" alt="Cat image" class="cat-slide">
            {{ end }}
        {{ else }}
            <p>No images available for the selected breed.</p>
        {{ end }}
    </div>

    <script>
        let currentIndex = 0;
        const images = document.querySelectorAll('.cat-slide');

        // Function to show the next image
        function showNextImage() {
            // Hide all images
            images.forEach(img => img.style.display = 'none');
            
            // Show the current image
            images[currentIndex].style.display = 'block';

            // Update the index for the next image
            currentIndex = (currentIndex + 1) % images.length; // Loops back to 0
        }

        // Initial call to show the first image
        if (images.length > 0) {
            showNextImage();
            setInterval(showNextImage, 3000); // Change image every 3 seconds
        }
    </script>
</body>
</html>
