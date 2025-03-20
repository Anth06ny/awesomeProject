<template>
  <div class="home-container">
    <section class="hero">
      <h1>Bienvenue à la Librairie Ludique 📚🎲</h1>
      <p>Découvrez et empruntez vos livres et jeux de plateau préférés.</p>
      <div class="buttons">
        <router-link to="/livres" class="btn">Explorer les Livres</router-link>
        <router-link to="/jeux" class="btn">Explorer les Jeux</router-link>
      </div>
    </section>

    <section class="presentation">
      <h2>Notre Collection</h2>
      <div class="items-grid">
        <div class="item">
          <img src="https://via.placeholder.com/150" alt="Livre 1">
          <p>Le Seigneur des Anneaux</p>
        </div>
        <div class="item">
          <img src="https://via.placeholder.com/150" alt="Livre 2">
          <p>1984 - George Orwell</p>
        </div>
        <div class="item">
          <img src="https://via.placeholder.com/150" alt="Jeu 1">
          <p>Catan</p>
        </div>
        <div class="item">
          <img src="https://via.placeholder.com/150" alt="Jeu 2">
          <p>Les Aventuriers du Rail</p>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "Accueil",
};
</script>

<style scoped>
.home-container {
  text-align: center;
  padding: 40px 20px;
}

.hero {
  background-color: #e6dbc5;
  padding: 50px;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.hero h1 {
  margin: 0;
  color: #6d4c41;
  font-size: 2.5em;
}

.hero p {
  font-size: 1.2em;
  color: #5a3e2b;
}

.buttons {
  margin-top: 20px;
}

.btn {
  display: inline-block;
  margin: 10px;
  padding: 12px 20px;
  background-color: #9c7e69;
  color: white;
  text-decoration: none;
  font-weight: bold;
  border-radius: 5px;
  transition: background 0.3s;
}

.btn:hover {
  background-color: #6d4c41;
}

.presentation {
  margin-top: 50px;
}

.presentation h2 {
  color: #6d4c41;
  font-size: 2em;
}

.items-grid {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 20px;
  margin-top: 20px;
}

.item {
  text-align: center;
}

.item img {
  border-radius: 10px;
  width: 150px;
  height: 150px;
  object-fit: cover;
  border: 3px solid #9c7e69;
}

.item p {
  margin-top: 10px;
  font-size: 1.1em;
  color: #5a3e2b;
}
</style>
