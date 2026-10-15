// Flyweight
// Flyweight is a structural design pattern that allows you to share objects to support a large number
// of similar objects efficiently.
// It is used to minimize memory usage by sharing common parts of state between multiple objects.

package structural

import "fmt"

// Flyweight
// The flyweight type below will be used to represent the shared state.
type Texture []byte

// Context
// The context type below will be used to represent the unique state of each object.
// In this case, it represents the position of the sprite in the game world.
// The context type contains a reference to the flyweight type.
type Sprite struct {
	X, Y    int
	Texture *Texture
}

// Flyweight Factory
// The factory creates and manages the flyweight objects.
// It is responsible for creating the flyweight objects and storing them in a map.
// In this case, we will only use a variable with registered textures.
var Textures = map[string]*Texture{
	"tree":  {0x01, 0x02, 0x03},
	"stone": {0x05, 0x06, 0x07},
}

// Test Flyweight
// The test function creates two sprites with the same texture.
// Note that the texture is shared between the two sprites.
// This means that the texture is not duplicated in memory.
func TestFlyweight() {
	treeTexture := Textures["tree"]
	tree1 := &Sprite{X: 0, Y: 0, Texture: treeTexture}
	tree2 := &Sprite{X: 1, Y: 1, Texture: treeTexture}
	// Check that the texture address is the same
	fmt.Println(tree1, tree2)                   // Output: &{0 0 0xc00000c080} &{1 1 0xc00000c080}
	fmt.Println(tree1.Texture == tree2.Texture) // Output: true
}
