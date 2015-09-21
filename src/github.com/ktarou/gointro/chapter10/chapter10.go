//Write your own Sleep function using time.After.
package main
import (
	"time"
	"math/rand"
	"fmt"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	p1 := Character{"player1", 100, 20.0}
	p2 := Character{"player2", 100, 20.0}
	c1 := make(chan Character)
	c2 := make(chan Character)
	go play(&p1, &p2, c1, r)
	go play(&p2, &p1, c2, r)
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Player 2 attacked", msg1)
			case msg2 := <-c2:
				fmt.Println("Player 1 attacked", msg2)
			case <- time.After(time.Second * 3):
				fmt.Println("Game Over")
				return
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}

type Character struct {
	Name string
	Hp   int
	Atk  float64
}

func ( c *Character) attack(c1 *Character, rand *rand.Rand) float64 {
	magnitude := float64(rand.Intn(100))/100;
	atk := c.Atk * magnitude
	c1.Hp -= int(atk)
	return magnitude
}

func play(c1 *Character, c2 *Character, ch chan Character, rand *rand.Rand) {
	for c1.Hp > 0 && c2.Hp > 0 {
		m := c1.attack(c2, rand)
		ch <- *c2
		time.Sleep(time.Millisecond * time.Duration(2000 * (1 - m)))
	}
	if (c1.Hp>0) {
		fmt.Println(c1.Name, " is alive")
		return
	}else{
		fmt.Println(c1.Name, " is dead")
		return
	}
}

