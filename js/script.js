function rechercherPokemon() {
    // Récupère le nom du Pokemon entré dans le formulaire
    var pokemon = document.getElementById("pokemon").value;
    
    // Envoie une requête à l'API PokeAPI pour récupérer les informations sur le Pokemon
    fetch(`https://pokeapi.co/api/v2/pokemon/${pokemon}`)
      .then(response => response.json())
      .then(data => {
        // Crée une carte avec les informations du Pokemon
        var resultat = document.getElementById("resultat");
        resultat.innerHTML = `
          <div>
            <h2>${data.name}</h2>
            <p>Type: ${data.types[0].type.name}</p>
            <img src="${data.sprites.front_default}">
          </div>
        `;
      })
      .catch(error => {
        // Affiche une erreur si le Pokemon n'est pas trouvé
        var resultat = document.getElementById("resultat");
        resultat.innerHTML = "<p>Pokemon non trouvé.</p>";
      });
  }
  