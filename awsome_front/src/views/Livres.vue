<template>
  <div class="livres-container">
    <h1>📚 Liste des Livres Disponibles</h1>

    <div v-if="loading" class="loading">Chargement des livres...</div>
    <div v-else-if="error" class="error">{{ error }}</div>

    <div v-else class="livres-grid">
      <div v-for="livre in livresFiltres" :key="livre.ID" :class="{ emprunte: livre.Status === 'indisponible' }"
           class="livre-card">
        <h2>{{ livre.Title }}</h2>
        <p><strong>Statut :</strong> <span
            :class="{'disponible': livre.Status === 'disponible', 'indisponible': livre.Status === 'indisponible'}">
          {{ livre.Status }}
        </span></p>

        <button v-if="livre.Status === 'disponible'" @click="emprunterLivre(livre.ID)">📖 Emprunter</button>
        <button v-else @click="rendreLivre(livre.ID)">🔄 Rendre</button>
      </div>
    </div>

  </div>
</template>

<script>
import apiClient from "@/api.js"; // Utilise le client centralisé

export default {
  name: "Livres",
  data() {
    return {
      livres: [],
      loading: true,
      error: null
    };
  },
  computed: {
    // Ne garde que les livres (exclut les jeux de plateau)
    livresFiltres() {
      return this.livres.filter(livre => livre.Type === "Livre");
    }
  },
  methods: {
    async fetchLivres() {
      try {
        const response = await apiClient.get("/resources"); // Récupère tous les livres et jeux
        this.livres = response.data;
      } catch (err) {
        this.error = "Impossible de charger les livres.";
      } finally {
        this.loading = false;
      }
    },
    async emprunterLivre(id) {
      try {
        await apiClient.put(`/resources/${id}/disable`); // Route pour emprunter un livre
        await this.fetchLivres(); // Rafraîchir la liste après emprunt
      } catch (err) {
        this.error = "Erreur lors de l'emprunt du livre.";
      }
    },
    async rendreLivre(id) {
      try {
        await apiClient.put(`/resources/${id}/enable`); // Route pour rendre un livre
        await this.fetchLivres(); // Rafraîchir la liste après retour
      } catch (err) {
        this.error = "Erreur lors du retour du livre.";
      }
    }
  },
  mounted() {
    this.fetchLivres();
  }
};
</script>

<style scoped>
.livres-container {
  text-align: center;
  padding: 20px;
}

.loading, .error {
  font-size: 1.2em;
  color: #6d4c41;
}

.livres-grid {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 20px;
  margin-top: 20px;
}

.livre-card {
  background: #f5f2e7;
  border-radius: 10px;
  padding: 20px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  width: 250px;
  text-align: center;
}

.livre-card h2 {
  margin-bottom: 10px;
  color: #6d4c41;
}

.livre-card p {
  margin: 10px 0;
}

.livre-card button {
  background-color: #9c7e69;
  color: white;
  border: none;
  padding: 10px;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s;
}

.livre-card button:hover {
  background-color: #6d4c41;
}

.livre-card button:disabled {
  background-color: #bbb;
  cursor: not-allowed;
}

.emprunte {
  opacity: 0.6;
}

.disponible {
  color: green;
}

.indisponible {
  color: red;
}
</style>
