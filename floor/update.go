package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	switch configuration.Global.FloorKind {
	case gridFloor:
		f.updateGridFloor(camXPos, camYPos)
	case fromFileFloor:
		f.updateFromFileFloor(camXPos, camYPos)
	case quadTreeFloor:
		f.updateQuadtreeFloor(camXPos, camYPos)
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(camXPos, camYPos int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absCamX := camXPos
			if absCamX < 0 {
				absCamX = -absCamX
			}
			absCamY := camYPos
			if absCamY < 0 {
				absCamY = -absCamY
			}
			f.content[y][x] = ((x + absCamX%2) + (y + absCamY%2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
func (f *Floor) updateFromFileFloor(camXPos, camYPos int) {
	// Calculez les coordonnées de la zone visible à l'écran.
	screenLeft := camXPos - configuration.Global.NumTileX/2
	screenRight := camXPos + configuration.Global.NumTileX/2
	screenTop := camYPos - configuration.Global.NumTileY/2
	screenBottom := camYPos + configuration.Global.NumTileY/2

	// Assurez-vous que les coordonnées de la zone visible restent dans les limites du terrain.
	// Vous pouvez définir des bornes minimales et maximales pour éviter de sortir des limites du terrain.
	if screenLeft < 0 {
		screenLeft = 0
	}
	if screenRight >= len(f.fullContent[0]) {
		screenRight = len(f.fullContent[0]) - 1
	}
	if screenTop < 0 {
		screenTop = 0
	}
	if screenBottom >= len(f.fullContent) {
		screenBottom = len(f.fullContent) - 1
	}

	// Remplissez le tableau content avec les valeurs du terrain chargé depuis le fichier.
	for i := 0; i < len(f.content); i++ {
		f.content[i] = make([]int, 0, len(f.content[0]))
		for j := screenLeft; j <= screenRight; j++ {
			if i >= screenTop && i <= screenBottom {
				if j >= 0 && j < len(f.fullContent[0]) {
					// Copiez la valeur du terrain depuis fullContent vers content.
					f.content[i] = append(f.content[i], f.fullContent[i][j])
				} else {
					// Coordonnées en dehors des limites du terrain. Utilisez -1 pour indiquer l'absence de terrain.
					f.content[i] = append(f.content[i], -1)
				}
			} else {
				// Coordonnées en dehors des limites du terrain. Utilisez -1 pour indiquer l'absence de terrain.
				f.content[i] = append(f.content[i], -1)
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
