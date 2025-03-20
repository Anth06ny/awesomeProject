<template>
  <div class="jeux-container">
    <h1>🎲 Liste des Jeux de Plateau Disponibles</h1>

    <div v-if="loading" class="loading">Chargement des jeux...</div>
    <div v-else-if="error" class="error">{{ error }}</div>

    <div v-else class="jeux-grid">
      <div v-for="jeu in jeuxFiltres" :key="jeu.ID" class="jeu-card" :class="{ emprunte: jeu.Status === 'indisponible' }">
        <h2>{{ jeu.Title }}</h2>
        <p><strong>Statut :</strong> <span :class="{'disponible': jeu.Status === 'disponible', 'emprunté': jeu.Status === 'indisponible'}">{{ jeu.Status }}</span></p>
        <button v-if="jeu.Status === 'disponible'" @click="emprunterJeu(jeu.ID)">🎮 Emprunter</button>
        <button v-else @click="rendreJeu(jeu.ID)">🔄 Rendre</button>
      </div>
    </div>
  </div>
</template>

<script>
import apiClient from "@/api.js"; // Utilise le client centralisé

export default {
  name: "Jeux",
  data() {
    return {
      jeux: [],
      loading: true,
      error: null
    };
  },
  computed: {
    // Ne garde que les livres (exclut les jeux de plateau)
    jeuxFiltres() {
      return this.jeux.filter(jeu => jeu.Type === "Jeu");
    }
  },
  methods: {
    async fetchJeux() {
      try {
        const response = await apiClient.get("/resources"); // Plus besoin d’écrire l’URL complète // Assure-toi que l'API backend est bien à cette adresse
        this.jeux = response.data;
      } catch (err) {
        this.error = "Impossible de charger les jeux.";
      } finally {
        this.loading = false;
      }
    },
    async emprunterJeu(id) {
      try {
        await apiClient.put(`/resources/${id}/disable`); // Route pour emprunter un livre
        await this.fetchJeux(); // Rafraîchir la liste après emprunt
      } catch (err) {
        this.error = "Erreur lors de l'emprunt du jeu.";
      }
    } ,
    async rendreJeu(id) {
      try {
        await apiClient.put(`/resources/${id}/enable`); // Route pour rendre un livre
        await this.fetchJeux(); // Rafraîchir la liste après emprunt
      } catch (err) {
        this.error = "Erreur lors de l'emprunt du jeu.";
      }
    }
  },
  mounted() {
    this.fetchJeux();
  }
};
</script>

<style scoped>
.jeux-container {
  text-align: center;
  padding: 20px;
}

.loading, .error {
  font-size: 1.2em;
  color: #6d4c41;
}

.jeux-grid {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 20px;
  margin-top: 20px;
}

.jeu-card {
  background: #f5f2e7;
  border-radius: 10px;
  padding: 20px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  width: 250px;
  text-align: center;
}

.jeu-card h2 {
  margin-bottom: 10px;
  color: #6d4c41;
}

.jeu-card p {
  margin: 10px 0;
}

.jeu-card button {
  background-color: #9c7e69;
  color: white;
  border: none;
  padding: 10px;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s;
}

.jeu-card button:hover {
  background-color: #6d4c41;
}

.jeu-card button:disabled {
  background-color: #bbb;
  cursor: not-allowed;
}

.emprunte {
  opacity: 0.6;
}

.disponible {
  color: green;
}

.emprunté {
  color: red;
}
</style>
