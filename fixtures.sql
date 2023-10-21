-- Création de la table "utilisateurs"
CREATE TABLE utilisateurs (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(50) NOT NULL,
    prenom VARCHAR(50),
    age INT
);

-- Création de la table "commandes"
CREATE TABLE commandes (
    id SERIAL PRIMARY KEY,
    utilisateur_id INT REFERENCES utilisateurs(id),
    date_commande DATE,
    montant DECIMAL(10, 2)
);

-- Insertion de données de test dans la table "utilisateurs"
INSERT INTO utilisateurs (nom, prenom, age) VALUES
    ('Doe', 'John', 30),
    ('Smith', 'Alice', 25),
    ('Johnson', 'Bob', 35);

-- Insertion de données de test dans la table "commandes"
INSERT INTO commandes (utilisateur_id, date_commande, montant) VALUES
    (1, '2023-10-01', 100.00),
    (2, '2023-10-02', 75.50),
    (1, '2023-10-03', 50.25);

-- Vous pouvez ajouter d'autres données de test si nécessaire
