package main

import (
	"groupie-tracker/requests"
	"math/rand"
)

func main() {

	Publickeys := []string{"ac98b57a03d4bc1b378ceae13aa1f470", "d9cbfed943ee91f7ef5f2bd9e93a5454", "3792dc6a0228717eef4fa462cfd63b25", "9bae353c2e1b4535f47d285120a3410b", "773fd5baab801c65c3195195f4a53579"}
	Privatekeys := []string{"dd535c1c275ff098a12e33b13dabafa14c489da0", "1947e7e5beef88b2f8c214b70a6c458abe89960d", "c93fe39aace8f4d323732ad14982bb3aad2d93d1", "a42e0ca2df4addb93ae8e743d36b82275eb125b1", "eb987ca2ca8d622c76b3a71e3c13114f7a4bc8fc"}

	randomts := rand.Intn(5432)
	h := requests.Hash(randomts, Publickeys, Privatekeys)
	println(h)
}
