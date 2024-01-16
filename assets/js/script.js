function showImageGallery() {
    var gallery = document.getElementById('image-gallery');
    gallery.style.display = 'flex';
}

function selectImage(selectedImage) {
    var preview = document.getElementById('selected-image-preview');
    var imageProfilField = document.getElementById('imageProfilField');

    while (preview.firstChild) {
        preview.removeChild(preview.firstChild);
    }

    var img = document.createElement('img');
    img.src = selectedImage.src;
    img.style.width = '100px';
    img.style.height = '100px';
    preview.appendChild(img);

    imageProfilField.value = selectedImage.src;

    const input = document.querySelector('.input-img');
    console.log(input.value)
    input.value = selectedImage.src;

}

function confirmDelete(nomPersonnage) {
    // Affiche une boîte de dialogue de confirmation
    if (confirm(`Confirmez-vous la suppression du personnage ${nomPersonnage} ?`)) {
        // Si l'utilisateur confirme, envoie une demande de suppression au serveur
        deleteCharacter(nomPersonnage);
    }
}