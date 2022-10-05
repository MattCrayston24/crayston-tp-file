package main

import (
    "fmt"
    "log"
    "os"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

func main() {
    content, err := os.ReadFile("Hangman.txt")
	// si err différent de nil on montre l'erreur a l'utilisateur
    if err != nil {
        log.Fatal(err)
    }
	// on utilise la fonction seed pour avoir
	// un aléatoire différent à chaque compilation
	// en lui donnant le temps actuel en seconde
	// depuis le 1er janvier 1970
	rand.Seed(time.Now().Unix())
    fmt.Println(enigme(content))
	fmt.Println(rand.Int())
}

func enigme(s []byte) string{
	var frags string 
	var tab []string
	var mot string
	// Split, ajout de mot dans un tableau de string
	for _, i:= range s{
		if string(i) != "\n"{
			mot += string(i)
		}else{
			// Enlève tout les espaces, retour à la ligne, tabulation etc...
			mot = strings.TrimSpace(mot) 
			tab = append(tab, mot)
			mot = ""
		}

	}
	// on récupère le premier et dernier élément du fichier
	frags += tab[0] + " "
	frags += tab[len(tab)-1] + " "

	// initialisation des index des fragments recherchés
	a := 0
	b := 0
	// dans la boucle on vérifie quand on arrive à before
	// dans la liste pour accéder à l'élément d'après
	// pour avoir le résultat du 3ème fragment
	// et on vérifie quand on arrive à now
	// dans la liste pour prendre l'élément d'avant
	// pour avoir le résultat du 4ème fragment
	for i:=0 ; i < len(tab) - 1; i++{
		if tab[i] == "before"{
			a, _ = strconv.Atoi(tab[i+1])
		}
		if tab[i] == "now"{
			c := tab[i-1]
			b = int(c[1]) / len(tab)
		}
	}
	frags += tab[a-1] + " "
	frags += tab[b-1] + " "


	return frags
}
