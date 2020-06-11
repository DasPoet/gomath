# gomath

<div
align="center">
     <a
    href ="https://goreportcard.com/report/github.com/daspoet/gomath">
    <img
        height="30" src="https://goreportcard.com/badge/github.com/daspoet/gomath?style=for-the-badge"/>
    </a>
    <a
    href="https://pkg.go.dev/github.com/daspoet/gomath">
    <img
        height="30" src="https://img.shields.io/badge/Made%20with-Go-%2300fff2?style=for-the-badge"/>
    </a>
    <img
        height="30" src="https://img.shields.io/badge/Contains-Math-%23fc0398?style=for-the-badge" />
</div>

## Installation

It is assumed that you have already worked with the Go evironment. If this is not the case, please see this [page first](https://golang.org/doc/install).

`go get github.com/daspoet/gomath`

## Functional use

### Imports

```golang
import github.com/daspoet/gomath
```

### [Matrices](https://pkg.go.dev/github.com/daspoet/gomath#Matrix)

#### Instantiating matrices

Notes:

* a new matrix is - by default - filled with 0s

Constructing a new matrix

```golang
gomath.NewMatrix(m, n)
```

<div align="center"
>
<img
    src="https://latex.codecogs.com/svg.latex?\begin%20{bmatrix}i_{0,0}%20&%20...%20&%20i_{0,n}%20\\...%20&%20...%20&%20...%20\\i_{m,0}%20&%20...%20&%20i_{m,n}\end%20{bmatrix}"/>
</div>

Constructing a new matrix that is already filled

```golang
gomath.NewFilledMatrix(m, n, element)
```

<div align="center"
>
<img
    src="https://latex.codecogs.com/svg.latex?\left.\begin%20{bmatrix}i_{0,0}%20&%20...%20&%20i_{0,n}%20\\...%20&%20...%20&%20...%20\\i_{m,0}%20&%20...%20&%20i_{m,n}\end%20{bmatrix}\right\}i_{a,b}%20=%20element%20\;%20for%20\;%200%20\,%20\leq%20\,%20a%20\,%20\leq%20\,%20m%20\;%20and%20\;%200%20\,%20\leq%20b%20\,%20\leq%20\,%20n"/>
</div>

Constructing a new square matrix

```golang
gomath.NewSquareMatrix(n, n)
```

<div align="center"
>
<img
    src="https://latex.codecogs.com/svg.latex?\begin%20{bmatrix}i_{0,0}%20&%20...%20&%20i_{0,n}%20\\%20%20%20%20...%20&%20...%20&%20%20...%20%20%20%20\\i_{n,0}%20&%20...%20&%20i_{n,n}\end%20{bmatrix}"/>
</div>

Constructing a new square matrix that is already filled

```golang
gomath.NewFilledSquareMatrix(n, element)
```

<div align="center"
>
<img
    src="https://latex.codecogs.com/svg.latex?\left.\begin%20{bmatrix}i_{0,0}%20&%20...%20&%20i_{0,n}%20\\%20%20%20...%20%20&%20...%20&%20%20...%20%20%20%20\\i_{n,0}%20&%20...%20&%20i_{n,n}\end%20{bmatrix}\right\}i_{a,b}%20=%20element%20\;%20for%20\;%200%20\,%20\leq%20\,%20a%20\,%20\leq%20\,%20n%20\;%20and%20\;%200%20\,%20\leq%20\,%20b%20\,%20\leq%20\,%20n"/>
</div>

Constructing a new identity matrix

```golang
gomath.NewIdentityMatrix(n)
```

<div align="center"
>
<img
    src="https://latex.codecogs.com/svg.latex?n%20\left\{\begin%20{bmatrix}1%20&%200%20&%200%20&...%20&%20\\0%20&%201%20&%200%20&%20...%20&%20\\0%20&%200%20&%201%20&%20...%20&%20\\0%20&%200%20&%20...%20&%201\end%20{bmatrix}\right."/>
</div>

### Vectors

Notes:

* a new vector is - by default - filled with 0s
* a vector containing n elements is simply an n x 1 matrix

Constructing a new vector

```golang
gomath.NewVector(n)
```

<div align="center"
>
<img
    src="https://latex.codecogs.com/svg.latex?\begin%20{bmatrix}i_0%20\\...%20\\i_n\end%20{bmatrix}"/>
</div>

Constructing a new vector that is already filled

```golang
gomath.NewFilledVector(n, element)
```

<div align="center"
>
<img
    src="https://latex.codecogs.com/svg.latex?\left.\begin%20{bmatrix}i_0%20\\...%20\\i_n\end%20{bmatrix}\right\}i_0%20=%20...%20=%20i_n%20=%20element"/>
</div>

### Linear systems

Matrices can also be used to find solutions to systems of linear equations.
