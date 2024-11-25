package main

import (
	"log"
	"strings"
	"time"

	"github.com/marcusljx/replayer"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	sampleList := strings.Fields("Lorem ipsum dolor sit amet, consectetur adipiscing elit. " +
		"In euismod mi condimentum tortor feugiat, posuere placerat felis porta. " +
		"Sed pellentesque neque dui, non molestie ante porta ut. " +
		"Duis imperdiet libero leo, ac consequat lectus fermentum non. " +
		"Suspendisse pellentesque feugiat posuere. Proin id lectus lorem. Morbi scelerisque feugiat erat quis tempus. " +
		"Sed convallis ipsum turpis, ut tempor felis vulputate vel. " +
		"Mauris accumsan enim a sem semper, ut posuere leo porttitor. Quisque mollis vulputate rutrum. " +
		"Pellentesque lacinia enim et tortor semper, vitae tempor massa iaculis. " +
		"Nullam odio nisl, semper ut finibus id, viverra vel orci.\n\n")

	config := &replayer.Configuration[string]{
		Source: &replayer.ListSource[string]{
			List:  sampleList,
			Index: 0,
		},
		GetTimestamp: func(i int, s string) time.Time {
			return time.Unix(int64(i), 0)
		},
		BufferSize: 10,
		Speed:      5,
	}

	player, err := config.Compile()
	if err != nil {
		panic(err)
	}

	player.Play(func(s string) error {
		log.Printf("out:%s", s)
		return nil
	})
}
