# Hexgrid

This is a Go library used to handle regular hexagons.

It's based on the algorithms described by [Amit Patel in his wonderful guide to
hexagons](http://www.redblobgames.com/grids/hexagons/implementation.html) -- I
highly recommend reading through that page.

## Usage
#### Importing

```go
import "github.com/seanhagen/hexgrid"
```

### Examples

#### Creating hexagons

```go
hexagonA := hexgrid.NewHex(1, 2) //at axial coordinates Q=1 R=2
hexagonB := hexgrid.NewHex(2, 3) //at axial coordinates Q=2 R=3
```

#### Measuring the distance (in hexagons) between one hexagon and another

```go
distance := hexagonA.DistanceTo(hexagonB)
```

#### Getting the array of hexagons on the path between one hexagon and another

```go
origin := hexgrid.NewHex(10,20)
destination := hexgrid.NewHex(30,40)
path := origin.LineDraw(destination) 
```


#### Creating a layout

```go
origin := hexgrid.Point{X: 0, Y: 0}     // The coordinate that corresponds to the center of hexagon 0,0
size := hexgrid.Point{X: 100, Y: 100}   // The length of an hexagon side => 100
layout: = hexgrid.Layout{Origin: origin, Size: size, Orientation: hexagon.OrientationFlat}
```

#### Obtaining the pixel that corresponds to a given hexagon

```go
hex := hexgrid.NewHex(1,0)             
pixel := hexgrid.HexToPixel(layout,hex)  // Pixel that corresponds to the center of hex 1,0 (in the given layout)
```


#### Obtaining the hexagon that contains the given pixel (and rounding it)

```go
point := hexgrid.Point{X: 10, Y: 20}
hex := PixelToHex(layout, point).Round()
```

## History

0.2: Combining multiple forks
0.1: First version

## Credits

* [Pedro Sousa](https://github.com/pmcxs), for the intial repo.
* [Igor Shmulyan](https://github.com/ishmulyan), [Brendan Le
  Glaunec](https://github.com/Ullaakut), [Sergey
  Kolunov](https://github.com/Metaur),
  [Puzzlemaker1](https://github.com/Puzzlemaker1), and [Patrick
  Fay](https://github.com/tigger0jk) for their forks.
* And of course [Amit Patel of Red Blob
  Games](http://www.redblobgames.com/grids/hexagons/implementation.html),
  without which we wouldn't have any of this.

## License

MIT
