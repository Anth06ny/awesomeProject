<template>
  <div id="app" class="app-container">
    <header class="header">
      <h1>📚 Librairie Ludique 🎲</h1>
      <nav>
        <router-link to="/">Accueil</router-link>
        <router-link to="/livres">Livres</router-link>
        <router-link to="/jeux">Jeux de plateau</router-link>
        <router-link to="/login">Connexion</router-link>
      </nav>
    </header>

    <main>
      <router-view></router-view>
    </main>

    <footer class="footer">
      <p>&copy; 2025 Librairie Ludique - Tous droits réservés.</p>
    </footer>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
/* Style global avec thème beige */
body {
  margin: 0;
  font-family: 'Arial', sans-serif;
  background-color: #f5f2e7;
  color: #4a4a4a;
}

.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.header {
  background-color: #d8c3a5;
  padding: 1rem;
  text-align: center;
}

.header h1 {
  margin: 0;
  color: #6d4c41;
}

nav {
  margin-top: 10px;
}

nav a {
  margin: 0 15px;
  text-decoration: none;
  color: #6d4c41;
  font-weight: bold;
}

nav a:hover {
  color: #9c7e69;
}

main {
  flex-grow: 1;
  padding: 20px;
}

.footer {
  background-color: #d8c3a5;
  text-align: center;
  padding: 10px;
  font-size: 14px;
  margin-top: auto;
}
</style>