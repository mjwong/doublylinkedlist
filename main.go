package main

import (
  "log"
	"fmt"
	"time"
  "errors"
	"github.com/goombaio/namegenerator"
	"gopkg.in/gookit/color.v1"
)

type song struct {
	name string
	artist string
  prev *song
	next *song
}

type playList struct {
	name string
	head *song
  last *song
	nowPlaying *song
}

func (p *playList) addSong(name, artist string) error {
	s := &song{
		name: name,
		artist: artist,
	}

	if p.head == nil {
		p.head = s
    s.prev = nil
	} else {
		currNode := p.head
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = s
    s.prev = currNode
    p.last = s
	}
	return nil
}

func (p *playList) countNodes() int {
	count := 0
	currNode := p.head
	for currNode != nil {
		currNode = currNode.next
		count += 1
	}
	return count
}

func (p *playList) getSong(index int) (*song, error) {
  if index < 0 || index > p.countNodes() {
    return nil, errors.New("Index out of range.")
  }

  i := 0
  currNode := p.head
  for currNode != nil {
    if i == index {
      return currNode, nil
    }
    i += 1
    currNode = currNode.next
  }
  return nil, nil
}

func (p *playList) delNode(node *song) {
  i := 0
  currNode := p.head

  for currNode != nil {
    if currNode.name == node.name {
      node.prev.next = node.next
      node.next.prev = node.prev
      break
    }
    currNode = currNode.next
    i += 1
  }
}

func (p *playList) displaySongs() error {
	currNode := p.head
	if currNode == nil {
		color.Yellow.Println("Playlist is empty.")
	} else {
		for currNode != nil {
			fmt.Printf("name: %25s. Addr: %v\n", color.Cyan.Render(currNode.name), color.Green.Render(currNode))
			currNode = currNode.next
		}
	}
	return nil
}

func (p *playList) displaySongsRev() error {
  currNode := p.last
	if currNode == nil {
		color.Yellow.Println("Playlist is empty.")
	} else {
		for currNode != nil {
			fmt.Printf("name: %25s. Curr: %v. Prev: %v\n", color.Cyan.Render(currNode.name), color.Green.Render(currNode), color.Green.Render(currNode.prev))
			currNode = currNode.prev
		}
	}
	return nil
}

func main() {
	p := &playList{}

	for i := 1; i <= 10; i++ {

		seed := time.Now().UTC().UnixNano()
		nameGenerator := namegenerator.NewNameGenerator(seed)

		name := nameGenerator.Generate()

		p.addSong(name, name)
	}

	p.displaySongs()

  fmt.Printf("\nDisplay in rev order\n\n")

  p.displaySongsRev()

  fmt.Printf("\nDelete no.5\n\n")

  node, err := p.getSong(5)

  fmt.Printf("Song: %s\n", node.name)
  
  if err != nil {
    fmt.Printf("Error finding song.\n")
    log.Fatal(err)
  }

  p.delNode(node)

  p.displaySongs()
  fmt.Printf("\nCount = %d\n", p.countNodes())
}
