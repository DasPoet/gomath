<<<<<<< HEAD
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

* a new matrix is - by default - filled with $0$s

Constructing a new matrix

```golang
gomath.NewMatrix(m, n)
```

$$
\begin {bmatrix}

i_{0,0} & ... & i_{0,n} \\
   ...  & ... &  ...    \\
i_{m,0} & ... & i_{m,n}

\end {bmatrix}
$$

<div align="center"
>
<img
    src="https://latex.codecogs.com/svg.latex?\color{white} \begin {bmatrix}i_{0,0} & ... & i_{0,n} \\...  & ... &  ...    \\i_{m,0} & ... & i_{m,n}\end {bmatrix}"/>
</div>

Constructing a new square matrix

```golang
gomath.NewSquareMatrix(n, n)
```

$$
\begin {bmatrix}

i_{0,0} & ... & i_{0,n} \\
    ... & ... &  ...    \\
i_{n,0} & ... & i_{n,n}

\end {bmatrix}
$$

Constructing a new matrix that is already filled

```golang
gomath.NewFilledMatrix(m, n, element)
```

$$
\left.
\begin {bmatrix}

i_{0,0} & ... & i_{0,n} \\
   ...  & ... &  ...    \\
i_{m,0} & ... & i_{m,n}

\end {bmatrix}
\right\}

i_{a,b} = element \; for \; 0 \, \leq \, a \, \leq \, m \; and \; 0 \, \leq \, b \, \leq \, n
$$

Constructing a new square matrix that is already filled

```golang
gomath.NewFilledSquareMatrix(n, element)
```

$$
\left.
\begin {bmatrix}

i_{0,0} & ... & i_{0,n} \\
   ...  & ... &  ...    \\
i_{n,0} & ... & i_{n,n}

\end {bmatrix}
\right\}

i_{a,b} = element \; for \; 0 \, \leq \, a \, \leq \, n \; and \; 0 \, \leq \, b \, \leq \, n
$$

Constructing a new identity matrix

```golang
gomath.NewIdentityMatrix(n)
```

$$
n \left\{
\begin {bmatrix}

1 & 0 & 0 &... & \\
0 & 1 & 0 & ... & \\
0 & 0 & 1 & ... & \\
0 & 0 & ... & 1

\end {bmatrix}
\right.
$$

### Vectors

Notes:

* a new vector is - by default - filled with $0$s
* a vector containing $n$ elements is simply an $n$ x $1$ matrix

Constructing a new vector

```golang
gomath.NewVector(n)
```

$$
\begin {bmatrix}

i_0 \\
... \\
i_n

\end {bmatrix}
$$

Constructing a new vector that is already filled

```golang
gomath.NewFilledVector()
```

$$
\left.
\begin {bmatrix}

i_0 \\
... \\
i_n

\end {bmatrix}
\right\}

i_0 = ... = i_n = element
$$

### Linear systems

Matrices can also be used to find solutions to systems of linear equations.
